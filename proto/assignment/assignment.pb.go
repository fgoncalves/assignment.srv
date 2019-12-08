// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/assignment.proto

package go_micro_srv_assignment

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

type Request struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExperimentName       string   `protobuf:"bytes,2,opt,name=experiment_name,json=experimentName,proto3" json:"experiment_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0feaacedf0ec52c7, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Request) GetExperimentName() string {
	if m != nil {
		return m.ExperimentName
	}
	return ""
}

type Response struct {
	Variant              string   `protobuf:"bytes,1,opt,name=variant,proto3" json:"variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0feaacedf0ec52c7, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.assignment.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.assignment.Response")
}

func init() { proto.RegisterFile("proto/assignment/assignment.proto", fileDescriptor_0feaacedf0ec52c7) }

var fileDescriptor_0feaacedf0ec52c7 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0x2e, 0xce, 0x4c, 0xcf, 0xcb, 0x4d, 0xcd, 0x2b, 0x41, 0x62, 0xea, 0x81,
	0xe5, 0x84, 0xc4, 0xd3, 0xf3, 0xf5, 0x72, 0x33, 0x93, 0x8b, 0xf2, 0xf5, 0x8a, 0x8b, 0xca, 0xf4,
	0x10, 0xd2, 0x4a, 0xde, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xe2, 0x5c,
	0xec, 0xa5, 0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x6c,
	0x20, 0xae, 0x67, 0x8a, 0x90, 0x3a, 0x17, 0x7f, 0x6a, 0x45, 0x41, 0x6a, 0x51, 0x26, 0x48, 0x47,
	0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x01, 0x1f, 0x42, 0xd8, 0x2f, 0x31, 0x37, 0x55,
	0x49, 0x85, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0xbd,
	0x2c, 0xb1, 0x28, 0x33, 0x31, 0xaf, 0x04, 0x6a, 0x1a, 0x8c, 0x6b, 0x14, 0xcb, 0xc5, 0xe5, 0x08,
	0x77, 0x80, 0x90, 0x3f, 0x17, 0x1b, 0x84, 0x27, 0xa4, 0xa0, 0x87, 0xc3, 0x91, 0x7a, 0x50, 0x17,
	0x4a, 0x29, 0xe2, 0x51, 0x01, 0xb1, 0x56, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x63, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x56, 0x49, 0x85, 0x1e, 0x16, 0x01, 0x00, 0x00,
}
