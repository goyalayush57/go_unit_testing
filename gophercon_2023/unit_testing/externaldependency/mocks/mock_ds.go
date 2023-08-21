// Code generated by MockGen. DO NOT EDIT.
// Source: gophercon_2023/unit_testing/externaldependency (interfaces: DatastoreService)

// Package external is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "gophercon_2023/unit_testing/externaldependency/model"
)

// MockDatastoreService is a mock of DatastoreService interface.
type MockDatastoreService struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreServiceMockRecorder
}

// MockDatastoreServiceMockRecorder is the mock recorder for MockDatastoreService.
type MockDatastoreServiceMockRecorder struct {
	mock *MockDatastoreService
}

// NewMockDatastoreService creates a new mock instance.
func NewMockDatastoreService(ctrl *gomock.Controller) *MockDatastoreService {
	mock := &MockDatastoreService{ctrl: ctrl}
	mock.recorder = &MockDatastoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatastoreService) EXPECT() *MockDatastoreServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDatastoreService) Delete(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
}

// Delete indicates an expected call of Delete.
func (mr *MockDatastoreServiceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatastoreService)(nil).Delete), arg0, arg1, arg2)
}

// Insert mocks base method.
func (m *MockDatastoreService) Insert(arg0 model.Row) (model.Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(model.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockDatastoreServiceMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDatastoreService)(nil).Insert), arg0)
}

// Query mocks base method.
func (m *MockDatastoreService) Query(arg0 string) []model.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].([]model.Row)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockDatastoreServiceMockRecorder) Query(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDatastoreService)(nil).Query), arg0)
}
