package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/btcsuite/btcd/rpcclient"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/rpc/client/http"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/grpc"

	tssd "github.com/axelarnetwork/tssd/pb"

	"github.com/axelarnetwork/axelar-core/x/balance"

	keyring "github.com/cosmos/cosmos-sdk/crypto/keys"

	"github.com/axelarnetwork/axelar-core/store"
	balanceKeeper "github.com/axelarnetwork/axelar-core/x/balance/keeper"
	balanceTypes "github.com/axelarnetwork/axelar-core/x/balance/types"
	"github.com/axelarnetwork/axelar-core/x/bitcoin"
	btcKeeper "github.com/axelarnetwork/axelar-core/x/bitcoin/keeper"
	btcTypes "github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	"github.com/axelarnetwork/axelar-core/x/broadcast"
	broadcastKeeper "github.com/axelarnetwork/axelar-core/x/broadcast/keeper"
	broadcastTypes "github.com/axelarnetwork/axelar-core/x/broadcast/types"
	"github.com/axelarnetwork/axelar-core/x/ethereum"
	ethKeeper "github.com/axelarnetwork/axelar-core/x/ethereum/keeper"
	ethTypes "github.com/axelarnetwork/axelar-core/x/ethereum/types"
	"github.com/axelarnetwork/axelar-core/x/snapshot"
	snapKeeper "github.com/axelarnetwork/axelar-core/x/snapshot/keeper"
	snapTypes "github.com/axelarnetwork/axelar-core/x/snapshot/types"
	"github.com/axelarnetwork/axelar-core/x/tss"
	tssKeeper "github.com/axelarnetwork/axelar-core/x/tss/keeper"
	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"
	"github.com/axelarnetwork/axelar-core/x/vote"
	voteKeeper "github.com/axelarnetwork/axelar-core/x/vote/keeper"
	voteTypes "github.com/axelarnetwork/axelar-core/x/vote/types"
)

const (
	appName = "axelar"
)

var (
	// DefaultCLIHome sets the default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.axelarcli")

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.axelard")

	// ModuleBasics is in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},

		tss.AppModuleBasic{},
		vote.AppModuleBasic{},
		bitcoin.AppModuleBasic{},
		ethereum.AppModuleBasic{},
		broadcast.AppModuleBasic{},
		snapshot.AppModuleBasic{},
		balance.AppModuleBasic{},
	)
	// account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		distr.ModuleName:          nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
	}
)

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	cdc = cdc.Seal()

	return cdc
}

// AxelarApp defines the axelar Cosmos app that runs all modules
type AxelarApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// Keepers
	accountKeeper   auth.AccountKeeper
	bankKeeper      bank.Keeper
	stakingKeeper   staking.Keeper
	slashingKeeper  slashing.Keeper
	distrKeeper     distr.Keeper
	supplyKeeper    supply.Keeper
	paramsKeeper    params.Keeper
	btcKeeper       btcKeeper.Keeper
	ethKeeper       ethKeeper.Keeper
	broadcastKeeper broadcastKeeper.Keeper
	tssKeeper       tssKeeper.Keeper
	votingKeeper    voteKeeper.Keeper
	snapKeeper      snapKeeper.Keeper
	balanceKeeper   balanceKeeper.Keeper

	// Module Manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager
}

// verify app interface at compile time
var _ simapp.App = &AxelarApp{}

// NewInitApp is a constructor function for axelarApp
func NewInitApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, axelarCfg Config, baseAppOptions ...func(*bam.BaseApp)) *AxelarApp {

	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey,
		auth.StoreKey,
		staking.StoreKey,
		supply.StoreKey,
		distr.StoreKey,
		slashing.StoreKey,
		params.StoreKey,
		voteTypes.StoreKey,
		broadcastTypes.StoreKey,
		btcTypes.StoreKey,
		ethTypes.StoreKey,
		snapTypes.StoreKey,
		tssTypes.StoreKey,
		balanceTypes.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)

	// Here you initialize your application with the store keys it requires
	var app = &AxelarApp{
		BaseApp:        bApp,
		cdc:            cdc,
		keys:           keys,
		tkeys:          tkeys,
		invCheckPeriod: invCheckPeriod,
	}

	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tkeys[params.TStoreKey])

	// Set specific subspaces
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	tssSubspace := app.paramsKeeper.Subspace(tssTypes.DefaultParamspace)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		keys[auth.StoreKey],
		authSubspace,
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		bankSubspace,
		app.ModuleAccountAddrs(),
	)

	// The SupplyKeeper collects transaction fees and renders them to the fee distribution module
	app.supplyKeeper = supply.NewKeeper(
		app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		maccPerms,
	)

	// The staking keeper
	stakingKeeper := staking.NewKeeper(
		app.cdc,
		keys[staking.StoreKey],
		app.supplyKeeper,
		stakingSubspace,
	)

	app.distrKeeper = distr.NewKeeper(
		app.cdc,
		keys[distr.StoreKey],
		distrSubspace,
		&stakingKeeper,
		app.supplyKeeper,
		auth.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		keys[slashing.StoreKey],
		&stakingKeeper,
		slashingSubspace,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distrKeeper.Hooks(),
			app.slashingKeeper.Hooks()),
	)

	var err error
	app.btcKeeper = btcKeeper.NewBtcKeeper(app.cdc, keys[btcTypes.StoreKey])

	app.ethKeeper = ethKeeper.NewEthKeeper(app.cdc, keys[ethTypes.StoreKey])

	app.snapKeeper = snapKeeper.NewKeeper(app.cdc, keys[snapTypes.StoreKey], app.stakingKeeper)

	keybase, err := keyring.NewKeyring(sdk.KeyringServiceName(), axelarCfg.ClientConfig.KeyringBackend, DefaultCLIHome, os.Stdin)
	if err != nil {
		tmos.Exit(err.Error())
	}
	abciClient, err := http.New(axelarCfg.TendermintNodeUri, "/websocket")
	if err != nil {
		tmos.Exit(err.Error())
	}
	app.broadcastKeeper, err = broadcastKeeper.NewKeeper(
		app.cdc,
		keys[broadcastTypes.StoreKey],
		store.NewSubjectiveStore(),
		keybase,
		app.accountKeeper,
		app.snapKeeper,
		abciClient,
		axelarCfg.ClientConfig,
		logger,
	)
	if err != nil {
		tmos.Exit(err.Error())
	}

	// TODO don't start gRPC unless I'm a validator?
	// start a gRPC client
	tssdServerAddress := axelarCfg.Host + ":" + axelarCfg.Port
	logger.Info(fmt.Sprintf("initiate connection to tssd gRPC server: %s", tssdServerAddress))
	conn, err := grpc.Dial(tssdServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		tmos.Exit(err.Error())
	}
	logger.Debug("successful connection to tssd gRPC server")

	client := tssd.NewGG18Client(conn)
	app.tssKeeper = tssKeeper.NewKeeper(app.cdc, keys[tssTypes.StoreKey], client, tssSubspace, app.broadcastKeeper)

	// Clean up tss grpc connection on process shutdown
	tmos.TrapSignal(logger, func() {
		logger.Debug("initiate Close")
		if err := conn.Close(); err != nil {
			logger.Error(sdkerrors.Wrap(err, "failure to close connection to server").Error())
			return
		}
		logger.Debug("successful Close")
	})

	app.votingKeeper = voteKeeper.NewKeeper(app.cdc, keys[voteTypes.StoreKey], store.NewSubjectiveStore(), app.snapKeeper, app.broadcastKeeper)

	// TODO: enable running node without an Ethereum bridge
	rpcETC, err := ethTypes.NewRPCClient(axelarCfg.EthRpcAddr)
	if err != nil {
		tmos.Exit(err.Error())
	}
	logger.Debug("Successfully connected to ethereum node")

	// Enable running a node with or without a Bitcoin bridge
	var rpcBTC *rpcclient.Client
	var btcModule bitcoin.AppModule
	if axelarCfg.WithBtcBridge {
		rpcBTC, err = btcTypes.NewRPCClient(axelarCfg.BtcConfig, logger)
		if err != nil {
			tmos.Exit(err.Error())
		}
		// BTC bridge opens a grpc connection. Clean it up on process shutdown
		tmos.TrapSignal(logger, rpcBTC.Shutdown)
		btcModule = bitcoin.NewAppModule(app.btcKeeper, app.votingKeeper, app.tssKeeper, rpcBTC)
	} else {
		btcModule = bitcoin.NewDummyAppModule(app.btcKeeper, app.votingKeeper)
	}

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distr.NewAppModule(app.distrKeeper, app.accountKeeper, app.supplyKeeper, app.stakingKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),

		snapshot.NewAppModule(app.snapKeeper),
		tss.NewAppModule(app.tssKeeper, app.snapKeeper, app.votingKeeper),
		vote.NewAppModule(app.votingKeeper),
		broadcast.NewAppModule(app.broadcastKeeper),
		ethereum.NewAppModule(app.ethKeeper, app.votingKeeper, app.tssKeeper, app.balanceKeeper, rpcETC),
		balance.NewAppModule(app.balanceKeeper),
		btcModule,
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(distr.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(staking.ModuleName, voteTypes.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	// NOTE: The genutils moodule must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		distr.ModuleName,
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		tssTypes.ModuleName,
		btcTypes.ModuleName,
		ethTypes.ModuleName,
		broadcastTypes.ModuleName,
		voteTypes.ModuleName,
		supply.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// InitChainer handles the chain initialization from a genesis file
func (app *AxelarApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	err := app.cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, genesisState)
}

// BeginBlocker calls the BeginBlock() function of every module at the beginning of a new block
func (app *AxelarApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker calls the EndBlock() function of every module at the end of a block
func (app *AxelarApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// LoadHeight loads the application version at a given height. It will panic if called
// more than once on a running baseapp.
func (app *AxelarApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// Codec returns simapp's codec
func (app *AxelarApp) Codec() *codec.Codec {
	return app.cdc
}

// SimulationManager implements the SimulationApp interface
func (app *AxelarApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *AxelarApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}
