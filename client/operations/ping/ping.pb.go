// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

package ping

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

type OpPingProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpPingProto) Reset()         { *m = OpPingProto{} }
func (m *OpPingProto) String() string { return proto.CompactTextString(m) }
func (*OpPingProto) ProtoMessage()    {}
func (*OpPingProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{0}
}

func (m *OpPingProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpPingProto.Unmarshal(m, b)
}
func (m *OpPingProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpPingProto.Marshal(b, m, deterministic)
}
func (m *OpPingProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpPingProto.Merge(m, src)
}
func (m *OpPingProto) XXX_Size() int {
	return xxx_messageInfo_OpPingProto.Size(m)
}
func (m *OpPingProto) XXX_DiscardUnknown() {
	xxx_messageInfo_OpPingProto.DiscardUnknown(m)
}

var xxx_messageInfo_OpPingProto proto.InternalMessageInfo

type ResultPingProto struct {
	SupportedVersionMaj  uint32   `protobuf:"varint,1,opt,name=supported_version_maj,json=supportedVersionMaj,proto3" json:"supported_version_maj,omitempty"`
	SupportedVersionMin  uint32   `protobuf:"varint,2,opt,name=supported_version_min,json=supportedVersionMin,proto3" json:"supported_version_min,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultPingProto) Reset()         { *m = ResultPingProto{} }
func (m *ResultPingProto) String() string { return proto.CompactTextString(m) }
func (*ResultPingProto) ProtoMessage()    {}
func (*ResultPingProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{1}
}

func (m *ResultPingProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultPingProto.Unmarshal(m, b)
}
func (m *ResultPingProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultPingProto.Marshal(b, m, deterministic)
}
func (m *ResultPingProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultPingProto.Merge(m, src)
}
func (m *ResultPingProto) XXX_Size() int {
	return xxx_messageInfo_ResultPingProto.Size(m)
}
func (m *ResultPingProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultPingProto.DiscardUnknown(m)
}

var xxx_messageInfo_ResultPingProto proto.InternalMessageInfo

func (m *ResultPingProto) GetSupportedVersionMaj() uint32 {
	if m != nil {
		return m.SupportedVersionMaj
	}
	return 0
}

func (m *ResultPingProto) GetSupportedVersionMin() uint32 {
	if m != nil {
		return m.SupportedVersionMin
	}
	return 0
}

func init() {
	proto.RegisterType((*OpPingProto)(nil), "ping.OpPingProto")
	proto.RegisterType((*ResultPingProto)(nil), "ping.ResultPingProto")
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_6d51d96c3ad891f5) }

var fileDescriptor_6d51d96c3ad891f5 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x78, 0xb9, 0xb8, 0xfd, 0x0b,
	0x02, 0x32, 0xf3, 0xd2, 0x03, 0x40, 0x82, 0x4a, 0x95, 0x5c, 0xfc, 0x41, 0xa9, 0xc5, 0xa5, 0x39,
	0x25, 0x70, 0x21, 0x21, 0x23, 0x2e, 0xd1, 0xe2, 0xd2, 0x82, 0x82, 0xfc, 0xa2, 0x92, 0xd4, 0x94,
	0xf8, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0xbc, 0xf8, 0xdc, 0xc4, 0x2c, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xde, 0x20, 0x61, 0xb8, 0x64, 0x18, 0x44, 0xce, 0x37, 0x31, 0x0b, 0x87, 0x9e, 0xcc, 0x3c,
	0x09, 0x26, 0x1c, 0x7a, 0x32, 0xf3, 0x92, 0xd8, 0xc0, 0xce, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xa4, 0x0f, 0x74, 0xa4, 0x00, 0x00, 0x00,
}
