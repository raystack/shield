// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	resource "github.com/odpf/shield/core/resource"
	mock "github.com/stretchr/testify/mock"
)

// ResourceService is an autogenerated mock type for the ResourceService type
type ResourceService struct {
	mock.Mock
}

type ResourceService_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceService) EXPECT() *ResourceService_Expecter {
	return &ResourceService_Expecter{mock: &_m.Mock}
}

// CheckAuthz provides a mock function with given fields: ctx, _a1, permissionName
func (_m *ResourceService) CheckAuthz(ctx context.Context, _a1 resource.Resource, permissionName string) (bool, error) {
	ret := _m.Called(ctx, _a1, permissionName)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource, string) (bool, error)); ok {
		return rf(ctx, _a1, permissionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource, string) bool); ok {
		r0 = rf(ctx, _a1, permissionName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.Resource, string) error); ok {
		r1 = rf(ctx, _a1, permissionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_CheckAuthz_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckAuthz'
type ResourceService_CheckAuthz_Call struct {
	*mock.Call
}

// CheckAuthz is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 resource.Resource
//   - permissionName string
func (_e *ResourceService_Expecter) CheckAuthz(ctx interface{}, _a1 interface{}, permissionName interface{}) *ResourceService_CheckAuthz_Call {
	return &ResourceService_CheckAuthz_Call{Call: _e.mock.On("CheckAuthz", ctx, _a1, permissionName)}
}

func (_c *ResourceService_CheckAuthz_Call) Run(run func(ctx context.Context, _a1 resource.Resource, permissionName string)) *ResourceService_CheckAuthz_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.Resource), args[2].(string))
	})
	return _c
}

func (_c *ResourceService_CheckAuthz_Call) Return(_a0 bool, _a1 error) *ResourceService_CheckAuthz_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_CheckAuthz_Call) RunAndReturn(run func(context.Context, resource.Resource, string) (bool, error)) *ResourceService_CheckAuthz_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *ResourceService) Create(ctx context.Context, _a1 resource.Resource) (resource.Resource, error) {
	ret := _m.Called(ctx, _a1)

	var r0 resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource) (resource.Resource, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource) resource.Resource); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(resource.Resource)
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.Resource) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ResourceService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 resource.Resource
func (_e *ResourceService_Expecter) Create(ctx interface{}, _a1 interface{}) *ResourceService_Create_Call {
	return &ResourceService_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *ResourceService_Create_Call) Run(run func(ctx context.Context, _a1 resource.Resource)) *ResourceService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.Resource))
	})
	return _c
}

func (_c *ResourceService_Create_Call) Return(_a0 resource.Resource, _a1 error) *ResourceService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_Create_Call) RunAndReturn(run func(context.Context, resource.Resource) (resource.Resource, error)) *ResourceService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *ResourceService) Get(ctx context.Context, id string) (resource.Resource, error) {
	ret := _m.Called(ctx, id)

	var r0 resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (resource.Resource, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) resource.Resource); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(resource.Resource)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ResourceService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ResourceService_Expecter) Get(ctx interface{}, id interface{}) *ResourceService_Get_Call {
	return &ResourceService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *ResourceService_Get_Call) Run(run func(ctx context.Context, id string)) *ResourceService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ResourceService_Get_Call) Return(_a0 resource.Resource, _a1 error) *ResourceService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_Get_Call) RunAndReturn(run func(context.Context, string) (resource.Resource, error)) *ResourceService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, flt
func (_m *ResourceService) List(ctx context.Context, flt resource.Filter) ([]resource.Resource, error) {
	ret := _m.Called(ctx, flt)

	var r0 []resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Filter) ([]resource.Resource, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.Filter) []resource.Resource); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]resource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ResourceService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - flt resource.Filter
func (_e *ResourceService_Expecter) List(ctx interface{}, flt interface{}) *ResourceService_List_Call {
	return &ResourceService_List_Call{Call: _e.mock.On("List", ctx, flt)}
}

func (_c *ResourceService_List_Call) Run(run func(ctx context.Context, flt resource.Filter)) *ResourceService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.Filter))
	})
	return _c
}

func (_c *ResourceService_List_Call) Return(_a0 []resource.Resource, _a1 error) *ResourceService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_List_Call) RunAndReturn(run func(context.Context, resource.Filter) ([]resource.Resource, error)) *ResourceService_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *ResourceService) Update(ctx context.Context, _a1 resource.Resource) (resource.Resource, error) {
	ret := _m.Called(ctx, _a1)

	var r0 resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource) (resource.Resource, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource) resource.Resource); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(resource.Resource)
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.Resource) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ResourceService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 resource.Resource
func (_e *ResourceService_Expecter) Update(ctx interface{}, _a1 interface{}) *ResourceService_Update_Call {
	return &ResourceService_Update_Call{Call: _e.mock.On("Update", ctx, _a1)}
}

func (_c *ResourceService_Update_Call) Run(run func(ctx context.Context, _a1 resource.Resource)) *ResourceService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.Resource))
	})
	return _c
}

func (_c *ResourceService_Update_Call) Return(_a0 resource.Resource, _a1 error) *ResourceService_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_Update_Call) RunAndReturn(run func(context.Context, resource.Resource) (resource.Resource, error)) *ResourceService_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewResourceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceService creates a new instance of ResourceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceService(t mockConstructorTestingTNewResourceService) *ResourceService {
	mock := &ResourceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
