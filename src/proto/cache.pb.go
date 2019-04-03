// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	cache.proto

It has these top-level messages:
	Request
	Response
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Response struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterType((*Response)(nil), "protocol.Response")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0xd2,
	0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x01, 0x17, 0x47, 0x50, 0x6a, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x10, 0x1f, 0x17, 0x53, 0x7e, 0x36, 0x58, 0x92, 0x23, 0x88, 0x29,
	0x3f, 0x5b, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x27, 0x08, 0xc2, 0x49, 0x62, 0x03, 0x1b, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x1f,
	0x4a, 0x16, 0x6e, 0x00, 0x00, 0x00,
}