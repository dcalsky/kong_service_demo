// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dcalsky/kong_service_demo/internal/service/session (interfaces: ISessionService)
//
// Generated by this command:
//
//	mockgen -package session --build_flags=--mod=mod --destination service_mock.go . ISessionService
//

// Package session is a generated GoMock package.
package session

import (
	context "context"
	reflect "reflect"

	dto "github.com/dcalsky/kong_service_demo/internal/model/dto"
	entity "github.com/dcalsky/kong_service_demo/internal/model/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockISessionService is a mock of ISessionService interface.
type MockISessionService struct {
	ctrl     *gomock.Controller
	recorder *MockISessionServiceMockRecorder
}

// MockISessionServiceMockRecorder is the mock recorder for MockISessionService.
type MockISessionServiceMockRecorder struct {
	mock *MockISessionService
}

// NewMockISessionService creates a new mock instance.
func NewMockISessionService(ctrl *gomock.Controller) *MockISessionService {
	mock := &MockISessionService{ctrl: ctrl}
	mock.recorder = &MockISessionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISessionService) EXPECT() *MockISessionServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockISessionService) Login(arg0 context.Context, arg1 dto.LoginRequest) dto.LoginResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(dto.LoginResponse)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockISessionServiceMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockISessionService)(nil).Login), arg0, arg1)
}

// Register mocks base method.
func (m *MockISessionService) Register(arg0 context.Context, arg1 dto.RegisterRequest) dto.RegisterResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(dto.RegisterResponse)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockISessionServiceMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockISessionService)(nil).Register), arg0, arg1)
}

// generateJwt mocks base method.
func (m *MockISessionService) generateJwt(arg0 entity.Account) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateJwt", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// generateJwt indicates an expected call of generateJwt.
func (mr *MockISessionServiceMockRecorder) generateJwt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateJwt", reflect.TypeOf((*MockISessionService)(nil).generateJwt), arg0)
}
