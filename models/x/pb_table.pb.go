// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_table.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PB_Chat struct {
	ChatId         int64  `protobuf:"varint,1,opt,name=ChatId" json:"ChatId,omitempty"`
	ChatKey        string `protobuf:"bytes,3,opt,name=ChatKey" json:"ChatKey,omitempty"`
	RoomTypeEnumId int32  `protobuf:"varint,5,opt,name=RoomTypeEnumId" json:"RoomTypeEnumId,omitempty"`
	UserId         int32  `protobuf:"varint,7,opt,name=UserId" json:"UserId,omitempty"`
	LastSeqSeen    int32  `protobuf:"varint,9,opt,name=LastSeqSeen" json:"LastSeqSeen,omitempty"`
	LastSeqDelete  int32  `protobuf:"varint,11,opt,name=LastSeqDelete" json:"LastSeqDelete,omitempty"`
	PeerUserId     int32  `protobuf:"varint,13,opt,name=PeerUserId" json:"PeerUserId,omitempty"`
	GroupId        int64  `protobuf:"varint,15,opt,name=GroupId" json:"GroupId,omitempty"`
	CreatedTime    int32  `protobuf:"varint,17,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	CurrentSeq     int32  `protobuf:"varint,19,opt,name=CurrentSeq" json:"CurrentSeq,omitempty"`
	UpdatedMs      int64  `protobuf:"varint,21,opt,name=UpdatedMs" json:"UpdatedMs,omitempty"`
}

func (m *PB_Chat) Reset()                    { *m = PB_Chat{} }
func (m *PB_Chat) String() string            { return proto.CompactTextString(m) }
func (*PB_Chat) ProtoMessage()               {}
func (*PB_Chat) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *PB_Chat) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *PB_Chat) GetChatKey() string {
	if m != nil {
		return m.ChatKey
	}
	return ""
}

func (m *PB_Chat) GetRoomTypeEnumId() int32 {
	if m != nil {
		return m.RoomTypeEnumId
	}
	return 0
}

func (m *PB_Chat) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_Chat) GetLastSeqSeen() int32 {
	if m != nil {
		return m.LastSeqSeen
	}
	return 0
}

func (m *PB_Chat) GetLastSeqDelete() int32 {
	if m != nil {
		return m.LastSeqDelete
	}
	return 0
}

func (m *PB_Chat) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *PB_Chat) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_Chat) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *PB_Chat) GetCurrentSeq() int32 {
	if m != nil {
		return m.CurrentSeq
	}
	return 0
}

func (m *PB_Chat) GetUpdatedMs() int64 {
	if m != nil {
		return m.UpdatedMs
	}
	return 0
}

type PB_DirectMessage struct {
	MessageId          int64  `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
	RoomKey            string `protobuf:"bytes,3,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId             int32  `protobuf:"varint,5,opt,name=UserId" json:"UserId,omitempty"`
	MessageFileId      int64  `protobuf:"varint,7,opt,name=MessageFileId" json:"MessageFileId,omitempty"`
	MessageTypeEnum    int32  `protobuf:"varint,9,opt,name=MessageTypeEnum" json:"MessageTypeEnum,omitempty"`
	Text               string `protobuf:"bytes,11,opt,name=Text" json:"Text,omitempty"`
	Time               int32  `protobuf:"varint,13,opt,name=Time" json:"Time,omitempty"`
	PeerReceivedTime   int32  `protobuf:"varint,15,opt,name=PeerReceivedTime" json:"PeerReceivedTime,omitempty"`
	PeerSeenTime       int32  `protobuf:"varint,17,opt,name=PeerSeenTime" json:"PeerSeenTime,omitempty"`
	DeliviryStatusEnum int32  `protobuf:"varint,19,opt,name=DeliviryStatusEnum" json:"DeliviryStatusEnum,omitempty"`
}

func (m *PB_DirectMessage) Reset()                    { *m = PB_DirectMessage{} }
func (m *PB_DirectMessage) String() string            { return proto.CompactTextString(m) }
func (*PB_DirectMessage) ProtoMessage()               {}
func (*PB_DirectMessage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *PB_DirectMessage) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_DirectMessage) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_DirectMessage) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_DirectMessage) GetMessageFileId() int64 {
	if m != nil {
		return m.MessageFileId
	}
	return 0
}

func (m *PB_DirectMessage) GetMessageTypeEnum() int32 {
	if m != nil {
		return m.MessageTypeEnum
	}
	return 0
}

func (m *PB_DirectMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PB_DirectMessage) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PB_DirectMessage) GetPeerReceivedTime() int32 {
	if m != nil {
		return m.PeerReceivedTime
	}
	return 0
}

func (m *PB_DirectMessage) GetPeerSeenTime() int32 {
	if m != nil {
		return m.PeerSeenTime
	}
	return 0
}

func (m *PB_DirectMessage) GetDeliviryStatusEnum() int32 {
	if m != nil {
		return m.DeliviryStatusEnum
	}
	return 0
}

func init() {
	proto.RegisterType((*PB_Chat)(nil), "PB_Chat")
	proto.RegisterType((*PB_DirectMessage)(nil), "PB_DirectMessage")
}

func init() { proto.RegisterFile("pb_table.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6e, 0xaa, 0x40,
	0x10, 0x80, 0x83, 0x1c, 0xf5, 0x30, 0x1e, 0xff, 0xf6, 0xe4, 0x9c, 0xec, 0x45, 0xd3, 0x10, 0x63,
	0x1a, 0xd2, 0x0b, 0x6e, 0xfa, 0x06, 0x6a, 0x7f, 0x48, 0x6b, 0x42, 0x50, 0x6f, 0x7a, 0x63, 0x40,
	0x26, 0x2d, 0x09, 0x08, 0x5d, 0x16, 0xa3, 0xaf, 0xda, 0xa4, 0xef, 0xd2, 0xec, 0xb2, 0x56, 0xb0,
	0xbd, 0x62, 0xe6, 0x9b, 0xdd, 0xd9, 0xd9, 0x6f, 0x81, 0x5e, 0x16, 0xac, 0xb9, 0x1f, 0xc4, 0x68,
	0x67, 0x2c, 0xe5, 0xe9, 0xe8, 0xa3, 0x01, 0x6d, 0x77, 0xb2, 0x9e, 0xbe, 0xfa, 0x9c, 0xfc, 0x87,
	0x96, 0xf8, 0x3a, 0x21, 0xd5, 0x4c, 0xcd, 0xd2, 0x3d, 0x95, 0x11, 0x0a, 0x6d, 0x11, 0x3d, 0xe2,
	0x81, 0xea, 0xa6, 0x66, 0x19, 0xde, 0x31, 0x25, 0x57, 0xd0, 0xf3, 0xd2, 0x34, 0x59, 0x1e, 0x32,
	0xbc, 0xdd, 0x16, 0x89, 0x13, 0xd2, 0xa6, 0xa9, 0x59, 0x4d, 0xef, 0x8c, 0x8a, 0xce, 0xab, 0x1c,
	0x99, 0x13, 0xd2, 0xb6, 0xac, 0xab, 0x8c, 0x98, 0xd0, 0x79, 0xf2, 0x73, 0xbe, 0xc0, 0xb7, 0x05,
	0xe2, 0x96, 0x1a, 0xb2, 0x58, 0x45, 0x64, 0x0c, 0x5d, 0x95, 0xce, 0x30, 0x46, 0x8e, 0xb4, 0x23,
	0xd7, 0xd4, 0x21, 0xb9, 0x04, 0x70, 0x11, 0x99, 0x3a, 0xa3, 0x2b, 0x97, 0x54, 0x88, 0xb8, 0xc1,
	0x3d, 0x4b, 0x8b, 0xcc, 0x09, 0x69, 0x5f, 0x5e, 0xed, 0x98, 0x8a, 0x09, 0xa6, 0x0c, 0x7d, 0x8e,
	0xe1, 0x32, 0x4a, 0x90, 0x0e, 0xcb, 0x09, 0x2a, 0x48, 0xf4, 0x9e, 0x16, 0x8c, 0xe1, 0x56, 0x9c,
	0x47, 0xff, 0x96, 0xbd, 0x4f, 0x84, 0x5c, 0x80, 0xb1, 0xca, 0x42, 0xb1, 0x7c, 0x9e, 0xd3, 0x7f,
	0xb2, 0xfb, 0x09, 0x8c, 0xde, 0x1b, 0x30, 0x70, 0x27, 0xeb, 0x59, 0xc4, 0x70, 0xc3, 0xe7, 0x98,
	0xe7, 0xfe, 0x0b, 0x8a, 0x2d, 0x2a, 0xfc, 0x72, 0x7d, 0x02, 0x62, 0x58, 0xa1, 0xaf, 0xa2, 0x5b,
	0xa5, 0x15, 0x8d, 0xcd, 0x9a, 0xc6, 0x31, 0x74, 0xd5, 0xf6, 0xbb, 0x28, 0x46, 0x65, 0x59, 0xf7,
	0xea, 0x90, 0x58, 0xd0, 0x57, 0xe0, 0xf8, 0x32, 0x4a, 0xf8, 0x39, 0x26, 0x04, 0x7e, 0x2d, 0x71,
	0xcf, 0xa5, 0x6b, 0xc3, 0x93, 0xb1, 0x64, 0xc2, 0x50, 0x29, 0x57, 0xc6, 0xe4, 0x1a, 0x06, 0x42,
	0xb2, 0x87, 0x1b, 0x8c, 0x76, 0xca, 0x60, 0x5f, 0xd6, 0xbf, 0x71, 0x32, 0x82, 0x3f, 0x82, 0x89,
	0x47, 0xad, 0x98, 0xae, 0x31, 0x62, 0x03, 0x99, 0x61, 0x1c, 0xed, 0x22, 0x76, 0x58, 0x70, 0x9f,
	0x17, 0xb9, 0x1c, 0xb2, 0x54, 0xfe, 0x43, 0x65, 0x32, 0x84, 0xdf, 0x11, 0xb3, 0x93, 0xdc, 0xce,
	0x82, 0x07, 0xdd, 0xd5, 0x9e, 0xb5, 0x7d, 0xd0, 0x92, 0xbf, 0xf5, 0xcd, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x73, 0x1f, 0x1f, 0x3f, 0xe8, 0x02, 0x00, 0x00,
}
