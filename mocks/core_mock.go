// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/core_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	go_ipfs_blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/runletapp/crabfs/interfaces"
)

// MockCore is a mock of Core interface
type MockCore struct {
	ctrl     *gomock.Controller
	recorder *MockCoreMockRecorder
}

// MockCoreMockRecorder is the mock recorder for MockCore
type MockCoreMockRecorder struct {
	mock *MockCore
}

// NewMockCore creates a new mock instance
func NewMockCore(ctrl *gomock.Controller) *MockCore {
	mock := &MockCore{ctrl: ctrl}
	mock.recorder = &MockCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCore) EXPECT() *MockCoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockCore) Get(ctx context.Context, publicKey interfaces.PubKey, bucket, filename string) (interfaces.Fetcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, publicKey, bucket, filename)
	ret0, _ := ret[0].(interfaces.Fetcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCoreMockRecorder) Get(ctx, publicKey, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCore)(nil).Get), ctx, publicKey, bucket, filename)
}

// Put mocks base method
func (m *MockCore) Put(ctx context.Context, privateKey interfaces.PrivKey, bucket, filename string, file io.Reader, mtime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, privateKey, bucket, filename, file, mtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockCoreMockRecorder) Put(ctx, privateKey, bucket, filename, file, mtime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCore)(nil).Put), ctx, privateKey, bucket, filename, file, mtime)
}

// Remove mocks base method
func (m *MockCore) Remove(ctx context.Context, privateKey interfaces.PrivKey, bucket, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, privateKey, bucket, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCoreMockRecorder) Remove(ctx, privateKey, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCore)(nil).Remove), ctx, privateKey, bucket, filename)
}

// GetID mocks base method
func (m *MockCore) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockCoreMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockCore)(nil).GetID))
}

// GetAddrs mocks base method
func (m *MockCore) GetAddrs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddrs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAddrs indicates an expected call of GetAddrs
func (mr *MockCoreMockRecorder) GetAddrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddrs", reflect.TypeOf((*MockCore)(nil).GetAddrs))
}

// Blockstore mocks base method
func (m *MockCore) Blockstore() go_ipfs_blockstore.Blockstore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Blockstore")
	ret0, _ := ret[0].(go_ipfs_blockstore.Blockstore)
	return ret0
}

// Blockstore indicates an expected call of Blockstore
func (mr *MockCoreMockRecorder) Blockstore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blockstore", reflect.TypeOf((*MockCore)(nil).Blockstore))
}

// Host mocks base method
func (m *MockCore) Host() interfaces.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(interfaces.Host)
	return ret0
}

// Host indicates an expected call of Host
func (mr *MockCoreMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockCore)(nil).Host))
}

// GarbageCollector mocks base method
func (m *MockCore) GarbageCollector() interfaces.GarbageCollector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GarbageCollector")
	ret0, _ := ret[0].(interfaces.GarbageCollector)
	return ret0
}

// GarbageCollector indicates an expected call of GarbageCollector
func (mr *MockCoreMockRecorder) GarbageCollector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GarbageCollector", reflect.TypeOf((*MockCore)(nil).GarbageCollector))
}

// Close mocks base method
func (m *MockCore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCore)(nil).Close))
}

// WithBucket mocks base method
func (m *MockCore) WithBucket(privateKey interfaces.PrivKey, bucket string) (interfaces.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithBucket", privateKey, bucket)
	ret0, _ := ret[0].(interfaces.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithBucket indicates an expected call of WithBucket
func (mr *MockCoreMockRecorder) WithBucket(privateKey, bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithBucket", reflect.TypeOf((*MockCore)(nil).WithBucket), privateKey, bucket)
}

// PublishPublicKey mocks base method
func (m *MockCore) PublishPublicKey(publicKey interfaces.PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishPublicKey", publicKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishPublicKey indicates an expected call of PublishPublicKey
func (mr *MockCoreMockRecorder) PublishPublicKey(publicKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishPublicKey", reflect.TypeOf((*MockCore)(nil).PublishPublicKey), publicKey)
}
