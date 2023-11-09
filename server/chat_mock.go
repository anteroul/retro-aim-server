// Code generated by mockery v2.35.2. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"
)

// MockChatHandler is an autogenerated mock type for the ChatHandler type
type MockChatHandler struct {
	mock.Mock
}

type MockChatHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockChatHandler) EXPECT() *MockChatHandler_Expecter {
	return &MockChatHandler_Expecter{mock: &_m.Mock}
}

// ChannelMsgToHostHandler provides a mock function with given fields: ctx, sess, room, snacPayloadIn
func (_m *MockChatHandler) ChannelMsgToHostHandler(ctx context.Context, sess *Session, room ChatRoom, snacPayloadIn oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*XMessage, error) {
	ret := _m.Called(ctx, sess, room, snacPayloadIn)

	var r0 *XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, ChatRoom, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*XMessage, error)); ok {
		return rf(ctx, sess, room, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, ChatRoom, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) *XMessage); ok {
		r0 = rf(ctx, sess, room, snacPayloadIn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*XMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, ChatRoom, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) error); ok {
		r1 = rf(ctx, sess, room, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChatHandler_ChannelMsgToHostHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelMsgToHostHandler'
type MockChatHandler_ChannelMsgToHostHandler_Call struct {
	*mock.Call
}

// ChannelMsgToHostHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - room ChatRoom
//   - snacPayloadIn oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost
func (_e *MockChatHandler_Expecter) ChannelMsgToHostHandler(ctx interface{}, sess interface{}, room interface{}, snacPayloadIn interface{}) *MockChatHandler_ChannelMsgToHostHandler_Call {
	return &MockChatHandler_ChannelMsgToHostHandler_Call{Call: _e.mock.On("ChannelMsgToHostHandler", ctx, sess, room, snacPayloadIn)}
}

func (_c *MockChatHandler_ChannelMsgToHostHandler_Call) Run(run func(ctx context.Context, sess *Session, room ChatRoom, snacPayloadIn oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost)) *MockChatHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(ChatRoom), args[3].(oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost))
	})
	return _c
}

func (_c *MockChatHandler_ChannelMsgToHostHandler_Call) Return(_a0 *XMessage, _a1 error) *MockChatHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChatHandler_ChannelMsgToHostHandler_Call) RunAndReturn(run func(context.Context, *Session, ChatRoom, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*XMessage, error)) *MockChatHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockChatHandler creates a new instance of MockChatHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChatHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChatHandler {
	mock := &MockChatHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
