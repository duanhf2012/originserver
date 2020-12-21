// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcproto/rpc.proto

package rpc

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PBRpcRequestData struct {
	Seq                  uint64   `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	ServiceMethod        string   `protobuf:"bytes,2,opt,name=ServiceMethod,proto3" json:"ServiceMethod,omitempty"`
	NoReply              bool     `protobuf:"varint,3,opt,name=NoReply,proto3" json:"NoReply,omitempty"`
	InParam              []byte   `protobuf:"bytes,4,opt,name=InParam,proto3" json:"InParam,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBRpcRequestData) Reset()         { *m = PBRpcRequestData{} }
func (m *PBRpcRequestData) String() string { return proto.CompactTextString(m) }
func (*PBRpcRequestData) ProtoMessage()    {}
func (*PBRpcRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0008ed1de5480352, []int{0}
}

func (m *PBRpcRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBRpcRequestData.Unmarshal(m, b)
}
func (m *PBRpcRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBRpcRequestData.Marshal(b, m, deterministic)
}
func (m *PBRpcRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBRpcRequestData.Merge(m, src)
}
func (m *PBRpcRequestData) XXX_Size() int {
	return xxx_messageInfo_PBRpcRequestData.Size(m)
}
func (m *PBRpcRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_PBRpcRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_PBRpcRequestData proto.InternalMessageInfo

func (m *PBRpcRequestData) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PBRpcRequestData) GetServiceMethod() string {
	if m != nil {
		return m.ServiceMethod
	}
	return ""
}

func (m *PBRpcRequestData) GetNoReply() bool {
	if m != nil {
		return m.NoReply
	}
	return false
}

func (m *PBRpcRequestData) GetInParam() []byte {
	if m != nil {
		return m.InParam
	}
	return nil
}

type PBRpcResponseData struct {
	Seq                  uint64   `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Reply                []byte   `protobuf:"bytes,3,opt,name=Reply,proto3" json:"Reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBRpcResponseData) Reset()         { *m = PBRpcResponseData{} }
func (m *PBRpcResponseData) String() string { return proto.CompactTextString(m) }
func (*PBRpcResponseData) ProtoMessage()    {}
func (*PBRpcResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0008ed1de5480352, []int{1}
}

func (m *PBRpcResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBRpcResponseData.Unmarshal(m, b)
}
func (m *PBRpcResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBRpcResponseData.Marshal(b, m, deterministic)
}
func (m *PBRpcResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBRpcResponseData.Merge(m, src)
}
func (m *PBRpcResponseData) XXX_Size() int {
	return xxx_messageInfo_PBRpcResponseData.Size(m)
}
func (m *PBRpcResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_PBRpcResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_PBRpcResponseData proto.InternalMessageInfo

func (m *PBRpcResponseData) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PBRpcResponseData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PBRpcResponseData) GetReply() []byte {
	if m != nil {
		return m.Reply
	}
	return nil
}

func init() {
	proto.RegisterType((*PBRpcRequestData)(nil), "rpc.PBRpcRequestData")
	proto.RegisterType((*PBRpcResponseData)(nil), "rpc.PBRpcResponseData")
}

func init() { proto.RegisterFile("rpcproto/rpc.proto", fileDescriptor_0008ed1de5480352) }

var fileDescriptor_0008ed1de5480352 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2a, 0x48, 0x2e,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0xb3, 0x84, 0x98, 0x8b, 0x0a, 0x92,
	0x95, 0xea, 0xb8, 0x04, 0x02, 0x9c, 0x82, 0x0a, 0x92, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x5c, 0x12, 0x4b, 0x12, 0x85, 0x04, 0xb8, 0x98, 0x83, 0x53, 0x0b, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x58, 0x82, 0x40, 0x4c, 0x21, 0x15, 0x2e, 0xde, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xdf,
	0xd4, 0x92, 0x8c, 0xfc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x54, 0x41, 0x21, 0x09,
	0x2e, 0x76, 0xbf, 0xfc, 0xa0, 0xd4, 0x82, 0x9c, 0x4a, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20,
	0x18, 0x17, 0x24, 0xe3, 0x99, 0x17, 0x90, 0x58, 0x94, 0x98, 0x2b, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1,
	0x13, 0x04, 0xe3, 0x2a, 0x05, 0x72, 0x09, 0x42, 0xed, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5,
	0xe1, 0x00, 0x11, 0x2e, 0x56, 0xd7, 0xa2, 0xa2, 0xfc, 0x22, 0xa8, 0xc5, 0x10, 0x0e, 0x48, 0x14,
	0x61, 0x1d, 0x4f, 0x10, 0x84, 0xe3, 0xc4, 0x1e, 0xc5, 0xaa, 0x67, 0x5d, 0x54, 0x90, 0x9c, 0xc4,
	0x06, 0xf6, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x5b, 0x74, 0xd8, 0xfd, 0x00, 0x00,
	0x00,
}