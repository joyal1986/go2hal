// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_alert is a generated GoMock package.
package alert

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// SendAlert mocks base method
func (m *MockService) SendAlert(ctx context.Context, message string) error {
	ret := m.ctrl.Call(m, "SendAlert", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAlert indicates an expected call of SendAlert
func (mr *MockServiceMockRecorder) SendAlert(ctx, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAlert", reflect.TypeOf((*MockService)(nil).SendAlert), ctx, message)
}

// SendNonTechnicalAlert mocks base method
func (m *MockService) SendNonTechnicalAlert(ctx context.Context, message string) error {
	ret := m.ctrl.Call(m, "SendNonTechnicalAlert", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendNonTechnicalAlert indicates an expected call of SendNonTechnicalAlert
func (mr *MockServiceMockRecorder) SendNonTechnicalAlert(ctx, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNonTechnicalAlert", reflect.TypeOf((*MockService)(nil).SendNonTechnicalAlert), ctx, message)
}

// SendHeartbeatGroupAlert mocks base method
func (m *MockService) SendHeartbeatGroupAlert(ctx context.Context, message string) error {
	ret := m.ctrl.Call(m, "SendHeartbeatGroupAlert", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeartbeatGroupAlert indicates an expected call of SendHeartbeatGroupAlert
func (mr *MockServiceMockRecorder) SendHeartbeatGroupAlert(ctx, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeartbeatGroupAlert", reflect.TypeOf((*MockService)(nil).SendHeartbeatGroupAlert), ctx, message)
}

// SendImageToAlertGroup mocks base method
func (m *MockService) SendImageToAlertGroup(ctx context.Context, image []byte) error {
	ret := m.ctrl.Call(m, "SendImageToAlertGroup", ctx, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendImageToAlertGroup indicates an expected call of SendImageToAlertGroup
func (mr *MockServiceMockRecorder) SendImageToAlertGroup(ctx, image interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendImageToAlertGroup", reflect.TypeOf((*MockService)(nil).SendImageToAlertGroup), ctx, image)
}

// SendImageToHeartbeatGroup mocks base method
func (m *MockService) SendImageToHeartbeatGroup(ctx context.Context, image []byte) error {
	ret := m.ctrl.Call(m, "SendImageToHeartbeatGroup", ctx, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendImageToHeartbeatGroup indicates an expected call of SendImageToHeartbeatGroup
func (mr *MockServiceMockRecorder) SendImageToHeartbeatGroup(ctx, image interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendImageToHeartbeatGroup", reflect.TypeOf((*MockService)(nil).SendImageToHeartbeatGroup), ctx, image)
}

// SendError mocks base method
func (m *MockService) SendError(ctx context.Context, err error) error {
	ret := m.ctrl.Call(m, "SendError", ctx, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendError indicates an expected call of SendError
func (mr *MockServiceMockRecorder) SendError(ctx, err interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendError", reflect.TypeOf((*MockService)(nil).SendError), ctx, err)
}
