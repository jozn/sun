// Code generated by protoc-gen-go.
// source: global.proto
// DO NOT EDIT!

/*
Package x is a generated protocol buffer package.

It is generated from these files:
	global.proto
	pb_msg_removed_server.proto
	pb_seen_msgs.proto

It has these top-level messages:
	PB_CommandToServer
	PB_CommandToClient
	PB_UserWithMe
	PB_Message
	PB_Response
	PB_Request
	PB_RequestMsgAddMany
	PB_ResponseMsgAddMany
	PB_Push
	PB_Result
	PB_PushMsgAddMany
	PB_ResultMsgAddMany
	PB_MsgEvent
	PB_PushMsgEvents
	PB_ResultMsgEvents
	PB_MsgRemovedFromServer
	PB_PushMsgsRemovedFromServer
	PB_ResultMsgsRemovedFromServer
	PB_MsgsSeenFromClient
	PB_MsgsSeenToClient
	PB_MsgSeen
	PB_RequestMsgsSeen
	PB_ResponseMsgsSeen
	PB_PushMsgsReceivedToPeer
	PB_ResultMsgsReceivedToPeer
*/
package x

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

type PB_CommandToServer struct {
	ClientCallId int64  `protobuf:"varint,1,opt,name=ClientCallId" json:"ClientCallId,omitempty"`
	Command      string `protobuf:"bytes,2,opt,name=Command" json:"Command,omitempty"`
	Data         []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *PB_CommandToServer) Reset()                    { *m = PB_CommandToServer{} }
func (m *PB_CommandToServer) String() string            { return proto.CompactTextString(m) }
func (*PB_CommandToServer) ProtoMessage()               {}
func (*PB_CommandToServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PB_CommandToServer) GetClientCallId() int64 {
	if m != nil {
		return m.ClientCallId
	}
	return 0
}

func (m *PB_CommandToServer) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *PB_CommandToServer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PB_CommandToClient struct {
	ServerCallId int64  `protobuf:"varint,1,opt,name=ServerCallId" json:"ServerCallId,omitempty"`
	Command      string `protobuf:"bytes,2,opt,name=Command" json:"Command,omitempty"`
	Data         []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *PB_CommandToClient) Reset()                    { *m = PB_CommandToClient{} }
func (m *PB_CommandToClient) String() string            { return proto.CompactTextString(m) }
func (*PB_CommandToClient) ProtoMessage()               {}
func (*PB_CommandToClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PB_CommandToClient) GetServerCallId() int64 {
	if m != nil {
		return m.ServerCallId
	}
	return 0
}

func (m *PB_CommandToClient) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *PB_CommandToClient) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PB_UserWithMe struct {
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
	FollowingType  int32  `protobuf:"varint,14,opt,name=FollowingType" json:"FollowingType,omitempty"`
}

func (m *PB_UserWithMe) Reset()                    { *m = PB_UserWithMe{} }
func (m *PB_UserWithMe) String() string            { return proto.CompactTextString(m) }
func (*PB_UserWithMe) ProtoMessage()               {}
func (*PB_UserWithMe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PB_UserWithMe) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_UserWithMe) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PB_UserWithMe) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *PB_UserWithMe) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *PB_UserWithMe) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PB_UserWithMe) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *PB_UserWithMe) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *PB_UserWithMe) GetPrivacyProfile() int32 {
	if m != nil {
		return m.PrivacyProfile
	}
	return 0
}

func (m *PB_UserWithMe) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PB_UserWithMe) GetUpdatedTime() int32 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *PB_UserWithMe) GetAppVersion() int32 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

func (m *PB_UserWithMe) GetFollowingType() int32 {
	if m != nil {
		return m.FollowingType
	}
	return 0
}

type PB_Message struct {
	MessageKey         string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey            string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId             int64  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	PeerId             int64  `protobuf:"varint,20,opt,name=PeerId" json:"PeerId,omitempty"`
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

func (m *PB_Message) Reset()                    { *m = PB_Message{} }
func (m *PB_Message) String() string            { return proto.CompactTextString(m) }
func (*PB_Message) ProtoMessage()               {}
func (*PB_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PB_Message) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PB_Message) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_Message) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_Message) GetPeerId() int64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *PB_Message) GetRoomTypeId() int32 {
	if m != nil {
		return m.RoomTypeId
	}
	return 0
}

func (m *PB_Message) GetMessageTypeId() int32 {
	if m != nil {
		return m.MessageTypeId
	}
	return 0
}

func (m *PB_Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PB_Message) GetExtraJson() string {
	if m != nil {
		return m.ExtraJson
	}
	return ""
}

func (m *PB_Message) GetIsByMe() int32 {
	if m != nil {
		return m.IsByMe
	}
	return 0
}

func (m *PB_Message) GetIsStared() int32 {
	if m != nil {
		return m.IsStared
	}
	return 0
}

func (m *PB_Message) GetCreatedMs() int64 {
	if m != nil {
		return m.CreatedMs
	}
	return 0
}

func (m *PB_Message) GetCreatedDeviceMs() int64 {
	if m != nil {
		return m.CreatedDeviceMs
	}
	return 0
}

func (m *PB_Message) GetSortId() int64 {
	if m != nil {
		return m.SortId
	}
	return 0
}

func (m *PB_Message) GetPeerSeenTime() int64 {
	if m != nil {
		return m.PeerSeenTime
	}
	return 0
}

func (m *PB_Message) GetServerReceivedTime() int64 {
	if m != nil {
		return m.ServerReceivedTime
	}
	return 0
}

func (m *PB_Message) GetServerDeletedTime() int64 {
	if m != nil {
		return m.ServerDeletedTime
	}
	return 0
}

func (m *PB_Message) GetISeenTime() int64 {
	if m != nil {
		return m.ISeenTime
	}
	return 0
}

func (m *PB_Message) GetToPush() int32 {
	if m != nil {
		return m.ToPush
	}
	return 0
}

func (m *PB_Message) GetMsgFile_LocalSrc() string {
	if m != nil {
		return m.MsgFile_LocalSrc
	}
	return ""
}

func (m *PB_Message) GetMsgFile_Status() int32 {
	if m != nil {
		return m.MsgFile_Status
	}
	return 0
}

type PB_Response struct {
	Status int32 `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
}

func (m *PB_Response) Reset()                    { *m = PB_Response{} }
func (m *PB_Response) String() string            { return proto.CompactTextString(m) }
func (*PB_Response) ProtoMessage()               {}
func (*PB_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PB_Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type PB_Request struct {
	Status int32 `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
}

func (m *PB_Request) Reset()                    { *m = PB_Request{} }
func (m *PB_Request) String() string            { return proto.CompactTextString(m) }
func (*PB_Request) ProtoMessage()               {}
func (*PB_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PB_Request) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

// //
type PB_RequestMsgAddMany struct {
	Request  *PB_Request   `protobuf:"bytes,1,opt,name=Request" json:"Request,omitempty"`
	Messages []*PB_Message `protobuf:"bytes,2,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *PB_RequestMsgAddMany) Reset()                    { *m = PB_RequestMsgAddMany{} }
func (m *PB_RequestMsgAddMany) String() string            { return proto.CompactTextString(m) }
func (*PB_RequestMsgAddMany) ProtoMessage()               {}
func (*PB_RequestMsgAddMany) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PB_RequestMsgAddMany) GetRequest() *PB_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PB_RequestMsgAddMany) GetMessages() []*PB_Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type PB_ResponseMsgAddMany struct {
	Response *PB_Response `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PB_ResponseMsgAddMany) Reset()                    { *m = PB_ResponseMsgAddMany{} }
func (m *PB_ResponseMsgAddMany) String() string            { return proto.CompactTextString(m) }
func (*PB_ResponseMsgAddMany) ProtoMessage()               {}
func (*PB_ResponseMsgAddMany) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PB_ResponseMsgAddMany) GetResponse() *PB_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

// ////////// Pushes ////////////
type PB_Push struct {
	Status int32 `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
}

func (m *PB_Push) Reset()                    { *m = PB_Push{} }
func (m *PB_Push) String() string            { return proto.CompactTextString(m) }
func (*PB_Push) ProtoMessage()               {}
func (*PB_Push) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PB_Push) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type PB_Result struct {
	Status int32 `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
}

func (m *PB_Result) Reset()                    { *m = PB_Result{} }
func (m *PB_Result) String() string            { return proto.CompactTextString(m) }
func (*PB_Result) ProtoMessage()               {}
func (*PB_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PB_Result) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

// /
type PB_PushMsgAddMany struct {
	Push     *PB_Push         `protobuf:"bytes,1,opt,name=Push" json:"Push,omitempty"`
	Messages []*PB_Message    `protobuf:"bytes,2,rep,name=Messages" json:"Messages,omitempty"`
	Users    []*PB_UserWithMe `protobuf:"bytes,3,rep,name=Users" json:"Users,omitempty"`
}

func (m *PB_PushMsgAddMany) Reset()                    { *m = PB_PushMsgAddMany{} }
func (m *PB_PushMsgAddMany) String() string            { return proto.CompactTextString(m) }
func (*PB_PushMsgAddMany) ProtoMessage()               {}
func (*PB_PushMsgAddMany) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PB_PushMsgAddMany) GetPush() *PB_Push {
	if m != nil {
		return m.Push
	}
	return nil
}

func (m *PB_PushMsgAddMany) GetMessages() []*PB_Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *PB_PushMsgAddMany) GetUsers() []*PB_UserWithMe {
	if m != nil {
		return m.Users
	}
	return nil
}

type PB_ResultMsgAddMany struct {
	Result *PB_Result `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *PB_ResultMsgAddMany) Reset()                    { *m = PB_ResultMsgAddMany{} }
func (m *PB_ResultMsgAddMany) String() string            { return proto.CompactTextString(m) }
func (*PB_ResultMsgAddMany) ProtoMessage()               {}
func (*PB_ResultMsgAddMany) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PB_ResultMsgAddMany) GetResult() *PB_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type PB_MsgEvent struct {
	MessageKey string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey    string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	PeerUserId int64  `protobuf:"varint,3,opt,name=PeerUserId" json:"PeerUserId,omitempty"`
	EventType  int32  `protobuf:"varint,4,opt,name=EventType" json:"EventType,omitempty"`
	AtTime     int64  `protobuf:"varint,10,opt,name=AtTime" json:"AtTime,omitempty"`
}

func (m *PB_MsgEvent) Reset()                    { *m = PB_MsgEvent{} }
func (m *PB_MsgEvent) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgEvent) ProtoMessage()               {}
func (*PB_MsgEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PB_MsgEvent) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PB_MsgEvent) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_MsgEvent) GetPeerUserId() int64 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *PB_MsgEvent) GetEventType() int32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

func (m *PB_MsgEvent) GetAtTime() int64 {
	if m != nil {
		return m.AtTime
	}
	return 0
}

type PB_PushMsgEvents struct {
	Push   *PB_Push       `protobuf:"bytes,1,opt,name=Push" json:"Push,omitempty"`
	Events []*PB_MsgEvent `protobuf:"bytes,2,rep,name=Events" json:"Events,omitempty"`
}

func (m *PB_PushMsgEvents) Reset()                    { *m = PB_PushMsgEvents{} }
func (m *PB_PushMsgEvents) String() string            { return proto.CompactTextString(m) }
func (*PB_PushMsgEvents) ProtoMessage()               {}
func (*PB_PushMsgEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PB_PushMsgEvents) GetPush() *PB_Push {
	if m != nil {
		return m.Push
	}
	return nil
}

func (m *PB_PushMsgEvents) GetEvents() []*PB_MsgEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type PB_ResultMsgEvents struct {
	Result *PB_Result `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *PB_ResultMsgEvents) Reset()                    { *m = PB_ResultMsgEvents{} }
func (m *PB_ResultMsgEvents) String() string            { return proto.CompactTextString(m) }
func (*PB_ResultMsgEvents) ProtoMessage()               {}
func (*PB_ResultMsgEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *PB_ResultMsgEvents) GetResult() *PB_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_CommandToServer)(nil), "PB_CommandToServer")
	proto.RegisterType((*PB_CommandToClient)(nil), "PB_CommandToClient")
	proto.RegisterType((*PB_UserWithMe)(nil), "PB_UserWithMe")
	proto.RegisterType((*PB_Message)(nil), "PB_Message")
	proto.RegisterType((*PB_Response)(nil), "PB_Response")
	proto.RegisterType((*PB_Request)(nil), "PB_Request")
	proto.RegisterType((*PB_RequestMsgAddMany)(nil), "PB_RequestMsgAddMany")
	proto.RegisterType((*PB_ResponseMsgAddMany)(nil), "PB_ResponseMsgAddMany")
	proto.RegisterType((*PB_Push)(nil), "PB_Push")
	proto.RegisterType((*PB_Result)(nil), "PB_Result")
	proto.RegisterType((*PB_PushMsgAddMany)(nil), "PB_PushMsgAddMany")
	proto.RegisterType((*PB_ResultMsgAddMany)(nil), "PB_ResultMsgAddMany")
	proto.RegisterType((*PB_MsgEvent)(nil), "PB_MsgEvent")
	proto.RegisterType((*PB_PushMsgEvents)(nil), "PB_PushMsgEvents")
	proto.RegisterType((*PB_ResultMsgEvents)(nil), "PB_ResultMsgEvents")
}

func init() { proto.RegisterFile("global.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xd1, 0x6e, 0x23, 0x35,
	0x14, 0x55, 0x3a, 0x4d, 0x9a, 0xdc, 0xa4, 0x69, 0xeb, 0xed, 0xae, 0x2c, 0xb4, 0x5a, 0x85, 0xa1,
	0x85, 0x20, 0xa1, 0x79, 0x80, 0x17, 0x78, 0xcc, 0xb4, 0x5b, 0x11, 0xd8, 0x41, 0x23, 0x27, 0x5d,
	0x24, 0x5e, 0x22, 0x37, 0xe3, 0xa6, 0x23, 0x4d, 0xc6, 0x61, 0xec, 0x84, 0xe6, 0x8d, 0x0f, 0xe0,
	0x17, 0xf8, 0x4a, 0x7e, 0x00, 0xf9, 0xda, 0x33, 0x99, 0x14, 0x2a, 0x96, 0x7d, 0x9b, 0x7b, 0xee,
	0xb9, 0x3e, 0xd7, 0xf7, 0xd8, 0x4e, 0xa0, 0xb7, 0xc8, 0xe4, 0x1d, 0xcf, 0x82, 0x55, 0x21, 0xb5,
	0xf4, 0xef, 0x81, 0xc4, 0xe1, 0xec, 0x4a, 0x2e, 0x97, 0x3c, 0x4f, 0xa6, 0x72, 0x22, 0x8a, 0x8d,
	0x28, 0x88, 0x0f, 0xbd, 0xab, 0x2c, 0x15, 0xb9, 0xbe, 0xe2, 0x59, 0x36, 0x4e, 0x68, 0x63, 0xd0,
	0x18, 0x7a, 0x6c, 0x0f, 0x23, 0x14, 0x8e, 0x5c, 0x19, 0x3d, 0x18, 0x34, 0x86, 0x1d, 0x56, 0x86,
	0x84, 0xc0, 0xe1, 0x35, 0xd7, 0x9c, 0x7a, 0x83, 0xc6, 0xb0, 0xc7, 0xf0, 0xfb, 0xa9, 0x8e, 0x5d,
	0xc9, 0xe8, 0x58, 0xc5, 0x7d, 0x9d, 0x3a, 0xf6, 0x3f, 0x75, 0xfe, 0x3a, 0x80, 0xe3, 0x38, 0x9c,
	0xdd, 0x2a, 0x51, 0xfc, 0x9c, 0xea, 0x87, 0x48, 0x90, 0x57, 0xd0, 0x32, 0x91, 0x5b, 0xbd, 0xc9,
	0x5c, 0x44, 0x3e, 0x81, 0xb6, 0xf9, 0xfa, 0x89, 0x2f, 0x85, 0x5b, 0xb8, 0x8a, 0xc9, 0x6b, 0xe8,
	0xdc, 0xa4, 0x85, 0xd2, 0x98, 0xf4, 0x30, 0xb9, 0x03, 0x4c, 0xe5, 0x3b, 0xee, 0x92, 0x87, 0xb6,
	0xb2, 0x8c, 0xc9, 0x39, 0x34, 0x47, 0x77, 0x72, 0xad, 0x69, 0x13, 0x13, 0x36, 0x30, 0x15, 0x37,
	0xeb, 0x2c, 0xc3, 0x8a, 0x96, 0xad, 0x28, 0x63, 0xa3, 0x35, 0xda, 0x70, 0xcd, 0x8b, 0xdb, 0x22,
	0xa3, 0x47, 0x56, 0xab, 0x02, 0xc8, 0xe7, 0xd0, 0x8f, 0x8b, 0x74, 0xc3, 0xe7, 0xdb, 0xb8, 0x90,
	0xf7, 0x69, 0x26, 0x68, 0x1b, 0x77, 0xf1, 0x04, 0x35, 0xba, 0xf1, 0x83, 0xcc, 0x05, 0xed, 0x58,
	0x5d, 0x0c, 0xc8, 0x00, 0xba, 0xb7, 0xab, 0x84, 0x6b, 0x91, 0x4c, 0xd3, 0xa5, 0xa0, 0x80, 0xa5,
	0x75, 0x88, 0xbc, 0x01, 0x18, 0xad, 0x56, 0xef, 0x45, 0xa1, 0x52, 0x99, 0xd3, 0x2e, 0x12, 0x6a,
	0x08, 0xb9, 0x80, 0xe3, 0x1b, 0x99, 0x65, 0xf2, 0xb7, 0x34, 0x5f, 0x4c, 0xb7, 0x2b, 0x41, 0xfb,
	0x48, 0xd9, 0x07, 0xfd, 0x3f, 0x9a, 0x00, 0x71, 0x38, 0x8b, 0x84, 0x52, 0x7c, 0x81, 0x8b, 0xba,
	0xcf, 0x1f, 0xc5, 0x16, 0xc7, 0xde, 0x61, 0x35, 0xc4, 0x58, 0xca, 0xa4, 0x5c, 0x9a, 0xa4, 0xb3,
	0xd4, 0x85, 0x35, 0xb3, 0x3c, 0x3c, 0x0a, 0xa5, 0x59, 0xaf, 0xa0, 0x15, 0x0b, 0xc4, 0xcf, 0x2d,
	0x6e, 0x23, 0xa3, 0x64, 0x4a, 0x4d, 0x13, 0xe3, 0x04, 0xcd, 0x68, 0xb2, 0x1a, 0x62, 0xda, 0x77,
	0xba, 0x8e, 0xd2, 0xb4, 0xed, 0xef, 0x81, 0xe6, 0x20, 0x4d, 0xc5, 0xa3, 0x76, 0xd6, 0xe0, 0xb7,
	0xb1, 0xe5, 0xed, 0xa3, 0x2e, 0xf8, 0x0f, 0x4a, 0xe6, 0xa5, 0x2d, 0x15, 0x60, 0xfa, 0x19, 0xab,
	0x70, 0x1b, 0x95, 0x76, 0xb8, 0xc8, 0x18, 0x3d, 0x56, 0x13, 0xcd, 0x0b, 0x91, 0xa0, 0x13, 0x4d,
	0x56, 0xc5, 0x66, 0xc5, 0xab, 0x42, 0x98, 0xc9, 0x47, 0x0a, 0xad, 0xf0, 0xd8, 0x0e, 0x20, 0x43,
	0x38, 0x71, 0xc1, 0xb5, 0xd8, 0xa4, 0x73, 0x11, 0x29, 0x74, 0xc3, 0x63, 0x4f, 0x61, 0xa3, 0x3d,
	0x91, 0x85, 0x1e, 0x27, 0xb4, 0x67, 0x67, 0x61, 0x23, 0x73, 0x99, 0xcc, 0x54, 0x26, 0x42, 0xe4,
	0xe8, 0xf6, 0xb1, 0xbd, 0x4c, 0x75, 0x8c, 0x04, 0x40, 0xec, 0xe5, 0x62, 0x62, 0x2e, 0xd2, 0x8d,
	0x3b, 0x17, 0x7d, 0x64, 0xfe, 0x4b, 0x86, 0x7c, 0x05, 0x67, 0x16, 0xbd, 0x16, 0x99, 0x28, 0x8f,
	0xd1, 0x09, 0xd2, 0xff, 0x99, 0x30, 0x3b, 0x1c, 0x57, 0xf2, 0xa7, 0x76, 0x87, 0x15, 0x60, 0xfa,
	0x9e, 0xca, 0x78, 0xad, 0x1e, 0xe8, 0x99, 0x9d, 0x99, 0x8d, 0xc8, 0x97, 0x70, 0x1a, 0xa9, 0xc5,
	0x4d, 0x9a, 0x89, 0xd9, 0x3b, 0x39, 0xe7, 0xd9, 0xa4, 0x98, 0x53, 0x82, 0x03, 0x3f, 0x71, 0x78,
	0x09, 0x93, 0x4b, 0xe8, 0x97, 0xd4, 0x89, 0xe6, 0x7a, 0xad, 0xe8, 0x0b, 0xe7, 0xa7, 0x45, 0x2d,
	0xe8, 0x5f, 0x42, 0x37, 0x0e, 0x67, 0x4c, 0xa8, 0x95, 0xcc, 0x15, 0x0a, 0x3b, 0xb6, 0x7b, 0x01,
	0x1c, 0xed, 0x02, 0x0f, 0x2d, 0x13, 0xbf, 0xae, 0x85, 0xd2, 0xcf, 0xb2, 0xee, 0xe1, 0x7c, 0xc7,
	0x8a, 0xd4, 0x62, 0x94, 0x24, 0x11, 0xcf, 0xb7, 0xe4, 0x12, 0x8e, 0x1c, 0x88, 0x05, 0xdd, 0xaf,
	0xbb, 0xc1, 0x8e, 0xc7, 0xca, 0x1c, 0xf9, 0x02, 0xda, 0xee, 0xb0, 0x29, 0x7a, 0x30, 0xf0, 0x4a,
	0x9e, 0xc3, 0x58, 0x95, 0xf4, 0x47, 0xf0, 0xb2, 0xd6, 0x74, 0x4d, 0x68, 0x08, 0xed, 0x12, 0x75,
	0x4a, 0xbd, 0xa0, 0xc6, 0x64, 0x55, 0xd6, 0xff, 0x14, 0x8e, 0xe2, 0x70, 0x86, 0x43, 0x7d, 0x6e,
	0x37, 0x9f, 0x41, 0xc7, 0xd6, 0xae, 0xb3, 0xe7, 0xb7, 0xfc, 0x7b, 0x03, 0xce, 0xdc, 0x42, 0xb5,
	0x3e, 0x5e, 0xc3, 0x21, 0xba, 0x67, 0x7b, 0x68, 0x07, 0x8e, 0xc1, 0x10, 0xfd, 0xe0, 0x7d, 0x92,
	0x0b, 0x68, 0x9a, 0x4b, 0xad, 0xa8, 0x87, 0xac, 0x7e, 0xb0, 0xf7, 0x5c, 0x33, 0x9b, 0xf4, 0xbf,
	0x83, 0x17, 0x55, 0x9f, 0xb5, 0x1e, 0x7c, 0x68, 0x59, 0xcc, 0x75, 0x01, 0x41, 0xc5, 0x62, 0x2e,
	0xe3, 0xff, 0xd9, 0x40, 0xfb, 0x23, 0xb5, 0x78, 0xbb, 0x31, 0x3f, 0x32, 0x1f, 0xff, 0x1a, 0xbd,
	0x01, 0x30, 0xb7, 0x67, 0xef, 0x45, 0xaa, 0x21, 0xf8, 0x46, 0x18, 0x09, 0x7c, 0x18, 0xed, 0xe3,
	0xb3, 0x03, 0xcc, 0x74, 0x47, 0xba, 0x7a, 0x77, 0x3d, 0xe6, 0x22, 0xff, 0x3d, 0x9c, 0xee, 0x86,
	0x8b, 0x74, 0xf5, 0x1f, 0xb3, 0xbd, 0x80, 0x96, 0xe5, 0xb9, 0xc9, 0xa2, 0xff, 0x65, 0x31, 0x73,
	0x39, 0xff, 0x5b, 0xfc, 0x89, 0xad, 0x46, 0xe6, 0x56, 0xfe, 0x80, 0x89, 0x85, 0x2f, 0xa1, 0x9d,
	0x16, 0xc1, 0x52, 0x05, 0xab, 0xbb, 0xf0, 0x20, 0x0e, 0xbf, 0xf7, 0xe2, 0xc6, 0x2f, 0x8d, 0xc7,
	0xbb, 0x16, 0xfe, 0x45, 0xf8, 0xe6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x54, 0x75, 0x2d,
	0x32, 0x08, 0x00, 0x00,
}
