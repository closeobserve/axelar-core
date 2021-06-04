// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/tss/rpc"
	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	"google.golang.org/grpc"
	"sync"
)

// Ensure, that ClientMock does implement rpc.Client.
// If this is not the case, regenerate this file with moq.
var _ rpc.Client = &ClientMock{}

// ClientMock is a mock implementation of rpc.Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked rpc.Client
// 		mockedClient := &ClientMock{
// 			KeygenFunc: func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_KeygenClient, error) {
// 				panic("mock out the Keygen method")
// 			},
// 			SignFunc: func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_SignClient, error) {
// 				panic("mock out the Sign method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires rpc.Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// KeygenFunc mocks the Keygen method.
	KeygenFunc func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_KeygenClient, error)

	// SignFunc mocks the Sign method.
	SignFunc func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_SignClient, error)

	// calls tracks calls to the methods.
	calls struct {
		// Keygen holds details about calls to the Keygen method.
		Keygen []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Sign holds details about calls to the Sign method.
		Sign []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
	}
	lockKeygen sync.RWMutex
	lockSign   sync.RWMutex
}

// Keygen calls KeygenFunc.
func (mock *ClientMock) Keygen(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_KeygenClient, error) {
	if mock.KeygenFunc == nil {
		panic("ClientMock.KeygenFunc: method is nil but Client.Keygen was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockKeygen.Lock()
	mock.calls.Keygen = append(mock.calls.Keygen, callInfo)
	mock.lockKeygen.Unlock()
	return mock.KeygenFunc(ctx, opts...)
}

// KeygenCalls gets all the calls that were made to Keygen.
// Check the length with:
//     len(mockedClient.KeygenCalls())
func (mock *ClientMock) KeygenCalls() []struct {
	Ctx  context.Context
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}
	mock.lockKeygen.RLock()
	calls = mock.calls.Keygen
	mock.lockKeygen.RUnlock()
	return calls
}

// Sign calls SignFunc.
func (mock *ClientMock) Sign(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_SignClient, error) {
	if mock.SignFunc == nil {
		panic("ClientMock.SignFunc: method is nil but Client.Sign was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockSign.Lock()
	mock.calls.Sign = append(mock.calls.Sign, callInfo)
	mock.lockSign.Unlock()
	return mock.SignFunc(ctx, opts...)
}

// SignCalls gets all the calls that were made to Sign.
// Check the length with:
//     len(mockedClient.SignCalls())
func (mock *ClientMock) SignCalls() []struct {
	Ctx  context.Context
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}
	mock.lockSign.RLock()
	calls = mock.calls.Sign
	mock.lockSign.RUnlock()
	return calls
}
