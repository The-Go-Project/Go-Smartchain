// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mps is a generated GoMock package.
package mps

import (
	context "context"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	state "github.com/ethereum/go-ethereum/core/state"
	types "github.com/ethereum/go-ethereum/core/types"
	trie "github.com/ethereum/go-ethereum/trie"
	gomock "github.com/golang/mock/gomock"
)

// MockPrivateStateManager is a mock of PrivateStateManager interface.
type MockPrivateStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateStateManagerMockRecorder
}

// MockPrivateStateManagerMockRecorder is the mock recorder for MockPrivateStateManager.
type MockPrivateStateManagerMockRecorder struct {
	mock *MockPrivateStateManager
}

// NewMockPrivateStateManager creates a new mock instance.
func NewMockPrivateStateManager(ctrl *gomock.Controller) *MockPrivateStateManager {
	mock := &MockPrivateStateManager{ctrl: ctrl}
	mock.recorder = &MockPrivateStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateStateManager) EXPECT() *MockPrivateStateManagerMockRecorder {
	return m.recorder
}

// CheckAt mocks base method.
func (m *MockPrivateStateManager) CheckAt(blockHash common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAt", blockHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAt indicates an expected call of CheckAt.
func (mr *MockPrivateStateManagerMockRecorder) CheckAt(blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAt", reflect.TypeOf((*MockPrivateStateManager)(nil).CheckAt), blockHash)
}

// NotIncludeAny mocks base method.
func (m *MockPrivateStateManager) NotIncludeAny(psm *PrivateStateMetadata, managedParties ...string) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{psm}
	for _, a := range managedParties {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotIncludeAny", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// NotIncludeAny indicates an expected call of NotIncludeAny.
func (mr *MockPrivateStateManagerMockRecorder) NotIncludeAny(psm interface{}, managedParties ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{psm}, managedParties...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotIncludeAny", reflect.TypeOf((*MockPrivateStateManager)(nil).NotIncludeAny), varargs...)
}

// PSIs mocks base method.
func (m *MockPrivateStateManager) PSIs() []types.PrivateStateIdentifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PSIs")
	ret0, _ := ret[0].([]types.PrivateStateIdentifier)
	return ret0
}

// PSIs indicates an expected call of PSIs.
func (mr *MockPrivateStateManagerMockRecorder) PSIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PSIs", reflect.TypeOf((*MockPrivateStateManager)(nil).PSIs))
}

// ResolveForManagedParty mocks base method.
func (m *MockPrivateStateManager) ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveForManagedParty", managedParty)
	ret0, _ := ret[0].(*PrivateStateMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveForManagedParty indicates an expected call of ResolveForManagedParty.
func (mr *MockPrivateStateManagerMockRecorder) ResolveForManagedParty(managedParty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveForManagedParty", reflect.TypeOf((*MockPrivateStateManager)(nil).ResolveForManagedParty), managedParty)
}

// ResolveForUserContext mocks base method.
func (m *MockPrivateStateManager) ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveForUserContext", ctx)
	ret0, _ := ret[0].(*PrivateStateMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveForUserContext indicates an expected call of ResolveForUserContext.
func (mr *MockPrivateStateManagerMockRecorder) ResolveForUserContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveForUserContext", reflect.TypeOf((*MockPrivateStateManager)(nil).ResolveForUserContext), ctx)
}

// StateRepository mocks base method.
func (m *MockPrivateStateManager) StateRepository(blockHash common.Hash) (PrivateStateRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateRepository", blockHash)
	ret0, _ := ret[0].(PrivateStateRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateRepository indicates an expected call of StateRepository.
func (mr *MockPrivateStateManagerMockRecorder) StateRepository(blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateRepository", reflect.TypeOf((*MockPrivateStateManager)(nil).StateRepository), blockHash)
}

// TrieDB mocks base method.
func (m *MockPrivateStateManager) TrieDB() *trie.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrieDB")
	ret0, _ := ret[0].(*trie.Database)
	return ret0
}

// TrieDB indicates an expected call of TrieDB.
func (mr *MockPrivateStateManagerMockRecorder) TrieDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrieDB", reflect.TypeOf((*MockPrivateStateManager)(nil).TrieDB))
}

// MockPrivateStateMetadataResolver is a mock of PrivateStateMetadataResolver interface.
type MockPrivateStateMetadataResolver struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateStateMetadataResolverMockRecorder
}

// MockPrivateStateMetadataResolverMockRecorder is the mock recorder for MockPrivateStateMetadataResolver.
type MockPrivateStateMetadataResolverMockRecorder struct {
	mock *MockPrivateStateMetadataResolver
}

// NewMockPrivateStateMetadataResolver creates a new mock instance.
func NewMockPrivateStateMetadataResolver(ctrl *gomock.Controller) *MockPrivateStateMetadataResolver {
	mock := &MockPrivateStateMetadataResolver{ctrl: ctrl}
	mock.recorder = &MockPrivateStateMetadataResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateStateMetadataResolver) EXPECT() *MockPrivateStateMetadataResolverMockRecorder {
	return m.recorder
}

// NotIncludeAny mocks base method.
func (m *MockPrivateStateMetadataResolver) NotIncludeAny(psm *PrivateStateMetadata, managedParties ...string) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{psm}
	for _, a := range managedParties {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotIncludeAny", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// NotIncludeAny indicates an expected call of NotIncludeAny.
func (mr *MockPrivateStateMetadataResolverMockRecorder) NotIncludeAny(psm interface{}, managedParties ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{psm}, managedParties...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotIncludeAny", reflect.TypeOf((*MockPrivateStateMetadataResolver)(nil).NotIncludeAny), varargs...)
}

// PSIs mocks base method.
func (m *MockPrivateStateMetadataResolver) PSIs() []types.PrivateStateIdentifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PSIs")
	ret0, _ := ret[0].([]types.PrivateStateIdentifier)
	return ret0
}

// PSIs indicates an expected call of PSIs.
func (mr *MockPrivateStateMetadataResolverMockRecorder) PSIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PSIs", reflect.TypeOf((*MockPrivateStateMetadataResolver)(nil).PSIs))
}

// ResolveForManagedParty mocks base method.
func (m *MockPrivateStateMetadataResolver) ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveForManagedParty", managedParty)
	ret0, _ := ret[0].(*PrivateStateMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveForManagedParty indicates an expected call of ResolveForManagedParty.
func (mr *MockPrivateStateMetadataResolverMockRecorder) ResolveForManagedParty(managedParty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveForManagedParty", reflect.TypeOf((*MockPrivateStateMetadataResolver)(nil).ResolveForManagedParty), managedParty)
}

// ResolveForUserContext mocks base method.
func (m *MockPrivateStateMetadataResolver) ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveForUserContext", ctx)
	ret0, _ := ret[0].(*PrivateStateMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveForUserContext indicates an expected call of ResolveForUserContext.
func (mr *MockPrivateStateMetadataResolverMockRecorder) ResolveForUserContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveForUserContext", reflect.TypeOf((*MockPrivateStateMetadataResolver)(nil).ResolveForUserContext), ctx)
}

// MockPrivateStateRepository is a mock of PrivateStateRepository interface.
type MockPrivateStateRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateStateRepositoryMockRecorder
}

// MockPrivateStateRepositoryMockRecorder is the mock recorder for MockPrivateStateRepository.
type MockPrivateStateRepositoryMockRecorder struct {
	mock *MockPrivateStateRepository
}

// NewMockPrivateStateRepository creates a new mock instance.
func NewMockPrivateStateRepository(ctrl *gomock.Controller) *MockPrivateStateRepository {
	mock := &MockPrivateStateRepository{ctrl: ctrl}
	mock.recorder = &MockPrivateStateRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateStateRepository) EXPECT() *MockPrivateStateRepositoryMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockPrivateStateRepository) Commit(isEIP158 bool, block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", isEIP158, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockPrivateStateRepositoryMockRecorder) Commit(isEIP158, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockPrivateStateRepository)(nil).Commit), isEIP158, block)
}

// CommitAndWrite mocks base method.
func (m *MockPrivateStateRepository) CommitAndWrite(isEIP158 bool, block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitAndWrite", isEIP158, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitAndWrite indicates an expected call of CommitAndWrite.
func (mr *MockPrivateStateRepositoryMockRecorder) CommitAndWrite(isEIP158, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitAndWrite", reflect.TypeOf((*MockPrivateStateRepository)(nil).CommitAndWrite), isEIP158, block)
}

// Copy mocks base method.
func (m *MockPrivateStateRepository) Copy() PrivateStateRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(PrivateStateRepository)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockPrivateStateRepositoryMockRecorder) Copy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockPrivateStateRepository)(nil).Copy))
}

// DefaultState mocks base method.
func (m *MockPrivateStateRepository) DefaultState() (*state.StateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultState")
	ret0, _ := ret[0].(*state.StateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultState indicates an expected call of DefaultState.
func (mr *MockPrivateStateRepositoryMockRecorder) DefaultState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultState", reflect.TypeOf((*MockPrivateStateRepository)(nil).DefaultState))
}

// DefaultStateMetadata mocks base method.
func (m *MockPrivateStateRepository) DefaultStateMetadata() *PrivateStateMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultStateMetadata")
	ret0, _ := ret[0].(*PrivateStateMetadata)
	return ret0
}

// DefaultStateMetadata indicates an expected call of DefaultStateMetadata.
func (mr *MockPrivateStateRepositoryMockRecorder) DefaultStateMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultStateMetadata", reflect.TypeOf((*MockPrivateStateRepository)(nil).DefaultStateMetadata))
}

// IsMPS mocks base method.
func (m *MockPrivateStateRepository) IsMPS() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMPS")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMPS indicates an expected call of IsMPS.
func (mr *MockPrivateStateRepositoryMockRecorder) IsMPS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMPS", reflect.TypeOf((*MockPrivateStateRepository)(nil).IsMPS))
}

// MergeReceipts mocks base method.
func (m *MockPrivateStateRepository) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeReceipts", pub, priv)
	ret0, _ := ret[0].(types.Receipts)
	return ret0
}

// MergeReceipts indicates an expected call of MergeReceipts.
func (mr *MockPrivateStateRepositoryMockRecorder) MergeReceipts(pub, priv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeReceipts", reflect.TypeOf((*MockPrivateStateRepository)(nil).MergeReceipts), pub, priv)
}

// Reset mocks base method.
func (m *MockPrivateStateRepository) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockPrivateStateRepositoryMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockPrivateStateRepository)(nil).Reset))
}

// StatePSI mocks base method.
func (m *MockPrivateStateRepository) StatePSI(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatePSI", psi)
	ret0, _ := ret[0].(*state.StateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatePSI indicates an expected call of StatePSI.
func (mr *MockPrivateStateRepositoryMockRecorder) StatePSI(psi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatePSI", reflect.TypeOf((*MockPrivateStateRepository)(nil).StatePSI), psi)
}
