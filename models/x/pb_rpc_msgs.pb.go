// Code generated by protoc-gen-go.
// source: pb_rpc_msgs.proto
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

type PB_ReqRpcAddMsg struct {
	Request *PB_Request `protobuf:"bytes,1,opt,name=Request" json:"Request,omitempty"`
	Message *PB_Message `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *PB_ReqRpcAddMsg) Reset()                    { *m = PB_ReqRpcAddMsg{} }
func (m *PB_ReqRpcAddMsg) String() string            { return proto.CompactTextString(m) }
func (*PB_ReqRpcAddMsg) ProtoMessage()               {}
func (*PB_ReqRpcAddMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PB_ReqRpcAddMsg) GetRequest() *PB_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PB_ReqRpcAddMsg) GetMessage() *PB_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type PB_ResRpcAddMsg struct {
	Response  *PB_Response `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
	ServerSrc string       `protobuf:"bytes,2,opt,name=ServerSrc" json:"ServerSrc,omitempty"`
}

func (m *PB_ResRpcAddMsg) Reset()                    { *m = PB_ResRpcAddMsg{} }
func (m *PB_ResRpcAddMsg) String() string            { return proto.CompactTextString(m) }
func (*PB_ResRpcAddMsg) ProtoMessage()               {}
func (*PB_ResRpcAddMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PB_ResRpcAddMsg) GetResponse() *PB_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PB_ResRpcAddMsg) GetServerSrc() string {
	if m != nil {
		return m.ServerSrc
	}
	return ""
}

func init() {
	proto.RegisterType((*PB_ReqRpcAddMsg)(nil), "PB_ReqRpcAddMsg")
	proto.RegisterType((*PB_ResRpcAddMsg)(nil), "PB_ResRpcAddMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RpcMsgs service

type RpcMsgsClient interface {
	// Sends a greeting
	UploadNewMsg(ctx context.Context, in *PB_Message, opts ...grpc.CallOption) (*PB_ResRpcAddMsg, error)
}

type rpcMsgsClient struct {
	cc *grpc.ClientConn
}

func NewRpcMsgsClient(cc *grpc.ClientConn) RpcMsgsClient {
	return &rpcMsgsClient{cc}
}

func (c *rpcMsgsClient) UploadNewMsg(ctx context.Context, in *PB_Message, opts ...grpc.CallOption) (*PB_ResRpcAddMsg, error) {
	out := new(PB_ResRpcAddMsg)
	err := grpc.Invoke(ctx, "/RpcMsgs/UploadNewMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RpcMsgs service

type RpcMsgsServer interface {
	// Sends a greeting
	UploadNewMsg(context.Context, *PB_Message) (*PB_ResRpcAddMsg, error)
}

func RegisterRpcMsgsServer(s *grpc.Server, srv RpcMsgsServer) {
	s.RegisterService(&_RpcMsgs_serviceDesc, srv)
}

func _RpcMsgs_UploadNewMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcMsgsServer).UploadNewMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcMsgs/UploadNewMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcMsgsServer).UploadNewMsg(ctx, req.(*PB_Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcMsgs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RpcMsgs",
	HandlerType: (*RpcMsgsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadNewMsg",
			Handler:    _RpcMsgs_UploadNewMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_rpc_msgs.proto",
}

func init() { proto.RegisterFile("pb_rpc_msgs.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0x47, 0x5d, 0x05, 0x93, 0x4c, 0x02, 0x31, 0x57, 0x85, 0x60, 0x21, 0x07, 0x42, 0xaa, 0x15,
	0x62, 0x67, 0x67, 0x2a, 0x9b, 0x93, 0x30, 0xc1, 0x42, 0x9b, 0xe5, 0x76, 0x6f, 0x58, 0x84, 0x8b,
	0x3b, 0xd9, 0x89, 0x7f, 0x3e, 0xbe, 0x5c, 0xf6, 0xce, 0x93, 0xb4, 0xef, 0xf7, 0xe0, 0x0d, 0x03,
	0x33, 0xb6, 0x26, 0xb2, 0x33, 0x3b, 0xf1, 0xa2, 0x39, 0x86, 0x43, 0x58, 0x4c, 0xd9, 0x1a, 0x5f,
	0x07, 0x5b, 0xd6, 0x09, 0xe4, 0x06, 0xa6, 0x9b, 0xb5, 0x41, 0xda, 0x23, 0xbb, 0xc7, 0xaa, 0x2a,
	0xc4, 0x67, 0xb7, 0x30, 0x40, 0xda, 0x7f, 0x92, 0x1c, 0xe6, 0xea, 0x46, 0x2d, 0xc7, 0xab, 0xb1,
	0x4e, 0x4a, 0x83, 0xb0, 0xdb, 0x1a, 0xad, 0x20, 0x91, 0xd2, 0xd3, 0xfc, 0xbc, 0xd7, 0x5a, 0x84,
	0xdd, 0x96, 0xbf, 0xb6, 0x01, 0xe9, 0x03, 0x4b, 0x18, 0x22, 0x09, 0x87, 0x0f, 0xa1, 0xb6, 0x30,
	0x49, 0x85, 0xc4, 0xf0, 0x6f, 0xcd, 0xae, 0x61, 0xb4, 0xa5, 0xf8, 0x45, 0x71, 0x1b, 0xdd, 0xb1,
	0x32, 0xc2, 0x1e, 0xac, 0x1e, 0x60, 0x80, 0xec, 0x0a, 0xf1, 0x92, 0xdd, 0xc1, 0xe4, 0x85, 0xeb,
	0x50, 0x56, 0xcf, 0xf4, 0xdd, 0x24, 0xfe, 0xdf, 0xb2, 0xb8, 0xd2, 0x27, 0x17, 0xe4, 0x67, 0xeb,
	0x19, 0x0c, 0xdf, 0xa3, 0xde, 0x89, 0x66, 0xfb, 0x74, 0xb1, 0x51, 0x6f, 0xea, 0xc7, 0x5e, 0x1e,
	0x3f, 0x72, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xd7, 0x17, 0x28, 0x37, 0x01, 0x00, 0x00,
}
