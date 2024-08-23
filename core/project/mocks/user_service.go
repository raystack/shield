// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/raystack/frontier/core/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserService) GetByID(ctx context.Context, id string) (user.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type UserService_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserService_Expecter) GetByID(ctx interface{}, id interface{}) *UserService_GetByID_Call {
	return &UserService_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *UserService_GetByID_Call) Run(run func(ctx context.Context, id string)) *UserService_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_GetByID_Call) Return(_a0 user.User, _a1 error) *UserService_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetByID_Call) RunAndReturn(run func(context.Context, string) (user.User, error)) *UserService_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, userIDs
func (_m *UserService) GetByIDs(ctx context.Context, userIDs []string) ([]user.User, error) {
	ret := _m.Called(ctx, userIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]user.User, error)); ok {
		return rf(ctx, userIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []user.User); ok {
		r0 = rf(ctx, userIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, userIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type UserService_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - userIDs []string
func (_e *UserService_Expecter) GetByIDs(ctx interface{}, userIDs interface{}) *UserService_GetByIDs_Call {
	return &UserService_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, userIDs)}
}

func (_c *UserService_GetByIDs_Call) Run(run func(ctx context.Context, userIDs []string)) *UserService_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *UserService_GetByIDs_Call) Return(_a0 []user.User, _a1 error) *UserService_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetByIDs_Call) RunAndReturn(run func(context.Context, []string) ([]user.User, error)) *UserService_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}