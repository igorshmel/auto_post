// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	port "auto_post/app/internal/adapters/port"
	dbo "auto_post/app/pkg/dbo"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExtractor is a mock of Extractor interface.
type MockExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockExtractorMockRecorder
}

// MockExtractorMockRecorder is the mock recorder for MockExtractor.
type MockExtractorMockRecorder struct {
	mock *MockExtractor
}

// NewMockExtractor creates a new mock instance.
func NewMockExtractor(ctrl *gomock.Controller) *MockExtractor {
	mock := &MockExtractor{ctrl: ctrl}
	mock.recorder = &MockExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtractor) EXPECT() *MockExtractorMockRecorder {
	return m.recorder
}

// GetByUUID mocks base method.
func (m *MockExtractor) GetByUUID(fileDBO *dbo.ManagerDBO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", fileDBO)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByUUID indicates an expected call of GetByUUID.
func (mr *MockExtractorMockRecorder) GetByUUID(fileDBO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockExtractor)(nil).GetByUUID), fileDBO)
}

// MockPersister is a mock of Persister interface.
type MockPersister struct {
	ctrl     *gomock.Controller
	recorder *MockPersisterMockRecorder
}

// MockPersisterMockRecorder is the mock recorder for MockPersister.
type MockPersisterMockRecorder struct {
	mock *MockPersister
}

// NewMockPersister creates a new mock instance.
func NewMockPersister(ctrl *gomock.Controller) *MockPersister {
	mock := &MockPersister{ctrl: ctrl}
	mock.recorder = &MockPersisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersister) EXPECT() *MockPersisterMockRecorder {
	return m.recorder
}

// SaveNewFile mocks base method.
func (m *MockPersister) CreateRecord(fileDBO *dbo.ManagerDBO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", fileDBO)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNewFile indicates an expected call of SaveNewFile.
func (mr *MockPersisterMockRecorder) SaveNewFile(fileDBO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockPersister)(nil).CreateRecord), fileDBO)
}

// UnitOfWork mocks base method.
func (m *MockPersister) UnitOfWork(arg0 func(port.Persister) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitOfWork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnitOfWork indicates an expected call of UnitOfWork.
func (mr *MockPersisterMockRecorder) UnitOfWork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitOfWork", reflect.TypeOf((*MockPersister)(nil).UnitOfWork), arg0)
}

// UpdateFileStatus mocks base method.
func (m *MockPersister) UpdateRecordStatus(fileDBO *dbo.ManagerDBO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecordStatus", fileDBO)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFileStatus indicates an expected call of UpdateFileStatus.
func (mr *MockPersisterMockRecorder) UpdateFileStatus(fileDBO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecordStatus", reflect.TypeOf((*MockPersister)(nil).UpdateRecordStatus), fileDBO)
}
