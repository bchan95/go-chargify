// Code generated by MockGen. DO NOT EDIT.
// Source: backend-services/chargify (interfaces: Client)

// Package mock_chargify is a generated GoMock package.
package test

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockClient) Delete(arg0 []byte, arg1 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockClient) Get(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0)
}

// Post mocks base method
func (m *MockClient) Post(arg0 []byte, arg1 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Post", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post
func (mr *MockClientMockRecorder) Post(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockClient)(nil).Post), arg0, arg1)
}

// Put mocks base method
func (m *MockClient) Put(arg0 []byte, arg1 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockClientMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), arg0, arg1)
}

func (mr *MockClient) GenerateSelfServiceLink(arg0 string, arg1 int64) string {
	return "mock-self-service"
}