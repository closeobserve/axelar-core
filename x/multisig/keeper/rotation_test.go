package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/assert"

	"github.com/axelarnetwork/axelar-core/app"
	"github.com/axelarnetwork/axelar-core/testutils/fake"
	testutilsrand "github.com/axelarnetwork/axelar-core/testutils/rand"
	"github.com/axelarnetwork/axelar-core/x/multisig/exported"
	exportedtestutils "github.com/axelarnetwork/axelar-core/x/multisig/exported/testutils"
	"github.com/axelarnetwork/axelar-core/x/multisig/keeper"
	"github.com/axelarnetwork/axelar-core/x/multisig/types"
	typestestutils "github.com/axelarnetwork/axelar-core/x/multisig/types/testutils"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	. "github.com/axelarnetwork/utils/test"
)

func TestKeeper(t *testing.T) {
	encCfg := app.MakeEncodingConfig()
	chainName := nexus.ChainName(testutilsrand.NormalizedStr(5))

	var (
		k      keeper.Keeper
		ctx    sdk.Context
		keyID1 exported.KeyID
		keyID2 exported.KeyID
	)

	givenKeeper := When("multisig keeper", func() {
		subspace := params.NewSubspace(encCfg.Codec, encCfg.Amino, sdk.NewKVStoreKey("paramsKey"), sdk.NewKVStoreKey("tparamsKey"), "multisig")
		k = keeper.NewKeeper(encCfg.Codec, sdk.NewKVStoreKey(types.StoreKey), subspace)
		ctx = testutilsrand.Context(fake.NewMultiStore())

		k.InitGenesis(ctx, types.DefaultGenesisState())
	})

	whenKeysExist := When("keys exist", func() {
		key := typestestutils.Key()
		keyID1 = key.GetID()

		k.SetKey(ctx, key)

		key = typestestutils.Key()
		keyID2 = key.GetID()

		k.SetKey(ctx, key)
	})

	t.Run("AssignKey", func(t *testing.T) {
		givenKeeper.
			Branch(
				Then("should fail if key does not exist", func(t *testing.T) {
					err := k.AssignKey(ctx, chainName, exportedtestutils.KeyID())
					assert.Error(t, err)
				}),

				whenKeysExist.
					Branch(
						When("a key is assigned", func() {
							k.AssignKey(ctx, chainName, keyID1)
						}).
							Then("should fail if assign the same key again", func(t *testing.T) {
								err := k.AssignKey(ctx, chainName, keyID1)
								assert.Error(t, err)
							}),

						When("a key is assigned", func() {
							k.AssignKey(ctx, chainName, keyID1)
						}).
							Then("should fail if assign another key", func(t *testing.T) {
								err := k.AssignKey(ctx, chainName, keyID2)
								assert.Error(t, err)
							}),

						Then("should succeed", func(t *testing.T) {
							err := k.AssignKey(ctx, chainName, keyID1)
							assert.NoError(t, err)

							actual, ok := k.GetNextKeyID(ctx, chainName)
							assert.True(t, ok)
							assert.Equal(t, keyID1, actual)

							_, ok = k.GetCurrentKeyID(ctx, chainName)
							assert.False(t, ok)
						}),
					),
			).
			Run(t)
	})

	t.Run("RotateKey", func(t *testing.T) {
		givenKeeper.
			When2(whenKeysExist).
			Branch(
				Then("should fail if no key is assigned", func(t *testing.T) {
					err := k.RotateKey(ctx, chainName)
					assert.Error(t, err)
				}),

				When("some key is assigned", func() {
					k.AssignKey(ctx, chainName, keyID1)
				}).
					Then("should succeed", func(t *testing.T) {
						err := k.RotateKey(ctx, chainName)
						assert.NoError(t, err)

						_, ok := k.GetNextKeyID(ctx, chainName)
						assert.False(t, ok)

						actual, ok := k.GetCurrentKeyID(ctx, chainName)
						assert.True(t, ok)
						assert.Equal(t, keyID1, actual)

						currentKey, ok := k.GetCurrentKey(ctx, chainName)
						assert.True(t, ok)
						assert.Equal(t, keyID1, currentKey.ID)
						assert.Equal(t, types.Active, currentKey.State)
					}),
			).
			Run(t)
	})
}
