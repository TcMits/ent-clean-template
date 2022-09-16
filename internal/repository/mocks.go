// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	io "io"
	reflect "reflect"

	ent "github.com/TcMits/ent-clean-template/ent"
	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockTransactionRepository) Start(arg0 context.Context) (*ent.Client, func() error, func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(*ent.Client)
	ret1, _ := ret[1].(func() error)
	ret2, _ := ret[2].(func() error)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Start indicates an expected call of Start.
func (mr *MockTransactionRepositoryMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTransactionRepository)(nil).Start), arg0)
}

// MockGetModelRepository is a mock of GetModelRepository interface.
type MockGetModelRepository[ModelType any, WhereInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockGetModelRepositoryMockRecorder[ModelType, WhereInputType]
}

// MockGetModelRepositoryMockRecorder is the mock recorder for MockGetModelRepository.
type MockGetModelRepositoryMockRecorder[ModelType any, WhereInputType any] struct {
	mock *MockGetModelRepository[ModelType, WhereInputType]
}

// NewMockGetModelRepository creates a new mock instance.
func NewMockGetModelRepository[ModelType any, WhereInputType any](ctrl *gomock.Controller) *MockGetModelRepository[ModelType, WhereInputType] {
	mock := &MockGetModelRepository[ModelType, WhereInputType]{ctrl: ctrl}
	mock.recorder = &MockGetModelRepositoryMockRecorder[ModelType, WhereInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetModelRepository[ModelType, WhereInputType]) EXPECT() *MockGetModelRepositoryMockRecorder[ModelType, WhereInputType] {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetModelRepository[ModelType, WhereInputType]) Get(arg0 context.Context, arg1 WhereInputType) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetModelRepositoryMockRecorder[ModelType, WhereInputType]) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetModelRepository[ModelType, WhereInputType])(nil).Get), arg0, arg1)
}

// MockCountModelRepository is a mock of CountModelRepository interface.
type MockCountModelRepository[WhereInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockCountModelRepositoryMockRecorder[WhereInputType]
}

// MockCountModelRepositoryMockRecorder is the mock recorder for MockCountModelRepository.
type MockCountModelRepositoryMockRecorder[WhereInputType any] struct {
	mock *MockCountModelRepository[WhereInputType]
}

// NewMockCountModelRepository creates a new mock instance.
func NewMockCountModelRepository[WhereInputType any](ctrl *gomock.Controller) *MockCountModelRepository[WhereInputType] {
	mock := &MockCountModelRepository[WhereInputType]{ctrl: ctrl}
	mock.recorder = &MockCountModelRepositoryMockRecorder[WhereInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCountModelRepository[WhereInputType]) EXPECT() *MockCountModelRepositoryMockRecorder[WhereInputType] {
	return m.recorder
}

// Count mocks base method.
func (m *MockCountModelRepository[WhereInputType]) Count(arg0 context.Context, arg1 WhereInputType) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCountModelRepositoryMockRecorder[WhereInputType]) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCountModelRepository[WhereInputType])(nil).Count), arg0, arg1)
}

// MockGetWithClientModelRepository is a mock of GetWithClientModelRepository interface.
type MockGetWithClientModelRepository[ModelType any, WhereInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockGetWithClientModelRepositoryMockRecorder[ModelType, WhereInputType]
}

// MockGetWithClientModelRepositoryMockRecorder is the mock recorder for MockGetWithClientModelRepository.
type MockGetWithClientModelRepositoryMockRecorder[ModelType any, WhereInputType any] struct {
	mock *MockGetWithClientModelRepository[ModelType, WhereInputType]
}

// NewMockGetWithClientModelRepository creates a new mock instance.
func NewMockGetWithClientModelRepository[ModelType any, WhereInputType any](ctrl *gomock.Controller) *MockGetWithClientModelRepository[ModelType, WhereInputType] {
	mock := &MockGetWithClientModelRepository[ModelType, WhereInputType]{ctrl: ctrl}
	mock.recorder = &MockGetWithClientModelRepositoryMockRecorder[ModelType, WhereInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetWithClientModelRepository[ModelType, WhereInputType]) EXPECT() *MockGetWithClientModelRepositoryMockRecorder[ModelType, WhereInputType] {
	return m.recorder
}

// GetWithClient mocks base method.
func (m *MockGetWithClientModelRepository[ModelType, WhereInputType]) GetWithClient(arg0 context.Context, arg1 *ent.Client, arg2 WhereInputType, arg3 bool) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithClient indicates an expected call of GetWithClient.
func (mr *MockGetWithClientModelRepositoryMockRecorder[ModelType, WhereInputType]) GetWithClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithClient", reflect.TypeOf((*MockGetWithClientModelRepository[ModelType, WhereInputType])(nil).GetWithClient), arg0, arg1, arg2, arg3)
}

// MockListModelRepository is a mock of ListModelRepository interface.
type MockListModelRepository[ModelType any, OrderInputType any, WhereInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockListModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType]
}

// MockListModelRepositoryMockRecorder is the mock recorder for MockListModelRepository.
type MockListModelRepositoryMockRecorder[ModelType any, OrderInputType any, WhereInputType any] struct {
	mock *MockListModelRepository[ModelType, OrderInputType, WhereInputType]
}

// NewMockListModelRepository creates a new mock instance.
func NewMockListModelRepository[ModelType any, OrderInputType any, WhereInputType any](ctrl *gomock.Controller) *MockListModelRepository[ModelType, OrderInputType, WhereInputType] {
	mock := &MockListModelRepository[ModelType, OrderInputType, WhereInputType]{ctrl: ctrl}
	mock.recorder = &MockListModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListModelRepository[ModelType, OrderInputType, WhereInputType]) EXPECT() *MockListModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType] {
	return m.recorder
}

// List mocks base method.
func (m *MockListModelRepository[ModelType, OrderInputType, WhereInputType]) List(arg0 context.Context, arg1, arg2 int, arg3 OrderInputType, arg4 WhereInputType) ([]ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockListModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType]) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockListModelRepository[ModelType, OrderInputType, WhereInputType])(nil).List), arg0, arg1, arg2, arg3, arg4)
}

// MockListWithClientModelRepository is a mock of ListWithClientModelRepository interface.
type MockListWithClientModelRepository[ModelType any, OrderInputType any, WhereInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockListWithClientModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType]
}

// MockListWithClientModelRepositoryMockRecorder is the mock recorder for MockListWithClientModelRepository.
type MockListWithClientModelRepositoryMockRecorder[ModelType any, OrderInputType any, WhereInputType any] struct {
	mock *MockListWithClientModelRepository[ModelType, OrderInputType, WhereInputType]
}

// NewMockListWithClientModelRepository creates a new mock instance.
func NewMockListWithClientModelRepository[ModelType any, OrderInputType any, WhereInputType any](ctrl *gomock.Controller) *MockListWithClientModelRepository[ModelType, OrderInputType, WhereInputType] {
	mock := &MockListWithClientModelRepository[ModelType, OrderInputType, WhereInputType]{ctrl: ctrl}
	mock.recorder = &MockListWithClientModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListWithClientModelRepository[ModelType, OrderInputType, WhereInputType]) EXPECT() *MockListWithClientModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType] {
	return m.recorder
}

// ListWithClient mocks base method.
func (m *MockListWithClientModelRepository[ModelType, OrderInputType, WhereInputType]) ListWithClient(arg0 context.Context, arg1 *ent.Client, arg2, arg3 int, arg4 OrderInputType, arg5 WhereInputType, arg6 bool) ([]ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWithClient", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWithClient indicates an expected call of ListWithClient.
func (mr *MockListWithClientModelRepositoryMockRecorder[ModelType, OrderInputType, WhereInputType]) ListWithClient(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWithClient", reflect.TypeOf((*MockListWithClientModelRepository[ModelType, OrderInputType, WhereInputType])(nil).ListWithClient), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MockCreateModelRepository is a mock of CreateModelRepository interface.
type MockCreateModelRepository[ModelType any, CreateInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockCreateModelRepositoryMockRecorder[ModelType, CreateInputType]
}

// MockCreateModelRepositoryMockRecorder is the mock recorder for MockCreateModelRepository.
type MockCreateModelRepositoryMockRecorder[ModelType any, CreateInputType any] struct {
	mock *MockCreateModelRepository[ModelType, CreateInputType]
}

// NewMockCreateModelRepository creates a new mock instance.
func NewMockCreateModelRepository[ModelType any, CreateInputType any](ctrl *gomock.Controller) *MockCreateModelRepository[ModelType, CreateInputType] {
	mock := &MockCreateModelRepository[ModelType, CreateInputType]{ctrl: ctrl}
	mock.recorder = &MockCreateModelRepositoryMockRecorder[ModelType, CreateInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateModelRepository[ModelType, CreateInputType]) EXPECT() *MockCreateModelRepositoryMockRecorder[ModelType, CreateInputType] {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreateModelRepository[ModelType, CreateInputType]) Create(arg0 context.Context, arg1 CreateInputType) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCreateModelRepositoryMockRecorder[ModelType, CreateInputType]) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreateModelRepository[ModelType, CreateInputType])(nil).Create), arg0, arg1)
}

// MockCreateWithClientModelRepository is a mock of CreateWithClientModelRepository interface.
type MockCreateWithClientModelRepository[ModelType any, CreateInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockCreateWithClientModelRepositoryMockRecorder[ModelType, CreateInputType]
}

// MockCreateWithClientModelRepositoryMockRecorder is the mock recorder for MockCreateWithClientModelRepository.
type MockCreateWithClientModelRepositoryMockRecorder[ModelType any, CreateInputType any] struct {
	mock *MockCreateWithClientModelRepository[ModelType, CreateInputType]
}

// NewMockCreateWithClientModelRepository creates a new mock instance.
func NewMockCreateWithClientModelRepository[ModelType any, CreateInputType any](ctrl *gomock.Controller) *MockCreateWithClientModelRepository[ModelType, CreateInputType] {
	mock := &MockCreateWithClientModelRepository[ModelType, CreateInputType]{ctrl: ctrl}
	mock.recorder = &MockCreateWithClientModelRepositoryMockRecorder[ModelType, CreateInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateWithClientModelRepository[ModelType, CreateInputType]) EXPECT() *MockCreateWithClientModelRepositoryMockRecorder[ModelType, CreateInputType] {
	return m.recorder
}

// CreateWithClient mocks base method.
func (m *MockCreateWithClientModelRepository[ModelType, CreateInputType]) CreateWithClient(arg0 context.Context, arg1 *ent.Client, arg2 CreateInputType) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithClient indicates an expected call of CreateWithClient.
func (mr *MockCreateWithClientModelRepositoryMockRecorder[ModelType, CreateInputType]) CreateWithClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithClient", reflect.TypeOf((*MockCreateWithClientModelRepository[ModelType, CreateInputType])(nil).CreateWithClient), arg0, arg1, arg2)
}

// MockUpdateModelRepository is a mock of UpdateModelRepository interface.
type MockUpdateModelRepository[ModelType any, UpdateInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateModelRepositoryMockRecorder[ModelType, UpdateInputType]
}

// MockUpdateModelRepositoryMockRecorder is the mock recorder for MockUpdateModelRepository.
type MockUpdateModelRepositoryMockRecorder[ModelType any, UpdateInputType any] struct {
	mock *MockUpdateModelRepository[ModelType, UpdateInputType]
}

// NewMockUpdateModelRepository creates a new mock instance.
func NewMockUpdateModelRepository[ModelType any, UpdateInputType any](ctrl *gomock.Controller) *MockUpdateModelRepository[ModelType, UpdateInputType] {
	mock := &MockUpdateModelRepository[ModelType, UpdateInputType]{ctrl: ctrl}
	mock.recorder = &MockUpdateModelRepositoryMockRecorder[ModelType, UpdateInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateModelRepository[ModelType, UpdateInputType]) EXPECT() *MockUpdateModelRepositoryMockRecorder[ModelType, UpdateInputType] {
	return m.recorder
}

// Update mocks base method.
func (m *MockUpdateModelRepository[ModelType, UpdateInputType]) Update(arg0 context.Context, arg1 ModelType, arg2 UpdateInputType) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUpdateModelRepositoryMockRecorder[ModelType, UpdateInputType]) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUpdateModelRepository[ModelType, UpdateInputType])(nil).Update), arg0, arg1, arg2)
}

// MockUpdateWithClientModelRepository is a mock of UpdateWithClientModelRepository interface.
type MockUpdateWithClientModelRepository[ModelType any, UpdateInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateWithClientModelRepositoryMockRecorder[ModelType, UpdateInputType]
}

// MockUpdateWithClientModelRepositoryMockRecorder is the mock recorder for MockUpdateWithClientModelRepository.
type MockUpdateWithClientModelRepositoryMockRecorder[ModelType any, UpdateInputType any] struct {
	mock *MockUpdateWithClientModelRepository[ModelType, UpdateInputType]
}

// NewMockUpdateWithClientModelRepository creates a new mock instance.
func NewMockUpdateWithClientModelRepository[ModelType any, UpdateInputType any](ctrl *gomock.Controller) *MockUpdateWithClientModelRepository[ModelType, UpdateInputType] {
	mock := &MockUpdateWithClientModelRepository[ModelType, UpdateInputType]{ctrl: ctrl}
	mock.recorder = &MockUpdateWithClientModelRepositoryMockRecorder[ModelType, UpdateInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateWithClientModelRepository[ModelType, UpdateInputType]) EXPECT() *MockUpdateWithClientModelRepositoryMockRecorder[ModelType, UpdateInputType] {
	return m.recorder
}

// UpdateWithClient mocks base method.
func (m *MockUpdateWithClientModelRepository[ModelType, UpdateInputType]) UpdateWithClient(arg0 context.Context, arg1 *ent.Client, arg2 ModelType, arg3 UpdateInputType) (ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWithClient indicates an expected call of UpdateWithClient.
func (mr *MockUpdateWithClientModelRepositoryMockRecorder[ModelType, UpdateInputType]) UpdateWithClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithClient", reflect.TypeOf((*MockUpdateWithClientModelRepository[ModelType, UpdateInputType])(nil).UpdateWithClient), arg0, arg1, arg2, arg3)
}

// MockDeleteModelRepository is a mock of DeleteModelRepository interface.
type MockDeleteModelRepository[ModelType any] struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteModelRepositoryMockRecorder[ModelType]
}

// MockDeleteModelRepositoryMockRecorder is the mock recorder for MockDeleteModelRepository.
type MockDeleteModelRepositoryMockRecorder[ModelType any] struct {
	mock *MockDeleteModelRepository[ModelType]
}

// NewMockDeleteModelRepository creates a new mock instance.
func NewMockDeleteModelRepository[ModelType any](ctrl *gomock.Controller) *MockDeleteModelRepository[ModelType] {
	mock := &MockDeleteModelRepository[ModelType]{ctrl: ctrl}
	mock.recorder = &MockDeleteModelRepositoryMockRecorder[ModelType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteModelRepository[ModelType]) EXPECT() *MockDeleteModelRepositoryMockRecorder[ModelType] {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeleteModelRepository[ModelType]) Delete(arg0 context.Context, arg1 ModelType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeleteModelRepositoryMockRecorder[ModelType]) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeleteModelRepository[ModelType])(nil).Delete), arg0, arg1)
}

// MockDeleteWithClientModelRepository is a mock of DeleteWithClientModelRepository interface.
type MockDeleteWithClientModelRepository[ModelType any] struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteWithClientModelRepositoryMockRecorder[ModelType]
}

// MockDeleteWithClientModelRepositoryMockRecorder is the mock recorder for MockDeleteWithClientModelRepository.
type MockDeleteWithClientModelRepositoryMockRecorder[ModelType any] struct {
	mock *MockDeleteWithClientModelRepository[ModelType]
}

// NewMockDeleteWithClientModelRepository creates a new mock instance.
func NewMockDeleteWithClientModelRepository[ModelType any](ctrl *gomock.Controller) *MockDeleteWithClientModelRepository[ModelType] {
	mock := &MockDeleteWithClientModelRepository[ModelType]{ctrl: ctrl}
	mock.recorder = &MockDeleteWithClientModelRepositoryMockRecorder[ModelType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteWithClientModelRepository[ModelType]) EXPECT() *MockDeleteWithClientModelRepositoryMockRecorder[ModelType] {
	return m.recorder
}

// DeleteWithClient mocks base method.
func (m *MockDeleteWithClientModelRepository[ModelType]) DeleteWithClient(arg0 context.Context, arg1 *ent.Client, arg2 ModelType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWithClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithClient indicates an expected call of DeleteWithClient.
func (mr *MockDeleteWithClientModelRepositoryMockRecorder[ModelType]) DeleteWithClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithClient", reflect.TypeOf((*MockDeleteWithClientModelRepository[ModelType])(nil).DeleteWithClient), arg0, arg1, arg2)
}

// MockLoginRepository is a mock of LoginRepository interface.
type MockLoginRepository[UserType any, WhereInputType any, LoginInputType any] struct {
	ctrl     *gomock.Controller
	recorder *MockLoginRepositoryMockRecorder[UserType, WhereInputType, LoginInputType]
}

// MockLoginRepositoryMockRecorder is the mock recorder for MockLoginRepository.
type MockLoginRepositoryMockRecorder[UserType any, WhereInputType any, LoginInputType any] struct {
	mock *MockLoginRepository[UserType, WhereInputType, LoginInputType]
}

// NewMockLoginRepository creates a new mock instance.
func NewMockLoginRepository[UserType any, WhereInputType any, LoginInputType any](ctrl *gomock.Controller) *MockLoginRepository[UserType, WhereInputType, LoginInputType] {
	mock := &MockLoginRepository[UserType, WhereInputType, LoginInputType]{ctrl: ctrl}
	mock.recorder = &MockLoginRepositoryMockRecorder[UserType, WhereInputType, LoginInputType]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginRepository[UserType, WhereInputType, LoginInputType]) EXPECT() *MockLoginRepositoryMockRecorder[UserType, WhereInputType, LoginInputType] {
	return m.recorder
}

// Login mocks base method.
func (m *MockLoginRepository[UserType, WhereInputType, LoginInputType]) Login(arg0 context.Context, arg1 LoginInputType) (UserType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(UserType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockLoginRepositoryMockRecorder[UserType, WhereInputType, LoginInputType]) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginRepository[UserType, WhereInputType, LoginInputType])(nil).Login), arg0, arg1)
}

// MockReadFileRepository is a mock of ReadFileRepository interface.
type MockReadFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReadFileRepositoryMockRecorder
}

// MockReadFileRepositoryMockRecorder is the mock recorder for MockReadFileRepository.
type MockReadFileRepositoryMockRecorder struct {
	mock *MockReadFileRepository
}

// NewMockReadFileRepository creates a new mock instance.
func NewMockReadFileRepository(ctrl *gomock.Controller) *MockReadFileRepository {
	mock := &MockReadFileRepository{ctrl: ctrl}
	mock.recorder = &MockReadFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadFileRepository) EXPECT() *MockReadFileRepositoryMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockReadFileRepository) Read(arg0 context.Context, arg1 string, arg2 io.Writer, arg3, arg4 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockReadFileRepositoryMockRecorder) Read(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadFileRepository)(nil).Read), arg0, arg1, arg2, arg3, arg4)
}

// MockExistFileRepository is a mock of ExistFileRepository interface.
type MockExistFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExistFileRepositoryMockRecorder
}

// MockExistFileRepositoryMockRecorder is the mock recorder for MockExistFileRepository.
type MockExistFileRepositoryMockRecorder struct {
	mock *MockExistFileRepository
}

// NewMockExistFileRepository creates a new mock instance.
func NewMockExistFileRepository(ctrl *gomock.Controller) *MockExistFileRepository {
	mock := &MockExistFileRepository{ctrl: ctrl}
	mock.recorder = &MockExistFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExistFileRepository) EXPECT() *MockExistFileRepositoryMockRecorder {
	return m.recorder
}

// Exist mocks base method.
func (m *MockExistFileRepository) Exist(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exist indicates an expected call of Exist.
func (mr *MockExistFileRepositoryMockRecorder) Exist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockExistFileRepository)(nil).Exist), arg0, arg1)
}

// MockWriteFileRepository is a mock of WriteFileRepository interface.
type MockWriteFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWriteFileRepositoryMockRecorder
}

// MockWriteFileRepositoryMockRecorder is the mock recorder for MockWriteFileRepository.
type MockWriteFileRepositoryMockRecorder struct {
	mock *MockWriteFileRepository
}

// NewMockWriteFileRepository creates a new mock instance.
func NewMockWriteFileRepository(ctrl *gomock.Controller) *MockWriteFileRepository {
	mock := &MockWriteFileRepository{ctrl: ctrl}
	mock.recorder = &MockWriteFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriteFileRepository) EXPECT() *MockWriteFileRepositoryMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockWriteFileRepository) Write(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockWriteFileRepositoryMockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriteFileRepository)(nil).Write), arg0, arg1, arg2, arg3)
}

// MockDeleteFileRepository is a mock of DeleteFileRepository interface.
type MockDeleteFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteFileRepositoryMockRecorder
}

// MockDeleteFileRepositoryMockRecorder is the mock recorder for MockDeleteFileRepository.
type MockDeleteFileRepositoryMockRecorder struct {
	mock *MockDeleteFileRepository
}

// NewMockDeleteFileRepository creates a new mock instance.
func NewMockDeleteFileRepository(ctrl *gomock.Controller) *MockDeleteFileRepository {
	mock := &MockDeleteFileRepository{ctrl: ctrl}
	mock.recorder = &MockDeleteFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteFileRepository) EXPECT() *MockDeleteFileRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeleteFileRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeleteFileRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeleteFileRepository)(nil).Delete), arg0, arg1)
}

// MockFileRepository is a mock of FileRepository interface.
type MockFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFileRepositoryMockRecorder
}

// MockFileRepositoryMockRecorder is the mock recorder for MockFileRepository.
type MockFileRepositoryMockRecorder struct {
	mock *MockFileRepository
}

// NewMockFileRepository creates a new mock instance.
func NewMockFileRepository(ctrl *gomock.Controller) *MockFileRepository {
	mock := &MockFileRepository{ctrl: ctrl}
	mock.recorder = &MockFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileRepository) EXPECT() *MockFileRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockFileRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFileRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFileRepository)(nil).Delete), arg0, arg1)
}

// Exist mocks base method.
func (m *MockFileRepository) Exist(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exist indicates an expected call of Exist.
func (mr *MockFileRepositoryMockRecorder) Exist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockFileRepository)(nil).Exist), arg0, arg1)
}

// Read mocks base method.
func (m *MockFileRepository) Read(arg0 context.Context, arg1 string, arg2 io.Writer, arg3, arg4 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFileRepositoryMockRecorder) Read(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileRepository)(nil).Read), arg0, arg1, arg2, arg3, arg4)
}

// Write mocks base method.
func (m *MockFileRepository) Write(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockFileRepositoryMockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileRepository)(nil).Write), arg0, arg1, arg2, arg3)
}
