// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	authenticate "github.com/raystack/frontier/core/authenticate"

	jwk "github.com/lestrrat-go/jwx/v2/jwk"

	mock "github.com/stretchr/testify/mock"
)

// AuthnService is an autogenerated mock type for the AuthnService type
type AuthnService struct {
	mock.Mock
}

type AuthnService_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthnService) EXPECT() *AuthnService_Expecter {
	return &AuthnService_Expecter{mock: &_m.Mock}
}

// BuildToken provides a mock function with given fields: ctx, principalID, metadata
func (_m *AuthnService) BuildToken(ctx context.Context, principalID string, metadata map[string]string) ([]byte, error) {
	ret := _m.Called(ctx, principalID, metadata)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) ([]byte, error)); ok {
		return rf(ctx, principalID, metadata)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) []byte); ok {
		r0 = rf(ctx, principalID, metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, principalID, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthnService_BuildToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BuildToken'
type AuthnService_BuildToken_Call struct {
	*mock.Call
}

// BuildToken is a helper method to define mock.On call
//   - ctx context.Context
//   - principalID string
//   - metadata map[string]string
func (_e *AuthnService_Expecter) BuildToken(ctx interface{}, principalID interface{}, metadata interface{}) *AuthnService_BuildToken_Call {
	return &AuthnService_BuildToken_Call{Call: _e.mock.On("BuildToken", ctx, principalID, metadata)}
}

func (_c *AuthnService_BuildToken_Call) Run(run func(ctx context.Context, principalID string, metadata map[string]string)) *AuthnService_BuildToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(map[string]string))
	})
	return _c
}

func (_c *AuthnService_BuildToken_Call) Return(_a0 []byte, _a1 error) *AuthnService_BuildToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthnService_BuildToken_Call) RunAndReturn(run func(context.Context, string, map[string]string) ([]byte, error)) *AuthnService_BuildToken_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *AuthnService) Close() {
	_m.Called()
}

// AuthnService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type AuthnService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *AuthnService_Expecter) Close() *AuthnService_Close_Call {
	return &AuthnService_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *AuthnService_Close_Call) Run(run func()) *AuthnService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AuthnService_Close_Call) Return() *AuthnService_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *AuthnService_Close_Call) RunAndReturn(run func()) *AuthnService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// FinishFlow provides a mock function with given fields: ctx, request
func (_m *AuthnService) FinishFlow(ctx context.Context, request authenticate.RegistrationFinishRequest) (*authenticate.RegistrationFinishResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *authenticate.RegistrationFinishResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authenticate.RegistrationFinishRequest) (*authenticate.RegistrationFinishResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authenticate.RegistrationFinishRequest) *authenticate.RegistrationFinishResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authenticate.RegistrationFinishResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authenticate.RegistrationFinishRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthnService_FinishFlow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FinishFlow'
type AuthnService_FinishFlow_Call struct {
	*mock.Call
}

// FinishFlow is a helper method to define mock.On call
//   - ctx context.Context
//   - request authenticate.RegistrationFinishRequest
func (_e *AuthnService_Expecter) FinishFlow(ctx interface{}, request interface{}) *AuthnService_FinishFlow_Call {
	return &AuthnService_FinishFlow_Call{Call: _e.mock.On("FinishFlow", ctx, request)}
}

func (_c *AuthnService_FinishFlow_Call) Run(run func(ctx context.Context, request authenticate.RegistrationFinishRequest)) *AuthnService_FinishFlow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(authenticate.RegistrationFinishRequest))
	})
	return _c
}

func (_c *AuthnService_FinishFlow_Call) Return(_a0 *authenticate.RegistrationFinishResponse, _a1 error) *AuthnService_FinishFlow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthnService_FinishFlow_Call) RunAndReturn(run func(context.Context, authenticate.RegistrationFinishRequest) (*authenticate.RegistrationFinishResponse, error)) *AuthnService_FinishFlow_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrincipal provides a mock function with given fields: ctx, via
func (_m *AuthnService) GetPrincipal(ctx context.Context, via ...authenticate.ClientAssertion) (authenticate.Principal, error) {
	_va := make([]interface{}, len(via))
	for _i := range via {
		_va[_i] = via[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 authenticate.Principal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...authenticate.ClientAssertion) (authenticate.Principal, error)); ok {
		return rf(ctx, via...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...authenticate.ClientAssertion) authenticate.Principal); ok {
		r0 = rf(ctx, via...)
	} else {
		r0 = ret.Get(0).(authenticate.Principal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...authenticate.ClientAssertion) error); ok {
		r1 = rf(ctx, via...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthnService_GetPrincipal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrincipal'
type AuthnService_GetPrincipal_Call struct {
	*mock.Call
}

// GetPrincipal is a helper method to define mock.On call
//   - ctx context.Context
//   - via ...authenticate.ClientAssertion
func (_e *AuthnService_Expecter) GetPrincipal(ctx interface{}, via ...interface{}) *AuthnService_GetPrincipal_Call {
	return &AuthnService_GetPrincipal_Call{Call: _e.mock.On("GetPrincipal",
		append([]interface{}{ctx}, via...)...)}
}

func (_c *AuthnService_GetPrincipal_Call) Run(run func(ctx context.Context, via ...authenticate.ClientAssertion)) *AuthnService_GetPrincipal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]authenticate.ClientAssertion, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(authenticate.ClientAssertion)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *AuthnService_GetPrincipal_Call) Return(_a0 authenticate.Principal, _a1 error) *AuthnService_GetPrincipal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthnService_GetPrincipal_Call) RunAndReturn(run func(context.Context, ...authenticate.ClientAssertion) (authenticate.Principal, error)) *AuthnService_GetPrincipal_Call {
	_c.Call.Return(run)
	return _c
}

// InitFlows provides a mock function with given fields: ctx
func (_m *AuthnService) InitFlows(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthnService_InitFlows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitFlows'
type AuthnService_InitFlows_Call struct {
	*mock.Call
}

// InitFlows is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AuthnService_Expecter) InitFlows(ctx interface{}) *AuthnService_InitFlows_Call {
	return &AuthnService_InitFlows_Call{Call: _e.mock.On("InitFlows", ctx)}
}

func (_c *AuthnService_InitFlows_Call) Run(run func(ctx context.Context)) *AuthnService_InitFlows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AuthnService_InitFlows_Call) Return(_a0 error) *AuthnService_InitFlows_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthnService_InitFlows_Call) RunAndReturn(run func(context.Context) error) *AuthnService_InitFlows_Call {
	_c.Call.Return(run)
	return _c
}

// JWKs provides a mock function with given fields: ctx
func (_m *AuthnService) JWKs(ctx context.Context) jwk.Set {
	ret := _m.Called(ctx)

	var r0 jwk.Set
	if rf, ok := ret.Get(0).(func(context.Context) jwk.Set); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwk.Set)
		}
	}

	return r0
}

// AuthnService_JWKs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'JWKs'
type AuthnService_JWKs_Call struct {
	*mock.Call
}

// JWKs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AuthnService_Expecter) JWKs(ctx interface{}) *AuthnService_JWKs_Call {
	return &AuthnService_JWKs_Call{Call: _e.mock.On("JWKs", ctx)}
}

func (_c *AuthnService_JWKs_Call) Run(run func(ctx context.Context)) *AuthnService_JWKs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AuthnService_JWKs_Call) Return(_a0 jwk.Set) *AuthnService_JWKs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthnService_JWKs_Call) RunAndReturn(run func(context.Context) jwk.Set) *AuthnService_JWKs_Call {
	_c.Call.Return(run)
	return _c
}

// StartFlow provides a mock function with given fields: ctx, request
func (_m *AuthnService) StartFlow(ctx context.Context, request authenticate.RegistrationStartRequest) (*authenticate.RegistrationStartResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *authenticate.RegistrationStartResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authenticate.RegistrationStartRequest) (*authenticate.RegistrationStartResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authenticate.RegistrationStartRequest) *authenticate.RegistrationStartResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authenticate.RegistrationStartResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authenticate.RegistrationStartRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthnService_StartFlow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartFlow'
type AuthnService_StartFlow_Call struct {
	*mock.Call
}

// StartFlow is a helper method to define mock.On call
//   - ctx context.Context
//   - request authenticate.RegistrationStartRequest
func (_e *AuthnService_Expecter) StartFlow(ctx interface{}, request interface{}) *AuthnService_StartFlow_Call {
	return &AuthnService_StartFlow_Call{Call: _e.mock.On("StartFlow", ctx, request)}
}

func (_c *AuthnService_StartFlow_Call) Run(run func(ctx context.Context, request authenticate.RegistrationStartRequest)) *AuthnService_StartFlow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(authenticate.RegistrationStartRequest))
	})
	return _c
}

func (_c *AuthnService_StartFlow_Call) Return(_a0 *authenticate.RegistrationStartResponse, _a1 error) *AuthnService_StartFlow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthnService_StartFlow_Call) RunAndReturn(run func(context.Context, authenticate.RegistrationStartRequest) (*authenticate.RegistrationStartResponse, error)) *AuthnService_StartFlow_Call {
	_c.Call.Return(run)
	return _c
}

// SupportedStrategies provides a mock function with given fields:
func (_m *AuthnService) SupportedStrategies() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// AuthnService_SupportedStrategies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SupportedStrategies'
type AuthnService_SupportedStrategies_Call struct {
	*mock.Call
}

// SupportedStrategies is a helper method to define mock.On call
func (_e *AuthnService_Expecter) SupportedStrategies() *AuthnService_SupportedStrategies_Call {
	return &AuthnService_SupportedStrategies_Call{Call: _e.mock.On("SupportedStrategies")}
}

func (_c *AuthnService_SupportedStrategies_Call) Run(run func()) *AuthnService_SupportedStrategies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AuthnService_SupportedStrategies_Call) Return(_a0 []string) *AuthnService_SupportedStrategies_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthnService_SupportedStrategies_Call) RunAndReturn(run func() []string) *AuthnService_SupportedStrategies_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthnService creates a new instance of AuthnService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthnService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthnService {
	mock := &AuthnService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
