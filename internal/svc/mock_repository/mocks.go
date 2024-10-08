// Code generated by MockGen. DO NOT EDIT.
// Source: internal/db/repositories.go
//
// Generated by this command:
//
//	mockgen -source=internal/db/repositories.go -destination=internal/svc/mock_repository/mocks.go -package mock_repository
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "RD-Clone-NAPI/internal/models"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), arg0, arg1)
}

// FindByUsername mocks base method.
func (m *MockUserRepository) FindByUsername(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockUserRepositoryMockRecorder) FindByUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockUserRepository)(nil).FindByUsername), arg0, arg1)
}

// Save mocks base method.
func (m *MockUserRepository) Save(arg0 context.Context, arg1 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserRepository) Update(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), arg0, arg1)
}

// MockTokenRepository is a mock of TokenRepository interface.
type MockTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRepositoryMockRecorder
}

// MockTokenRepositoryMockRecorder is the mock recorder for MockTokenRepository.
type MockTokenRepositoryMockRecorder struct {
	mock *MockTokenRepository
}

// NewMockTokenRepository creates a new mock instance.
func NewMockTokenRepository(ctrl *gomock.Controller) *MockTokenRepository {
	mock := &MockTokenRepository{ctrl: ctrl}
	mock.recorder = &MockTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenRepository) EXPECT() *MockTokenRepositoryMockRecorder {
	return m.recorder
}

// FindByToken mocks base method.
func (m *MockTokenRepository) FindByToken(arg0 context.Context, arg1 string) (*models.VerificationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", arg0, arg1)
	ret0, _ := ret[0].(*models.VerificationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockTokenRepositoryMockRecorder) FindByToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockTokenRepository)(nil).FindByToken), arg0, arg1)
}

// Save mocks base method.
func (m *MockTokenRepository) Save(arg0 context.Context, arg1 *models.VerificationToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTokenRepositoryMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTokenRepository)(nil).Save), arg0, arg1)
}

// MockRefreshTokenRepository is a mock of RefreshTokenRepository interface.
type MockRefreshTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenRepositoryMockRecorder
}

// MockRefreshTokenRepositoryMockRecorder is the mock recorder for MockRefreshTokenRepository.
type MockRefreshTokenRepositoryMockRecorder struct {
	mock *MockRefreshTokenRepository
}

// NewMockRefreshTokenRepository creates a new mock instance.
func NewMockRefreshTokenRepository(ctrl *gomock.Controller) *MockRefreshTokenRepository {
	mock := &MockRefreshTokenRepository{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenRepository) EXPECT() *MockRefreshTokenRepositoryMockRecorder {
	return m.recorder
}

// FindByToken mocks base method.
func (m *MockRefreshTokenRepository) FindByToken(arg0 context.Context, arg1 string) (*models.RefreshToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", arg0, arg1)
	ret0, _ := ret[0].(*models.RefreshToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockRefreshTokenRepositoryMockRecorder) FindByToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockRefreshTokenRepository)(nil).FindByToken), arg0, arg1)
}

// Save mocks base method.
func (m *MockRefreshTokenRepository) Save(arg0 context.Context, arg1 *models.RefreshToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRefreshTokenRepositoryMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRefreshTokenRepository)(nil).Save), arg0, arg1)
}
