// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_views.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// {realm} ,
type PB_ChatView struct {
	// option realm = "1";
	ChatKey              string       `protobuf:"bytes,1,opt,name=ChatKey" json:"ChatKey,omitempty"`
	RoomKey              string       `protobuf:"bytes,3,opt,name=RoomKey" json:"RoomKey,omitempty"`
	RoomTypeEnumId       int32        `protobuf:"varint,5,opt,name=RoomTypeEnumId" json:"RoomTypeEnumId,omitempty"`
	UserId               int32        `protobuf:"varint,7,opt,name=UserId" json:"UserId,omitempty"`
	PeerUserId           int32        `protobuf:"varint,9,opt,name=PeerUserId" json:"PeerUserId,omitempty"`
	GroupId              int64        `protobuf:"varint,11,opt,name=GroupId" json:"GroupId,omitempty"`
	CreatedSe            int32        `protobuf:"varint,13,opt,name=CreatedSe" json:"CreatedSe,omitempty"`
	UpdatedMs            int64        `protobuf:"varint,15,opt,name=UpdatedMs" json:"UpdatedMs,omitempty"`
	LastMessageId        int64        `protobuf:"varint,17,opt,name=LastMessageId" json:"LastMessageId,omitempty"`
	LastDeletedMessageId int64        `protobuf:"varint,19,opt,name=LastDeletedMessageId" json:"LastDeletedMessageId,omitempty"`
	LastSeenMessageId    int64        `protobuf:"varint,21,opt,name=LastSeenMessageId" json:"LastSeenMessageId,omitempty"`
	LastSeqSeen          int32        `protobuf:"varint,23,opt,name=LastSeqSeen" json:"LastSeqSeen,omitempty"`
	LastSeqDelete        int32        `protobuf:"varint,25,opt,name=LastSeqDelete" json:"LastSeqDelete,omitempty"`
	CurrentSeq           int32        `protobuf:"varint,27,opt,name=CurrentSeq" json:"CurrentSeq,omitempty"`
	UserView             *PB_UserView `protobuf:"bytes,100,opt,name=UserView" json:"UserView,omitempty"`
	// local
	SharedMediaCount   int32           `protobuf:"varint,200,opt,name=SharedMediaCount" json:"SharedMediaCount,omitempty"`
	UnseenCount        int32           `protobuf:"varint,205,opt,name=UnseenCount" json:"UnseenCount,omitempty"`
	FirstUnreadMessage *PB_MessageView `protobuf:"bytes,210,opt,name=FirstUnreadMessage" json:"FirstUnreadMessage,omitempty"`
	LastMessage        *PB_MessageView `protobuf:"bytes,212,opt,name=LastMessage" json:"LastMessage,omitempty"`
}

func (m *PB_ChatView) Reset()                    { *m = PB_ChatView{} }
func (m *PB_ChatView) String() string            { return proto.CompactTextString(m) }
func (*PB_ChatView) ProtoMessage()               {}
func (*PB_ChatView) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *PB_ChatView) GetChatKey() string {
	if m != nil {
		return m.ChatKey
	}
	return ""
}

func (m *PB_ChatView) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_ChatView) GetRoomTypeEnumId() int32 {
	if m != nil {
		return m.RoomTypeEnumId
	}
	return 0
}

func (m *PB_ChatView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_ChatView) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *PB_ChatView) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_ChatView) GetCreatedSe() int32 {
	if m != nil {
		return m.CreatedSe
	}
	return 0
}

func (m *PB_ChatView) GetUpdatedMs() int64 {
	if m != nil {
		return m.UpdatedMs
	}
	return 0
}

func (m *PB_ChatView) GetLastMessageId() int64 {
	if m != nil {
		return m.LastMessageId
	}
	return 0
}

func (m *PB_ChatView) GetLastDeletedMessageId() int64 {
	if m != nil {
		return m.LastDeletedMessageId
	}
	return 0
}

func (m *PB_ChatView) GetLastSeenMessageId() int64 {
	if m != nil {
		return m.LastSeenMessageId
	}
	return 0
}

func (m *PB_ChatView) GetLastSeqSeen() int32 {
	if m != nil {
		return m.LastSeqSeen
	}
	return 0
}

func (m *PB_ChatView) GetLastSeqDelete() int32 {
	if m != nil {
		return m.LastSeqDelete
	}
	return 0
}

func (m *PB_ChatView) GetCurrentSeq() int32 {
	if m != nil {
		return m.CurrentSeq
	}
	return 0
}

func (m *PB_ChatView) GetUserView() *PB_UserView {
	if m != nil {
		return m.UserView
	}
	return nil
}

func (m *PB_ChatView) GetSharedMediaCount() int32 {
	if m != nil {
		return m.SharedMediaCount
	}
	return 0
}

func (m *PB_ChatView) GetUnseenCount() int32 {
	if m != nil {
		return m.UnseenCount
	}
	return 0
}

func (m *PB_ChatView) GetFirstUnreadMessage() *PB_MessageView {
	if m != nil {
		return m.FirstUnreadMessage
	}
	return nil
}

func (m *PB_ChatView) GetLastMessage() *PB_MessageView {
	if m != nil {
		return m.LastMessage
	}
	return nil
}

// {realm} ,
type PB_MessageView struct {
	MessageId            int64               `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
	MessageKey           string              `protobuf:"bytes,3,opt,name=MessageKey" json:"MessageKey,omitempty"`
	RoomKey              string              `protobuf:"bytes,5,opt,name=RoomKey" json:"RoomKey,omitempty"`
	UserId               int32               `protobuf:"varint,7,opt,name=UserId" json:"UserId,omitempty"`
	MessageFileId        int64               `protobuf:"varint,9,opt,name=MessageFileId" json:"MessageFileId,omitempty"`
	MessageTypeEnumId    int32               `protobuf:"varint,11,opt,name=MessageTypeEnumId" json:"MessageTypeEnumId,omitempty"`
	Text                 string              `protobuf:"bytes,13,opt,name=Text" json:"Text,omitempty"`
	CreatedSe            int32               `protobuf:"varint,15,opt,name=CreatedSe" json:"CreatedSe,omitempty"`
	PeerReceivedTime     int32               `protobuf:"varint,17,opt,name=PeerReceivedTime" json:"PeerReceivedTime,omitempty"`
	PeerSeenTime         int32               `protobuf:"varint,19,opt,name=PeerSeenTime" json:"PeerSeenTime,omitempty"`
	DeliviryStatusEnumId int32               `protobuf:"varint,21,opt,name=DeliviryStatusEnumId" json:"DeliviryStatusEnumId,omitempty"`
	ChatKey              string              `protobuf:"bytes,30,opt,name=ChatKey" json:"ChatKey,omitempty"`
	RoomTypeEnumId       int32               `protobuf:"varint,31,opt,name=RoomTypeEnumId" json:"RoomTypeEnumId,omitempty"`
	IsByMe               bool                `protobuf:"varint,33,opt,name=IsByMe" json:"IsByMe,omitempty"`
	RemoteId             int64               `protobuf:"varint,50,opt,name=RemoteId" json:"RemoteId,omitempty"`
	MessageFileView      *PB_MessageFileView `protobuf:"bytes,100,opt,name=MessageFileView" json:"MessageFileView,omitempty"`
}

func (m *PB_MessageView) Reset()                    { *m = PB_MessageView{} }
func (m *PB_MessageView) String() string            { return proto.CompactTextString(m) }
func (*PB_MessageView) ProtoMessage()               {}
func (*PB_MessageView) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *PB_MessageView) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_MessageView) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *PB_MessageView) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_MessageView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_MessageView) GetMessageFileId() int64 {
	if m != nil {
		return m.MessageFileId
	}
	return 0
}

func (m *PB_MessageView) GetMessageTypeEnumId() int32 {
	if m != nil {
		return m.MessageTypeEnumId
	}
	return 0
}

func (m *PB_MessageView) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PB_MessageView) GetCreatedSe() int32 {
	if m != nil {
		return m.CreatedSe
	}
	return 0
}

func (m *PB_MessageView) GetPeerReceivedTime() int32 {
	if m != nil {
		return m.PeerReceivedTime
	}
	return 0
}

func (m *PB_MessageView) GetPeerSeenTime() int32 {
	if m != nil {
		return m.PeerSeenTime
	}
	return 0
}

func (m *PB_MessageView) GetDeliviryStatusEnumId() int32 {
	if m != nil {
		return m.DeliviryStatusEnumId
	}
	return 0
}

func (m *PB_MessageView) GetChatKey() string {
	if m != nil {
		return m.ChatKey
	}
	return ""
}

func (m *PB_MessageView) GetRoomTypeEnumId() int32 {
	if m != nil {
		return m.RoomTypeEnumId
	}
	return 0
}

func (m *PB_MessageView) GetIsByMe() bool {
	if m != nil {
		return m.IsByMe
	}
	return false
}

func (m *PB_MessageView) GetRemoteId() int64 {
	if m != nil {
		return m.RemoteId
	}
	return 0
}

func (m *PB_MessageView) GetMessageFileView() *PB_MessageFileView {
	if m != nil {
		return m.MessageFileView
	}
	return nil
}

// {realm} ,
type PB_MessageFileView struct {
	//    option realm = "1";
	MessageFileId  int64  `protobuf:"varint,1,opt,name=MessageFileId" json:"MessageFileId,omitempty"`
	MessageFileKey string `protobuf:"bytes,32,opt,name=MessageFileKey" json:"MessageFileKey,omitempty"`
	OriginalUserId int32  `protobuf:"varint,2,opt,name=OriginalUserId" json:"OriginalUserId,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Size           int32  `protobuf:"varint,5,opt,name=Size" json:"Size,omitempty"`
	FileTypeEnumId int32  `protobuf:"varint,7,opt,name=FileTypeEnumId" json:"FileTypeEnumId,omitempty"`
	Width          int32  `protobuf:"varint,9,opt,name=Width" json:"Width,omitempty"`
	Height         int32  `protobuf:"varint,11,opt,name=Height" json:"Height,omitempty"`
	Duration       int32  `protobuf:"varint,13,opt,name=Duration" json:"Duration,omitempty"`
	Extension      string `protobuf:"bytes,15,opt,name=Extension" json:"Extension,omitempty"`
	HashMd5        string `protobuf:"bytes,17,opt,name=HashMd5" json:"HashMd5,omitempty"`
	HashAccess     int64  `protobuf:"varint,19,opt,name=HashAccess" json:"HashAccess,omitempty"`
	CreatedSe      int32  `protobuf:"varint,21,opt,name=CreatedSe" json:"CreatedSe,omitempty"`
	ServerSrc      string `protobuf:"bytes,23,opt,name=ServerSrc" json:"ServerSrc,omitempty"`
	// del???
	ServerPath      string `protobuf:"bytes,25,opt,name=ServerPath" json:"ServerPath,omitempty"`
	ServerThumbPath string `protobuf:"bytes,27,opt,name=ServerThumbPath" json:"ServerThumbPath,omitempty"`
	BucketId        string `protobuf:"bytes,29,opt,name=BucketId" json:"BucketId,omitempty"`
	ServerId        int32  `protobuf:"varint,31,opt,name=ServerId" json:"ServerId,omitempty"`
	CanDel          int32  `protobuf:"varint,33,opt,name=CanDel" json:"CanDel,omitempty"`
	// just views
	ServerThumbLocalSrc string `protobuf:"bytes,50,opt,name=ServerThumbLocalSrc" json:"ServerThumbLocalSrc,omitempty"`
	// local
	RemoteMessageFileId int64  `protobuf:"varint,100,opt,name=RemoteMessageFileId" json:"RemoteMessageFileId,omitempty"`
	LocalSrc            string `protobuf:"bytes,101,opt,name=LocalSrc" json:"LocalSrc,omitempty"`
	ThumbLocalSrc       string `protobuf:"bytes,104,opt,name=ThumbLocalSrc" json:"ThumbLocalSrc,omitempty"`
	MessageFileStatusId string `protobuf:"bytes,106,opt,name=MessageFileStatusId" json:"MessageFileStatusId,omitempty"`
}

func (m *PB_MessageFileView) Reset()                    { *m = PB_MessageFileView{} }
func (m *PB_MessageFileView) String() string            { return proto.CompactTextString(m) }
func (*PB_MessageFileView) ProtoMessage()               {}
func (*PB_MessageFileView) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *PB_MessageFileView) GetMessageFileId() int64 {
	if m != nil {
		return m.MessageFileId
	}
	return 0
}

func (m *PB_MessageFileView) GetMessageFileKey() string {
	if m != nil {
		return m.MessageFileKey
	}
	return ""
}

func (m *PB_MessageFileView) GetOriginalUserId() int32 {
	if m != nil {
		return m.OriginalUserId
	}
	return 0
}

func (m *PB_MessageFileView) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PB_MessageFileView) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PB_MessageFileView) GetFileTypeEnumId() int32 {
	if m != nil {
		return m.FileTypeEnumId
	}
	return 0
}

func (m *PB_MessageFileView) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *PB_MessageFileView) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PB_MessageFileView) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PB_MessageFileView) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *PB_MessageFileView) GetHashMd5() string {
	if m != nil {
		return m.HashMd5
	}
	return ""
}

func (m *PB_MessageFileView) GetHashAccess() int64 {
	if m != nil {
		return m.HashAccess
	}
	return 0
}

func (m *PB_MessageFileView) GetCreatedSe() int32 {
	if m != nil {
		return m.CreatedSe
	}
	return 0
}

func (m *PB_MessageFileView) GetServerSrc() string {
	if m != nil {
		return m.ServerSrc
	}
	return ""
}

func (m *PB_MessageFileView) GetServerPath() string {
	if m != nil {
		return m.ServerPath
	}
	return ""
}

func (m *PB_MessageFileView) GetServerThumbPath() string {
	if m != nil {
		return m.ServerThumbPath
	}
	return ""
}

func (m *PB_MessageFileView) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

func (m *PB_MessageFileView) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *PB_MessageFileView) GetCanDel() int32 {
	if m != nil {
		return m.CanDel
	}
	return 0
}

func (m *PB_MessageFileView) GetServerThumbLocalSrc() string {
	if m != nil {
		return m.ServerThumbLocalSrc
	}
	return ""
}

func (m *PB_MessageFileView) GetRemoteMessageFileId() int64 {
	if m != nil {
		return m.RemoteMessageFileId
	}
	return 0
}

func (m *PB_MessageFileView) GetLocalSrc() string {
	if m != nil {
		return m.LocalSrc
	}
	return ""
}

func (m *PB_MessageFileView) GetThumbLocalSrc() string {
	if m != nil {
		return m.ThumbLocalSrc
	}
	return ""
}

func (m *PB_MessageFileView) GetMessageFileStatusId() string {
	if m != nil {
		return m.MessageFileStatusId
	}
	return ""
}

// {realm} ,
type PB_UserView struct {
	UserId         int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	UserName       string `protobuf:"bytes,3,opt,name=UserName" json:"UserName,omitempty"`
	FirstName      string `protobuf:"bytes,7,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName       string `protobuf:"bytes,9,opt,name=LastName" json:"LastName,omitempty"`
	About          string `protobuf:"bytes,11,opt,name=About" json:"About,omitempty"`
	FullName       string `protobuf:"bytes,13,opt,name=FullName" json:"FullName,omitempty"`
	AvatarUrl      string `protobuf:"bytes,15,opt,name=AvatarUrl" json:"AvatarUrl,omitempty"`
	PrivacyProfile int32  `protobuf:"varint,17,opt,name=PrivacyProfile" json:"PrivacyProfile,omitempty"`
	// string Phone = 19;
	// string Email = 21;
	IsDeleted        int32 `protobuf:"varint,23,opt,name=IsDeleted" json:"IsDeleted,omitempty"`
	FollowersCount   int32 `protobuf:"varint,29,opt,name=FollowersCount" json:"FollowersCount,omitempty"`
	FollowingCount   int32 `protobuf:"varint,31,opt,name=FollowingCount" json:"FollowingCount,omitempty"`
	PostsCount       int32 `protobuf:"varint,33,opt,name=PostsCount" json:"PostsCount,omitempty"`
	UpdatedTime      int32 `protobuf:"varint,49,opt,name=UpdatedTime" json:"UpdatedTime,omitempty"`
	AppVersion       int32 `protobuf:"varint,59,opt,name=AppVersion" json:"AppVersion,omitempty"`
	LastActivityTime int32 `protobuf:"varint,61,opt,name=LastActivityTime" json:"LastActivityTime,omitempty"`
	FollowingType    int32 `protobuf:"varint,62,opt,name=FollowingType" json:"FollowingType,omitempty"`
}

func (m *PB_UserView) Reset()                    { *m = PB_UserView{} }
func (m *PB_UserView) String() string            { return proto.CompactTextString(m) }
func (*PB_UserView) ProtoMessage()               {}
func (*PB_UserView) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *PB_UserView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_UserView) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PB_UserView) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *PB_UserView) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *PB_UserView) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PB_UserView) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *PB_UserView) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *PB_UserView) GetPrivacyProfile() int32 {
	if m != nil {
		return m.PrivacyProfile
	}
	return 0
}

func (m *PB_UserView) GetIsDeleted() int32 {
	if m != nil {
		return m.IsDeleted
	}
	return 0
}

func (m *PB_UserView) GetFollowersCount() int32 {
	if m != nil {
		return m.FollowersCount
	}
	return 0
}

func (m *PB_UserView) GetFollowingCount() int32 {
	if m != nil {
		return m.FollowingCount
	}
	return 0
}

func (m *PB_UserView) GetPostsCount() int32 {
	if m != nil {
		return m.PostsCount
	}
	return 0
}

func (m *PB_UserView) GetUpdatedTime() int32 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *PB_UserView) GetAppVersion() int32 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

func (m *PB_UserView) GetLastActivityTime() int32 {
	if m != nil {
		return m.LastActivityTime
	}
	return 0
}

func (m *PB_UserView) GetFollowingType() int32 {
	if m != nil {
		return m.FollowingType
	}
	return 0
}

func init() {
	proto.RegisterType((*PB_ChatView)(nil), "PB_ChatView")
	proto.RegisterType((*PB_MessageView)(nil), "PB_MessageView")
	proto.RegisterType((*PB_MessageFileView)(nil), "PB_MessageFileView")
	proto.RegisterType((*PB_UserView)(nil), "PB_UserView")
}

func init() { proto.RegisterFile("pb_views.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 1065 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0xd6, 0xd2, 0x4b, 0xdb, 0x38, 0xd7, 0xe6, 0xea, 0x5e, 0x61, 0xe9, 0x71, 0xbd, 0xb4, 0xaa,
	0x4e, 0x11, 0xa0, 0x0a, 0x82, 0x78, 0x42, 0x87, 0x48, 0xd2, 0x2b, 0x8d, 0xb8, 0x42, 0xb4, 0x69,
	0x0e, 0x89, 0x97, 0x6a, 0x93, 0x1d, 0x1a, 0xc3, 0x66, 0x37, 0xe7, 0xf5, 0xa6, 0x0d, 0xff, 0x83,
	0x3f, 0xc1, 0x9f, 0xe0, 0x95, 0x17, 0x5e, 0x10, 0x12, 0x7f, 0x07, 0xcd, 0xd8, 0xbb, 0xeb, 0x4d,
	0xd2, 0x37, 0x7f, 0xdf, 0x8c, 0xc7, 0xf6, 0xf8, 0x9b, 0xb1, 0xd9, 0xee, 0x6c, 0x74, 0x33, 0x17,
	0x70, 0x97, 0x9c, 0xcd, 0x64, 0xac, 0xe2, 0x43, 0xc4, 0x10, 0xa5, 0x53, 0x83, 0x4f, 0xfe, 0xab,
	0xb0, 0x5a, 0xbf, 0x73, 0xd3, 0x9d, 0xf8, 0xea, 0xad, 0x80, 0x3b, 0xee, 0xb2, 0x2d, 0x1c, 0x7f,
	0x07, 0x0b, 0xd7, 0x69, 0x38, 0xcd, 0xaa, 0x97, 0x41, 0xb4, 0x78, 0x71, 0x3c, 0x45, 0xcb, 0x86,
	0xb6, 0x18, 0xc8, 0x5f, 0xb2, 0x5d, 0x1c, 0x5e, 0x2f, 0x66, 0xf0, 0x3a, 0x4a, 0xa7, 0xbd, 0xc0,
	0xad, 0x34, 0x9c, 0x66, 0xc5, 0x5b, 0x62, 0xf9, 0xfb, 0x6c, 0x73, 0x98, 0x80, 0xec, 0x05, 0xee,
	0x16, 0xd9, 0x0d, 0xe2, 0x47, 0x8c, 0xf5, 0x01, 0xa4, 0xb1, 0x55, 0xc9, 0x66, 0x31, 0xb8, 0xf2,
	0xb7, 0x32, 0x4e, 0x67, 0xbd, 0xc0, 0xad, 0x35, 0x9c, 0xe6, 0x86, 0x97, 0x41, 0xfe, 0x11, 0xab,
	0x76, 0x25, 0xf8, 0x0a, 0x82, 0x01, 0xb8, 0x3b, 0x34, 0xb1, 0x20, 0xd0, 0x3a, 0x9c, 0x05, 0x08,
	0xae, 0x12, 0xb7, 0x4e, 0x33, 0x0b, 0x82, 0x9f, 0xb2, 0x9d, 0x37, 0x7e, 0xa2, 0xae, 0x20, 0x49,
	0xfc, 0x5b, 0xe8, 0x05, 0xee, 0x1e, 0x79, 0x94, 0x49, 0xde, 0x62, 0x4f, 0x91, 0x38, 0x87, 0x10,
	0x70, 0x5a, 0xee, 0xbc, 0x4f, 0xce, 0x6b, 0x6d, 0xfc, 0x53, 0xb6, 0x87, 0xfc, 0x00, 0x20, 0x2a,
	0x26, 0x1c, 0xd0, 0x84, 0x55, 0x03, 0x6f, 0xb0, 0x9a, 0x26, 0xdf, 0x21, 0xef, 0x7e, 0x40, 0xa7,
	0xb0, 0xa9, 0x6c, 0xa7, 0x03, 0x78, 0xa7, 0x97, 0x72, 0x3f, 0x24, 0x9f, 0x32, 0x89, 0x59, 0xec,
	0xa6, 0x52, 0x42, 0x84, 0x9c, 0xfb, 0x4c, 0x67, 0xb1, 0x60, 0x78, 0x93, 0x6d, 0x63, 0x3e, 0xf1,
	0x96, 0xdd, 0xa0, 0xe1, 0x34, 0x6b, 0xad, 0xc7, 0x67, 0xfd, 0xce, 0x4d, 0xc6, 0x79, 0xb9, 0x95,
	0x7f, 0xc2, 0x9e, 0x0c, 0x26, 0xbe, 0xc4, 0x23, 0x05, 0xc2, 0xef, 0xc6, 0x69, 0xa4, 0xdc, 0xbf,
	0x1c, 0x0a, 0xb8, 0x62, 0xe0, 0xc7, 0xac, 0x36, 0x8c, 0x12, 0x80, 0x48, 0xfb, 0xfd, 0xad, 0xfd,
	0x6c, 0x8e, 0x7f, 0xc3, 0xf8, 0x85, 0x90, 0x89, 0x1a, 0x46, 0x12, 0xfc, 0x2c, 0x4f, 0xee, 0x3f,
	0x0e, 0x6d, 0xa2, 0x8e, 0x9b, 0x30, 0x1c, 0xed, 0x63, 0x8d, 0x2f, 0x6f, 0xe9, 0x1c, 0x65, 0x53,
	0xff, 0x7d, 0x60, 0xaa, 0xed, 0x74, 0xf2, 0xe7, 0x23, 0xb6, 0x5b, 0xb6, 0xa3, 0x20, 0x8a, 0x0b,
	0x71, 0xb4, 0x20, 0x8a, 0x8b, 0x38, 0x62, 0xcc, 0x80, 0x42, 0xe3, 0x16, 0x63, 0x17, 0x40, 0xa5,
	0x5c, 0x00, 0x0f, 0x09, 0xfb, 0x94, 0xed, 0x98, 0xf9, 0x17, 0x22, 0x04, 0xa3, 0xed, 0x0d, 0xaf,
	0x4c, 0xa2, 0x5c, 0x0c, 0x61, 0x55, 0x50, 0x8d, 0x02, 0xad, 0x1a, 0x38, 0x67, 0x8f, 0xae, 0xe1,
	0x5e, 0x91, 0xda, 0xab, 0x1e, 0x8d, 0xcb, 0x65, 0x50, 0x5f, 0x2e, 0x83, 0x8f, 0xd9, 0x13, 0x2c,
	0x26, 0x0f, 0xc6, 0x20, 0xe6, 0x10, 0x5c, 0x8b, 0x29, 0x90, 0xd6, 0x2b, 0xde, 0x0a, 0xcf, 0x4f,
	0xd8, 0x63, 0xe4, 0x50, 0x76, 0xe4, 0xb7, 0x4f, 0x7e, 0x25, 0x0e, 0x4b, 0xe2, 0x1c, 0x42, 0x31,
	0x17, 0x72, 0x31, 0x50, 0xbe, 0x4a, 0x13, 0xb3, 0xe5, 0x03, 0xf2, 0x5d, 0x6b, 0xb3, 0xdb, 0xca,
	0x51, 0xb9, 0xad, 0xac, 0x36, 0x8f, 0x17, 0x0f, 0x35, 0x8f, 0x5e, 0xd2, 0x59, 0x5c, 0x81, 0x7b,
	0xdc, 0x70, 0x9a, 0xdb, 0x9e, 0x41, 0xfc, 0x90, 0x6d, 0x7b, 0x30, 0x8d, 0x15, 0xa6, 0xb7, 0x45,
	0xe9, 0xcd, 0x31, 0x7f, 0xc5, 0xea, 0x56, 0xaa, 0x2d, 0xe5, 0xef, 0x5b, 0xca, 0xc9, 0x4c, 0xde,
	0xb2, 0xef, 0xc9, 0x1f, 0x9b, 0x8c, 0xaf, 0xfa, 0xad, 0xde, 0xaa, 0xb3, 0xee, 0x56, 0x5f, 0xb2,
	0x5d, 0x8b, 0xc0, 0x83, 0x37, 0xe8, 0xe0, 0x4b, 0x2c, 0xfa, 0xfd, 0x20, 0xc5, 0xad, 0x88, 0xfc,
	0xd0, 0x68, 0xe8, 0x3d, 0x7d, 0xfe, 0x32, 0x8b, 0xf7, 0xfe, 0xbd, 0x3f, 0x05, 0xa3, 0x4b, 0x1a,
	0x23, 0x37, 0x10, 0xbf, 0x81, 0x69, 0xb7, 0x34, 0xc6, 0x78, 0x18, 0xda, 0xca, 0xa7, 0xd6, 0xe4,
	0x12, 0xcb, 0x9f, 0xb2, 0xca, 0x8f, 0x22, 0x50, 0x13, 0xd3, 0x6f, 0x35, 0xc0, 0x2c, 0x5f, 0x82,
	0xb8, 0x9d, 0x28, 0x23, 0x40, 0x83, 0x30, 0xcb, 0xe7, 0xa9, 0xf4, 0x95, 0x88, 0x23, 0xd3, 0x67,
	0x73, 0x8c, 0xea, 0x7b, 0x7d, 0xaf, 0x20, 0x4a, 0xd0, 0x58, 0xa7, 0xed, 0x15, 0x04, 0xde, 0xfc,
	0xa5, 0x9f, 0x4c, 0xae, 0x82, 0x2f, 0x49, 0x74, 0x55, 0x2f, 0x83, 0x58, 0x6f, 0x38, 0x6c, 0x8f,
	0xc7, 0x90, 0x24, 0xa6, 0xa1, 0x5a, 0x4c, 0x59, 0xd5, 0x07, 0x6b, 0x9a, 0xfb, 0x00, 0xe4, 0x1c,
	0xe4, 0x40, 0x8e, 0xa9, 0x69, 0x56, 0xbd, 0x82, 0xc0, 0xd8, 0x1a, 0xf4, 0x7d, 0x35, 0xa1, 0x7e,
	0x59, 0xf5, 0x2c, 0x86, 0x37, 0x59, 0x5d, 0xa3, 0xeb, 0x49, 0x3a, 0x1d, 0x91, 0xd3, 0x33, 0x72,
	0x5a, 0xa6, 0xf1, 0xe4, 0x9d, 0x74, 0xfc, 0x2b, 0xa8, 0x5e, 0xe0, 0x3e, 0x27, 0x97, 0x1c, 0xa3,
	0x4d, 0xbb, 0xe7, 0xaa, 0xcd, 0x31, 0x66, 0xb2, 0xeb, 0x47, 0xe7, 0x10, 0x92, 0x5e, 0x2b, 0x9e,
	0x41, 0xfc, 0x33, 0xb6, 0x6f, 0x2d, 0xf1, 0x26, 0x1e, 0xfb, 0x21, 0x9e, 0xa0, 0x45, 0xa1, 0xd7,
	0x99, 0x70, 0x86, 0x56, 0x74, 0x59, 0x75, 0x01, 0x25, 0x6c, 0x9d, 0x09, 0xf7, 0x95, 0x07, 0x06,
	0xbd, 0xe7, 0x3c, 0xda, 0x29, 0xdb, 0x29, 0xaf, 0x3c, 0x21, 0x87, 0x9d, 0x95, 0x35, 0xad, 0x90,
	0xba, 0x94, 0x7b, 0x81, 0xfb, 0x8b, 0xde, 0xe5, 0x1a, 0xd3, 0xc9, 0xef, 0x8f, 0xe8, 0x23, 0x91,
	0x3f, 0x22, 0x45, 0x4f, 0x74, 0x4a, 0x3d, 0xf1, 0x50, 0x3f, 0x43, 0x96, 0x96, 0x73, 0x8c, 0x77,
	0x4a, 0xcd, 0x9f, 0x8c, 0x5b, 0xfa, 0x4e, 0x73, 0x82, 0x4e, 0xe5, 0x1b, 0x63, 0xd5, 0x9c, 0xca,
	0x60, 0x54, 0x73, 0x7b, 0x14, 0xa7, 0x5a, 0xb6, 0x55, 0x4f, 0x03, 0x9c, 0x71, 0x91, 0x86, 0x21,
	0xcd, 0xd0, 0xfd, 0x32, 0xc7, 0xb8, 0x56, 0x7b, 0xee, 0x2b, 0x5f, 0x0e, 0x65, 0x98, 0xa9, 0x36,
	0x27, 0xb0, 0x8a, 0xfa, 0x52, 0xcc, 0xfd, 0xf1, 0xa2, 0x2f, 0xe3, 0x9f, 0x45, 0x98, 0x75, 0xcc,
	0x25, 0x16, 0xa3, 0xf4, 0x12, 0xf3, 0x01, 0x30, 0x4f, 0x77, 0x41, 0x50, 0x2d, 0xc6, 0x61, 0x18,
	0xdf, 0x81, 0x4c, 0xf4, 0xf3, 0xf8, 0xdc, 0xd4, 0x62, 0x89, 0x2d, 0xfc, 0x44, 0x74, 0xab, 0xfd,
	0x5e, 0xd8, 0x7e, 0x19, 0x4b, 0x1f, 0xa5, 0x38, 0x51, 0x26, 0xd6, 0xb1, 0xf9, 0x28, 0xe5, 0x0c,
	0x7e, 0x25, 0xcc, 0xff, 0x86, 0x9a, 0xf7, 0xe7, 0xe6, 0x29, 0x2e, 0x28, 0x8c, 0xd0, 0x9e, 0xcd,
	0xde, 0x82, 0xa4, 0x62, 0xfd, 0x4a, 0x47, 0x28, 0x18, 0x7c, 0x2b, 0x30, 0xa7, 0xed, 0xb1, 0x12,
	0x73, 0xa1, 0x16, 0x14, 0xe6, 0x95, 0x7e, 0x2b, 0x96, 0x79, 0x54, 0x52, 0xbe, 0x3f, 0x6c, 0x2c,
	0xee, 0xd7, 0xfa, 0x5b, 0x52, 0x22, 0x3b, 0x7b, 0x6c, 0x5b, 0xc8, 0x33, 0xfc, 0x6f, 0x8e, 0x2e,
	0x37, 0xfa, 0xce, 0x4f, 0xce, 0xfd, 0x68, 0x93, 0xbe, 0x9e, 0x5f, 0xfc, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xaf, 0x7a, 0xeb, 0x5c, 0x9c, 0x0a, 0x00, 0x00,
}
