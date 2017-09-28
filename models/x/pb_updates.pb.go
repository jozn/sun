// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_updates.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UpdateLogEnum int32

const (
	UpdateLogEnum_UNKNOWN_UPDATE_LOG                  UpdateLogEnum = 0
	UpdateLogEnum_UPDATE_NEW_MESSAGE                  UpdateLogEnum = 1
	UpdateLogEnum_UPDATE_MESSAGE_ID                   UpdateLogEnum = 2
	UpdateLogEnum_UPDATE_SEE_NMESSAGES                UpdateLogEnum = 3
	UpdateLogEnum_UPDATE_DELIVIERD_MESSAGES           UpdateLogEnum = 4
	UpdateLogEnum_UPDATE_DELETED_FROM_SERVERM_ESSAGES UpdateLogEnum = 5
	UpdateLogEnum_UPDATE_DELETE_MESSAGES              UpdateLogEnum = 6
	UpdateLogEnum_UPDATE_ROOM_ACTIONDOING             UpdateLogEnum = 10
	UpdateLogEnum_UPDATE_USER_BLOCKED                 UpdateLogEnum = 11
	UpdateLogEnum_UPDATE_NOTIFY_SETTINGS              UpdateLogEnum = 12
	UpdateLogEnum_UPDATE_SERVICE_NOTIFICATION         UpdateLogEnum = 15
)

var UpdateLogEnum_name = map[int32]string{
	0:  "UNKNOWN_UPDATE_LOG",
	1:  "UPDATE_NEW_MESSAGE",
	2:  "UPDATE_MESSAGE_ID",
	3:  "UPDATE_SEE_NMESSAGES",
	4:  "UPDATE_DELIVIERD_MESSAGES",
	5:  "UPDATE_DELETED_FROM_SERVERM_ESSAGES",
	6:  "UPDATE_DELETE_MESSAGES",
	10: "UPDATE_ROOM_ACTIONDOING",
	11: "UPDATE_USER_BLOCKED",
	12: "UPDATE_NOTIFY_SETTINGS",
	15: "UPDATE_SERVICE_NOTIFICATION",
}
var UpdateLogEnum_value = map[string]int32{
	"UNKNOWN_UPDATE_LOG":                  0,
	"UPDATE_NEW_MESSAGE":                  1,
	"UPDATE_MESSAGE_ID":                   2,
	"UPDATE_SEE_NMESSAGES":                3,
	"UPDATE_DELIVIERD_MESSAGES":           4,
	"UPDATE_DELETED_FROM_SERVERM_ESSAGES": 5,
	"UPDATE_DELETE_MESSAGES":              6,
	"UPDATE_ROOM_ACTIONDOING":             10,
	"UPDATE_USER_BLOCKED":                 11,
	"UPDATE_NOTIFY_SETTINGS":              12,
	"UPDATE_SERVICE_NOTIFICATION":         15,
}

func (x UpdateLogEnum) String() string {
	return proto.EnumName(UpdateLogEnum_name, int32(x))
}
func (UpdateLogEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type PB_UpdateGroupParticipants struct {
}

func (m *PB_UpdateGroupParticipants) Reset()                    { *m = PB_UpdateGroupParticipants{} }
func (m *PB_UpdateGroupParticipants) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateGroupParticipants) ProtoMessage()               {}
func (*PB_UpdateGroupParticipants) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type PB_UpdateNotifySettings struct {
}

func (m *PB_UpdateNotifySettings) Reset()                    { *m = PB_UpdateNotifySettings{} }
func (m *PB_UpdateNotifySettings) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateNotifySettings) ProtoMessage()               {}
func (*PB_UpdateNotifySettings) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

type PB_UpdateServiceNotification struct {
}

func (m *PB_UpdateServiceNotification) Reset()                    { *m = PB_UpdateServiceNotification{} }
func (m *PB_UpdateServiceNotification) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateServiceNotification) ProtoMessage()               {}
func (*PB_UpdateServiceNotification) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

// //////////////////////////////////////////////
type PB_UpdateMessageMeta struct {
	MessageId int64 `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
	AtTime    int64 `protobuf:"varint,2,opt,name=AtTime" json:"AtTime,omitempty"`
}

func (m *PB_UpdateMessageMeta) Reset()                    { *m = PB_UpdateMessageMeta{} }
func (m *PB_UpdateMessageMeta) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateMessageMeta) ProtoMessage()               {}
func (*PB_UpdateMessageMeta) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *PB_UpdateMessageMeta) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_UpdateMessageMeta) GetAtTime() int64 {
	if m != nil {
		return m.AtTime
	}
	return 0
}

type PB_UpdateMessageId struct {
	OldMessageId int64 `protobuf:"varint,1,opt,name=OldMessageId" json:"OldMessageId,omitempty"`
	NewMessageId int64 `protobuf:"varint,2,opt,name=NewMessageId" json:"NewMessageId,omitempty"`
}

func (m *PB_UpdateMessageId) Reset()                    { *m = PB_UpdateMessageId{} }
func (m *PB_UpdateMessageId) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateMessageId) ProtoMessage()               {}
func (*PB_UpdateMessageId) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *PB_UpdateMessageId) GetOldMessageId() int64 {
	if m != nil {
		return m.OldMessageId
	}
	return 0
}

func (m *PB_UpdateMessageId) GetNewMessageId() int64 {
	if m != nil {
		return m.NewMessageId
	}
	return 0
}

type PB_UpdateMessageToEdit struct {
	MessageId int64  `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
	NewText   string `protobuf:"bytes,2,opt,name=NewText" json:"NewText,omitempty"`
}

func (m *PB_UpdateMessageToEdit) Reset()                    { *m = PB_UpdateMessageToEdit{} }
func (m *PB_UpdateMessageToEdit) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateMessageToEdit) ProtoMessage()               {}
func (*PB_UpdateMessageToEdit) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *PB_UpdateMessageToEdit) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_UpdateMessageToEdit) GetNewText() string {
	if m != nil {
		return m.NewText
	}
	return ""
}

type PB_UpdateMessageToDelete struct {
	MessageId int64 `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
}

func (m *PB_UpdateMessageToDelete) Reset()                    { *m = PB_UpdateMessageToDelete{} }
func (m *PB_UpdateMessageToDelete) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateMessageToDelete) ProtoMessage()               {}
func (*PB_UpdateMessageToDelete) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *PB_UpdateMessageToDelete) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PB_UpdateRoomActionDoing struct {
	//    int64 ChatId = 1;
	RoomKey    string              `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	ActionType RoomActionDoingEnum `protobuf:"varint,3,opt,name=ActionType,enum=RoomActionDoingEnum" json:"ActionType,omitempty"`
}

func (m *PB_UpdateRoomActionDoing) Reset()                    { *m = PB_UpdateRoomActionDoing{} }
func (m *PB_UpdateRoomActionDoing) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateRoomActionDoing) ProtoMessage()               {}
func (*PB_UpdateRoomActionDoing) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *PB_UpdateRoomActionDoing) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_UpdateRoomActionDoing) GetActionType() RoomActionDoingEnum {
	if m != nil {
		return m.ActionType
	}
	return RoomActionDoingEnum_UNKNOWN_ROOM_ACTION_DOING
}

type PB_UpdateUserBlocked struct {
	UserId  int32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	Blocked bool  `protobuf:"varint,2,opt,name=Blocked" json:"Blocked,omitempty"`
}

func (m *PB_UpdateUserBlocked) Reset()                    { *m = PB_UpdateUserBlocked{} }
func (m *PB_UpdateUserBlocked) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateUserBlocked) ProtoMessage()               {}
func (*PB_UpdateUserBlocked) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *PB_UpdateUserBlocked) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_UpdateUserBlocked) GetBlocked() bool {
	if m != nil {
		return m.Blocked
	}
	return false
}

// todo: remove this - use seprate updates below
type PB_PushHolderView struct {
	// new mssages
	NewMessages       []*PB_MessageView           `protobuf:"bytes,1,rep,name=NewMessages" json:"NewMessages,omitempty"`
	ChatFiles         []*PB_MessageFileView       `protobuf:"bytes,2,rep,name=ChatFiles" json:"ChatFiles,omitempty"`
	Chats             []*PB_ChatView              `protobuf:"bytes,3,rep,name=Chats" json:"Chats,omitempty"`
	Users             []*PB_UserView              `protobuf:"bytes,6,rep,name=Users" json:"Users,omitempty"`
	MessagesChangeIds []*PB_UpdateMessageId       `protobuf:"bytes,10,rep,name=MessagesChangeIds" json:"MessagesChangeIds,omitempty"`
	MessagesToUpdate  []*PB_UpdateMessageToEdit   `protobuf:"bytes,11,rep,name=MessagesToUpdate" json:"MessagesToUpdate,omitempty"`
	MessagesToDelete  []*PB_UpdateMessageToDelete `protobuf:"bytes,12,rep,name=MessagesToDelete" json:"MessagesToDelete,omitempty"`
	// messages meta info
	MessagesDelivierdToServer []*PB_UpdateMessageMeta `protobuf:"bytes,20,rep,name=MessagesDelivierdToServer" json:"MessagesDelivierdToServer,omitempty"`
	MessagesDelivierdToPeer   []*PB_UpdateMessageMeta `protobuf:"bytes,21,rep,name=MessagesDelivierdToPeer" json:"MessagesDelivierdToPeer,omitempty"`
	MessagesSeenByPeer        []*PB_UpdateMessageMeta `protobuf:"bytes,22,rep,name=MessagesSeenByPeer" json:"MessagesSeenByPeer,omitempty"`
	MessagesDeletedFromServer []*PB_UpdateMessageMeta `protobuf:"bytes,23,rep,name=MessagesDeletedFromServer" json:"MessagesDeletedFromServer,omitempty"`
	// room updates
	RoomActionDoing []*PB_UpdateRoomActionDoing `protobuf:"bytes,30,rep,name=RoomActionDoing" json:"RoomActionDoing,omitempty"`
	// users
	UserBlockedByMe []*PB_UpdateUserBlocked `protobuf:"bytes,50,rep,name=UserBlockedByMe" json:"UserBlockedByMe,omitempty"`
	UserBlockedMe   []*PB_UpdateUserBlocked `protobuf:"bytes,51,rep,name=UserBlockedMe" json:"UserBlockedMe,omitempty"`
}

func (m *PB_PushHolderView) Reset()                    { *m = PB_PushHolderView{} }
func (m *PB_PushHolderView) String() string            { return proto.CompactTextString(m) }
func (*PB_PushHolderView) ProtoMessage()               {}
func (*PB_PushHolderView) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *PB_PushHolderView) GetNewMessages() []*PB_MessageView {
	if m != nil {
		return m.NewMessages
	}
	return nil
}

func (m *PB_PushHolderView) GetChatFiles() []*PB_MessageFileView {
	if m != nil {
		return m.ChatFiles
	}
	return nil
}

func (m *PB_PushHolderView) GetChats() []*PB_ChatView {
	if m != nil {
		return m.Chats
	}
	return nil
}

func (m *PB_PushHolderView) GetUsers() []*PB_UserView {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesChangeIds() []*PB_UpdateMessageId {
	if m != nil {
		return m.MessagesChangeIds
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesToUpdate() []*PB_UpdateMessageToEdit {
	if m != nil {
		return m.MessagesToUpdate
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesToDelete() []*PB_UpdateMessageToDelete {
	if m != nil {
		return m.MessagesToDelete
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesDelivierdToServer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDelivierdToServer
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesDelivierdToPeer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDelivierdToPeer
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesSeenByPeer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesSeenByPeer
	}
	return nil
}

func (m *PB_PushHolderView) GetMessagesDeletedFromServer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDeletedFromServer
	}
	return nil
}

func (m *PB_PushHolderView) GetRoomActionDoing() []*PB_UpdateRoomActionDoing {
	if m != nil {
		return m.RoomActionDoing
	}
	return nil
}

func (m *PB_PushHolderView) GetUserBlockedByMe() []*PB_UpdateUserBlocked {
	if m != nil {
		return m.UserBlockedByMe
	}
	return nil
}

func (m *PB_PushHolderView) GetUserBlockedMe() []*PB_UpdateUserBlocked {
	if m != nil {
		return m.UserBlockedMe
	}
	return nil
}

// ////////////////// updates used in rpc_sync //////////////////////
// dep
type PB_DirectUpdatesView struct {
	// new mssages
	NewMessages       []*PB_MessageView           `protobuf:"bytes,1,rep,name=NewMessages" json:"NewMessages,omitempty"`
	ChatFiles         []*PB_MessageFileView       `protobuf:"bytes,2,rep,name=ChatFiles" json:"ChatFiles,omitempty"`
	Chats             []*PB_ChatView              `protobuf:"bytes,3,rep,name=Chats" json:"Chats,omitempty"`
	Users             []*PB_UserView              `protobuf:"bytes,6,rep,name=Users" json:"Users,omitempty"`
	MessagesChangeIds []*PB_UpdateMessageId       `protobuf:"bytes,10,rep,name=MessagesChangeIds" json:"MessagesChangeIds,omitempty"`
	MessagesToUpdate  []*PB_UpdateMessageToEdit   `protobuf:"bytes,11,rep,name=MessagesToUpdate" json:"MessagesToUpdate,omitempty"`
	MessagesToDelete  []*PB_UpdateMessageToDelete `protobuf:"bytes,12,rep,name=MessagesToDelete" json:"MessagesToDelete,omitempty"`
	// messages meta info
	MessagesDelivierdToServer []*PB_UpdateMessageMeta `protobuf:"bytes,20,rep,name=MessagesDelivierdToServer" json:"MessagesDelivierdToServer,omitempty"`
	MessagesDelivierdToPeer   []*PB_UpdateMessageMeta `protobuf:"bytes,21,rep,name=MessagesDelivierdToPeer" json:"MessagesDelivierdToPeer,omitempty"`
	MessagesSeenByPeer        []*PB_UpdateMessageMeta `protobuf:"bytes,22,rep,name=MessagesSeenByPeer" json:"MessagesSeenByPeer,omitempty"`
	MessagesDeletedFromServer []*PB_UpdateMessageMeta `protobuf:"bytes,23,rep,name=MessagesDeletedFromServer" json:"MessagesDeletedFromServer,omitempty"`
	// room updates
	RoomActionDoing []*PB_UpdateRoomActionDoing `protobuf:"bytes,30,rep,name=RoomActionDoing" json:"RoomActionDoing,omitempty"`
}

func (m *PB_DirectUpdatesView) Reset()                    { *m = PB_DirectUpdatesView{} }
func (m *PB_DirectUpdatesView) String() string            { return proto.CompactTextString(m) }
func (*PB_DirectUpdatesView) ProtoMessage()               {}
func (*PB_DirectUpdatesView) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *PB_DirectUpdatesView) GetNewMessages() []*PB_MessageView {
	if m != nil {
		return m.NewMessages
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetChatFiles() []*PB_MessageFileView {
	if m != nil {
		return m.ChatFiles
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetChats() []*PB_ChatView {
	if m != nil {
		return m.Chats
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetUsers() []*PB_UserView {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesChangeIds() []*PB_UpdateMessageId {
	if m != nil {
		return m.MessagesChangeIds
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesToUpdate() []*PB_UpdateMessageToEdit {
	if m != nil {
		return m.MessagesToUpdate
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesToDelete() []*PB_UpdateMessageToDelete {
	if m != nil {
		return m.MessagesToDelete
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesDelivierdToServer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDelivierdToServer
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesDelivierdToPeer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDelivierdToPeer
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesSeenByPeer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesSeenByPeer
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetMessagesDeletedFromServer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDeletedFromServer
	}
	return nil
}

func (m *PB_DirectUpdatesView) GetRoomActionDoing() []*PB_UpdateRoomActionDoing {
	if m != nil {
		return m.RoomActionDoing
	}
	return nil
}

// dep
type PB_GeneralUpdatesView struct {
	// users
	UserBlockedByMe []*PB_UpdateUserBlocked `protobuf:"bytes,50,rep,name=UserBlockedByMe" json:"UserBlockedByMe,omitempty"`
	UserBlockedMe   []*PB_UpdateUserBlocked `protobuf:"bytes,51,rep,name=UserBlockedMe" json:"UserBlockedMe,omitempty"`
}

func (m *PB_GeneralUpdatesView) Reset()                    { *m = PB_GeneralUpdatesView{} }
func (m *PB_GeneralUpdatesView) String() string            { return proto.CompactTextString(m) }
func (*PB_GeneralUpdatesView) ProtoMessage()               {}
func (*PB_GeneralUpdatesView) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *PB_GeneralUpdatesView) GetUserBlockedByMe() []*PB_UpdateUserBlocked {
	if m != nil {
		return m.UserBlockedByMe
	}
	return nil
}

func (m *PB_GeneralUpdatesView) GetUserBlockedMe() []*PB_UpdateUserBlocked {
	if m != nil {
		return m.UserBlockedMe
	}
	return nil
}

// dep
type PB_NotifyUpdatesView struct {
	// users
	UserBlockedByMe []*PB_UpdateUserBlocked `protobuf:"bytes,50,rep,name=UserBlockedByMe" json:"UserBlockedByMe,omitempty"`
	UserBlockedMe   []*PB_UpdateUserBlocked `protobuf:"bytes,51,rep,name=UserBlockedMe" json:"UserBlockedMe,omitempty"`
}

func (m *PB_NotifyUpdatesView) Reset()                    { *m = PB_NotifyUpdatesView{} }
func (m *PB_NotifyUpdatesView) String() string            { return proto.CompactTextString(m) }
func (*PB_NotifyUpdatesView) ProtoMessage()               {}
func (*PB_NotifyUpdatesView) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *PB_NotifyUpdatesView) GetUserBlockedByMe() []*PB_UpdateUserBlocked {
	if m != nil {
		return m.UserBlockedByMe
	}
	return nil
}

func (m *PB_NotifyUpdatesView) GetUserBlockedMe() []*PB_UpdateUserBlocked {
	if m != nil {
		return m.UserBlockedMe
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_UpdateGroupParticipants)(nil), "PB_UpdateGroupParticipants")
	proto.RegisterType((*PB_UpdateNotifySettings)(nil), "PB_UpdateNotifySettings")
	proto.RegisterType((*PB_UpdateServiceNotification)(nil), "PB_UpdateServiceNotification")
	proto.RegisterType((*PB_UpdateMessageMeta)(nil), "PB_UpdateMessageMeta")
	proto.RegisterType((*PB_UpdateMessageId)(nil), "PB_UpdateMessageId")
	proto.RegisterType((*PB_UpdateMessageToEdit)(nil), "PB_UpdateMessageToEdit")
	proto.RegisterType((*PB_UpdateMessageToDelete)(nil), "PB_UpdateMessageToDelete")
	proto.RegisterType((*PB_UpdateRoomActionDoing)(nil), "PB_UpdateRoomActionDoing")
	proto.RegisterType((*PB_UpdateUserBlocked)(nil), "PB_UpdateUserBlocked")
	proto.RegisterType((*PB_PushHolderView)(nil), "PB_PushHolderView")
	proto.RegisterType((*PB_DirectUpdatesView)(nil), "PB_DirectUpdatesView")
	proto.RegisterType((*PB_GeneralUpdatesView)(nil), "PB_GeneralUpdatesView")
	proto.RegisterType((*PB_NotifyUpdatesView)(nil), "PB_NotifyUpdatesView")
	proto.RegisterEnum("UpdateLogEnum", UpdateLogEnum_name, UpdateLogEnum_value)
}

func init() { proto.RegisterFile("pb_updates.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 870 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0x4f, 0x4f, 0xe3, 0x46,
	0x18, 0xc6, 0x1b, 0x52, 0xd8, 0xe5, 0x05, 0x16, 0x33, 0x0b, 0x89, 0x61, 0xe9, 0x16, 0xb9, 0x87,
	0xa2, 0x1e, 0x22, 0x2d, 0xdb, 0x43, 0xa5, 0x1e, 0xaa, 0x24, 0x1e, 0x82, 0x45, 0x62, 0x5b, 0x63,
	0x87, 0x55, 0xab, 0x4a, 0x23, 0x27, 0x79, 0x0b, 0x6e, 0x13, 0x3b, 0xb2, 0x1d, 0xd8, 0x7c, 0x8f,
	0xee, 0xa5, 0x5f, 0xa8, 0x5f, 0xab, 0x9a, 0xb1, 0x9d, 0x38, 0xe6, 0x4f, 0x7b, 0xeb, 0x1e, 0x38,
	0xbe, 0xcf, 0xf3, 0x7b, 0x1f, 0xcf, 0x78, 0xe6, 0x95, 0x06, 0x94, 0xe9, 0x80, 0xcf, 0xa6, 0x23,
	0x2f, 0xc1, 0xb8, 0x31, 0x8d, 0xc2, 0x24, 0x3c, 0x7a, 0x35, 0x1d, 0x70, 0x0c, 0x66, 0x93, 0xbc,
	0xde, 0x99, 0x0e, 0xf8, 0xf0, 0xc6, 0x4b, 0x0a, 0xf6, 0xad, 0x8f, 0x77, 0x99, 0xad, 0x1d, 0xc3,
	0x91, 0xdd, 0xe2, 0x7d, 0x19, 0xd1, 0x89, 0xc2, 0xd9, 0xd4, 0xf6, 0xa2, 0xc4, 0x1f, 0xfa, 0x53,
	0x2f, 0x48, 0x62, 0xed, 0x10, 0xea, 0x0b, 0xd7, 0x0c, 0x13, 0xff, 0xb7, 0xb9, 0x83, 0x49, 0xe2,
	0x07, 0xd7, 0xb1, 0xf6, 0x16, 0x8e, 0x17, 0x96, 0x83, 0xd1, 0xad, 0x3f, 0x4c, 0x09, 0x7f, 0xe8,
	0x25, 0x7e, 0x18, 0x68, 0x5d, 0xd8, 0x5f, 0xf8, 0x3d, 0x8c, 0x63, 0xef, 0x1a, 0x7b, 0x98, 0x78,
	0xe4, 0x18, 0x36, 0xb3, 0xd2, 0x18, 0xa9, 0x95, 0x93, 0xca, 0x69, 0x95, 0x2d, 0x05, 0x52, 0x83,
	0x8d, 0x66, 0xe2, 0xfa, 0x13, 0x54, 0xd7, 0xa4, 0x95, 0x55, 0xda, 0xaf, 0x40, 0xca, 0x69, 0xc6,
	0x88, 0x68, 0xb0, 0x6d, 0x8d, 0x47, 0xe5, 0xb8, 0x15, 0x4d, 0x30, 0x26, 0xde, 0x2d, 0x99, 0x34,
	0x77, 0x45, 0xd3, 0x6c, 0xa8, 0x95, 0xd3, 0xdd, 0x90, 0x8e, 0xfc, 0xe4, 0x5f, 0x56, 0xab, 0xc2,
	0x0b, 0x13, 0xef, 0x5c, 0xfc, 0x98, 0xc8, 0xd8, 0x4d, 0x96, 0x97, 0xda, 0x0f, 0xa0, 0xde, 0x4f,
	0xd4, 0x71, 0x8c, 0x09, 0x3e, 0x9d, 0xa9, 0xfd, 0x5e, 0xe8, 0x64, 0x61, 0x38, 0x69, 0x0e, 0xc5,
	0xef, 0xd4, 0x43, 0x3f, 0xb8, 0x16, 0xdf, 0x13, 0xd2, 0x25, 0xce, 0xf3, 0xef, 0x65, 0x25, 0xf9,
	0x1e, 0x20, 0x05, 0xdd, 0xf9, 0x14, 0xd5, 0xea, 0x49, 0xe5, 0xf4, 0xd5, 0xd9, 0x7e, 0xa3, 0xd4,
	0x4f, 0x83, 0xd9, 0x84, 0x15, 0x38, 0xed, 0xa2, 0x70, 0x46, 0xfd, 0x18, 0xa3, 0xd6, 0x38, 0x1c,
	0xfe, 0x81, 0xf2, 0x14, 0x44, 0x99, 0x2d, 0x6f, 0x9d, 0x65, 0x95, 0xf8, 0x7e, 0x86, 0xc8, 0xef,
	0xbf, 0x64, 0x79, 0xa9, 0x7d, 0x7a, 0x01, 0x7b, 0x76, 0x8b, 0xdb, 0xb3, 0xf8, 0xe6, 0x22, 0x1c,
	0x8f, 0x30, 0xba, 0xf2, 0xf1, 0x8e, 0xbc, 0x83, 0xad, 0xe5, 0x7f, 0x8e, 0xd5, 0xca, 0x49, 0xf5,
	0x74, 0xeb, 0x6c, 0xb7, 0x61, 0xb7, 0x78, 0xa6, 0x09, 0x8a, 0x15, 0x19, 0xf2, 0x0e, 0x36, 0xdb,
	0x37, 0x5e, 0x72, 0xee, 0x8f, 0x31, 0x56, 0xd7, 0x64, 0xc3, 0xeb, 0x42, 0x83, 0xd0, 0x65, 0xd3,
	0x92, 0x22, 0x1a, 0xac, 0x8b, 0x22, 0x56, 0xab, 0x12, 0xdf, 0x16, 0xb8, 0x10, 0x24, 0x97, 0x5a,
	0x82, 0x11, 0x7b, 0x88, 0xd5, 0x8d, 0x25, 0x23, 0x84, 0x94, 0x91, 0x16, 0x69, 0xc2, 0x5e, 0xbe,
	0x8c, 0xf6, 0x8d, 0x17, 0x88, 0xd3, 0x88, 0x55, 0x58, 0x2e, 0xa1, 0x74, 0xfb, 0xd8, 0x7d, 0x9a,
	0xb4, 0x41, 0xc9, 0x45, 0x37, 0x4c, 0x79, 0x75, 0x4b, 0x26, 0xd4, 0x1b, 0x0f, 0xdf, 0x30, 0x76,
	0xaf, 0x81, 0xd0, 0x62, 0x48, 0x7a, 0x67, 0xd4, 0x6d, 0x19, 0x72, 0xd8, 0x78, 0xec, 0x52, 0xb1,
	0x7b, 0x2d, 0xc4, 0x81, 0xc3, 0x5c, 0xd3, 0x71, 0xec, 0xdf, 0xfa, 0x18, 0x8d, 0xdc, 0x50, 0x8c,
	0x2a, 0x46, 0xea, 0xbe, 0xcc, 0x3b, 0x68, 0x3c, 0x34, 0xa2, 0xec, 0xf1, 0x3e, 0x62, 0x41, 0xfd,
	0x01, 0xd3, 0x46, 0x8c, 0xd4, 0x83, 0xa7, 0x22, 0x1f, 0xeb, 0x22, 0x14, 0x48, 0x6e, 0x39, 0x88,
	0x41, 0x6b, 0x2e, 0xb3, 0x6a, 0x4f, 0x65, 0x3d, 0xd0, 0x50, 0xda, 0x2c, 0x26, 0x38, 0x3a, 0x8f,
	0xc2, 0x49, 0xb6, 0xd9, 0xfa, 0x7f, 0xdd, 0xec, 0x6a, 0x1f, 0x69, 0xc3, 0x6e, 0x69, 0x82, 0xd4,
	0xb7, 0xe5, 0x73, 0x28, 0x01, 0xac, 0xdc, 0x41, 0x7e, 0x82, 0xdd, 0xc2, 0x68, 0xb5, 0xe6, 0x3d,
	0x54, 0xcf, 0xca, 0xeb, 0x29, 0x00, 0xac, 0x4c, 0x93, 0x1f, 0x61, 0xa7, 0x20, 0xf5, 0x50, 0x7d,
	0xff, 0x54, 0xfb, 0x2a, 0xab, 0xfd, 0xb5, 0x21, 0x47, 0x5c, 0xf7, 0x23, 0x1c, 0x26, 0x29, 0x1d,
	0x3f, 0x8f, 0xe6, 0xf3, 0x68, 0x3e, 0x8f, 0x66, 0x70, 0xad, 0x7d, 0xaa, 0xc0, 0x81, 0xdd, 0xe2,
	0x1d, 0x0c, 0x30, 0xf2, 0xc6, 0xc5, 0xe9, 0xf8, 0x7f, 0x87, 0xf6, 0xcf, 0x8a, 0x1c, 0xda, 0xf4,
	0xc1, 0xf5, 0xd9, 0x2c, 0xeb, 0xbb, 0xbf, 0xd7, 0x60, 0x27, 0x85, 0xba, 0xa1, 0x7c, 0x4b, 0x90,
	0x1a, 0x90, 0xbe, 0x79, 0x69, 0x5a, 0x1f, 0x4c, 0xde, 0xb7, 0xf5, 0xa6, 0x4b, 0x79, 0xd7, 0xea,
	0x28, 0x5f, 0x48, 0x3d, 0xad, 0x4d, 0xfa, 0x81, 0xf7, 0xa8, 0xe3, 0x34, 0x3b, 0x54, 0xa9, 0x90,
	0x03, 0xd8, 0xcb, 0xf4, 0x4c, 0xe3, 0x86, 0xae, 0xac, 0x11, 0x15, 0xf6, 0x33, 0xd9, 0xa1, 0x94,
	0x9b, 0x99, 0xe7, 0x28, 0x55, 0xf2, 0x15, 0x1c, 0x66, 0x8e, 0x4e, 0xbb, 0xc6, 0x95, 0x41, 0x99,
	0xce, 0x17, 0xf6, 0x97, 0xe4, 0x5b, 0xf8, 0x66, 0x69, 0x53, 0x97, 0xea, 0xfc, 0x9c, 0x59, 0x3d,
	0xee, 0x50, 0x76, 0x45, 0x59, 0x8f, 0xe7, 0xe0, 0x3a, 0x39, 0x82, 0xda, 0x0a, 0xb8, 0x0c, 0xd9,
	0x20, 0x6f, 0xa0, 0x9e, 0x79, 0xcc, 0xb2, 0x7a, 0xbc, 0xd9, 0x76, 0x0d, 0xcb, 0xd4, 0x2d, 0xc3,
	0xec, 0x28, 0x40, 0xea, 0xf0, 0x3a, 0x33, 0xfb, 0x0e, 0x65, 0xbc, 0xd5, 0xb5, 0xda, 0x97, 0x54,
	0x57, 0xb6, 0x0a, 0x89, 0xa6, 0xe5, 0x1a, 0xe7, 0x3f, 0x73, 0x87, 0xba, 0xae, 0x61, 0x76, 0x1c,
	0x65, 0x9b, 0x7c, 0x0d, 0x6f, 0x16, 0xfb, 0x61, 0x57, 0x46, 0x3b, 0x63, 0x8c, 0x76, 0x53, 0x44,
	0x2b, 0xbb, 0xad, 0x3d, 0x78, 0xe9, 0x47, 0x0d, 0xf1, 0x44, 0x1f, 0x5c, 0x54, 0xed, 0xca, 0x2f,
	0x95, 0x8f, 0x83, 0x0d, 0xf9, 0x1c, 0x7f, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0xe9,
	0xfc, 0xcc, 0xd1, 0x0b, 0x00, 0x00,
}
