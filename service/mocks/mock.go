// Code generated by MockGen. DO NOT EDIT.

// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	bytes "bytes"
	context "context"
	reflect "reflect"

	model "github.com/ASeegull/edriver-space/model"
	service "github.com/ASeegull/edriver-space/service"
	gomock "github.com/golang/mock/gomock"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// AddDriverLicence mocks base method.
func (m *MockUsers) AddDriverLicence(arg0 context.Context, arg1 service.AddDriverLicenceInput, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDriverLicence", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDriverLicence indicates an expected call of AddDriverLicence.
func (mr *MockUsersMockRecorder) AddDriverLicence(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDriverLicence", reflect.TypeOf((*MockUsers)(nil).AddDriverLicence), arg0, arg1, arg2)
}

// DeleteSession mocks base method.
func (m *MockUsers) DeleteSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockUsersMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockUsers)(nil).DeleteSession), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockUsers) GetUserById(arg0 context.Context, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUsersMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUsers)(nil).GetUserById), arg0, arg1)
}

// RefreshTokens mocks base method.
func (m *MockUsers) RefreshTokens(arg0 context.Context, arg1 string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", arg0, arg1)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockUsersMockRecorder) RefreshTokens(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockUsers)(nil).RefreshTokens), arg0, arg1)
}

// SignIn mocks base method.
func (m *MockUsers) SignIn(arg0 context.Context, arg1 service.UserSignInInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0, arg1)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUsersMockRecorder) SignIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUsers)(nil).SignIn), arg0, arg1)
}

// SignUp mocks base method.
func (m *MockUsers) SignUp(arg0 context.Context, arg1 service.UserSignUpInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0, arg1)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUsersMockRecorder) SignUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUsers)(nil).SignUp), arg0, arg1)
}

// MockUploader is a mock of Uploader interface.
type MockUploader struct {
	ctrl     *gomock.Controller
	recorder *MockUploaderMockRecorder
}

// MockUploaderMockRecorder is the mock recorder for MockUploader.
type MockUploaderMockRecorder struct {
	mock *MockUploader
}

// NewMockUploader creates a new mock instance.
func NewMockUploader(ctrl *gomock.Controller) *MockUploader {
	mock := &MockUploader{ctrl: ctrl}
	mock.recorder = &MockUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploader) EXPECT() *MockUploaderMockRecorder {
	return m.recorder
}

// ReadFinesExcel mocks base method.
func (m *MockUploader) ReadFinesExcel(ctx context.Context, r *bytes.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFinesExcel", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFinesExcel indicates an expected call of ReadFinesExcel.
func (mr *MockUploaderMockRecorder) ReadFinesExcel(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFinesExcel", reflect.TypeOf((*MockUploader)(nil).ReadFinesExcel), ctx, r)
}

// XMLFinesService mocks base method.
func (m *MockUploader) XMLFinesService(ctx context.Context, data model.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XMLFinesService", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// XMLFinesService indicates an expected call of XMLFinesService.
func (mr *MockUploaderMockRecorder) XMLFinesService(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XMLFinesService", reflect.TypeOf((*MockUploader)(nil).XMLFinesService), ctx, data)
}