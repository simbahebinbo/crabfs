// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BlockMetadata struct {
	Cid                  []byte   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	PaddingStart         int64    `protobuf:"varint,4,opt,name=paddingStart,proto3" json:"paddingStart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMetadata.Unmarshal(m, b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
}
func (m *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(m, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return xxx_messageInfo_BlockMetadata.Size(m)
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetCid() []byte {
	if m != nil {
		return m.Cid
	}
	return nil
}

func (m *BlockMetadata) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *BlockMetadata) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BlockMetadata) GetPaddingStart() int64 {
	if m != nil {
		return m.PaddingStart
	}
	return 0
}

type CrabObject struct {
	Name                 string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Blocks               map[int64]*BlockMetadata `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mtime                string                   `protobuf:"bytes,3,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Size                 int64                    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Key                  []byte                   `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Delete               bool                     `protobuf:"varint,6,opt,name=delete,proto3" json:"delete,omitempty"`
	Lock                 []byte                   `protobuf:"bytes,7,opt,name=lock,proto3" json:"lock,omitempty"`
	CacheTTL             uint64                   `protobuf:"varint,8,opt,name=cacheTTL,proto3" json:"cacheTTL,omitempty"`
	BucketID             string                   `protobuf:"bytes,9,opt,name=bucketID,proto3" json:"bucketID,omitempty"`
	Timestamp            string                   `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CrabObject) Reset()         { *m = CrabObject{} }
func (m *CrabObject) String() string { return proto.CompactTextString(m) }
func (*CrabObject) ProtoMessage()    {}
func (*CrabObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *CrabObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrabObject.Unmarshal(m, b)
}
func (m *CrabObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrabObject.Marshal(b, m, deterministic)
}
func (m *CrabObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrabObject.Merge(m, src)
}
func (m *CrabObject) XXX_Size() int {
	return xxx_messageInfo_CrabObject.Size(m)
}
func (m *CrabObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CrabObject.DiscardUnknown(m)
}

var xxx_messageInfo_CrabObject proto.InternalMessageInfo

func (m *CrabObject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CrabObject) GetBlocks() map[int64]*BlockMetadata {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *CrabObject) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *CrabObject) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CrabObject) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CrabObject) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (m *CrabObject) GetLock() []byte {
	if m != nil {
		return m.Lock
	}
	return nil
}

func (m *CrabObject) GetCacheTTL() uint64 {
	if m != nil {
		return m.CacheTTL
	}
	return 0
}

func (m *CrabObject) GetBucketID() string {
	if m != nil {
		return m.BucketID
	}
	return ""
}

func (m *CrabObject) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockMetadata)(nil), "protos.BlockMetadata")
	proto.RegisterType((*CrabObject)(nil), "protos.CrabObject")
	proto.RegisterMapType((map[int64]*BlockMetadata)(nil), "protos.CrabObject.BlocksEntry")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xdb, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xd9, 0x24, 0xcd, 0xd7, 0x4c, 0xfb, 0x41, 0x59, 0x54, 0x96, 0x22, 0x12, 0x7a, 0x15,
	0x10, 0x7a, 0x51, 0x41, 0xc4, 0x4b, 0x0f, 0x17, 0x82, 0xa2, 0xac, 0x7d, 0x81, 0xcd, 0x66, 0xd0,
	0x98, 0x26, 0x29, 0xd9, 0xad, 0x50, 0x9f, 0xd8, 0xc7, 0x90, 0x9d, 0xed, 0xf1, 0x2a, 0xff, 0xff,
	0x9c, 0xf2, 0x9b, 0x1d, 0x18, 0xea, 0xb6, 0xae, 0xdb, 0x66, 0xba, 0xec, 0x5a, 0xdb, 0xf2, 0x98,
	0x3e, 0x66, 0xd2, 0xc2, 0xff, 0xbb, 0x45, 0xab, 0xab, 0x17, 0xb4, 0xaa, 0x50, 0x56, 0xf1, 0x11,
	0x84, 0xba, 0x2c, 0x04, 0x4b, 0x59, 0x36, 0x94, 0x4e, 0xf2, 0x13, 0xe8, 0x19, 0xab, 0x3a, 0x2b,
	0x82, 0x94, 0x65, 0xa1, 0xf4, 0x86, 0x73, 0x88, 0x4c, 0xf9, 0x83, 0x22, 0xa4, 0x20, 0x69, 0x3e,
	0x81, 0xe1, 0x52, 0x15, 0x45, 0xd9, 0x7c, 0xbc, 0x53, 0x43, 0x44, 0xb9, 0xa3, 0xd8, 0xe4, 0x37,
	0x00, 0xb8, 0xef, 0x54, 0xfe, 0x9a, 0x7f, 0xa1, 0xa6, 0x31, 0x8d, 0xaa, 0x91, 0xfe, 0x97, 0x48,
	0xd2, 0xfc, 0x1a, 0xe2, 0xdc, 0x31, 0x19, 0x11, 0xa4, 0x61, 0x36, 0x98, 0x5d, 0x78, 0x66, 0x33,
	0xdd, 0xf7, 0x4d, 0x09, 0xda, 0x3c, 0x36, 0xb6, 0x5b, 0xcb, 0x4d, 0xb5, 0x03, 0xad, 0x6d, 0x59,
	0x7b, 0xa6, 0x44, 0x7a, 0xb3, 0x03, 0x8d, 0x0e, 0x40, 0x47, 0x10, 0x56, 0xb8, 0x16, 0x3d, 0xbf,
	0x64, 0x85, 0x6b, 0x7e, 0x06, 0x71, 0x81, 0x0b, 0xb4, 0x28, 0xe2, 0x94, 0x65, 0x7d, 0xb9, 0x71,
	0xae, 0xdb, 0x0d, 0x17, 0xff, 0xa8, 0x94, 0x34, 0x1f, 0x43, 0x5f, 0x2b, 0xfd, 0x89, 0xf3, 0xf9,
	0xb3, 0xe8, 0xa7, 0x2c, 0x8b, 0xe4, 0xce, 0xbb, 0x5c, 0xbe, 0xd2, 0x15, 0xda, 0xa7, 0x07, 0x91,
	0x10, 0xc6, 0xce, 0xf3, 0x73, 0x48, 0x1c, 0x91, 0xb1, 0xaa, 0x5e, 0x0a, 0xa0, 0xe4, 0x3e, 0x30,
	0x7e, 0x83, 0xc1, 0xc1, 0x52, 0x5b, 0x44, 0x46, 0xd4, 0x84, 0x78, 0x09, 0xbd, 0x6f, 0xb5, 0x58,
	0x21, 0xdd, 0x61, 0x30, 0x3b, 0xdd, 0xbe, 0xca, 0xd1, 0xfd, 0xa4, 0xaf, 0xb9, 0x0d, 0x6e, 0x58,
	0xee, 0x6f, 0x7c, 0xf5, 0x17, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x6e, 0xa4, 0x33, 0xfa, 0x01, 0x00,
	0x00,
}
