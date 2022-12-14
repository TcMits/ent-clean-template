// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSerializeModelUseCase is a mock of SerializeModelUseCase interface.
type MockSerializeModelUseCase[ModelType any, SerializedType any] struct {
	ctrl     *gomock.Controller
	recorder *MockSerializeModelUseCaseMockRecorder[ModelType, SerializedType]
}

// MockSerializeModelUseCaseMockRecorder is the mock recorder for MockSerializeModelUseCase.
type MockSerializeModelUseCaseMockRecorder[ModelType any, SerializedType any] struct {
	mock *MockSerializeModelUseCase[ModelType, SerializedType]
}

// NewMockSerializeModelUseCase creates a new mock instance.
func NewMockSerializeModelUseCase[ModelType any, SerializedType any](ctrl *gomock.Controller) *MockSerializeModelUseCase[ModelType, SerializedType] {
	mock := &MockSerializeModelUseCase[ModelType, SerializedType]{ctrl: ctrl}
	mock.recorder = &MockSerializeModelUseCaseMockRecorder[ModelType, SerializedType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSerializeModelUseCase[ModelType, SerializedType]) EXPECT() *MockSerializeModelUseCaseMockRecorder[ModelType, SerializedType] {
	return m.recorder
}

// Serialize mocks base method.
func (m *MockSerializeModelUseCase[ModelType, SerializedType]) Serialize(arg0 context.Context, arg1 ModelType) SerializedType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serialize", arg0, arg1)
	ret0, _ := ret[0].(SerializedType)
	return ret0
}

// Serialize indicates an expected call of Serialize.
func (mr *MockSerializeModelUseCaseMockRecorder[ModelType, SerializedType]) Serialize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serialize", reflect.TypeOf((*MockSerializeModelUseCase[ModelType, SerializedType])(nil).Serialize), arg0, arg1)
}

// MockListModelUseCase is a mock of ListModelUseCase interface.
type MockListModelUseCase[ModelType any, OrderInput any, WhereInput any] struct {
	ctrl     *gomock.Controller
	recorder *MockListModelUseCaseMockRecorder[ModelType, OrderInput, WhereInput]
}

// MockListModelUseCaseMockRecorder is the mock recorder for MockListModelUseCase.
type MockListModelUseCaseMockRecorder[ModelType any, OrderInput any, WhereInput any] struct {
	mock *MockListModelUseCase[ModelType, OrderInput, WhereInput]
}

// NewMockListModelUseCase creates a new mock instance.
func NewMockListModelUseCase[ModelType any, OrderInput any, WhereInput any](ctrl *gomock.Controller) *MockListModelUseCase[ModelType, OrderInput, WhereInput] {
	mock := &MockListModelUseCase[ModelType, OrderInput, WhereInput]{ctrl: ctrl}
	mock.recorder = &MockListModelUseCaseMockRecorder[ModelType, OrderInput, WhereInput]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListModelUseCase[ModelType, OrderInput, WhereInput]) EXPECT() *MockListModelUseCaseMockRecorder[ModelType, OrderInput, WhereInput] {
	return m.recorder
}

// List mocks base method.
func (m *MockListModelUseCase[ModelType, OrderInput, WhereInput]) List(arg0 context.Context, arg1, arg2 *int, arg3 OrderInput, arg4 WhereInput) ([]ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockListModelUseCaseMockRecorder[ModelType, OrderInput, WhereInput]) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockListModelUseCase[ModelType, OrderInput, WhereInput])(nil).List), arg0, arg1, arg2, arg3, arg4)
}

// MockGetModelUseCase is a mock of GetModelUseCase interface.
type MockGetModelUseCase[ModelType any, WhereInput any] struct {
	ctrl     *gomock.Controller
	recorder *MockGetModelUseCaseMockRecorder[ModelType, WhereInput]
}

// MockGetModelUseCaseMockRecorder is the mock recorder for MockGetModelUseCase.
type MockGetModelUseCaseMockRecorder[ModelType any, WhereInput any] struct {
	mock *MockGetModelUseCase[ModelType, WhereInput]
}

// NewMockGetModelUseCase creates a new mock instance.
func NewMockGetModelUseCase[ModelType any, WhereInput any](ctrl *gomock.Controller) *MockGetModelUseCase[ModelType, WhereInput] {
	mock := &MockGetModelUseCase[ModelType, WhereInput]{ctrl: ctrl}
	mock.recorder = &MockGetModelUseCaseMockRecorder[ModelType, WhereInput]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetModelUseCase[ModelType, WhereInput]) EXPECT() *MockGetModelUseCaseMockRecorder[ModelType, WhereInput] {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetModelUseCase[ModelType, WhereInput]) Get(arg0 context.Context, arg1 WhereInput) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetModelUseCaseMockRecorder[ModelType, WhereInput]) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetModelUseCase[ModelType, WhereInput])(nil).Get), arg0, arg1)
}

// MockCountModelUseCase is a mock of CountModelUseCase interface.
type MockCountModelUseCase[WhereInput any] struct {
	ctrl     *gomock.Controller
	recorder *MockCountModelUseCaseMockRecorder[WhereInput]
}

// MockCountModelUseCaseMockRecorder is the mock recorder for MockCountModelUseCase.
type MockCountModelUseCaseMockRecorder[WhereInput any] struct {
	mock *MockCountModelUseCase[WhereInput]
}

// NewMockCountModelUseCase creates a new mock instance.
func NewMockCountModelUseCase[WhereInput any](ctrl *gomock.Controller) *MockCountModelUseCase[WhereInput] {
	mock := &MockCountModelUseCase[WhereInput]{ctrl: ctrl}
	mock.recorder = &MockCountModelUseCaseMockRecorder[WhereInput]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCountModelUseCase[WhereInput]) EXPECT() *MockCountModelUseCaseMockRecorder[WhereInput] {
	return m.recorder
}

// Count mocks base method.
func (m *MockCountModelUseCase[WhereInput]) Count(arg0 context.Context, arg1 WhereInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCountModelUseCaseMockRecorder[WhereInput]) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCountModelUseCase[WhereInput])(nil).Count), arg0, arg1)
}

// MockCreateModelUseCase is a mock of CreateModelUseCase interface.
type MockCreateModelUseCase[ModelType any, CreateInput any] struct {
	ctrl     *gomock.Controller
	recorder *MockCreateModelUseCaseMockRecorder[ModelType, CreateInput]
}

// MockCreateModelUseCaseMockRecorder is the mock recorder for MockCreateModelUseCase.
type MockCreateModelUseCaseMockRecorder[ModelType any, CreateInput any] struct {
	mock *MockCreateModelUseCase[ModelType, CreateInput]
}

// NewMockCreateModelUseCase creates a new mock instance.
func NewMockCreateModelUseCase[ModelType any, CreateInput any](ctrl *gomock.Controller) *MockCreateModelUseCase[ModelType, CreateInput] {
	mock := &MockCreateModelUseCase[ModelType, CreateInput]{ctrl: ctrl}
	mock.recorder = &MockCreateModelUseCaseMockRecorder[ModelType, CreateInput]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateModelUseCase[ModelType, CreateInput]) EXPECT() *MockCreateModelUseCaseMockRecorder[ModelType, CreateInput] {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreateModelUseCase[ModelType, CreateInput]) Create(arg0 context.Context, arg1 CreateInput) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCreateModelUseCaseMockRecorder[ModelType, CreateInput]) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreateModelUseCase[ModelType, CreateInput])(nil).Create), arg0, arg1)
}

// MockGetAndUpdateModelUseCase is a mock of GetAndUpdateModelUseCase interface.
type MockGetAndUpdateModelUseCase[ModelType any, WhereInput any, UpdateInput any] struct {
	ctrl     *gomock.Controller
	recorder *MockGetAndUpdateModelUseCaseMockRecorder[ModelType, WhereInput, UpdateInput]
}

// MockGetAndUpdateModelUseCaseMockRecorder is the mock recorder for MockGetAndUpdateModelUseCase.
type MockGetAndUpdateModelUseCaseMockRecorder[ModelType any, WhereInput any, UpdateInput any] struct {
	mock *MockGetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput]
}

// NewMockGetAndUpdateModelUseCase creates a new mock instance.
func NewMockGetAndUpdateModelUseCase[ModelType any, WhereInput any, UpdateInput any](ctrl *gomock.Controller) *MockGetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput] {
	mock := &MockGetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput]{ctrl: ctrl}
	mock.recorder = &MockGetAndUpdateModelUseCaseMockRecorder[ModelType, WhereInput, UpdateInput]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput]) EXPECT() *MockGetAndUpdateModelUseCaseMockRecorder[ModelType, WhereInput, UpdateInput] {
	return m.recorder
}

// GetAndUpdate mocks base method.
func (m *MockGetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput]) GetAndUpdate(arg0 context.Context, arg1 WhereInput, arg2 UpdateInput) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAndUpdate indicates an expected call of GetAndUpdate.
func (mr *MockGetAndUpdateModelUseCaseMockRecorder[ModelType, WhereInput, UpdateInput]) GetAndUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndUpdate", reflect.TypeOf((*MockGetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput])(nil).GetAndUpdate), arg0, arg1, arg2)
}

// MockGetAndDeleteModelUseCase is a mock of GetAndDeleteModelUseCase interface.
type MockGetAndDeleteModelUseCase[ModelType any, WhereInput any] struct {
	ctrl     *gomock.Controller
	recorder *MockGetAndDeleteModelUseCaseMockRecorder[ModelType, WhereInput]
}

// MockGetAndDeleteModelUseCaseMockRecorder is the mock recorder for MockGetAndDeleteModelUseCase.
type MockGetAndDeleteModelUseCaseMockRecorder[ModelType any, WhereInput any] struct {
	mock *MockGetAndDeleteModelUseCase[ModelType, WhereInput]
}

// NewMockGetAndDeleteModelUseCase creates a new mock instance.
func NewMockGetAndDeleteModelUseCase[ModelType any, WhereInput any](ctrl *gomock.Controller) *MockGetAndDeleteModelUseCase[ModelType, WhereInput] {
	mock := &MockGetAndDeleteModelUseCase[ModelType, WhereInput]{ctrl: ctrl}
	mock.recorder = &MockGetAndDeleteModelUseCaseMockRecorder[ModelType, WhereInput]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAndDeleteModelUseCase[ModelType, WhereInput]) EXPECT() *MockGetAndDeleteModelUseCaseMockRecorder[ModelType, WhereInput] {
	return m.recorder
}

// GetAndDelete mocks base method.
func (m *MockGetAndDeleteModelUseCase[ModelType, WhereInput]) GetAndDelete(arg0 context.Context, arg1 WhereInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAndDelete indicates an expected call of GetAndDelete.
func (mr *MockGetAndDeleteModelUseCaseMockRecorder[ModelType, WhereInput]) GetAndDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndDelete", reflect.TypeOf((*MockGetAndDeleteModelUseCase[ModelType, WhereInput])(nil).GetAndDelete), arg0, arg1)
}

// MockUserPermissionCheckerUseCase is a mock of UserPermissionCheckerUseCase interface.
type MockUserPermissionCheckerUseCase[UserType any] struct {
	ctrl     *gomock.Controller
	recorder *MockUserPermissionCheckerUseCaseMockRecorder[UserType]
}

// MockUserPermissionCheckerUseCaseMockRecorder is the mock recorder for MockUserPermissionCheckerUseCase.
type MockUserPermissionCheckerUseCaseMockRecorder[UserType any] struct {
	mock *MockUserPermissionCheckerUseCase[UserType]
}

// NewMockUserPermissionCheckerUseCase creates a new mock instance.
func NewMockUserPermissionCheckerUseCase[UserType any](ctrl *gomock.Controller) *MockUserPermissionCheckerUseCase[UserType] {
	mock := &MockUserPermissionCheckerUseCase[UserType]{ctrl: ctrl}
	mock.recorder = &MockUserPermissionCheckerUseCaseMockRecorder[UserType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserPermissionCheckerUseCase[UserType]) EXPECT() *MockUserPermissionCheckerUseCaseMockRecorder[UserType] {
	return m.recorder
}

// And mocks base method.
func (m *MockUserPermissionCheckerUseCase[UserType]) And(arg0 UserPermissionCheckerUseCase[UserType]) UserPermissionCheckerUseCase[UserType] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "And", arg0)
	ret0, _ := ret[0].(UserPermissionCheckerUseCase[UserType])
	return ret0
}

// And indicates an expected call of And.
func (mr *MockUserPermissionCheckerUseCaseMockRecorder[UserType]) And(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "And", reflect.TypeOf((*MockUserPermissionCheckerUseCase[UserType])(nil).And), arg0)
}

// Check mocks base method.
func (m *MockUserPermissionCheckerUseCase[UserType]) Check(arg0 context.Context, arg1 UserType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockUserPermissionCheckerUseCaseMockRecorder[UserType]) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockUserPermissionCheckerUseCase[UserType])(nil).Check), arg0, arg1)
}

// Or mocks base method.
func (m *MockUserPermissionCheckerUseCase[UserType]) Or(arg0 UserPermissionCheckerUseCase[UserType]) UserPermissionCheckerUseCase[UserType] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Or", arg0)
	ret0, _ := ret[0].(UserPermissionCheckerUseCase[UserType])
	return ret0
}

// Or indicates an expected call of Or.
func (mr *MockUserPermissionCheckerUseCaseMockRecorder[UserType]) Or(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*MockUserPermissionCheckerUseCase[UserType])(nil).Or), arg0)
}

// MockLoginUseCase is a mock of LoginUseCase interface.
type MockLoginUseCase[LoginInputType any, JWTAuthenticatedPayloadType any, RefreshTokenInputType any, UserType any] struct {
	ctrl     *gomock.Controller
	recorder *MockLoginUseCaseMockRecorder[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]
}

// MockLoginUseCaseMockRecorder is the mock recorder for MockLoginUseCase.
type MockLoginUseCaseMockRecorder[LoginInputType any, JWTAuthenticatedPayloadType any, RefreshTokenInputType any, UserType any] struct {
	mock *MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]
}

// NewMockLoginUseCase creates a new mock instance.
func NewMockLoginUseCase[LoginInputType any, JWTAuthenticatedPayloadType any, RefreshTokenInputType any, UserType any](ctrl *gomock.Controller) *MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType] {
	mock := &MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]{ctrl: ctrl}
	mock.recorder = &MockLoginUseCaseMockRecorder[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) EXPECT() *MockLoginUseCaseMockRecorder[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType] {
	return m.recorder
}

// Login mocks base method.
func (m *MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) Login(arg0 context.Context, arg1 LoginInputType) (JWTAuthenticatedPayloadType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(JWTAuthenticatedPayloadType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockLoginUseCaseMockRecorder[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType])(nil).Login), arg0, arg1)
}

// RefreshToken mocks base method.
func (m *MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) RefreshToken(arg0 context.Context, arg1 RefreshTokenInputType) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockLoginUseCaseMockRecorder[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) RefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType])(nil).RefreshToken), arg0, arg1)
}

// VerifyToken mocks base method.
func (m *MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) VerifyToken(arg0 context.Context, arg1 string) (UserType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", arg0, arg1)
	ret0, _ := ret[0].(UserType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockLoginUseCaseMockRecorder[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType]) VerifyToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockLoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType])(nil).VerifyToken), arg0, arg1)
}
