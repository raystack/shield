// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	relation "github.com/raystack/frontier/core/relation"
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

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

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

// LookupResources provides a mock function with given fields: ctx, rel
func (_m *RelationService) LookupResources(ctx context.Context, rel relation.Relation) ([]string, error) {
	ret := _m.Called(ctx, rel)

	if len(ret) == 0 {
		panic("no return value specified for LookupResources")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) ([]string, error)); ok {
		return rf(ctx, rel)
	}
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) []string); ok {
		r0 = rf(ctx, rel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, relation.Relation) error); ok {
		r1 = rf(ctx, rel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RelationService_LookupResources_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LookupResources'
type RelationService_LookupResources_Call struct {
	*mock.Call
}

// LookupResources is a helper method to define mock.On call
//   - ctx context.Context
//   - rel relation.Relation
func (_e *RelationService_Expecter) LookupResources(ctx interface{}, rel interface{}) *RelationService_LookupResources_Call {
	return &RelationService_LookupResources_Call{Call: _e.mock.On("LookupResources", ctx, rel)}
}

func (_c *RelationService_LookupResources_Call) Run(run func(ctx context.Context, rel relation.Relation)) *RelationService_LookupResources_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(relation.Relation))
	})
	return _c
}

func (_c *RelationService_LookupResources_Call) Return(_a0 []string, _a1 error) *RelationService_LookupResources_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RelationService_LookupResources_Call) RunAndReturn(run func(context.Context, relation.Relation) ([]string, error)) *RelationService_LookupResources_Call {
	_c.Call.Return(run)
	return _c
}

// LookupSubjects provides a mock function with given fields: ctx, rel
func (_m *RelationService) LookupSubjects(ctx context.Context, rel relation.Relation) ([]string, error) {
	ret := _m.Called(ctx, rel)

	if len(ret) == 0 {
		panic("no return value specified for LookupSubjects")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) ([]string, error)); ok {
		return rf(ctx, rel)
	}
	if rf, ok := ret.Get(0).(func(context.Context, relation.Relation) []string); ok {
		r0 = rf(ctx, rel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, relation.Relation) error); ok {
		r1 = rf(ctx, rel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RelationService_LookupSubjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LookupSubjects'
type RelationService_LookupSubjects_Call struct {
	*mock.Call
}

// LookupSubjects is a helper method to define mock.On call
//   - ctx context.Context
//   - rel relation.Relation
func (_e *RelationService_Expecter) LookupSubjects(ctx interface{}, rel interface{}) *RelationService_LookupSubjects_Call {
	return &RelationService_LookupSubjects_Call{Call: _e.mock.On("LookupSubjects", ctx, rel)}
}

func (_c *RelationService_LookupSubjects_Call) Run(run func(ctx context.Context, rel relation.Relation)) *RelationService_LookupSubjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(relation.Relation))
	})
	return _c
}

func (_c *RelationService_LookupSubjects_Call) Return(_a0 []string, _a1 error) *RelationService_LookupSubjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RelationService_LookupSubjects_Call) RunAndReturn(run func(context.Context, relation.Relation) ([]string, error)) *RelationService_LookupSubjects_Call {
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
