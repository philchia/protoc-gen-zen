// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/proto/test.proto

/*
	Package http is a generated protocol buffer package.

	It is generated from these files:
		test/proto/test.proto

	It has these top-level messages:
		ReqPing
		RespPing
		ReqGetUser
		RespGetUser
*/
package http

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	zen "github.com/philchia/zen"
	http "net/http"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Foo int32

const (
	Foo_zero Foo = 0
	Foo_one  Foo = 1
	Foo_two  Foo = 2
)

var Foo_name = map[int32]string{
	0: "zero",
	1: "one",
	2: "two",
}
var Foo_value = map[string]int32{
	"zero": 0,
	"one":  1,
	"two":  2,
}

func (x Foo) String() string {
	return proto1.EnumName(Foo_name, int32(x))
}
func (Foo) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReqPing struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *ReqPing) Reset()                    { *m = ReqPing{} }
func (m *ReqPing) String() string            { return proto1.CompactTextString(m) }
func (*ReqPing) ProtoMessage()               {}
func (*ReqPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqPing) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RespPing struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *RespPing) Reset()                    { *m = RespPing{} }
func (m *RespPing) String() string            { return proto1.CompactTextString(m) }
func (*RespPing) ProtoMessage()               {}
func (*RespPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RespPing) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ReqGetUser struct {
}

func (m *ReqGetUser) Reset()                    { *m = ReqGetUser{} }
func (m *ReqGetUser) String() string            { return proto1.CompactTextString(m) }
func (*ReqGetUser) ProtoMessage()               {}
func (*ReqGetUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RespGetUser struct {
}

func (m *RespGetUser) Reset()                    { *m = RespGetUser{} }
func (m *RespGetUser) String() string            { return proto1.CompactTextString(m) }
func (*RespGetUser) ProtoMessage()               {}
func (*RespGetUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto1.RegisterType((*ReqPing)(nil), "proto.ReqPing")
	proto1.RegisterType((*RespPing)(nil), "proto.RespPing")
	proto1.RegisterType((*ReqGetUser)(nil), "proto.ReqGetUser")
	proto1.RegisterType((*RespGetUser)(nil), "proto.RespGetUser")
	proto1.RegisterEnum("proto.Foo", Foo_name, Foo_value)
}

type TestService interface {
	Ping(zen.Context, *ReqPing) (*RespPing, error)
	GetUser(zen.Context, *ReqGetUser) (*RespGetUser, error)
}

func RegisterTestServer(router zen.Router, server TestService) {
	router.Post("test.ping", func(ctx zen.Context) {
		var req = new(ReqPing)
		if err := ctx.BindJSON(req); err != nil {
			ctx.WriteStatus(http.StatusBadRequest)
			return
		}
		resp, err := server.Ping(ctx, req)
		if err != nil {
			ctx.WriteStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(resp)
	})

	router.Post("test.getuser", func(ctx zen.Context) {
		var req = new(ReqGetUser)
		if err := ctx.BindJSON(req); err != nil {
			ctx.WriteStatus(http.StatusBadRequest)
			return
		}
		resp, err := server.GetUser(ctx, req)
		if err != nil {
			ctx.WriteStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(resp)
	})

}

func init() { proto1.RegisterFile("test/proto/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x49, 0x2d, 0x2e,
	0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0x31, 0xf5, 0xc0, 0x4c, 0x21, 0x56, 0x30, 0xa5,
	0x24, 0xcd, 0xc5, 0x1e, 0x94, 0x5a, 0x18, 0x90, 0x99, 0x97, 0x2e, 0x24, 0xc0, 0xc5, 0x9c, 0x5b,
	0x9c, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x2a, 0xc9, 0x70, 0x71, 0x04, 0xa5,
	0x16, 0x17, 0xe0, 0x90, 0xe5, 0xe1, 0xe2, 0x0a, 0x4a, 0x2d, 0x74, 0x4f, 0x2d, 0x09, 0x2d, 0x4e,
	0x2d, 0x52, 0xe2, 0xe5, 0xe2, 0x06, 0xa9, 0x85, 0x72, 0xb5, 0x14, 0xb9, 0x98, 0xdd, 0xf2, 0xf3,
	0x85, 0x38, 0xb8, 0x58, 0xaa, 0x52, 0x8b, 0xf2, 0x05, 0x18, 0x84, 0xd8, 0xb9, 0x98, 0xf3, 0xf3,
	0x52, 0x05, 0x18, 0x41, 0x8c, 0x92, 0xf2, 0x7c, 0x01, 0x26, 0xa3, 0x44, 0x2e, 0x96, 0x90, 0xd4,
	0xe2, 0x12, 0x21, 0x75, 0x2e, 0x16, 0xb0, 0x0d, 0x7c, 0x10, 0x97, 0xe9, 0x41, 0xdd, 0x23, 0xc5,
	0x0f, 0xe7, 0x43, 0x9d, 0x60, 0xc0, 0xc5, 0x0e, 0x35, 0x5e, 0x48, 0x10, 0xa1, 0x16, 0x2a, 0x24,
	0x25, 0x84, 0xa4, 0x1c, 0x2a, 0x96, 0xc4, 0x06, 0x16, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0xee, 0x15, 0x91, 0x04, 0x01, 0x00, 0x00,
}
