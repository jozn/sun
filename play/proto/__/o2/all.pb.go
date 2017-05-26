// Code generated by protoc-gen-go.
// source: all.proto
// DO NOT EDIT!

/*
Package all is a generated protocol buffer package.

It is generated from these files:
	all.proto

It has these top-level messages:
	PBUserWithMe
	PBMessage
*/
package all

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PBUserWithMe struct {
	UserId         int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	UserName       string `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	FirstName      string `protobuf:"bytes,3,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName       string `protobuf:"bytes,4,opt,name=LastName" json:"LastName,omitempty"`
	About          string `protobuf:"bytes,5,opt,name=About" json:"About,omitempty"`
	FullName       string `protobuf:"bytes,6,opt,name=FullName" json:"FullName,omitempty"`
	AvatarUrl      string `protobuf:"bytes,7,opt,name=AvatarUrl" json:"AvatarUrl,omitempty"`
	PrivacyProfile int32  `protobuf:"varint,8,opt,name=PrivacyProfile" json:"PrivacyProfile,omitempty"`
	Phone          string `protobuf:"bytes,9,opt,name=Phone" json:"Phone,omitempty"`
	UpdatedTime    int32  `protobuf:"varint,10,opt,name=UpdatedTime" json:"UpdatedTime,omitempty"`
	AppVersion     int32  `protobuf:"varint,11,opt,name=AppVersion" json:"AppVersion,omitempty"`
}

func (m *PBUserWithMe) Reset()                    { *m = PBUserWithMe{} }
func (m *PBUserWithMe) String() string            { return proto.CompactTextString(m) }
func (*PBUserWithMe) ProtoMessage()               {}
func (*PBUserWithMe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PBUserWithMe) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PBUserWithMe) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PBUserWithMe) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *PBUserWithMe) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *PBUserWithMe) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PBUserWithMe) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *PBUserWithMe) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *PBUserWithMe) GetPrivacyProfile() int32 {
	if m != nil {
		return m.PrivacyProfile
	}
	return 0
}

func (m *PBUserWithMe) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PBUserWithMe) GetUpdatedTime() int32 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *PBUserWithMe) GetAppVersion() int32 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

type PBMessage struct {
	MessageKey         string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey            string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId             int32  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	RoomTypeId         int32  `protobuf:"varint,4,opt,name=RoomTypeId" json:"RoomTypeId,omitempty"`
	MessageTypeId      int32  `protobuf:"varint,5,opt,name=MessageTypeId" json:"MessageTypeId,omitempty"`
	Text               string `protobuf:"bytes,6,opt,name=Text" json:"Text,omitempty"`
	ExtraJson          string `protobuf:"bytes,7,opt,name=ExtraJson" json:"ExtraJson,omitempty"`
	IsByMe             int32  `protobuf:"varint,8,opt,name=IsByMe" json:"IsByMe,omitempty"`
	IsStared           int32  `protobuf:"varint,9,opt,name=IsStared" json:"IsStared,omitempty"`
	CreatedMs          int64  `protobuf:"varint,10,opt,name=CreatedMs" json:"CreatedMs,omitempty"`
	CreatedDeviceMs    int64  `protobuf:"varint,11,opt,name=CreatedDeviceMs" json:"CreatedDeviceMs,omitempty"`
	SortId             int64  `protobuf:"varint,12,opt,name=SortId" json:"SortId,omitempty"`
	PeerSeenTime       int64  `protobuf:"varint,13,opt,name=PeerSeenTime" json:"PeerSeenTime,omitempty"`
	ServerReceivedTime int64  `protobuf:"varint,14,opt,name=ServerReceivedTime" json:"ServerReceivedTime,omitempty"`
	ServerDeletedTime  int64  `protobuf:"varint,15,opt,name=ServerDeletedTime" json:"ServerDeletedTime,omitempty"`
	ISeenTime          int64  `protobuf:"varint,16,opt,name=ISeenTime" json:"ISeenTime,omitempty"`
	ToPush             int32  `protobuf:"varint,17,opt,name=ToPush" json:"ToPush,omitempty"`
	MsgFile_LocalSrc   string `protobuf:"bytes,18,opt,name=MsgFile_LocalSrc,json=MsgFileLocalSrc" json:"MsgFile_LocalSrc,omitempty"`
	MsgFile_Status     int32  `protobuf:"varint,19,opt,name=MsgFile_Status,json=MsgFileStatus" json:"MsgFile_Status,omitempty"`
}

func (m *PBMessage) Reset()                    { *m = PBMessage{} }
func (m *PBMessage) String() string            { return proto.CompactTextString(m) }
func (*PBMessage) ProtoMessage()               {}
func (*PBMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PBMessage) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PBMessage) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PBMessage) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PBMessage) GetRoomTypeId() int32 {
	if m != nil {
		return m.RoomTypeId
	}
	return 0
}

func (m *PBMessage) GetMessageTypeId() int32 {
	if m != nil {
		return m.MessageTypeId
	}
	return 0
}

func (m *PBMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PBMessage) GetExtraJson() string {
	if m != nil {
		return m.ExtraJson
	}
	return ""
}

func (m *PBMessage) GetIsByMe() int32 {
	if m != nil {
		return m.IsByMe
	}
	return 0
}

func (m *PBMessage) GetIsStared() int32 {
	if m != nil {
		return m.IsStared
	}
	return 0
}

func (m *PBMessage) GetCreatedMs() int64 {
	if m != nil {
		return m.CreatedMs
	}
	return 0
}

func (m *PBMessage) GetCreatedDeviceMs() int64 {
	if m != nil {
		return m.CreatedDeviceMs
	}
	return 0
}

func (m *PBMessage) GetSortId() int64 {
	if m != nil {
		return m.SortId
	}
	return 0
}

func (m *PBMessage) GetPeerSeenTime() int64 {
	if m != nil {
		return m.PeerSeenTime
	}
	return 0
}

func (m *PBMessage) GetServerReceivedTime() int64 {
	if m != nil {
		return m.ServerReceivedTime
	}
	return 0
}

func (m *PBMessage) GetServerDeletedTime() int64 {
	if m != nil {
		return m.ServerDeletedTime
	}
	return 0
}

func (m *PBMessage) GetISeenTime() int64 {
	if m != nil {
		return m.ISeenTime
	}
	return 0
}

func (m *PBMessage) GetToPush() int32 {
	if m != nil {
		return m.ToPush
	}
	return 0
}

func (m *PBMessage) GetMsgFile_LocalSrc() string {
	if m != nil {
		return m.MsgFile_LocalSrc
	}
	return ""
}

func (m *PBMessage) GetMsgFile_Status() int32 {
	if m != nil {
		return m.MsgFile_Status
	}
	return 0
}

func init() {
	proto.RegisterType((*PBUserWithMe)(nil), "PBUserWithMe")
	proto.RegisterType((*PBMessage)(nil), "PBMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *PBMessage, opts ...grpc.CallOption) (*PBUserWithMe, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *PBMessage, opts ...grpc.CallOption) (*PBUserWithMe, error) {
	out := new(PBUserWithMe)
	err := grpc.Invoke(ctx, "/Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *PBMessage) (*PBUserWithMe, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PBMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*PBMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "all.proto",
}

func init() { proto.RegisterFile("all.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x94, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x69, 0x53, 0xb7, 0xf5, 0xf4, 0xbc, 0x20, 0xb4, 0xaa, 0x50, 0x55, 0x45, 0x1c, 0x8a,
	0x84, 0x7c, 0x51, 0x9e, 0xa0, 0xa6, 0x14, 0x0c, 0x35, 0xb2, 0xec, 0x14, 0x2e, 0xd1, 0xd6, 0x1e,
	0x1a, 0x4b, 0xeb, 0xac, 0xb5, 0xbb, 0x89, 0x9a, 0x77, 0xe4, 0x6d, 0x78, 0x01, 0xb4, 0x07, 0xbb,
	0x4e, 0xe1, 0x6e, 0xff, 0x6f, 0xfe, 0xf1, 0x78, 0xff, 0x71, 0x02, 0x21, 0xe3, 0x3c, 0x6a, 0xa5,
	0xd0, 0x62, 0xfc, 0x7b, 0x1d, 0x76, 0xb3, 0xf8, 0x46, 0xa1, 0xfc, 0x51, 0xeb, 0x69, 0x8a, 0xe4,
	0x39, 0x6c, 0x1a, 0x95, 0x54, 0x74, 0xed, 0x74, 0xed, 0x2c, 0xc8, 0xbd, 0x22, 0xc7, 0xb0, 0x6d,
	0x4e, 0xdf, 0x58, 0x83, 0x74, 0xfd, 0x74, 0xed, 0x2c, 0xcc, 0x7b, 0x4d, 0x5e, 0x40, 0x78, 0x55,
	0x4b, 0xa5, 0x6d, 0x71, 0x64, 0x8b, 0x0f, 0xc0, 0x74, 0x5e, 0x33, 0x5f, 0xdc, 0x70, 0x9d, 0x9d,
	0x26, 0xcf, 0x20, 0xb8, 0xb8, 0x15, 0x73, 0x4d, 0x03, 0x5b, 0x70, 0xc2, 0x74, 0x5c, 0xcd, 0x39,
	0xb7, 0x1d, 0x9b, 0xae, 0xa3, 0xd3, 0x66, 0xd6, 0xc5, 0x82, 0x69, 0x26, 0x6f, 0x24, 0xa7, 0x5b,
	0x6e, 0x56, 0x0f, 0xc8, 0x6b, 0xd8, 0xcf, 0x64, 0xbd, 0x60, 0xe5, 0x32, 0x93, 0xe2, 0x57, 0xcd,
	0x91, 0x6e, 0xdb, 0x5b, 0x3c, 0xa2, 0x66, 0x6e, 0x36, 0x15, 0x33, 0xa4, 0xa1, 0x9b, 0x6b, 0x05,
	0x39, 0x85, 0x9d, 0x9b, 0xb6, 0x62, 0x1a, 0xab, 0x49, 0xdd, 0x20, 0x05, 0xdb, 0x3a, 0x44, 0xe4,
	0x04, 0xe0, 0xa2, 0x6d, 0xbf, 0xa3, 0x54, 0xb5, 0x98, 0xd1, 0x1d, 0x6b, 0x18, 0x90, 0xf1, 0x9f,
	0x0d, 0x08, 0xb3, 0x38, 0x45, 0xa5, 0xd8, 0x9d, 0x75, 0xfb, 0xe3, 0x57, 0x5c, 0xda, 0x3c, 0xc3,
	0x7c, 0x40, 0x08, 0x85, 0xad, 0x5c, 0x88, 0xc6, 0x14, 0x5d, 0xa4, 0x9d, 0x1c, 0x6c, 0x61, 0xb4,
	0xb2, 0x85, 0x13, 0x00, 0x63, 0x99, 0x2c, 0x5b, 0x4c, 0x2a, 0x9b, 0x66, 0x90, 0x0f, 0x08, 0x79,
	0x09, 0x7b, 0xfe, 0xf9, 0xde, 0x12, 0x58, 0xcb, 0x2a, 0x24, 0x04, 0x36, 0x26, 0x78, 0xaf, 0x7d,
	0xb6, 0xf6, 0x6c, 0x72, 0xfd, 0x78, 0xaf, 0x25, 0xfb, 0xa2, 0xc4, 0xac, 0xcb, 0xb5, 0x07, 0xe6,
	0x7d, 0x12, 0x15, 0x2f, 0xd3, 0x2e, 0x4f, 0xaf, 0xcc, 0xa6, 0x12, 0x55, 0x68, 0x26, 0xb1, 0xb2,
	0x51, 0x06, 0x79, 0xaf, 0xcd, 0x13, 0x3f, 0x48, 0x34, 0xd1, 0xa5, 0xca, 0x66, 0x39, 0xca, 0x1f,
	0x00, 0x39, 0x83, 0x03, 0x2f, 0x2e, 0x71, 0x51, 0x97, 0x98, 0x2a, 0x1b, 0xe7, 0x28, 0x7f, 0x8c,
	0xcd, 0xec, 0x42, 0x48, 0x9d, 0x54, 0x74, 0xd7, 0x1a, 0xbc, 0x22, 0x63, 0xd8, 0xcd, 0x10, 0x65,
	0x81, 0x38, 0xb3, 0xeb, 0xda, 0xb3, 0xd5, 0x15, 0x46, 0x22, 0x20, 0x05, 0xca, 0x05, 0xca, 0x1c,
	0x4b, 0xac, 0x17, 0x7e, 0xb1, 0xfb, 0xd6, 0xf9, 0x9f, 0x0a, 0x79, 0x07, 0x47, 0x8e, 0x5e, 0x22,
	0xc7, 0xee, 0x3b, 0x38, 0xb0, 0xf6, 0x7f, 0x0b, 0xe6, 0x86, 0x49, 0x3f, 0xfe, 0xd0, 0xdd, 0xb0,
	0x07, 0xe6, 0xbd, 0x27, 0x22, 0x9b, 0xab, 0x29, 0x3d, 0x72, 0x99, 0x39, 0x45, 0xde, 0xc2, 0x61,
	0xaa, 0xee, 0xae, 0x6a, 0x8e, 0x3f, 0xaf, 0x45, 0xc9, 0x78, 0x21, 0x4b, 0x4a, 0x6c, 0xe0, 0x07,
	0x9e, 0x77, 0x98, 0xbc, 0x82, 0xfd, 0xce, 0x5a, 0x68, 0xa6, 0xe7, 0x8a, 0x3e, 0xf5, 0xfb, 0x74,
	0xd4, 0xc1, 0xf3, 0x73, 0xd8, 0xfa, 0x24, 0x11, 0x35, 0x4a, 0xf2, 0x06, 0xb6, 0x0b, 0xb6, 0xfc,
	0x8c, 0x9c, 0x0b, 0x02, 0x51, 0xff, 0x29, 0x1e, 0xef, 0x45, 0xc3, 0x5f, 0xf9, 0xf8, 0x49, 0x7c,
	0x02, 0xb4, 0x14, 0x4d, 0xd4, 0x30, 0x59, 0x89, 0x46, 0x31, 0xc9, 0x22, 0x25, 0xca, 0x9a, 0xf1,
	0xa8, 0xbd, 0x8d, 0xd7, 0xb3, 0xf8, 0x76, 0xd3, 0xfe, 0x3f, 0xbc, 0xff, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x36, 0xfa, 0x8a, 0x2c, 0x04, 0x00, 0x00,
}