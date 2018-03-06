// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_machineLearning is a generated GoMock package.
package machineLearning

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// SaveInputRecord mocks base method
func (m *MockStore) SaveInputRecord(reqType string, date time.Time, fields map[string]interface{}) string {
	ret := m.ctrl.Call(m, "SaveInputRecord", reqType, date, fields)
	ret0, _ := ret[0].(string)
	return ret0
}

// SaveInputRecord indicates an expected call of SaveInputRecord
func (mr *MockStoreMockRecorder) SaveInputRecord(reqType, date, fields interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInputRecord", reflect.TypeOf((*MockStore)(nil).SaveInputRecord), reqType, date, fields)
}

// SaveAction mocks base method
func (m *MockStore) SaveAction(requestId, action string, date time.Time, fields map[string]interface{}) {
	m.ctrl.Call(m, "SaveAction", requestId, action, date, fields)
}

// SaveAction indicates an expected call of SaveAction
func (mr *MockStoreMockRecorder) SaveAction(requestId, action, date, fields interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAction", reflect.TypeOf((*MockStore)(nil).SaveAction), requestId, action, date, fields)
}