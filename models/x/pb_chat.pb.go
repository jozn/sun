// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_chat.proto

/*
Package x is a generated protocol buffer package.

It is generated from these files:
	pb_chat.proto
	pb_del.proto
	pb_enums.proto
	pb_global.proto
	pb_rpc_auth2.proto
	pb_rpc_msg.proto
	pb_rpc_sync.proto
	pb_rpc_user.proto
	pb_table.proto
	pb_updates.proto
	pb_views.proto

It has these top-level messages:
	PB_RoomInlineView
	PB_MessageForwardedFrom
	PB_GroupView
	PB_GroupMemberView
	PB_MessageFileView__DEp
	PB_ReqLastChangesForTheRoom
	PB_ResponseLastChangesForTheRoom
	PB_RequestSetLastSeenMessages
	PB_ResponseSetLastSeenMessages
	PB_MessagesCollections
	PB_DirectLogView
	PB_DirectLog
	PB_PushDirectLogViewsMany
	PB_UserWithMe
	PB_Message
	PB_MsgFile
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
	PB_MsgsSeenFromClient
	PB_MsgSeen
	PB_RequestMsgsSeen
	PB_ResponseMsgsSeen
	PB_ReqRpcAddMsg
	PB_ResRpcAddMsg
	GeoLocation
	RoomMessageLog
	RoomMessageForwardFrom
	RoomDraft
	ChatRoom
	Pagination
	PB_CommandToServer
	PB_CommandToClient
	PB_CommandReachedToServer
	PB_CommandReachedToClient
	PB_ResponseToClient
	PB_UserParam_CheckUserName2
	PB_UserResponse_CheckUserName2
	PB_MsgParam_AddNewTextMessage
	PB_MsgResponse_AddNewTextMessage
	PB_MsgParam_AddNewMessage
	PB_MsgResponse_AddNewMessage
	PB_MsgParam_SetRoomActionDoing
	PB_MsgResponse_SetRoomActionDoing
	PB_MsgParam_GetMessagesByIds
	PB_MsgResponse_GetMessagesByIds
	PB_MsgParam_GetMessagesHistory
	PB_MsgResponse_GetMessagesHistory
	PB_MsgParam_SetChatMessagesRangeAsSeen
	PB_MsgResponse_SetChatMessagesRangeAsSeen
	PB_MsgParam_DeleteChatHistory
	PB_MsgResponse_DeleteChatHistory
	PB_MsgParam_DeleteMessagesByIds
	PB_MsgResponse_DeleteMessagesByIds
	PB_MsgParam_SetMessagesAsReceived
	PB_MsgResponse_SetMessagesAsReceived
	PB_MsgParam_ForwardMessages
	PB_MsgResponse_ForwardMessages
	PB_MsgParam_EditMessage
	PB_MsgResponse_EditMessage
	PB_MsgParam_BroadcastNewMessage
	PB_MsgResponse_BroadcastNewMessage
	PB_MsgParam_GetFreshChatList
	PB_MsgResponse_GetFreshChatList
	PB_MsgParam_GetFreshRoomMessagesList
	PB_MsgResponse_GetFreshRoomMessagesList
	PB_MsgParam_GetFreshAllDirectMessagesList
	PB_MsgResponse_GetFreshAllDirectMessagesList
	PB_MsgParam_Echo
	PB_MsgResponse_PB_MsgParam_Echo
	PB_SyncParam_GetDirectUpdates
	PB_SyncResponse_GetDirectUpdates
	PB_SyncParam_GetGeneralUpdates
	PB_SyncResponse_GetGeneralUpdates
	PB_SyncParam_GetNotifyUpdates
	PB_SyncResponse_GetNotifyUpdates
	PB_SyncParam_SetLastSyncDirectUpdateId
	PB_SyncResponse_SetLastSyncDirectUpdateId
	PB_SyncParam_SetLastSyncGeneralUpdateId
	PB_SyncResponse_SetLastSyncGeneralUpdateId
	PB_SyncParam_SetLastSyncNotifyUpdateId
	PB_SyncResponse_SetLastSyncNotifyUpdateId
	PB_UserParam_BlockUser
	PB_UserOfflineResponse_BlockUser
	PB_UserParam_UnBlockUser
	PB_UserOfflineResponse_UnBlockUser
	PB_UserParam_BlockedList
	PB_UserResponse_BlockedList
	PB_UserParam_UpdateAbout
	PB_UserOfflineResponse_UpdateAbout
	PB_UserParam_UpdateUserName
	PB_UserOfflineResponse_UpdateUserName
	PB_UserParam_ChangeAvatar
	PB_UserOfflineResponse_ChangeAvatar
	PB_UserParam_ChangePrivacy
	PB_UserResponseOffline_ChangePrivacy
	PB_UserParam_CheckUserName
	PB_UserResponse_CheckUserName
	UserView
	PB_Chat
	PB_DirectMessage
	PB_MessageFile
	PB_User
	PB_UpdateGroupParticipants
	PB_UpdateNotifySettings
	PB_UpdateServiceNotification
	PB_UpdateMessageMeta
	PB_UpdateMessageId
	PB_UpdateMessageToEdit
	PB_UpdateMessageToDelete
	PB_UpdateRoomActionDoing
	PB_UpdateUserBlocked
	PB_PushHolderView
	PB_DirectUpdatesView
	PB_GeneralUpdatesView
	PB_NotifyUpdatesView
	PB_ChatView
	PB_MessageView
	PB_MessageFileView
	PB_UserView
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoomLogTypeEnum int32

const (
	RoomLogTypeEnum_UNKNOWN_ROOM_UPDATE_LOG     RoomLogTypeEnum = 0
	RoomLogTypeEnum_NEW_DIRECT_MESSAGE          RoomLogTypeEnum = 1
	RoomLogTypeEnum_CHANGE_MESSAGE_ID           RoomLogTypeEnum = 2
	RoomLogTypeEnum_MESSAGE_RECIVED_TO_SERVER   RoomLogTypeEnum = 3
	RoomLogTypeEnum_MESSAGE_DELIVIERD_TO_PEER   RoomLogTypeEnum = 4
	RoomLogTypeEnum_MESSAGE_SEEN_BY_PEER        RoomLogTypeEnum = 5
	RoomLogTypeEnum_MESSAGE_DELETED_FROM_SERVER RoomLogTypeEnum = 6
	RoomLogTypeEnum_MESSAGE_UPDATE_BY_USER      RoomLogTypeEnum = 7
	RoomLogTypeEnum_MESSAGE_DELETE_BY_USER      RoomLogTypeEnum = 8
	RoomLogTypeEnum_ROOM_ACTION_DOING           RoomLogTypeEnum = 10
)

var RoomLogTypeEnum_name = map[int32]string{
	0:  "UNKNOWN_ROOM_UPDATE_LOG",
	1:  "NEW_DIRECT_MESSAGE",
	2:  "CHANGE_MESSAGE_ID",
	3:  "MESSAGE_RECIVED_TO_SERVER",
	4:  "MESSAGE_DELIVIERD_TO_PEER",
	5:  "MESSAGE_SEEN_BY_PEER",
	6:  "MESSAGE_DELETED_FROM_SERVER",
	7:  "MESSAGE_UPDATE_BY_USER",
	8:  "MESSAGE_DELETE_BY_USER",
	10: "ROOM_ACTION_DOING",
}
var RoomLogTypeEnum_value = map[string]int32{
	"UNKNOWN_ROOM_UPDATE_LOG":     0,
	"NEW_DIRECT_MESSAGE":          1,
	"CHANGE_MESSAGE_ID":           2,
	"MESSAGE_RECIVED_TO_SERVER":   3,
	"MESSAGE_DELIVIERD_TO_PEER":   4,
	"MESSAGE_SEEN_BY_PEER":        5,
	"MESSAGE_DELETED_FROM_SERVER": 6,
	"MESSAGE_UPDATE_BY_USER":      7,
	"MESSAGE_DELETE_BY_USER":      8,
	"ROOM_ACTION_DOING":           10,
}

func (x RoomLogTypeEnum) String() string {
	return proto.EnumName(RoomLogTypeEnum_name, int32(x))
}
func (RoomLogTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PB_RoomInlineView struct {
	RoomId       int64        `protobuf:"varint,1,opt,name=RoomId" json:"RoomId,omitempty"`
	RoomTypeEnum RoomTypeEnum `protobuf:"varint,2,opt,name=RoomTypeEnum,enum=RoomTypeEnum" json:"RoomTypeEnum,omitempty"`
}

func (m *PB_RoomInlineView) Reset()                    { *m = PB_RoomInlineView{} }
func (m *PB_RoomInlineView) String() string            { return proto.CompactTextString(m) }
func (*PB_RoomInlineView) ProtoMessage()               {}
func (*PB_RoomInlineView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PB_RoomInlineView) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PB_RoomInlineView) GetRoomTypeEnum() RoomTypeEnum {
	if m != nil {
		return m.RoomTypeEnum
	}
	return RoomTypeEnum_UNKNOWN_ROOM_TYPE
}

type PB_MessageForwardedFrom struct {
	RoomId       int64        `protobuf:"varint,1,opt,name=RoomId" json:"RoomId,omitempty"`
	RoomTypeEnum RoomTypeEnum `protobuf:"varint,2,opt,name=RoomTypeEnum,enum=RoomTypeEnum" json:"RoomTypeEnum,omitempty"`
	MessageId    int64        `protobuf:"varint,3,opt,name=MessageId" json:"MessageId,omitempty"`
	MessageSeq   int32        `protobuf:"varint,4,opt,name=MessageSeq" json:"MessageSeq,omitempty"`
}

func (m *PB_MessageForwardedFrom) Reset()                    { *m = PB_MessageForwardedFrom{} }
func (m *PB_MessageForwardedFrom) String() string            { return proto.CompactTextString(m) }
func (*PB_MessageForwardedFrom) ProtoMessage()               {}
func (*PB_MessageForwardedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PB_MessageForwardedFrom) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PB_MessageForwardedFrom) GetRoomTypeEnum() RoomTypeEnum {
	if m != nil {
		return m.RoomTypeEnum
	}
	return RoomTypeEnum_UNKNOWN_ROOM_TYPE
}

func (m *PB_MessageForwardedFrom) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_MessageForwardedFrom) GetMessageSeq() int32 {
	if m != nil {
		return m.MessageSeq
	}
	return 0
}

type PB_GroupView struct {
	GroupId          int64  `protobuf:"varint,1,opt,name=GroupId" json:"GroupId,omitempty"`
	GroupName        string `protobuf:"bytes,3,opt,name=GroupName" json:"GroupName,omitempty"`
	MembersCount     int32  `protobuf:"varint,5,opt,name=MembersCount" json:"MembersCount,omitempty"`
	GroupPrivacyEnum int32  `protobuf:"varint,7,opt,name=GroupPrivacyEnum" json:"GroupPrivacyEnum,omitempty"`
	CreatorUserId    int32  `protobuf:"varint,9,opt,name=CreatorUserId" json:"CreatorUserId,omitempty"`
	CreatedTime      int32  `protobuf:"varint,11,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	UpdatedMs        int64  `protobuf:"varint,13,opt,name=UpdatedMs" json:"UpdatedMs,omitempty"`
	CurrentSeq       int32  `protobuf:"varint,15,opt,name=CurrentSeq" json:"CurrentSeq,omitempty"`
}

func (m *PB_GroupView) Reset()                    { *m = PB_GroupView{} }
func (m *PB_GroupView) String() string            { return proto.CompactTextString(m) }
func (*PB_GroupView) ProtoMessage()               {}
func (*PB_GroupView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PB_GroupView) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_GroupView) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *PB_GroupView) GetMembersCount() int32 {
	if m != nil {
		return m.MembersCount
	}
	return 0
}

func (m *PB_GroupView) GetGroupPrivacyEnum() int32 {
	if m != nil {
		return m.GroupPrivacyEnum
	}
	return 0
}

func (m *PB_GroupView) GetCreatorUserId() int32 {
	if m != nil {
		return m.CreatorUserId
	}
	return 0
}

func (m *PB_GroupView) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *PB_GroupView) GetUpdatedMs() int64 {
	if m != nil {
		return m.UpdatedMs
	}
	return 0
}

func (m *PB_GroupView) GetCurrentSeq() int32 {
	if m != nil {
		return m.CurrentSeq
	}
	return 0
}

type PB_GroupMemberView struct {
	Id            int64 `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	GroupId       int64 `protobuf:"varint,2,opt,name=GroupId" json:"GroupId,omitempty"`
	UserId        int32 `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	ByUserId      int32 `protobuf:"varint,4,opt,name=ByUserId" json:"ByUserId,omitempty"`
	GroupRoleEnum int32 `protobuf:"varint,5,opt,name=GroupRoleEnum" json:"GroupRoleEnum,omitempty"`
	CreatedTime   int32 `protobuf:"varint,6,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
}

func (m *PB_GroupMemberView) Reset()                    { *m = PB_GroupMemberView{} }
func (m *PB_GroupMemberView) String() string            { return proto.CompactTextString(m) }
func (*PB_GroupMemberView) ProtoMessage()               {}
func (*PB_GroupMemberView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PB_GroupMemberView) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PB_GroupMemberView) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_GroupMemberView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_GroupMemberView) GetByUserId() int32 {
	if m != nil {
		return m.ByUserId
	}
	return 0
}

func (m *PB_GroupMemberView) GetGroupRoleEnum() int32 {
	if m != nil {
		return m.GroupRoleEnum
	}
	return 0
}

func (m *PB_GroupMemberView) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

type PB_MessageFileView__DEp struct {
	MessageFileId int64  `protobuf:"varint,1,opt,name=MessageFileId" json:"MessageFileId,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Size          int32  `protobuf:"varint,5,opt,name=Size" json:"Size,omitempty"`
	FileTypeEnum  int32  `protobuf:"varint,7,opt,name=FileTypeEnum" json:"FileTypeEnum,omitempty"`
	MimeType      string `protobuf:"bytes,9,opt,name=MimeType" json:"MimeType,omitempty"`
	Width         int32  `protobuf:"varint,11,opt,name=Width" json:"Width,omitempty"`
	Height        int32  `protobuf:"varint,13,opt,name=Height" json:"Height,omitempty"`
	Duration      int32  `protobuf:"varint,15,opt,name=Duration" json:"Duration,omitempty"`
	Extension     string `protobuf:"bytes,17,opt,name=Extension" json:"Extension,omitempty"`
	//    ???? ThumbData = 19;
	ThumbData64     string `protobuf:"bytes,21,opt,name=ThumbData64" json:"ThumbData64,omitempty"`
	ServerSrc       string `protobuf:"bytes,23,opt,name=ServerSrc" json:"ServerSrc,omitempty"`
	ServerPath      string `protobuf:"bytes,25,opt,name=ServerPath" json:"ServerPath,omitempty"`
	ServerThumbPath string `protobuf:"bytes,27,opt,name=ServerThumbPath" json:"ServerThumbPath,omitempty"`
	BucketId        string `protobuf:"bytes,29,opt,name=BucketId" json:"BucketId,omitempty"`
	ServerId        int32  `protobuf:"varint,31,opt,name=ServerId" json:"ServerId,omitempty"`
	CanDel          int32  `protobuf:"varint,33,opt,name=CanDel" json:"CanDel,omitempty"`
	CreatedTime     int32  `protobuf:"varint,35,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
}

func (m *PB_MessageFileView__DEp) Reset()                    { *m = PB_MessageFileView__DEp{} }
func (m *PB_MessageFileView__DEp) String() string            { return proto.CompactTextString(m) }
func (*PB_MessageFileView__DEp) ProtoMessage()               {}
func (*PB_MessageFileView__DEp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PB_MessageFileView__DEp) GetMessageFileId() int64 {
	if m != nil {
		return m.MessageFileId
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetFileTypeEnum() int32 {
	if m != nil {
		return m.FileTypeEnum
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetThumbData64() string {
	if m != nil {
		return m.ThumbData64
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetServerSrc() string {
	if m != nil {
		return m.ServerSrc
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetServerPath() string {
	if m != nil {
		return m.ServerPath
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetServerThumbPath() string {
	if m != nil {
		return m.ServerThumbPath
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

func (m *PB_MessageFileView__DEp) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetCanDel() int32 {
	if m != nil {
		return m.CanDel
	}
	return 0
}

func (m *PB_MessageFileView__DEp) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

// //
type PB_ReqLastChangesForTheRoom struct {
	RoomId      int64 `protobuf:"varint,1,opt,name=RoomId" json:"RoomId,omitempty"`
	LastLogId   int64 `protobuf:"varint,2,opt,name=LastLogId" json:"LastLogId,omitempty"`
	LastHaveSeq int32 `protobuf:"varint,3,opt,name=LastHaveSeq" json:"LastHaveSeq,omitempty"`
}

func (m *PB_ReqLastChangesForTheRoom) Reset()                    { *m = PB_ReqLastChangesForTheRoom{} }
func (m *PB_ReqLastChangesForTheRoom) String() string            { return proto.CompactTextString(m) }
func (*PB_ReqLastChangesForTheRoom) ProtoMessage()               {}
func (*PB_ReqLastChangesForTheRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PB_ReqLastChangesForTheRoom) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PB_ReqLastChangesForTheRoom) GetLastLogId() int64 {
	if m != nil {
		return m.LastLogId
	}
	return 0
}

func (m *PB_ReqLastChangesForTheRoom) GetLastHaveSeq() int32 {
	if m != nil {
		return m.LastHaveSeq
	}
	return 0
}

type PB_ResponseLastChangesForTheRoom struct {
	Messages []*PB_MessageView `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *PB_ResponseLastChangesForTheRoom) Reset()         { *m = PB_ResponseLastChangesForTheRoom{} }
func (m *PB_ResponseLastChangesForTheRoom) String() string { return proto.CompactTextString(m) }
func (*PB_ResponseLastChangesForTheRoom) ProtoMessage()    {}
func (*PB_ResponseLastChangesForTheRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

func (m *PB_ResponseLastChangesForTheRoom) GetMessages() []*PB_MessageView {
	if m != nil {
		return m.Messages
	}
	return nil
}

// //
type PB_RequestSetLastSeenMessages struct {
	RoomId        int64 `protobuf:"varint,1,opt,name=RoomId" json:"RoomId,omitempty"`
	FromMessageId int64 `protobuf:"varint,2,opt,name=FromMessageId" json:"FromMessageId,omitempty"`
	ToMessageId   int32 `protobuf:"varint,3,opt,name=ToMessageId" json:"ToMessageId,omitempty"`
	AtTimeMs      int64 `protobuf:"varint,4,opt,name=AtTimeMs" json:"AtTimeMs,omitempty"`
}

func (m *PB_RequestSetLastSeenMessages) Reset()                    { *m = PB_RequestSetLastSeenMessages{} }
func (m *PB_RequestSetLastSeenMessages) String() string            { return proto.CompactTextString(m) }
func (*PB_RequestSetLastSeenMessages) ProtoMessage()               {}
func (*PB_RequestSetLastSeenMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PB_RequestSetLastSeenMessages) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PB_RequestSetLastSeenMessages) GetFromMessageId() int64 {
	if m != nil {
		return m.FromMessageId
	}
	return 0
}

func (m *PB_RequestSetLastSeenMessages) GetToMessageId() int32 {
	if m != nil {
		return m.ToMessageId
	}
	return 0
}

func (m *PB_RequestSetLastSeenMessages) GetAtTimeMs() int64 {
	if m != nil {
		return m.AtTimeMs
	}
	return 0
}

type PB_ResponseSetLastSeenMessages struct {
	Messages []*PB_MessageView `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *PB_ResponseSetLastSeenMessages) Reset()                    { *m = PB_ResponseSetLastSeenMessages{} }
func (m *PB_ResponseSetLastSeenMessages) String() string            { return proto.CompactTextString(m) }
func (*PB_ResponseSetLastSeenMessages) ProtoMessage()               {}
func (*PB_ResponseSetLastSeenMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PB_ResponseSetLastSeenMessages) GetMessages() []*PB_MessageView {
	if m != nil {
		return m.Messages
	}
	return nil
}

type PB_MessagesCollections struct {
	DirectMessagesIds   []int64 `protobuf:"varint,10,rep,packed,name=DirectMessagesIds" json:"DirectMessagesIds,omitempty"`
	GroupMessagesIds    []int64 `protobuf:"varint,20,rep,packed,name=GroupMessagesIds" json:"GroupMessagesIds,omitempty"`
	BroadCatMessagesIds []int64 `protobuf:"varint,21,rep,packed,name=BroadCatMessagesIds" json:"BroadCatMessagesIds,omitempty"`
}

func (m *PB_MessagesCollections) Reset()                    { *m = PB_MessagesCollections{} }
func (m *PB_MessagesCollections) String() string            { return proto.CompactTextString(m) }
func (*PB_MessagesCollections) ProtoMessage()               {}
func (*PB_MessagesCollections) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PB_MessagesCollections) GetDirectMessagesIds() []int64 {
	if m != nil {
		return m.DirectMessagesIds
	}
	return nil
}

func (m *PB_MessagesCollections) GetGroupMessagesIds() []int64 {
	if m != nil {
		return m.GroupMessagesIds
	}
	return nil
}

func (m *PB_MessagesCollections) GetBroadCatMessagesIds() []int64 {
	if m != nil {
		return m.BroadCatMessagesIds
	}
	return nil
}

type PB_DirectLogView struct {
	Row        *PB_DirectLog   `protobuf:"bytes,1,opt,name=Row" json:"Row,omitempty"`
	NewMessage *PB_MessageView `protobuf:"bytes,2,opt,name=NewMessage" json:"NewMessage,omitempty"`
}

func (m *PB_DirectLogView) Reset()                    { *m = PB_DirectLogView{} }
func (m *PB_DirectLogView) String() string            { return proto.CompactTextString(m) }
func (*PB_DirectLogView) ProtoMessage()               {}
func (*PB_DirectLogView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PB_DirectLogView) GetRow() *PB_DirectLog {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *PB_DirectLogView) GetNewMessage() *PB_MessageView {
	if m != nil {
		return m.NewMessage
	}
	return nil
}

type PB_DirectLog struct {
	Id            int64  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	ToUserId      int32  `protobuf:"varint,3,opt,name=ToUserId" json:"ToUserId,omitempty"`
	MessageId     int64  `protobuf:"varint,5,opt,name=MessageId" json:"MessageId,omitempty"`
	ChatId        int64  `protobuf:"varint,7,opt,name=ChatId" json:"ChatId,omitempty"`
	PeerUserId    int32  `protobuf:"varint,9,opt,name=PeerUserId" json:"PeerUserId,omitempty"`
	EventType     int32  `protobuf:"varint,11,opt,name=EventType" json:"EventType,omitempty"`
	RoomLogTypeId int32  `protobuf:"varint,13,opt,name=RoomLogTypeId" json:"RoomLogTypeId,omitempty"`
	FromSeq       int32  `protobuf:"varint,15,opt,name=FromSeq" json:"FromSeq,omitempty"`
	ToSeq         int32  `protobuf:"varint,17,opt,name=ToSeq" json:"ToSeq,omitempty"`
	ExtraPB       []byte `protobuf:"bytes,19,opt,name=ExtraPB,proto3" json:"ExtraPB,omitempty"`
	ExtraJson     string `protobuf:"bytes,21,opt,name=ExtraJson" json:"ExtraJson,omitempty"`
	AtTimeMs      int64  `protobuf:"varint,23,opt,name=AtTimeMs" json:"AtTimeMs,omitempty"`
}

func (m *PB_DirectLog) Reset()                    { *m = PB_DirectLog{} }
func (m *PB_DirectLog) String() string            { return proto.CompactTextString(m) }
func (*PB_DirectLog) ProtoMessage()               {}
func (*PB_DirectLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PB_DirectLog) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PB_DirectLog) GetToUserId() int32 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *PB_DirectLog) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_DirectLog) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *PB_DirectLog) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *PB_DirectLog) GetEventType() int32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

func (m *PB_DirectLog) GetRoomLogTypeId() int32 {
	if m != nil {
		return m.RoomLogTypeId
	}
	return 0
}

func (m *PB_DirectLog) GetFromSeq() int32 {
	if m != nil {
		return m.FromSeq
	}
	return 0
}

func (m *PB_DirectLog) GetToSeq() int32 {
	if m != nil {
		return m.ToSeq
	}
	return 0
}

func (m *PB_DirectLog) GetExtraPB() []byte {
	if m != nil {
		return m.ExtraPB
	}
	return nil
}

func (m *PB_DirectLog) GetExtraJson() string {
	if m != nil {
		return m.ExtraJson
	}
	return ""
}

func (m *PB_DirectLog) GetAtTimeMs() int64 {
	if m != nil {
		return m.AtTimeMs
	}
	return 0
}

type PB_PushDirectLogViewsMany struct {
	Rows []*PB_DirectLogView `protobuf:"bytes,1,rep,name=Rows" json:"Rows,omitempty"`
}

func (m *PB_PushDirectLogViewsMany) Reset()                    { *m = PB_PushDirectLogViewsMany{} }
func (m *PB_PushDirectLogViewsMany) String() string            { return proto.CompactTextString(m) }
func (*PB_PushDirectLogViewsMany) ProtoMessage()               {}
func (*PB_PushDirectLogViewsMany) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PB_PushDirectLogViewsMany) GetRows() []*PB_DirectLogView {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_RoomInlineView)(nil), "PB_RoomInlineView")
	proto.RegisterType((*PB_MessageForwardedFrom)(nil), "PB_MessageForwardedFrom")
	proto.RegisterType((*PB_GroupView)(nil), "PB_GroupView")
	proto.RegisterType((*PB_GroupMemberView)(nil), "PB_GroupMemberView")
	proto.RegisterType((*PB_MessageFileView__DEp)(nil), "PB_MessageFileView__DEp")
	proto.RegisterType((*PB_ReqLastChangesForTheRoom)(nil), "PB_ReqLastChangesForTheRoom")
	proto.RegisterType((*PB_ResponseLastChangesForTheRoom)(nil), "PB_ResponseLastChangesForTheRoom")
	proto.RegisterType((*PB_RequestSetLastSeenMessages)(nil), "PB_RequestSetLastSeenMessages")
	proto.RegisterType((*PB_ResponseSetLastSeenMessages)(nil), "PB_ResponseSetLastSeenMessages")
	proto.RegisterType((*PB_MessagesCollections)(nil), "PB_MessagesCollections")
	proto.RegisterType((*PB_DirectLogView)(nil), "PB_DirectLogView")
	proto.RegisterType((*PB_DirectLog)(nil), "PB_DirectLog")
	proto.RegisterType((*PB_PushDirectLogViewsMany)(nil), "PB_PushDirectLogViewsMany")
	proto.RegisterEnum("RoomLogTypeEnum", RoomLogTypeEnum_name, RoomLogTypeEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPC_MessageReq service

type RPC_MessageReqClient interface {
	GetLastChnagesForRoom(ctx context.Context, in *PB_ReqLastChangesForTheRoom, opts ...grpc.CallOption) (*PB_ResponseLastChangesForTheRoom, error)
}

type rPC_MessageReqClient struct {
	cc *grpc.ClientConn
}

func NewRPC_MessageReqClient(cc *grpc.ClientConn) RPC_MessageReqClient {
	return &rPC_MessageReqClient{cc}
}

func (c *rPC_MessageReqClient) GetLastChnagesForRoom(ctx context.Context, in *PB_ReqLastChangesForTheRoom, opts ...grpc.CallOption) (*PB_ResponseLastChangesForTheRoom, error) {
	out := new(PB_ResponseLastChangesForTheRoom)
	err := grpc.Invoke(ctx, "/RPC_MessageReq/GetLastChnagesForRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC_MessageReq service

type RPC_MessageReqServer interface {
	GetLastChnagesForRoom(context.Context, *PB_ReqLastChangesForTheRoom) (*PB_ResponseLastChangesForTheRoom, error)
}

func RegisterRPC_MessageReqServer(s *grpc.Server, srv RPC_MessageReqServer) {
	s.RegisterService(&_RPC_MessageReq_serviceDesc, srv)
}

func _RPC_MessageReq_GetLastChnagesForRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_ReqLastChangesForTheRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MessageReqServer).GetLastChnagesForRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_MessageReq/GetLastChnagesForRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MessageReqServer).GetLastChnagesForRoom(ctx, req.(*PB_ReqLastChangesForTheRoom))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_MessageReq_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPC_MessageReq",
	HandlerType: (*RPC_MessageReqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastChnagesForRoom",
			Handler:    _RPC_MessageReq_GetLastChnagesForRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_chat.proto",
}

// Client API for RPC_MessageReqOffline service

type RPC_MessageReqOfflineClient interface {
	SetLastSeen(ctx context.Context, in *PB_RequestSetLastSeenMessages, opts ...grpc.CallOption) (*PB_ResponseSetLastSeenMessages, error)
}

type rPC_MessageReqOfflineClient struct {
	cc *grpc.ClientConn
}

func NewRPC_MessageReqOfflineClient(cc *grpc.ClientConn) RPC_MessageReqOfflineClient {
	return &rPC_MessageReqOfflineClient{cc}
}

func (c *rPC_MessageReqOfflineClient) SetLastSeen(ctx context.Context, in *PB_RequestSetLastSeenMessages, opts ...grpc.CallOption) (*PB_ResponseSetLastSeenMessages, error) {
	out := new(PB_ResponseSetLastSeenMessages)
	err := grpc.Invoke(ctx, "/RPC_MessageReqOffline/SetLastSeen", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC_MessageReqOffline service

type RPC_MessageReqOfflineServer interface {
	SetLastSeen(context.Context, *PB_RequestSetLastSeenMessages) (*PB_ResponseSetLastSeenMessages, error)
}

func RegisterRPC_MessageReqOfflineServer(s *grpc.Server, srv RPC_MessageReqOfflineServer) {
	s.RegisterService(&_RPC_MessageReqOffline_serviceDesc, srv)
}

func _RPC_MessageReqOffline_SetLastSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_RequestSetLastSeenMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_MessageReqOfflineServer).SetLastSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_MessageReqOffline/SetLastSeen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_MessageReqOfflineServer).SetLastSeen(ctx, req.(*PB_RequestSetLastSeenMessages))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_MessageReqOffline_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPC_MessageReqOffline",
	HandlerType: (*RPC_MessageReqOfflineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLastSeen",
			Handler:    _RPC_MessageReqOffline_SetLastSeen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_chat.proto",
}

func init() { proto.RegisterFile("pb_chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x72, 0xe3, 0x44,
	0x13, 0xfd, 0x64, 0xc7, 0xf9, 0xe9, 0xc4, 0x89, 0x3d, 0xbb, 0x49, 0xb4, 0xce, 0xee, 0xc6, 0xeb,
	0x6f, 0xa9, 0x72, 0x2d, 0x94, 0x81, 0x40, 0x71, 0x1f, 0xfd, 0xc4, 0x11, 0xf8, 0x47, 0x35, 0x76,
	0xb2, 0x2c, 0x17, 0xa8, 0x64, 0x6b, 0xd6, 0x16, 0xd8, 0x1a, 0x47, 0x92, 0xe3, 0x84, 0x57, 0xe0,
	0x11, 0xe0, 0x8a, 0x82, 0x97, 0xe0, 0x49, 0x78, 0x1c, 0x6a, 0x66, 0x24, 0x59, 0xb2, 0xb3, 0xa1,
	0xa8, 0xe2, 0x4e, 0xe7, 0x74, 0xcf, 0x8c, 0xba, 0xfb, 0xe8, 0x8c, 0xa0, 0x38, 0x1b, 0x58, 0xc3,
	0xb1, 0x1d, 0x36, 0x66, 0x3e, 0x0d, 0x69, 0x65, 0x7f, 0x36, 0xb0, 0x88, 0x37, 0x9f, 0x06, 0x29,
	0x7c, 0xeb, 0x92, 0x45, 0x84, 0x6b, 0xdf, 0x43, 0xd9, 0x54, 0x2c, 0x4c, 0xe9, 0xd4, 0xf0, 0x26,
	0xae, 0x47, 0xae, 0x5d, 0xb2, 0x40, 0x47, 0xb0, 0xc9, 0x19, 0x47, 0x96, 0xaa, 0x52, 0x3d, 0x8f,
	0x23, 0x84, 0x3e, 0x87, 0x3d, 0xf6, 0xd4, 0xbf, 0x9f, 0x11, 0xdd, 0x9b, 0x4f, 0xe5, 0x5c, 0x55,
	0xaa, 0xef, 0x9f, 0x15, 0x1b, 0x69, 0x12, 0x67, 0x52, 0x6a, 0xbf, 0x49, 0x70, 0x6c, 0x2a, 0x56,
	0x9b, 0x04, 0x81, 0x3d, 0x22, 0x17, 0xd4, 0x5f, 0xd8, 0xbe, 0x43, 0x9c, 0x0b, 0x9f, 0x4e, 0xff,
	0xc3, 0x63, 0xd0, 0x73, 0xd8, 0x89, 0x8e, 0x30, 0x1c, 0x39, 0xcf, 0x77, 0x5b, 0x12, 0xe8, 0x25,
	0x40, 0x04, 0x7a, 0xe4, 0x46, 0xde, 0xa8, 0x4a, 0xf5, 0x02, 0x4e, 0x31, 0xb5, 0x5f, 0x73, 0xb0,
	0x67, 0x2a, 0x56, 0xd3, 0xa7, 0xf3, 0x19, 0x6f, 0x80, 0x0c, 0x5b, 0x1c, 0x24, 0xaf, 0x16, 0x43,
	0x76, 0x10, 0x7f, 0xec, 0xd8, 0x53, 0xc2, 0x0f, 0xda, 0xc1, 0x4b, 0x02, 0xd5, 0x60, 0xaf, 0x4d,
	0xa6, 0x03, 0xe2, 0x07, 0x2a, 0x9d, 0x7b, 0xa1, 0x5c, 0xe0, 0x47, 0x65, 0x38, 0xf4, 0x06, 0x4a,
	0x7c, 0x81, 0xe9, 0xbb, 0xb7, 0xf6, 0xf0, 0x9e, 0x57, 0xb8, 0xc5, 0xf3, 0xd6, 0x78, 0xf4, 0x1a,
	0x8a, 0xaa, 0x4f, 0xec, 0x90, 0xfa, 0x57, 0x01, 0xf1, 0x0d, 0x47, 0xde, 0xe1, 0x89, 0x59, 0x12,
	0x55, 0x61, 0x97, 0x13, 0xc4, 0xe9, 0xbb, 0x53, 0x22, 0xef, 0xf2, 0x9c, 0x34, 0xc5, 0xde, 0xfa,
	0x6a, 0xe6, 0x30, 0xd8, 0x0e, 0xe4, 0xa2, 0x68, 0x4f, 0x42, 0xb0, 0xf6, 0xa8, 0x73, 0xdf, 0x27,
	0x5e, 0xc8, 0xda, 0x73, 0x20, 0xda, 0xb3, 0x64, 0x6a, 0x7f, 0x4a, 0x80, 0xe2, 0xf6, 0x88, 0x52,
	0x78, 0x93, 0xf6, 0x21, 0x97, 0xf4, 0x27, 0x67, 0x38, 0xe9, 0xa6, 0xe5, 0xb2, 0x4d, 0x3b, 0x82,
	0xcd, 0xe8, 0xfd, 0xf3, 0x7c, 0xf3, 0x08, 0xa1, 0x0a, 0x6c, 0x2b, 0xf7, 0x51, 0x44, 0x4c, 0x25,
	0xc1, 0xac, 0x74, 0xbe, 0x1c, 0xd3, 0x89, 0x50, 0x81, 0xe8, 0x65, 0x96, 0x5c, 0x2d, 0x7d, 0x73,
	0xad, 0xf4, 0xda, 0xcf, 0x1b, 0x19, 0x01, 0xba, 0x13, 0x2e, 0x71, 0xcb, 0xd2, 0xf4, 0x19, 0x3b,
	0x23, 0xc5, 0x27, 0xc5, 0x64, 0x49, 0x84, 0x60, 0x23, 0x35, 0x6d, 0xfe, 0xcc, 0xb8, 0x9e, 0xfb,
	0x13, 0x89, 0x5e, 0x8a, 0x3f, 0xb3, 0xe1, 0xb3, 0x15, 0x89, 0x6c, 0xc5, 0x50, 0x33, 0x1c, 0xab,
	0xb8, 0xed, 0x4e, 0x39, 0xe6, 0xb3, 0xdc, 0xc1, 0x09, 0x46, 0x4f, 0xa1, 0xf0, 0xd6, 0x75, 0xc2,
	0x71, 0x34, 0x40, 0x01, 0x58, 0xef, 0x2e, 0x89, 0x3b, 0x1a, 0x87, 0x7c, 0x6e, 0x05, 0x1c, 0x21,
	0xb6, 0x93, 0x36, 0xf7, 0xed, 0xd0, 0xa5, 0x5e, 0x34, 0xb2, 0x04, 0xb3, 0x71, 0xeb, 0x77, 0x21,
	0xf1, 0x02, 0x16, 0x2c, 0x0b, 0x91, 0x26, 0x04, 0xeb, 0x59, 0x7f, 0x3c, 0x9f, 0x0e, 0x34, 0x3b,
	0xb4, 0xbf, 0xfa, 0x52, 0x3e, 0xe4, 0xf1, 0x34, 0xc5, 0xd6, 0xf7, 0x88, 0x7f, 0x4b, 0xfc, 0x9e,
	0x3f, 0x94, 0x8f, 0xc5, 0xfa, 0x84, 0x60, 0x72, 0x11, 0xc0, 0xb4, 0xc3, 0xb1, 0xfc, 0x8c, 0x87,
	0x53, 0x0c, 0xaa, 0xc3, 0x81, 0x40, 0x7c, 0x4b, 0x9e, 0x74, 0xc2, 0x93, 0x56, 0x69, 0x3e, 0xff,
	0xf9, 0xf0, 0x47, 0x12, 0x1a, 0x8e, 0xfc, 0x42, 0x74, 0x23, 0xc6, 0x2c, 0x26, 0xd2, 0x0d, 0x47,
	0x3e, 0x15, 0xf5, 0xc5, 0x98, 0xf5, 0x44, 0xb5, 0x3d, 0x8d, 0x4c, 0xe4, 0x57, 0xa2, 0x27, 0x02,
	0xad, 0xaa, 0xe1, 0xff, 0xeb, 0x6a, 0x98, 0xc3, 0x09, 0xb3, 0x3b, 0x72, 0xd3, 0xb2, 0x83, 0x50,
	0x1d, 0xdb, 0xde, 0x88, 0x04, 0x17, 0xd4, 0xef, 0x8f, 0x09, 0xb3, 0x93, 0x0f, 0x3a, 0xd2, 0x73,
	0xd8, 0x61, 0x0b, 0x5a, 0x74, 0x94, 0x88, 0x7b, 0x49, 0xb0, 0x63, 0x19, 0xb8, 0xb4, 0x6f, 0xb9,
	0xbf, 0x08, 0x8d, 0xa7, 0xa9, 0x5a, 0x17, 0xaa, 0xfc, 0xd8, 0x60, 0x46, 0xbd, 0x80, 0x3c, 0x7c,
	0xf6, 0xc7, 0xb0, 0x1d, 0xe9, 0x2e, 0x90, 0xa5, 0x6a, 0xbe, 0xbe, 0x7b, 0x76, 0xd0, 0x58, 0x0a,
	0x97, 0x89, 0x16, 0x27, 0x09, 0xb5, 0x5f, 0x24, 0x78, 0x21, 0x0a, 0x99, 0x93, 0x20, 0xec, 0x91,
	0x90, 0xed, 0xd9, 0x23, 0xc4, 0x8b, 0x33, 0x3e, 0x58, 0xca, 0x6b, 0x28, 0x32, 0xf3, 0x5d, 0xba,
	0xa5, 0x28, 0x27, 0x4b, 0x72, 0x8d, 0xd0, 0xac, 0xa3, 0x16, 0x70, 0x9a, 0x62, 0xf3, 0x39, 0x0f,
	0x59, 0x4f, 0xdb, 0x01, 0xff, 0x76, 0xf3, 0x38, 0xc1, 0xb5, 0x36, 0xbc, 0x4c, 0x95, 0xfb, 0xd0,
	0xdb, 0xfd, 0xab, 0x62, 0x7f, 0x97, 0xe0, 0x68, 0x19, 0x0c, 0x54, 0x3a, 0x99, 0x90, 0x21, 0x13,
	0x7a, 0x80, 0x3e, 0x81, 0xb2, 0xe6, 0xfa, 0x64, 0x18, 0xc6, 0x41, 0xc3, 0x09, 0x64, 0xa8, 0xe6,
	0xeb, 0x79, 0xbc, 0x1e, 0x48, 0xac, 0x37, 0x9d, 0xfc, 0x94, 0x27, 0xaf, 0xf1, 0xe8, 0x33, 0x78,
	0xa2, 0xf8, 0xd4, 0x76, 0x54, 0x3b, 0xb3, 0xf7, 0x21, 0x4f, 0x7f, 0x28, 0x54, 0x73, 0xa0, 0x64,
	0x2a, 0x96, 0x38, 0xb5, 0x45, 0x47, 0xdc, 0x23, 0x4f, 0x21, 0x8f, 0xe9, 0x82, 0x8f, 0x60, 0xf7,
	0xac, 0xd8, 0x48, 0xc7, 0x31, 0x8b, 0xa0, 0x4f, 0x01, 0x3a, 0x64, 0x11, 0x6d, 0xc3, 0x67, 0xf1,
	0x40, 0x2b, 0x52, 0x29, 0xb5, 0xbf, 0xc4, 0x5d, 0x95, 0x6c, 0xb3, 0x66, 0xc3, 0x15, 0xd8, 0xee,
	0xd3, 0x8c, 0xdd, 0x26, 0x38, 0x7b, 0x4d, 0x16, 0x56, 0xaf, 0x49, 0xf6, 0x59, 0x8d, 0x6d, 0xf6,
	0x31, 0x6e, 0x09, 0xc9, 0x08, 0xc4, 0x3e, 0x78, 0x93, 0x90, 0xec, 0x15, 0x94, 0x62, 0xb8, 0xdd,
	0xdc, 0x12, 0x2f, 0xe4, 0xae, 0x26, 0xcc, 0x6b, 0x49, 0x30, 0xc1, 0x31, 0xe9, 0xb5, 0xe8, 0x88,
	0x41, 0xc3, 0x89, 0x7c, 0x2c, 0x4b, 0xb2, 0xcb, 0x83, 0x29, 0x70, 0x79, 0x01, 0xc5, 0x90, 0xd9,
	0x62, 0x9f, 0x32, 0xbe, 0x2c, 0x6c, 0x91, 0x03, 0x96, 0xaf, 0xdf, 0x85, 0xbe, 0x6d, 0x2a, 0xf2,
	0x93, 0xaa, 0x54, 0xdf, 0xc3, 0x31, 0x8c, 0xcc, 0xcf, 0xb7, 0xbf, 0x0e, 0xa8, 0x17, 0x99, 0xdb,
	0x92, 0xc8, 0xc8, 0xf6, 0x78, 0x45, 0xb6, 0x0a, 0x3c, 0x33, 0x15, 0xcb, 0x9c, 0x07, 0xe3, 0xcc,
	0x10, 0x83, 0xb6, 0xed, 0xdd, 0xa3, 0x8f, 0x60, 0x03, 0xd3, 0x45, 0xac, 0xd6, 0x72, 0x63, 0x75,
	0xd4, 0x98, 0x87, 0xdf, 0xfc, 0x91, 0x83, 0x83, 0x54, 0x65, 0xdc, 0xf4, 0x4f, 0xe0, 0xf8, 0xaa,
	0xf3, 0x4d, 0xa7, 0xfb, 0xb6, 0x63, 0xe1, 0x6e, 0xb7, 0x6d, 0x5d, 0x99, 0xda, 0x79, 0x5f, 0xb7,
	0x5a, 0xdd, 0x66, 0xe9, 0x7f, 0xe8, 0x08, 0x50, 0x47, 0x7f, 0x6b, 0x69, 0x06, 0xd6, 0xd5, 0xbe,
	0xd5, 0xd6, 0x7b, 0xbd, 0xf3, 0xa6, 0x5e, 0x92, 0xd0, 0x21, 0x94, 0xd5, 0xcb, 0xf3, 0x4e, 0x53,
	0x8f, 0x39, 0xcb, 0xd0, 0x4a, 0x39, 0xf4, 0x02, 0x9e, 0xc5, 0x18, 0xeb, 0xaa, 0x71, 0xad, 0x6b,
	0x56, 0xbf, 0x6b, 0xf5, 0x74, 0x7c, 0xad, 0xe3, 0x52, 0x3e, 0x1d, 0xd6, 0xf4, 0x96, 0x71, 0x6d,
	0xe8, 0x98, 0x27, 0x98, 0xba, 0x8e, 0x4b, 0x1b, 0x48, 0x86, 0xa7, 0x71, 0xb8, 0xa7, 0xeb, 0x1d,
	0x4b, 0x79, 0x27, 0x22, 0x05, 0x74, 0x0a, 0x27, 0xa9, 0x85, 0x7a, 0x5f, 0xd7, 0xac, 0x0b, 0xdc,
	0x6d, 0xc7, 0x3b, 0x6f, 0xa2, 0x0a, 0x1c, 0xc5, 0x09, 0xd1, 0xfb, 0x2b, 0xef, 0xac, 0xab, 0x9e,
	0x8e, 0x4b, 0x5b, 0xe9, 0x98, 0x58, 0x9c, 0xc4, 0xb6, 0x59, 0x1d, 0xbc, 0xe8, 0x73, 0xb5, 0x6f,
	0x74, 0x3b, 0x96, 0xd6, 0x35, 0x3a, 0xcd, 0x12, 0x9c, 0xfd, 0x00, 0xfb, 0xd8, 0x54, 0x63, 0x95,
	0x63, 0x72, 0x83, 0xbe, 0x85, 0xc3, 0xa6, 0x70, 0x0a, 0x75, 0xec, 0xd9, 0xc2, 0x1b, 0xb9, 0x31,
	0x3e, 0x6f, 0x3c, 0x62, 0xd9, 0x95, 0x57, 0x8d, 0x7f, 0x72, 0xd6, 0xb3, 0x11, 0x1c, 0x66, 0xcf,
	0xea, 0xbe, 0x7f, 0xcf, 0xfe, 0x75, 0x51, 0x07, 0x76, 0x53, 0xe6, 0x84, 0x5e, 0x36, 0x1e, 0xb5,
	0xd4, 0xca, 0x69, 0xe3, 0x71, 0x57, 0x53, 0xca, 0xb0, 0xed, 0xfa, 0x0d, 0xf6, 0xaf, 0x3d, 0xb8,
	0xcc, 0x9b, 0xd2, 0x77, 0xd2, 0xdd, 0x60, 0x93, 0xff, 0x66, 0x7f, 0xf1, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x24, 0x53, 0x6e, 0xfd, 0x97, 0x0b, 0x00, 0x00,
}
