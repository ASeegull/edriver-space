// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

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
func (m *MockUsers) AddDriverLicence(ctx context.Context, input service.AddDriverLicenceInput, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDriverLicence", ctx, input, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDriverLicence indicates an expected call of AddDriverLicence.
func (mr *MockUsersMockRecorder) AddDriverLicence(ctx, input, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDriverLicence", reflect.TypeOf((*MockUsers)(nil).AddDriverLicence), ctx, input, userId)
}

// DeleteSession mocks base method.
func (m *MockUsers) DeleteSession(ctx context.Context, sessionId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, sessionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockUsersMockRecorder) DeleteSession(ctx, sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockUsers)(nil).DeleteSession), ctx, sessionId)
}

// GetFines mocks base method.
func (m *MockUsers) GetFines(ctx context.Context, userId string) (model.Fines, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFines", ctx, userId)
	ret0, _ := ret[0].(model.Fines)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFines indicates an expected call of GetFines.
func (mr *MockUsersMockRecorder) GetFines(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFines", reflect.TypeOf((*MockUsers)(nil).GetFines), ctx, userId)
}

// GetUserById mocks base method.
func (m *MockUsers) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, userId)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUsersMockRecorder) GetUserById(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUsers)(nil).GetUserById), ctx, userId)
}

// RefreshTokens mocks base method.
func (m *MockUsers) RefreshTokens(ctx context.Context, sessionId string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", ctx, sessionId)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockUsersMockRecorder) RefreshTokens(ctx, sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockUsers)(nil).RefreshTokens), ctx, sessionId)
}

// SignIn mocks base method.
func (m *MockUsers) SignIn(ctx context.Context, user service.UserSignInInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, user)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUsersMockRecorder) SignIn(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUsers)(nil).SignIn), ctx, user)
}

// SignUp mocks base method.
func (m *MockUsers) SignUp(ctx context.Context, user service.UserSignUpInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, user)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUsersMockRecorder) SignUp(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUsers)(nil).SignUp), ctx, user)
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

// MockCars is a mock of Cars interface.
type MockCars struct {
	ctrl     *gomock.Controller
	recorder *MockCarsMockRecorder
}

// MockCarsMockRecorder is the mock recorder for MockCars.
type MockCarsMockRecorder struct {
	mock *MockCars
}

// NewMockCars creates a new mock instance.
func NewMockCars(ctrl *gomock.Controller) *MockCars {
	mock := &MockCars{ctrl: ctrl}
	mock.recorder = &MockCarsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCars) EXPECT() *MockCarsMockRecorder {
	return m.recorder
}

// CreateCar mocks base method.
func (m *MockCars) CreateCar(ctx context.Context, car *model.Car) (*model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCar", ctx, car)
	ret0, _ := ret[0].(*model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCar indicates an expected call of CreateCar.
func (mr *MockCarsMockRecorder) CreateCar(ctx, car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCar", reflect.TypeOf((*MockCars)(nil).CreateCar), ctx, car)
}

// DeleteCar mocks base method.
func (m *MockCars) DeleteCar(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCar", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCar indicates an expected call of DeleteCar.
func (mr *MockCarsMockRecorder) DeleteCar(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCar", reflect.TypeOf((*MockCars)(nil).DeleteCar), ctx, id)
}

// GetCar mocks base method.
func (m *MockCars) GetCar(ctx context.Context, id string) (*model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCar", ctx, id)
	ret0, _ := ret[0].(*model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCar indicates an expected call of GetCar.
func (mr *MockCarsMockRecorder) GetCar(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCar", reflect.TypeOf((*MockCars)(nil).GetCar), ctx, id)
}

// GetCars mocks base method.
func (m *MockCars) GetCars(ctx context.Context) (*[]model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCars", ctx)
	ret0, _ := ret[0].(*[]model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCars indicates an expected call of GetCars.
func (mr *MockCarsMockRecorder) GetCars(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCars", reflect.TypeOf((*MockCars)(nil).GetCars), ctx)
}

// MockDrivers is a mock of Drivers interface.
type MockDrivers struct {
	ctrl     *gomock.Controller
	recorder *MockDriversMockRecorder
}

// MockDriversMockRecorder is the mock recorder for MockDrivers.
type MockDriversMockRecorder struct {
	mock *MockDrivers
}

// NewMockDrivers creates a new mock instance.
func NewMockDrivers(ctrl *gomock.Controller) *MockDrivers {
	mock := &MockDrivers{ctrl: ctrl}
	mock.recorder = &MockDriversMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDrivers) EXPECT() *MockDriversMockRecorder {
	return m.recorder
}

// CreateDriver mocks base method.
func (m *MockDrivers) CreateDriver(ctx context.Context, driver *model.Driver) (*model.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDriver", ctx, driver)
	ret0, _ := ret[0].(*model.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDriver indicates an expected call of CreateDriver.
func (mr *MockDriversMockRecorder) CreateDriver(ctx, driver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDriver", reflect.TypeOf((*MockDrivers)(nil).CreateDriver), ctx, driver)
}

// DeleteDriver mocks base method.
func (m *MockDrivers) DeleteDriver(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDriver", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDriver indicates an expected call of DeleteDriver.
func (mr *MockDriversMockRecorder) DeleteDriver(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDriver", reflect.TypeOf((*MockDrivers)(nil).DeleteDriver), ctx, id)
}

// GetDriver mocks base method.
func (m *MockDrivers) GetDriver(ctx context.Context, id string) (*model.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriver", ctx, id)
	ret0, _ := ret[0].(*model.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriver indicates an expected call of GetDriver.
func (mr *MockDriversMockRecorder) GetDriver(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriver", reflect.TypeOf((*MockDrivers)(nil).GetDriver), ctx, id)
}

// GetDrivers mocks base method.
func (m *MockDrivers) GetDrivers(ctx context.Context) (*[]model.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDrivers", ctx)
	ret0, _ := ret[0].(*[]model.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDrivers indicates an expected call of GetDrivers.
func (mr *MockDriversMockRecorder) GetDrivers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDrivers", reflect.TypeOf((*MockDrivers)(nil).GetDrivers), ctx)
}