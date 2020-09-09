// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/wasm/tools/wasme/cli/pkg/deploy (interfaces: Provider)

// Package mock_deploy is a generated GoMock package.
package mock_deploy

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/wasm/tools/wasme/cli/operator/api/wasme.io/v1"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// ApplyFilter mocks base method
func (m *MockProvider) ApplyFilter(arg0 *v1.FilterSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyFilter", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyFilter indicates an expected call of ApplyFilter
func (mr *MockProviderMockRecorder) ApplyFilter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFilter", reflect.TypeOf((*MockProvider)(nil).ApplyFilter), arg0)
}

// RemoveFilter mocks base method
func (m *MockProvider) RemoveFilter(arg0 *v1.FilterSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFilter", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFilter indicates an expected call of RemoveFilter
func (mr *MockProviderMockRecorder) RemoveFilter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFilter", reflect.TypeOf((*MockProvider)(nil).RemoveFilter), arg0)
}
