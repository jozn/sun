// Code generated by protoc-gen-go.
// source: pb_rpc_messages.proto
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

// AddNewTextMessage
type PB_MsgParam_AddNewTextMessage struct {
	Text             string                   `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
	PeerId           int32                    `protobuf:"varint,5,opt,name=PeerId" json:"PeerId,omitempty"`
	Time             int32                    `protobuf:"varint,6,opt,name=Time" json:"Time,omitempty"`
	ReplyToMessageId int64                    `protobuf:"varint,10,opt,name=ReplyToMessageId" json:"ReplyToMessageId,omitempty"`
	Forward          *PB_MessageForwardedFrom `protobuf:"bytes,14,opt,name=Forward" json:"Forward,omitempty"`
}

func (m *PB_MsgParam_AddNewTextMessage) Reset()                    { *m = PB_MsgParam_AddNewTextMessage{} }
func (m *PB_MsgParam_AddNewTextMessage) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_AddNewTextMessage) ProtoMessage()               {}
func (*PB_MsgParam_AddNewTextMessage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *PB_MsgParam_AddNewTextMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PB_MsgParam_AddNewTextMessage) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *PB_MsgParam_AddNewTextMessage) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PB_MsgParam_AddNewTextMessage) GetReplyToMessageId() int64 {
	if m != nil {
		return m.ReplyToMessageId
	}
	return 0
}

func (m *PB_MsgParam_AddNewTextMessage) GetForward() *PB_MessageForwardedFrom {
	if m != nil {
		return m.Forward
	}
	return nil
}

type PB_MsgResponse_AddNewTextMessage struct {
}

func (m *PB_MsgResponse_AddNewTextMessage) Reset()         { *m = PB_MsgResponse_AddNewTextMessage{} }
func (m *PB_MsgResponse_AddNewTextMessage) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_AddNewTextMessage) ProtoMessage()    {}
func (*PB_MsgResponse_AddNewTextMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{1}
}

// SetRoomActionDoing
type PB_MsgParam_SetRoomActionDoing struct {
}

func (m *PB_MsgParam_SetRoomActionDoing) Reset()                    { *m = PB_MsgParam_SetRoomActionDoing{} }
func (m *PB_MsgParam_SetRoomActionDoing) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_SetRoomActionDoing) ProtoMessage()               {}
func (*PB_MsgParam_SetRoomActionDoing) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type PB_MsgResponse_SetRoomActionDoing struct {
}

func (m *PB_MsgResponse_SetRoomActionDoing) Reset()         { *m = PB_MsgResponse_SetRoomActionDoing{} }
func (m *PB_MsgResponse_SetRoomActionDoing) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_SetRoomActionDoing) ProtoMessage()    {}
func (*PB_MsgResponse_SetRoomActionDoing) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{3}
}

// GetMessagesByIds
type PB_MsgParam_GetMessagesByIds struct {
}

func (m *PB_MsgParam_GetMessagesByIds) Reset()                    { *m = PB_MsgParam_GetMessagesByIds{} }
func (m *PB_MsgParam_GetMessagesByIds) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_GetMessagesByIds) ProtoMessage()               {}
func (*PB_MsgParam_GetMessagesByIds) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type PB_MsgResponse_GetMessagesByIds struct {
}

func (m *PB_MsgResponse_GetMessagesByIds) Reset()                    { *m = PB_MsgResponse_GetMessagesByIds{} }
func (m *PB_MsgResponse_GetMessagesByIds) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgResponse_GetMessagesByIds) ProtoMessage()               {}
func (*PB_MsgResponse_GetMessagesByIds) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

// GetMessagesHistory
type PB_MsgParam_GetMessagesHistory struct {
}

func (m *PB_MsgParam_GetMessagesHistory) Reset()                    { *m = PB_MsgParam_GetMessagesHistory{} }
func (m *PB_MsgParam_GetMessagesHistory) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_GetMessagesHistory) ProtoMessage()               {}
func (*PB_MsgParam_GetMessagesHistory) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type PB_MsgResponse_GetMessagesHistory struct {
}

func (m *PB_MsgResponse_GetMessagesHistory) Reset()         { *m = PB_MsgResponse_GetMessagesHistory{} }
func (m *PB_MsgResponse_GetMessagesHistory) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_GetMessagesHistory) ProtoMessage()    {}
func (*PB_MsgResponse_GetMessagesHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{7}
}

// ///////
type PB_MsgParam_SetMessagesRangeAsSeen struct {
}

func (m *PB_MsgParam_SetMessagesRangeAsSeen) Reset()         { *m = PB_MsgParam_SetMessagesRangeAsSeen{} }
func (m *PB_MsgParam_SetMessagesRangeAsSeen) String() string { return proto.CompactTextString(m) }
func (*PB_MsgParam_SetMessagesRangeAsSeen) ProtoMessage()    {}
func (*PB_MsgParam_SetMessagesRangeAsSeen) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{8}
}

type PB_MsgResponse_SetMessagesRangeAsSeen struct {
}

func (m *PB_MsgResponse_SetMessagesRangeAsSeen) Reset()         { *m = PB_MsgResponse_SetMessagesRangeAsSeen{} }
func (m *PB_MsgResponse_SetMessagesRangeAsSeen) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_SetMessagesRangeAsSeen) ProtoMessage()    {}
func (*PB_MsgResponse_SetMessagesRangeAsSeen) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{9}
}

// ///////////////////////////////////
type PB_MsgParam_DeleteRoomHistory struct {
}

func (m *PB_MsgParam_DeleteRoomHistory) Reset()                    { *m = PB_MsgParam_DeleteRoomHistory{} }
func (m *PB_MsgParam_DeleteRoomHistory) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_DeleteRoomHistory) ProtoMessage()               {}
func (*PB_MsgParam_DeleteRoomHistory) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

type PB_MsgResponse_DeleteRoomHistory struct {
}

func (m *PB_MsgResponse_DeleteRoomHistory) Reset()         { *m = PB_MsgResponse_DeleteRoomHistory{} }
func (m *PB_MsgResponse_DeleteRoomHistory) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_DeleteRoomHistory) ProtoMessage()    {}
func (*PB_MsgResponse_DeleteRoomHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{11}
}

// ///////////////////////////////////
type PB_MsgParam_DeleteMessagesByIds struct {
}

func (m *PB_MsgParam_DeleteMessagesByIds) Reset()         { *m = PB_MsgParam_DeleteMessagesByIds{} }
func (m *PB_MsgParam_DeleteMessagesByIds) String() string { return proto.CompactTextString(m) }
func (*PB_MsgParam_DeleteMessagesByIds) ProtoMessage()    {}
func (*PB_MsgParam_DeleteMessagesByIds) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{12}
}

type PB_MsgResponse_DeleteMessagesByIds struct {
}

func (m *PB_MsgResponse_DeleteMessagesByIds) Reset()         { *m = PB_MsgResponse_DeleteMessagesByIds{} }
func (m *PB_MsgResponse_DeleteMessagesByIds) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_DeleteMessagesByIds) ProtoMessage()    {}
func (*PB_MsgResponse_DeleteMessagesByIds) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{13}
}

// ///////////////////////////////////
type PB_MsgParam_SetMessagesAsReceived struct {
}

func (m *PB_MsgParam_SetMessagesAsReceived) Reset()         { *m = PB_MsgParam_SetMessagesAsReceived{} }
func (m *PB_MsgParam_SetMessagesAsReceived) String() string { return proto.CompactTextString(m) }
func (*PB_MsgParam_SetMessagesAsReceived) ProtoMessage()    {}
func (*PB_MsgParam_SetMessagesAsReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{14}
}

type PB_MsgResponse_SetMessagesAsReceived struct {
}

func (m *PB_MsgResponse_SetMessagesAsReceived) Reset()         { *m = PB_MsgResponse_SetMessagesAsReceived{} }
func (m *PB_MsgResponse_SetMessagesAsReceived) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_SetMessagesAsReceived) ProtoMessage()    {}
func (*PB_MsgResponse_SetMessagesAsReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{15}
}

// ///////////////////////////////////
type PB_MsgParam_ForwardMessages struct {
}

func (m *PB_MsgParam_ForwardMessages) Reset()                    { *m = PB_MsgParam_ForwardMessages{} }
func (m *PB_MsgParam_ForwardMessages) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_ForwardMessages) ProtoMessage()               {}
func (*PB_MsgParam_ForwardMessages) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

type PB_MsgResponse_ForwardMessages struct {
}

func (m *PB_MsgResponse_ForwardMessages) Reset()                    { *m = PB_MsgResponse_ForwardMessages{} }
func (m *PB_MsgResponse_ForwardMessages) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgResponse_ForwardMessages) ProtoMessage()               {}
func (*PB_MsgResponse_ForwardMessages) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

// ///////////////////////////////////
type PB_MsgParam_EditMessage struct {
}

func (m *PB_MsgParam_EditMessage) Reset()                    { *m = PB_MsgParam_EditMessage{} }
func (m *PB_MsgParam_EditMessage) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgParam_EditMessage) ProtoMessage()               {}
func (*PB_MsgParam_EditMessage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

type PB_MsgResponse_EditMessage struct {
}

func (m *PB_MsgResponse_EditMessage) Reset()                    { *m = PB_MsgResponse_EditMessage{} }
func (m *PB_MsgResponse_EditMessage) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgResponse_EditMessage) ProtoMessage()               {}
func (*PB_MsgResponse_EditMessage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

// ///////////////////////////////////
type PB_MsgParam_BroadcastNewMessage struct {
}

func (m *PB_MsgParam_BroadcastNewMessage) Reset()         { *m = PB_MsgParam_BroadcastNewMessage{} }
func (m *PB_MsgParam_BroadcastNewMessage) String() string { return proto.CompactTextString(m) }
func (*PB_MsgParam_BroadcastNewMessage) ProtoMessage()    {}
func (*PB_MsgParam_BroadcastNewMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{20}
}

type PB_MsgResponse_BroadcastNewMessage struct {
}

func (m *PB_MsgResponse_BroadcastNewMessage) Reset()         { *m = PB_MsgResponse_BroadcastNewMessage{} }
func (m *PB_MsgResponse_BroadcastNewMessage) String() string { return proto.CompactTextString(m) }
func (*PB_MsgResponse_BroadcastNewMessage) ProtoMessage()    {}
func (*PB_MsgResponse_BroadcastNewMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{21}
}

func init() {
	proto.RegisterType((*PB_MsgParam_AddNewTextMessage)(nil), "PB_MsgParam_AddNewTextMessage")
	proto.RegisterType((*PB_MsgResponse_AddNewTextMessage)(nil), "PB_MsgResponse_AddNewTextMessage")
	proto.RegisterType((*PB_MsgParam_SetRoomActionDoing)(nil), "PB_MsgParam_SetRoomActionDoing")
	proto.RegisterType((*PB_MsgResponse_SetRoomActionDoing)(nil), "PB_MsgResponse_SetRoomActionDoing")
	proto.RegisterType((*PB_MsgParam_GetMessagesByIds)(nil), "PB_MsgParam_GetMessagesByIds")
	proto.RegisterType((*PB_MsgResponse_GetMessagesByIds)(nil), "PB_MsgResponse_GetMessagesByIds")
	proto.RegisterType((*PB_MsgParam_GetMessagesHistory)(nil), "PB_MsgParam_GetMessagesHistory")
	proto.RegisterType((*PB_MsgResponse_GetMessagesHistory)(nil), "PB_MsgResponse_GetMessagesHistory")
	proto.RegisterType((*PB_MsgParam_SetMessagesRangeAsSeen)(nil), "PB_MsgParam_SetMessagesRangeAsSeen")
	proto.RegisterType((*PB_MsgResponse_SetMessagesRangeAsSeen)(nil), "PB_MsgResponse_SetMessagesRangeAsSeen")
	proto.RegisterType((*PB_MsgParam_DeleteRoomHistory)(nil), "PB_MsgParam_DeleteRoomHistory")
	proto.RegisterType((*PB_MsgResponse_DeleteRoomHistory)(nil), "PB_MsgResponse_DeleteRoomHistory")
	proto.RegisterType((*PB_MsgParam_DeleteMessagesByIds)(nil), "PB_MsgParam_DeleteMessagesByIds")
	proto.RegisterType((*PB_MsgResponse_DeleteMessagesByIds)(nil), "PB_MsgResponse_DeleteMessagesByIds")
	proto.RegisterType((*PB_MsgParam_SetMessagesAsReceived)(nil), "PB_MsgParam_SetMessagesAsReceived")
	proto.RegisterType((*PB_MsgResponse_SetMessagesAsReceived)(nil), "PB_MsgResponse_SetMessagesAsReceived")
	proto.RegisterType((*PB_MsgParam_ForwardMessages)(nil), "PB_MsgParam_ForwardMessages")
	proto.RegisterType((*PB_MsgResponse_ForwardMessages)(nil), "PB_MsgResponse_ForwardMessages")
	proto.RegisterType((*PB_MsgParam_EditMessage)(nil), "PB_MsgParam_EditMessage")
	proto.RegisterType((*PB_MsgResponse_EditMessage)(nil), "PB_MsgResponse_EditMessage")
	proto.RegisterType((*PB_MsgParam_BroadcastNewMessage)(nil), "PB_MsgParam_BroadcastNewMessage")
	proto.RegisterType((*PB_MsgResponse_BroadcastNewMessage)(nil), "PB_MsgResponse_BroadcastNewMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPC_Msg service

type RPC_MsgClient interface {
	AddNewTextMessage(ctx context.Context, in *PB_MsgParam_AddNewTextMessage, opts ...grpc.CallOption) (*PB_MsgResponse_AddNewTextMessage, error)
	SetRoomActionDoing(ctx context.Context, in *PB_MsgParam_SetRoomActionDoing, opts ...grpc.CallOption) (*PB_MsgResponse_SetRoomActionDoing, error)
	GetMessagesByIds(ctx context.Context, in *PB_MsgParam_GetMessagesByIds, opts ...grpc.CallOption) (*PB_MsgResponse_GetMessagesByIds, error)
	GetMessagesHistory(ctx context.Context, in *PB_MsgParam_GetMessagesHistory, opts ...grpc.CallOption) (*PB_MsgResponse_GetMessagesHistory, error)
	SetMessagesRangeAsSeen(ctx context.Context, in *PB_MsgParam_SetMessagesRangeAsSeen, opts ...grpc.CallOption) (*PB_MsgResponse_SetMessagesRangeAsSeen, error)
	DeleteRoomHistory(ctx context.Context, in *PB_MsgParam_DeleteRoomHistory, opts ...grpc.CallOption) (*PB_MsgResponse_DeleteRoomHistory, error)
	DeleteMessagesByIds(ctx context.Context, in *PB_MsgParam_DeleteMessagesByIds, opts ...grpc.CallOption) (*PB_MsgResponse_DeleteMessagesByIds, error)
	SetMessagesAsReceived(ctx context.Context, in *PB_MsgParam_SetMessagesAsReceived, opts ...grpc.CallOption) (*PB_MsgResponse_SetMessagesAsReceived, error)
	ForwardMessages(ctx context.Context, in *PB_MsgParam_ForwardMessages, opts ...grpc.CallOption) (*PB_MsgResponse_ForwardMessages, error)
	EditMessage(ctx context.Context, in *PB_MsgParam_EditMessage, opts ...grpc.CallOption) (*PB_MsgResponse_EditMessage, error)
	BroadcastNewMessage(ctx context.Context, in *PB_MsgParam_BroadcastNewMessage, opts ...grpc.CallOption) (*PB_MsgResponse_BroadcastNewMessage, error)
}

type rPC_MsgClient struct {
	cc *grpc.ClientConn
}

func NewRPC_MsgClient(cc *grpc.ClientConn) RPC_MsgClient {
	return &rPC_MsgClient{cc}
}

func (c *rPC_MsgClient) AddNewTextMessage(ctx context.Context, in *PB_MsgParam_AddNewTextMessage, opts ...grpc.CallOption) (*PB_MsgResponse_AddNewTextMessage, error) {
	out := new(PB_MsgResponse_AddNewTextMessage)
	err := grpc.Invoke(ctx, "/RPC_Msg/AddNewTextMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) SetRoomActionDoing(ctx context.Context, in *PB_MsgParam_SetRoomActionDoing, opts ...grpc.CallOption) (*PB_MsgResponse_SetRoomActionDoing, error) {
	out := new(PB_MsgResponse_SetRoomActionDoing)
	err := grpc.Invoke(ctx, "/RPC_Msg/SetRoomActionDoing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) GetMessagesByIds(ctx context.Context, in *PB_MsgParam_GetMessagesByIds, opts ...grpc.CallOption) (*PB_MsgResponse_GetMessagesByIds, error) {
	out := new(PB_MsgResponse_GetMessagesByIds)
	err := grpc.Invoke(ctx, "/RPC_Msg/GetMessagesByIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) GetMessagesHistory(ctx context.Context, in *PB_MsgParam_GetMessagesHistory, opts ...grpc.CallOption) (*PB_MsgResponse_GetMessagesHistory, error) {
	out := new(PB_MsgResponse_GetMessagesHistory)
	err := grpc.Invoke(ctx, "/RPC_Msg/GetMessagesHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) SetMessagesRangeAsSeen(ctx context.Context, in *PB_MsgParam_SetMessagesRangeAsSeen, opts ...grpc.CallOption) (*PB_MsgResponse_SetMessagesRangeAsSeen, error) {
	out := new(PB_MsgResponse_SetMessagesRangeAsSeen)
	err := grpc.Invoke(ctx, "/RPC_Msg/SetMessagesRangeAsSeen", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) DeleteRoomHistory(ctx context.Context, in *PB_MsgParam_DeleteRoomHistory, opts ...grpc.CallOption) (*PB_MsgResponse_DeleteRoomHistory, error) {
	out := new(PB_MsgResponse_DeleteRoomHistory)
	err := grpc.Invoke(ctx, "/RPC_Msg/DeleteRoomHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) DeleteMessagesByIds(ctx context.Context, in *PB_MsgParam_DeleteMessagesByIds, opts ...grpc.CallOption) (*PB_MsgResponse_DeleteMessagesByIds, error) {
	out := new(PB_MsgResponse_DeleteMessagesByIds)
	err := grpc.Invoke(ctx, "/RPC_Msg/DeleteMessagesByIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) SetMessagesAsReceived(ctx context.Context, in *PB_MsgParam_SetMessagesAsReceived, opts ...grpc.CallOption) (*PB_MsgResponse_SetMessagesAsReceived, error) {
	out := new(PB_MsgResponse_SetMessagesAsReceived)
	err := grpc.Invoke(ctx, "/RPC_Msg/SetMessagesAsReceived", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) ForwardMessages(ctx context.Context, in *PB_MsgParam_ForwardMessages, opts ...grpc.CallOption) (*PB_MsgResponse_ForwardMessages, error) {
	out := new(PB_MsgResponse_ForwardMessages)
	err := grpc.Invoke(ctx, "/RPC_Msg/ForwardMessages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) EditMessage(ctx context.Context, in *PB_MsgParam_EditMessage, opts ...grpc.CallOption) (*PB_MsgResponse_EditMessage, error) {
	out := new(PB_MsgResponse_EditMessage)
	err := grpc.Invoke(ctx, "/RPC_Msg/EditMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPC_MsgClient) BroadcastNewMessage(ctx context.Context, in *PB_MsgParam_BroadcastNewMessage, opts ...grpc.CallOption) (*PB_MsgResponse_BroadcastNewMessage, error) {
	out := new(PB_MsgResponse_BroadcastNewMessage)
	err := grpc.Invoke(ctx, "/RPC_Msg/BroadcastNewMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC_Msg service

type RPC_MsgServer interface {
	AddNewTextMessage(context.Context, *PB_MsgParam_AddNewTextMessage) (*PB_MsgResponse_AddNewTextMessage, error)
	SetRoomActionDoing(context.Context, *PB_MsgParam_SetRoomActionDoing) (*PB_MsgResponse_SetRoomActionDoing, error)
	GetMessagesByIds(context.Context, *PB_MsgParam_GetMessagesByIds) (*PB_MsgResponse_GetMessagesByIds, error)
	GetMessagesHistory(context.Context, *PB_MsgParam_GetMessagesHistory) (*PB_MsgResponse_GetMessagesHistory, error)
	SetMessagesRangeAsSeen(context.Context, *PB_MsgParam_SetMessagesRangeAsSeen) (*PB_MsgResponse_SetMessagesRangeAsSeen, error)
	DeleteRoomHistory(context.Context, *PB_MsgParam_DeleteRoomHistory) (*PB_MsgResponse_DeleteRoomHistory, error)
	DeleteMessagesByIds(context.Context, *PB_MsgParam_DeleteMessagesByIds) (*PB_MsgResponse_DeleteMessagesByIds, error)
	SetMessagesAsReceived(context.Context, *PB_MsgParam_SetMessagesAsReceived) (*PB_MsgResponse_SetMessagesAsReceived, error)
	ForwardMessages(context.Context, *PB_MsgParam_ForwardMessages) (*PB_MsgResponse_ForwardMessages, error)
	EditMessage(context.Context, *PB_MsgParam_EditMessage) (*PB_MsgResponse_EditMessage, error)
	BroadcastNewMessage(context.Context, *PB_MsgParam_BroadcastNewMessage) (*PB_MsgResponse_BroadcastNewMessage, error)
}

func RegisterRPC_MsgServer(s *grpc.Server, srv RPC_MsgServer) {
	s.RegisterService(&_RPC_Msg_serviceDesc, srv)
}

func _RPC_Msg_AddNewTextMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_AddNewTextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).AddNewTextMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/AddNewTextMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).AddNewTextMessage(ctx, req.(*PB_MsgParam_AddNewTextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_SetRoomActionDoing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_SetRoomActionDoing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).SetRoomActionDoing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/SetRoomActionDoing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).SetRoomActionDoing(ctx, req.(*PB_MsgParam_SetRoomActionDoing))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_GetMessagesByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_GetMessagesByIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).GetMessagesByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/GetMessagesByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).GetMessagesByIds(ctx, req.(*PB_MsgParam_GetMessagesByIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_GetMessagesHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_GetMessagesHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).GetMessagesHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/GetMessagesHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).GetMessagesHistory(ctx, req.(*PB_MsgParam_GetMessagesHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_SetMessagesRangeAsSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_SetMessagesRangeAsSeen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).SetMessagesRangeAsSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/SetMessagesRangeAsSeen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).SetMessagesRangeAsSeen(ctx, req.(*PB_MsgParam_SetMessagesRangeAsSeen))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_DeleteRoomHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_DeleteRoomHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).DeleteRoomHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/DeleteRoomHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).DeleteRoomHistory(ctx, req.(*PB_MsgParam_DeleteRoomHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_DeleteMessagesByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_DeleteMessagesByIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).DeleteMessagesByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/DeleteMessagesByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).DeleteMessagesByIds(ctx, req.(*PB_MsgParam_DeleteMessagesByIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_SetMessagesAsReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_SetMessagesAsReceived)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).SetMessagesAsReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/SetMessagesAsReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).SetMessagesAsReceived(ctx, req.(*PB_MsgParam_SetMessagesAsReceived))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_ForwardMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_ForwardMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).ForwardMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/ForwardMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).ForwardMessages(ctx, req.(*PB_MsgParam_ForwardMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_EditMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/EditMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).EditMessage(ctx, req.(*PB_MsgParam_EditMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Msg_BroadcastNewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_MsgParam_BroadcastNewMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MsgServer).BroadcastNewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Msg/BroadcastNewMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MsgServer).BroadcastNewMessage(ctx, req.(*PB_MsgParam_BroadcastNewMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPC_Msg",
	HandlerType: (*RPC_MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewTextMessage",
			Handler:    _RPC_Msg_AddNewTextMessage_Handler,
		},
		{
			MethodName: "SetRoomActionDoing",
			Handler:    _RPC_Msg_SetRoomActionDoing_Handler,
		},
		{
			MethodName: "GetMessagesByIds",
			Handler:    _RPC_Msg_GetMessagesByIds_Handler,
		},
		{
			MethodName: "GetMessagesHistory",
			Handler:    _RPC_Msg_GetMessagesHistory_Handler,
		},
		{
			MethodName: "SetMessagesRangeAsSeen",
			Handler:    _RPC_Msg_SetMessagesRangeAsSeen_Handler,
		},
		{
			MethodName: "DeleteRoomHistory",
			Handler:    _RPC_Msg_DeleteRoomHistory_Handler,
		},
		{
			MethodName: "DeleteMessagesByIds",
			Handler:    _RPC_Msg_DeleteMessagesByIds_Handler,
		},
		{
			MethodName: "SetMessagesAsReceived",
			Handler:    _RPC_Msg_SetMessagesAsReceived_Handler,
		},
		{
			MethodName: "ForwardMessages",
			Handler:    _RPC_Msg_ForwardMessages_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _RPC_Msg_EditMessage_Handler,
		},
		{
			MethodName: "BroadcastNewMessage",
			Handler:    _RPC_Msg_BroadcastNewMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_rpc_messages.proto",
}

func init() { proto.RegisterFile("pb_rpc_messages.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6e, 0x9b, 0x4c,
	0x10, 0xc5, 0x85, 0xf2, 0xc5, 0xfe, 0x3a, 0x51, 0xd3, 0x64, 0xab, 0xa4, 0x94, 0xf8, 0x0f, 0xc6,
	0x49, 0x6a, 0xf5, 0x62, 0x2f, 0xdc, 0x27, 0xb0, 0xeb, 0xa6, 0xf1, 0x45, 0x22, 0x84, 0xa3, 0x4a,
	0xad, 0x2a, 0x21, 0x30, 0x23, 0x17, 0x29, 0xb0, 0x88, 0xa5, 0x75, 0xfc, 0x74, 0x7d, 0x98, 0xbe,
	0x48, 0x85, 0x6d, 0x5c, 0xbc, 0xbb, 0x36, 0xb9, 0x63, 0x99, 0xc3, 0x61, 0xce, 0xec, 0xfe, 0xb4,
	0x70, 0x96, 0xf8, 0x6e, 0x9a, 0x4c, 0xdd, 0x08, 0x39, 0xf7, 0x66, 0xc8, 0x69, 0x92, 0xb2, 0x8c,
	0x19, 0xc7, 0x89, 0xef, 0x62, 0xfc, 0x33, 0x2a, 0xd6, 0x2f, 0x13, 0xdf, 0x9d, 0xfe, 0xf0, 0xb2,
	0xd5, 0xd2, 0xfa, 0xad, 0x41, 0xd3, 0x1e, 0xba, 0x77, 0x7c, 0x66, 0x7b, 0xa9, 0x17, 0xb9, 0x83,
	0x20, 0xb8, 0xc7, 0xf9, 0x03, 0x3e, 0x65, 0x77, 0x2b, 0x1f, 0x42, 0xe0, 0xbf, 0x7c, 0xa9, 0x6b,
	0xa6, 0xd6, 0x7b, 0xe1, 0x2c, 0x9f, 0xc9, 0x39, 0xd4, 0x6c, 0xc4, 0x74, 0x1c, 0xe8, 0x87, 0xa6,
	0xd6, 0x3b, 0x74, 0xd6, 0xab, 0xa5, 0x36, 0x8c, 0x50, 0xaf, 0x2d, 0xdf, 0x2e, 0x9f, 0xc9, 0x7b,
	0x38, 0x71, 0x30, 0x79, 0x5c, 0x3c, 0xb0, 0xb5, 0xe3, 0x38, 0xd0, 0xc1, 0xd4, 0x7a, 0x07, 0x8e,
	0xf4, 0x9e, 0xf4, 0xa1, 0x7e, 0xc3, 0xd2, 0xb9, 0x97, 0x06, 0xfa, 0xb1, 0xa9, 0xf5, 0x8e, 0xfa,
	0x3a, 0xcd, 0x9b, 0x5b, 0xd5, 0xd7, 0x15, 0x0c, 0x6e, 0x52, 0x16, 0x39, 0x85, 0xd0, 0xb2, 0xc0,
	0x5c, 0x05, 0x70, 0x90, 0x27, 0x2c, 0xe6, 0x28, 0x67, 0xb0, 0x4c, 0x68, 0x95, 0x43, 0x4e, 0x30,
	0x73, 0x18, 0x8b, 0x06, 0xd3, 0x2c, 0x64, 0xf1, 0x88, 0x85, 0xf1, 0xcc, 0xea, 0x42, 0x47, 0x70,
	0x51, 0x88, 0x5a, 0xd0, 0x28, 0xdb, 0x7c, 0xc6, 0xe2, 0x07, 0x7c, 0xb8, 0x18, 0x07, 0xdc, 0xea,
	0x40, 0x5b, 0x30, 0x91, 0x24, 0x42, 0x27, 0xa5, 0xfa, 0x6d, 0xc8, 0x33, 0x96, 0x2e, 0x14, 0x9d,
	0x28, 0x44, 0x97, 0x60, 0x09, 0x81, 0x0a, 0x85, 0xe3, 0xc5, 0x33, 0x1c, 0xf0, 0x09, 0x62, 0x6c,
	0xbd, 0x83, 0x2b, 0x39, 0x94, 0x4a, 0xd8, 0xde, 0x3e, 0x04, 0x23, 0x7c, 0xc4, 0x0c, 0xf3, 0xf4,
	0xc5, 0xff, 0xe4, 0x21, 0xcb, 0x9a, 0x4d, 0xfa, 0xb2, 0xc9, 0x76, 0xfa, 0x4d, 0xdb, 0x82, 0xcd,
	0xb6, 0x6a, 0x33, 0x01, 0x29, 0xdc, 0x80, 0x3b, 0x38, 0xc5, 0xf0, 0x17, 0x06, 0xd6, 0x35, 0x5c,
	0xee, 0xce, 0x56, 0xd2, 0x35, 0xe1, 0xa2, 0x6c, 0xb6, 0x3e, 0x35, 0x85, 0xf0, 0xdf, 0x7e, 0x6c,
	0x6c, 0x44, 0xc5, 0x5b, 0x78, 0x53, 0x36, 0xf8, 0x14, 0x84, 0x9b, 0x63, 0xd5, 0x00, 0x43, 0xf8,
	0xb8, 0x5c, 0x15, 0xe6, 0x31, 0x4c, 0x99, 0x17, 0x4c, 0x3d, 0x9e, 0xdd, 0xe3, 0xbc, 0x90, 0xc8,
	0xf3, 0x50, 0xa8, 0xfa, 0x7f, 0xea, 0x50, 0x77, 0xec, 0x8f, 0xb9, 0x8e, 0x7c, 0x81, 0x53, 0x19,
	0xd1, 0x16, 0xdd, 0x8b, 0xb0, 0xd1, 0xa1, 0x55, 0x84, 0x90, 0xaf, 0x40, 0xe4, 0x03, 0x4f, 0xda,
	0x74, 0x3f, 0x36, 0x86, 0x45, 0x2b, 0xa9, 0x21, 0x13, 0x38, 0x11, 0x31, 0x20, 0x4d, 0xba, 0x0f,
	0x24, 0xc3, 0xa4, 0x15, 0x1c, 0xe5, 0xfd, 0xca, 0x58, 0x08, 0xfd, 0xca, 0x02, 0xb9, 0x5f, 0x85,
	0x09, 0xc2, 0xb9, 0x1a, 0x13, 0xd2, 0xa5, 0xd5, 0xd0, 0x19, 0xd7, 0xf4, 0x59, 0xcc, 0xe5, 0x3b,
	0x29, 0x31, 0x24, 0xec, 0xa4, 0x54, 0x97, 0x77, 0x52, 0xb6, 0xf8, 0x0e, 0xaf, 0x15, 0x50, 0x11,
	0x93, 0x56, 0xc0, 0x69, 0x74, 0x69, 0x35, 0x9b, 0xc4, 0x87, 0x33, 0x25, 0x67, 0xc4, 0xa2, 0x95,
	0xcc, 0x1a, 0x57, 0xf4, 0x39, 0xc8, 0x12, 0x1b, 0x5e, 0x09, 0x10, 0x92, 0x06, 0xdd, 0x03, 0xb1,
	0xd1, 0xa6, 0xfb, 0x19, 0x26, 0x23, 0x38, 0x2a, 0x91, 0x49, 0x74, 0xba, 0x83, 0x68, 0xe3, 0x82,
	0xee, 0x06, 0x3a, 0x9f, 0xac, 0x02, 0x4f, 0x61, 0xb2, 0x0a, 0x85, 0x3c, 0x59, 0x85, 0x68, 0x78,
	0x0a, 0xff, 0x87, 0x29, 0xcd, 0xef, 0x69, 0xff, 0xf6, 0xc0, 0xd6, 0xbe, 0x69, 0x4f, 0x7e, 0x6d,
	0x79, 0x47, 0x7f, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x08, 0x87, 0xb9, 0x69, 0xdb, 0x07, 0x00,
	0x00,
}
