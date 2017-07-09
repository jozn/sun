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
func (*PB_MsgsSeenFromClient) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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

type PB_MsgSeen struct {
	MessageKey string `protobuf:"bytes,1,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey    string `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	AtTime     int64  `protobuf:"varint,10,opt,name=AtTime" json:"AtTime,omitempty"`
}

func (m *PB_MsgSeen) Reset()                    { *m = PB_MsgSeen{} }
func (m *PB_MsgSeen) String() string            { return proto.CompactTextString(m) }
func (*PB_MsgSeen) ProtoMessage()               {}
func (*PB_MsgSeen) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
func (*PB_RequestMsgsSeen) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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
func (*PB_ResponseMsgsSeen) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *PB_ResponseMsgsSeen) GetResponse() *PB_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_MsgsSeenFromClient)(nil), "PB_MsgsSeenFromClient")
	proto.RegisterType((*PB_MsgSeen)(nil), "PB_MsgSeen")
	proto.RegisterType((*PB_RequestMsgsSeen)(nil), "PB_RequestMsgsSeen")
	proto.RegisterType((*PB_ResponseMsgsSeen)(nil), "PB_ResponseMsgsSeen")
}

func init() { proto.RegisterFile("pb_seen_msgs.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0x56, 0x76, 0xd7, 0x59, 0x41, 0x8c, 0x28, 0xc1, 0x83, 0x2e, 0x05, 0xa1, 0xa7,
	0x1c, 0xf4, 0x01, 0xc4, 0x0a, 0xa2, 0xc8, 0x42, 0x89, 0x7a, 0x11, 0x21, 0x6c, 0x70, 0x28, 0x85,
	0xa6, 0x89, 0x9d, 0x2a, 0x7a, 0xf4, 0xcd, 0xa5, 0x69, 0xb3, 0xee, 0x13, 0x78, 0x9c, 0xef, 0xfb,
	0x33, 0x33, 0x49, 0x80, 0x7b, 0xa3, 0x09, 0xb1, 0xd1, 0x96, 0x4a, 0x92, 0xbe, 0x75, 0x9d, 0x3b,
	0xd9, 0xf7, 0x46, 0x97, 0xb5, 0x33, 0xeb, 0x7a, 0x00, 0xe9, 0x0f, 0x83, 0xa3, 0x22, 0xd7, 0x2b,
	0x2a, 0xe9, 0x11, 0xb1, 0xb9, 0x6d, 0x9d, 0xbd, 0xa9, 0x2b, 0x6c, 0x3a, 0x7e, 0x0a, 0xb0, 0x42,
	0xa2, 0x75, 0x89, 0x0f, 0xf8, 0x2d, 0xd8, 0x92, 0x65, 0xbb, 0x6a, 0x8b, 0x70, 0x01, 0x33, 0xe5,
	0x9c, 0xed, 0xe5, 0x24, 0xc8, 0x58, 0xf2, 0x63, 0x98, 0x3e, 0x13, 0xb6, 0xf7, 0x6f, 0x22, 0x59,
	0xb2, 0x2c, 0x51, 0x63, 0xd5, 0xf3, 0xeb, 0xee, 0xa9, 0xb2, 0x28, 0x60, 0xe0, 0x43, 0x95, 0x7e,
	0x02, 0x0c, 0x2b, 0xf4, 0x1b, 0xfc, 0xe3, 0xdc, 0x57, 0xe0, 0x45, 0xae, 0x15, 0xbe, 0x7f, 0x20,
	0x75, 0xf1, 0x05, 0xf8, 0x39, 0xcc, 0x46, 0x14, 0x86, 0x2f, 0x2e, 0x16, 0xf2, 0x2f, 0xa5, 0xa2,
	0xe3, 0x67, 0xb0, 0xd3, 0xc7, 0xc5, 0x64, 0x99, 0xc4, 0xcc, 0x78, 0x03, 0x15, 0x44, 0x7a, 0x05,
	0x87, 0xe1, 0x1c, 0x79, 0xd7, 0x10, 0x6e, 0xda, 0x67, 0x30, 0x8f, 0x6c, 0xec, 0xbf, 0x27, 0xb7,
	0x72, 0x6a, 0x63, 0xf3, 0x03, 0x98, 0x57, 0xad, 0xb4, 0x24, 0xbd, 0xb9, 0x4b, 0x0a, 0xf6, 0xc2,
	0xbe, 0xcc, 0x34, 0x7c, 0xda, 0xe5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x2c, 0x05, 0xca,
	0xdb, 0x01, 0x00, 0x00,
}
