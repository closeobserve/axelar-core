// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"crypto/ecdsa"
	"github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	voting "github.com/axelarnetwork/axelar-core/x/vote/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that VoterMock does implement types.Voter.
// If this is not the case, regenerate this file with moq.
var _ types.Voter = &VoterMock{}

// VoterMock is a mock implementation of types.Voter.
//
//     func TestSomethingThatUsesVoter(t *testing.T) {
//
//         // make and configure a mocked types.Voter
//         mockedVoter := &VoterMock{
//             InitPollFunc: func(ctx sdk.Context, poll voting.PollMeta) error {
// 	               panic("mock out the InitPoll method")
//             },
//             RecordVoteFunc: func(ctx sdk.Context, vote voting.MsgVote) error {
// 	               panic("mock out the RecordVote method")
//             },
//             ResultFunc: func(ctx sdk.Context, poll voting.PollMeta) voting.VotingData {
// 	               panic("mock out the Result method")
//             },
//             TallyVoteFunc: func(ctx sdk.Context, vote voting.MsgVote) error {
// 	               panic("mock out the TallyVote method")
//             },
//         }
//
//         // use mockedVoter in code that requires types.Voter
//         // and then make assertions.
//
//     }
type VoterMock struct {
	// InitPollFunc mocks the InitPoll method.
	InitPollFunc func(ctx sdk.Context, poll voting.PollMeta) error

	// RecordVoteFunc mocks the RecordVote method.
	RecordVoteFunc func(ctx sdk.Context, vote voting.MsgVote) error

	// ResultFunc mocks the Result method.
	ResultFunc func(ctx sdk.Context, poll voting.PollMeta) voting.VotingData

	// TallyVoteFunc mocks the TallyVote method.
	TallyVoteFunc func(ctx sdk.Context, vote voting.MsgVote) error

	// calls tracks calls to the methods.
	calls struct {
		// InitPoll holds details about calls to the InitPoll method.
		InitPoll []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Poll is the poll argument value.
			Poll voting.PollMeta
		}
		// RecordVote holds details about calls to the RecordVote method.
		RecordVote []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Vote is the vote argument value.
			Vote voting.MsgVote
		}
		// Result holds details about calls to the Result method.
		Result []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Poll is the poll argument value.
			Poll voting.PollMeta
		}
		// TallyVote holds details about calls to the TallyVote method.
		TallyVote []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Vote is the vote argument value.
			Vote voting.MsgVote
		}
	}
	lockInitPoll   sync.RWMutex
	lockRecordVote sync.RWMutex
	lockResult     sync.RWMutex
	lockTallyVote  sync.RWMutex
}

// InitPoll calls InitPollFunc.
func (mock *VoterMock) InitPoll(ctx sdk.Context, poll voting.PollMeta) error {
	if mock.InitPollFunc == nil {
		panic("VoterMock.InitPollFunc: method is nil but Voter.InitPoll was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockInitPoll.Lock()
	mock.calls.InitPoll = append(mock.calls.InitPoll, callInfo)
	mock.lockInitPoll.Unlock()
	return mock.InitPollFunc(ctx, poll)
}

// InitPollCalls gets all the calls that were made to InitPoll.
// Check the length with:
//     len(mockedVoter.InitPollCalls())
func (mock *VoterMock) InitPollCalls() []struct {
	Ctx  sdk.Context
	Poll voting.PollMeta
} {
	var calls []struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}
	mock.lockInitPoll.RLock()
	calls = mock.calls.InitPoll
	mock.lockInitPoll.RUnlock()
	return calls
}

// RecordVote calls RecordVoteFunc.
func (mock *VoterMock) RecordVote(ctx sdk.Context, vote voting.MsgVote) error {
	if mock.RecordVoteFunc == nil {
		panic("VoterMock.RecordVoteFunc: method is nil but Voter.RecordVote was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}{
		Ctx:  ctx,
		Vote: vote,
	}
	mock.lockRecordVote.Lock()
	mock.calls.RecordVote = append(mock.calls.RecordVote, callInfo)
	mock.lockRecordVote.Unlock()
	return mock.RecordVoteFunc(ctx, vote)
}

// RecordVoteCalls gets all the calls that were made to RecordVote.
// Check the length with:
//     len(mockedVoter.RecordVoteCalls())
func (mock *VoterMock) RecordVoteCalls() []struct {
	Ctx  sdk.Context
	Vote voting.MsgVote
} {
	var calls []struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}
	mock.lockRecordVote.RLock()
	calls = mock.calls.RecordVote
	mock.lockRecordVote.RUnlock()
	return calls
}

// Result calls ResultFunc.
func (mock *VoterMock) Result(ctx sdk.Context, poll voting.PollMeta) voting.VotingData {
	if mock.ResultFunc == nil {
		panic("VoterMock.ResultFunc: method is nil but Voter.Result was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockResult.Lock()
	mock.calls.Result = append(mock.calls.Result, callInfo)
	mock.lockResult.Unlock()
	return mock.ResultFunc(ctx, poll)
}

// ResultCalls gets all the calls that were made to Result.
// Check the length with:
//     len(mockedVoter.ResultCalls())
func (mock *VoterMock) ResultCalls() []struct {
	Ctx  sdk.Context
	Poll voting.PollMeta
} {
	var calls []struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}
	mock.lockResult.RLock()
	calls = mock.calls.Result
	mock.lockResult.RUnlock()
	return calls
}

// TallyVote calls TallyVoteFunc.
func (mock *VoterMock) TallyVote(ctx sdk.Context, vote voting.MsgVote) error {
	if mock.TallyVoteFunc == nil {
		panic("VoterMock.TallyVoteFunc: method is nil but Voter.TallyVote was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}{
		Ctx:  ctx,
		Vote: vote,
	}
	mock.lockTallyVote.Lock()
	mock.calls.TallyVote = append(mock.calls.TallyVote, callInfo)
	mock.lockTallyVote.Unlock()
	return mock.TallyVoteFunc(ctx, vote)
}

// TallyVoteCalls gets all the calls that were made to TallyVote.
// Check the length with:
//     len(mockedVoter.TallyVoteCalls())
func (mock *VoterMock) TallyVoteCalls() []struct {
	Ctx  sdk.Context
	Vote voting.MsgVote
} {
	var calls []struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}
	mock.lockTallyVote.RLock()
	calls = mock.calls.TallyVote
	mock.lockTallyVote.RUnlock()
	return calls
}

// Ensure, that SignerMock does implement types.Signer.
// If this is not the case, regenerate this file with moq.
var _ types.Signer = &SignerMock{}

// SignerMock is a mock implementation of types.Signer.
//
//     func TestSomethingThatUsesSigner(t *testing.T) {
//
//         // make and configure a mocked types.Signer
//         mockedSigner := &SignerMock{
//             GetCurrentMasterKeyFunc: func(ctx sdk.Context, chain string) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetCurrentMasterKey method")
//             },
//             GetKeyFunc: func(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetKey method")
//             },
//             GetKeyForSigIDFunc: func(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetKeyForSigID method")
//             },
//             GetNextMasterKeyFunc: func(ctx sdk.Context, chain string) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetNextMasterKey method")
//             },
//             GetSigFunc: func(ctx sdk.Context, sigID string) (tss.Signature, bool) {
// 	               panic("mock out the GetSig method")
//             },
//         }
//
//         // use mockedSigner in code that requires types.Signer
//         // and then make assertions.
//
//     }
type SignerMock struct {
	// GetCurrentMasterKeyFunc mocks the GetCurrentMasterKey method.
	GetCurrentMasterKeyFunc func(ctx sdk.Context, chain string) (ecdsa.PublicKey, bool)

	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool)

	// GetKeyForSigIDFunc mocks the GetKeyForSigID method.
	GetKeyForSigIDFunc func(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool)

	// GetNextMasterKeyFunc mocks the GetNextMasterKey method.
	GetNextMasterKeyFunc func(ctx sdk.Context, chain string) (ecdsa.PublicKey, bool)

	// GetSigFunc mocks the GetSig method.
	GetSigFunc func(ctx sdk.Context, sigID string) (tss.Signature, bool)

	// calls tracks calls to the methods.
	calls struct {
		// GetCurrentMasterKey holds details about calls to the GetCurrentMasterKey method.
		GetCurrentMasterKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain string
		}
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// KeyID is the keyID argument value.
			KeyID string
		}
		// GetKeyForSigID holds details about calls to the GetKeyForSigID method.
		GetKeyForSigID []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SigID is the sigID argument value.
			SigID string
		}
		// GetNextMasterKey holds details about calls to the GetNextMasterKey method.
		GetNextMasterKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain string
		}
		// GetSig holds details about calls to the GetSig method.
		GetSig []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SigID is the sigID argument value.
			SigID string
		}
	}
	lockGetCurrentMasterKey sync.RWMutex
	lockGetKey              sync.RWMutex
	lockGetKeyForSigID      sync.RWMutex
	lockGetNextMasterKey    sync.RWMutex
	lockGetSig              sync.RWMutex
}

// GetCurrentMasterKey calls GetCurrentMasterKeyFunc.
func (mock *SignerMock) GetCurrentMasterKey(ctx sdk.Context, chain string) (ecdsa.PublicKey, bool) {
	if mock.GetCurrentMasterKeyFunc == nil {
		panic("SignerMock.GetCurrentMasterKeyFunc: method is nil but Signer.GetCurrentMasterKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain string
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetCurrentMasterKey.Lock()
	mock.calls.GetCurrentMasterKey = append(mock.calls.GetCurrentMasterKey, callInfo)
	mock.lockGetCurrentMasterKey.Unlock()
	return mock.GetCurrentMasterKeyFunc(ctx, chain)
}

// GetCurrentMasterKeyCalls gets all the calls that were made to GetCurrentMasterKey.
// Check the length with:
//     len(mockedSigner.GetCurrentMasterKeyCalls())
func (mock *SignerMock) GetCurrentMasterKeyCalls() []struct {
	Ctx   sdk.Context
	Chain string
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain string
	}
	mock.lockGetCurrentMasterKey.RLock()
	calls = mock.calls.GetCurrentMasterKey
	mock.lockGetCurrentMasterKey.RUnlock()
	return calls
}

// GetKey calls GetKeyFunc.
func (mock *SignerMock) GetKey(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool) {
	if mock.GetKeyFunc == nil {
		panic("SignerMock.GetKeyFunc: method is nil but Signer.GetKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		KeyID string
	}{
		Ctx:   ctx,
		KeyID: keyID,
	}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	return mock.GetKeyFunc(ctx, keyID)
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//     len(mockedSigner.GetKeyCalls())
func (mock *SignerMock) GetKeyCalls() []struct {
	Ctx   sdk.Context
	KeyID string
} {
	var calls []struct {
		Ctx   sdk.Context
		KeyID string
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}

// GetKeyForSigID calls GetKeyForSigIDFunc.
func (mock *SignerMock) GetKeyForSigID(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool) {
	if mock.GetKeyForSigIDFunc == nil {
		panic("SignerMock.GetKeyForSigIDFunc: method is nil but Signer.GetKeyForSigID was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		SigID string
	}{
		Ctx:   ctx,
		SigID: sigID,
	}
	mock.lockGetKeyForSigID.Lock()
	mock.calls.GetKeyForSigID = append(mock.calls.GetKeyForSigID, callInfo)
	mock.lockGetKeyForSigID.Unlock()
	return mock.GetKeyForSigIDFunc(ctx, sigID)
}

// GetKeyForSigIDCalls gets all the calls that were made to GetKeyForSigID.
// Check the length with:
//     len(mockedSigner.GetKeyForSigIDCalls())
func (mock *SignerMock) GetKeyForSigIDCalls() []struct {
	Ctx   sdk.Context
	SigID string
} {
	var calls []struct {
		Ctx   sdk.Context
		SigID string
	}
	mock.lockGetKeyForSigID.RLock()
	calls = mock.calls.GetKeyForSigID
	mock.lockGetKeyForSigID.RUnlock()
	return calls
}

// GetNextMasterKey calls GetNextMasterKeyFunc.
func (mock *SignerMock) GetNextMasterKey(ctx sdk.Context, chain string) (ecdsa.PublicKey, bool) {
	if mock.GetNextMasterKeyFunc == nil {
		panic("SignerMock.GetNextMasterKeyFunc: method is nil but Signer.GetNextMasterKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain string
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetNextMasterKey.Lock()
	mock.calls.GetNextMasterKey = append(mock.calls.GetNextMasterKey, callInfo)
	mock.lockGetNextMasterKey.Unlock()
	return mock.GetNextMasterKeyFunc(ctx, chain)
}

// GetNextMasterKeyCalls gets all the calls that were made to GetNextMasterKey.
// Check the length with:
//     len(mockedSigner.GetNextMasterKeyCalls())
func (mock *SignerMock) GetNextMasterKeyCalls() []struct {
	Ctx   sdk.Context
	Chain string
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain string
	}
	mock.lockGetNextMasterKey.RLock()
	calls = mock.calls.GetNextMasterKey
	mock.lockGetNextMasterKey.RUnlock()
	return calls
}

// GetSig calls GetSigFunc.
func (mock *SignerMock) GetSig(ctx sdk.Context, sigID string) (tss.Signature, bool) {
	if mock.GetSigFunc == nil {
		panic("SignerMock.GetSigFunc: method is nil but Signer.GetSig was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		SigID string
	}{
		Ctx:   ctx,
		SigID: sigID,
	}
	mock.lockGetSig.Lock()
	mock.calls.GetSig = append(mock.calls.GetSig, callInfo)
	mock.lockGetSig.Unlock()
	return mock.GetSigFunc(ctx, sigID)
}

// GetSigCalls gets all the calls that were made to GetSig.
// Check the length with:
//     len(mockedSigner.GetSigCalls())
func (mock *SignerMock) GetSigCalls() []struct {
	Ctx   sdk.Context
	SigID string
} {
	var calls []struct {
		Ctx   sdk.Context
		SigID string
	}
	mock.lockGetSig.RLock()
	calls = mock.calls.GetSig
	mock.lockGetSig.RUnlock()
	return calls
}
