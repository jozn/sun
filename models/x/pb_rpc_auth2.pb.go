// Code generated by protoc-gen-go.
// source: pb_rpc_auth2.proto
// DO NOT EDIT!

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PB_UserParam_CheckUserName2 struct {
}

func (m *PB_UserParam_CheckUserName2) Reset()                    { *m = PB_UserParam_CheckUserName2{} }
func (m *PB_UserParam_CheckUserName2) String() string            { return proto.CompactTextString(m) }
func (*PB_UserParam_CheckUserName2) ProtoMessage()               {}
func (*PB_UserParam_CheckUserName2) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type PB_UserResponse_CheckUserName2 struct {
}

func (m *PB_UserResponse_CheckUserName2) Reset()                    { *m = PB_UserResponse_CheckUserName2{} }
func (m *PB_UserResponse_CheckUserName2) String() string            { return proto.CompactTextString(m) }
func (*PB_UserResponse_CheckUserName2) ProtoMessage()               {}
func (*PB_UserResponse_CheckUserName2) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func init() {
	proto.RegisterType((*PB_UserParam_CheckUserName2)(nil), "PB_UserParam_CheckUserName2")
	proto.RegisterType((*PB_UserResponse_CheckUserName2)(nil), "PB_UserResponse_CheckUserName2")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPC_Auth service

type RPC_AuthClient interface {
	CheckPhone(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
	SendCode(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
	SendCodeToSms(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
	SendCodeToTelgram(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
	SingUp(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
	SingIn(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
	LogOut(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error)
}

type rPC_AuthClient struct {
	cc *grpc.ClientConn
}

func NewRPC_AuthClient(cc *grpc.ClientConn) RPC_AuthClient {
	return &rPC_AuthClient{cc}
}

func (c *rPC_AuthClient) CheckPhone(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/CheckPhone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_AuthClient) SendCode(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/SendCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_AuthClient) SendCodeToSms(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/SendCodeToSms", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_AuthClient) SendCodeToTelgram(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/SendCodeToTelgram", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_AuthClient) SingUp(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/SingUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_AuthClient) SingIn(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/SingIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_AuthClient) LogOut(ctx context.Context, in *PB_UserParam_CheckUserName2, opts ...grpc.CallOption) (*PB_UserResponse_CheckUserName2, error) {
	out := new(PB_UserResponse_CheckUserName2)
	err := grpc.Invoke(ctx, "/RPC_Auth/LogOut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC_Auth service

type RPC_AuthServer interface {
	CheckPhone(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
	SendCode(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
	SendCodeToSms(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
	SendCodeToTelgram(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
	SingUp(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
	SingIn(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
	LogOut(context.Context, *PB_UserParam_CheckUserName2) (*PB_UserResponse_CheckUserName2, error)
}

func RegisterRPC_AuthServer(s *grpc.Server, srv RPC_AuthServer) {
	s.RegisterService(&_RPC_Auth_serviceDesc, srv)
}

func _RPC_Auth_CheckPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).CheckPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/CheckPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).CheckPhone(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Auth_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).SendCode(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Auth_SendCodeToSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).SendCodeToSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/SendCodeToSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).SendCodeToSms(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Auth_SendCodeToTelgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).SendCodeToTelgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/SendCodeToTelgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).SendCodeToTelgram(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Auth_SingUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).SingUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/SingUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).SingUp(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Auth_SingIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).SingIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/SingIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).SingIn(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Auth_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_UserParam_CheckUserName2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_AuthServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Auth/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_AuthServer).LogOut(ctx, req.(*PB_UserParam_CheckUserName2))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPC_Auth",
	HandlerType: (*RPC_AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPhone",
			Handler:    _RPC_Auth_CheckPhone_Handler,
		},
		{
			MethodName: "SendCode",
			Handler:    _RPC_Auth_SendCode_Handler,
		},
		{
			MethodName: "SendCodeToSms",
			Handler:    _RPC_Auth_SendCodeToSms_Handler,
		},
		{
			MethodName: "SendCodeToTelgram",
			Handler:    _RPC_Auth_SendCodeToTelgram_Handler,
		},
		{
			MethodName: "SingUp",
			Handler:    _RPC_Auth_SingUp_Handler,
		},
		{
			MethodName: "SingIn",
			Handler:    _RPC_Auth_SingIn_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _RPC_Auth_LogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_rpc_auth2.proto",
}

func init() { proto.RegisterFile("pb_rpc_auth2.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0x8a, 0x2f,
	0x2a, 0x48, 0x8e, 0x4f, 0x2c, 0x2d, 0xc9, 0x30, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2,
	0x2b, 0x48, 0x8a, 0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0x86, 0xf0, 0x95, 0x64, 0xb9, 0xa4, 0x03, 0x9c,
	0xe2, 0x43, 0x8b, 0x53, 0x8b, 0x02, 0x12, 0x8b, 0x12, 0x73, 0xe3, 0x9d, 0x33, 0x52, 0x93, 0xb3,
	0x41, 0x5c, 0xbf, 0xc4, 0xdc, 0x54, 0x23, 0x25, 0x05, 0x2e, 0x39, 0xa8, 0x74, 0x50, 0x6a, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x9a, 0x0a, 0xa3, 0x85, 0x2c, 0x5c, 0x1c, 0x41, 0x01, 0xce, 0xf1,
	0x8e, 0xa5, 0x25, 0x19, 0x42, 0xde, 0x5c, 0x5c, 0x60, 0xe9, 0x80, 0x8c, 0xfc, 0xbc, 0x54, 0x21,
	0x19, 0x3d, 0x3c, 0x46, 0x4b, 0xc9, 0xeb, 0xe1, 0x37, 0x59, 0xc8, 0x93, 0x8b, 0x23, 0x38, 0x35,
	0x2f, 0xc5, 0x39, 0x3f, 0x85, 0x62, 0xa3, 0xfc, 0xb8, 0x78, 0x61, 0x46, 0x85, 0xe4, 0x07, 0xe7,
	0x16, 0x53, 0x6a, 0x5e, 0x10, 0x97, 0x20, 0xc2, 0xbc, 0x90, 0xd4, 0x9c, 0xf4, 0xa2, 0xc4, 0x5c,
	0x4a, 0xcd, 0x74, 0xe7, 0x62, 0x0b, 0xce, 0xcc, 0x4b, 0x0f, 0x2d, 0xa0, 0x92, 0x41, 0x9e, 0x79,
	0x54, 0x30, 0xc8, 0x27, 0x3f, 0xdd, 0xbf, 0xb4, 0x84, 0x42, 0x83, 0x9c, 0x04, 0xb9, 0x38, 0x32,
	0x8b, 0xf4, 0x40, 0x69, 0x2e, 0xc9, 0x83, 0x39, 0x80, 0x31, 0x8a, 0xb1, 0x22, 0x89, 0x0d, 0x9c,
	0xfc, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xfe, 0xc7, 0x74, 0xa4, 0x02, 0x00, 0x00,
}
