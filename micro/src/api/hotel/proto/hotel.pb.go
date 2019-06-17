// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hotel.proto

package hotel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "srv/profile/proto"
	proto2 "srv/rate/proto"
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
	InDate               string   `protobuf:"bytes,1,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string   `protobuf:"bytes,2,opt,name=outDate,proto3" json:"outDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_169099f938a8a4b0, []int{0}
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

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Response struct {
	Hotels               []*proto1.Hotel    `protobuf:"bytes,1,rep,name=hotels,proto3" json:"hotels,omitempty"`
	RatePlans            []*proto2.RatePlan `protobuf:"bytes,2,rep,name=ratePlans,proto3" json:"ratePlans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_169099f938a8a4b0, []int{1}
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

func (m *Response) GetHotels() []*proto1.Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

func (m *Response) GetRatePlans() []*proto2.RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "hotel.Request")
	proto.RegisterType((*Response)(nil), "hotel.Response")
}

func init() { proto.RegisterFile("hotel.proto", fileDescriptor_169099f938a8a4b0) }

var fileDescriptor_169099f938a8a4b0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x8a, 0x83, 0x30,
	0x10, 0x86, 0x57, 0x17, 0x75, 0x1d, 0xc1, 0x85, 0x1c, 0x96, 0xac, 0x97, 0x8a, 0x87, 0x22, 0xa5,
	0x28, 0xe8, 0xb1, 0xd7, 0x1e, 0x7a, 0x2c, 0x79, 0x82, 0x5a, 0x48, 0xa9, 0x20, 0xc6, 0x26, 0xb1,
	0xcf, 0x5f, 0x32, 0x49, 0xda, 0x5b, 0xfe, 0xf9, 0xfe, 0x90, 0x2f, 0x03, 0xd9, 0x5d, 0x68, 0x3e,
	0x35, 0x8b, 0x14, 0x5a, 0x90, 0x08, 0x43, 0xb1, 0x51, 0xf2, 0xd9, 0x2e, 0x52, 0xdc, 0xc6, 0x89,
	0xb7, 0x48, 0x7c, 0xb2, 0xbd, 0xe2, 0xdf, 0x14, 0xe4, 0xa0, 0x3d, 0x35, 0x47, 0x8b, 0xaa, 0x03,
	0x24, 0x8c, 0x3f, 0x56, 0xae, 0x34, 0xf9, 0x83, 0x78, 0x9c, 0x8f, 0x83, 0xe6, 0x34, 0x28, 0x83,
	0x3a, 0x65, 0x2e, 0x11, 0x0a, 0x89, 0x58, 0x35, 0x82, 0x10, 0x81, 0x8f, 0xd5, 0x05, 0x7e, 0x18,
	0x57, 0x8b, 0x98, 0x15, 0x27, 0x5b, 0x88, 0xd1, 0x46, 0xd1, 0xa0, 0xfc, 0xae, 0xb3, 0x2e, 0x6f,
	0xbc, 0xc3, 0xc9, 0x8c, 0x99, 0xa3, 0x64, 0x0f, 0xa9, 0x79, 0xfe, 0x3c, 0x0d, 0xb3, 0xa2, 0xa1,
	0xab, 0xa2, 0x10, 0x73, 0x63, 0xf6, 0x29, 0x74, 0x3d, 0x44, 0x78, 0x9d, 0xec, 0x20, 0x32, 0x5c,
	0x91, 0xbc, 0xb1, 0x1b, 0x70, 0xd6, 0xc5, 0xef, 0x3b, 0x5b, 0x91, 0xea, 0xeb, 0x1a, 0xe3, 0xd7,
	0xfa, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x09, 0x3d, 0x42, 0x2c, 0x01, 0x00, 0x00,
}