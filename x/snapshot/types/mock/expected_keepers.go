// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/snapshot/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkExported "github.com/cosmos/cosmos-sdk/x/staking/exported"
	"sync"
)

// Ensure, that StakingKeeperMock does implement types.StakingKeeper.
// If this is not the case, regenerate this file with moq.
var _ types.StakingKeeper = &StakingKeeperMock{}

// StakingKeeperMock is a mock implementation of types.StakingKeeper.
//
//     func TestSomethingThatUsesStakingKeeper(t *testing.T) {
//
//         // make and configure a mocked types.StakingKeeper
//         mockedStakingKeeper := &StakingKeeperMock{
//             GetLastTotalPowerFunc: func(ctx sdk.Context) sdk.Int {
// 	               panic("mock out the GetLastTotalPower method")
//             },
//             IterateLastValidatorsFunc: func(ctx sdk.Context, fn func(index int64, validator sdkExported.ValidatorI) (stop bool))  {
// 	               panic("mock out the IterateLastValidators method")
//             },
//             ValidatorFunc: func(ctx sdk.Context, addr sdk.ValAddress) sdkExported.ValidatorI {
// 	               panic("mock out the Validator method")
//             },
//         }
//
//         // use mockedStakingKeeper in code that requires types.StakingKeeper
//         // and then make assertions.
//
//     }
type StakingKeeperMock struct {
	// GetLastTotalPowerFunc mocks the GetLastTotalPower method.
	GetLastTotalPowerFunc func(ctx sdk.Context) sdk.Int

	// IterateLastValidatorsFunc mocks the IterateLastValidators method.
	IterateLastValidatorsFunc func(ctx sdk.Context, fn func(index int64, validator sdkExported.ValidatorI) (stop bool))

	// ValidatorFunc mocks the Validator method.
	ValidatorFunc func(ctx sdk.Context, addr sdk.ValAddress) sdkExported.ValidatorI

	// calls tracks calls to the methods.
	calls struct {
		// GetLastTotalPower holds details about calls to the GetLastTotalPower method.
		GetLastTotalPower []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
		// IterateLastValidators holds details about calls to the IterateLastValidators method.
		IterateLastValidators []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Fn is the fn argument value.
			Fn func(index int64, validator sdkExported.ValidatorI) (stop bool)
		}
		// Validator holds details about calls to the Validator method.
		Validator []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Addr is the addr argument value.
			Addr sdk.ValAddress
		}
	}
	lockGetLastTotalPower     sync.RWMutex
	lockIterateLastValidators sync.RWMutex
	lockValidator             sync.RWMutex
}

// GetLastTotalPower calls GetLastTotalPowerFunc.
func (mock *StakingKeeperMock) GetLastTotalPower(ctx sdk.Context) sdk.Int {
	if mock.GetLastTotalPowerFunc == nil {
		panic("StakingKeeperMock.GetLastTotalPowerFunc: method is nil but StakingKeeper.GetLastTotalPower was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLastTotalPower.Lock()
	mock.calls.GetLastTotalPower = append(mock.calls.GetLastTotalPower, callInfo)
	mock.lockGetLastTotalPower.Unlock()
	return mock.GetLastTotalPowerFunc(ctx)
}

// GetLastTotalPowerCalls gets all the calls that were made to GetLastTotalPower.
// Check the length with:
//     len(mockedStakingKeeper.GetLastTotalPowerCalls())
func (mock *StakingKeeperMock) GetLastTotalPowerCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetLastTotalPower.RLock()
	calls = mock.calls.GetLastTotalPower
	mock.lockGetLastTotalPower.RUnlock()
	return calls
}

// IterateLastValidators calls IterateLastValidatorsFunc.
func (mock *StakingKeeperMock) IterateLastValidators(ctx sdk.Context, fn func(index int64, validator sdkExported.ValidatorI) (stop bool)) {
	if mock.IterateLastValidatorsFunc == nil {
		panic("StakingKeeperMock.IterateLastValidatorsFunc: method is nil but StakingKeeper.IterateLastValidators was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
		Fn  func(index int64, validator sdkExported.ValidatorI) (stop bool)
	}{
		Ctx: ctx,
		Fn:  fn,
	}
	mock.lockIterateLastValidators.Lock()
	mock.calls.IterateLastValidators = append(mock.calls.IterateLastValidators, callInfo)
	mock.lockIterateLastValidators.Unlock()
	mock.IterateLastValidatorsFunc(ctx, fn)
}

// IterateLastValidatorsCalls gets all the calls that were made to IterateLastValidators.
// Check the length with:
//     len(mockedStakingKeeper.IterateLastValidatorsCalls())
func (mock *StakingKeeperMock) IterateLastValidatorsCalls() []struct {
	Ctx sdk.Context
	Fn  func(index int64, validator sdkExported.ValidatorI) (stop bool)
} {
	var calls []struct {
		Ctx sdk.Context
		Fn  func(index int64, validator sdkExported.ValidatorI) (stop bool)
	}
	mock.lockIterateLastValidators.RLock()
	calls = mock.calls.IterateLastValidators
	mock.lockIterateLastValidators.RUnlock()
	return calls
}

// Validator calls ValidatorFunc.
func (mock *StakingKeeperMock) Validator(ctx sdk.Context, addr sdk.ValAddress) sdkExported.ValidatorI {
	if mock.ValidatorFunc == nil {
		panic("StakingKeeperMock.ValidatorFunc: method is nil but StakingKeeper.Validator was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Addr sdk.ValAddress
	}{
		Ctx:  ctx,
		Addr: addr,
	}
	mock.lockValidator.Lock()
	mock.calls.Validator = append(mock.calls.Validator, callInfo)
	mock.lockValidator.Unlock()
	return mock.ValidatorFunc(ctx, addr)
}

// ValidatorCalls gets all the calls that were made to Validator.
// Check the length with:
//     len(mockedStakingKeeper.ValidatorCalls())
func (mock *StakingKeeperMock) ValidatorCalls() []struct {
	Ctx  sdk.Context
	Addr sdk.ValAddress
} {
	var calls []struct {
		Ctx  sdk.Context
		Addr sdk.ValAddress
	}
	mock.lockValidator.RLock()
	calls = mock.calls.Validator
	mock.lockValidator.RUnlock()
	return calls
}
