// Code generated by mockery v2.35.2. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"
)

// MockOServiceHandler is an autogenerated mock type for the OServiceHandler type
type MockOServiceHandler struct {
	mock.Mock
}

type MockOServiceHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOServiceHandler) EXPECT() *MockOServiceHandler_Expecter {
	return &MockOServiceHandler_Expecter{mock: &_m.Mock}
}

// ClientVersionsHandler provides a mock function with given fields: ctx, snacPayloadIn
func (_m *MockOServiceHandler) ClientVersionsHandler(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions) XMessage {
	ret := _m.Called(ctx, snacPayloadIn)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNAC_0x01_0x17_OServiceClientVersions) XMessage); ok {
		r0 = rf(ctx, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockOServiceHandler_ClientVersionsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientVersionsHandler'
type MockOServiceHandler_ClientVersionsHandler_Call struct {
	*mock.Call
}

// ClientVersionsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions
func (_e *MockOServiceHandler_Expecter) ClientVersionsHandler(ctx interface{}, snacPayloadIn interface{}) *MockOServiceHandler_ClientVersionsHandler_Call {
	return &MockOServiceHandler_ClientVersionsHandler_Call{Call: _e.mock.On("ClientVersionsHandler", ctx, snacPayloadIn)}
}

func (_c *MockOServiceHandler_ClientVersionsHandler_Call) Run(run func(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions)) *MockOServiceHandler_ClientVersionsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x17_OServiceClientVersions))
	})
	return _c
}

func (_c *MockOServiceHandler_ClientVersionsHandler_Call) Return(_a0 XMessage) *MockOServiceHandler_ClientVersionsHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceHandler_ClientVersionsHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x17_OServiceClientVersions) XMessage) *MockOServiceHandler_ClientVersionsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// IdleNotificationHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockOServiceHandler) IdleNotificationHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification) error {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x01_0x11_OServiceIdleNotification) error); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOServiceHandler_IdleNotificationHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IdleNotificationHandler'
type MockOServiceHandler_IdleNotificationHandler_Call struct {
	*mock.Call
}

// IdleNotificationHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification
func (_e *MockOServiceHandler_Expecter) IdleNotificationHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockOServiceHandler_IdleNotificationHandler_Call {
	return &MockOServiceHandler_IdleNotificationHandler_Call{Call: _e.mock.On("IdleNotificationHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockOServiceHandler_IdleNotificationHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification)) *MockOServiceHandler_IdleNotificationHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x01_0x11_OServiceIdleNotification))
	})
	return _c
}

func (_c *MockOServiceHandler_IdleNotificationHandler_Call) Return(_a0 error) *MockOServiceHandler_IdleNotificationHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceHandler_IdleNotificationHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x01_0x11_OServiceIdleNotification) error) *MockOServiceHandler_IdleNotificationHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsQueryHandler provides a mock function with given fields: ctx
func (_m *MockOServiceHandler) RateParamsQueryHandler(ctx context.Context) XMessage {
	ret := _m.Called(ctx)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context) XMessage); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockOServiceHandler_RateParamsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsQueryHandler'
type MockOServiceHandler_RateParamsQueryHandler_Call struct {
	*mock.Call
}

// RateParamsQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockOServiceHandler_Expecter) RateParamsQueryHandler(ctx interface{}) *MockOServiceHandler_RateParamsQueryHandler_Call {
	return &MockOServiceHandler_RateParamsQueryHandler_Call{Call: _e.mock.On("RateParamsQueryHandler", ctx)}
}

func (_c *MockOServiceHandler_RateParamsQueryHandler_Call) Run(run func(ctx context.Context)) *MockOServiceHandler_RateParamsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockOServiceHandler_RateParamsQueryHandler_Call) Return(_a0 XMessage) *MockOServiceHandler_RateParamsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceHandler_RateParamsQueryHandler_Call) RunAndReturn(run func(context.Context) XMessage) *MockOServiceHandler_RateParamsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsSubAddHandler provides a mock function with given fields: _a0, _a1
func (_m *MockOServiceHandler) RateParamsSubAddHandler(_a0 context.Context, _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd) {
	_m.Called(_a0, _a1)
}

// MockOServiceHandler_RateParamsSubAddHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsSubAddHandler'
type MockOServiceHandler_RateParamsSubAddHandler_Call struct {
	*mock.Call
}

// RateParamsSubAddHandler is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd
func (_e *MockOServiceHandler_Expecter) RateParamsSubAddHandler(_a0 interface{}, _a1 interface{}) *MockOServiceHandler_RateParamsSubAddHandler_Call {
	return &MockOServiceHandler_RateParamsSubAddHandler_Call{Call: _e.mock.On("RateParamsSubAddHandler", _a0, _a1)}
}

func (_c *MockOServiceHandler_RateParamsSubAddHandler_Call) Run(run func(_a0 context.Context, _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *MockOServiceHandler_RateParamsSubAddHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd))
	})
	return _c
}

func (_c *MockOServiceHandler_RateParamsSubAddHandler_Call) Return() *MockOServiceHandler_RateParamsSubAddHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockOServiceHandler_RateParamsSubAddHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *MockOServiceHandler_RateParamsSubAddHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserInfoFieldsHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockOServiceHandler) SetUserInfoFieldsHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (XMessage, error) {
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

// MockOServiceHandler_SetUserInfoFieldsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserInfoFieldsHandler'
type MockOServiceHandler_SetUserInfoFieldsHandler_Call struct {
	*mock.Call
}

// SetUserInfoFieldsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields
func (_e *MockOServiceHandler_Expecter) SetUserInfoFieldsHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockOServiceHandler_SetUserInfoFieldsHandler_Call {
	return &MockOServiceHandler_SetUserInfoFieldsHandler_Call{Call: _e.mock.On("SetUserInfoFieldsHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockOServiceHandler_SetUserInfoFieldsHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields)) *MockOServiceHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields))
	})
	return _c
}

func (_c *MockOServiceHandler_SetUserInfoFieldsHandler_Call) Return(_a0 XMessage, _a1 error) *MockOServiceHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOServiceHandler_SetUserInfoFieldsHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (XMessage, error)) *MockOServiceHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// UserInfoQueryHandler provides a mock function with given fields: ctx, sess
func (_m *MockOServiceHandler) UserInfoQueryHandler(ctx context.Context, sess *Session) XMessage {
	ret := _m.Called(ctx, sess)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context, *Session) XMessage); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockOServiceHandler_UserInfoQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfoQueryHandler'
type MockOServiceHandler_UserInfoQueryHandler_Call struct {
	*mock.Call
}

// UserInfoQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
func (_e *MockOServiceHandler_Expecter) UserInfoQueryHandler(ctx interface{}, sess interface{}) *MockOServiceHandler_UserInfoQueryHandler_Call {
	return &MockOServiceHandler_UserInfoQueryHandler_Call{Call: _e.mock.On("UserInfoQueryHandler", ctx, sess)}
}

func (_c *MockOServiceHandler_UserInfoQueryHandler_Call) Run(run func(ctx context.Context, sess *Session)) *MockOServiceHandler_UserInfoQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session))
	})
	return _c
}

func (_c *MockOServiceHandler_UserInfoQueryHandler_Call) Return(_a0 XMessage) *MockOServiceHandler_UserInfoQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOServiceHandler_UserInfoQueryHandler_Call) RunAndReturn(run func(context.Context, *Session) XMessage) *MockOServiceHandler_UserInfoQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOServiceHandler creates a new instance of MockOServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOServiceHandler {
	mock := &MockOServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
