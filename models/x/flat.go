package x

import "ms/sun/helper"

type GeoLocation_Flat struct {
	Lat float64

	Lon float64
}

type RoomMessageLog_Flat struct {
	typ RoomMessageLogEnum

	TargetUserId int

	ByUserId int
}

type RoomMessageForwardFrom_Flat struct {
	RoomId int

	MessageId int

	RoomTypeEnum int
}

type RoomDraft_Flat struct {
	Message string

	ReplyTo int
}

type ChatRoom_Flat struct {
}

type Pagination_Flat struct {
	Offset int

	Limit int
}

type PB_RoomInlineView_Flat struct {
	RoomId int

	RoomTypeEnum RoomTypeEnum
}

type PB_MessageForwardedFrom_Flat struct {
	RoomId int

	RoomTypeEnum RoomTypeEnum

	MessageId int

	MessageSeq int
}

type PB_MessageView_Flat struct {
	MessageId int

	RoomKey string

	UserId int

	MessageFileId int

	MessageTypeEnum int

	Text string

	CreatedMs int

	PeerReceivedTime int

	PeerSeenTime int

	DeliviryStatusEnum int

	PeerUserId int
}

type PB_GroupView_Flat struct {
	GroupId int

	GroupName string

	MembersCount int

	GroupPrivacyEnum int

	CreatorUserId int

	CreatedTime int

	UpdatedMs int

	CurrentSeq int
}

type PB_GroupMemberView_Flat struct {
	Id int

	GroupId int

	UserId int

	ByUserId int

	GroupRoleEnum int

	CreatedTime int
}

type PB_MessageFileView_Flat struct {
	MessageFileId int

	Name string

	Size int

	FileTypeEnum int

	MimeType string

	Width int

	Height int

	Duration int

	Extension string

	ThumbData64 string

	ServerSrc string

	ServerPath string

	ServerThumbPath string

	BucketId string

	ServerId int

	CanDel int

	CreatedTime int
}

type PB_ReqLastChangesForTheRoom_Flat struct {
	RoomId int

	LastLogId int

	LastHaveSeq int
}

type PB_ResponseLastChangesForTheRoom_Flat struct {
	Messages []PB_MessageView
}

type PB_RequestSetLastSeenMessages_Flat struct {
	RoomId int

	FromMessageId int

	ToMessageId int

	AtTimeMs int
}

type PB_ResponseSetLastSeenMessages_Flat struct {
	Messages []PB_MessageView
}

type PB_MessagesCollections_Flat struct {
	DirectMessagesIds   []int
	GroupMessagesIds    []int
	BroadCatMessagesIds []int
}

type PB_CommandToServer_Flat struct {
	ClientCallId int

	Command string

	Data []byte
}

type PB_CommandToClient_Flat struct {
	ServerCallId int

	Command string

	Data []byte
}

type PB_CommandReceivedToServer_Flat struct {
	ClientCallId int
}

type PB_CommandReceivedToClient_Flat struct {
	ServerCallId int
}

type PB_UserWithMe_Flat struct {
	UserId int

	UserName string

	FirstName string

	LastName string

	About string

	FullName string

	AvatarUrl string

	PrivacyProfile int

	Phone string

	UpdatedTime int

	AppVersion int

	FollowingType int
}

type PB_Message_Flat struct {
	MessageKey string

	RoomKey string

	UserId int

	PeerId int

	RoomTypeId int

	MessageTypeId int

	Text string

	ExtraJson string

	IsByMe int

	IsStared int

	CreatedMs int

	CreatedDeviceMs int

	SortId int

	PeerSeenTime int

	ServerReceivedTime int

	ServerDeletedTime int

	ISeenTime int

	ToPush int

	MsgFile_LocalSrc string

	MsgFile_Status int

	File PB_MsgFile
}

type PB_MsgFile_Flat struct {
	Name string

	Size int

	FileType int

	MimeType string

	Width int

	Height int

	Duration int

	Extension string

	ThumbData []byte

	Data []byte

	ServerSrc string
}

type PB_Response_Flat struct {
	Status int
}

type PB_Request_Flat struct {
	Status int
}

type PB_RequestMsgAddMany_Flat struct {
	Request PB_Request

	Messages []PB_Message
}

type PB_ResponseMsgAddMany_Flat struct {
	Response PB_Response
}

type PB_Push_Flat struct {
	Status int
}

type PB_Result_Flat struct {
	Status int
}

type PB_PushMsgAddMany_Flat struct {
	Push PB_Push

	Messages []PB_Message
	Users    []PB_UserWithMe
}

type PB_ResultMsgAddMany_Flat struct {
	Result PB_Result
}

type PB_MsgEvent_Flat struct {
	MessageKey string

	RoomKey string

	PeerUserId int

	EventType int

	AtTime int
}

type PB_PushMsgEvents_Flat struct {
	Push PB_Push

	Events []PB_MsgEvent
}

type PB_ResultMsgEvents_Flat struct {
	Result PB_Result
}

type PB_UserParam_CheckUserName2_Flat struct {
}

type PB_UserResponse_CheckUserName2_Flat struct {
}

type PB_MsgParam_AddNewTextMessage_Flat struct {
	Text string

	PeerId int

	Time int

	ReplyToMessageId int

	Forward PB_MessageForwardedFrom
}

type PB_MsgResponse_AddNewTextMessage_Flat struct {
}

type PB_MsgParam_SetRoomActionDoing_Flat struct {
	GroupId int

	DirectRoomKey string

	ActionType RoomActionDoingEnum
}

type PB_MsgResponse_SetRoomActionDoing_Flat struct {
}

type PB_MsgParam_GetMessagesByIds_Flat struct {
	MessagesCollections PB_MessagesCollections
}

type PB_MsgResponse_GetMessagesByIds_Flat struct {
	MessagesViews []PB_MessageView
}

type PB_MsgParam_GetMessagesHistory_Flat struct {
	ChatId int

	FromSeq int

	EndSeq int
}

type PB_MsgResponse_GetMessagesHistory_Flat struct {
	MessagesViews []PB_MessageView
}

type PB_MsgParam_SetChatMessagesRangeAsSeen_Flat struct {
	ChatId int

	FromSeq int

	EndSeq int

	SeenTimeMs int
}

type PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat struct {
}

type PB_MsgParam_DeleteChatHistory_Flat struct {
	ChatId int

	ToSeq int
}

type PB_MsgResponse_DeleteChatHistory_Flat struct {
}

type PB_MsgParam_DeleteMessagesByIds_Flat struct {
	MessagesCollections PB_MessagesCollections
}

type PB_MsgResponse_DeleteMessagesByIds_Flat struct {
}

type PB_MsgParam_SetMessagesAsReceived_Flat struct {
	MessagesCollections PB_MessagesCollections
}

type PB_MsgResponse_SetMessagesAsReceived_Flat struct {
}

type PB_MsgParam_ForwardMessages_Flat struct {
	MessagesCollections PB_MessagesCollections

	ToDirectChatIds []int
	ToGroupChatIds  []int
}

type PB_MsgResponse_ForwardMessages_Flat struct {
}

type PB_MsgParam_EditMessage_Flat struct {
	ChatId int

	RoomType RoomTypeEnum

	MessageId int

	NewText string
}

type PB_MsgResponse_EditMessage_Flat struct {
}

type PB_MsgParam_BroadcastNewMessage_Flat struct {
	Text string

	PeerId int

	Time int

	Forward PB_MessageForwardedFrom
}

type PB_MsgResponse_BroadcastNewMessage_Flat struct {
}

type PB_MsgParam_Echo_Flat struct {
	Text string
}

type PB_MsgResponse_PB_MsgParam_Echo_Flat struct {
	Text string
}

type PB_UserParam_BlockUser_Flat struct {
	UserId int

	UserName string
}

type PB_UserOfflineResponse_BlockUser_Flat struct {
	ByUserId int

	TargetUserId int

	TargetUserName string
}

type PB_UserParam_UnBlockUser_Flat struct {
	Offset int

	Limit int
}

type PB_UserOfflineResponse_UnBlockUser_Flat struct {
	Users []UserView
}

type PB_UserParam_BlockedList_Flat struct {
	UserId int

	UserName string
}

type PB_UserResponse_BlockedList_Flat struct {
	ByUserId int

	TargetUserId int

	TargetUserName string
}

type PB_UserParam_UpdateAbout_Flat struct {
	NewAbout string
}

type PB_UserOfflineResponse_UpdateAbout_Flat struct {
	UserId int

	NewAbout string
}

type PB_UserParam_UpdateUserName_Flat struct {
	NewUserName string
}

type PB_UserOfflineResponse_UpdateUserName_Flat struct {
	UserId int

	NewUserName string
}

type PB_UserParam_ChangeAvatar_Flat struct {
	None bool

	ImageData2 []byte
}

type PB_UserOfflineResponse_ChangeAvatar_Flat struct {
}

type PB_UserParam_ChangePrivacy_Flat struct {
	Level ProfilePrivacyLevelEnum
}

type PB_UserResponseOffline_ChangePrivacy_Flat struct {
}

type PB_UserParam_CheckUserName_Flat struct {
	Level ProfilePrivacyLevelEnum
}

type PB_UserResponse_CheckUserName_Flat struct {
}

type UserView_Flat struct {
}

type PB_UpdateNewMessage_Flat struct {
	Message PB_MessageView
}

type PB_UpdateMessageId_Flat struct {
	OldMessageId int

	NewMessageId int
}

type PB_UpdateSeenMessages_Flat struct {
	MessageIds []int
	AtTime     int
}

type PB_UpdateDelivierdMessages_Flat struct {
	MessageIds []int
	AtTime     int
}

type PB_UpdateDeletedFromServerMessages_Flat struct {
	MessageIds []int
	AtTime     int
}

type PB_UpdateDeleteMessages_Flat struct {
	MessageIds []int
}

type PB_UpdateRestoreMessage_Flat struct {
	MessageIds []int
}

type PB_UpdateRoomActionDoing_Flat struct {
	RoomKey string

	ActionType RoomActionDoingEnum
}

type PB_UpdateGroupParticipants_Flat struct {
}

type PB_UpdateUserBlocked_Flat struct {
	UserId int

	Blocked bool
}

type PB_UpdateNotifySettings_Flat struct {
}

type PB_UpdateServiceNotification_Flat struct {
}

type PB_UpdateEditMessage_Flat struct {
	MessageId int

	NewText string
}

///// to_flat ///

func (m *GeoLocation) ToFlat() *GeoLocation_Flat {
	r := &GeoLocation_Flat{
		Lat: float64(m.Lat),
		Lon: float64(m.Lon),
	}
	return r
}

func (m *RoomMessageLog) ToFlat() *RoomMessageLog_Flat {
	r := &RoomMessageLog_Flat{

		TargetUserId: int(m.TargetUserId),
		ByUserId:     int(m.ByUserId),
	}
	return r
}

func (m *RoomMessageForwardFrom) ToFlat() *RoomMessageForwardFrom_Flat {
	r := &RoomMessageForwardFrom_Flat{
		RoomId:       int(m.RoomId),
		MessageId:    int(m.MessageId),
		RoomTypeEnum: int(m.RoomTypeEnum),
	}
	return r
}

func (m *RoomDraft) ToFlat() *RoomDraft_Flat {
	r := &RoomDraft_Flat{
		Message: m.Message,
		ReplyTo: int(m.ReplyTo),
	}
	return r
}

func (m *ChatRoom) ToFlat() *ChatRoom_Flat {
	r := &ChatRoom_Flat{}
	return r
}

func (m *Pagination) ToFlat() *Pagination_Flat {
	r := &Pagination_Flat{
		Offset: int(m.Offset),
		Limit:  int(m.Limit),
	}
	return r
}

func (m *PB_RoomInlineView) ToFlat() *PB_RoomInlineView_Flat {
	r := &PB_RoomInlineView_Flat{
		RoomId: int(m.RoomId),
	}
	return r
}

func (m *PB_MessageForwardedFrom) ToFlat() *PB_MessageForwardedFrom_Flat {
	r := &PB_MessageForwardedFrom_Flat{
		RoomId: int(m.RoomId),

		MessageId:  int(m.MessageId),
		MessageSeq: int(m.MessageSeq),
	}
	return r
}

func (m *PB_MessageView) ToFlat() *PB_MessageView_Flat {
	r := &PB_MessageView_Flat{
		MessageId:          int(m.MessageId),
		RoomKey:            m.RoomKey,
		UserId:             int(m.UserId),
		MessageFileId:      int(m.MessageFileId),
		MessageTypeEnum:    int(m.MessageTypeEnum),
		Text:               m.Text,
		CreatedMs:          int(m.CreatedMs),
		PeerReceivedTime:   int(m.PeerReceivedTime),
		PeerSeenTime:       int(m.PeerSeenTime),
		DeliviryStatusEnum: int(m.DeliviryStatusEnum),
		PeerUserId:         int(m.PeerUserId),
	}
	return r
}

func (m *PB_GroupView) ToFlat() *PB_GroupView_Flat {
	r := &PB_GroupView_Flat{
		GroupId:          int(m.GroupId),
		GroupName:        m.GroupName,
		MembersCount:     int(m.MembersCount),
		GroupPrivacyEnum: int(m.GroupPrivacyEnum),
		CreatorUserId:    int(m.CreatorUserId),
		CreatedTime:      int(m.CreatedTime),
		UpdatedMs:        int(m.UpdatedMs),
		CurrentSeq:       int(m.CurrentSeq),
	}
	return r
}

func (m *PB_GroupMemberView) ToFlat() *PB_GroupMemberView_Flat {
	r := &PB_GroupMemberView_Flat{
		Id:            int(m.Id),
		GroupId:       int(m.GroupId),
		UserId:        int(m.UserId),
		ByUserId:      int(m.ByUserId),
		GroupRoleEnum: int(m.GroupRoleEnum),
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

func (m *PB_MessageFileView) ToFlat() *PB_MessageFileView_Flat {
	r := &PB_MessageFileView_Flat{
		MessageFileId:   int(m.MessageFileId),
		Name:            m.Name,
		Size:            int(m.Size),
		FileTypeEnum:    int(m.FileTypeEnum),
		MimeType:        m.MimeType,
		Width:           int(m.Width),
		Height:          int(m.Height),
		Duration:        int(m.Duration),
		Extension:       m.Extension,
		ThumbData64:     m.ThumbData64,
		ServerSrc:       m.ServerSrc,
		ServerPath:      m.ServerPath,
		ServerThumbPath: m.ServerThumbPath,
		BucketId:        m.BucketId,
		ServerId:        int(m.ServerId),
		CanDel:          int(m.CanDel),
		CreatedTime:     int(m.CreatedTime),
	}
	return r
}

func (m *PB_ReqLastChangesForTheRoom) ToFlat() *PB_ReqLastChangesForTheRoom_Flat {
	r := &PB_ReqLastChangesForTheRoom_Flat{
		RoomId:      int(m.RoomId),
		LastLogId:   int(m.LastLogId),
		LastHaveSeq: int(m.LastHaveSeq),
	}
	return r
}

func (m *PB_ResponseLastChangesForTheRoom) ToFlat() *PB_ResponseLastChangesForTheRoom_Flat {
	r := &PB_ResponseLastChangesForTheRoom_Flat{}
	return r
}

func (m *PB_RequestSetLastSeenMessages) ToFlat() *PB_RequestSetLastSeenMessages_Flat {
	r := &PB_RequestSetLastSeenMessages_Flat{
		RoomId:        int(m.RoomId),
		FromMessageId: int(m.FromMessageId),
		ToMessageId:   int(m.ToMessageId),
		AtTimeMs:      int(m.AtTimeMs),
	}
	return r
}

func (m *PB_ResponseSetLastSeenMessages) ToFlat() *PB_ResponseSetLastSeenMessages_Flat {
	r := &PB_ResponseSetLastSeenMessages_Flat{}
	return r
}

func (m *PB_MessagesCollections) ToFlat() *PB_MessagesCollections_Flat {
	r := &PB_MessagesCollections_Flat{
		DirectMessagesIds:   helper.SliceInt64ToInt(m.DirectMessagesIds),
		GroupMessagesIds:    helper.SliceInt64ToInt(m.GroupMessagesIds),
		BroadCatMessagesIds: helper.SliceInt64ToInt(m.BroadCatMessagesIds),
	}
	return r
}

func (m *PB_CommandToServer) ToFlat() *PB_CommandToServer_Flat {
	r := &PB_CommandToServer_Flat{
		ClientCallId: int(m.ClientCallId),
		Command:      m.Command,
		Data:         []byte(m.Data),
	}
	return r
}

func (m *PB_CommandToClient) ToFlat() *PB_CommandToClient_Flat {
	r := &PB_CommandToClient_Flat{
		ServerCallId: int(m.ServerCallId),
		Command:      m.Command,
		Data:         []byte(m.Data),
	}
	return r
}

func (m *PB_CommandReceivedToServer) ToFlat() *PB_CommandReceivedToServer_Flat {
	r := &PB_CommandReceivedToServer_Flat{
		ClientCallId: int(m.ClientCallId),
	}
	return r
}

func (m *PB_CommandReceivedToClient) ToFlat() *PB_CommandReceivedToClient_Flat {
	r := &PB_CommandReceivedToClient_Flat{
		ServerCallId: int(m.ServerCallId),
	}
	return r
}

func (m *PB_UserWithMe) ToFlat() *PB_UserWithMe_Flat {
	r := &PB_UserWithMe_Flat{
		UserId:         int(m.UserId),
		UserName:       m.UserName,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		About:          m.About,
		FullName:       m.FullName,
		AvatarUrl:      m.AvatarUrl,
		PrivacyProfile: int(m.PrivacyProfile),
		Phone:          m.Phone,
		UpdatedTime:    int(m.UpdatedTime),
		AppVersion:     int(m.AppVersion),
		FollowingType:  int(m.FollowingType),
	}
	return r
}

func (m *PB_Message) ToFlat() *PB_Message_Flat {
	r := &PB_Message_Flat{
		MessageKey:         m.MessageKey,
		RoomKey:            m.RoomKey,
		UserId:             int(m.UserId),
		PeerId:             int(m.PeerId),
		RoomTypeId:         int(m.RoomTypeId),
		MessageTypeId:      int(m.MessageTypeId),
		Text:               m.Text,
		ExtraJson:          m.ExtraJson,
		IsByMe:             int(m.IsByMe),
		IsStared:           int(m.IsStared),
		CreatedMs:          int(m.CreatedMs),
		CreatedDeviceMs:    int(m.CreatedDeviceMs),
		SortId:             int(m.SortId),
		PeerSeenTime:       int(m.PeerSeenTime),
		ServerReceivedTime: int(m.ServerReceivedTime),
		ServerDeletedTime:  int(m.ServerDeletedTime),
		ISeenTime:          int(m.ISeenTime),
		ToPush:             int(m.ToPush),
		MsgFile_LocalSrc:   m.MsgFile_LocalSrc,
		MsgFile_Status:     int(m.MsgFile_Status),
	}
	return r
}

func (m *PB_MsgFile) ToFlat() *PB_MsgFile_Flat {
	r := &PB_MsgFile_Flat{
		Name:      m.Name,
		Size:      int(m.Size),
		FileType:  int(m.FileType),
		MimeType:  m.MimeType,
		Width:     int(m.Width),
		Height:    int(m.Height),
		Duration:  int(m.Duration),
		Extension: m.Extension,
		ThumbData: []byte(m.ThumbData),
		Data:      []byte(m.Data),
		ServerSrc: m.ServerSrc,
	}
	return r
}

func (m *PB_Response) ToFlat() *PB_Response_Flat {
	r := &PB_Response_Flat{
		Status: int(m.Status),
	}
	return r
}

func (m *PB_Request) ToFlat() *PB_Request_Flat {
	r := &PB_Request_Flat{
		Status: int(m.Status),
	}
	return r
}

func (m *PB_RequestMsgAddMany) ToFlat() *PB_RequestMsgAddMany_Flat {
	r := &PB_RequestMsgAddMany_Flat{}
	return r
}

func (m *PB_ResponseMsgAddMany) ToFlat() *PB_ResponseMsgAddMany_Flat {
	r := &PB_ResponseMsgAddMany_Flat{}
	return r
}

func (m *PB_Push) ToFlat() *PB_Push_Flat {
	r := &PB_Push_Flat{
		Status: int(m.Status),
	}
	return r
}

func (m *PB_Result) ToFlat() *PB_Result_Flat {
	r := &PB_Result_Flat{
		Status: int(m.Status),
	}
	return r
}

func (m *PB_PushMsgAddMany) ToFlat() *PB_PushMsgAddMany_Flat {
	r := &PB_PushMsgAddMany_Flat{}
	return r
}

func (m *PB_ResultMsgAddMany) ToFlat() *PB_ResultMsgAddMany_Flat {
	r := &PB_ResultMsgAddMany_Flat{}
	return r
}

func (m *PB_MsgEvent) ToFlat() *PB_MsgEvent_Flat {
	r := &PB_MsgEvent_Flat{
		MessageKey: m.MessageKey,
		RoomKey:    m.RoomKey,
		PeerUserId: int(m.PeerUserId),
		EventType:  int(m.EventType),
		AtTime:     int(m.AtTime),
	}
	return r
}

func (m *PB_PushMsgEvents) ToFlat() *PB_PushMsgEvents_Flat {
	r := &PB_PushMsgEvents_Flat{}
	return r
}

func (m *PB_ResultMsgEvents) ToFlat() *PB_ResultMsgEvents_Flat {
	r := &PB_ResultMsgEvents_Flat{}
	return r
}

func (m *PB_UserParam_CheckUserName2) ToFlat() *PB_UserParam_CheckUserName2_Flat {
	r := &PB_UserParam_CheckUserName2_Flat{}
	return r
}

func (m *PB_UserResponse_CheckUserName2) ToFlat() *PB_UserResponse_CheckUserName2_Flat {
	r := &PB_UserResponse_CheckUserName2_Flat{}
	return r
}

func (m *PB_MsgParam_AddNewTextMessage) ToFlat() *PB_MsgParam_AddNewTextMessage_Flat {
	r := &PB_MsgParam_AddNewTextMessage_Flat{
		Text:             m.Text,
		PeerId:           int(m.PeerId),
		Time:             int(m.Time),
		ReplyToMessageId: int(m.ReplyToMessageId),
	}
	return r
}

func (m *PB_MsgResponse_AddNewTextMessage) ToFlat() *PB_MsgResponse_AddNewTextMessage_Flat {
	r := &PB_MsgResponse_AddNewTextMessage_Flat{}
	return r
}

func (m *PB_MsgParam_SetRoomActionDoing) ToFlat() *PB_MsgParam_SetRoomActionDoing_Flat {
	r := &PB_MsgParam_SetRoomActionDoing_Flat{
		GroupId:       int(m.GroupId),
		DirectRoomKey: m.DirectRoomKey,
	}
	return r
}

func (m *PB_MsgResponse_SetRoomActionDoing) ToFlat() *PB_MsgResponse_SetRoomActionDoing_Flat {
	r := &PB_MsgResponse_SetRoomActionDoing_Flat{}
	return r
}

func (m *PB_MsgParam_GetMessagesByIds) ToFlat() *PB_MsgParam_GetMessagesByIds_Flat {
	r := &PB_MsgParam_GetMessagesByIds_Flat{}
	return r
}

func (m *PB_MsgResponse_GetMessagesByIds) ToFlat() *PB_MsgResponse_GetMessagesByIds_Flat {
	r := &PB_MsgResponse_GetMessagesByIds_Flat{}
	return r
}

func (m *PB_MsgParam_GetMessagesHistory) ToFlat() *PB_MsgParam_GetMessagesHistory_Flat {
	r := &PB_MsgParam_GetMessagesHistory_Flat{
		ChatId:  int(m.ChatId),
		FromSeq: int(m.FromSeq),
		EndSeq:  int(m.EndSeq),
	}
	return r
}

func (m *PB_MsgResponse_GetMessagesHistory) ToFlat() *PB_MsgResponse_GetMessagesHistory_Flat {
	r := &PB_MsgResponse_GetMessagesHistory_Flat{}
	return r
}

func (m *PB_MsgParam_SetChatMessagesRangeAsSeen) ToFlat() *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat {
	r := &PB_MsgParam_SetChatMessagesRangeAsSeen_Flat{
		ChatId:     int(m.ChatId),
		FromSeq:    int(m.FromSeq),
		EndSeq:     int(m.EndSeq),
		SeenTimeMs: int(m.SeenTimeMs),
	}
	return r
}

func (m *PB_MsgResponse_SetChatMessagesRangeAsSeen) ToFlat() *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat {
	r := &PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat{}
	return r
}

func (m *PB_MsgParam_DeleteChatHistory) ToFlat() *PB_MsgParam_DeleteChatHistory_Flat {
	r := &PB_MsgParam_DeleteChatHistory_Flat{
		ChatId: int(m.ChatId),
		ToSeq:  int(m.ToSeq),
	}
	return r
}

func (m *PB_MsgResponse_DeleteChatHistory) ToFlat() *PB_MsgResponse_DeleteChatHistory_Flat {
	r := &PB_MsgResponse_DeleteChatHistory_Flat{}
	return r
}

func (m *PB_MsgParam_DeleteMessagesByIds) ToFlat() *PB_MsgParam_DeleteMessagesByIds_Flat {
	r := &PB_MsgParam_DeleteMessagesByIds_Flat{}
	return r
}

func (m *PB_MsgResponse_DeleteMessagesByIds) ToFlat() *PB_MsgResponse_DeleteMessagesByIds_Flat {
	r := &PB_MsgResponse_DeleteMessagesByIds_Flat{}
	return r
}

func (m *PB_MsgParam_SetMessagesAsReceived) ToFlat() *PB_MsgParam_SetMessagesAsReceived_Flat {
	r := &PB_MsgParam_SetMessagesAsReceived_Flat{}
	return r
}

func (m *PB_MsgResponse_SetMessagesAsReceived) ToFlat() *PB_MsgResponse_SetMessagesAsReceived_Flat {
	r := &PB_MsgResponse_SetMessagesAsReceived_Flat{}
	return r
}

func (m *PB_MsgParam_ForwardMessages) ToFlat() *PB_MsgParam_ForwardMessages_Flat {
	r := &PB_MsgParam_ForwardMessages_Flat{

		ToDirectChatIds: helper.SliceInt64ToInt(m.ToDirectChatIds),
		ToGroupChatIds:  helper.SliceInt64ToInt(m.ToGroupChatIds),
	}
	return r
}

func (m *PB_MsgResponse_ForwardMessages) ToFlat() *PB_MsgResponse_ForwardMessages_Flat {
	r := &PB_MsgResponse_ForwardMessages_Flat{}
	return r
}

func (m *PB_MsgParam_EditMessage) ToFlat() *PB_MsgParam_EditMessage_Flat {
	r := &PB_MsgParam_EditMessage_Flat{
		ChatId: int(m.ChatId),

		MessageId: int(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

func (m *PB_MsgResponse_EditMessage) ToFlat() *PB_MsgResponse_EditMessage_Flat {
	r := &PB_MsgResponse_EditMessage_Flat{}
	return r
}

func (m *PB_MsgParam_BroadcastNewMessage) ToFlat() *PB_MsgParam_BroadcastNewMessage_Flat {
	r := &PB_MsgParam_BroadcastNewMessage_Flat{
		Text:   m.Text,
		PeerId: int(m.PeerId),
		Time:   int(m.Time),
	}
	return r
}

func (m *PB_MsgResponse_BroadcastNewMessage) ToFlat() *PB_MsgResponse_BroadcastNewMessage_Flat {
	r := &PB_MsgResponse_BroadcastNewMessage_Flat{}
	return r
}

func (m *PB_MsgParam_Echo) ToFlat() *PB_MsgParam_Echo_Flat {
	r := &PB_MsgParam_Echo_Flat{
		Text: m.Text,
	}
	return r
}

func (m *PB_MsgResponse_PB_MsgParam_Echo) ToFlat() *PB_MsgResponse_PB_MsgParam_Echo_Flat {
	r := &PB_MsgResponse_PB_MsgParam_Echo_Flat{
		Text: m.Text,
	}
	return r
}

func (m *PB_UserParam_BlockUser) ToFlat() *PB_UserParam_BlockUser_Flat {
	r := &PB_UserParam_BlockUser_Flat{
		UserId:   int(m.UserId),
		UserName: m.UserName,
	}
	return r
}

func (m *PB_UserOfflineResponse_BlockUser) ToFlat() *PB_UserOfflineResponse_BlockUser_Flat {
	r := &PB_UserOfflineResponse_BlockUser_Flat{
		ByUserId:       int(m.ByUserId),
		TargetUserId:   int(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

func (m *PB_UserParam_UnBlockUser) ToFlat() *PB_UserParam_UnBlockUser_Flat {
	r := &PB_UserParam_UnBlockUser_Flat{
		Offset: int(m.Offset),
		Limit:  int(m.Limit),
	}
	return r
}

func (m *PB_UserOfflineResponse_UnBlockUser) ToFlat() *PB_UserOfflineResponse_UnBlockUser_Flat {
	r := &PB_UserOfflineResponse_UnBlockUser_Flat{}
	return r
}

func (m *PB_UserParam_BlockedList) ToFlat() *PB_UserParam_BlockedList_Flat {
	r := &PB_UserParam_BlockedList_Flat{
		UserId:   int(m.UserId),
		UserName: m.UserName,
	}
	return r
}

func (m *PB_UserResponse_BlockedList) ToFlat() *PB_UserResponse_BlockedList_Flat {
	r := &PB_UserResponse_BlockedList_Flat{
		ByUserId:       int(m.ByUserId),
		TargetUserId:   int(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

func (m *PB_UserParam_UpdateAbout) ToFlat() *PB_UserParam_UpdateAbout_Flat {
	r := &PB_UserParam_UpdateAbout_Flat{
		NewAbout: m.NewAbout,
	}
	return r
}

func (m *PB_UserOfflineResponse_UpdateAbout) ToFlat() *PB_UserOfflineResponse_UpdateAbout_Flat {
	r := &PB_UserOfflineResponse_UpdateAbout_Flat{
		UserId:   int(m.UserId),
		NewAbout: m.NewAbout,
	}
	return r
}

func (m *PB_UserParam_UpdateUserName) ToFlat() *PB_UserParam_UpdateUserName_Flat {
	r := &PB_UserParam_UpdateUserName_Flat{
		NewUserName: m.NewUserName,
	}
	return r
}

func (m *PB_UserOfflineResponse_UpdateUserName) ToFlat() *PB_UserOfflineResponse_UpdateUserName_Flat {
	r := &PB_UserOfflineResponse_UpdateUserName_Flat{
		UserId:      int(m.UserId),
		NewUserName: m.NewUserName,
	}
	return r
}

func (m *PB_UserParam_ChangeAvatar) ToFlat() *PB_UserParam_ChangeAvatar_Flat {
	r := &PB_UserParam_ChangeAvatar_Flat{
		None:       m.None,
		ImageData2: []byte(m.ImageData2),
	}
	return r
}

func (m *PB_UserOfflineResponse_ChangeAvatar) ToFlat() *PB_UserOfflineResponse_ChangeAvatar_Flat {
	r := &PB_UserOfflineResponse_ChangeAvatar_Flat{}
	return r
}

func (m *PB_UserParam_ChangePrivacy) ToFlat() *PB_UserParam_ChangePrivacy_Flat {
	r := &PB_UserParam_ChangePrivacy_Flat{}
	return r
}

func (m *PB_UserResponseOffline_ChangePrivacy) ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
	r := &PB_UserResponseOffline_ChangePrivacy_Flat{}
	return r
}

func (m *PB_UserParam_CheckUserName) ToFlat() *PB_UserParam_CheckUserName_Flat {
	r := &PB_UserParam_CheckUserName_Flat{}
	return r
}

func (m *PB_UserResponse_CheckUserName) ToFlat() *PB_UserResponse_CheckUserName_Flat {
	r := &PB_UserResponse_CheckUserName_Flat{}
	return r
}

func (m *UserView) ToFlat() *UserView_Flat {
	r := &UserView_Flat{}
	return r
}

func (m *PB_UpdateNewMessage) ToFlat() *PB_UpdateNewMessage_Flat {
	r := &PB_UpdateNewMessage_Flat{}
	return r
}

func (m *PB_UpdateMessageId) ToFlat() *PB_UpdateMessageId_Flat {
	r := &PB_UpdateMessageId_Flat{
		OldMessageId: int(m.OldMessageId),
		NewMessageId: int(m.NewMessageId),
	}
	return r
}

func (m *PB_UpdateSeenMessages) ToFlat() *PB_UpdateSeenMessages_Flat {
	r := &PB_UpdateSeenMessages_Flat{
		MessageIds: helper.SliceInt64ToInt(m.MessageIds),
		AtTime:     int(m.AtTime),
	}
	return r
}

func (m *PB_UpdateDelivierdMessages) ToFlat() *PB_UpdateDelivierdMessages_Flat {
	r := &PB_UpdateDelivierdMessages_Flat{
		MessageIds: helper.SliceInt64ToInt(m.MessageIds),
		AtTime:     int(m.AtTime),
	}
	return r
}

func (m *PB_UpdateDeletedFromServerMessages) ToFlat() *PB_UpdateDeletedFromServerMessages_Flat {
	r := &PB_UpdateDeletedFromServerMessages_Flat{
		MessageIds: helper.SliceInt64ToInt(m.MessageIds),
		AtTime:     int(m.AtTime),
	}
	return r
}

func (m *PB_UpdateDeleteMessages) ToFlat() *PB_UpdateDeleteMessages_Flat {
	r := &PB_UpdateDeleteMessages_Flat{
		MessageIds: helper.SliceInt64ToInt(m.MessageIds),
	}
	return r
}

func (m *PB_UpdateRestoreMessage) ToFlat() *PB_UpdateRestoreMessage_Flat {
	r := &PB_UpdateRestoreMessage_Flat{
		MessageIds: helper.SliceInt64ToInt(m.MessageIds),
	}
	return r
}

func (m *PB_UpdateRoomActionDoing) ToFlat() *PB_UpdateRoomActionDoing_Flat {
	r := &PB_UpdateRoomActionDoing_Flat{
		RoomKey: m.RoomKey,
	}
	return r
}

func (m *PB_UpdateGroupParticipants) ToFlat() *PB_UpdateGroupParticipants_Flat {
	r := &PB_UpdateGroupParticipants_Flat{}
	return r
}

func (m *PB_UpdateUserBlocked) ToFlat() *PB_UpdateUserBlocked_Flat {
	r := &PB_UpdateUserBlocked_Flat{
		UserId:  int(m.UserId),
		Blocked: m.Blocked,
	}
	return r
}

func (m *PB_UpdateNotifySettings) ToFlat() *PB_UpdateNotifySettings_Flat {
	r := &PB_UpdateNotifySettings_Flat{}
	return r
}

func (m *PB_UpdateServiceNotification) ToFlat() *PB_UpdateServiceNotification_Flat {
	r := &PB_UpdateServiceNotification_Flat{}
	return r
}

func (m *PB_UpdateEditMessage) ToFlat() *PB_UpdateEditMessage_Flat {
	r := &PB_UpdateEditMessage_Flat{
		MessageId: int(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

///// from_flat ///

func (m *GeoLocation_Flat) ToPB() *GeoLocation {
	r := &GeoLocation{
		Lat: m.Lat,
		Lon: m.Lon,
	}
	return r
}

func (m *RoomMessageLog_Flat) ToPB() *RoomMessageLog {
	r := &RoomMessageLog{

		TargetUserId: uint64(m.TargetUserId),
		ByUserId:     uint64(m.ByUserId),
	}
	return r
}

func (m *RoomMessageForwardFrom_Flat) ToPB() *RoomMessageForwardFrom {
	r := &RoomMessageForwardFrom{
		RoomId:       uint64(m.RoomId),
		MessageId:    uint64(m.MessageId),
		RoomTypeEnum: uint32(m.RoomTypeEnum),
	}
	return r
}

func (m *RoomDraft_Flat) ToPB() *RoomDraft {
	r := &RoomDraft{
		Message: m.Message,
		ReplyTo: uint64(m.ReplyTo),
	}
	return r
}

func (m *ChatRoom_Flat) ToPB() *ChatRoom {
	r := &ChatRoom{}
	return r
}

func (m *Pagination_Flat) ToPB() *Pagination {
	r := &Pagination{
		Offset: uint32(m.Offset),
		Limit:  uint32(m.Limit),
	}
	return r
}

func (m *PB_RoomInlineView_Flat) ToPB() *PB_RoomInlineView {
	r := &PB_RoomInlineView{
		RoomId: int64(m.RoomId),
	}
	return r
}

func (m *PB_MessageForwardedFrom_Flat) ToPB() *PB_MessageForwardedFrom {
	r := &PB_MessageForwardedFrom{
		RoomId: int64(m.RoomId),

		MessageId:  int64(m.MessageId),
		MessageSeq: int32(m.MessageSeq),
	}
	return r
}

func (m *PB_MessageView_Flat) ToPB() *PB_MessageView {
	r := &PB_MessageView{
		MessageId:          int64(m.MessageId),
		RoomKey:            m.RoomKey,
		UserId:             int32(m.UserId),
		MessageFileId:      int64(m.MessageFileId),
		MessageTypeEnum:    int32(m.MessageTypeEnum),
		Text:               m.Text,
		CreatedMs:          int64(m.CreatedMs),
		PeerReceivedTime:   int32(m.PeerReceivedTime),
		PeerSeenTime:       int32(m.PeerSeenTime),
		DeliviryStatusEnum: int32(m.DeliviryStatusEnum),
		PeerUserId:         int32(m.PeerUserId),
	}
	return r
}

func (m *PB_GroupView_Flat) ToPB() *PB_GroupView {
	r := &PB_GroupView{
		GroupId:          int64(m.GroupId),
		GroupName:        m.GroupName,
		MembersCount:     int32(m.MembersCount),
		GroupPrivacyEnum: int32(m.GroupPrivacyEnum),
		CreatorUserId:    int32(m.CreatorUserId),
		CreatedTime:      int32(m.CreatedTime),
		UpdatedMs:        int64(m.UpdatedMs),
		CurrentSeq:       int32(m.CurrentSeq),
	}
	return r
}

func (m *PB_GroupMemberView_Flat) ToPB() *PB_GroupMemberView {
	r := &PB_GroupMemberView{
		Id:            int64(m.Id),
		GroupId:       int64(m.GroupId),
		UserId:        int32(m.UserId),
		ByUserId:      int32(m.ByUserId),
		GroupRoleEnum: int32(m.GroupRoleEnum),
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

func (m *PB_MessageFileView_Flat) ToPB() *PB_MessageFileView {
	r := &PB_MessageFileView{
		MessageFileId:   int64(m.MessageFileId),
		Name:            m.Name,
		Size:            int32(m.Size),
		FileTypeEnum:    int32(m.FileTypeEnum),
		MimeType:        m.MimeType,
		Width:           int32(m.Width),
		Height:          int32(m.Height),
		Duration:        int32(m.Duration),
		Extension:       m.Extension,
		ThumbData64:     m.ThumbData64,
		ServerSrc:       m.ServerSrc,
		ServerPath:      m.ServerPath,
		ServerThumbPath: m.ServerThumbPath,
		BucketId:        m.BucketId,
		ServerId:        int32(m.ServerId),
		CanDel:          int32(m.CanDel),
		CreatedTime:     int32(m.CreatedTime),
	}
	return r
}

func (m *PB_ReqLastChangesForTheRoom_Flat) ToPB() *PB_ReqLastChangesForTheRoom {
	r := &PB_ReqLastChangesForTheRoom{
		RoomId:      int64(m.RoomId),
		LastLogId:   int64(m.LastLogId),
		LastHaveSeq: int32(m.LastHaveSeq),
	}
	return r
}

func (m *PB_ResponseLastChangesForTheRoom_Flat) ToPB() *PB_ResponseLastChangesForTheRoom {
	r := &PB_ResponseLastChangesForTheRoom{}
	return r
}

func (m *PB_RequestSetLastSeenMessages_Flat) ToPB() *PB_RequestSetLastSeenMessages {
	r := &PB_RequestSetLastSeenMessages{
		RoomId:        int64(m.RoomId),
		FromMessageId: int64(m.FromMessageId),
		ToMessageId:   int32(m.ToMessageId),
		AtTimeMs:      int64(m.AtTimeMs),
	}
	return r
}

func (m *PB_ResponseSetLastSeenMessages_Flat) ToPB() *PB_ResponseSetLastSeenMessages {
	r := &PB_ResponseSetLastSeenMessages{}
	return r
}

func (m *PB_MessagesCollections_Flat) ToPB() *PB_MessagesCollections {
	r := &PB_MessagesCollections{
		DirectMessagesIds:   helper.SliceIntToInt64(m.DirectMessagesIds),
		GroupMessagesIds:    helper.SliceIntToInt64(m.GroupMessagesIds),
		BroadCatMessagesIds: helper.SliceIntToInt64(m.BroadCatMessagesIds),
	}
	return r
}

func (m *PB_CommandToServer_Flat) ToPB() *PB_CommandToServer {
	r := &PB_CommandToServer{
		ClientCallId: int64(m.ClientCallId),
		Command:      m.Command,
		Data:         m.Data,
	}
	return r
}

func (m *PB_CommandToClient_Flat) ToPB() *PB_CommandToClient {
	r := &PB_CommandToClient{
		ServerCallId: int64(m.ServerCallId),
		Command:      m.Command,
		Data:         m.Data,
	}
	return r
}

func (m *PB_CommandReceivedToServer_Flat) ToPB() *PB_CommandReceivedToServer {
	r := &PB_CommandReceivedToServer{
		ClientCallId: int64(m.ClientCallId),
	}
	return r
}

func (m *PB_CommandReceivedToClient_Flat) ToPB() *PB_CommandReceivedToClient {
	r := &PB_CommandReceivedToClient{
		ServerCallId: int64(m.ServerCallId),
	}
	return r
}

func (m *PB_UserWithMe_Flat) ToPB() *PB_UserWithMe {
	r := &PB_UserWithMe{
		UserId:         int32(m.UserId),
		UserName:       m.UserName,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		About:          m.About,
		FullName:       m.FullName,
		AvatarUrl:      m.AvatarUrl,
		PrivacyProfile: int32(m.PrivacyProfile),
		Phone:          m.Phone,
		UpdatedTime:    int32(m.UpdatedTime),
		AppVersion:     int32(m.AppVersion),
		FollowingType:  int32(m.FollowingType),
	}
	return r
}

func (m *PB_Message_Flat) ToPB() *PB_Message {
	r := &PB_Message{
		MessageKey:         m.MessageKey,
		RoomKey:            m.RoomKey,
		UserId:             int64(m.UserId),
		PeerId:             int64(m.PeerId),
		RoomTypeId:         int32(m.RoomTypeId),
		MessageTypeId:      int32(m.MessageTypeId),
		Text:               m.Text,
		ExtraJson:          m.ExtraJson,
		IsByMe:             int32(m.IsByMe),
		IsStared:           int32(m.IsStared),
		CreatedMs:          int64(m.CreatedMs),
		CreatedDeviceMs:    int64(m.CreatedDeviceMs),
		SortId:             int64(m.SortId),
		PeerSeenTime:       int64(m.PeerSeenTime),
		ServerReceivedTime: int64(m.ServerReceivedTime),
		ServerDeletedTime:  int64(m.ServerDeletedTime),
		ISeenTime:          int64(m.ISeenTime),
		ToPush:             int32(m.ToPush),
		MsgFile_LocalSrc:   m.MsgFile_LocalSrc,
		MsgFile_Status:     int32(m.MsgFile_Status),
	}
	return r
}

func (m *PB_MsgFile_Flat) ToPB() *PB_MsgFile {
	r := &PB_MsgFile{
		Name:      m.Name,
		Size:      int64(m.Size),
		FileType:  int32(m.FileType),
		MimeType:  m.MimeType,
		Width:     int32(m.Width),
		Height:    int32(m.Height),
		Duration:  int32(m.Duration),
		Extension: m.Extension,
		ThumbData: m.ThumbData,
		Data:      m.Data,
		ServerSrc: m.ServerSrc,
	}
	return r
}

func (m *PB_Response_Flat) ToPB() *PB_Response {
	r := &PB_Response{
		Status: int32(m.Status),
	}
	return r
}

func (m *PB_Request_Flat) ToPB() *PB_Request {
	r := &PB_Request{
		Status: int32(m.Status),
	}
	return r
}

func (m *PB_RequestMsgAddMany_Flat) ToPB() *PB_RequestMsgAddMany {
	r := &PB_RequestMsgAddMany{}
	return r
}

func (m *PB_ResponseMsgAddMany_Flat) ToPB() *PB_ResponseMsgAddMany {
	r := &PB_ResponseMsgAddMany{}
	return r
}

func (m *PB_Push_Flat) ToPB() *PB_Push {
	r := &PB_Push{
		Status: int32(m.Status),
	}
	return r
}

func (m *PB_Result_Flat) ToPB() *PB_Result {
	r := &PB_Result{
		Status: int32(m.Status),
	}
	return r
}

func (m *PB_PushMsgAddMany_Flat) ToPB() *PB_PushMsgAddMany {
	r := &PB_PushMsgAddMany{}
	return r
}

func (m *PB_ResultMsgAddMany_Flat) ToPB() *PB_ResultMsgAddMany {
	r := &PB_ResultMsgAddMany{}
	return r
}

func (m *PB_MsgEvent_Flat) ToPB() *PB_MsgEvent {
	r := &PB_MsgEvent{
		MessageKey: m.MessageKey,
		RoomKey:    m.RoomKey,
		PeerUserId: int64(m.PeerUserId),
		EventType:  int32(m.EventType),
		AtTime:     int64(m.AtTime),
	}
	return r
}

func (m *PB_PushMsgEvents_Flat) ToPB() *PB_PushMsgEvents {
	r := &PB_PushMsgEvents{}
	return r
}

func (m *PB_ResultMsgEvents_Flat) ToPB() *PB_ResultMsgEvents {
	r := &PB_ResultMsgEvents{}
	return r
}

func (m *PB_UserParam_CheckUserName2_Flat) ToPB() *PB_UserParam_CheckUserName2 {
	r := &PB_UserParam_CheckUserName2{}
	return r
}

func (m *PB_UserResponse_CheckUserName2_Flat) ToPB() *PB_UserResponse_CheckUserName2 {
	r := &PB_UserResponse_CheckUserName2{}
	return r
}

func (m *PB_MsgParam_AddNewTextMessage_Flat) ToPB() *PB_MsgParam_AddNewTextMessage {
	r := &PB_MsgParam_AddNewTextMessage{
		Text:             m.Text,
		PeerId:           int32(m.PeerId),
		Time:             int32(m.Time),
		ReplyToMessageId: int64(m.ReplyToMessageId),
	}
	return r
}

func (m *PB_MsgResponse_AddNewTextMessage_Flat) ToPB() *PB_MsgResponse_AddNewTextMessage {
	r := &PB_MsgResponse_AddNewTextMessage{}
	return r
}

func (m *PB_MsgParam_SetRoomActionDoing_Flat) ToPB() *PB_MsgParam_SetRoomActionDoing {
	r := &PB_MsgParam_SetRoomActionDoing{
		GroupId:       int64(m.GroupId),
		DirectRoomKey: m.DirectRoomKey,
	}
	return r
}

func (m *PB_MsgResponse_SetRoomActionDoing_Flat) ToPB() *PB_MsgResponse_SetRoomActionDoing {
	r := &PB_MsgResponse_SetRoomActionDoing{}
	return r
}

func (m *PB_MsgParam_GetMessagesByIds_Flat) ToPB() *PB_MsgParam_GetMessagesByIds {
	r := &PB_MsgParam_GetMessagesByIds{}
	return r
}

func (m *PB_MsgResponse_GetMessagesByIds_Flat) ToPB() *PB_MsgResponse_GetMessagesByIds {
	r := &PB_MsgResponse_GetMessagesByIds{}
	return r
}

func (m *PB_MsgParam_GetMessagesHistory_Flat) ToPB() *PB_MsgParam_GetMessagesHistory {
	r := &PB_MsgParam_GetMessagesHistory{
		ChatId:  int64(m.ChatId),
		FromSeq: int32(m.FromSeq),
		EndSeq:  int32(m.EndSeq),
	}
	return r
}

func (m *PB_MsgResponse_GetMessagesHistory_Flat) ToPB() *PB_MsgResponse_GetMessagesHistory {
	r := &PB_MsgResponse_GetMessagesHistory{}
	return r
}

func (m *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat) ToPB() *PB_MsgParam_SetChatMessagesRangeAsSeen {
	r := &PB_MsgParam_SetChatMessagesRangeAsSeen{
		ChatId:     int64(m.ChatId),
		FromSeq:    int32(m.FromSeq),
		EndSeq:     int32(m.EndSeq),
		SeenTimeMs: int64(m.SeenTimeMs),
	}
	return r
}

func (m *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat) ToPB() *PB_MsgResponse_SetChatMessagesRangeAsSeen {
	r := &PB_MsgResponse_SetChatMessagesRangeAsSeen{}
	return r
}

func (m *PB_MsgParam_DeleteChatHistory_Flat) ToPB() *PB_MsgParam_DeleteChatHistory {
	r := &PB_MsgParam_DeleteChatHistory{
		ChatId: int64(m.ChatId),
		ToSeq:  int32(m.ToSeq),
	}
	return r
}

func (m *PB_MsgResponse_DeleteChatHistory_Flat) ToPB() *PB_MsgResponse_DeleteChatHistory {
	r := &PB_MsgResponse_DeleteChatHistory{}
	return r
}

func (m *PB_MsgParam_DeleteMessagesByIds_Flat) ToPB() *PB_MsgParam_DeleteMessagesByIds {
	r := &PB_MsgParam_DeleteMessagesByIds{}
	return r
}

func (m *PB_MsgResponse_DeleteMessagesByIds_Flat) ToPB() *PB_MsgResponse_DeleteMessagesByIds {
	r := &PB_MsgResponse_DeleteMessagesByIds{}
	return r
}

func (m *PB_MsgParam_SetMessagesAsReceived_Flat) ToPB() *PB_MsgParam_SetMessagesAsReceived {
	r := &PB_MsgParam_SetMessagesAsReceived{}
	return r
}

func (m *PB_MsgResponse_SetMessagesAsReceived_Flat) ToPB() *PB_MsgResponse_SetMessagesAsReceived {
	r := &PB_MsgResponse_SetMessagesAsReceived{}
	return r
}

func (m *PB_MsgParam_ForwardMessages_Flat) ToPB() *PB_MsgParam_ForwardMessages {
	r := &PB_MsgParam_ForwardMessages{

		ToDirectChatIds: helper.SliceIntToInt64(m.ToDirectChatIds),
		ToGroupChatIds:  helper.SliceIntToInt64(m.ToGroupChatIds),
	}
	return r
}

func (m *PB_MsgResponse_ForwardMessages_Flat) ToPB() *PB_MsgResponse_ForwardMessages {
	r := &PB_MsgResponse_ForwardMessages{}
	return r
}

func (m *PB_MsgParam_EditMessage_Flat) ToPB() *PB_MsgParam_EditMessage {
	r := &PB_MsgParam_EditMessage{
		ChatId: int64(m.ChatId),

		MessageId: int64(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

func (m *PB_MsgResponse_EditMessage_Flat) ToPB() *PB_MsgResponse_EditMessage {
	r := &PB_MsgResponse_EditMessage{}
	return r
}

func (m *PB_MsgParam_BroadcastNewMessage_Flat) ToPB() *PB_MsgParam_BroadcastNewMessage {
	r := &PB_MsgParam_BroadcastNewMessage{
		Text:   m.Text,
		PeerId: int32(m.PeerId),
		Time:   int32(m.Time),
	}
	return r
}

func (m *PB_MsgResponse_BroadcastNewMessage_Flat) ToPB() *PB_MsgResponse_BroadcastNewMessage {
	r := &PB_MsgResponse_BroadcastNewMessage{}
	return r
}

func (m *PB_MsgParam_Echo_Flat) ToPB() *PB_MsgParam_Echo {
	r := &PB_MsgParam_Echo{
		Text: m.Text,
	}
	return r
}

func (m *PB_MsgResponse_PB_MsgParam_Echo_Flat) ToPB() *PB_MsgResponse_PB_MsgParam_Echo {
	r := &PB_MsgResponse_PB_MsgParam_Echo{
		Text: m.Text,
	}
	return r
}

func (m *PB_UserParam_BlockUser_Flat) ToPB() *PB_UserParam_BlockUser {
	r := &PB_UserParam_BlockUser{
		UserId:   int32(m.UserId),
		UserName: m.UserName,
	}
	return r
}

func (m *PB_UserOfflineResponse_BlockUser_Flat) ToPB() *PB_UserOfflineResponse_BlockUser {
	r := &PB_UserOfflineResponse_BlockUser{
		ByUserId:       int32(m.ByUserId),
		TargetUserId:   int32(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

func (m *PB_UserParam_UnBlockUser_Flat) ToPB() *PB_UserParam_UnBlockUser {
	r := &PB_UserParam_UnBlockUser{
		Offset: int32(m.Offset),
		Limit:  int32(m.Limit),
	}
	return r
}

func (m *PB_UserOfflineResponse_UnBlockUser_Flat) ToPB() *PB_UserOfflineResponse_UnBlockUser {
	r := &PB_UserOfflineResponse_UnBlockUser{}
	return r
}

func (m *PB_UserParam_BlockedList_Flat) ToPB() *PB_UserParam_BlockedList {
	r := &PB_UserParam_BlockedList{
		UserId:   int32(m.UserId),
		UserName: m.UserName,
	}
	return r
}

func (m *PB_UserResponse_BlockedList_Flat) ToPB() *PB_UserResponse_BlockedList {
	r := &PB_UserResponse_BlockedList{
		ByUserId:       int32(m.ByUserId),
		TargetUserId:   int32(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

func (m *PB_UserParam_UpdateAbout_Flat) ToPB() *PB_UserParam_UpdateAbout {
	r := &PB_UserParam_UpdateAbout{
		NewAbout: m.NewAbout,
	}
	return r
}

func (m *PB_UserOfflineResponse_UpdateAbout_Flat) ToPB() *PB_UserOfflineResponse_UpdateAbout {
	r := &PB_UserOfflineResponse_UpdateAbout{
		UserId:   int32(m.UserId),
		NewAbout: m.NewAbout,
	}
	return r
}

func (m *PB_UserParam_UpdateUserName_Flat) ToPB() *PB_UserParam_UpdateUserName {
	r := &PB_UserParam_UpdateUserName{
		NewUserName: m.NewUserName,
	}
	return r
}

func (m *PB_UserOfflineResponse_UpdateUserName_Flat) ToPB() *PB_UserOfflineResponse_UpdateUserName {
	r := &PB_UserOfflineResponse_UpdateUserName{
		UserId:      int32(m.UserId),
		NewUserName: m.NewUserName,
	}
	return r
}

func (m *PB_UserParam_ChangeAvatar_Flat) ToPB() *PB_UserParam_ChangeAvatar {
	r := &PB_UserParam_ChangeAvatar{
		None:       m.None,
		ImageData2: m.ImageData2,
	}
	return r
}

func (m *PB_UserOfflineResponse_ChangeAvatar_Flat) ToPB() *PB_UserOfflineResponse_ChangeAvatar {
	r := &PB_UserOfflineResponse_ChangeAvatar{}
	return r
}

func (m *PB_UserParam_ChangePrivacy_Flat) ToPB() *PB_UserParam_ChangePrivacy {
	r := &PB_UserParam_ChangePrivacy{}
	return r
}

func (m *PB_UserResponseOffline_ChangePrivacy_Flat) ToPB() *PB_UserResponseOffline_ChangePrivacy {
	r := &PB_UserResponseOffline_ChangePrivacy{}
	return r
}

func (m *PB_UserParam_CheckUserName_Flat) ToPB() *PB_UserParam_CheckUserName {
	r := &PB_UserParam_CheckUserName{}
	return r
}

func (m *PB_UserResponse_CheckUserName_Flat) ToPB() *PB_UserResponse_CheckUserName {
	r := &PB_UserResponse_CheckUserName{}
	return r
}

func (m *UserView_Flat) ToPB() *UserView {
	r := &UserView{}
	return r
}

func (m *PB_UpdateNewMessage_Flat) ToPB() *PB_UpdateNewMessage {
	r := &PB_UpdateNewMessage{}
	return r
}

func (m *PB_UpdateMessageId_Flat) ToPB() *PB_UpdateMessageId {
	r := &PB_UpdateMessageId{
		OldMessageId: int64(m.OldMessageId),
		NewMessageId: int64(m.NewMessageId),
	}
	return r
}

func (m *PB_UpdateSeenMessages_Flat) ToPB() *PB_UpdateSeenMessages {
	r := &PB_UpdateSeenMessages{
		MessageIds: helper.SliceIntToInt64(m.MessageIds),
		AtTime:     int64(m.AtTime),
	}
	return r
}

func (m *PB_UpdateDelivierdMessages_Flat) ToPB() *PB_UpdateDelivierdMessages {
	r := &PB_UpdateDelivierdMessages{
		MessageIds: helper.SliceIntToInt64(m.MessageIds),
		AtTime:     int64(m.AtTime),
	}
	return r
}

func (m *PB_UpdateDeletedFromServerMessages_Flat) ToPB() *PB_UpdateDeletedFromServerMessages {
	r := &PB_UpdateDeletedFromServerMessages{
		MessageIds: helper.SliceIntToInt64(m.MessageIds),
		AtTime:     int64(m.AtTime),
	}
	return r
}

func (m *PB_UpdateDeleteMessages_Flat) ToPB() *PB_UpdateDeleteMessages {
	r := &PB_UpdateDeleteMessages{
		MessageIds: helper.SliceIntToInt64(m.MessageIds),
	}
	return r
}

func (m *PB_UpdateRestoreMessage_Flat) ToPB() *PB_UpdateRestoreMessage {
	r := &PB_UpdateRestoreMessage{
		MessageIds: helper.SliceIntToInt64(m.MessageIds),
	}
	return r
}

func (m *PB_UpdateRoomActionDoing_Flat) ToPB() *PB_UpdateRoomActionDoing {
	r := &PB_UpdateRoomActionDoing{
		RoomKey: m.RoomKey,
	}
	return r
}

func (m *PB_UpdateGroupParticipants_Flat) ToPB() *PB_UpdateGroupParticipants {
	r := &PB_UpdateGroupParticipants{}
	return r
}

func (m *PB_UpdateUserBlocked_Flat) ToPB() *PB_UpdateUserBlocked {
	r := &PB_UpdateUserBlocked{
		UserId:  int32(m.UserId),
		Blocked: m.Blocked,
	}
	return r
}

func (m *PB_UpdateNotifySettings_Flat) ToPB() *PB_UpdateNotifySettings {
	r := &PB_UpdateNotifySettings{}
	return r
}

func (m *PB_UpdateServiceNotification_Flat) ToPB() *PB_UpdateServiceNotification {
	r := &PB_UpdateServiceNotification{}
	return r
}

func (m *PB_UpdateEditMessage_Flat) ToPB() *PB_UpdateEditMessage {
	r := &PB_UpdateEditMessage{
		MessageId: int64(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}
