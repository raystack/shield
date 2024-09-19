// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CreditService is an autogenerated mock type for the CreditService type
type CreditService struct {
	mock.Mock
}

type CreditService_Expecter struct {
	mock *mock.Mock
}

func (_m *CreditService) EXPECT() *CreditService_Expecter {
	return &CreditService_Expecter{mock: &_m.Mock}
}

// GetBalance provides a mock function with given fields: ctx, id
func (_m *CreditService) GetBalance(ctx context.Context, id string) (int64, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreditService_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type CreditService_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *CreditService_Expecter) GetBalance(ctx interface{}, id interface{}) *CreditService_GetBalance_Call {
	return &CreditService_GetBalance_Call{Call: _e.mock.On("GetBalance", ctx, id)}
}

func (_c *CreditService_GetBalance_Call) Run(run func(ctx context.Context, id string)) *CreditService_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CreditService_GetBalance_Call) Return(_a0 int64, _a1 error) *CreditService_GetBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CreditService_GetBalance_Call) RunAndReturn(run func(context.Context, string) (int64, error)) *CreditService_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewCreditService creates a new instance of CreditService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreditService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreditService {
	mock := &CreditService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
