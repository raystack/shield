// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"

	permission "github.com/raystack/frontier/core/permission"
	mock "github.com/stretchr/testify/mock"
)

// PermissionService is an autogenerated mock type for the PermissionService type
type PermissionService struct {
	mock.Mock
}

type PermissionService_Expecter struct {
	mock *mock.Mock
}

func (_m *PermissionService) EXPECT() *PermissionService_Expecter {
	return &PermissionService_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *PermissionService) Get(ctx context.Context, id string) (permission.Permission, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 permission.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (permission.Permission, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) permission.Permission); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(permission.Permission)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermissionService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PermissionService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *PermissionService_Expecter) Get(ctx interface{}, id interface{}) *PermissionService_Get_Call {
	return &PermissionService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *PermissionService_Get_Call) Run(run func(ctx context.Context, id string)) *PermissionService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *PermissionService_Get_Call) Return(_a0 permission.Permission, _a1 error) *PermissionService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PermissionService_Get_Call) RunAndReturn(run func(context.Context, string) (permission.Permission, error)) *PermissionService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewPermissionService creates a new instance of PermissionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPermissionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PermissionService {
	mock := &PermissionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}