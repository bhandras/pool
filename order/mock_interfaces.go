// Code generated by MockGen. DO NOT EDIT.
// Source: order/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=order/interfaces.go -package=order -destination=order/mock_interfaces.go
//

// Package order is a generated GoMock package.
package order

import (
	context "context"
	reflect "reflect"

	btcutil "github.com/btcsuite/btcd/btcutil"
	account "github.com/lightninglabs/pool/account"
	terms "github.com/lightninglabs/pool/terms"
	gomock "go.uber.org/mock/gomock"
)

// MockOrder is a mock of Order interface.
type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

// MockOrderMockRecorder is the mock recorder for MockOrder.
type MockOrderMockRecorder struct {
	mock *MockOrder
}

// NewMockOrder creates a new mock instance.
func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

// Details mocks base method.
func (m *MockOrder) Details() *Kit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Details")
	ret0, _ := ret[0].(*Kit)
	return ret0
}

// Details indicates an expected call of Details.
func (mr *MockOrderMockRecorder) Details() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Details", reflect.TypeOf((*MockOrder)(nil).Details))
}

// Digest mocks base method.
func (m *MockOrder) Digest() ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Digest")
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Digest indicates an expected call of Digest.
func (mr *MockOrderMockRecorder) Digest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Digest", reflect.TypeOf((*MockOrder)(nil).Digest))
}

// Nonce mocks base method.
func (m *MockOrder) Nonce() Nonce {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nonce")
	ret0, _ := ret[0].(Nonce)
	return ret0
}

// Nonce indicates an expected call of Nonce.
func (mr *MockOrderMockRecorder) Nonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockOrder)(nil).Nonce))
}

// ReservedValue mocks base method.
func (m *MockOrder) ReservedValue(feeSchedule terms.FeeSchedule, accountVersion account.Version) btcutil.Amount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReservedValue", feeSchedule, accountVersion)
	ret0, _ := ret[0].(btcutil.Amount)
	return ret0
}

// ReservedValue indicates an expected call of ReservedValue.
func (mr *MockOrderMockRecorder) ReservedValue(feeSchedule, accountVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReservedValue", reflect.TypeOf((*MockOrder)(nil).ReservedValue), feeSchedule, accountVersion)
}

// Type mocks base method.
func (m *MockOrder) Type() Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockOrderMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockOrder)(nil).Type))
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(arg0 Nonce) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0)
}

// GetOrder mocks base method.
func (m *MockStore) GetOrder(arg0 Nonce) (Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0)
	ret0, _ := ret[0].(Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockStoreMockRecorder) GetOrder(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockStore)(nil).GetOrder), arg0)
}

// GetOrders mocks base method.
func (m *MockStore) GetOrders() ([]Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders")
	ret0, _ := ret[0].([]Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockStoreMockRecorder) GetOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockStore)(nil).GetOrders))
}

// MarkBatchComplete mocks base method.
func (m *MockStore) MarkBatchComplete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkBatchComplete")
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkBatchComplete indicates an expected call of MarkBatchComplete.
func (mr *MockStoreMockRecorder) MarkBatchComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkBatchComplete", reflect.TypeOf((*MockStore)(nil).MarkBatchComplete))
}

// StorePendingBatch mocks base method.
func (m *MockStore) StorePendingBatch(arg0 *Batch, orders []Nonce, orderModifiers [][]Modifier, accounts []*account.Account, accountModifiers [][]account.Modifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorePendingBatch", arg0, orders, orderModifiers, accounts, accountModifiers)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorePendingBatch indicates an expected call of StorePendingBatch.
func (mr *MockStoreMockRecorder) StorePendingBatch(arg0, orders, orderModifiers, accounts, accountModifiers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorePendingBatch", reflect.TypeOf((*MockStore)(nil).StorePendingBatch), arg0, orders, orderModifiers, accounts, accountModifiers)
}

// SubmitOrder mocks base method.
func (m *MockStore) SubmitOrder(arg0 Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitOrder indicates an expected call of SubmitOrder.
func (mr *MockStoreMockRecorder) SubmitOrder(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockStore)(nil).SubmitOrder), arg0)
}

// UpdateOrder mocks base method.
func (m *MockStore) UpdateOrder(arg0 Nonce, arg1 ...Modifier) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrder", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStoreMockRecorder) UpdateOrder(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStore)(nil).UpdateOrder), varargs...)
}

// UpdateOrders mocks base method.
func (m *MockStore) UpdateOrders(arg0 []Nonce, arg1 [][]Modifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrders", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrders indicates an expected call of UpdateOrders.
func (mr *MockStoreMockRecorder) UpdateOrders(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrders", reflect.TypeOf((*MockStore)(nil).UpdateOrders), arg0, arg1)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// BatchFinalize mocks base method.
func (m *MockManager) BatchFinalize(batchID BatchID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchFinalize", batchID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchFinalize indicates an expected call of BatchFinalize.
func (mr *MockManagerMockRecorder) BatchFinalize(batchID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchFinalize", reflect.TypeOf((*MockManager)(nil).BatchFinalize), batchID)
}

// BatchSign mocks base method.
func (m *MockManager) BatchSign() (BatchSignature, AccountNonces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchSign")
	ret0, _ := ret[0].(BatchSignature)
	ret1, _ := ret[1].(AccountNonces)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BatchSign indicates an expected call of BatchSign.
func (mr *MockManagerMockRecorder) BatchSign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSign", reflect.TypeOf((*MockManager)(nil).BatchSign))
}

// HasPendingBatch mocks base method.
func (m *MockManager) HasPendingBatch() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPendingBatch")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPendingBatch indicates an expected call of HasPendingBatch.
func (mr *MockManagerMockRecorder) HasPendingBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPendingBatch", reflect.TypeOf((*MockManager)(nil).HasPendingBatch))
}

// OrderMatchValidate mocks base method.
func (m *MockManager) OrderMatchValidate(batch *Batch, bestHeight uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderMatchValidate", batch, bestHeight)
	ret0, _ := ret[0].(error)
	return ret0
}

// OrderMatchValidate indicates an expected call of OrderMatchValidate.
func (mr *MockManagerMockRecorder) OrderMatchValidate(batch, bestHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderMatchValidate", reflect.TypeOf((*MockManager)(nil).OrderMatchValidate), batch, bestHeight)
}

// OurNodePubkey mocks base method.
func (m *MockManager) OurNodePubkey() ([33]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OurNodePubkey")
	ret0, _ := ret[0].([33]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OurNodePubkey indicates an expected call of OurNodePubkey.
func (mr *MockManagerMockRecorder) OurNodePubkey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OurNodePubkey", reflect.TypeOf((*MockManager)(nil).OurNodePubkey))
}

// PendingBatch mocks base method.
func (m *MockManager) PendingBatch() *Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingBatch")
	ret0, _ := ret[0].(*Batch)
	return ret0
}

// PendingBatch indicates an expected call of PendingBatch.
func (mr *MockManagerMockRecorder) PendingBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingBatch", reflect.TypeOf((*MockManager)(nil).PendingBatch))
}

// PrepareOrder mocks base method.
func (m *MockManager) PrepareOrder(ctx context.Context, order Order, acct *account.Account, terms *terms.AuctioneerTerms) (*ServerOrderParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareOrder", ctx, order, acct, terms)
	ret0, _ := ret[0].(*ServerOrderParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareOrder indicates an expected call of PrepareOrder.
func (mr *MockManagerMockRecorder) PrepareOrder(ctx, order, acct, terms any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareOrder", reflect.TypeOf((*MockManager)(nil).PrepareOrder), ctx, order, acct, terms)
}

// Start mocks base method.
func (m *MockManager) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop))
}
