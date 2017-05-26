// Code generated by protoc-gen-go.
// source: pb_seen_msgs.proto
// DO NOT EDIT!

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PB_MsgsSeenFromClient struct {
	MessageKey string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey    string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	AtTime     int64  `protobuf:"varint,10,opt,name=AtTime" json:"AtTime,omitempty"`
}

func (m *PB_MsgsSeenFromClient) Reset()                    { *m = PB_MsgsSeenFromClient{} }
func (m *PB_MsgsSeenFromClient) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgsSeenFromClient) ProtoMessage()               {}
func (*PB_MsgsSeenFromClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PB_MsgsSeenFromClient) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PB_MsgsSeenFromClient) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_MsgsSeenFromClient) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_MsgsSeenFromClient) GetAtTime() int64 {
	if m != nil {
		return m.AtTime
	}
	return 0
}

type PB_MsgsSeenToClient struct {
	MessageKey string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey    string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	AtTime     int64  `protobuf:"varint,10,opt,name=AtTime" json:"AtTime,omitempty"`
}

func (m *PB_MsgsSeenToClient) Reset()                    { *m = PB_MsgsSeenToClient{} }
func (m *PB_MsgsSeenToClient) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgsSeenToClient) ProtoMessage()               {}
func (*PB_MsgsSeenToClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PB_MsgsSeenToClient) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PB_MsgsSeenToClient) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_MsgsSeenToClient) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_MsgsSeenToClient) GetAtTime() int64 {
	if m != nil {
		return m.AtTime
	}
	return 0
}

type PB_MsgSeen struct {
	MessageKey string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey    string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	AtTime     int64  `protobuf:"varint,10,opt,name=AtTime" json:"AtTime,omitempty"`
}

func (m *PB_MsgSeen) Reset()                    { *m = PB_MsgSeen{} }
func (m *PB_MsgSeen) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgSeen) ProtoMessage()               {}
func (*PB_MsgSeen) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *PB_MsgSeen) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PB_MsgSeen) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_MsgSeen) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_MsgSeen) GetAtTime() int64 {
	if m != nil {
		return m.AtTime
	}
	return 0
}

type PB_RequestMsgsSeen struct {
	Request *PB_Request   `protobuf:"bytes,1,opt,name=Request" json:"Request,omitempty"`
	Seen    []*PB_MsgSeen `protobuf:"bytes,2,rep,name=Seen" json:"Seen,omitempty"`
}

func (m *PB_RequestMsgsSeen) Reset()                    { *m = PB_RequestMsgsSeen{} }
func (m *PB_RequestMsgsSeen) String() string            { return proto.CompactTextString(m) }
func (*PB_RequestMsgsSeen) ProtoMessage()               {}
func (*PB_RequestMsgsSeen) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *PB_RequestMsgsSeen) GetRequest() *PB_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PB_RequestMsgsSeen) GetSeen() []*PB_MsgSeen {
	if m != nil {
		return m.Seen
	}
	return nil
}

type PB_ResponseMsgsSeen struct {
	Response *PB_Response `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PB_ResponseMsgsSeen) Reset()                    { *m = PB_ResponseMsgsSeen{} }
func (m *PB_ResponseMsgsSeen) String() string            { return proto.CompactTextString(m) }
func (*PB_ResponseMsgsSeen) ProtoMessage()               {}
func (*PB_ResponseMsgsSeen) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *PB_ResponseMsgsSeen) GetResponse() *PB_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type PB_PushMsgsReceivedToPeer struct {
	Push *PB_Push      `protobuf:"bytes,1,opt,name=Push" json:"Push,omitempty"`
	Seen []*PB_MsgSeen `protobuf:"bytes,2,rep,name=Seen" json:"Seen,omitempty"`
}

func (m *PB_PushMsgsReceivedToPeer) Reset()                    { *m = PB_PushMsgsReceivedToPeer{} }
func (m *PB_PushMsgsReceivedToPeer) String() string            { return proto.CompactTextString(m) }
func (*PB_PushMsgsReceivedToPeer) ProtoMessage()               {}
func (*PB_PushMsgsReceivedToPeer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *PB_PushMsgsReceivedToPeer) GetPush() *PB_Push {
	if m != nil {
		return m.Push
	}
	return nil
}

func (m *PB_PushMsgsReceivedToPeer) GetSeen() []*PB_MsgSeen {
	if m != nil {
		return m.Seen
	}
	return nil
}

type PB_ResultMsgsReceivedToPeer struct {
	Result *PB_Result `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *PB_ResultMsgsReceivedToPeer) Reset()                    { *m = PB_ResultMsgsReceivedToPeer{} }
func (m *PB_ResultMsgsReceivedToPeer) String() string            { return proto.CompactTextString(m) }
func (*PB_ResultMsgsReceivedToPeer) ProtoMessage()               {}
func (*PB_ResultMsgsReceivedToPeer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *PB_ResultMsgsReceivedToPeer) GetResult() *PB_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_MsgsSeenFromClient)(nil), "PB_MsgsSeenFromClient")
	proto.RegisterType((*PB_MsgsSeenToClient)(nil), "PB_MsgsSeenToClient")
	proto.RegisterType((*PB_MsgSeen)(nil), "PB_MsgSeen")
	proto.RegisterType((*PB_RequestMsgsSeen)(nil), "PB_RequestMsgsSeen")
	proto.RegisterType((*PB_ResponseMsgsSeen)(nil), "PB_ResponseMsgsSeen")
	proto.RegisterType((*PB_PushMsgsReceivedToPeer)(nil), "PB_PushMsgsReceivedToPeer")
	proto.RegisterType((*PB_ResultMsgsReceivedToPeer)(nil), "PB_ResultMsgsReceivedToPeer")
}

func init() { proto.RegisterFile("pb_seen_msgs.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x53, 0xda, 0x3a, 0xed, 0xc5, 0x15, 0x25, 0xfe, 0x41, 0x4b, 0x40, 0xe8, 0x69,
	0x0f, 0xfa, 0x01, 0xa4, 0x11, 0x44, 0x91, 0x42, 0x58, 0xeb, 0xa5, 0x08, 0xa1, 0xb1, 0x43, 0x0c,
	0x24, 0xd9, 0x98, 0x49, 0x8b, 0x9e, 0xc4, 0x6f, 0x2e, 0xfb, 0x27, 0x35, 0x07, 0xc1, 0x5b, 0x8f,
	0xf3, 0x7e, 0x2f, 0x6f, 0xde, 0x84, 0x05, 0x56, 0xc6, 0x11, 0x21, 0x16, 0x51, 0x4e, 0x09, 0xf1,
	0xb2, 0x92, 0xb5, 0x3c, 0x19, 0x25, 0x99, 0x8c, 0x97, 0x99, 0x99, 0xfc, 0x6f, 0x07, 0x0e, 0xc3,
	0x20, 0x9a, 0x51, 0x42, 0x4f, 0x88, 0xc5, 0x5d, 0x25, 0xf3, 0xdb, 0x2c, 0xc5, 0xa2, 0x66, 0xe7,
	0x00, 0x33, 0x24, 0x5a, 0x26, 0xf8, 0x88, 0x9f, 0x9e, 0x33, 0x76, 0x26, 0x7b, 0xa2, 0xa5, 0x30,
	0x0f, 0xfa, 0x42, 0xca, 0x5c, 0xc1, 0x8e, 0x86, 0xcd, 0xc8, 0x8e, 0xa0, 0xf7, 0x4c, 0x58, 0x3d,
	0xac, 0x3c, 0x77, 0xec, 0x4c, 0x5c, 0x61, 0x27, 0xa5, 0x4f, 0xeb, 0x79, 0x9a, 0xa3, 0x07, 0x46,
	0x37, 0x93, 0xff, 0x05, 0x07, 0xad, 0x0a, 0x73, 0xb9, 0xf3, 0x02, 0x1b, 0x00, 0x53, 0x40, 0xed,
	0xdf, 0xe1, 0xde, 0x17, 0x60, 0x61, 0x10, 0x09, 0x7c, 0x5f, 0x23, 0xd5, 0xcd, 0xfd, 0xec, 0x12,
	0xfa, 0x56, 0xd2, 0xcb, 0x87, 0x57, 0x43, 0xfe, 0xeb, 0x12, 0x0d, 0x63, 0x17, 0xd0, 0x55, 0x76,
	0xaf, 0x33, 0x76, 0x1b, 0x8f, 0xbd, 0x40, 0x68, 0xe0, 0xdf, 0xe8, 0xdf, 0x2a, 0x90, 0x4a, 0x59,
	0x10, 0x6e, 0xe3, 0x27, 0x30, 0x68, 0x34, 0x9b, 0x3f, 0xe2, 0x2d, 0x9f, 0xd8, 0x52, 0x7f, 0x01,
	0xc7, 0x61, 0x10, 0x85, 0x6b, 0x7a, 0x53, 0x1f, 0x0b, 0x7c, 0xc5, 0x74, 0x83, 0xab, 0xb9, 0x0c,
	0x11, 0x2b, 0x76, 0x06, 0x5d, 0x45, 0x6c, 0xc4, 0x80, 0x5b, 0xa7, 0xd0, 0xea, 0xff, 0xe5, 0xa6,
	0x70, 0x6a, 0x96, 0xae, 0xb3, 0xfa, 0x8f, 0x74, 0x1f, 0x7a, 0x86, 0xd9, 0x7c, 0xe0, 0x5b, 0xb7,
	0xb0, 0x24, 0xd8, 0x87, 0x41, 0x5a, 0xf1, 0x9c, 0x78, 0x19, 0xdf, 0xbb, 0xa1, 0xb3, 0x70, 0x3e,
	0xe2, 0x9e, 0x7e, 0xd4, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xbc, 0xfb, 0x2b, 0xf8,
	0x02, 0x00, 0x00,
}