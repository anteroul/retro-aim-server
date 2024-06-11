// Code generated by mockery v2.40.1. DO NOT EDIT.

package foodgroup

import (
	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"
)

// mockLegacyBuddyListManager is an autogenerated mock type for the LegacyBuddyListManager type
type mockLegacyBuddyListManager struct {
	mock.Mock
}

type mockLegacyBuddyListManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockLegacyBuddyListManager) EXPECT() *mockLegacyBuddyListManager_Expecter {
	return &mockLegacyBuddyListManager_Expecter{mock: &_m.Mock}
}

// AddBuddy provides a mock function with given fields: userScreenName, buddyScreenName
func (_m *mockLegacyBuddyListManager) AddBuddy(userScreenName state.IdentScreenName, buddyScreenName state.IdentScreenName) {
	_m.Called(userScreenName, buddyScreenName)
}

// mockLegacyBuddyListManager_AddBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBuddy'
type mockLegacyBuddyListManager_AddBuddy_Call struct {
	*mock.Call
}

// AddBuddy is a helper method to define mock.On call
//   - userScreenName state.IdentScreenName
//   - buddyScreenName state.IdentScreenName
func (_e *mockLegacyBuddyListManager_Expecter) AddBuddy(userScreenName interface{}, buddyScreenName interface{}) *mockLegacyBuddyListManager_AddBuddy_Call {
	return &mockLegacyBuddyListManager_AddBuddy_Call{Call: _e.mock.On("AddBuddy", userScreenName, buddyScreenName)}
}

func (_c *mockLegacyBuddyListManager_AddBuddy_Call) Run(run func(userScreenName state.IdentScreenName, buddyScreenName state.IdentScreenName)) *mockLegacyBuddyListManager_AddBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLegacyBuddyListManager_AddBuddy_Call) Return() *mockLegacyBuddyListManager_AddBuddy_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockLegacyBuddyListManager_AddBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName)) *mockLegacyBuddyListManager_AddBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// Buddies provides a mock function with given fields: userScreenName
func (_m *mockLegacyBuddyListManager) Buddies(userScreenName state.IdentScreenName) []state.IdentScreenName {
	ret := _m.Called(userScreenName)

	if len(ret) == 0 {
		panic("no return value specified for Buddies")
	}

	var r0 []state.IdentScreenName
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) []state.IdentScreenName); ok {
		r0 = rf(userScreenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.IdentScreenName)
		}
	}

	return r0
}

// mockLegacyBuddyListManager_Buddies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Buddies'
type mockLegacyBuddyListManager_Buddies_Call struct {
	*mock.Call
}

// Buddies is a helper method to define mock.On call
//   - userScreenName state.IdentScreenName
func (_e *mockLegacyBuddyListManager_Expecter) Buddies(userScreenName interface{}) *mockLegacyBuddyListManager_Buddies_Call {
	return &mockLegacyBuddyListManager_Buddies_Call{Call: _e.mock.On("Buddies", userScreenName)}
}

func (_c *mockLegacyBuddyListManager_Buddies_Call) Run(run func(userScreenName state.IdentScreenName)) *mockLegacyBuddyListManager_Buddies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLegacyBuddyListManager_Buddies_Call) Return(_a0 []state.IdentScreenName) *mockLegacyBuddyListManager_Buddies_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLegacyBuddyListManager_Buddies_Call) RunAndReturn(run func(state.IdentScreenName) []state.IdentScreenName) *mockLegacyBuddyListManager_Buddies_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBuddy provides a mock function with given fields: userScreenName, buddyScreenName
func (_m *mockLegacyBuddyListManager) DeleteBuddy(userScreenName state.IdentScreenName, buddyScreenName state.IdentScreenName) {
	_m.Called(userScreenName, buddyScreenName)
}

// mockLegacyBuddyListManager_DeleteBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBuddy'
type mockLegacyBuddyListManager_DeleteBuddy_Call struct {
	*mock.Call
}

// DeleteBuddy is a helper method to define mock.On call
//   - userScreenName state.IdentScreenName
//   - buddyScreenName state.IdentScreenName
func (_e *mockLegacyBuddyListManager_Expecter) DeleteBuddy(userScreenName interface{}, buddyScreenName interface{}) *mockLegacyBuddyListManager_DeleteBuddy_Call {
	return &mockLegacyBuddyListManager_DeleteBuddy_Call{Call: _e.mock.On("DeleteBuddy", userScreenName, buddyScreenName)}
}

func (_c *mockLegacyBuddyListManager_DeleteBuddy_Call) Run(run func(userScreenName state.IdentScreenName, buddyScreenName state.IdentScreenName)) *mockLegacyBuddyListManager_DeleteBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLegacyBuddyListManager_DeleteBuddy_Call) Return() *mockLegacyBuddyListManager_DeleteBuddy_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockLegacyBuddyListManager_DeleteBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName)) *mockLegacyBuddyListManager_DeleteBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: userScreenName
func (_m *mockLegacyBuddyListManager) DeleteUser(userScreenName state.IdentScreenName) {
	_m.Called(userScreenName)
}

// mockLegacyBuddyListManager_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type mockLegacyBuddyListManager_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - userScreenName state.IdentScreenName
func (_e *mockLegacyBuddyListManager_Expecter) DeleteUser(userScreenName interface{}) *mockLegacyBuddyListManager_DeleteUser_Call {
	return &mockLegacyBuddyListManager_DeleteUser_Call{Call: _e.mock.On("DeleteUser", userScreenName)}
}

func (_c *mockLegacyBuddyListManager_DeleteUser_Call) Run(run func(userScreenName state.IdentScreenName)) *mockLegacyBuddyListManager_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLegacyBuddyListManager_DeleteUser_Call) Return() *mockLegacyBuddyListManager_DeleteUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockLegacyBuddyListManager_DeleteUser_Call) RunAndReturn(run func(state.IdentScreenName)) *mockLegacyBuddyListManager_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// WhoAddedUser provides a mock function with given fields: userScreenName
func (_m *mockLegacyBuddyListManager) WhoAddedUser(userScreenName state.IdentScreenName) []state.IdentScreenName {
	ret := _m.Called(userScreenName)

	if len(ret) == 0 {
		panic("no return value specified for WhoAddedUser")
	}

	var r0 []state.IdentScreenName
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) []state.IdentScreenName); ok {
		r0 = rf(userScreenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.IdentScreenName)
		}
	}

	return r0
}

// mockLegacyBuddyListManager_WhoAddedUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WhoAddedUser'
type mockLegacyBuddyListManager_WhoAddedUser_Call struct {
	*mock.Call
}

// WhoAddedUser is a helper method to define mock.On call
//   - userScreenName state.IdentScreenName
func (_e *mockLegacyBuddyListManager_Expecter) WhoAddedUser(userScreenName interface{}) *mockLegacyBuddyListManager_WhoAddedUser_Call {
	return &mockLegacyBuddyListManager_WhoAddedUser_Call{Call: _e.mock.On("WhoAddedUser", userScreenName)}
}

func (_c *mockLegacyBuddyListManager_WhoAddedUser_Call) Run(run func(userScreenName state.IdentScreenName)) *mockLegacyBuddyListManager_WhoAddedUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLegacyBuddyListManager_WhoAddedUser_Call) Return(_a0 []state.IdentScreenName) *mockLegacyBuddyListManager_WhoAddedUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLegacyBuddyListManager_WhoAddedUser_Call) RunAndReturn(run func(state.IdentScreenName) []state.IdentScreenName) *mockLegacyBuddyListManager_WhoAddedUser_Call {
	_c.Call.Return(run)
	return _c
}

// newMockLegacyBuddyListManager creates a new instance of mockLegacyBuddyListManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockLegacyBuddyListManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockLegacyBuddyListManager {
	mock := &mockLegacyBuddyListManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
