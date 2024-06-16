// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dcalsky/kong_service_demo/internal/adapter/repo (interfaces: IKongServiceRepo)
//
// Generated by this command:
//
//	mockgen -package repo --build_flags=--mod=mod --destination kong_service_mock.go . IKongServiceRepo
//

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	reflect "reflect"

	dto "github.com/dcalsky/kong_service_demo/internal/model/dto"
	entity "github.com/dcalsky/kong_service_demo/internal/model/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIKongServiceRepo is a mock of IKongServiceRepo interface.
type MockIKongServiceRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIKongServiceRepoMockRecorder
}

// MockIKongServiceRepoMockRecorder is the mock recorder for MockIKongServiceRepo.
type MockIKongServiceRepoMockRecorder struct {
	mock *MockIKongServiceRepo
}

// NewMockIKongServiceRepo creates a new mock instance.
func NewMockIKongServiceRepo(ctrl *gomock.Controller) *MockIKongServiceRepo {
	mock := &MockIKongServiceRepo{ctrl: ctrl}
	mock.recorder = &MockIKongServiceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIKongServiceRepo) EXPECT() *MockIKongServiceRepoMockRecorder {
	return m.recorder
}

// CountServicesVersionAmount mocks base method.
func (m *MockIKongServiceRepo) CountServicesVersionAmount(arg0 context.Context, arg1 []entity.KongServiceId) (map[entity.KongServiceId]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountServicesVersionAmount", arg0, arg1)
	ret0, _ := ret[0].(map[entity.KongServiceId]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountServicesVersionAmount indicates an expected call of CountServicesVersionAmount.
func (mr *MockIKongServiceRepoMockRecorder) CountServicesVersionAmount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountServicesVersionAmount", reflect.TypeOf((*MockIKongServiceRepo)(nil).CountServicesVersionAmount), arg0, arg1)
}

// CreateService mocks base method.
func (m *MockIKongServiceRepo) CreateService(arg0 context.Context, arg1 *entity.KongService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService.
func (mr *MockIKongServiceRepoMockRecorder) CreateService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockIKongServiceRepo)(nil).CreateService), arg0, arg1)
}

// CreateServiceVersion mocks base method.
func (m *MockIKongServiceRepo) CreateServiceVersion(arg0 context.Context, arg1 *entity.KongServiceVersion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceVersion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateServiceVersion indicates an expected call of CreateServiceVersion.
func (mr *MockIKongServiceRepoMockRecorder) CreateServiceVersion(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceVersion", reflect.TypeOf((*MockIKongServiceRepo)(nil).CreateServiceVersion), arg0, arg1)
}

// DeleteService mocks base method.
func (m *MockIKongServiceRepo) DeleteService(arg0 context.Context, arg1 entity.KongServiceId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockIKongServiceRepoMockRecorder) DeleteService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockIKongServiceRepo)(nil).DeleteService), arg0, arg1)
}

// DescribeService mocks base method.
func (m *MockIKongServiceRepo) DescribeService(arg0 context.Context, arg1 entity.KongServiceId) (*entity.KongService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeService", arg0, arg1)
	ret0, _ := ret[0].(*entity.KongService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeService indicates an expected call of DescribeService.
func (mr *MockIKongServiceRepoMockRecorder) DescribeService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeService", reflect.TypeOf((*MockIKongServiceRepo)(nil).DescribeService), arg0, arg1)
}

// DescribeServiceVersion mocks base method.
func (m *MockIKongServiceRepo) DescribeServiceVersion(arg0 context.Context, arg1 entity.KongServiceVersionId) (*entity.KongServiceVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeServiceVersion", arg0, arg1)
	ret0, _ := ret[0].(*entity.KongServiceVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServiceVersion indicates an expected call of DescribeServiceVersion.
func (mr *MockIKongServiceRepoMockRecorder) DescribeServiceVersion(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServiceVersion", reflect.TypeOf((*MockIKongServiceRepo)(nil).DescribeServiceVersion), arg0, arg1)
}

// ListServices mocks base method.
func (m *MockIKongServiceRepo) ListServices(arg0 context.Context, arg1 ListServicesRequest) ([]entity.KongService, dto.PagingResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0, arg1)
	ret0, _ := ret[0].([]entity.KongService)
	ret1, _ := ret[1].(dto.PagingResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListServices indicates an expected call of ListServices.
func (mr *MockIKongServiceRepoMockRecorder) ListServices(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockIKongServiceRepo)(nil).ListServices), arg0, arg1)
}

// ListVersionsByServiceId mocks base method.
func (m *MockIKongServiceRepo) ListVersionsByServiceId(arg0 context.Context, arg1 entity.KongServiceId) ([]entity.KongServiceVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVersionsByServiceId", arg0, arg1)
	ret0, _ := ret[0].([]entity.KongServiceVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVersionsByServiceId indicates an expected call of ListVersionsByServiceId.
func (mr *MockIKongServiceRepoMockRecorder) ListVersionsByServiceId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVersionsByServiceId", reflect.TypeOf((*MockIKongServiceRepo)(nil).ListVersionsByServiceId), arg0, arg1)
}

// ReplaceService mocks base method.
func (m *MockIKongServiceRepo) ReplaceService(arg0 context.Context, arg1 entity.KongService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplaceService indicates an expected call of ReplaceService.
func (mr *MockIKongServiceRepoMockRecorder) ReplaceService(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceService", reflect.TypeOf((*MockIKongServiceRepo)(nil).ReplaceService), arg0, arg1)
}
