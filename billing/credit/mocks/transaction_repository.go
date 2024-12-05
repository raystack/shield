// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	credit "github.com/raystack/frontier/billing/credit"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

type TransactionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TransactionRepository) EXPECT() *TransactionRepository_Expecter {
	return &TransactionRepository_Expecter{mock: &_m.Mock}
}

// CreateEntry provides a mock function with given fields: ctx, debit, _a2
func (_m *TransactionRepository) CreateEntry(ctx context.Context, debit credit.Transaction, _a2 credit.Transaction) ([]credit.Transaction, error) {
	ret := _m.Called(ctx, debit, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateEntry")
	}

	var r0 []credit.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, credit.Transaction, credit.Transaction) ([]credit.Transaction, error)); ok {
		return rf(ctx, debit, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, credit.Transaction, credit.Transaction) []credit.Transaction); ok {
		r0 = rf(ctx, debit, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]credit.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, credit.Transaction, credit.Transaction) error); ok {
		r1 = rf(ctx, debit, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionRepository_CreateEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEntry'
type TransactionRepository_CreateEntry_Call struct {
	*mock.Call
}

// CreateEntry is a helper method to define mock.On call
//   - ctx context.Context
//   - debit credit.Transaction
//   - _a2 credit.Transaction
func (_e *TransactionRepository_Expecter) CreateEntry(ctx interface{}, debit interface{}, _a2 interface{}) *TransactionRepository_CreateEntry_Call {
	return &TransactionRepository_CreateEntry_Call{Call: _e.mock.On("CreateEntry", ctx, debit, _a2)}
}

func (_c *TransactionRepository_CreateEntry_Call) Run(run func(ctx context.Context, debit credit.Transaction, _a2 credit.Transaction)) *TransactionRepository_CreateEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(credit.Transaction), args[2].(credit.Transaction))
	})
	return _c
}

func (_c *TransactionRepository_CreateEntry_Call) Return(_a0 []credit.Transaction, _a1 error) *TransactionRepository_CreateEntry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransactionRepository_CreateEntry_Call) RunAndReturn(run func(context.Context, credit.Transaction, credit.Transaction) ([]credit.Transaction, error)) *TransactionRepository_CreateEntry_Call {
	_c.Call.Return(run)
	return _c
}

// GetBalance provides a mock function with given fields: ctx, id
func (_m *TransactionRepository) GetBalance(ctx context.Context, id string) (int64, error) {
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

// TransactionRepository_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type TransactionRepository_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *TransactionRepository_Expecter) GetBalance(ctx interface{}, id interface{}) *TransactionRepository_GetBalance_Call {
	return &TransactionRepository_GetBalance_Call{Call: _e.mock.On("GetBalance", ctx, id)}
}

func (_c *TransactionRepository_GetBalance_Call) Run(run func(ctx context.Context, id string)) *TransactionRepository_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TransactionRepository_GetBalance_Call) Return(_a0 int64, _a1 error) *TransactionRepository_GetBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransactionRepository_GetBalance_Call) RunAndReturn(run func(context.Context, string) (int64, error)) *TransactionRepository_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetBalanceForRange provides a mock function with given fields: ctx, accountID, start, end
func (_m *TransactionRepository) GetBalanceForRange(ctx context.Context, accountID string, start time.Time, end time.Time) (int64, error) {
	ret := _m.Called(ctx, accountID, start, end)

	if len(ret) == 0 {
		panic("no return value specified for GetBalanceForRange")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, time.Time) (int64, error)); ok {
		return rf(ctx, accountID, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, time.Time) int64); ok {
		r0 = rf(ctx, accountID, start, end)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time, time.Time) error); ok {
		r1 = rf(ctx, accountID, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionRepository_GetBalanceForRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalanceForRange'
type TransactionRepository_GetBalanceForRange_Call struct {
	*mock.Call
}

// GetBalanceForRange is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID string
//   - start time.Time
//   - end time.Time
func (_e *TransactionRepository_Expecter) GetBalanceForRange(ctx interface{}, accountID interface{}, start interface{}, end interface{}) *TransactionRepository_GetBalanceForRange_Call {
	return &TransactionRepository_GetBalanceForRange_Call{Call: _e.mock.On("GetBalanceForRange", ctx, accountID, start, end)}
}

func (_c *TransactionRepository_GetBalanceForRange_Call) Run(run func(ctx context.Context, accountID string, start time.Time, end time.Time)) *TransactionRepository_GetBalanceForRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Time), args[3].(time.Time))
	})
	return _c
}

func (_c *TransactionRepository_GetBalanceForRange_Call) Return(_a0 int64, _a1 error) *TransactionRepository_GetBalanceForRange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransactionRepository_GetBalanceForRange_Call) RunAndReturn(run func(context.Context, string, time.Time, time.Time) (int64, error)) *TransactionRepository_GetBalanceForRange_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TransactionRepository) GetByID(ctx context.Context, id string) (credit.Transaction, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 credit.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (credit.Transaction, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) credit.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(credit.Transaction)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type TransactionRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *TransactionRepository_Expecter) GetByID(ctx interface{}, id interface{}) *TransactionRepository_GetByID_Call {
	return &TransactionRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *TransactionRepository_GetByID_Call) Run(run func(ctx context.Context, id string)) *TransactionRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TransactionRepository_GetByID_Call) Return(_a0 credit.Transaction, _a1 error) *TransactionRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransactionRepository_GetByID_Call) RunAndReturn(run func(context.Context, string) (credit.Transaction, error)) *TransactionRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, flt
func (_m *TransactionRepository) List(ctx context.Context, flt credit.Filter) ([]credit.Transaction, error) {
	ret := _m.Called(ctx, flt)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []credit.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, credit.Filter) ([]credit.Transaction, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, credit.Filter) []credit.Transaction); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]credit.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, credit.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type TransactionRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - flt credit.Filter
func (_e *TransactionRepository_Expecter) List(ctx interface{}, flt interface{}) *TransactionRepository_List_Call {
	return &TransactionRepository_List_Call{Call: _e.mock.On("List", ctx, flt)}
}

func (_c *TransactionRepository_List_Call) Run(run func(ctx context.Context, flt credit.Filter)) *TransactionRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(credit.Filter))
	})
	return _c
}

func (_c *TransactionRepository_List_Call) Return(_a0 []credit.Transaction, _a1 error) *TransactionRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransactionRepository_List_Call) RunAndReturn(run func(context.Context, credit.Filter) ([]credit.Transaction, error)) *TransactionRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
