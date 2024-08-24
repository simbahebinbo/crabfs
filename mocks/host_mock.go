// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/host_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	go_cid "github.com/ipfs/go-cid"
	go_libp2p_peerstore "github.com/libp2p/go-libp2p-peerstore"
	crypto "github.com/simbahebinbo/crabfs/crypto"
	"github.com/simbahebinbo/crabfs/interfaces"
	protos "github.com/simbahebinbo/crabfs/protos"
)

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHost) EXPECT() *MockHostMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockHost) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockHostMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHost)(nil).Close))
}

// Announce mocks base method
func (m *MockHost) Announce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Announce")
	ret0, _ := ret[0].(error)
	return ret0
}

// Announce indicates an expected call of Announce
func (mr *MockHostMockRecorder) Announce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Announce", reflect.TypeOf((*MockHost)(nil).Announce))
}

// GetSwarmPublicKey mocks base method
func (m *MockHost) GetSwarmPublicKey(ctx context.Context, hash string) (crypto.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSwarmPublicKey", ctx, hash)
	ret0, _ := ret[0].(crypto.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSwarmPublicKey indicates an expected call of GetSwarmPublicKey
func (mr *MockHostMockRecorder) GetSwarmPublicKey(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSwarmPublicKey", reflect.TypeOf((*MockHost)(nil).GetSwarmPublicKey), ctx, hash)
}

// Publish mocks base method
func (m *MockHost) Publish(ctx context.Context, privateKey crypto.PrivKey, cipherKey []byte, bucket, filename string, blockMap interfaces.BlockMap, mtime time.Time, size int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockHostMockRecorder) Publish(ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockHost)(nil).Publish), ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size)
}

// PublishAndLock mocks base method
func (m *MockHost) PublishAndLock(ctx context.Context, privateKey crypto.PrivKey, cipherKey []byte, bucket, filename string, blockMap interfaces.BlockMap, mtime time.Time, size int64) (*protos.LockToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAndLock", ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size)
	ret0, _ := ret[0].(*protos.LockToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAndLock indicates an expected call of PublishAndLock
func (mr *MockHostMockRecorder) PublishAndLock(ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAndLock", reflect.TypeOf((*MockHost)(nil).PublishAndLock), ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size)
}

// PublishWithCacheTTL mocks base method
func (m *MockHost) PublishWithCacheTTL(ctx context.Context, privateKey crypto.PrivKey, cipherKey []byte, bucket, filename string, blockMap interfaces.BlockMap, mtime time.Time, size int64, ttl uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithCacheTTL", ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishWithCacheTTL indicates an expected call of PublishWithCacheTTL
func (mr *MockHostMockRecorder) PublishWithCacheTTL(ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithCacheTTL", reflect.TypeOf((*MockHost)(nil).PublishWithCacheTTL), ctx, privateKey, cipherKey, bucket, filename, blockMap, mtime, size, ttl)
}

// Remove mocks base method
func (m *MockHost) Remove(ctx context.Context, privateKey crypto.PrivKey, bucket, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, privateKey, bucket, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockHostMockRecorder) Remove(ctx, privateKey, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockHost)(nil).Remove), ctx, privateKey, bucket, filename)
}

// GetContent mocks base method
func (m *MockHost) GetContent(ctx context.Context, publicKey crypto.PubKey, bucket, filename string) (*protos.CrabObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", ctx, publicKey, bucket, filename)
	ret0, _ := ret[0].(*protos.CrabObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent
func (mr *MockHostMockRecorder) GetContent(ctx, publicKey, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockHost)(nil).GetContent), ctx, publicKey, bucket, filename)
}

// Lock mocks base method
func (m *MockHost) Lock(ctx context.Context, privateKey crypto.PrivKey, bucket, filename string) (*protos.LockToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, privateKey, bucket, filename)
	ret0, _ := ret[0].(*protos.LockToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock
func (mr *MockHostMockRecorder) Lock(ctx, privateKey, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockHost)(nil).Lock), ctx, privateKey, bucket, filename)
}

// Unlock mocks base method
func (m *MockHost) Unlock(ctx context.Context, privateKey crypto.PrivKey, bucket, filename string, token *protos.LockToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", ctx, privateKey, bucket, filename, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock
func (mr *MockHostMockRecorder) Unlock(ctx, privateKey, bucket, filename, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockHost)(nil).Unlock), ctx, privateKey, bucket, filename, token)
}

// IsLocked mocks base method
func (m *MockHost) IsLocked(ctx context.Context, publicKey crypto.PubKey, bucket, filename string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLocked", ctx, publicKey, bucket, filename)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLocked indicates an expected call of IsLocked
func (mr *MockHostMockRecorder) IsLocked(ctx, publicKey, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLocked", reflect.TypeOf((*MockHost)(nil).IsLocked), ctx, publicKey, bucket, filename)
}

// FindProviders mocks base method
func (m *MockHost) FindProviders(ctx context.Context, blockMeta *protos.BlockMetadata) <-chan go_libp2p_peerstore.PeerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProviders", ctx, blockMeta)
	ret0, _ := ret[0].(<-chan go_libp2p_peerstore.PeerInfo)
	return ret0
}

// FindProviders indicates an expected call of FindProviders
func (mr *MockHostMockRecorder) FindProviders(ctx, blockMeta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProviders", reflect.TypeOf((*MockHost)(nil).FindProviders), ctx, blockMeta)
}

// CreateBlockStream mocks base method
func (m *MockHost) CreateBlockStream(ctx context.Context, blockMeta *protos.BlockMetadata, peer *go_libp2p_peerstore.PeerInfo) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlockStream", ctx, blockMeta, peer)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBlockStream indicates an expected call of CreateBlockStream
func (mr *MockHostMockRecorder) CreateBlockStream(ctx, blockMeta, peer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlockStream", reflect.TypeOf((*MockHost)(nil).CreateBlockStream), ctx, blockMeta, peer)
}

// GetID mocks base method
func (m *MockHost) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockHostMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockHost)(nil).GetID))
}

// GetAddrs mocks base method
func (m *MockHost) GetAddrs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddrs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAddrs indicates an expected call of GetAddrs
func (mr *MockHostMockRecorder) GetAddrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddrs", reflect.TypeOf((*MockHost)(nil).GetAddrs))
}

// Reprovide mocks base method
func (m *MockHost) Reprovide(ctx context.Context, withBlocks bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reprovide", ctx, withBlocks)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reprovide indicates an expected call of Reprovide
func (mr *MockHostMockRecorder) Reprovide(ctx, withBlocks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reprovide", reflect.TypeOf((*MockHost)(nil).Reprovide), ctx, withBlocks)
}

// PutPublicKey mocks base method
func (m *MockHost) PutPublicKey(publicKey crypto.PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPublicKey", publicKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutPublicKey indicates an expected call of PutPublicKey
func (mr *MockHostMockRecorder) PutPublicKey(publicKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPublicKey", reflect.TypeOf((*MockHost)(nil).PutPublicKey), publicKey)
}

// Provide mocks base method
func (m *MockHost) Provide(ctx context.Context, cid go_cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide", ctx, cid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Provide indicates an expected call of Provide
func (mr *MockHostMockRecorder) Provide(ctx, cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockHost)(nil).Provide), ctx, cid)
}
