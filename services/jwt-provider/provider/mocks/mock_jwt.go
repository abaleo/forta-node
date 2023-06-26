// Code generated by MockGen. DO NOT EDIT.
// Source: services/jwt-provider/provider/jwt.go

// Package mock_provider is a generated GoMock package.
package mock_provider

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJWTProvider is a mock of JWTProvider interface.
type MockJWTProvider struct {
	ctrl     *gomock.Controller
	recorder *MockJWTProviderMockRecorder
}

// MockJWTProviderMockRecorder is the mock recorder for MockJWTProvider.
type MockJWTProviderMockRecorder struct {
	mock *MockJWTProvider
}

// NewMockJWTProvider creates a new mock instance.
func NewMockJWTProvider(ctrl *gomock.Controller) *MockJWTProvider {
	mock := &MockJWTProvider{ctrl: ctrl}
	mock.recorder = &MockJWTProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTProvider) EXPECT() *MockJWTProviderMockRecorder {
	return m.recorder
}

// CreateJWT mocks base method.
func (m *MockJWTProvider) CreateJWT(ctx context.Context, ipAddress string, claims map[string]interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWT", ctx, ipAddress, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJWT indicates an expected call of CreateJWT.
func (mr *MockJWTProviderMockRecorder) CreateJWT(ctx, ipAddress, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWT", reflect.TypeOf((*MockJWTProvider)(nil).CreateJWT), ctx, ipAddress, claims)
}
