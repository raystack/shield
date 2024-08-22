// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	group "github.com/raystack/frontier/core/group"
	mock "github.com/stretchr/testify/mock"
)

// GroupService is an autogenerated mock type for the GroupService type
type GroupService struct {
	mock.Mock
}

type GroupService_Expecter struct {
	mock *mock.Mock
}

func (_m *GroupService) EXPECT() *GroupService_Expecter {
	return &GroupService_Expecter{mock: &_m.Mock}
}

// GetByIDs provides a mock function with given fields: ctx, ids
func (_m *GroupService) GetByIDs(ctx context.Context, ids []string) ([]group.Group, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 []group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]group.Group, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []group.Group); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupService_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type GroupService_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []string
func (_e *GroupService_Expecter) GetByIDs(ctx interface{}, ids interface{}) *GroupService_GetByIDs_Call {
	return &GroupService_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, ids)}
}

func (_c *GroupService_GetByIDs_Call) Run(run func(ctx context.Context, ids []string)) *GroupService_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *GroupService_GetByIDs_Call) Return(_a0 []group.Group, _a1 error) *GroupService_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GroupService_GetByIDs_Call) RunAndReturn(run func(context.Context, []string) ([]group.Group, error)) *GroupService_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// ListByUser provides a mock function with given fields: ctx, principalID, principalType, flt
func (_m *GroupService) ListByUser(ctx context.Context, principalID string, principalType string, flt group.Filter) ([]group.Group, error) {
	ret := _m.Called(ctx, principalID, principalType, flt)

	if len(ret) == 0 {
		panic("no return value specified for ListByUser")
	}

	var r0 []group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, group.Filter) ([]group.Group, error)); ok {
		return rf(ctx, principalID, principalType, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, group.Filter) []group.Group); ok {
		r0 = rf(ctx, principalID, principalType, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, group.Filter) error); ok {
		r1 = rf(ctx, principalID, principalType, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupService_ListByUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByUser'
type GroupService_ListByUser_Call struct {
	*mock.Call
}

// ListByUser is a helper method to define mock.On call
//   - ctx context.Context
//   - principalID string
//   - principalType string
//   - flt group.Filter
func (_e *GroupService_Expecter) ListByUser(ctx interface{}, principalID interface{}, principalType interface{}, flt interface{}) *GroupService_ListByUser_Call {
	return &GroupService_ListByUser_Call{Call: _e.mock.On("ListByUser", ctx, principalID, principalType, flt)}
}

func (_c *GroupService_ListByUser_Call) Run(run func(ctx context.Context, principalID string, principalType string, flt group.Filter)) *GroupService_ListByUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(group.Filter))
	})
	return _c
}

func (_c *GroupService_ListByUser_Call) Return(_a0 []group.Group, _a1 error) *GroupService_ListByUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GroupService_ListByUser_Call) RunAndReturn(run func(context.Context, string, string, group.Filter) ([]group.Group, error)) *GroupService_ListByUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewGroupService creates a new instance of GroupService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupService(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupService {
	mock := &GroupService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
