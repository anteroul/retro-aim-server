// Code generated by mockery v2.35.2. DO NOT EDIT.

package server

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	oscar "github.com/mkaminski/goaim/oscar"
)

// MockOServiceChatHandler is an autogenerated mock type for the OServiceChatHandler type
type MockOServiceChatHandler struct {
	mock.Mock
}

type MockOServiceChatHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOServiceChatHandler) EXPECT() *MockOServiceChatHandler_Expecter {
	return &MockOServiceChatHandler_Expecter{mock: &_m.Mock}
}

// ClientOnlineHandler provides a mock function with given fields: ctx, snacPayloadIn, sess, room
func (_m *MockOServiceChatHandler) ClientOnlineHandler(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x02_OServiceClientOnline, sess *Session, room ChatRoom) error {
	ret := _m.Called(ctx, snacPayloadIn, sess, room)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNAC_0x01_0x02_OServiceClientOnline, *Session, ChatRoom) error); ok {
		r0 = rf(ctx, snacPayloadIn, sess, room)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOServiceChatHandler_ClientOnlineHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientOnlineHandler'
type MockOServiceChatHandler_ClientOnlineHandler_Call struct {
	*mock.Call
}

// ClientOnlineHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - snacPayloadIn oscar.SNAC_0x01_0x02_OServiceClientOnline
//   - sess *Session
//   - room ChatRoom
func (_e *MockOServiceChatHandler_Expecter) ClientOnlineHandler(ctx interface{}, snacPayloadIn interface{}, sess interface{}, room interface{}) *MockOServiceChatHandler_ClientOnlineHandler_Call {
	return &MockOServiceChatHandler_ClientOnlineHandler_Call{Call: _e.mock.On("ClientOnlineHandler", ctx, snacPayloadIn, sess, room)}
}

func (_c *MockOServiceChatHandler_ClientOnlineHandler_Call) Run(run func(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x02_OServiceClientOnline, sess *Session, room ChatRoom)) *MockOServiceChatHandler_ClientOnlineHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x02_OServiceClientOnline), args[2].(*Session), args[3].(ChatRoom))
	})
	return _c
}

func (_c *MockOServiceChatHandler_ClientOnlineHandler_Call) Return(_a0 error) *MockOServiceChatHandler_ClientOnlineHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceChatHandler_ClientOnlineHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x02_OServiceClientOnline, *Session, ChatRoom) error) *MockOServiceChatHandler_ClientOnlineHandler_Call {
	_c.Call.Return(run)
	return _c
}

// ClientVersionsHandler provides a mock function with given fields: ctx, snacPayloadIn
func (_m *MockOServiceChatHandler) ClientVersionsHandler(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions) XMessage {
	ret := _m.Called(ctx, snacPayloadIn)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNAC_0x01_0x17_OServiceClientVersions) XMessage); ok {
		r0 = rf(ctx, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockOServiceChatHandler_ClientVersionsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientVersionsHandler'
type MockOServiceChatHandler_ClientVersionsHandler_Call struct {
	*mock.Call
}

// ClientVersionsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions
func (_e *MockOServiceChatHandler_Expecter) ClientVersionsHandler(ctx interface{}, snacPayloadIn interface{}) *MockOServiceChatHandler_ClientVersionsHandler_Call {
	return &MockOServiceChatHandler_ClientVersionsHandler_Call{Call: _e.mock.On("ClientVersionsHandler", ctx, snacPayloadIn)}
}

func (_c *MockOServiceChatHandler_ClientVersionsHandler_Call) Run(run func(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions)) *MockOServiceChatHandler_ClientVersionsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x17_OServiceClientVersions))
	})
	return _c
}

func (_c *MockOServiceChatHandler_ClientVersionsHandler_Call) Return(_a0 XMessage) *MockOServiceChatHandler_ClientVersionsHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceChatHandler_ClientVersionsHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x17_OServiceClientVersions) XMessage) *MockOServiceChatHandler_ClientVersionsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// IdleNotificationHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockOServiceChatHandler) IdleNotificationHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification) error {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x01_0x11_OServiceIdleNotification) error); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOServiceChatHandler_IdleNotificationHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IdleNotificationHandler'
type MockOServiceChatHandler_IdleNotificationHandler_Call struct {
	*mock.Call
}

// IdleNotificationHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification
func (_e *MockOServiceChatHandler_Expecter) IdleNotificationHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockOServiceChatHandler_IdleNotificationHandler_Call {
	return &MockOServiceChatHandler_IdleNotificationHandler_Call{Call: _e.mock.On("IdleNotificationHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockOServiceChatHandler_IdleNotificationHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification)) *MockOServiceChatHandler_IdleNotificationHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x01_0x11_OServiceIdleNotification))
	})
	return _c
}

func (_c *MockOServiceChatHandler_IdleNotificationHandler_Call) Return(_a0 error) *MockOServiceChatHandler_IdleNotificationHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceChatHandler_IdleNotificationHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x01_0x11_OServiceIdleNotification) error) *MockOServiceChatHandler_IdleNotificationHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsQueryHandler provides a mock function with given fields: ctx
func (_m *MockOServiceChatHandler) RateParamsQueryHandler(ctx context.Context) XMessage {
	ret := _m.Called(ctx)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context) XMessage); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockOServiceChatHandler_RateParamsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsQueryHandler'
type MockOServiceChatHandler_RateParamsQueryHandler_Call struct {
	*mock.Call
}

// RateParamsQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockOServiceChatHandler_Expecter) RateParamsQueryHandler(ctx interface{}) *MockOServiceChatHandler_RateParamsQueryHandler_Call {
	return &MockOServiceChatHandler_RateParamsQueryHandler_Call{Call: _e.mock.On("RateParamsQueryHandler", ctx)}
}

func (_c *MockOServiceChatHandler_RateParamsQueryHandler_Call) Run(run func(ctx context.Context)) *MockOServiceChatHandler_RateParamsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockOServiceChatHandler_RateParamsQueryHandler_Call) Return(_a0 XMessage) *MockOServiceChatHandler_RateParamsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceChatHandler_RateParamsQueryHandler_Call) RunAndReturn(run func(context.Context) XMessage) *MockOServiceChatHandler_RateParamsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsSubAddHandler provides a mock function with given fields: _a0, _a1
func (_m *MockOServiceChatHandler) RateParamsSubAddHandler(_a0 context.Context, _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd) {
	_m.Called(_a0, _a1)
}

// MockOServiceChatHandler_RateParamsSubAddHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsSubAddHandler'
type MockOServiceChatHandler_RateParamsSubAddHandler_Call struct {
	*mock.Call
}

// RateParamsSubAddHandler is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd
func (_e *MockOServiceChatHandler_Expecter) RateParamsSubAddHandler(_a0 interface{}, _a1 interface{}) *MockOServiceChatHandler_RateParamsSubAddHandler_Call {
	return &MockOServiceChatHandler_RateParamsSubAddHandler_Call{Call: _e.mock.On("RateParamsSubAddHandler", _a0, _a1)}
}

func (_c *MockOServiceChatHandler_RateParamsSubAddHandler_Call) Run(run func(_a0 context.Context, _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *MockOServiceChatHandler_RateParamsSubAddHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd))
	})
	return _c
}

func (_c *MockOServiceChatHandler_RateParamsSubAddHandler_Call) Return() *MockOServiceChatHandler_RateParamsSubAddHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockOServiceChatHandler_RateParamsSubAddHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *MockOServiceChatHandler_RateParamsSubAddHandler_Call {
	_c.Call.Return(run)
	return _c
}

// ServiceRequestHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockOServiceChatHandler) ServiceRequestHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x04_OServiceServiceRequest) (XMessage, error) {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x01_0x04_OServiceServiceRequest) (XMessage, error)); ok {
		return rf(ctx, sess, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x01_0x04_OServiceServiceRequest) XMessage); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, oscar.SNAC_0x01_0x04_OServiceServiceRequest) error); ok {
		r1 = rf(ctx, sess, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOServiceChatHandler_ServiceRequestHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServiceRequestHandler'
type MockOServiceChatHandler_ServiceRequestHandler_Call struct {
	*mock.Call
}

// ServiceRequestHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x01_0x04_OServiceServiceRequest
func (_e *MockOServiceChatHandler_Expecter) ServiceRequestHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockOServiceChatHandler_ServiceRequestHandler_Call {
	return &MockOServiceChatHandler_ServiceRequestHandler_Call{Call: _e.mock.On("ServiceRequestHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockOServiceChatHandler_ServiceRequestHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x04_OServiceServiceRequest)) *MockOServiceChatHandler_ServiceRequestHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x01_0x04_OServiceServiceRequest))
	})
	return _c
}

func (_c *MockOServiceChatHandler_ServiceRequestHandler_Call) Return(_a0 XMessage, _a1 error) *MockOServiceChatHandler_ServiceRequestHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOServiceChatHandler_ServiceRequestHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x01_0x04_OServiceServiceRequest) (XMessage, error)) *MockOServiceChatHandler_ServiceRequestHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserInfoFieldsHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockOServiceChatHandler) SetUserInfoFieldsHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (XMessage, error) {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (XMessage, error)); ok {
		return rf(ctx, sess, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) XMessage); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) error); ok {
		r1 = rf(ctx, sess, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOServiceChatHandler_SetUserInfoFieldsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserInfoFieldsHandler'
type MockOServiceChatHandler_SetUserInfoFieldsHandler_Call struct {
	*mock.Call
}

// SetUserInfoFieldsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields
func (_e *MockOServiceChatHandler_Expecter) SetUserInfoFieldsHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	return &MockOServiceChatHandler_SetUserInfoFieldsHandler_Call{Call: _e.mock.On("SetUserInfoFieldsHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields)) *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields))
	})
	return _c
}

func (_c *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call) Return(_a0 XMessage, _a1 error) *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (XMessage, error)) *MockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// UserInfoQueryHandler provides a mock function with given fields: ctx, sess
func (_m *MockOServiceChatHandler) UserInfoQueryHandler(ctx context.Context, sess *Session) XMessage {
	ret := _m.Called(ctx, sess)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context, *Session) XMessage); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockOServiceChatHandler_UserInfoQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfoQueryHandler'
type MockOServiceChatHandler_UserInfoQueryHandler_Call struct {
	*mock.Call
}

// UserInfoQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
func (_e *MockOServiceChatHandler_Expecter) UserInfoQueryHandler(ctx interface{}, sess interface{}) *MockOServiceChatHandler_UserInfoQueryHandler_Call {
	return &MockOServiceChatHandler_UserInfoQueryHandler_Call{Call: _e.mock.On("UserInfoQueryHandler", ctx, sess)}
}

func (_c *MockOServiceChatHandler_UserInfoQueryHandler_Call) Run(run func(ctx context.Context, sess *Session)) *MockOServiceChatHandler_UserInfoQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session))
	})
	return _c
}

func (_c *MockOServiceChatHandler_UserInfoQueryHandler_Call) Return(_a0 XMessage) *MockOServiceChatHandler_UserInfoQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceChatHandler_UserInfoQueryHandler_Call) RunAndReturn(run func(context.Context, *Session) XMessage) *MockOServiceChatHandler_UserInfoQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// WriteOServiceHostOnline provides a mock function with given fields: w, sequence
func (_m *MockOServiceChatHandler) WriteOServiceHostOnline(w io.Writer, sequence *uint32) error {
	ret := _m.Called(w, sequence)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer, *uint32) error); ok {
		r0 = rf(w, sequence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOServiceChatHandler_WriteOServiceHostOnline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteOServiceHostOnline'
type MockOServiceChatHandler_WriteOServiceHostOnline_Call struct {
	*mock.Call
}

// WriteOServiceHostOnline is a helper method to define mock.On call
//   - w io.Writer
//   - sequence *uint32
func (_e *MockOServiceChatHandler_Expecter) WriteOServiceHostOnline(w interface{}, sequence interface{}) *MockOServiceChatHandler_WriteOServiceHostOnline_Call {
	return &MockOServiceChatHandler_WriteOServiceHostOnline_Call{Call: _e.mock.On("WriteOServiceHostOnline", w, sequence)}
}

func (_c *MockOServiceChatHandler_WriteOServiceHostOnline_Call) Run(run func(w io.Writer, sequence *uint32)) *MockOServiceChatHandler_WriteOServiceHostOnline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Writer), args[1].(*uint32))
	})
	return _c
}

func (_c *MockOServiceChatHandler_WriteOServiceHostOnline_Call) Return(_a0 error) *MockOServiceChatHandler_WriteOServiceHostOnline_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceChatHandler_WriteOServiceHostOnline_Call) RunAndReturn(run func(io.Writer, *uint32) error) *MockOServiceChatHandler_WriteOServiceHostOnline_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOServiceChatHandler creates a new instance of MockOServiceChatHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOServiceChatHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOServiceChatHandler {
	mock := &MockOServiceChatHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
