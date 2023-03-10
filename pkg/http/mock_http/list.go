// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/http/list.go

// Package mock_http is a generated GoMock package.
package mock_http

import (
	service "btc/pkg/service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockListService is a mock of ListService interface.
type MockListService struct {
	ctrl     *gomock.Controller
	recorder *MockListServiceMockRecorder
}

// MockListServiceMockRecorder is the mock recorder for MockListService.
type MockListServiceMockRecorder struct {
	mock *MockListService
}

// NewMockListService creates a new mock instance.
func NewMockListService(ctrl *gomock.Controller) *MockListService {
	mock := &MockListService{ctrl: ctrl}
	mock.recorder = &MockListServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListService) EXPECT() *MockListServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockListService) List(params *service.ListServiceParams) ([]*service.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", params)
	ret0, _ := ret[0].([]*service.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockListServiceMockRecorder) List(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockListService)(nil).List), params)
}
