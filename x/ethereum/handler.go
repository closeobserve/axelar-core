package ethereum

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	balance "github.com/axelarnetwork/axelar-core/x/balance/exported"
	"github.com/axelarnetwork/axelar-core/x/ethereum/keeper"
	"github.com/axelarnetwork/axelar-core/x/ethereum/types"
	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	"github.com/axelarnetwork/axelar-core/x/vote/exported"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	crypto "github.com/ethereum/go-ethereum/crypto"
)

func NewHandler(k keeper.Keeper, rpc types.RPCClient, v types.Voter, s types.Signer, snap snapshot.Snapshotter, balancer types.Balancer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgVerifyTx:
			return handleMsgVerifyTx(ctx, k, rpc, v, msg)
		case *types.MsgVoteVerifiedTx:
			return handleMsgVoteVerifiedTx(ctx, k, v, msg)
		case types.MsgSignTx:
			return handleMsgSignTx(ctx, k, s, snap, msg)
		case types.MsgSignPendingTransfersTx:
			return handleMsgSignPendingTransfersTx(ctx, k, s, snap, balancer, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg))
		}
	}
}

func handleMsgSignPendingTransfersTx(ctx sdk.Context, k keeper.Keeper, signer types.Signer, snap snapshot.Snapshotter, balancer types.Balancer, msg types.MsgSignPendingTransfersTx) (*sdk.Result, error) {
	pendingTransfers := balancer.GetPendingTransfersForChain(ctx, balance.Ethereum)
	addresses, denoms, amounts := flatTransfers(pendingTransfers)
	chainID := k.GetParams(ctx).Network.Params().ChainID
	commandID := calculateCommandID(pendingTransfers)

	data, err := types.CreateExecuteMintData(chainID, commandID, addresses, denoms, amounts)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrEthereum, err.Error())
	}

	keyID, ok := signer.GetCurrentMasterKeyID(ctx, balance.Ethereum)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrEthereum, "no master key for chain %s found", balance.Ethereum)
	}

	snapshot, ok := snap.GetLatestSnapshot(ctx)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrEthereum, "no snapshot found")
	}

	commandIDHex := hex.EncodeToString(commandID[:])
	k.Logger(ctx).Info(fmt.Sprintf("storing data for mint command %s", commandIDHex))
	k.SetCommandData(ctx, commandID, data)

	k.Logger(ctx).Info(fmt.Sprintf("signing mint command [%s] for pending transfers to chain %s", commandIDHex, balance.Ethereum))
	signHash := types.GetEthereumSignHash(data)

	err = signer.StartSign(ctx, keyID, commandIDHex, signHash.Bytes(), snapshot.Validators)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrEthereum, err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeModule),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(types.AttributeCommandID, commandIDHex),
		),
	)

	return &sdk.Result{
		Data:   signHash.Bytes(),
		Log:    fmt.Sprintf("successfully started signing protocol for chain %s pending transfers", balance.Ethereum),
		Events: ctx.EventManager().Events(),
	}, nil
}

// TODO: Remove this function after https://github.com/axelarnetwork/ethereum-bridge/issues/3 is implemented
func calculateCommandID(transfers []balance.CrossChainTransfer) [32]byte {
	var result [32]byte
	var hash []byte

	for _, transfer := range transfers {
		idByte := make([]byte, 8)
		binary.LittleEndian.PutUint64(idByte, transfer.ID)

		hash = crypto.Keccak256(hash, idByte)
	}

	copy(result[:], hash[:32])

	return result
}

func flatTransfers(transfers []balance.CrossChainTransfer) ([]string, []string, []*big.Int) {
	var addresses []string
	var denoms []string
	var amounts []*big.Int

	for _, transfer := range transfers {
		addresses = append(addresses, transfer.Recipient.Address)
		denoms = append(denoms, transfer.Amount.Denom)
		amounts = append(amounts, transfer.Amount.Amount.BigInt())
	}

	return addresses, denoms, amounts
}

func handleMsgVerifyTx(ctx sdk.Context, k keeper.Keeper, rpc types.RPCClient, v types.Voter, msg types.MsgVerifyTx) (*sdk.Result, error) {
	k.Logger(ctx).Debug("verifying ethereum transaction")
	tx := msg.UnmarshaledTx()
	txID := tx.Hash().String()

	poll := exported.PollMeta{Module: types.ModuleName, Type: msg.Type(), ID: txID}
	if err := v.InitPoll(ctx, poll); err != nil {
		return nil, sdkerrors.Wrap(types.ErrEthereum, err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeModule),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(types.AttributeTxID, txID),
		),
	)

	k.SetUnverifiedTx(ctx, txID, tx)

	/*
	 Anyone not able to verify the transaction will automatically record a negative vote,
	 but only validators will later send out that vote.
	*/

	if err := verifyTx(ctx, k, rpc, tx); err != nil {
		k.Logger(ctx).Debug(sdkerrors.Wrapf(err, "expected transaction (%s) could not be verified", txID).Error())
		if err := v.RecordVote(ctx, &types.MsgVoteVerifiedTx{PollMeta: poll, VotingData: false}); err != nil {
			k.Logger(ctx).Error(sdkerrors.Wrap(err, "voting failed").Error())
			return &sdk.Result{
				Log:    err.Error(),
				Data:   k.Codec().MustMarshalBinaryLengthPrefixed(false),
				Events: ctx.EventManager().Events(),
			}, nil
		}
		return &sdk.Result{
			Log:    err.Error(),
			Data:   k.Codec().MustMarshalBinaryLengthPrefixed(false),
			Events: ctx.EventManager().Events(),
		}, nil
	} else {
		if err := v.RecordVote(ctx, &types.MsgVoteVerifiedTx{PollMeta: poll, VotingData: true}); err != nil {
			k.Logger(ctx).Error(sdkerrors.Wrap(err, "voting failed").Error())
			return &sdk.Result{
				Log:    err.Error(),
				Data:   k.Codec().MustMarshalBinaryLengthPrefixed(false),
				Events: ctx.EventManager().Events(),
			}, nil
		}
		return &sdk.Result{
			Log:    "successfully verified transaction",
			Data:   k.Codec().MustMarshalBinaryLengthPrefixed(true),
			Events: ctx.EventManager().Events(),
		}, nil
	}

}

// This can be used as a potential hook to immediately act on a poll being decided by the vote
func handleMsgVoteVerifiedTx(ctx sdk.Context, k keeper.Keeper, v types.Voter, msg *types.MsgVoteVerifiedTx) (*sdk.Result, error) {
	event := sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeModule),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		sdk.NewAttribute(types.AttributePoll, msg.PollMeta.String()),
		sdk.NewAttribute(types.AttributeVotingData, strconv.FormatBool(msg.VotingData)),
	)

	if err := v.TallyVote(ctx, msg); err != nil {
		return nil, err
	}

	if confirmed := v.Result(ctx, msg.Poll()); confirmed != nil {
		if err := k.ProcessVerificationResult(ctx, msg.PollMeta.ID, confirmed.(bool)); err != nil {

			return nil, err
		}

		v.DeletePoll(ctx, msg.Poll())

		event = event.AppendAttributes(sdk.NewAttribute(types.AttributePollConfirmed, strconv.FormatBool(confirmed.(bool))))
	}

	ctx.EventManager().EmitEvent(event)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSignTx(ctx sdk.Context, k keeper.Keeper, signer types.Signer, snap snapshot.Snapshotter, msg types.MsgSignTx) (*sdk.Result, error) {
	tx := msg.UnmarshaledTx()
	txID := tx.Hash().String()
	k.SetRawTx(ctx, txID, tx)
	k.Logger(ctx).Info(fmt.Sprintf("storing raw tx %s", txID))
	hash, err := k.GetHashToSign(ctx, txID)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrEthereum, err.Error())
	}

	k.Logger(ctx).Info(fmt.Sprintf("ethereum tx [%s] to sign: %s", txID, k.Codec().MustMarshalJSON(hash)))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeModule),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(types.AttributeTxID, txID),
		),
	)

	keyID, ok := signer.GetCurrentMasterKeyID(ctx, balance.Ethereum)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrEthereum, "no master key for chain %s found", balance.Ethereum)
	}

	s, ok := snap.GetLatestSnapshot(ctx)
	if !ok {
		return nil, sdkerrors.Wrap(types.ErrEthereum, "no snapshot found")
	}
	err = signer.StartSign(ctx, keyID, hash.String(), hash.Bytes(), s.Validators)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrEthereum, err.Error())
	}

	return &sdk.Result{
		Data:   []byte(txID),
		Log:    fmt.Sprintf("successfully started signing protocol for transaction with ID %s.", txID),
		Events: ctx.EventManager().Events(),
	}, nil
}

func verifyTx(ctx sdk.Context, k keeper.Keeper, rpc types.RPCClient, expectedTx *ethTypes.Transaction) error {
	hash := expectedTx.Hash()
	receipt, err := rpc.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return sdkerrors.Wrap(err, "could not retrieve Ethereum receipt")
	}

	blockNumber, err := rpc.BlockNumber(context.Background())
	if err != nil {
		return sdkerrors.Wrap(err, "could not retrieve Ethereum block number")
	}

	if (blockNumber - receipt.BlockNumber.Uint64()) < k.GetRequiredConfirmationHeight(ctx) {
		return fmt.Errorf("not enough confirmations yet")
	}
	return nil
}
