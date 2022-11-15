// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ce_client/ce_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockContainerEngineClient is a mock of ContainerEngineClient interface.
type MockContainerEngineClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerEngineClientMockRecorder
}

// MockContainerEngineClientMockRecorder is the mock recorder for MockContainerEngineClient.
type MockContainerEngineClientMockRecorder struct {
	mock *MockContainerEngineClient
}

// NewMockContainerEngineClient creates a new mock instance.
func NewMockContainerEngineClient(ctrl *gomock.Controller) *MockContainerEngineClient {
	mock := &MockContainerEngineClient{ctrl: ctrl}
	mock.recorder = &MockContainerEngineClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerEngineClient) EXPECT() *MockContainerEngineClientMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockContainerEngineClient) Build(filepath, name string, tags []string, quay_img_exp string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", filepath, name, tags)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockContainerEngineClientMockRecorder) Build(filepath, name, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockContainerEngineClient)(nil).Build), filepath, name, tags)
}

// Push mocks base method.
func (m *MockContainerEngineClient) Push(destination, username, password, registry_address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", destination, username, password, registry_address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockContainerEngineClientMockRecorder) Push(destination, username, password, registry_address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockContainerEngineClient)(nil).Push), destination, username, password, registry_address)
}

// Tag mocks base method.
func (m *MockContainerEngineClient) Tag(image_tag, destination string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag", image_tag, destination)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockContainerEngineClientMockRecorder) Tag(image_tag, destination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockContainerEngineClient)(nil).Tag), image_tag, destination)
}
