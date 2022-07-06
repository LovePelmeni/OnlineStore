// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LovePelmeni/OnlineStore/EmailService/emails (interfaces: EmailSenderInterface)

// Package mock_emails is a generated GoMock package.
package mock_emails

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailSenderInterface is a mock of EmailSenderInterface interface.
type MockEmailSenderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEmailSenderInterfaceMockRecorder
}

// MockEmailSenderInterfaceMockRecorder is the mock recorder for MockEmailSenderInterface.
type MockEmailSenderInterfaceMockRecorder struct {
	mock *MockEmailSenderInterface
}

// NewMockEmailSenderInterface creates a new mock instance.
func NewMockEmailSenderInterface(ctrl *gomock.Controller) *MockEmailSenderInterface {
	mock := &MockEmailSenderInterface{ctrl: ctrl}
	mock.recorder = &MockEmailSenderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailSenderInterface) EXPECT() *MockEmailSenderInterfaceMockRecorder {
	return m.recorder
}

// CustomerEmail mocks base method.
func (m *MockEmailSenderInterface) CustomerEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomerEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

// CustomerEmail indicates an expected call of CustomerEmail.
func (mr *MockEmailSenderInterfaceMockRecorder) CustomerEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomerEmail", reflect.TypeOf((*MockEmailSenderInterface)(nil).CustomerEmail))
}

// Message mocks base method.
func (m *MockEmailSenderInterface) Message() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(string)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockEmailSenderInterfaceMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockEmailSenderInterface)(nil).Message))
}

// SendEmail mocks base method.
func (m *MockEmailSenderInterface) SendEmail() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockEmailSenderInterfaceMockRecorder) SendEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockEmailSenderInterface)(nil).SendEmail))
}

// SendOrderEmail mocks base method.
func (m *MockEmailSenderInterface) SendOrderEmail() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendOrderEmail")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendOrderEmail indicates an expected call of SendOrderEmail.
func (mr *MockEmailSenderInterfaceMockRecorder) SendOrderEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendOrderEmail", reflect.TypeOf((*MockEmailSenderInterface)(nil).SendOrderEmail))
}