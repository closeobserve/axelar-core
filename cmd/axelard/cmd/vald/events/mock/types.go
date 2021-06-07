// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/axelarnetwork/tm-events/pkg/pubsub"
	"github.com/tendermint/tendermint/rpc/core/types"
	"sync"
)

// Ensure, that SignClientMock does implement SignClient.
// If this is not the case, regenerate this file with moq.
var _ SignClient = &SignClientMock{}

// SignClientMock is a mock implementation of SignClient.
//
// 	func TestSomethingThatUsesSignClient(t *testing.T) {
//
// 		// make and configure a mocked SignClient
// 		mockedSignClient := &SignClientMock{
// 			BlockFunc: func(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
// 				panic("mock out the Block method")
// 			},
// 			BlockByHashFunc: func(ctx context.Context, hash []byte) (*coretypes.ResultBlock, error) {
// 				panic("mock out the BlockByHash method")
// 			},
// 			BlockResultsFunc: func(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
// 				panic("mock out the BlockResults method")
// 			},
// 			BlockSearchFunc: func(ctx context.Context, query string, page *int, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error) {
// 				panic("mock out the BlockSearch method")
// 			},
// 			CommitFunc: func(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
// 				panic("mock out the Commit method")
// 			},
// 			TxFunc: func(ctx context.Context, hash []byte, prove bool) (*coretypes.ResultTx, error) {
// 				panic("mock out the Tx method")
// 			},
// 			TxSearchFunc: func(ctx context.Context, query string, prove bool, page *int, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
// 				panic("mock out the TxSearch method")
// 			},
// 			ValidatorsFunc: func(ctx context.Context, height *int64, page *int, perPage *int) (*coretypes.ResultValidators, error) {
// 				panic("mock out the Validators method")
// 			},
// 		}
//
// 		// use mockedSignClient in code that requires SignClient
// 		// and then make assertions.
//
// 	}
type SignClientMock struct {
	// BlockFunc mocks the Block method.
	BlockFunc func(ctx context.Context, height *int64) (*coretypes.ResultBlock, error)

	// BlockByHashFunc mocks the BlockByHash method.
	BlockByHashFunc func(ctx context.Context, hash []byte) (*coretypes.ResultBlock, error)

	// BlockResultsFunc mocks the BlockResults method.
	BlockResultsFunc func(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error)

	// BlockSearchFunc mocks the BlockSearch method.
	BlockSearchFunc func(ctx context.Context, query string, page *int, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error)

	// CommitFunc mocks the Commit method.
	CommitFunc func(ctx context.Context, height *int64) (*coretypes.ResultCommit, error)

	// TxFunc mocks the Tx method.
	TxFunc func(ctx context.Context, hash []byte, prove bool) (*coretypes.ResultTx, error)

	// TxSearchFunc mocks the TxSearch method.
	TxSearchFunc func(ctx context.Context, query string, prove bool, page *int, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error)

	// ValidatorsFunc mocks the Validators method.
	ValidatorsFunc func(ctx context.Context, height *int64, page *int, perPage *int) (*coretypes.ResultValidators, error)

	// calls tracks calls to the methods.
	calls struct {
		// Block holds details about calls to the Block method.
		Block []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Height is the height argument value.
			Height *int64
		}
		// BlockByHash holds details about calls to the BlockByHash method.
		BlockByHash []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Hash is the hash argument value.
			Hash []byte
		}
		// BlockResults holds details about calls to the BlockResults method.
		BlockResults []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Height is the height argument value.
			Height *int64
		}
		// BlockSearch holds details about calls to the BlockSearch method.
		BlockSearch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Page is the page argument value.
			Page *int
			// PerPage is the perPage argument value.
			PerPage *int
			// OrderBy is the orderBy argument value.
			OrderBy string
		}
		// Commit holds details about calls to the Commit method.
		Commit []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Height is the height argument value.
			Height *int64
		}
		// Tx holds details about calls to the Tx method.
		Tx []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Hash is the hash argument value.
			Hash []byte
			// Prove is the prove argument value.
			Prove bool
		}
		// TxSearch holds details about calls to the TxSearch method.
		TxSearch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Prove is the prove argument value.
			Prove bool
			// Page is the page argument value.
			Page *int
			// PerPage is the perPage argument value.
			PerPage *int
			// OrderBy is the orderBy argument value.
			OrderBy string
		}
		// Validators holds details about calls to the Validators method.
		Validators []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Height is the height argument value.
			Height *int64
			// Page is the page argument value.
			Page *int
			// PerPage is the perPage argument value.
			PerPage *int
		}
	}
	lockBlock        sync.RWMutex
	lockBlockByHash  sync.RWMutex
	lockBlockResults sync.RWMutex
	lockBlockSearch  sync.RWMutex
	lockCommit       sync.RWMutex
	lockTx           sync.RWMutex
	lockTxSearch     sync.RWMutex
	lockValidators   sync.RWMutex
}

// Block calls BlockFunc.
func (mock *SignClientMock) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	if mock.BlockFunc == nil {
		panic("SignClientMock.BlockFunc: method is nil but SignClient.Block was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Height *int64
	}{
		Ctx:    ctx,
		Height: height,
	}
	mock.lockBlock.Lock()
	mock.calls.Block = append(mock.calls.Block, callInfo)
	mock.lockBlock.Unlock()
	return mock.BlockFunc(ctx, height)
}

// BlockCalls gets all the calls that were made to Block.
// Check the length with:
//     len(mockedSignClient.BlockCalls())
func (mock *SignClientMock) BlockCalls() []struct {
	Ctx    context.Context
	Height *int64
} {
	var calls []struct {
		Ctx    context.Context
		Height *int64
	}
	mock.lockBlock.RLock()
	calls = mock.calls.Block
	mock.lockBlock.RUnlock()
	return calls
}

// BlockByHash calls BlockByHashFunc.
func (mock *SignClientMock) BlockByHash(ctx context.Context, hash []byte) (*coretypes.ResultBlock, error) {
	if mock.BlockByHashFunc == nil {
		panic("SignClientMock.BlockByHashFunc: method is nil but SignClient.BlockByHash was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Hash []byte
	}{
		Ctx:  ctx,
		Hash: hash,
	}
	mock.lockBlockByHash.Lock()
	mock.calls.BlockByHash = append(mock.calls.BlockByHash, callInfo)
	mock.lockBlockByHash.Unlock()
	return mock.BlockByHashFunc(ctx, hash)
}

// BlockByHashCalls gets all the calls that were made to BlockByHash.
// Check the length with:
//     len(mockedSignClient.BlockByHashCalls())
func (mock *SignClientMock) BlockByHashCalls() []struct {
	Ctx  context.Context
	Hash []byte
} {
	var calls []struct {
		Ctx  context.Context
		Hash []byte
	}
	mock.lockBlockByHash.RLock()
	calls = mock.calls.BlockByHash
	mock.lockBlockByHash.RUnlock()
	return calls
}

// BlockResults calls BlockResultsFunc.
func (mock *SignClientMock) BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	if mock.BlockResultsFunc == nil {
		panic("SignClientMock.BlockResultsFunc: method is nil but SignClient.BlockResults was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Height *int64
	}{
		Ctx:    ctx,
		Height: height,
	}
	mock.lockBlockResults.Lock()
	mock.calls.BlockResults = append(mock.calls.BlockResults, callInfo)
	mock.lockBlockResults.Unlock()
	return mock.BlockResultsFunc(ctx, height)
}

// BlockResultsCalls gets all the calls that were made to BlockResults.
// Check the length with:
//     len(mockedSignClient.BlockResultsCalls())
func (mock *SignClientMock) BlockResultsCalls() []struct {
	Ctx    context.Context
	Height *int64
} {
	var calls []struct {
		Ctx    context.Context
		Height *int64
	}
	mock.lockBlockResults.RLock()
	calls = mock.calls.BlockResults
	mock.lockBlockResults.RUnlock()
	return calls
}

// BlockSearch calls BlockSearchFunc.
func (mock *SignClientMock) BlockSearch(ctx context.Context, query string, page *int, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error) {
	if mock.BlockSearchFunc == nil {
		panic("SignClientMock.BlockSearchFunc: method is nil but SignClient.BlockSearch was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Query   string
		Page    *int
		PerPage *int
		OrderBy string
	}{
		Ctx:     ctx,
		Query:   query,
		Page:    page,
		PerPage: perPage,
		OrderBy: orderBy,
	}
	mock.lockBlockSearch.Lock()
	mock.calls.BlockSearch = append(mock.calls.BlockSearch, callInfo)
	mock.lockBlockSearch.Unlock()
	return mock.BlockSearchFunc(ctx, query, page, perPage, orderBy)
}

// BlockSearchCalls gets all the calls that were made to BlockSearch.
// Check the length with:
//     len(mockedSignClient.BlockSearchCalls())
func (mock *SignClientMock) BlockSearchCalls() []struct {
	Ctx     context.Context
	Query   string
	Page    *int
	PerPage *int
	OrderBy string
} {
	var calls []struct {
		Ctx     context.Context
		Query   string
		Page    *int
		PerPage *int
		OrderBy string
	}
	mock.lockBlockSearch.RLock()
	calls = mock.calls.BlockSearch
	mock.lockBlockSearch.RUnlock()
	return calls
}

// Commit calls CommitFunc.
func (mock *SignClientMock) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	if mock.CommitFunc == nil {
		panic("SignClientMock.CommitFunc: method is nil but SignClient.Commit was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Height *int64
	}{
		Ctx:    ctx,
		Height: height,
	}
	mock.lockCommit.Lock()
	mock.calls.Commit = append(mock.calls.Commit, callInfo)
	mock.lockCommit.Unlock()
	return mock.CommitFunc(ctx, height)
}

// CommitCalls gets all the calls that were made to Commit.
// Check the length with:
//     len(mockedSignClient.CommitCalls())
func (mock *SignClientMock) CommitCalls() []struct {
	Ctx    context.Context
	Height *int64
} {
	var calls []struct {
		Ctx    context.Context
		Height *int64
	}
	mock.lockCommit.RLock()
	calls = mock.calls.Commit
	mock.lockCommit.RUnlock()
	return calls
}

// Tx calls TxFunc.
func (mock *SignClientMock) Tx(ctx context.Context, hash []byte, prove bool) (*coretypes.ResultTx, error) {
	if mock.TxFunc == nil {
		panic("SignClientMock.TxFunc: method is nil but SignClient.Tx was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Hash  []byte
		Prove bool
	}{
		Ctx:   ctx,
		Hash:  hash,
		Prove: prove,
	}
	mock.lockTx.Lock()
	mock.calls.Tx = append(mock.calls.Tx, callInfo)
	mock.lockTx.Unlock()
	return mock.TxFunc(ctx, hash, prove)
}

// TxCalls gets all the calls that were made to Tx.
// Check the length with:
//     len(mockedSignClient.TxCalls())
func (mock *SignClientMock) TxCalls() []struct {
	Ctx   context.Context
	Hash  []byte
	Prove bool
} {
	var calls []struct {
		Ctx   context.Context
		Hash  []byte
		Prove bool
	}
	mock.lockTx.RLock()
	calls = mock.calls.Tx
	mock.lockTx.RUnlock()
	return calls
}

// TxSearch calls TxSearchFunc.
func (mock *SignClientMock) TxSearch(ctx context.Context, query string, prove bool, page *int, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	if mock.TxSearchFunc == nil {
		panic("SignClientMock.TxSearchFunc: method is nil but SignClient.TxSearch was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Query   string
		Prove   bool
		Page    *int
		PerPage *int
		OrderBy string
	}{
		Ctx:     ctx,
		Query:   query,
		Prove:   prove,
		Page:    page,
		PerPage: perPage,
		OrderBy: orderBy,
	}
	mock.lockTxSearch.Lock()
	mock.calls.TxSearch = append(mock.calls.TxSearch, callInfo)
	mock.lockTxSearch.Unlock()
	return mock.TxSearchFunc(ctx, query, prove, page, perPage, orderBy)
}

// TxSearchCalls gets all the calls that were made to TxSearch.
// Check the length with:
//     len(mockedSignClient.TxSearchCalls())
func (mock *SignClientMock) TxSearchCalls() []struct {
	Ctx     context.Context
	Query   string
	Prove   bool
	Page    *int
	PerPage *int
	OrderBy string
} {
	var calls []struct {
		Ctx     context.Context
		Query   string
		Prove   bool
		Page    *int
		PerPage *int
		OrderBy string
	}
	mock.lockTxSearch.RLock()
	calls = mock.calls.TxSearch
	mock.lockTxSearch.RUnlock()
	return calls
}

// Validators calls ValidatorsFunc.
func (mock *SignClientMock) Validators(ctx context.Context, height *int64, page *int, perPage *int) (*coretypes.ResultValidators, error) {
	if mock.ValidatorsFunc == nil {
		panic("SignClientMock.ValidatorsFunc: method is nil but SignClient.Validators was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Height  *int64
		Page    *int
		PerPage *int
	}{
		Ctx:     ctx,
		Height:  height,
		Page:    page,
		PerPage: perPage,
	}
	mock.lockValidators.Lock()
	mock.calls.Validators = append(mock.calls.Validators, callInfo)
	mock.lockValidators.Unlock()
	return mock.ValidatorsFunc(ctx, height, page, perPage)
}

// ValidatorsCalls gets all the calls that were made to Validators.
// Check the length with:
//     len(mockedSignClient.ValidatorsCalls())
func (mock *SignClientMock) ValidatorsCalls() []struct {
	Ctx     context.Context
	Height  *int64
	Page    *int
	PerPage *int
} {
	var calls []struct {
		Ctx     context.Context
		Height  *int64
		Page    *int
		PerPage *int
	}
	mock.lockValidators.RLock()
	calls = mock.calls.Validators
	mock.lockValidators.RUnlock()
	return calls
}

// Ensure, that QueryMock does implement Query.
// If this is not the case, regenerate this file with moq.
var _ Query = &QueryMock{}

// QueryMock is a mock implementation of Query.
//
// 	func TestSomethingThatUsesQuery(t *testing.T) {
//
// 		// make and configure a mocked Query
// 		mockedQuery := &QueryMock{
// 			MatchesFunc: func(events map[string][]string) (bool, error) {
// 				panic("mock out the Matches method")
// 			},
// 			StringFunc: func() string {
// 				panic("mock out the String method")
// 			},
// 		}
//
// 		// use mockedQuery in code that requires Query
// 		// and then make assertions.
//
// 	}
type QueryMock struct {
	// MatchesFunc mocks the Matches method.
	MatchesFunc func(events map[string][]string) (bool, error)

	// StringFunc mocks the String method.
	StringFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Matches holds details about calls to the Matches method.
		Matches []struct {
			// Events is the events argument value.
			Events map[string][]string
		}
		// String holds details about calls to the String method.
		String []struct {
		}
	}
	lockMatches sync.RWMutex
	lockString  sync.RWMutex
}

// Matches calls MatchesFunc.
func (mock *QueryMock) Matches(events map[string][]string) (bool, error) {
	if mock.MatchesFunc == nil {
		panic("QueryMock.MatchesFunc: method is nil but Query.Matches was just called")
	}
	callInfo := struct {
		Events map[string][]string
	}{
		Events: events,
	}
	mock.lockMatches.Lock()
	mock.calls.Matches = append(mock.calls.Matches, callInfo)
	mock.lockMatches.Unlock()
	return mock.MatchesFunc(events)
}

// MatchesCalls gets all the calls that were made to Matches.
// Check the length with:
//     len(mockedQuery.MatchesCalls())
func (mock *QueryMock) MatchesCalls() []struct {
	Events map[string][]string
} {
	var calls []struct {
		Events map[string][]string
	}
	mock.lockMatches.RLock()
	calls = mock.calls.Matches
	mock.lockMatches.RUnlock()
	return calls
}

// String calls StringFunc.
func (mock *QueryMock) String() string {
	if mock.StringFunc == nil {
		panic("QueryMock.StringFunc: method is nil but Query.String was just called")
	}
	callInfo := struct {
	}{}
	mock.lockString.Lock()
	mock.calls.String = append(mock.calls.String, callInfo)
	mock.lockString.Unlock()
	return mock.StringFunc()
}

// StringCalls gets all the calls that were made to String.
// Check the length with:
//     len(mockedQuery.StringCalls())
func (mock *QueryMock) StringCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockString.RLock()
	calls = mock.calls.String
	mock.lockString.RUnlock()
	return calls
}

// Ensure, that BusMock does implement Bus.
// If this is not the case, regenerate this file with moq.
var _ Bus = &BusMock{}

// BusMock is a mock implementation of Bus.
//
// 	func TestSomethingThatUsesBus(t *testing.T) {
//
// 		// make and configure a mocked Bus
// 		mockedBus := &BusMock{
// 			CloseFunc: func()  {
// 				panic("mock out the Close method")
// 			},
// 			DoneFunc: func() <-chan struct{} {
// 				panic("mock out the Done method")
// 			},
// 			PublishFunc: func(event pubsub.Event) error {
// 				panic("mock out the Publish method")
// 			},
// 			SubscribeFunc: func() (pubsub.Subscriber, error) {
// 				panic("mock out the Subscribe method")
// 			},
// 		}
//
// 		// use mockedBus in code that requires Bus
// 		// and then make assertions.
//
// 	}
type BusMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func()

	// DoneFunc mocks the Done method.
	DoneFunc func() <-chan struct{}

	// PublishFunc mocks the Publish method.
	PublishFunc func(event pubsub.Event) error

	// SubscribeFunc mocks the Subscribe method.
	SubscribeFunc func() (pubsub.Subscriber, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Done holds details about calls to the Done method.
		Done []struct {
		}
		// Publish holds details about calls to the Publish method.
		Publish []struct {
			// Event is the event argument value.
			Event pubsub.Event
		}
		// Subscribe holds details about calls to the Subscribe method.
		Subscribe []struct {
		}
	}
	lockClose     sync.RWMutex
	lockDone      sync.RWMutex
	lockPublish   sync.RWMutex
	lockSubscribe sync.RWMutex
}

// Close calls CloseFunc.
func (mock *BusMock) Close() {
	if mock.CloseFunc == nil {
		panic("BusMock.CloseFunc: method is nil but Bus.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedBus.CloseCalls())
func (mock *BusMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Done calls DoneFunc.
func (mock *BusMock) Done() <-chan struct{} {
	if mock.DoneFunc == nil {
		panic("BusMock.DoneFunc: method is nil but Bus.Done was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDone.Lock()
	mock.calls.Done = append(mock.calls.Done, callInfo)
	mock.lockDone.Unlock()
	return mock.DoneFunc()
}

// DoneCalls gets all the calls that were made to Done.
// Check the length with:
//     len(mockedBus.DoneCalls())
func (mock *BusMock) DoneCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDone.RLock()
	calls = mock.calls.Done
	mock.lockDone.RUnlock()
	return calls
}

// Publish calls PublishFunc.
func (mock *BusMock) Publish(event pubsub.Event) error {
	if mock.PublishFunc == nil {
		panic("BusMock.PublishFunc: method is nil but Bus.Publish was just called")
	}
	callInfo := struct {
		Event pubsub.Event
	}{
		Event: event,
	}
	mock.lockPublish.Lock()
	mock.calls.Publish = append(mock.calls.Publish, callInfo)
	mock.lockPublish.Unlock()
	return mock.PublishFunc(event)
}

// PublishCalls gets all the calls that were made to Publish.
// Check the length with:
//     len(mockedBus.PublishCalls())
func (mock *BusMock) PublishCalls() []struct {
	Event pubsub.Event
} {
	var calls []struct {
		Event pubsub.Event
	}
	mock.lockPublish.RLock()
	calls = mock.calls.Publish
	mock.lockPublish.RUnlock()
	return calls
}

// Subscribe calls SubscribeFunc.
func (mock *BusMock) Subscribe() (pubsub.Subscriber, error) {
	if mock.SubscribeFunc == nil {
		panic("BusMock.SubscribeFunc: method is nil but Bus.Subscribe was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSubscribe.Lock()
	mock.calls.Subscribe = append(mock.calls.Subscribe, callInfo)
	mock.lockSubscribe.Unlock()
	return mock.SubscribeFunc()
}

// SubscribeCalls gets all the calls that were made to Subscribe.
// Check the length with:
//     len(mockedBus.SubscribeCalls())
func (mock *BusMock) SubscribeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSubscribe.RLock()
	calls = mock.calls.Subscribe
	mock.lockSubscribe.RUnlock()
	return calls
}

// Ensure, that SubscriberMock does implement Subscriber.
// If this is not the case, regenerate this file with moq.
var _ Subscriber = &SubscriberMock{}

// SubscriberMock is a mock implementation of Subscriber.
//
// 	func TestSomethingThatUsesSubscriber(t *testing.T) {
//
// 		// make and configure a mocked Subscriber
// 		mockedSubscriber := &SubscriberMock{
// 			CloneFunc: func() (pubsub.Subscriber, error) {
// 				panic("mock out the Clone method")
// 			},
// 			CloseFunc: func()  {
// 				panic("mock out the Close method")
// 			},
// 			DoneFunc: func() <-chan struct{} {
// 				panic("mock out the Done method")
// 			},
// 			EventsFunc: func() <-chan pubsub.Event {
// 				panic("mock out the Events method")
// 			},
// 		}
//
// 		// use mockedSubscriber in code that requires Subscriber
// 		// and then make assertions.
//
// 	}
type SubscriberMock struct {
	// CloneFunc mocks the Clone method.
	CloneFunc func() (pubsub.Subscriber, error)

	// CloseFunc mocks the Close method.
	CloseFunc func()

	// DoneFunc mocks the Done method.
	DoneFunc func() <-chan struct{}

	// EventsFunc mocks the Events method.
	EventsFunc func() <-chan pubsub.Event

	// calls tracks calls to the methods.
	calls struct {
		// Clone holds details about calls to the Clone method.
		Clone []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Done holds details about calls to the Done method.
		Done []struct {
		}
		// Events holds details about calls to the Events method.
		Events []struct {
		}
	}
	lockClone  sync.RWMutex
	lockClose  sync.RWMutex
	lockDone   sync.RWMutex
	lockEvents sync.RWMutex
}

// Clone calls CloneFunc.
func (mock *SubscriberMock) Clone() (pubsub.Subscriber, error) {
	if mock.CloneFunc == nil {
		panic("SubscriberMock.CloneFunc: method is nil but Subscriber.Clone was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClone.Lock()
	mock.calls.Clone = append(mock.calls.Clone, callInfo)
	mock.lockClone.Unlock()
	return mock.CloneFunc()
}

// CloneCalls gets all the calls that were made to Clone.
// Check the length with:
//     len(mockedSubscriber.CloneCalls())
func (mock *SubscriberMock) CloneCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClone.RLock()
	calls = mock.calls.Clone
	mock.lockClone.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *SubscriberMock) Close() {
	if mock.CloseFunc == nil {
		panic("SubscriberMock.CloseFunc: method is nil but Subscriber.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedSubscriber.CloseCalls())
func (mock *SubscriberMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Done calls DoneFunc.
func (mock *SubscriberMock) Done() <-chan struct{} {
	if mock.DoneFunc == nil {
		panic("SubscriberMock.DoneFunc: method is nil but Subscriber.Done was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDone.Lock()
	mock.calls.Done = append(mock.calls.Done, callInfo)
	mock.lockDone.Unlock()
	return mock.DoneFunc()
}

// DoneCalls gets all the calls that were made to Done.
// Check the length with:
//     len(mockedSubscriber.DoneCalls())
func (mock *SubscriberMock) DoneCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDone.RLock()
	calls = mock.calls.Done
	mock.lockDone.RUnlock()
	return calls
}

// Events calls EventsFunc.
func (mock *SubscriberMock) Events() <-chan pubsub.Event {
	if mock.EventsFunc == nil {
		panic("SubscriberMock.EventsFunc: method is nil but Subscriber.Events was just called")
	}
	callInfo := struct {
	}{}
	mock.lockEvents.Lock()
	mock.calls.Events = append(mock.calls.Events, callInfo)
	mock.lockEvents.Unlock()
	return mock.EventsFunc()
}

// EventsCalls gets all the calls that were made to Events.
// Check the length with:
//     len(mockedSubscriber.EventsCalls())
func (mock *SubscriberMock) EventsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockEvents.RLock()
	calls = mock.calls.Events
	mock.lockEvents.RUnlock()
	return calls
}

// Ensure, that ReadWriteSeekTruncateCloserMock does implement ReadWriteSeekTruncateCloser.
// If this is not the case, regenerate this file with moq.
var _ ReadWriteSeekTruncateCloser = &ReadWriteSeekTruncateCloserMock{}

// ReadWriteSeekTruncateCloserMock is a mock implementation of ReadWriteSeekTruncateCloser.
//
// 	func TestSomethingThatUsesReadWriteSeekTruncateCloser(t *testing.T) {
//
// 		// make and configure a mocked ReadWriteSeekTruncateCloser
// 		mockedReadWriteSeekTruncateCloser := &ReadWriteSeekTruncateCloserMock{
// 			CloseFunc: func() error {
// 				panic("mock out the Close method")
// 			},
// 			ReadFunc: func(p []byte) (int, error) {
// 				panic("mock out the Read method")
// 			},
// 			SeekFunc: func(offset int64, whence int) (int64, error) {
// 				panic("mock out the Seek method")
// 			},
// 			TruncateFunc: func(size int64) error {
// 				panic("mock out the Truncate method")
// 			},
// 			WriteFunc: func(p []byte) (int, error) {
// 				panic("mock out the Write method")
// 			},
// 		}
//
// 		// use mockedReadWriteSeekTruncateCloser in code that requires ReadWriteSeekTruncateCloser
// 		// and then make assertions.
//
// 	}
type ReadWriteSeekTruncateCloserMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ReadFunc mocks the Read method.
	ReadFunc func(p []byte) (int, error)

	// SeekFunc mocks the Seek method.
	SeekFunc func(offset int64, whence int) (int64, error)

	// TruncateFunc mocks the Truncate method.
	TruncateFunc func(size int64) error

	// WriteFunc mocks the Write method.
	WriteFunc func(p []byte) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Read holds details about calls to the Read method.
		Read []struct {
			// P is the p argument value.
			P []byte
		}
		// Seek holds details about calls to the Seek method.
		Seek []struct {
			// Offset is the offset argument value.
			Offset int64
			// Whence is the whence argument value.
			Whence int
		}
		// Truncate holds details about calls to the Truncate method.
		Truncate []struct {
			// Size is the size argument value.
			Size int64
		}
		// Write holds details about calls to the Write method.
		Write []struct {
			// P is the p argument value.
			P []byte
		}
	}
	lockClose    sync.RWMutex
	lockRead     sync.RWMutex
	lockSeek     sync.RWMutex
	lockTruncate sync.RWMutex
	lockWrite    sync.RWMutex
}

// Close calls CloseFunc.
func (mock *ReadWriteSeekTruncateCloserMock) Close() error {
	if mock.CloseFunc == nil {
		panic("ReadWriteSeekTruncateCloserMock.CloseFunc: method is nil but ReadWriteSeekTruncateCloser.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedReadWriteSeekTruncateCloser.CloseCalls())
func (mock *ReadWriteSeekTruncateCloserMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Read calls ReadFunc.
func (mock *ReadWriteSeekTruncateCloserMock) Read(p []byte) (int, error) {
	if mock.ReadFunc == nil {
		panic("ReadWriteSeekTruncateCloserMock.ReadFunc: method is nil but ReadWriteSeekTruncateCloser.Read was just called")
	}
	callInfo := struct {
		P []byte
	}{
		P: p,
	}
	mock.lockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	mock.lockRead.Unlock()
	return mock.ReadFunc(p)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedReadWriteSeekTruncateCloser.ReadCalls())
func (mock *ReadWriteSeekTruncateCloserMock) ReadCalls() []struct {
	P []byte
} {
	var calls []struct {
		P []byte
	}
	mock.lockRead.RLock()
	calls = mock.calls.Read
	mock.lockRead.RUnlock()
	return calls
}

// Seek calls SeekFunc.
func (mock *ReadWriteSeekTruncateCloserMock) Seek(offset int64, whence int) (int64, error) {
	if mock.SeekFunc == nil {
		panic("ReadWriteSeekTruncateCloserMock.SeekFunc: method is nil but ReadWriteSeekTruncateCloser.Seek was just called")
	}
	callInfo := struct {
		Offset int64
		Whence int
	}{
		Offset: offset,
		Whence: whence,
	}
	mock.lockSeek.Lock()
	mock.calls.Seek = append(mock.calls.Seek, callInfo)
	mock.lockSeek.Unlock()
	return mock.SeekFunc(offset, whence)
}

// SeekCalls gets all the calls that were made to Seek.
// Check the length with:
//     len(mockedReadWriteSeekTruncateCloser.SeekCalls())
func (mock *ReadWriteSeekTruncateCloserMock) SeekCalls() []struct {
	Offset int64
	Whence int
} {
	var calls []struct {
		Offset int64
		Whence int
	}
	mock.lockSeek.RLock()
	calls = mock.calls.Seek
	mock.lockSeek.RUnlock()
	return calls
}

// Truncate calls TruncateFunc.
func (mock *ReadWriteSeekTruncateCloserMock) Truncate(size int64) error {
	if mock.TruncateFunc == nil {
		panic("ReadWriteSeekTruncateCloserMock.TruncateFunc: method is nil but ReadWriteSeekTruncateCloser.Truncate was just called")
	}
	callInfo := struct {
		Size int64
	}{
		Size: size,
	}
	mock.lockTruncate.Lock()
	mock.calls.Truncate = append(mock.calls.Truncate, callInfo)
	mock.lockTruncate.Unlock()
	return mock.TruncateFunc(size)
}

// TruncateCalls gets all the calls that were made to Truncate.
// Check the length with:
//     len(mockedReadWriteSeekTruncateCloser.TruncateCalls())
func (mock *ReadWriteSeekTruncateCloserMock) TruncateCalls() []struct {
	Size int64
} {
	var calls []struct {
		Size int64
	}
	mock.lockTruncate.RLock()
	calls = mock.calls.Truncate
	mock.lockTruncate.RUnlock()
	return calls
}

// Write calls WriteFunc.
func (mock *ReadWriteSeekTruncateCloserMock) Write(p []byte) (int, error) {
	if mock.WriteFunc == nil {
		panic("ReadWriteSeekTruncateCloserMock.WriteFunc: method is nil but ReadWriteSeekTruncateCloser.Write was just called")
	}
	callInfo := struct {
		P []byte
	}{
		P: p,
	}
	mock.lockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	mock.lockWrite.Unlock()
	return mock.WriteFunc(p)
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//     len(mockedReadWriteSeekTruncateCloser.WriteCalls())
func (mock *ReadWriteSeekTruncateCloserMock) WriteCalls() []struct {
	P []byte
} {
	var calls []struct {
		P []byte
	}
	mock.lockWrite.RLock()
	calls = mock.calls.Write
	mock.lockWrite.RUnlock()
	return calls
}
