// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nivl/go-filestorage (interfaces: FileStorage)

// Package mockfilestorage is a generated GoMock package.
package mockfilestorage

import (
	filestorage "github.com/Nivl/go-filestorage"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockFileStorage is a mock of FileStorage interface
type MockFileStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageMockRecorder
}

// MockFileStorageMockRecorder is the mock recorder for MockFileStorage
type MockFileStorageMockRecorder struct {
	mock *MockFileStorage
}

// NewMockFileStorage creates a new mock instance
func NewMockFileStorage(ctrl *gomock.Controller) *MockFileStorage {
	mock := &MockFileStorage{ctrl: ctrl}
	mock.recorder = &MockFileStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStorage) EXPECT() *MockFileStorageMockRecorder {
	return m.recorder
}

// Attributes mocks base method
func (m *MockFileStorage) Attributes(arg0 string) (*filestorage.FileAttributes, error) {
	ret := m.ctrl.Call(m, "Attributes", arg0)
	ret0, _ := ret[0].(*filestorage.FileAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attributes indicates an expected call of Attributes
func (mr *MockFileStorageMockRecorder) Attributes(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attributes", reflect.TypeOf((*MockFileStorage)(nil).Attributes), arg0)
}

// Delete mocks base method
func (m *MockFileStorage) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFileStorageMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFileStorage)(nil).Delete), arg0)
}

// Exists mocks base method
func (m *MockFileStorage) Exists(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockFileStorageMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockFileStorage)(nil).Exists), arg0)
}

// ID mocks base method
func (m *MockFileStorage) ID() string {
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockFileStorageMockRecorder) ID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockFileStorage)(nil).ID))
}

// Read mocks base method
func (m *MockFileStorage) Read(arg0 string) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockFileStorageMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileStorage)(nil).Read), arg0)
}

// SetAttributes mocks base method
func (m *MockFileStorage) SetAttributes(arg0 string, arg1 *filestorage.UpdatableFileAttributes) (*filestorage.FileAttributes, error) {
	ret := m.ctrl.Call(m, "SetAttributes", arg0, arg1)
	ret0, _ := ret[0].(*filestorage.FileAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAttributes indicates an expected call of SetAttributes
func (mr *MockFileStorageMockRecorder) SetAttributes(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAttributes", reflect.TypeOf((*MockFileStorage)(nil).SetAttributes), arg0, arg1)
}

// SetBucket mocks base method
func (m *MockFileStorage) SetBucket(arg0 string) error {
	ret := m.ctrl.Call(m, "SetBucket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBucket indicates an expected call of SetBucket
func (mr *MockFileStorageMockRecorder) SetBucket(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBucket", reflect.TypeOf((*MockFileStorage)(nil).SetBucket), arg0)
}

// URL mocks base method
func (m *MockFileStorage) URL(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "URL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URL indicates an expected call of URL
func (mr *MockFileStorageMockRecorder) URL(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockFileStorage)(nil).URL), arg0)
}

// Write mocks base method
func (m *MockFileStorage) Write(arg0 io.Reader, arg1 string) error {
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockFileStorageMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileStorage)(nil).Write), arg0, arg1)
}

// WriteIfNotExist mocks base method
func (m *MockFileStorage) WriteIfNotExist(arg0 io.Reader, arg1 string) (bool, string, error) {
	ret := m.ctrl.Call(m, "WriteIfNotExist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WriteIfNotExist indicates an expected call of WriteIfNotExist
func (mr *MockFileStorageMockRecorder) WriteIfNotExist(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIfNotExist", reflect.TypeOf((*MockFileStorage)(nil).WriteIfNotExist), arg0, arg1)
}
