// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	relation "github.com/raystack/frontier/core/relation"
	mock "github.com/stretchr/testify/mock"
)

// RelationService is an autogenerated mock type for the RelationService type
type RelationService struct {
	mock.Mock
}

type RelationService_Expecter struct {
	mock *mock.Mock
}

func (_m *RelationService) EXPECT() *RelationService_Expecter {
	return &RelationService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, rel
func (_m *RelationService) Create(ctx context.Context, rel relation.Relation) (relation.Relation, error) {
	ret := _m.Called(ctx, rel)

	var r0 relation.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) (relation.Relation, error)); ok {
		return rf(ctx, rel)
	}
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) relation.Relation); ok {
		r0 = rf(ctx, rel)
	} else {
		r0 = ret.Get(0).(relation.Relation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, relation.Relation) error); ok {
		r1 = rf(ctx, rel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RelationService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type RelationService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - rel relation.Relation
func (_e *RelationService_Expecter) Create(ctx interface{}, rel interface{}) *RelationService_Create_Call {
	return &RelationService_Create_Call{Call: _e.mock.On("Create", ctx, rel)}
}

func (_c *RelationService_Create_Call) Run(run func(ctx context.Context, rel relation.Relation)) *RelationService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(relation.Relation))
	})
	return _c
}

func (_c *RelationService_Create_Call) Return(_a0 relation.Relation, _a1 error) *RelationService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RelationService_Create_Call) RunAndReturn(run func(context.Context, relation.Relation) (relation.Relation, error)) *RelationService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, rel
func (_m *RelationService) Delete(ctx context.Context, rel relation.Relation) error {
	ret := _m.Called(ctx, rel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) error); ok {
		r0 = rf(ctx, rel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RelationService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type RelationService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - rel relation.Relation
func (_e *RelationService_Expecter) Delete(ctx interface{}, rel interface{}) *RelationService_Delete_Call {
	return &RelationService_Delete_Call{Call: _e.mock.On("Delete", ctx, rel)}
}

func (_c *RelationService_Delete_Call) Run(run func(ctx context.Context, rel relation.Relation)) *RelationService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(relation.Relation))
	})
	return _c
}

func (_c *RelationService_Delete_Call) Return(_a0 error) *RelationService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RelationService_Delete_Call) RunAndReturn(run func(context.Context, relation.Relation) error) *RelationService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *RelationService) Get(ctx context.Context, id string) (relation.Relation, error) {
	ret := _m.Called(ctx, id)

	var r0 relation.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (relation.Relation, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) relation.Relation); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(relation.Relation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RelationService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RelationService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *RelationService_Expecter) Get(ctx interface{}, id interface{}) *RelationService_Get_Call {
	return &RelationService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *RelationService_Get_Call) Run(run func(ctx context.Context, id string)) *RelationService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RelationService_Get_Call) Return(_a0 relation.Relation, _a1 error) *RelationService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RelationService_Get_Call) RunAndReturn(run func(context.Context, string) (relation.Relation, error)) *RelationService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *RelationService) List(ctx context.Context) ([]relation.Relation, error) {
	ret := _m.Called(ctx)

	var r0 []relation.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]relation.Relation, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []relation.Relation); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]relation.Relation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RelationService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type RelationService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RelationService_Expecter) List(ctx interface{}) *RelationService_List_Call {
	return &RelationService_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *RelationService_List_Call) Run(run func(ctx context.Context)) *RelationService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RelationService_List_Call) Return(_a0 []relation.Relation, _a1 error) *RelationService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RelationService_List_Call) RunAndReturn(run func(context.Context) ([]relation.Relation, error)) *RelationService_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewRelationService creates a new instance of RelationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRelationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *RelationService {
	mock := &RelationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
