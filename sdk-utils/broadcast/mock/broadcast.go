// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/axelarnetwork/axelar-core/sdk-utils/broadcast"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that BroadcasterMock does implement broadcast.Broadcaster.
// If this is not the case, regenerate this file with moq.
var _ broadcast.Broadcaster = &BroadcasterMock{}

// BroadcasterMock is a mock implementation of broadcast.Broadcaster.
//
// 	func TestSomethingThatUsesBroadcaster(t *testing.T) {
//
// 		// make and configure a mocked broadcast.Broadcaster
// 		mockedBroadcaster := &BroadcasterMock{
// 			BroadcastFunc: func(ctx context.Context, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
// 				panic("mock out the Broadcast method")
// 			},
// 		}
//
// 		// use mockedBroadcaster in code that requires broadcast.Broadcaster
// 		// and then make assertions.
//
// 	}
type BroadcasterMock struct {
	// BroadcastFunc mocks the Broadcast method.
	BroadcastFunc func(ctx context.Context, msgs ...sdk.Msg) (*sdk.TxResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// Broadcast holds details about calls to the Broadcast method.
		Broadcast []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msgs is the msgs argument value.
			Msgs []sdk.Msg
		}
	}
	lockBroadcast sync.RWMutex
}

// Broadcast calls BroadcastFunc.
func (mock *BroadcasterMock) Broadcast(ctx context.Context, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	if mock.BroadcastFunc == nil {
		panic("BroadcasterMock.BroadcastFunc: method is nil but Broadcaster.Broadcast was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Msgs []sdk.Msg
	}{
		Ctx:  ctx,
		Msgs: msgs,
	}
	mock.lockBroadcast.Lock()
	mock.calls.Broadcast = append(mock.calls.Broadcast, callInfo)
	mock.lockBroadcast.Unlock()
	return mock.BroadcastFunc(ctx, msgs...)
}

// BroadcastCalls gets all the calls that were made to Broadcast.
// Check the length with:
//     len(mockedBroadcaster.BroadcastCalls())
func (mock *BroadcasterMock) BroadcastCalls() []struct {
	Ctx  context.Context
	Msgs []sdk.Msg
} {
	var calls []struct {
		Ctx  context.Context
		Msgs []sdk.Msg
	}
	mock.lockBroadcast.RLock()
	calls = mock.calls.Broadcast
	mock.lockBroadcast.RUnlock()
	return calls
}
