// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_chef is a generated GoMock package.
package chef

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
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

// GetChefClientDetails mocks base method
func (m *MockStore) GetChefClientDetails() (ChefClient, error) {
	ret := m.ctrl.Call(m, "GetChefClientDetails")
	ret0, _ := ret[0].(ChefClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChefClientDetails indicates an expected call of GetChefClientDetails
func (mr *MockStoreMockRecorder) GetChefClientDetails() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChefClientDetails", reflect.TypeOf((*MockStore)(nil).GetChefClientDetails))
}

// IsChefConfigured mocks base method
func (m *MockStore) IsChefConfigured() (bool, error) {
	ret := m.ctrl.Call(m, "IsChefConfigured")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsChefConfigured indicates an expected call of IsChefConfigured
func (mr *MockStoreMockRecorder) IsChefConfigured() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsChefConfigured", reflect.TypeOf((*MockStore)(nil).IsChefConfigured))
}

// AddChefEnvironment mocks base method
func (m *MockStore) AddChefEnvironment(environment, friendlyName string) {
	m.ctrl.Call(m, "AddChefEnvironment", environment, friendlyName)
}

// AddChefEnvironment indicates an expected call of AddChefEnvironment
func (mr *MockStoreMockRecorder) AddChefEnvironment(environment, friendlyName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChefEnvironment", reflect.TypeOf((*MockStore)(nil).AddChefEnvironment), environment, friendlyName)
}

// GetChefEnvironments mocks base method
func (m *MockStore) GetChefEnvironments() ([]ChefEnvironment, error) {
	ret := m.ctrl.Call(m, "GetChefEnvironments")
	ret0, _ := ret[0].([]ChefEnvironment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChefEnvironments indicates an expected call of GetChefEnvironments
func (mr *MockStoreMockRecorder) GetChefEnvironments() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChefEnvironments", reflect.TypeOf((*MockStore)(nil).GetChefEnvironments))
}

// GetEnvironmentFromFriendlyName mocks base method
func (m *MockStore) GetEnvironmentFromFriendlyName(recipe string) (string, error) {
	ret := m.ctrl.Call(m, "GetEnvironmentFromFriendlyName", recipe)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironmentFromFriendlyName indicates an expected call of GetEnvironmentFromFriendlyName
func (mr *MockStoreMockRecorder) GetEnvironmentFromFriendlyName(recipe interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironmentFromFriendlyName", reflect.TypeOf((*MockStore)(nil).GetEnvironmentFromFriendlyName), recipe)
}

// AddRecipe mocks base method
func (m *MockStore) AddRecipe(recipeName, friendlyName string) error {
	ret := m.ctrl.Call(m, "AddRecipe", recipeName, friendlyName)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRecipe indicates an expected call of AddRecipe
func (mr *MockStoreMockRecorder) AddRecipe(recipeName, friendlyName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecipe", reflect.TypeOf((*MockStore)(nil).AddRecipe), recipeName, friendlyName)
}

// GetRecipes mocks base method
func (m *MockStore) GetRecipes() ([]Recipe, error) {
	ret := m.ctrl.Call(m, "GetRecipes")
	ret0, _ := ret[0].([]Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipes indicates an expected call of GetRecipes
func (mr *MockStoreMockRecorder) GetRecipes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipes", reflect.TypeOf((*MockStore)(nil).GetRecipes))
}

// GetRecipeFromFriendlyName mocks base method
func (m *MockStore) GetRecipeFromFriendlyName(recipe string) (string, error) {
	ret := m.ctrl.Call(m, "GetRecipeFromFriendlyName", recipe)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipeFromFriendlyName indicates an expected call of GetRecipeFromFriendlyName
func (mr *MockStoreMockRecorder) GetRecipeFromFriendlyName(recipe interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipeFromFriendlyName", reflect.TypeOf((*MockStore)(nil).GetRecipeFromFriendlyName), recipe)
}

// addChefClient mocks base method
func (m *MockStore) addChefClient(name, url, key string) {
	m.ctrl.Call(m, "addChefClient", name, url, key)
}

// addChefClient indicates an expected call of addChefClient
func (mr *MockStoreMockRecorder) addChefClient(name, url, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addChefClient", reflect.TypeOf((*MockStore)(nil).addChefClient), name, url, key)
}
