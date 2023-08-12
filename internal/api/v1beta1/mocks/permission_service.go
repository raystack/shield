// Code generated by mockery v2.32.4. DO NOT EDIT.

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

// List provides a mock function with given fields: ctx, filter
func (_m *PermissionService) List(ctx context.Context, filter permission.Filter) ([]permission.Permission, error) {
	ret := _m.Called(ctx, filter)

	var r0 []permission.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, permission.Filter) ([]permission.Permission, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, permission.Filter) []permission.Permission); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]permission.Permission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, permission.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermissionService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type PermissionService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - filter permission.Filter
func (_e *PermissionService_Expecter) List(ctx interface{}, filter interface{}) *PermissionService_List_Call {
	return &PermissionService_List_Call{Call: _e.mock.On("List", ctx, filter)}
}

func (_c *PermissionService_List_Call) Run(run func(ctx context.Context, filter permission.Filter)) *PermissionService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(permission.Filter))
	})
	return _c
}

func (_c *PermissionService_List_Call) Return(_a0 []permission.Permission, _a1 error) *PermissionService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PermissionService_List_Call) RunAndReturn(run func(context.Context, permission.Filter) ([]permission.Permission, error)) *PermissionService_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, perm
func (_m *PermissionService) Update(ctx context.Context, perm permission.Permission) (permission.Permission, error) {
	ret := _m.Called(ctx, perm)

	var r0 permission.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, permission.Permission) (permission.Permission, error)); ok {
		return rf(ctx, perm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, permission.Permission) permission.Permission); ok {
		r0 = rf(ctx, perm)
	} else {
		r0 = ret.Get(0).(permission.Permission)
	}

	if rf, ok := ret.Get(1).(func(context.Context, permission.Permission) error); ok {
		r1 = rf(ctx, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermissionService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PermissionService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - perm permission.Permission
func (_e *PermissionService_Expecter) Update(ctx interface{}, perm interface{}) *PermissionService_Update_Call {
	return &PermissionService_Update_Call{Call: _e.mock.On("Update", ctx, perm)}
}

func (_c *PermissionService_Update_Call) Run(run func(ctx context.Context, perm permission.Permission)) *PermissionService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(permission.Permission))
	})
	return _c
}

func (_c *PermissionService_Update_Call) Return(_a0 permission.Permission, _a1 error) *PermissionService_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PermissionService_Update_Call) RunAndReturn(run func(context.Context, permission.Permission) (permission.Permission, error)) *PermissionService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Upsert provides a mock function with given fields: ctx, perm
func (_m *PermissionService) Upsert(ctx context.Context, perm permission.Permission) (permission.Permission, error) {
	ret := _m.Called(ctx, perm)

	var r0 permission.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, permission.Permission) (permission.Permission, error)); ok {
		return rf(ctx, perm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, permission.Permission) permission.Permission); ok {
		r0 = rf(ctx, perm)
	} else {
		r0 = ret.Get(0).(permission.Permission)
	}

	if rf, ok := ret.Get(1).(func(context.Context, permission.Permission) error); ok {
		r1 = rf(ctx, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermissionService_Upsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upsert'
type PermissionService_Upsert_Call struct {
	*mock.Call
}

// Upsert is a helper method to define mock.On call
//   - ctx context.Context
//   - perm permission.Permission
func (_e *PermissionService_Expecter) Upsert(ctx interface{}, perm interface{}) *PermissionService_Upsert_Call {
	return &PermissionService_Upsert_Call{Call: _e.mock.On("Upsert", ctx, perm)}
}

func (_c *PermissionService_Upsert_Call) Run(run func(ctx context.Context, perm permission.Permission)) *PermissionService_Upsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(permission.Permission))
	})
	return _c
}

func (_c *PermissionService_Upsert_Call) Return(_a0 permission.Permission, _a1 error) *PermissionService_Upsert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PermissionService_Upsert_Call) RunAndReturn(run func(context.Context, permission.Permission) (permission.Permission, error)) *PermissionService_Upsert_Call {
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
