// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dcalsky/kong_service_demo/internal/service/kong_service (interfaces: IKongService)
//
// Generated by this command:
//
//	mockgen -package kong_service --build_flags=--mod=mod --destination service_mock.go . IKongService
//

// Package kong_service is a generated GoMock package.
package kong_service

import (
	context "context"
	reflect "reflect"

	base "github.com/dcalsky/kong_service_demo/internal/base"
	dto "github.com/dcalsky/kong_service_demo/internal/model/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockIKongService is a mock of IKongService interface.
type MockIKongService struct {
	ctrl     *gomock.Controller
	recorder *MockIKongServiceMockRecorder
}

// MockIKongServiceMockRecorder is the mock recorder for MockIKongService.
type MockIKongServiceMockRecorder struct {
	mock *MockIKongService
}

// NewMockIKongService creates a new mock instance.
func NewMockIKongService(ctrl *gomock.Controller) *MockIKongService {
	mock := &MockIKongService{ctrl: ctrl}
	mock.recorder = &MockIKongServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIKongService) EXPECT() *MockIKongServiceMockRecorder {
	return m.recorder
}

// CreateKongService mocks base method.
func (m *MockIKongService) CreateKongService(arg0 context.Context, arg1 base.KongArgs, arg2 dto.CreateKongServiceRequest) dto.CreateKongServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKongService", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.CreateKongServiceResponse)
	return ret0
}

// CreateKongService indicates an expected call of CreateKongService.
func (mr *MockIKongServiceMockRecorder) CreateKongService(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKongService", reflect.TypeOf((*MockIKongService)(nil).CreateKongService), arg0, arg1, arg2)
}

// CreateKongServiceVersion mocks base method.
func (m *MockIKongService) CreateKongServiceVersion(arg0 context.Context, arg1 base.KongArgs, arg2 dto.CreateKongServiceVersionRequest) dto.CreateKongServiceVersionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKongServiceVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.CreateKongServiceVersionResponse)
	return ret0
}

// CreateKongServiceVersion indicates an expected call of CreateKongServiceVersion.
func (mr *MockIKongServiceMockRecorder) CreateKongServiceVersion(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKongServiceVersion", reflect.TypeOf((*MockIKongService)(nil).CreateKongServiceVersion), arg0, arg1, arg2)
}

// DeleteKongService mocks base method.
func (m *MockIKongService) DeleteKongService(arg0 context.Context, arg1 base.KongArgs, arg2 dto.DeleteKongServiceRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteKongService", arg0, arg1, arg2)
}

// DeleteKongService indicates an expected call of DeleteKongService.
func (mr *MockIKongServiceMockRecorder) DeleteKongService(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKongService", reflect.TypeOf((*MockIKongService)(nil).DeleteKongService), arg0, arg1, arg2)
}

// DescribeKongService mocks base method.
func (m *MockIKongService) DescribeKongService(arg0 context.Context, arg1 base.KongArgs, arg2 dto.DescribeKongServiceRequest) dto.DescribeKongServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeKongService", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.DescribeKongServiceResponse)
	return ret0
}

// DescribeKongService indicates an expected call of DescribeKongService.
func (mr *MockIKongServiceMockRecorder) DescribeKongService(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeKongService", reflect.TypeOf((*MockIKongService)(nil).DescribeKongService), arg0, arg1, arg2)
}

// ListKongServices mocks base method.
func (m *MockIKongService) ListKongServices(arg0 context.Context, arg1 base.KongArgs, arg2 dto.ListKongServicesRequest) dto.ListKongServicesResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKongServices", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.ListKongServicesResponse)
	return ret0
}

// ListKongServices indicates an expected call of ListKongServices.
func (mr *MockIKongServiceMockRecorder) ListKongServices(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKongServices", reflect.TypeOf((*MockIKongService)(nil).ListKongServices), arg0, arg1, arg2)
}

// SwitchKongServiceVersion mocks base method.
func (m *MockIKongService) SwitchKongServiceVersion(arg0 context.Context, arg1 base.KongArgs, arg2 dto.SwitchKongServiceVersionRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SwitchKongServiceVersion", arg0, arg1, arg2)
}

// SwitchKongServiceVersion indicates an expected call of SwitchKongServiceVersion.
func (mr *MockIKongServiceMockRecorder) SwitchKongServiceVersion(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchKongServiceVersion", reflect.TypeOf((*MockIKongService)(nil).SwitchKongServiceVersion), arg0, arg1, arg2)
}

// UpdateKongService mocks base method.
func (m *MockIKongService) UpdateKongService(arg0 context.Context, arg1 base.KongArgs, arg2 dto.UpdateKongServiceRequest) dto.UpdateKongServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKongService", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.UpdateKongServiceResponse)
	return ret0
}

// UpdateKongService indicates an expected call of UpdateKongService.
func (mr *MockIKongServiceMockRecorder) UpdateKongService(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKongService", reflect.TypeOf((*MockIKongService)(nil).UpdateKongService), arg0, arg1, arg2)
}
