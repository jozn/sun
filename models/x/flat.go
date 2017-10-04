package x

import "ms/sun/helper"

type PB_RoomInlineView_Flat struct {
	RoomId       int
	RoomTypeEnum RoomTypeEnum
}

//ToPB
func (m *PB_RoomInlineView) ToFlat() *PB_RoomInlineView_Flat {
	r := &PB_RoomInlineView_Flat{
		RoomId: int(m.RoomId),
	}
	return r
}

//ToPB
func (m *PB_RoomInlineView_Flat) ToPB() *PB_RoomInlineView {
	r := &PB_RoomInlineView{
		RoomId: int64(m.RoomId),
	}
	return r
}

//folding
var PB_RoomInlineView__FOlD = &PB_RoomInlineView{
	RoomId: 0,
}

type PB_MessageForwardedFrom_Flat struct {
	RoomId       int
	RoomTypeEnum RoomTypeEnum
	MessageId    int
	MessageSeq   int
}

//ToPB
func (m *PB_MessageForwardedFrom) ToFlat() *PB_MessageForwardedFrom_Flat {
	r := &PB_MessageForwardedFrom_Flat{
		RoomId: int(m.RoomId),

		MessageId:  int(m.MessageId),
		MessageSeq: int(m.MessageSeq),
	}
	return r
}

//ToPB
func (m *PB_MessageForwardedFrom_Flat) ToPB() *PB_MessageForwardedFrom {
	r := &PB_MessageForwardedFrom{
		RoomId: int64(m.RoomId),

		MessageId:  int64(m.MessageId),
		MessageSeq: int32(m.MessageSeq),
	}
	return r
}

//folding
var PB_MessageForwardedFrom__FOlD = &PB_MessageForwardedFrom{
	RoomId: 0,

	MessageId:  0,
	MessageSeq: 0,
}

type PB_GroupView_Flat struct {
	GroupId          int
	GroupName        string
	MembersCount     int
	GroupPrivacyEnum int
	CreatorUserId    int
	CreatedTime      int
	UpdatedMs        int
	CurrentSeq       int
}

//ToPB
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

//ToPB
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

//folding
var PB_GroupView__FOlD = &PB_GroupView{
	GroupId:          0,
	GroupName:        "",
	MembersCount:     0,
	GroupPrivacyEnum: 0,
	CreatorUserId:    0,
	CreatedTime:      0,
	UpdatedMs:        0,
	CurrentSeq:       0,
}

type PB_GroupMemberView_Flat struct {
	Id            int
	GroupId       int
	UserId        int
	ByUserId      int
	GroupRoleEnum int
	CreatedTime   int
}

//ToPB
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

//ToPB
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

//folding
var PB_GroupMemberView__FOlD = &PB_GroupMemberView{
	Id:            0,
	GroupId:       0,
	UserId:        0,
	ByUserId:      0,
	GroupRoleEnum: 0,
	CreatedTime:   0,
}

type PB_MessageFileView__DEp_Flat struct {
	MessageFileId   int
	Name            string
	Size            int
	FileTypeEnum    int
	MimeType        string
	Width           int
	Height          int
	Duration        int
	Extension       string
	ThumbData64     string
	ServerSrc       string
	ServerPath      string
	ServerThumbPath string
	BucketId        string
	ServerId        int
	CanDel          int
	CreatedTime     int
}

//ToPB
func (m *PB_MessageFileView__DEp) ToFlat() *PB_MessageFileView__DEp_Flat {
	r := &PB_MessageFileView__DEp_Flat{
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

//ToPB
func (m *PB_MessageFileView__DEp_Flat) ToPB() *PB_MessageFileView__DEp {
	r := &PB_MessageFileView__DEp{
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

//folding
var PB_MessageFileView__DEp__FOlD = &PB_MessageFileView__DEp{
	MessageFileId:   0,
	Name:            "",
	Size:            0,
	FileTypeEnum:    0,
	MimeType:        "",
	Width:           0,
	Height:          0,
	Duration:        0,
	Extension:       "",
	ThumbData64:     "",
	ServerSrc:       "",
	ServerPath:      "",
	ServerThumbPath: "",
	BucketId:        "",
	ServerId:        0,
	CanDel:          0,
	CreatedTime:     0,
}

type PB_ReqLastChangesForTheRoom_Flat struct {
	RoomId      int
	LastLogId   int
	LastHaveSeq int
}

//ToPB
func (m *PB_ReqLastChangesForTheRoom) ToFlat() *PB_ReqLastChangesForTheRoom_Flat {
	r := &PB_ReqLastChangesForTheRoom_Flat{
		RoomId:      int(m.RoomId),
		LastLogId:   int(m.LastLogId),
		LastHaveSeq: int(m.LastHaveSeq),
	}
	return r
}

//ToPB
func (m *PB_ReqLastChangesForTheRoom_Flat) ToPB() *PB_ReqLastChangesForTheRoom {
	r := &PB_ReqLastChangesForTheRoom{
		RoomId:      int64(m.RoomId),
		LastLogId:   int64(m.LastLogId),
		LastHaveSeq: int32(m.LastHaveSeq),
	}
	return r
}

//folding
var PB_ReqLastChangesForTheRoom__FOlD = &PB_ReqLastChangesForTheRoom{
	RoomId:      0,
	LastLogId:   0,
	LastHaveSeq: 0,
}

type PB_ResponseLastChangesForTheRoom_Flat struct {
	Messages []PB_MessageView
}

//ToPB
func (m *PB_ResponseLastChangesForTheRoom) ToFlat() *PB_ResponseLastChangesForTheRoom_Flat {
	r := &PB_ResponseLastChangesForTheRoom_Flat{}
	return r
}

//ToPB
func (m *PB_ResponseLastChangesForTheRoom_Flat) ToPB() *PB_ResponseLastChangesForTheRoom {
	r := &PB_ResponseLastChangesForTheRoom{}
	return r
}

//folding
var PB_ResponseLastChangesForTheRoom__FOlD = &PB_ResponseLastChangesForTheRoom{}

type PB_RequestSetLastSeenMessages_Flat struct {
	RoomId        int
	FromMessageId int
	ToMessageId   int
	AtTimeMs      int
}

//ToPB
func (m *PB_RequestSetLastSeenMessages) ToFlat() *PB_RequestSetLastSeenMessages_Flat {
	r := &PB_RequestSetLastSeenMessages_Flat{
		RoomId:        int(m.RoomId),
		FromMessageId: int(m.FromMessageId),
		ToMessageId:   int(m.ToMessageId),
		AtTimeMs:      int(m.AtTimeMs),
	}
	return r
}

//ToPB
func (m *PB_RequestSetLastSeenMessages_Flat) ToPB() *PB_RequestSetLastSeenMessages {
	r := &PB_RequestSetLastSeenMessages{
		RoomId:        int64(m.RoomId),
		FromMessageId: int64(m.FromMessageId),
		ToMessageId:   int32(m.ToMessageId),
		AtTimeMs:      int64(m.AtTimeMs),
	}
	return r
}

//folding
var PB_RequestSetLastSeenMessages__FOlD = &PB_RequestSetLastSeenMessages{
	RoomId:        0,
	FromMessageId: 0,
	ToMessageId:   0,
	AtTimeMs:      0,
}

type PB_ResponseSetLastSeenMessages_Flat struct {
	Messages []PB_MessageView
}

//ToPB
func (m *PB_ResponseSetLastSeenMessages) ToFlat() *PB_ResponseSetLastSeenMessages_Flat {
	r := &PB_ResponseSetLastSeenMessages_Flat{}
	return r
}

//ToPB
func (m *PB_ResponseSetLastSeenMessages_Flat) ToPB() *PB_ResponseSetLastSeenMessages {
	r := &PB_ResponseSetLastSeenMessages{}
	return r
}

//folding
var PB_ResponseSetLastSeenMessages__FOlD = &PB_ResponseSetLastSeenMessages{}

type PB_MessagesCollections_Flat struct {
	DirectMessagesIds   []int
	GroupMessagesIds    []int
	BroadCatMessagesIds []int
}

//ToPB
func (m *PB_MessagesCollections) ToFlat() *PB_MessagesCollections_Flat {
	r := &PB_MessagesCollections_Flat{
		DirectMessagesIds:   helper.SliceInt64ToInt(m.DirectMessagesIds),
		GroupMessagesIds:    helper.SliceInt64ToInt(m.GroupMessagesIds),
		BroadCatMessagesIds: helper.SliceInt64ToInt(m.BroadCatMessagesIds),
	}
	return r
}

//ToPB
func (m *PB_MessagesCollections_Flat) ToPB() *PB_MessagesCollections {
	r := &PB_MessagesCollections{
		DirectMessagesIds:   helper.SliceIntToInt64(m.DirectMessagesIds),
		GroupMessagesIds:    helper.SliceIntToInt64(m.GroupMessagesIds),
		BroadCatMessagesIds: helper.SliceIntToInt64(m.BroadCatMessagesIds),
	}
	return r
}

//folding
var PB_MessagesCollections__FOlD = &PB_MessagesCollections{}

type PB_DirectLogView_Flat struct {
	Row        PB_DirectLog
	NewMessage PB_MessageView
}

//ToPB
func (m *PB_DirectLogView) ToFlat() *PB_DirectLogView_Flat {
	r := &PB_DirectLogView_Flat{}
	return r
}

//ToPB
func (m *PB_DirectLogView_Flat) ToPB() *PB_DirectLogView {
	r := &PB_DirectLogView{}
	return r
}

//folding
var PB_DirectLogView__FOlD = &PB_DirectLogView{}

type PB_DirectLog_Flat struct {
	Id            int
	ToUserId      int
	MessageId     int
	ChatId        int
	PeerUserId    int
	EventType     int
	RoomLogTypeId int
	FromSeq       int
	ToSeq         int
	ExtraPB       []byte
	ExtraJson     string
	AtTimeMs      int
}

//ToPB
func (m *PB_DirectLog) ToFlat() *PB_DirectLog_Flat {
	r := &PB_DirectLog_Flat{
		Id:            int(m.Id),
		ToUserId:      int(m.ToUserId),
		MessageId:     int(m.MessageId),
		ChatId:        int(m.ChatId),
		PeerUserId:    int(m.PeerUserId),
		EventType:     int(m.EventType),
		RoomLogTypeId: int(m.RoomLogTypeId),
		FromSeq:       int(m.FromSeq),
		ToSeq:         int(m.ToSeq),
		ExtraPB:       []byte(m.ExtraPB),
		ExtraJson:     m.ExtraJson,
		AtTimeMs:      int(m.AtTimeMs),
	}
	return r
}

//ToPB
func (m *PB_DirectLog_Flat) ToPB() *PB_DirectLog {
	r := &PB_DirectLog{
		Id:            int64(m.Id),
		ToUserId:      int32(m.ToUserId),
		MessageId:     int64(m.MessageId),
		ChatId:        int64(m.ChatId),
		PeerUserId:    int32(m.PeerUserId),
		EventType:     int32(m.EventType),
		RoomLogTypeId: int32(m.RoomLogTypeId),
		FromSeq:       int32(m.FromSeq),
		ToSeq:         int32(m.ToSeq),
		ExtraPB:       m.ExtraPB,
		ExtraJson:     m.ExtraJson,
		AtTimeMs:      int64(m.AtTimeMs),
	}
	return r
}

//folding
var PB_DirectLog__FOlD = &PB_DirectLog{
	Id:            0,
	ToUserId:      0,
	MessageId:     0,
	ChatId:        0,
	PeerUserId:    0,
	EventType:     0,
	RoomLogTypeId: 0,
	FromSeq:       0,
	ToSeq:         0,
	ExtraPB:       []byte{},
	ExtraJson:     "",
	AtTimeMs:      0,
}

type PB_PushDirectLogViewsMany_Flat struct {
	Rows []PB_DirectLogView
}

//ToPB
func (m *PB_PushDirectLogViewsMany) ToFlat() *PB_PushDirectLogViewsMany_Flat {
	r := &PB_PushDirectLogViewsMany_Flat{}
	return r
}

//ToPB
func (m *PB_PushDirectLogViewsMany_Flat) ToPB() *PB_PushDirectLogViewsMany {
	r := &PB_PushDirectLogViewsMany{}
	return r
}

//folding
var PB_PushDirectLogViewsMany__FOlD = &PB_PushDirectLogViewsMany{}

type PB_UserWithMe_Flat struct {
	UserId         int
	UserName       string
	FirstName      string
	LastName       string
	About          string
	FullName       string
	AvatarUrl      string
	PrivacyProfile int
	Phone          string
	UpdatedTime    int
	AppVersion     int
	FollowingType  int
}

//ToPB
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

//ToPB
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

//folding
var PB_UserWithMe__FOlD = &PB_UserWithMe{
	UserId:         0,
	UserName:       "",
	FirstName:      "",
	LastName:       "",
	About:          "",
	FullName:       "",
	AvatarUrl:      "",
	PrivacyProfile: 0,
	Phone:          "",
	UpdatedTime:    0,
	AppVersion:     0,
	FollowingType:  0,
}

type PB_Message_Flat struct {
	MessageKey         string
	RoomKey            string
	UserId             int
	PeerId             int
	RoomTypeId         int
	MessageTypeId      int
	Text               string
	ExtraJson          string
	IsByMe             int
	IsStared           int
	CreatedMs          int
	CreatedDeviceMs    int
	SortId             int
	PeerSeenTime       int
	ServerReceivedTime int
	ServerDeletedTime  int
	ISeenTime          int
	ToPush             int
	MsgFile_LocalSrc   string
	MsgFile_Status     int
	File               PB_MsgFile
}

//ToPB
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

//ToPB
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

//folding
var PB_Message__FOlD = &PB_Message{
	MessageKey:         "",
	RoomKey:            "",
	UserId:             0,
	PeerId:             0,
	RoomTypeId:         0,
	MessageTypeId:      0,
	Text:               "",
	ExtraJson:          "",
	IsByMe:             0,
	IsStared:           0,
	CreatedMs:          0,
	CreatedDeviceMs:    0,
	SortId:             0,
	PeerSeenTime:       0,
	ServerReceivedTime: 0,
	ServerDeletedTime:  0,
	ISeenTime:          0,
	ToPush:             0,
	MsgFile_LocalSrc:   "",
	MsgFile_Status:     0,
}

type PB_MsgFile_Flat struct {
	Name      string
	Size      int
	FileType  int
	MimeType  string
	Width     int
	Height    int
	Duration  int
	Extension string
	ThumbData []byte
	Data      []byte
	ServerSrc string
}

//ToPB
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

//ToPB
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

//folding
var PB_MsgFile__FOlD = &PB_MsgFile{
	Name:      "",
	Size:      0,
	FileType:  0,
	MimeType:  "",
	Width:     0,
	Height:    0,
	Duration:  0,
	Extension: "",
	ThumbData: []byte{},
	Data:      []byte{},
	ServerSrc: "",
}

type PB_Response_Flat struct {
	Status int
}

//ToPB
func (m *PB_Response) ToFlat() *PB_Response_Flat {
	r := &PB_Response_Flat{
		Status: int(m.Status),
	}
	return r
}

//ToPB
func (m *PB_Response_Flat) ToPB() *PB_Response {
	r := &PB_Response{
		Status: int32(m.Status),
	}
	return r
}

//folding
var PB_Response__FOlD = &PB_Response{
	Status: 0,
}

type PB_Request_Flat struct {
	Status int
}

//ToPB
func (m *PB_Request) ToFlat() *PB_Request_Flat {
	r := &PB_Request_Flat{
		Status: int(m.Status),
	}
	return r
}

//ToPB
func (m *PB_Request_Flat) ToPB() *PB_Request {
	r := &PB_Request{
		Status: int32(m.Status),
	}
	return r
}

//folding
var PB_Request__FOlD = &PB_Request{
	Status: 0,
}

type PB_RequestMsgAddMany_Flat struct {
	Request  PB_Request
	Messages []PB_Message
}

//ToPB
func (m *PB_RequestMsgAddMany) ToFlat() *PB_RequestMsgAddMany_Flat {
	r := &PB_RequestMsgAddMany_Flat{}
	return r
}

//ToPB
func (m *PB_RequestMsgAddMany_Flat) ToPB() *PB_RequestMsgAddMany {
	r := &PB_RequestMsgAddMany{}
	return r
}

//folding
var PB_RequestMsgAddMany__FOlD = &PB_RequestMsgAddMany{}

type PB_ResponseMsgAddMany_Flat struct {
	Response PB_Response
}

//ToPB
func (m *PB_ResponseMsgAddMany) ToFlat() *PB_ResponseMsgAddMany_Flat {
	r := &PB_ResponseMsgAddMany_Flat{}
	return r
}

//ToPB
func (m *PB_ResponseMsgAddMany_Flat) ToPB() *PB_ResponseMsgAddMany {
	r := &PB_ResponseMsgAddMany{}
	return r
}

//folding
var PB_ResponseMsgAddMany__FOlD = &PB_ResponseMsgAddMany{}

type PB_Push_Flat struct {
	Status int
}

//ToPB
func (m *PB_Push) ToFlat() *PB_Push_Flat {
	r := &PB_Push_Flat{
		Status: int(m.Status),
	}
	return r
}

//ToPB
func (m *PB_Push_Flat) ToPB() *PB_Push {
	r := &PB_Push{
		Status: int32(m.Status),
	}
	return r
}

//folding
var PB_Push__FOlD = &PB_Push{
	Status: 0,
}

type PB_Result_Flat struct {
	Status int
}

//ToPB
func (m *PB_Result) ToFlat() *PB_Result_Flat {
	r := &PB_Result_Flat{
		Status: int(m.Status),
	}
	return r
}

//ToPB
func (m *PB_Result_Flat) ToPB() *PB_Result {
	r := &PB_Result{
		Status: int32(m.Status),
	}
	return r
}

//folding
var PB_Result__FOlD = &PB_Result{
	Status: 0,
}

type PB_PushMsgAddMany_Flat struct {
	Push     PB_Push
	Messages []PB_Message
	Users    []PB_UserWithMe
}

//ToPB
func (m *PB_PushMsgAddMany) ToFlat() *PB_PushMsgAddMany_Flat {
	r := &PB_PushMsgAddMany_Flat{}
	return r
}

//ToPB
func (m *PB_PushMsgAddMany_Flat) ToPB() *PB_PushMsgAddMany {
	r := &PB_PushMsgAddMany{}
	return r
}

//folding
var PB_PushMsgAddMany__FOlD = &PB_PushMsgAddMany{}

type PB_ResultMsgAddMany_Flat struct {
	Result PB_Result
}

//ToPB
func (m *PB_ResultMsgAddMany) ToFlat() *PB_ResultMsgAddMany_Flat {
	r := &PB_ResultMsgAddMany_Flat{}
	return r
}

//ToPB
func (m *PB_ResultMsgAddMany_Flat) ToPB() *PB_ResultMsgAddMany {
	r := &PB_ResultMsgAddMany{}
	return r
}

//folding
var PB_ResultMsgAddMany__FOlD = &PB_ResultMsgAddMany{}

type PB_MsgEvent_Flat struct {
	MessageKey string
	RoomKey    string
	PeerUserId int
	EventType  int
	AtTime     int
}

//ToPB
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

//ToPB
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

//folding
var PB_MsgEvent__FOlD = &PB_MsgEvent{
	MessageKey: "",
	RoomKey:    "",
	PeerUserId: 0,
	EventType:  0,
	AtTime:     0,
}

type PB_PushMsgEvents_Flat struct {
	Push   PB_Push
	Events []PB_MsgEvent
}

//ToPB
func (m *PB_PushMsgEvents) ToFlat() *PB_PushMsgEvents_Flat {
	r := &PB_PushMsgEvents_Flat{}
	return r
}

//ToPB
func (m *PB_PushMsgEvents_Flat) ToPB() *PB_PushMsgEvents {
	r := &PB_PushMsgEvents{}
	return r
}

//folding
var PB_PushMsgEvents__FOlD = &PB_PushMsgEvents{}

type PB_ResultMsgEvents_Flat struct {
	Result PB_Result
}

//ToPB
func (m *PB_ResultMsgEvents) ToFlat() *PB_ResultMsgEvents_Flat {
	r := &PB_ResultMsgEvents_Flat{}
	return r
}

//ToPB
func (m *PB_ResultMsgEvents_Flat) ToPB() *PB_ResultMsgEvents {
	r := &PB_ResultMsgEvents{}
	return r
}

//folding
var PB_ResultMsgEvents__FOlD = &PB_ResultMsgEvents{}

type PB_MsgsSeenFromClient_Flat struct {
	MessageKey string
	RoomKey    string
	UserId     int
	AtTime     int
}

//ToPB
func (m *PB_MsgsSeenFromClient) ToFlat() *PB_MsgsSeenFromClient_Flat {
	r := &PB_MsgsSeenFromClient_Flat{
		MessageKey: m.MessageKey,
		RoomKey:    m.RoomKey,
		UserId:     int(m.UserId),
		AtTime:     int(m.AtTime),
	}
	return r
}

//ToPB
func (m *PB_MsgsSeenFromClient_Flat) ToPB() *PB_MsgsSeenFromClient {
	r := &PB_MsgsSeenFromClient{
		MessageKey: m.MessageKey,
		RoomKey:    m.RoomKey,
		UserId:     int64(m.UserId),
		AtTime:     int64(m.AtTime),
	}
	return r
}

//folding
var PB_MsgsSeenFromClient__FOlD = &PB_MsgsSeenFromClient{
	MessageKey: "",
	RoomKey:    "",
	UserId:     0,
	AtTime:     0,
}

type PB_MsgSeen_Flat struct {
	MessageKey string
	RoomKey    string
	UserId     int
	AtTime     int
}

//ToPB
func (m *PB_MsgSeen) ToFlat() *PB_MsgSeen_Flat {
	r := &PB_MsgSeen_Flat{
		MessageKey: m.MessageKey,
		RoomKey:    m.RoomKey,
		UserId:     int(m.UserId),
		AtTime:     int(m.AtTime),
	}
	return r
}

//ToPB
func (m *PB_MsgSeen_Flat) ToPB() *PB_MsgSeen {
	r := &PB_MsgSeen{
		MessageKey: m.MessageKey,
		RoomKey:    m.RoomKey,
		UserId:     int64(m.UserId),
		AtTime:     int64(m.AtTime),
	}
	return r
}

//folding
var PB_MsgSeen__FOlD = &PB_MsgSeen{
	MessageKey: "",
	RoomKey:    "",
	UserId:     0,
	AtTime:     0,
}

type PB_RequestMsgsSeen_Flat struct {
	Request PB_Request
	Seen    []PB_MsgSeen
}

//ToPB
func (m *PB_RequestMsgsSeen) ToFlat() *PB_RequestMsgsSeen_Flat {
	r := &PB_RequestMsgsSeen_Flat{}
	return r
}

//ToPB
func (m *PB_RequestMsgsSeen_Flat) ToPB() *PB_RequestMsgsSeen {
	r := &PB_RequestMsgsSeen{}
	return r
}

//folding
var PB_RequestMsgsSeen__FOlD = &PB_RequestMsgsSeen{}

type PB_ResponseMsgsSeen_Flat struct {
	Response PB_Response
}

//ToPB
func (m *PB_ResponseMsgsSeen) ToFlat() *PB_ResponseMsgsSeen_Flat {
	r := &PB_ResponseMsgsSeen_Flat{}
	return r
}

//ToPB
func (m *PB_ResponseMsgsSeen_Flat) ToPB() *PB_ResponseMsgsSeen {
	r := &PB_ResponseMsgsSeen{}
	return r
}

//folding
var PB_ResponseMsgsSeen__FOlD = &PB_ResponseMsgsSeen{}

type PB_ReqRpcAddMsg_Flat struct {
	Request PB_Request
	Message PB_Message
}

//ToPB
func (m *PB_ReqRpcAddMsg) ToFlat() *PB_ReqRpcAddMsg_Flat {
	r := &PB_ReqRpcAddMsg_Flat{}
	return r
}

//ToPB
func (m *PB_ReqRpcAddMsg_Flat) ToPB() *PB_ReqRpcAddMsg {
	r := &PB_ReqRpcAddMsg{}
	return r
}

//folding
var PB_ReqRpcAddMsg__FOlD = &PB_ReqRpcAddMsg{}

type PB_ResRpcAddMsg_Flat struct {
	Response  PB_Response
	ServerSrc string
}

//ToPB
func (m *PB_ResRpcAddMsg) ToFlat() *PB_ResRpcAddMsg_Flat {
	r := &PB_ResRpcAddMsg_Flat{

		ServerSrc: m.ServerSrc,
	}
	return r
}

//ToPB
func (m *PB_ResRpcAddMsg_Flat) ToPB() *PB_ResRpcAddMsg {
	r := &PB_ResRpcAddMsg{

		ServerSrc: m.ServerSrc,
	}
	return r
}

//folding
var PB_ResRpcAddMsg__FOlD = &PB_ResRpcAddMsg{

	ServerSrc: "",
}

type GeoLocation_Flat struct {
	Lat float64
	Lon float64
}

//ToPB
func (m *GeoLocation) ToFlat() *GeoLocation_Flat {
	r := &GeoLocation_Flat{
		Lat: float64(m.Lat),
		Lon: float64(m.Lon),
	}
	return r
}

//ToPB
func (m *GeoLocation_Flat) ToPB() *GeoLocation {
	r := &GeoLocation{
		Lat: m.Lat,
		Lon: m.Lon,
	}
	return r
}

//folding
var GeoLocation__FOlD = &GeoLocation{
	Lat: 0.0,
	Lon: 0.0,
}

type RoomMessageLog_Flat struct {
	typ          RoomMessageLogEnum
	TargetUserId int
	ByUserId     int
}

//ToPB
func (m *RoomMessageLog) ToFlat() *RoomMessageLog_Flat {
	r := &RoomMessageLog_Flat{

		TargetUserId: int(m.TargetUserId),
		ByUserId:     int(m.ByUserId),
	}
	return r
}

//ToPB
func (m *RoomMessageLog_Flat) ToPB() *RoomMessageLog {
	r := &RoomMessageLog{

		TargetUserId: uint64(m.TargetUserId),
		ByUserId:     uint64(m.ByUserId),
	}
	return r
}

//folding
var RoomMessageLog__FOlD = &RoomMessageLog{

	TargetUserId: 0,
	ByUserId:     0,
}

type RoomMessageForwardFrom_Flat struct {
	RoomId       int
	MessageId    int
	RoomTypeEnum int
}

//ToPB
func (m *RoomMessageForwardFrom) ToFlat() *RoomMessageForwardFrom_Flat {
	r := &RoomMessageForwardFrom_Flat{
		RoomId:       int(m.RoomId),
		MessageId:    int(m.MessageId),
		RoomTypeEnum: int(m.RoomTypeEnum),
	}
	return r
}

//ToPB
func (m *RoomMessageForwardFrom_Flat) ToPB() *RoomMessageForwardFrom {
	r := &RoomMessageForwardFrom{
		RoomId:       uint64(m.RoomId),
		MessageId:    uint64(m.MessageId),
		RoomTypeEnum: uint32(m.RoomTypeEnum),
	}
	return r
}

//folding
var RoomMessageForwardFrom__FOlD = &RoomMessageForwardFrom{
	RoomId:       0,
	MessageId:    0,
	RoomTypeEnum: 0,
}

type RoomDraft_Flat struct {
	Message string
	ReplyTo int
}

//ToPB
func (m *RoomDraft) ToFlat() *RoomDraft_Flat {
	r := &RoomDraft_Flat{
		Message: m.Message,
		ReplyTo: int(m.ReplyTo),
	}
	return r
}

//ToPB
func (m *RoomDraft_Flat) ToPB() *RoomDraft {
	r := &RoomDraft{
		Message: m.Message,
		ReplyTo: uint64(m.ReplyTo),
	}
	return r
}

//folding
var RoomDraft__FOlD = &RoomDraft{
	Message: "",
	ReplyTo: 0,
}

type ChatRoom_Flat struct {
}

//ToPB
func (m *ChatRoom) ToFlat() *ChatRoom_Flat {
	r := &ChatRoom_Flat{}
	return r
}

//ToPB
func (m *ChatRoom_Flat) ToPB() *ChatRoom {
	r := &ChatRoom{}
	return r
}

//folding
var ChatRoom__FOlD = &ChatRoom{}

type Pagination_Flat struct {
	Offset int
	Limit  int
}

//ToPB
func (m *Pagination) ToFlat() *Pagination_Flat {
	r := &Pagination_Flat{
		Offset: int(m.Offset),
		Limit:  int(m.Limit),
	}
	return r
}

//ToPB
func (m *Pagination_Flat) ToPB() *Pagination {
	r := &Pagination{
		Offset: uint32(m.Offset),
		Limit:  uint32(m.Limit),
	}
	return r
}

//folding
var Pagination__FOlD = &Pagination{
	Offset: 0,
	Limit:  0,
}

type PB_CommandToServer_Flat struct {
	ClientCallId   int
	Command        string
	RespondReached bool
	Data           []byte
}

//ToPB
func (m *PB_CommandToServer) ToFlat() *PB_CommandToServer_Flat {
	r := &PB_CommandToServer_Flat{
		ClientCallId:   int(m.ClientCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_CommandToServer_Flat) ToPB() *PB_CommandToServer {
	r := &PB_CommandToServer{
		ClientCallId:   int64(m.ClientCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           m.Data,
	}
	return r
}

//folding
var PB_CommandToServer__FOlD = &PB_CommandToServer{
	ClientCallId:   0,
	Command:        "",
	RespondReached: false,
	Data:           []byte{},
}

type PB_CommandToClient_Flat struct {
	ServerCallId   int
	Command        string
	RespondReached bool
	Data           []byte
}

//ToPB
func (m *PB_CommandToClient) ToFlat() *PB_CommandToClient_Flat {
	r := &PB_CommandToClient_Flat{
		ServerCallId:   int(m.ServerCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_CommandToClient_Flat) ToPB() *PB_CommandToClient {
	r := &PB_CommandToClient{
		ServerCallId:   int64(m.ServerCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           m.Data,
	}
	return r
}

//folding
var PB_CommandToClient__FOlD = &PB_CommandToClient{
	ServerCallId:   0,
	Command:        "",
	RespondReached: false,
	Data:           []byte{},
}

type PB_CommandReachedToServer_Flat struct {
	ClientCallId int
}

//ToPB
func (m *PB_CommandReachedToServer) ToFlat() *PB_CommandReachedToServer_Flat {
	r := &PB_CommandReachedToServer_Flat{
		ClientCallId: int(m.ClientCallId),
	}
	return r
}

//ToPB
func (m *PB_CommandReachedToServer_Flat) ToPB() *PB_CommandReachedToServer {
	r := &PB_CommandReachedToServer{
		ClientCallId: int64(m.ClientCallId),
	}
	return r
}

//folding
var PB_CommandReachedToServer__FOlD = &PB_CommandReachedToServer{
	ClientCallId: 0,
}

type PB_CommandReachedToClient_Flat struct {
	ServerCallId int
}

//ToPB
func (m *PB_CommandReachedToClient) ToFlat() *PB_CommandReachedToClient_Flat {
	r := &PB_CommandReachedToClient_Flat{
		ServerCallId: int(m.ServerCallId),
	}
	return r
}

//ToPB
func (m *PB_CommandReachedToClient_Flat) ToPB() *PB_CommandReachedToClient {
	r := &PB_CommandReachedToClient{
		ServerCallId: int64(m.ServerCallId),
	}
	return r
}

//folding
var PB_CommandReachedToClient__FOlD = &PB_CommandReachedToClient{
	ServerCallId: 0,
}

type PB_ResponseToClient_Flat struct {
	ClientCallId int
	PBClass      string
	RpcFullName  string
	Data         []byte
}

//ToPB
func (m *PB_ResponseToClient) ToFlat() *PB_ResponseToClient_Flat {
	r := &PB_ResponseToClient_Flat{
		ClientCallId: int(m.ClientCallId),
		PBClass:      m.PBClass,
		RpcFullName:  m.RpcFullName,
		Data:         []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_ResponseToClient_Flat) ToPB() *PB_ResponseToClient {
	r := &PB_ResponseToClient{
		ClientCallId: int64(m.ClientCallId),
		PBClass:      m.PBClass,
		RpcFullName:  m.RpcFullName,
		Data:         m.Data,
	}
	return r
}

//folding
var PB_ResponseToClient__FOlD = &PB_ResponseToClient{
	ClientCallId: 0,
	PBClass:      "",
	RpcFullName:  "",
	Data:         []byte{},
}

type PB_UserParam_CheckUserName2_Flat struct {
}

//ToPB
func (m *PB_UserParam_CheckUserName2) ToFlat() *PB_UserParam_CheckUserName2_Flat {
	r := &PB_UserParam_CheckUserName2_Flat{}
	return r
}

//ToPB
func (m *PB_UserParam_CheckUserName2_Flat) ToPB() *PB_UserParam_CheckUserName2 {
	r := &PB_UserParam_CheckUserName2{}
	return r
}

//folding
var PB_UserParam_CheckUserName2__FOlD = &PB_UserParam_CheckUserName2{}

type PB_UserResponse_CheckUserName2_Flat struct {
}

//ToPB
func (m *PB_UserResponse_CheckUserName2) ToFlat() *PB_UserResponse_CheckUserName2_Flat {
	r := &PB_UserResponse_CheckUserName2_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponse_CheckUserName2_Flat) ToPB() *PB_UserResponse_CheckUserName2 {
	r := &PB_UserResponse_CheckUserName2{}
	return r
}

//folding
var PB_UserResponse_CheckUserName2__FOlD = &PB_UserResponse_CheckUserName2{}

type PB_MsgParam_AddNewTextMessage_Flat struct {
	Text             string
	PeerId           int
	Time             int
	ReplyToMessageId int
	Forward          PB_MessageForwardedFrom
}

//ToPB
func (m *PB_MsgParam_AddNewTextMessage) ToFlat() *PB_MsgParam_AddNewTextMessage_Flat {
	r := &PB_MsgParam_AddNewTextMessage_Flat{
		Text:             m.Text,
		PeerId:           int(m.PeerId),
		Time:             int(m.Time),
		ReplyToMessageId: int(m.ReplyToMessageId),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_AddNewTextMessage_Flat) ToPB() *PB_MsgParam_AddNewTextMessage {
	r := &PB_MsgParam_AddNewTextMessage{
		Text:             m.Text,
		PeerId:           int32(m.PeerId),
		Time:             int32(m.Time),
		ReplyToMessageId: int64(m.ReplyToMessageId),
	}
	return r
}

//folding
var PB_MsgParam_AddNewTextMessage__FOlD = &PB_MsgParam_AddNewTextMessage{
	Text:             "",
	PeerId:           0,
	Time:             0,
	ReplyToMessageId: 0,
}

type PB_MsgResponse_AddNewTextMessage_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_AddNewTextMessage) ToFlat() *PB_MsgResponse_AddNewTextMessage_Flat {
	r := &PB_MsgResponse_AddNewTextMessage_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_AddNewTextMessage_Flat) ToPB() *PB_MsgResponse_AddNewTextMessage {
	r := &PB_MsgResponse_AddNewTextMessage{}
	return r
}

//folding
var PB_MsgResponse_AddNewTextMessage__FOlD = &PB_MsgResponse_AddNewTextMessage{}

type PB_MsgParam_AddNewMessage_Flat struct {
	Text             string
	PeerId           int
	Time             int
	ReplyToMessageId int
	Forward          PB_MessageForwardedFrom
	RoomMessageType  RoomMessageTypeEnum
	Blob             []byte
	MessageView      PB_MessageView
}

//ToPB
func (m *PB_MsgParam_AddNewMessage) ToFlat() *PB_MsgParam_AddNewMessage_Flat {
	r := &PB_MsgParam_AddNewMessage_Flat{
		Text:             m.Text,
		PeerId:           int(m.PeerId),
		Time:             int(m.Time),
		ReplyToMessageId: int(m.ReplyToMessageId),

		Blob: []byte(m.Blob),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_AddNewMessage_Flat) ToPB() *PB_MsgParam_AddNewMessage {
	r := &PB_MsgParam_AddNewMessage{
		Text:             m.Text,
		PeerId:           int32(m.PeerId),
		Time:             int32(m.Time),
		ReplyToMessageId: int64(m.ReplyToMessageId),

		Blob: m.Blob,
	}
	return r
}

//folding
var PB_MsgParam_AddNewMessage__FOlD = &PB_MsgParam_AddNewMessage{
	Text:             "",
	PeerId:           0,
	Time:             0,
	ReplyToMessageId: 0,

	Blob: []byte{},
}

type PB_MsgResponse_AddNewMessage_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_AddNewMessage) ToFlat() *PB_MsgResponse_AddNewMessage_Flat {
	r := &PB_MsgResponse_AddNewMessage_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_AddNewMessage_Flat) ToPB() *PB_MsgResponse_AddNewMessage {
	r := &PB_MsgResponse_AddNewMessage{}
	return r
}

//folding
var PB_MsgResponse_AddNewMessage__FOlD = &PB_MsgResponse_AddNewMessage{}

type PB_MsgParam_SetRoomActionDoing_Flat struct {
	GroupId       int
	DirectRoomKey string
	ActionType    RoomActionDoingEnum
}

//ToPB
func (m *PB_MsgParam_SetRoomActionDoing) ToFlat() *PB_MsgParam_SetRoomActionDoing_Flat {
	r := &PB_MsgParam_SetRoomActionDoing_Flat{
		GroupId:       int(m.GroupId),
		DirectRoomKey: m.DirectRoomKey,
	}
	return r
}

//ToPB
func (m *PB_MsgParam_SetRoomActionDoing_Flat) ToPB() *PB_MsgParam_SetRoomActionDoing {
	r := &PB_MsgParam_SetRoomActionDoing{
		GroupId:       int64(m.GroupId),
		DirectRoomKey: m.DirectRoomKey,
	}
	return r
}

//folding
var PB_MsgParam_SetRoomActionDoing__FOlD = &PB_MsgParam_SetRoomActionDoing{
	GroupId:       0,
	DirectRoomKey: "",
}

type PB_MsgResponse_SetRoomActionDoing_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_SetRoomActionDoing) ToFlat() *PB_MsgResponse_SetRoomActionDoing_Flat {
	r := &PB_MsgResponse_SetRoomActionDoing_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_SetRoomActionDoing_Flat) ToPB() *PB_MsgResponse_SetRoomActionDoing {
	r := &PB_MsgResponse_SetRoomActionDoing{}
	return r
}

//folding
var PB_MsgResponse_SetRoomActionDoing__FOlD = &PB_MsgResponse_SetRoomActionDoing{}

type PB_MsgParam_GetMessagesByIds_Flat struct {
	MessagesCollections PB_MessagesCollections
}

//ToPB
func (m *PB_MsgParam_GetMessagesByIds) ToFlat() *PB_MsgParam_GetMessagesByIds_Flat {
	r := &PB_MsgParam_GetMessagesByIds_Flat{}
	return r
}

//ToPB
func (m *PB_MsgParam_GetMessagesByIds_Flat) ToPB() *PB_MsgParam_GetMessagesByIds {
	r := &PB_MsgParam_GetMessagesByIds{}
	return r
}

//folding
var PB_MsgParam_GetMessagesByIds__FOlD = &PB_MsgParam_GetMessagesByIds{}

type PB_MsgResponse_GetMessagesByIds_Flat struct {
	MessagesViews []PB_MessageView
}

//ToPB
func (m *PB_MsgResponse_GetMessagesByIds) ToFlat() *PB_MsgResponse_GetMessagesByIds_Flat {
	r := &PB_MsgResponse_GetMessagesByIds_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_GetMessagesByIds_Flat) ToPB() *PB_MsgResponse_GetMessagesByIds {
	r := &PB_MsgResponse_GetMessagesByIds{}
	return r
}

//folding
var PB_MsgResponse_GetMessagesByIds__FOlD = &PB_MsgResponse_GetMessagesByIds{}

type PB_MsgParam_GetMessagesHistory_Flat struct {
	ChatKey string
	FromSeq int
	EndSeq  int
}

//ToPB
func (m *PB_MsgParam_GetMessagesHistory) ToFlat() *PB_MsgParam_GetMessagesHistory_Flat {
	r := &PB_MsgParam_GetMessagesHistory_Flat{
		ChatKey: m.ChatKey,
		FromSeq: int(m.FromSeq),
		EndSeq:  int(m.EndSeq),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_GetMessagesHistory_Flat) ToPB() *PB_MsgParam_GetMessagesHistory {
	r := &PB_MsgParam_GetMessagesHistory{
		ChatKey: m.ChatKey,
		FromSeq: int32(m.FromSeq),
		EndSeq:  int32(m.EndSeq),
	}
	return r
}

//folding
var PB_MsgParam_GetMessagesHistory__FOlD = &PB_MsgParam_GetMessagesHistory{
	ChatKey: "",
	FromSeq: 0,
	EndSeq:  0,
}

type PB_MsgResponse_GetMessagesHistory_Flat struct {
	MessagesViews []PB_MessageView
}

//ToPB
func (m *PB_MsgResponse_GetMessagesHistory) ToFlat() *PB_MsgResponse_GetMessagesHistory_Flat {
	r := &PB_MsgResponse_GetMessagesHistory_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_GetMessagesHistory_Flat) ToPB() *PB_MsgResponse_GetMessagesHistory {
	r := &PB_MsgResponse_GetMessagesHistory{}
	return r
}

//folding
var PB_MsgResponse_GetMessagesHistory__FOlD = &PB_MsgResponse_GetMessagesHistory{}

type PB_MsgParam_SetChatMessagesRangeAsSeen_Flat struct {
	ChatKey    string
	FromSeq    int
	EndSeq     int
	SeenTimeMs int
}

//ToPB
func (m *PB_MsgParam_SetChatMessagesRangeAsSeen) ToFlat() *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat {
	r := &PB_MsgParam_SetChatMessagesRangeAsSeen_Flat{
		ChatKey:    m.ChatKey,
		FromSeq:    int(m.FromSeq),
		EndSeq:     int(m.EndSeq),
		SeenTimeMs: int(m.SeenTimeMs),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat) ToPB() *PB_MsgParam_SetChatMessagesRangeAsSeen {
	r := &PB_MsgParam_SetChatMessagesRangeAsSeen{
		ChatKey:    m.ChatKey,
		FromSeq:    int32(m.FromSeq),
		EndSeq:     int32(m.EndSeq),
		SeenTimeMs: int64(m.SeenTimeMs),
	}
	return r
}

//folding
var PB_MsgParam_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgParam_SetChatMessagesRangeAsSeen{
	ChatKey:    "",
	FromSeq:    0,
	EndSeq:     0,
	SeenTimeMs: 0,
}

type PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_SetChatMessagesRangeAsSeen) ToFlat() *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat {
	r := &PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat) ToPB() *PB_MsgResponse_SetChatMessagesRangeAsSeen {
	r := &PB_MsgResponse_SetChatMessagesRangeAsSeen{}
	return r
}

//folding
var PB_MsgResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgResponse_SetChatMessagesRangeAsSeen{}

type PB_MsgParam_DeleteChatHistory_Flat struct {
	ChatKey string
	ToSeq   int
}

//ToPB
func (m *PB_MsgParam_DeleteChatHistory) ToFlat() *PB_MsgParam_DeleteChatHistory_Flat {
	r := &PB_MsgParam_DeleteChatHistory_Flat{
		ChatKey: m.ChatKey,
		ToSeq:   int(m.ToSeq),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_DeleteChatHistory_Flat) ToPB() *PB_MsgParam_DeleteChatHistory {
	r := &PB_MsgParam_DeleteChatHistory{
		ChatKey: m.ChatKey,
		ToSeq:   int32(m.ToSeq),
	}
	return r
}

//folding
var PB_MsgParam_DeleteChatHistory__FOlD = &PB_MsgParam_DeleteChatHistory{
	ChatKey: "",
	ToSeq:   0,
}

type PB_MsgResponse_DeleteChatHistory_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_DeleteChatHistory) ToFlat() *PB_MsgResponse_DeleteChatHistory_Flat {
	r := &PB_MsgResponse_DeleteChatHistory_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_DeleteChatHistory_Flat) ToPB() *PB_MsgResponse_DeleteChatHistory {
	r := &PB_MsgResponse_DeleteChatHistory{}
	return r
}

//folding
var PB_MsgResponse_DeleteChatHistory__FOlD = &PB_MsgResponse_DeleteChatHistory{}

type PB_MsgParam_DeleteMessagesByIds_Flat struct {
	MessagesCollections PB_MessagesCollections
}

//ToPB
func (m *PB_MsgParam_DeleteMessagesByIds) ToFlat() *PB_MsgParam_DeleteMessagesByIds_Flat {
	r := &PB_MsgParam_DeleteMessagesByIds_Flat{}
	return r
}

//ToPB
func (m *PB_MsgParam_DeleteMessagesByIds_Flat) ToPB() *PB_MsgParam_DeleteMessagesByIds {
	r := &PB_MsgParam_DeleteMessagesByIds{}
	return r
}

//folding
var PB_MsgParam_DeleteMessagesByIds__FOlD = &PB_MsgParam_DeleteMessagesByIds{}

type PB_MsgResponse_DeleteMessagesByIds_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_DeleteMessagesByIds) ToFlat() *PB_MsgResponse_DeleteMessagesByIds_Flat {
	r := &PB_MsgResponse_DeleteMessagesByIds_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_DeleteMessagesByIds_Flat) ToPB() *PB_MsgResponse_DeleteMessagesByIds {
	r := &PB_MsgResponse_DeleteMessagesByIds{}
	return r
}

//folding
var PB_MsgResponse_DeleteMessagesByIds__FOlD = &PB_MsgResponse_DeleteMessagesByIds{}

type PB_MsgParam_SetMessagesAsReceived_Flat struct {
	MessagesCollections PB_MessagesCollections
}

//ToPB
func (m *PB_MsgParam_SetMessagesAsReceived) ToFlat() *PB_MsgParam_SetMessagesAsReceived_Flat {
	r := &PB_MsgParam_SetMessagesAsReceived_Flat{}
	return r
}

//ToPB
func (m *PB_MsgParam_SetMessagesAsReceived_Flat) ToPB() *PB_MsgParam_SetMessagesAsReceived {
	r := &PB_MsgParam_SetMessagesAsReceived{}
	return r
}

//folding
var PB_MsgParam_SetMessagesAsReceived__FOlD = &PB_MsgParam_SetMessagesAsReceived{}

type PB_MsgResponse_SetMessagesAsReceived_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_SetMessagesAsReceived) ToFlat() *PB_MsgResponse_SetMessagesAsReceived_Flat {
	r := &PB_MsgResponse_SetMessagesAsReceived_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_SetMessagesAsReceived_Flat) ToPB() *PB_MsgResponse_SetMessagesAsReceived {
	r := &PB_MsgResponse_SetMessagesAsReceived{}
	return r
}

//folding
var PB_MsgResponse_SetMessagesAsReceived__FOlD = &PB_MsgResponse_SetMessagesAsReceived{}

type PB_MsgParam_ForwardMessages_Flat struct {
	MessagesCollections PB_MessagesCollections
	ToDirectChatKeys    []string
	ToGroupChatIds      []int
}

//ToPB
func (m *PB_MsgParam_ForwardMessages) ToFlat() *PB_MsgParam_ForwardMessages_Flat {
	r := &PB_MsgParam_ForwardMessages_Flat{

		ToDirectChatKeys: m.ToDirectChatKeys,
		ToGroupChatIds:   helper.SliceInt64ToInt(m.ToGroupChatIds),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_ForwardMessages_Flat) ToPB() *PB_MsgParam_ForwardMessages {
	r := &PB_MsgParam_ForwardMessages{

		ToDirectChatKeys: m.ToDirectChatKeys,
		ToGroupChatIds:   helper.SliceIntToInt64(m.ToGroupChatIds),
	}
	return r
}

//folding
var PB_MsgParam_ForwardMessages__FOlD = &PB_MsgParam_ForwardMessages{}

type PB_MsgResponse_ForwardMessages_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_ForwardMessages) ToFlat() *PB_MsgResponse_ForwardMessages_Flat {
	r := &PB_MsgResponse_ForwardMessages_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_ForwardMessages_Flat) ToPB() *PB_MsgResponse_ForwardMessages {
	r := &PB_MsgResponse_ForwardMessages{}
	return r
}

//folding
var PB_MsgResponse_ForwardMessages__FOlD = &PB_MsgResponse_ForwardMessages{}

type PB_MsgParam_EditMessage_Flat struct {
	ChatKey   string
	RoomType  RoomTypeEnum
	MessageId int
	NewText   string
}

//ToPB
func (m *PB_MsgParam_EditMessage) ToFlat() *PB_MsgParam_EditMessage_Flat {
	r := &PB_MsgParam_EditMessage_Flat{
		ChatKey: m.ChatKey,

		MessageId: int(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

//ToPB
func (m *PB_MsgParam_EditMessage_Flat) ToPB() *PB_MsgParam_EditMessage {
	r := &PB_MsgParam_EditMessage{
		ChatKey: m.ChatKey,

		MessageId: int64(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

//folding
var PB_MsgParam_EditMessage__FOlD = &PB_MsgParam_EditMessage{
	ChatKey: "",

	MessageId: 0,
	NewText:   "",
}

type PB_MsgResponse_EditMessage_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_EditMessage) ToFlat() *PB_MsgResponse_EditMessage_Flat {
	r := &PB_MsgResponse_EditMessage_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_EditMessage_Flat) ToPB() *PB_MsgResponse_EditMessage {
	r := &PB_MsgResponse_EditMessage{}
	return r
}

//folding
var PB_MsgResponse_EditMessage__FOlD = &PB_MsgResponse_EditMessage{}

type PB_MsgParam_BroadcastNewMessage_Flat struct {
	Text    string
	PeerId  int
	Time    int
	Forward PB_MessageForwardedFrom
}

//ToPB
func (m *PB_MsgParam_BroadcastNewMessage) ToFlat() *PB_MsgParam_BroadcastNewMessage_Flat {
	r := &PB_MsgParam_BroadcastNewMessage_Flat{
		Text:   m.Text,
		PeerId: int(m.PeerId),
		Time:   int(m.Time),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_BroadcastNewMessage_Flat) ToPB() *PB_MsgParam_BroadcastNewMessage {
	r := &PB_MsgParam_BroadcastNewMessage{
		Text:   m.Text,
		PeerId: int32(m.PeerId),
		Time:   int32(m.Time),
	}
	return r
}

//folding
var PB_MsgParam_BroadcastNewMessage__FOlD = &PB_MsgParam_BroadcastNewMessage{
	Text:   "",
	PeerId: 0,
	Time:   0,
}

type PB_MsgResponse_BroadcastNewMessage_Flat struct {
}

//ToPB
func (m *PB_MsgResponse_BroadcastNewMessage) ToFlat() *PB_MsgResponse_BroadcastNewMessage_Flat {
	r := &PB_MsgResponse_BroadcastNewMessage_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_BroadcastNewMessage_Flat) ToPB() *PB_MsgResponse_BroadcastNewMessage {
	r := &PB_MsgResponse_BroadcastNewMessage{}
	return r
}

//folding
var PB_MsgResponse_BroadcastNewMessage__FOlD = &PB_MsgResponse_BroadcastNewMessage{}

type PB_MsgParam_GetFreshChatList_Flat struct {
}

//ToPB
func (m *PB_MsgParam_GetFreshChatList) ToFlat() *PB_MsgParam_GetFreshChatList_Flat {
	r := &PB_MsgParam_GetFreshChatList_Flat{}
	return r
}

//ToPB
func (m *PB_MsgParam_GetFreshChatList_Flat) ToPB() *PB_MsgParam_GetFreshChatList {
	r := &PB_MsgParam_GetFreshChatList{}
	return r
}

//folding
var PB_MsgParam_GetFreshChatList__FOlD = &PB_MsgParam_GetFreshChatList{}

type PB_MsgResponse_GetFreshChatList_Flat struct {
	Chats []PB_ChatView
	Users []PB_UserView
}

//ToPB
func (m *PB_MsgResponse_GetFreshChatList) ToFlat() *PB_MsgResponse_GetFreshChatList_Flat {
	r := &PB_MsgResponse_GetFreshChatList_Flat{}
	return r
}

//ToPB
func (m *PB_MsgResponse_GetFreshChatList_Flat) ToPB() *PB_MsgResponse_GetFreshChatList {
	r := &PB_MsgResponse_GetFreshChatList{}
	return r
}

//folding
var PB_MsgResponse_GetFreshChatList__FOlD = &PB_MsgResponse_GetFreshChatList{}

type PB_MsgParam_GetFreshRoomMessagesList_Flat struct {
	ChatKey string
	RoomKey string
	Last    int
}

//ToPB
func (m *PB_MsgParam_GetFreshRoomMessagesList) ToFlat() *PB_MsgParam_GetFreshRoomMessagesList_Flat {
	r := &PB_MsgParam_GetFreshRoomMessagesList_Flat{
		ChatKey: m.ChatKey,
		RoomKey: m.RoomKey,
		Last:    int(m.Last),
	}
	return r
}

//ToPB
func (m *PB_MsgParam_GetFreshRoomMessagesList_Flat) ToPB() *PB_MsgParam_GetFreshRoomMessagesList {
	r := &PB_MsgParam_GetFreshRoomMessagesList{
		ChatKey: m.ChatKey,
		RoomKey: m.RoomKey,
		Last:    int64(m.Last),
	}
	return r
}

//folding
var PB_MsgParam_GetFreshRoomMessagesList__FOlD = &PB_MsgParam_GetFreshRoomMessagesList{
	ChatKey: "",
	RoomKey: "",
	Last:    0,
}

type PB_MsgResponse_GetFreshRoomMessagesList_Flat struct {
	Messages []PB_MessageView
	HasMore  bool
}

//ToPB
func (m *PB_MsgResponse_GetFreshRoomMessagesList) ToFlat() *PB_MsgResponse_GetFreshRoomMessagesList_Flat {
	r := &PB_MsgResponse_GetFreshRoomMessagesList_Flat{

		HasMore: m.HasMore,
	}
	return r
}

//ToPB
func (m *PB_MsgResponse_GetFreshRoomMessagesList_Flat) ToPB() *PB_MsgResponse_GetFreshRoomMessagesList {
	r := &PB_MsgResponse_GetFreshRoomMessagesList{

		HasMore: m.HasMore,
	}
	return r
}

//folding
var PB_MsgResponse_GetFreshRoomMessagesList__FOlD = &PB_MsgResponse_GetFreshRoomMessagesList{

	HasMore: false,
}

type PB_MsgParam_Echo_Flat struct {
	Text string
}

//ToPB
func (m *PB_MsgParam_Echo) ToFlat() *PB_MsgParam_Echo_Flat {
	r := &PB_MsgParam_Echo_Flat{
		Text: m.Text,
	}
	return r
}

//ToPB
func (m *PB_MsgParam_Echo_Flat) ToPB() *PB_MsgParam_Echo {
	r := &PB_MsgParam_Echo{
		Text: m.Text,
	}
	return r
}

//folding
var PB_MsgParam_Echo__FOlD = &PB_MsgParam_Echo{
	Text: "",
}

type PB_MsgResponse_PB_MsgParam_Echo_Flat struct {
	Text string
}

//ToPB
func (m *PB_MsgResponse_PB_MsgParam_Echo) ToFlat() *PB_MsgResponse_PB_MsgParam_Echo_Flat {
	r := &PB_MsgResponse_PB_MsgParam_Echo_Flat{
		Text: m.Text,
	}
	return r
}

//ToPB
func (m *PB_MsgResponse_PB_MsgParam_Echo_Flat) ToPB() *PB_MsgResponse_PB_MsgParam_Echo {
	r := &PB_MsgResponse_PB_MsgParam_Echo{
		Text: m.Text,
	}
	return r
}

//folding
var PB_MsgResponse_PB_MsgParam_Echo__FOlD = &PB_MsgResponse_PB_MsgParam_Echo{
	Text: "",
}

type PB_SyncParam_GetDirectUpdates_Flat struct {
	LastId int
}

//ToPB
func (m *PB_SyncParam_GetDirectUpdates) ToFlat() *PB_SyncParam_GetDirectUpdates_Flat {
	r := &PB_SyncParam_GetDirectUpdates_Flat{
		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncParam_GetDirectUpdates_Flat) ToPB() *PB_SyncParam_GetDirectUpdates {
	r := &PB_SyncParam_GetDirectUpdates{
		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncParam_GetDirectUpdates__FOlD = &PB_SyncParam_GetDirectUpdates{
	LastId: 0,
}

type PB_SyncResponse_GetDirectUpdates_Flat struct {
	NewMessages               []PB_MessageView
	ChatFiles                 []PB_MessageFileView
	Chats                     []PB_ChatView
	Users                     []PB_UserView
	MessagesChangeIds         []PB_UpdateMessageId
	MessagesToUpdate          []PB_UpdateMessageToEdit
	MessagesToDelete          []PB_UpdateMessageToDelete
	MessagesDelivierdToServer []PB_UpdateMessageMeta
	MessagesDelivierdToPeer   []PB_UpdateMessageMeta
	MessagesSeenByPeer        []PB_UpdateMessageMeta
	MessagesDeletedFromServer []PB_UpdateMessageMeta
	RoomActionDoing           []PB_UpdateRoomActionDoing
	LastId                    int
}

//ToPB
func (m *PB_SyncResponse_GetDirectUpdates) ToFlat() *PB_SyncResponse_GetDirectUpdates_Flat {
	r := &PB_SyncResponse_GetDirectUpdates_Flat{

		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncResponse_GetDirectUpdates_Flat) ToPB() *PB_SyncResponse_GetDirectUpdates {
	r := &PB_SyncResponse_GetDirectUpdates{

		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncResponse_GetDirectUpdates__FOlD = &PB_SyncResponse_GetDirectUpdates{

	LastId: 0,
}

type PB_SyncParam_GetGeneralUpdates_Flat struct {
	LastId int
}

//ToPB
func (m *PB_SyncParam_GetGeneralUpdates) ToFlat() *PB_SyncParam_GetGeneralUpdates_Flat {
	r := &PB_SyncParam_GetGeneralUpdates_Flat{
		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncParam_GetGeneralUpdates_Flat) ToPB() *PB_SyncParam_GetGeneralUpdates {
	r := &PB_SyncParam_GetGeneralUpdates{
		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncParam_GetGeneralUpdates__FOlD = &PB_SyncParam_GetGeneralUpdates{
	LastId: 0,
}

type PB_SyncResponse_GetGeneralUpdates_Flat struct {
	UserBlockedByMe []PB_UpdateUserBlocked
	UserBlockedMe   []PB_UpdateUserBlocked
	LastId          int
}

//ToPB
func (m *PB_SyncResponse_GetGeneralUpdates) ToFlat() *PB_SyncResponse_GetGeneralUpdates_Flat {
	r := &PB_SyncResponse_GetGeneralUpdates_Flat{

		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncResponse_GetGeneralUpdates_Flat) ToPB() *PB_SyncResponse_GetGeneralUpdates {
	r := &PB_SyncResponse_GetGeneralUpdates{

		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncResponse_GetGeneralUpdates__FOlD = &PB_SyncResponse_GetGeneralUpdates{

	LastId: 0,
}

type PB_SyncParam_GetNotifyUpdates_Flat struct {
	LastId int
}

//ToPB
func (m *PB_SyncParam_GetNotifyUpdates) ToFlat() *PB_SyncParam_GetNotifyUpdates_Flat {
	r := &PB_SyncParam_GetNotifyUpdates_Flat{
		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncParam_GetNotifyUpdates_Flat) ToPB() *PB_SyncParam_GetNotifyUpdates {
	r := &PB_SyncParam_GetNotifyUpdates{
		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncParam_GetNotifyUpdates__FOlD = &PB_SyncParam_GetNotifyUpdates{
	LastId: 0,
}

type PB_SyncResponse_GetNotifyUpdates_Flat struct {
	Updates PB_NotifyUpdatesView
	LastId  int
}

//ToPB
func (m *PB_SyncResponse_GetNotifyUpdates) ToFlat() *PB_SyncResponse_GetNotifyUpdates_Flat {
	r := &PB_SyncResponse_GetNotifyUpdates_Flat{

		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncResponse_GetNotifyUpdates_Flat) ToPB() *PB_SyncResponse_GetNotifyUpdates {
	r := &PB_SyncResponse_GetNotifyUpdates{

		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncResponse_GetNotifyUpdates__FOlD = &PB_SyncResponse_GetNotifyUpdates{

	LastId: 0,
}

type PB_SyncParam_SetLastSyncDirectUpdateId_Flat struct {
	LastId int
}

//ToPB
func (m *PB_SyncParam_SetLastSyncDirectUpdateId) ToFlat() *PB_SyncParam_SetLastSyncDirectUpdateId_Flat {
	r := &PB_SyncParam_SetLastSyncDirectUpdateId_Flat{
		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncParam_SetLastSyncDirectUpdateId_Flat) ToPB() *PB_SyncParam_SetLastSyncDirectUpdateId {
	r := &PB_SyncParam_SetLastSyncDirectUpdateId{
		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncParam_SetLastSyncDirectUpdateId__FOlD = &PB_SyncParam_SetLastSyncDirectUpdateId{
	LastId: 0,
}

type PB_SyncResponse_SetLastSyncDirectUpdateId_Flat struct {
}

//ToPB
func (m *PB_SyncResponse_SetLastSyncDirectUpdateId) ToFlat() *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat {
	r := &PB_SyncResponse_SetLastSyncDirectUpdateId_Flat{}
	return r
}

//ToPB
func (m *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat) ToPB() *PB_SyncResponse_SetLastSyncDirectUpdateId {
	r := &PB_SyncResponse_SetLastSyncDirectUpdateId{}
	return r
}

//folding
var PB_SyncResponse_SetLastSyncDirectUpdateId__FOlD = &PB_SyncResponse_SetLastSyncDirectUpdateId{}

type PB_SyncParam_SetLastSyncGeneralUpdateId_Flat struct {
	LastId int
}

//ToPB
func (m *PB_SyncParam_SetLastSyncGeneralUpdateId) ToFlat() *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat {
	r := &PB_SyncParam_SetLastSyncGeneralUpdateId_Flat{
		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat) ToPB() *PB_SyncParam_SetLastSyncGeneralUpdateId {
	r := &PB_SyncParam_SetLastSyncGeneralUpdateId{
		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncParam_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncParam_SetLastSyncGeneralUpdateId{
	LastId: 0,
}

type PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat struct {
}

//ToPB
func (m *PB_SyncResponse_SetLastSyncGeneralUpdateId) ToFlat() *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat {
	r := &PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat{}
	return r
}

//ToPB
func (m *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat) ToPB() *PB_SyncResponse_SetLastSyncGeneralUpdateId {
	r := &PB_SyncResponse_SetLastSyncGeneralUpdateId{}
	return r
}

//folding
var PB_SyncResponse_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncResponse_SetLastSyncGeneralUpdateId{}

type PB_SyncParam_SetLastSyncNotifyUpdateId_Flat struct {
	LastId int
}

//ToPB
func (m *PB_SyncParam_SetLastSyncNotifyUpdateId) ToFlat() *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat {
	r := &PB_SyncParam_SetLastSyncNotifyUpdateId_Flat{
		LastId: int(m.LastId),
	}
	return r
}

//ToPB
func (m *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat) ToPB() *PB_SyncParam_SetLastSyncNotifyUpdateId {
	r := &PB_SyncParam_SetLastSyncNotifyUpdateId{
		LastId: int64(m.LastId),
	}
	return r
}

//folding
var PB_SyncParam_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncParam_SetLastSyncNotifyUpdateId{
	LastId: 0,
}

type PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat struct {
}

//ToPB
func (m *PB_SyncResponse_SetLastSyncNotifyUpdateId) ToFlat() *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat {
	r := &PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat{}
	return r
}

//ToPB
func (m *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat) ToPB() *PB_SyncResponse_SetLastSyncNotifyUpdateId {
	r := &PB_SyncResponse_SetLastSyncNotifyUpdateId{}
	return r
}

//folding
var PB_SyncResponse_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncResponse_SetLastSyncNotifyUpdateId{}

type PB_UserParam_BlockUser_Flat struct {
	UserId   int
	UserName string
}

//ToPB
func (m *PB_UserParam_BlockUser) ToFlat() *PB_UserParam_BlockUser_Flat {
	r := &PB_UserParam_BlockUser_Flat{
		UserId:   int(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//ToPB
func (m *PB_UserParam_BlockUser_Flat) ToPB() *PB_UserParam_BlockUser {
	r := &PB_UserParam_BlockUser{
		UserId:   int32(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//folding
var PB_UserParam_BlockUser__FOlD = &PB_UserParam_BlockUser{
	UserId:   0,
	UserName: "",
}

type PB_UserOfflineResponse_BlockUser_Flat struct {
	ByUserId       int
	TargetUserId   int
	TargetUserName string
}

//ToPB
func (m *PB_UserOfflineResponse_BlockUser) ToFlat() *PB_UserOfflineResponse_BlockUser_Flat {
	r := &PB_UserOfflineResponse_BlockUser_Flat{
		ByUserId:       int(m.ByUserId),
		TargetUserId:   int(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//ToPB
func (m *PB_UserOfflineResponse_BlockUser_Flat) ToPB() *PB_UserOfflineResponse_BlockUser {
	r := &PB_UserOfflineResponse_BlockUser{
		ByUserId:       int32(m.ByUserId),
		TargetUserId:   int32(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//folding
var PB_UserOfflineResponse_BlockUser__FOlD = &PB_UserOfflineResponse_BlockUser{
	ByUserId:       0,
	TargetUserId:   0,
	TargetUserName: "",
}

type PB_UserParam_UnBlockUser_Flat struct {
	Offset int
	Limit  int
}

//ToPB
func (m *PB_UserParam_UnBlockUser) ToFlat() *PB_UserParam_UnBlockUser_Flat {
	r := &PB_UserParam_UnBlockUser_Flat{
		Offset: int(m.Offset),
		Limit:  int(m.Limit),
	}
	return r
}

//ToPB
func (m *PB_UserParam_UnBlockUser_Flat) ToPB() *PB_UserParam_UnBlockUser {
	r := &PB_UserParam_UnBlockUser{
		Offset: int32(m.Offset),
		Limit:  int32(m.Limit),
	}
	return r
}

//folding
var PB_UserParam_UnBlockUser__FOlD = &PB_UserParam_UnBlockUser{
	Offset: 0,
	Limit:  0,
}

type PB_UserOfflineResponse_UnBlockUser_Flat struct {
	Users []UserView
}

//ToPB
func (m *PB_UserOfflineResponse_UnBlockUser) ToFlat() *PB_UserOfflineResponse_UnBlockUser_Flat {
	r := &PB_UserOfflineResponse_UnBlockUser_Flat{}
	return r
}

//ToPB
func (m *PB_UserOfflineResponse_UnBlockUser_Flat) ToPB() *PB_UserOfflineResponse_UnBlockUser {
	r := &PB_UserOfflineResponse_UnBlockUser{}
	return r
}

//folding
var PB_UserOfflineResponse_UnBlockUser__FOlD = &PB_UserOfflineResponse_UnBlockUser{}

type PB_UserParam_BlockedList_Flat struct {
	UserId   int
	UserName string
}

//ToPB
func (m *PB_UserParam_BlockedList) ToFlat() *PB_UserParam_BlockedList_Flat {
	r := &PB_UserParam_BlockedList_Flat{
		UserId:   int(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//ToPB
func (m *PB_UserParam_BlockedList_Flat) ToPB() *PB_UserParam_BlockedList {
	r := &PB_UserParam_BlockedList{
		UserId:   int32(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//folding
var PB_UserParam_BlockedList__FOlD = &PB_UserParam_BlockedList{
	UserId:   0,
	UserName: "",
}

type PB_UserResponse_BlockedList_Flat struct {
	ByUserId       int
	TargetUserId   int
	TargetUserName string
}

//ToPB
func (m *PB_UserResponse_BlockedList) ToFlat() *PB_UserResponse_BlockedList_Flat {
	r := &PB_UserResponse_BlockedList_Flat{
		ByUserId:       int(m.ByUserId),
		TargetUserId:   int(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//ToPB
func (m *PB_UserResponse_BlockedList_Flat) ToPB() *PB_UserResponse_BlockedList {
	r := &PB_UserResponse_BlockedList{
		ByUserId:       int32(m.ByUserId),
		TargetUserId:   int32(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//folding
var PB_UserResponse_BlockedList__FOlD = &PB_UserResponse_BlockedList{
	ByUserId:       0,
	TargetUserId:   0,
	TargetUserName: "",
}

type PB_UserParam_UpdateAbout_Flat struct {
	NewAbout string
}

//ToPB
func (m *PB_UserParam_UpdateAbout) ToFlat() *PB_UserParam_UpdateAbout_Flat {
	r := &PB_UserParam_UpdateAbout_Flat{
		NewAbout: m.NewAbout,
	}
	return r
}

//ToPB
func (m *PB_UserParam_UpdateAbout_Flat) ToPB() *PB_UserParam_UpdateAbout {
	r := &PB_UserParam_UpdateAbout{
		NewAbout: m.NewAbout,
	}
	return r
}

//folding
var PB_UserParam_UpdateAbout__FOlD = &PB_UserParam_UpdateAbout{
	NewAbout: "",
}

type PB_UserOfflineResponse_UpdateAbout_Flat struct {
	UserId   int
	NewAbout string
}

//ToPB
func (m *PB_UserOfflineResponse_UpdateAbout) ToFlat() *PB_UserOfflineResponse_UpdateAbout_Flat {
	r := &PB_UserOfflineResponse_UpdateAbout_Flat{
		UserId:   int(m.UserId),
		NewAbout: m.NewAbout,
	}
	return r
}

//ToPB
func (m *PB_UserOfflineResponse_UpdateAbout_Flat) ToPB() *PB_UserOfflineResponse_UpdateAbout {
	r := &PB_UserOfflineResponse_UpdateAbout{
		UserId:   int32(m.UserId),
		NewAbout: m.NewAbout,
	}
	return r
}

//folding
var PB_UserOfflineResponse_UpdateAbout__FOlD = &PB_UserOfflineResponse_UpdateAbout{
	UserId:   0,
	NewAbout: "",
}

type PB_UserParam_UpdateUserName_Flat struct {
	NewUserName string
}

//ToPB
func (m *PB_UserParam_UpdateUserName) ToFlat() *PB_UserParam_UpdateUserName_Flat {
	r := &PB_UserParam_UpdateUserName_Flat{
		NewUserName: m.NewUserName,
	}
	return r
}

//ToPB
func (m *PB_UserParam_UpdateUserName_Flat) ToPB() *PB_UserParam_UpdateUserName {
	r := &PB_UserParam_UpdateUserName{
		NewUserName: m.NewUserName,
	}
	return r
}

//folding
var PB_UserParam_UpdateUserName__FOlD = &PB_UserParam_UpdateUserName{
	NewUserName: "",
}

type PB_UserOfflineResponse_UpdateUserName_Flat struct {
	UserId      int
	NewUserName string
}

//ToPB
func (m *PB_UserOfflineResponse_UpdateUserName) ToFlat() *PB_UserOfflineResponse_UpdateUserName_Flat {
	r := &PB_UserOfflineResponse_UpdateUserName_Flat{
		UserId:      int(m.UserId),
		NewUserName: m.NewUserName,
	}
	return r
}

//ToPB
func (m *PB_UserOfflineResponse_UpdateUserName_Flat) ToPB() *PB_UserOfflineResponse_UpdateUserName {
	r := &PB_UserOfflineResponse_UpdateUserName{
		UserId:      int32(m.UserId),
		NewUserName: m.NewUserName,
	}
	return r
}

//folding
var PB_UserOfflineResponse_UpdateUserName__FOlD = &PB_UserOfflineResponse_UpdateUserName{
	UserId:      0,
	NewUserName: "",
}

type PB_UserParam_ChangeAvatar_Flat struct {
	None       bool
	ImageData2 []byte
}

//ToPB
func (m *PB_UserParam_ChangeAvatar) ToFlat() *PB_UserParam_ChangeAvatar_Flat {
	r := &PB_UserParam_ChangeAvatar_Flat{
		None:       m.None,
		ImageData2: []byte(m.ImageData2),
	}
	return r
}

//ToPB
func (m *PB_UserParam_ChangeAvatar_Flat) ToPB() *PB_UserParam_ChangeAvatar {
	r := &PB_UserParam_ChangeAvatar{
		None:       m.None,
		ImageData2: m.ImageData2,
	}
	return r
}

//folding
var PB_UserParam_ChangeAvatar__FOlD = &PB_UserParam_ChangeAvatar{
	None:       false,
	ImageData2: []byte{},
}

type PB_UserOfflineResponse_ChangeAvatar_Flat struct {
}

//ToPB
func (m *PB_UserOfflineResponse_ChangeAvatar) ToFlat() *PB_UserOfflineResponse_ChangeAvatar_Flat {
	r := &PB_UserOfflineResponse_ChangeAvatar_Flat{}
	return r
}

//ToPB
func (m *PB_UserOfflineResponse_ChangeAvatar_Flat) ToPB() *PB_UserOfflineResponse_ChangeAvatar {
	r := &PB_UserOfflineResponse_ChangeAvatar{}
	return r
}

//folding
var PB_UserOfflineResponse_ChangeAvatar__FOlD = &PB_UserOfflineResponse_ChangeAvatar{}

type PB_UserParam_ChangePrivacy_Flat struct {
	Level ProfilePrivacyLevelEnum
}

//ToPB
func (m *PB_UserParam_ChangePrivacy) ToFlat() *PB_UserParam_ChangePrivacy_Flat {
	r := &PB_UserParam_ChangePrivacy_Flat{}
	return r
}

//ToPB
func (m *PB_UserParam_ChangePrivacy_Flat) ToPB() *PB_UserParam_ChangePrivacy {
	r := &PB_UserParam_ChangePrivacy{}
	return r
}

//folding
var PB_UserParam_ChangePrivacy__FOlD = &PB_UserParam_ChangePrivacy{}

type PB_UserResponseOffline_ChangePrivacy_Flat struct {
}

//ToPB
func (m *PB_UserResponseOffline_ChangePrivacy) ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
	r := &PB_UserResponseOffline_ChangePrivacy_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponseOffline_ChangePrivacy_Flat) ToPB() *PB_UserResponseOffline_ChangePrivacy {
	r := &PB_UserResponseOffline_ChangePrivacy{}
	return r
}

//folding
var PB_UserResponseOffline_ChangePrivacy__FOlD = &PB_UserResponseOffline_ChangePrivacy{}

type PB_UserParam_CheckUserName_Flat struct {
	Level ProfilePrivacyLevelEnum
}

//ToPB
func (m *PB_UserParam_CheckUserName) ToFlat() *PB_UserParam_CheckUserName_Flat {
	r := &PB_UserParam_CheckUserName_Flat{}
	return r
}

//ToPB
func (m *PB_UserParam_CheckUserName_Flat) ToPB() *PB_UserParam_CheckUserName {
	r := &PB_UserParam_CheckUserName{}
	return r
}

//folding
var PB_UserParam_CheckUserName__FOlD = &PB_UserParam_CheckUserName{}

type PB_UserResponse_CheckUserName_Flat struct {
}

//ToPB
func (m *PB_UserResponse_CheckUserName) ToFlat() *PB_UserResponse_CheckUserName_Flat {
	r := &PB_UserResponse_CheckUserName_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponse_CheckUserName_Flat) ToPB() *PB_UserResponse_CheckUserName {
	r := &PB_UserResponse_CheckUserName{}
	return r
}

//folding
var PB_UserResponse_CheckUserName__FOlD = &PB_UserResponse_CheckUserName{}

type UserView_Flat struct {
}

//ToPB
func (m *UserView) ToFlat() *UserView_Flat {
	r := &UserView_Flat{}
	return r
}

//ToPB
func (m *UserView_Flat) ToPB() *UserView {
	r := &UserView{}
	return r
}

//folding
var UserView__FOlD = &UserView{}

type PB_Chat_Flat struct {
	ChatId         int
	ChatKey        string
	RoomTypeEnumId int
	UserId         int
	LastSeqSeen    int
	LastSeqDelete  int
	PeerUserId     int
	GroupId        int
	CreatedTime    int
	CurrentSeq     int
	UpdatedMs      int
}

//ToPB
func (m *PB_Chat) ToFlat() *PB_Chat_Flat {
	r := &PB_Chat_Flat{
		ChatId:         int(m.ChatId),
		ChatKey:        m.ChatKey,
		RoomTypeEnumId: int(m.RoomTypeEnumId),
		UserId:         int(m.UserId),
		LastSeqSeen:    int(m.LastSeqSeen),
		LastSeqDelete:  int(m.LastSeqDelete),
		PeerUserId:     int(m.PeerUserId),
		GroupId:        int(m.GroupId),
		CreatedTime:    int(m.CreatedTime),
		CurrentSeq:     int(m.CurrentSeq),
		UpdatedMs:      int(m.UpdatedMs),
	}
	return r
}

//ToPB
func (m *PB_Chat_Flat) ToPB() *PB_Chat {
	r := &PB_Chat{
		ChatId:         int64(m.ChatId),
		ChatKey:        m.ChatKey,
		RoomTypeEnumId: int32(m.RoomTypeEnumId),
		UserId:         int32(m.UserId),
		LastSeqSeen:    int32(m.LastSeqSeen),
		LastSeqDelete:  int32(m.LastSeqDelete),
		PeerUserId:     int32(m.PeerUserId),
		GroupId:        int64(m.GroupId),
		CreatedTime:    int32(m.CreatedTime),
		CurrentSeq:     int32(m.CurrentSeq),
		UpdatedMs:      int64(m.UpdatedMs),
	}
	return r
}

//folding
var PB_Chat__FOlD = &PB_Chat{
	ChatId:         0,
	ChatKey:        "",
	RoomTypeEnumId: 0,
	UserId:         0,
	LastSeqSeen:    0,
	LastSeqDelete:  0,
	PeerUserId:     0,
	GroupId:        0,
	CreatedTime:    0,
	CurrentSeq:     0,
	UpdatedMs:      0,
}

type PB_DirectMessage_Flat struct {
	MessageId          int
	RoomKey            string
	UserId             int
	MessageFileId      int
	MessageTypeEnum    int
	Text               string
	Time               int
	PeerReceivedTime   int
	PeerSeenTime       int
	DeliviryStatusEnum int
}

//ToPB
func (m *PB_DirectMessage) ToFlat() *PB_DirectMessage_Flat {
	r := &PB_DirectMessage_Flat{
		MessageId:          int(m.MessageId),
		RoomKey:            m.RoomKey,
		UserId:             int(m.UserId),
		MessageFileId:      int(m.MessageFileId),
		MessageTypeEnum:    int(m.MessageTypeEnum),
		Text:               m.Text,
		Time:               int(m.Time),
		PeerReceivedTime:   int(m.PeerReceivedTime),
		PeerSeenTime:       int(m.PeerSeenTime),
		DeliviryStatusEnum: int(m.DeliviryStatusEnum),
	}
	return r
}

//ToPB
func (m *PB_DirectMessage_Flat) ToPB() *PB_DirectMessage {
	r := &PB_DirectMessage{
		MessageId:          int64(m.MessageId),
		RoomKey:            m.RoomKey,
		UserId:             int32(m.UserId),
		MessageFileId:      int64(m.MessageFileId),
		MessageTypeEnum:    int32(m.MessageTypeEnum),
		Text:               m.Text,
		Time:               int32(m.Time),
		PeerReceivedTime:   int32(m.PeerReceivedTime),
		PeerSeenTime:       int32(m.PeerSeenTime),
		DeliviryStatusEnum: int32(m.DeliviryStatusEnum),
	}
	return r
}

//folding
var PB_DirectMessage__FOlD = &PB_DirectMessage{
	MessageId:          0,
	RoomKey:            "",
	UserId:             0,
	MessageFileId:      0,
	MessageTypeEnum:    0,
	Text:               "",
	Time:               0,
	PeerReceivedTime:   0,
	PeerSeenTime:       0,
	DeliviryStatusEnum: 0,
}

type PB_MessageFile_Flat struct {
	MessageFileId   int
	Name            string
	Size            int
	FileTypeEnumId  int
	Width           int
	Height          int
	Duration        int
	Extension       string
	HashMd5         string
	HashAccess      int
	CreatedSe       int
	ServerSrc       string
	ServerPath      string
	ServerThumbPath string
	BucketId        string
	ServerId        int
	CanDel          int
}

//ToPB
func (m *PB_MessageFile) ToFlat() *PB_MessageFile_Flat {
	r := &PB_MessageFile_Flat{
		MessageFileId:   int(m.MessageFileId),
		Name:            m.Name,
		Size:            int(m.Size),
		FileTypeEnumId:  int(m.FileTypeEnumId),
		Width:           int(m.Width),
		Height:          int(m.Height),
		Duration:        int(m.Duration),
		Extension:       m.Extension,
		HashMd5:         m.HashMd5,
		HashAccess:      int(m.HashAccess),
		CreatedSe:       int(m.CreatedSe),
		ServerSrc:       m.ServerSrc,
		ServerPath:      m.ServerPath,
		ServerThumbPath: m.ServerThumbPath,
		BucketId:        m.BucketId,
		ServerId:        int(m.ServerId),
		CanDel:          int(m.CanDel),
	}
	return r
}

//ToPB
func (m *PB_MessageFile_Flat) ToPB() *PB_MessageFile {
	r := &PB_MessageFile{
		MessageFileId:   int64(m.MessageFileId),
		Name:            m.Name,
		Size:            int32(m.Size),
		FileTypeEnumId:  int32(m.FileTypeEnumId),
		Width:           int32(m.Width),
		Height:          int32(m.Height),
		Duration:        int32(m.Duration),
		Extension:       m.Extension,
		HashMd5:         m.HashMd5,
		HashAccess:      int64(m.HashAccess),
		CreatedSe:       int32(m.CreatedSe),
		ServerSrc:       m.ServerSrc,
		ServerPath:      m.ServerPath,
		ServerThumbPath: m.ServerThumbPath,
		BucketId:        m.BucketId,
		ServerId:        int32(m.ServerId),
		CanDel:          int32(m.CanDel),
	}
	return r
}

//folding
var PB_MessageFile__FOlD = &PB_MessageFile{
	MessageFileId:   0,
	Name:            "",
	Size:            0,
	FileTypeEnumId:  0,
	Width:           0,
	Height:          0,
	Duration:        0,
	Extension:       "",
	HashMd5:         "",
	HashAccess:      0,
	CreatedSe:       0,
	ServerSrc:       "",
	ServerPath:      "",
	ServerThumbPath: "",
	BucketId:        "",
	ServerId:        0,
	CanDel:          0,
}

type PB_User_Flat struct {
	Id                   int
	UserName             string
	UserNameLower        string
	FirstName            string
	LastName             string
	About                string
	FullName             string
	AvatarUrl            string
	PrivacyProfile       int
	Phone                string
	Email                string
	IsDeleted            int
	PasswordHash         string
	PasswordSalt         string
	FollowersCount       int
	FollowingCount       int
	PostsCount           int
	MediaCount           int
	LikesCount           int
	ResharedCount        int
	LastActionTime       int
	LastPostTime         int
	PrimaryFollowingList int
	CreatedTime          int
	UpdatedTime          int
	SessionUuid          string
	DeviceUuid           string
	LastWifiMacAddress   string
	LastNetworkType      string
	AppVersion           int
	LastActivityTime     int
	LastLoginTime        int
	LastIpAddress        string
}

//ToPB
func (m *PB_User) ToFlat() *PB_User_Flat {
	r := &PB_User_Flat{
		Id:                   int(m.Id),
		UserName:             m.UserName,
		UserNameLower:        m.UserNameLower,
		FirstName:            m.FirstName,
		LastName:             m.LastName,
		About:                m.About,
		FullName:             m.FullName,
		AvatarUrl:            m.AvatarUrl,
		PrivacyProfile:       int(m.PrivacyProfile),
		Phone:                m.Phone,
		Email:                m.Email,
		IsDeleted:            int(m.IsDeleted),
		PasswordHash:         m.PasswordHash,
		PasswordSalt:         m.PasswordSalt,
		FollowersCount:       int(m.FollowersCount),
		FollowingCount:       int(m.FollowingCount),
		PostsCount:           int(m.PostsCount),
		MediaCount:           int(m.MediaCount),
		LikesCount:           int(m.LikesCount),
		ResharedCount:        int(m.ResharedCount),
		LastActionTime:       int(m.LastActionTime),
		LastPostTime:         int(m.LastPostTime),
		PrimaryFollowingList: int(m.PrimaryFollowingList),
		CreatedTime:          int(m.CreatedTime),
		UpdatedTime:          int(m.UpdatedTime),
		SessionUuid:          m.SessionUuid,
		DeviceUuid:           m.DeviceUuid,
		LastWifiMacAddress:   m.LastWifiMacAddress,
		LastNetworkType:      m.LastNetworkType,
		AppVersion:           int(m.AppVersion),
		LastActivityTime:     int(m.LastActivityTime),
		LastLoginTime:        int(m.LastLoginTime),
		LastIpAddress:        m.LastIpAddress,
	}
	return r
}

//ToPB
func (m *PB_User_Flat) ToPB() *PB_User {
	r := &PB_User{
		Id:                   int32(m.Id),
		UserName:             m.UserName,
		UserNameLower:        m.UserNameLower,
		FirstName:            m.FirstName,
		LastName:             m.LastName,
		About:                m.About,
		FullName:             m.FullName,
		AvatarUrl:            m.AvatarUrl,
		PrivacyProfile:       int32(m.PrivacyProfile),
		Phone:                m.Phone,
		Email:                m.Email,
		IsDeleted:            int32(m.IsDeleted),
		PasswordHash:         m.PasswordHash,
		PasswordSalt:         m.PasswordSalt,
		FollowersCount:       int32(m.FollowersCount),
		FollowingCount:       int32(m.FollowingCount),
		PostsCount:           int32(m.PostsCount),
		MediaCount:           int32(m.MediaCount),
		LikesCount:           int32(m.LikesCount),
		ResharedCount:        int32(m.ResharedCount),
		LastActionTime:       int32(m.LastActionTime),
		LastPostTime:         int32(m.LastPostTime),
		PrimaryFollowingList: int32(m.PrimaryFollowingList),
		CreatedTime:          int32(m.CreatedTime),
		UpdatedTime:          int32(m.UpdatedTime),
		SessionUuid:          m.SessionUuid,
		DeviceUuid:           m.DeviceUuid,
		LastWifiMacAddress:   m.LastWifiMacAddress,
		LastNetworkType:      m.LastNetworkType,
		AppVersion:           int32(m.AppVersion),
		LastActivityTime:     int32(m.LastActivityTime),
		LastLoginTime:        int32(m.LastLoginTime),
		LastIpAddress:        m.LastIpAddress,
	}
	return r
}

//folding
var PB_User__FOlD = &PB_User{
	Id:                   0,
	UserName:             "",
	UserNameLower:        "",
	FirstName:            "",
	LastName:             "",
	About:                "",
	FullName:             "",
	AvatarUrl:            "",
	PrivacyProfile:       0,
	Phone:                "",
	Email:                "",
	IsDeleted:            0,
	PasswordHash:         "",
	PasswordSalt:         "",
	FollowersCount:       0,
	FollowingCount:       0,
	PostsCount:           0,
	MediaCount:           0,
	LikesCount:           0,
	ResharedCount:        0,
	LastActionTime:       0,
	LastPostTime:         0,
	PrimaryFollowingList: 0,
	CreatedTime:          0,
	UpdatedTime:          0,
	SessionUuid:          "",
	DeviceUuid:           "",
	LastWifiMacAddress:   "",
	LastNetworkType:      "",
	AppVersion:           0,
	LastActivityTime:     0,
	LastLoginTime:        0,
	LastIpAddress:        "",
}

type PB_UpdateGroupParticipants_Flat struct {
}

//ToPB
func (m *PB_UpdateGroupParticipants) ToFlat() *PB_UpdateGroupParticipants_Flat {
	r := &PB_UpdateGroupParticipants_Flat{}
	return r
}

//ToPB
func (m *PB_UpdateGroupParticipants_Flat) ToPB() *PB_UpdateGroupParticipants {
	r := &PB_UpdateGroupParticipants{}
	return r
}

//folding
var PB_UpdateGroupParticipants__FOlD = &PB_UpdateGroupParticipants{}

type PB_UpdateNotifySettings_Flat struct {
}

//ToPB
func (m *PB_UpdateNotifySettings) ToFlat() *PB_UpdateNotifySettings_Flat {
	r := &PB_UpdateNotifySettings_Flat{}
	return r
}

//ToPB
func (m *PB_UpdateNotifySettings_Flat) ToPB() *PB_UpdateNotifySettings {
	r := &PB_UpdateNotifySettings{}
	return r
}

//folding
var PB_UpdateNotifySettings__FOlD = &PB_UpdateNotifySettings{}

type PB_UpdateServiceNotification_Flat struct {
}

//ToPB
func (m *PB_UpdateServiceNotification) ToFlat() *PB_UpdateServiceNotification_Flat {
	r := &PB_UpdateServiceNotification_Flat{}
	return r
}

//ToPB
func (m *PB_UpdateServiceNotification_Flat) ToPB() *PB_UpdateServiceNotification {
	r := &PB_UpdateServiceNotification{}
	return r
}

//folding
var PB_UpdateServiceNotification__FOlD = &PB_UpdateServiceNotification{}

type PB_UpdateMessageMeta_Flat struct {
	MessageId int
	AtTime    int
}

//ToPB
func (m *PB_UpdateMessageMeta) ToFlat() *PB_UpdateMessageMeta_Flat {
	r := &PB_UpdateMessageMeta_Flat{
		MessageId: int(m.MessageId),
		AtTime:    int(m.AtTime),
	}
	return r
}

//ToPB
func (m *PB_UpdateMessageMeta_Flat) ToPB() *PB_UpdateMessageMeta {
	r := &PB_UpdateMessageMeta{
		MessageId: int64(m.MessageId),
		AtTime:    int64(m.AtTime),
	}
	return r
}

//folding
var PB_UpdateMessageMeta__FOlD = &PB_UpdateMessageMeta{
	MessageId: 0,
	AtTime:    0,
}

type PB_UpdateMessageId_Flat struct {
	OldMessageId int
	NewMessageId int
}

//ToPB
func (m *PB_UpdateMessageId) ToFlat() *PB_UpdateMessageId_Flat {
	r := &PB_UpdateMessageId_Flat{
		OldMessageId: int(m.OldMessageId),
		NewMessageId: int(m.NewMessageId),
	}
	return r
}

//ToPB
func (m *PB_UpdateMessageId_Flat) ToPB() *PB_UpdateMessageId {
	r := &PB_UpdateMessageId{
		OldMessageId: int64(m.OldMessageId),
		NewMessageId: int64(m.NewMessageId),
	}
	return r
}

//folding
var PB_UpdateMessageId__FOlD = &PB_UpdateMessageId{
	OldMessageId: 0,
	NewMessageId: 0,
}

type PB_UpdateMessageToEdit_Flat struct {
	MessageId int
	NewText   string
}

//ToPB
func (m *PB_UpdateMessageToEdit) ToFlat() *PB_UpdateMessageToEdit_Flat {
	r := &PB_UpdateMessageToEdit_Flat{
		MessageId: int(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

//ToPB
func (m *PB_UpdateMessageToEdit_Flat) ToPB() *PB_UpdateMessageToEdit {
	r := &PB_UpdateMessageToEdit{
		MessageId: int64(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

//folding
var PB_UpdateMessageToEdit__FOlD = &PB_UpdateMessageToEdit{
	MessageId: 0,
	NewText:   "",
}

type PB_UpdateMessageToDelete_Flat struct {
	MessageId int
}

//ToPB
func (m *PB_UpdateMessageToDelete) ToFlat() *PB_UpdateMessageToDelete_Flat {
	r := &PB_UpdateMessageToDelete_Flat{
		MessageId: int(m.MessageId),
	}
	return r
}

//ToPB
func (m *PB_UpdateMessageToDelete_Flat) ToPB() *PB_UpdateMessageToDelete {
	r := &PB_UpdateMessageToDelete{
		MessageId: int64(m.MessageId),
	}
	return r
}

//folding
var PB_UpdateMessageToDelete__FOlD = &PB_UpdateMessageToDelete{
	MessageId: 0,
}

type PB_UpdateRoomActionDoing_Flat struct {
	RoomKey    string
	ActionType RoomActionDoingEnum
}

//ToPB
func (m *PB_UpdateRoomActionDoing) ToFlat() *PB_UpdateRoomActionDoing_Flat {
	r := &PB_UpdateRoomActionDoing_Flat{
		RoomKey: m.RoomKey,
	}
	return r
}

//ToPB
func (m *PB_UpdateRoomActionDoing_Flat) ToPB() *PB_UpdateRoomActionDoing {
	r := &PB_UpdateRoomActionDoing{
		RoomKey: m.RoomKey,
	}
	return r
}

//folding
var PB_UpdateRoomActionDoing__FOlD = &PB_UpdateRoomActionDoing{
	RoomKey: "",
}

type PB_UpdateUserBlocked_Flat struct {
	UserId  int
	Blocked bool
}

//ToPB
func (m *PB_UpdateUserBlocked) ToFlat() *PB_UpdateUserBlocked_Flat {
	r := &PB_UpdateUserBlocked_Flat{
		UserId:  int(m.UserId),
		Blocked: m.Blocked,
	}
	return r
}

//ToPB
func (m *PB_UpdateUserBlocked_Flat) ToPB() *PB_UpdateUserBlocked {
	r := &PB_UpdateUserBlocked{
		UserId:  int32(m.UserId),
		Blocked: m.Blocked,
	}
	return r
}

//folding
var PB_UpdateUserBlocked__FOlD = &PB_UpdateUserBlocked{
	UserId:  0,
	Blocked: false,
}

type PB_PushHolderView_Flat struct {
	NewMessages               []PB_MessageView
	ChatFiles                 []PB_MessageFileView
	Chats                     []PB_ChatView
	Users                     []PB_UserView
	MessagesChangeIds         []PB_UpdateMessageId
	MessagesToUpdate          []PB_UpdateMessageToEdit
	MessagesToDelete          []PB_UpdateMessageToDelete
	MessagesDelivierdToServer []PB_UpdateMessageMeta
	MessagesDelivierdToPeer   []PB_UpdateMessageMeta
	MessagesSeenByPeer        []PB_UpdateMessageMeta
	MessagesDeletedFromServer []PB_UpdateMessageMeta
	RoomActionDoing           []PB_UpdateRoomActionDoing
	UserBlockedByMe           []PB_UpdateUserBlocked
	UserBlockedMe             []PB_UpdateUserBlocked
}

//ToPB
func (m *PB_PushHolderView) ToFlat() *PB_PushHolderView_Flat {
	r := &PB_PushHolderView_Flat{}
	return r
}

//ToPB
func (m *PB_PushHolderView_Flat) ToPB() *PB_PushHolderView {
	r := &PB_PushHolderView{}
	return r
}

//folding
var PB_PushHolderView__FOlD = &PB_PushHolderView{}

type PB_DirectUpdatesView_Flat struct {
	NewMessages               []PB_MessageView
	ChatFiles                 []PB_MessageFileView
	Chats                     []PB_ChatView
	Users                     []PB_UserView
	MessagesChangeIds         []PB_UpdateMessageId
	MessagesToUpdate          []PB_UpdateMessageToEdit
	MessagesToDelete          []PB_UpdateMessageToDelete
	MessagesDelivierdToServer []PB_UpdateMessageMeta
	MessagesDelivierdToPeer   []PB_UpdateMessageMeta
	MessagesSeenByPeer        []PB_UpdateMessageMeta
	MessagesDeletedFromServer []PB_UpdateMessageMeta
	RoomActionDoing           []PB_UpdateRoomActionDoing
}

//ToPB
func (m *PB_DirectUpdatesView) ToFlat() *PB_DirectUpdatesView_Flat {
	r := &PB_DirectUpdatesView_Flat{}
	return r
}

//ToPB
func (m *PB_DirectUpdatesView_Flat) ToPB() *PB_DirectUpdatesView {
	r := &PB_DirectUpdatesView{}
	return r
}

//folding
var PB_DirectUpdatesView__FOlD = &PB_DirectUpdatesView{}

type PB_GeneralUpdatesView_Flat struct {
	UserBlockedByMe []PB_UpdateUserBlocked
	UserBlockedMe   []PB_UpdateUserBlocked
}

//ToPB
func (m *PB_GeneralUpdatesView) ToFlat() *PB_GeneralUpdatesView_Flat {
	r := &PB_GeneralUpdatesView_Flat{}
	return r
}

//ToPB
func (m *PB_GeneralUpdatesView_Flat) ToPB() *PB_GeneralUpdatesView {
	r := &PB_GeneralUpdatesView{}
	return r
}

//folding
var PB_GeneralUpdatesView__FOlD = &PB_GeneralUpdatesView{}

type PB_NotifyUpdatesView_Flat struct {
	UserBlockedByMe []PB_UpdateUserBlocked
	UserBlockedMe   []PB_UpdateUserBlocked
}

//ToPB
func (m *PB_NotifyUpdatesView) ToFlat() *PB_NotifyUpdatesView_Flat {
	r := &PB_NotifyUpdatesView_Flat{}
	return r
}

//ToPB
func (m *PB_NotifyUpdatesView_Flat) ToPB() *PB_NotifyUpdatesView {
	r := &PB_NotifyUpdatesView{}
	return r
}

//folding
var PB_NotifyUpdatesView__FOlD = &PB_NotifyUpdatesView{}

type PB_ChatView_Flat struct {
	ChatKey              string
	ChatId               int
	RoomTypeEnumId       int
	UserId               int
	PeerUserId           int
	GroupId              int
	CreatedTime          int
	UpdatedMs            int
	LastMessageId        int
	LastDeletedMessageId int
	LastSeenMessageId    int
	LastSeqSeen          int
	LastSeqDelete        int
	CurrentSeq           int
	UserView             PB_UserView
	SharedMediaCount     int
	UnseenCount          int
	FirstUnreadMessage   PB_MessageView
	LastMessage          PB_MessageView
}

//ToPB
func (m *PB_ChatView) ToFlat() *PB_ChatView_Flat {
	r := &PB_ChatView_Flat{
		ChatKey:              m.ChatKey,
		ChatId:               int(m.ChatId),
		RoomTypeEnumId:       int(m.RoomTypeEnumId),
		UserId:               int(m.UserId),
		PeerUserId:           int(m.PeerUserId),
		GroupId:              int(m.GroupId),
		CreatedTime:          int(m.CreatedTime),
		UpdatedMs:            int(m.UpdatedMs),
		LastMessageId:        int(m.LastMessageId),
		LastDeletedMessageId: int(m.LastDeletedMessageId),
		LastSeenMessageId:    int(m.LastSeenMessageId),
		LastSeqSeen:          int(m.LastSeqSeen),
		LastSeqDelete:        int(m.LastSeqDelete),
		CurrentSeq:           int(m.CurrentSeq),

		SharedMediaCount: int(m.SharedMediaCount),
		UnseenCount:      int(m.UnseenCount),
	}
	return r
}

//ToPB
func (m *PB_ChatView_Flat) ToPB() *PB_ChatView {
	r := &PB_ChatView{
		ChatKey:              m.ChatKey,
		ChatId:               int64(m.ChatId),
		RoomTypeEnumId:       int32(m.RoomTypeEnumId),
		UserId:               int32(m.UserId),
		PeerUserId:           int32(m.PeerUserId),
		GroupId:              int64(m.GroupId),
		CreatedTime:          int32(m.CreatedTime),
		UpdatedMs:            int64(m.UpdatedMs),
		LastMessageId:        int64(m.LastMessageId),
		LastDeletedMessageId: int64(m.LastDeletedMessageId),
		LastSeenMessageId:    int64(m.LastSeenMessageId),
		LastSeqSeen:          int32(m.LastSeqSeen),
		LastSeqDelete:        int32(m.LastSeqDelete),
		CurrentSeq:           int32(m.CurrentSeq),

		SharedMediaCount: int32(m.SharedMediaCount),
		UnseenCount:      int32(m.UnseenCount),
	}
	return r
}

//folding
var PB_ChatView__FOlD = &PB_ChatView{
	ChatKey:              "",
	ChatId:               0,
	RoomTypeEnumId:       0,
	UserId:               0,
	PeerUserId:           0,
	GroupId:              0,
	CreatedTime:          0,
	UpdatedMs:            0,
	LastMessageId:        0,
	LastDeletedMessageId: 0,
	LastSeenMessageId:    0,
	LastSeqSeen:          0,
	LastSeqDelete:        0,
	CurrentSeq:           0,

	SharedMediaCount: 0,
	UnseenCount:      0,
}

type PB_MessageView_Flat struct {
	MessageId            int
	RoomKey              string
	UserId               int
	MessageFileId        int
	MessageTypeEnumId    int
	Text                 string
	Time                 int
	PeerReceivedTime     int
	PeerSeenTime         int
	DeliviryStatusEnumId int
	ChatId               int
	RoomTypeEnumId       int
	IsByMe               bool
	RemoteId             int
	MessageFileView      PB_MessageFileView
}

//ToPB
func (m *PB_MessageView) ToFlat() *PB_MessageView_Flat {
	r := &PB_MessageView_Flat{
		MessageId:            int(m.MessageId),
		RoomKey:              m.RoomKey,
		UserId:               int(m.UserId),
		MessageFileId:        int(m.MessageFileId),
		MessageTypeEnumId:    int(m.MessageTypeEnumId),
		Text:                 m.Text,
		Time:                 int(m.Time),
		PeerReceivedTime:     int(m.PeerReceivedTime),
		PeerSeenTime:         int(m.PeerSeenTime),
		DeliviryStatusEnumId: int(m.DeliviryStatusEnumId),
		ChatId:               int(m.ChatId),
		RoomTypeEnumId:       int(m.RoomTypeEnumId),
		IsByMe:               m.IsByMe,
		RemoteId:             int(m.RemoteId),
	}
	return r
}

//ToPB
func (m *PB_MessageView_Flat) ToPB() *PB_MessageView {
	r := &PB_MessageView{
		MessageId:            uint64(m.MessageId),
		RoomKey:              m.RoomKey,
		UserId:               int32(m.UserId),
		MessageFileId:        int64(m.MessageFileId),
		MessageTypeEnumId:    int32(m.MessageTypeEnumId),
		Text:                 m.Text,
		Time:                 int32(m.Time),
		PeerReceivedTime:     int32(m.PeerReceivedTime),
		PeerSeenTime:         int32(m.PeerSeenTime),
		DeliviryStatusEnumId: int32(m.DeliviryStatusEnumId),
		ChatId:               int64(m.ChatId),
		RoomTypeEnumId:       int32(m.RoomTypeEnumId),
		IsByMe:               m.IsByMe,
		RemoteId:             int64(m.RemoteId),
	}
	return r
}

//folding
var PB_MessageView__FOlD = &PB_MessageView{
	MessageId:            0,
	RoomKey:              "",
	UserId:               0,
	MessageFileId:        0,
	MessageTypeEnumId:    0,
	Text:                 "",
	Time:                 0,
	PeerReceivedTime:     0,
	PeerSeenTime:         0,
	DeliviryStatusEnumId: 0,
	ChatId:               0,
	RoomTypeEnumId:       0,
	IsByMe:               false,
	RemoteId:             0,
}

type PB_MessageFileView_Flat struct {
	MessageFileId       int
	OriginalUserId      int
	Name                string
	Size                int
	FileTypeEnumId      int
	Width               int
	Height              int
	Duration            int
	Extension           string
	HashMd5             string
	HashAccess          int
	CreatedSe           int
	ServerSrc           string
	ServerPath          string
	ServerThumbPath     string
	BucketId            string
	ServerId            int
	CanDel              int
	ServerThumbLocalSrc string
	RemoteMessageFileId int
	LocalSrc            string
	ThumbLocalSrc       string
	MessageFileStatusId string
}

//ToPB
func (m *PB_MessageFileView) ToFlat() *PB_MessageFileView_Flat {
	r := &PB_MessageFileView_Flat{
		MessageFileId:       int(m.MessageFileId),
		OriginalUserId:      int(m.OriginalUserId),
		Name:                m.Name,
		Size:                int(m.Size),
		FileTypeEnumId:      int(m.FileTypeEnumId),
		Width:               int(m.Width),
		Height:              int(m.Height),
		Duration:            int(m.Duration),
		Extension:           m.Extension,
		HashMd5:             m.HashMd5,
		HashAccess:          int(m.HashAccess),
		CreatedSe:           int(m.CreatedSe),
		ServerSrc:           m.ServerSrc,
		ServerPath:          m.ServerPath,
		ServerThumbPath:     m.ServerThumbPath,
		BucketId:            m.BucketId,
		ServerId:            int(m.ServerId),
		CanDel:              int(m.CanDel),
		ServerThumbLocalSrc: m.ServerThumbLocalSrc,
		RemoteMessageFileId: int(m.RemoteMessageFileId),
		LocalSrc:            m.LocalSrc,
		ThumbLocalSrc:       m.ThumbLocalSrc,
		MessageFileStatusId: m.MessageFileStatusId,
	}
	return r
}

//ToPB
func (m *PB_MessageFileView_Flat) ToPB() *PB_MessageFileView {
	r := &PB_MessageFileView{
		MessageFileId:       int64(m.MessageFileId),
		OriginalUserId:      int32(m.OriginalUserId),
		Name:                m.Name,
		Size:                int32(m.Size),
		FileTypeEnumId:      int32(m.FileTypeEnumId),
		Width:               int32(m.Width),
		Height:              int32(m.Height),
		Duration:            int32(m.Duration),
		Extension:           m.Extension,
		HashMd5:             m.HashMd5,
		HashAccess:          int64(m.HashAccess),
		CreatedSe:           int32(m.CreatedSe),
		ServerSrc:           m.ServerSrc,
		ServerPath:          m.ServerPath,
		ServerThumbPath:     m.ServerThumbPath,
		BucketId:            m.BucketId,
		ServerId:            int32(m.ServerId),
		CanDel:              int32(m.CanDel),
		ServerThumbLocalSrc: m.ServerThumbLocalSrc,
		RemoteMessageFileId: int64(m.RemoteMessageFileId),
		LocalSrc:            m.LocalSrc,
		ThumbLocalSrc:       m.ThumbLocalSrc,
		MessageFileStatusId: m.MessageFileStatusId,
	}
	return r
}

//folding
var PB_MessageFileView__FOlD = &PB_MessageFileView{
	MessageFileId:       0,
	OriginalUserId:      0,
	Name:                "",
	Size:                0,
	FileTypeEnumId:      0,
	Width:               0,
	Height:              0,
	Duration:            0,
	Extension:           "",
	HashMd5:             "",
	HashAccess:          0,
	CreatedSe:           0,
	ServerSrc:           "",
	ServerPath:          "",
	ServerThumbPath:     "",
	BucketId:            "",
	ServerId:            0,
	CanDel:              0,
	ServerThumbLocalSrc: "",
	RemoteMessageFileId: 0,
	LocalSrc:            "",
	ThumbLocalSrc:       "",
	MessageFileStatusId: "",
}

type PB_UserView_Flat struct {
	UserId           int
	UserName         string
	FirstName        string
	LastName         string
	About            string
	FullName         string
	AvatarUrl        string
	PrivacyProfile   int
	IsDeleted        int
	FollowersCount   int
	FollowingCount   int
	PostsCount       int
	UpdatedTime      int
	AppVersion       int
	LastActivityTime int
	FollowingType    int
}

//ToPB
func (m *PB_UserView) ToFlat() *PB_UserView_Flat {
	r := &PB_UserView_Flat{
		UserId:           int(m.UserId),
		UserName:         m.UserName,
		FirstName:        m.FirstName,
		LastName:         m.LastName,
		About:            m.About,
		FullName:         m.FullName,
		AvatarUrl:        m.AvatarUrl,
		PrivacyProfile:   int(m.PrivacyProfile),
		IsDeleted:        int(m.IsDeleted),
		FollowersCount:   int(m.FollowersCount),
		FollowingCount:   int(m.FollowingCount),
		PostsCount:       int(m.PostsCount),
		UpdatedTime:      int(m.UpdatedTime),
		AppVersion:       int(m.AppVersion),
		LastActivityTime: int(m.LastActivityTime),
		FollowingType:    int(m.FollowingType),
	}
	return r
}

//ToPB
func (m *PB_UserView_Flat) ToPB() *PB_UserView {
	r := &PB_UserView{
		UserId:           int32(m.UserId),
		UserName:         m.UserName,
		FirstName:        m.FirstName,
		LastName:         m.LastName,
		About:            m.About,
		FullName:         m.FullName,
		AvatarUrl:        m.AvatarUrl,
		PrivacyProfile:   int32(m.PrivacyProfile),
		IsDeleted:        int32(m.IsDeleted),
		FollowersCount:   int32(m.FollowersCount),
		FollowingCount:   int32(m.FollowingCount),
		PostsCount:       int32(m.PostsCount),
		UpdatedTime:      int32(m.UpdatedTime),
		AppVersion:       int32(m.AppVersion),
		LastActivityTime: int32(m.LastActivityTime),
		FollowingType:    int32(m.FollowingType),
	}
	return r
}

//folding
var PB_UserView__FOlD = &PB_UserView{
	UserId:           0,
	UserName:         "",
	FirstName:        "",
	LastName:         "",
	About:            "",
	FullName:         "",
	AvatarUrl:        "",
	PrivacyProfile:   0,
	IsDeleted:        0,
	FollowersCount:   0,
	FollowingCount:   0,
	PostsCount:       0,
	UpdatedTime:      0,
	AppVersion:       0,
	LastActivityTime: 0,
	FollowingType:    0,
}

/*
///// to_flat ///

func(m *PB_RoomInlineView)ToFlat() *PB_RoomInlineView_Flat {
r := &PB_RoomInlineView_Flat{
    RoomId:  int(m.RoomId) ,

}
return r
}

func(m *PB_MessageForwardedFrom)ToFlat() *PB_MessageForwardedFrom_Flat {
r := &PB_MessageForwardedFrom_Flat{
    RoomId:  int(m.RoomId) ,

    MessageId:  int(m.MessageId) ,
    MessageSeq:  int(m.MessageSeq) ,
}
return r
}

func(m *PB_GroupView)ToFlat() *PB_GroupView_Flat {
r := &PB_GroupView_Flat{
    GroupId:  int(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int(m.MembersCount) ,
    GroupPrivacyEnum:  int(m.GroupPrivacyEnum) ,
    CreatorUserId:  int(m.CreatorUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    CurrentSeq:  int(m.CurrentSeq) ,
}
return r
}

func(m *PB_GroupMemberView)ToFlat() *PB_GroupMemberView_Flat {
r := &PB_GroupMemberView_Flat{
    Id:  int(m.Id) ,
    GroupId:  int(m.GroupId) ,
    UserId:  int(m.UserId) ,
    ByUserId:  int(m.ByUserId) ,
    GroupRoleEnum:  int(m.GroupRoleEnum) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_MessageFileView__DEp)ToFlat() *PB_MessageFileView__DEp_Flat {
r := &PB_MessageFileView__DEp_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnum:  int(m.FileTypeEnum) ,
    MimeType:  m.MimeType ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData64:  m.ThumbData64 ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_ReqLastChangesForTheRoom)ToFlat() *PB_ReqLastChangesForTheRoom_Flat {
r := &PB_ReqLastChangesForTheRoom_Flat{
    RoomId:  int(m.RoomId) ,
    LastLogId:  int(m.LastLogId) ,
    LastHaveSeq:  int(m.LastHaveSeq) ,
}
return r
}

func(m *PB_ResponseLastChangesForTheRoom)ToFlat() *PB_ResponseLastChangesForTheRoom_Flat {
r := &PB_ResponseLastChangesForTheRoom_Flat{

}
return r
}

func(m *PB_RequestSetLastSeenMessages)ToFlat() *PB_RequestSetLastSeenMessages_Flat {
r := &PB_RequestSetLastSeenMessages_Flat{
    RoomId:  int(m.RoomId) ,
    FromMessageId:  int(m.FromMessageId) ,
    ToMessageId:  int(m.ToMessageId) ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}

func(m *PB_ResponseSetLastSeenMessages)ToFlat() *PB_ResponseSetLastSeenMessages_Flat {
r := &PB_ResponseSetLastSeenMessages_Flat{

}
return r
}

func(m *PB_MessagesCollections)ToFlat() *PB_MessagesCollections_Flat {
r := &PB_MessagesCollections_Flat{
    DirectMessagesIds:  helper.SliceInt64ToInt(m.DirectMessagesIds) ,
    GroupMessagesIds:  helper.SliceInt64ToInt(m.GroupMessagesIds) ,
    BroadCatMessagesIds:  helper.SliceInt64ToInt(m.BroadCatMessagesIds) ,
}
return r
}

func(m *PB_DirectLogView)ToFlat() *PB_DirectLogView_Flat {
r := &PB_DirectLogView_Flat{


}
return r
}

func(m *PB_DirectLog)ToFlat() *PB_DirectLog_Flat {
r := &PB_DirectLog_Flat{
    Id:  int(m.Id) ,
    ToUserId:  int(m.ToUserId) ,
    MessageId:  int(m.MessageId) ,
    ChatId:  int(m.ChatId) ,
    PeerUserId:  int(m.PeerUserId) ,
    EventType:  int(m.EventType) ,
    RoomLogTypeId:  int(m.RoomLogTypeId) ,
    FromSeq:  int(m.FromSeq) ,
    ToSeq:  int(m.ToSeq) ,
    ExtraPB:  []byte(m.ExtraPB) ,
    ExtraJson:  m.ExtraJson ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}

func(m *PB_PushDirectLogViewsMany)ToFlat() *PB_PushDirectLogViewsMany_Flat {
r := &PB_PushDirectLogViewsMany_Flat{

}
return r
}

func(m *PB_UserWithMe)ToFlat() *PB_UserWithMe_Flat {
r := &PB_UserWithMe_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    UpdatedTime:  int(m.UpdatedTime) ,
    AppVersion:  int(m.AppVersion) ,
    FollowingType:  int(m.FollowingType) ,
}
return r
}

func(m *PB_Message)ToFlat() *PB_Message_Flat {
r := &PB_Message_Flat{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    PeerId:  int(m.PeerId) ,
    RoomTypeId:  int(m.RoomTypeId) ,
    MessageTypeId:  int(m.MessageTypeId) ,
    Text:  m.Text ,
    ExtraJson:  m.ExtraJson ,
    IsByMe:  int(m.IsByMe) ,
    IsStared:  int(m.IsStared) ,
    CreatedMs:  int(m.CreatedMs) ,
    CreatedDeviceMs:  int(m.CreatedDeviceMs) ,
    SortId:  int(m.SortId) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    ServerReceivedTime:  int(m.ServerReceivedTime) ,
    ServerDeletedTime:  int(m.ServerDeletedTime) ,
    ISeenTime:  int(m.ISeenTime) ,
    ToPush:  int(m.ToPush) ,
    MsgFile_LocalSrc:  m.MsgFile_LocalSrc ,
    MsgFile_Status:  int(m.MsgFile_Status) ,

}
return r
}

func(m *PB_MsgFile)ToFlat() *PB_MsgFile_Flat {
r := &PB_MsgFile_Flat{
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileType:  int(m.FileType) ,
    MimeType:  m.MimeType ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData:  []byte(m.ThumbData) ,
    Data:  []byte(m.Data) ,
    ServerSrc:  m.ServerSrc ,
}
return r
}

func(m *PB_Response)ToFlat() *PB_Response_Flat {
r := &PB_Response_Flat{
    Status:  int(m.Status) ,
}
return r
}

func(m *PB_Request)ToFlat() *PB_Request_Flat {
r := &PB_Request_Flat{
    Status:  int(m.Status) ,
}
return r
}

func(m *PB_RequestMsgAddMany)ToFlat() *PB_RequestMsgAddMany_Flat {
r := &PB_RequestMsgAddMany_Flat{


}
return r
}

func(m *PB_ResponseMsgAddMany)ToFlat() *PB_ResponseMsgAddMany_Flat {
r := &PB_ResponseMsgAddMany_Flat{

}
return r
}

func(m *PB_Push)ToFlat() *PB_Push_Flat {
r := &PB_Push_Flat{
    Status:  int(m.Status) ,
}
return r
}

func(m *PB_Result)ToFlat() *PB_Result_Flat {
r := &PB_Result_Flat{
    Status:  int(m.Status) ,
}
return r
}

func(m *PB_PushMsgAddMany)ToFlat() *PB_PushMsgAddMany_Flat {
r := &PB_PushMsgAddMany_Flat{



}
return r
}

func(m *PB_ResultMsgAddMany)ToFlat() *PB_ResultMsgAddMany_Flat {
r := &PB_ResultMsgAddMany_Flat{

}
return r
}

func(m *PB_MsgEvent)ToFlat() *PB_MsgEvent_Flat {
r := &PB_MsgEvent_Flat{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    PeerUserId:  int(m.PeerUserId) ,
    EventType:  int(m.EventType) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_PushMsgEvents)ToFlat() *PB_PushMsgEvents_Flat {
r := &PB_PushMsgEvents_Flat{


}
return r
}

func(m *PB_ResultMsgEvents)ToFlat() *PB_ResultMsgEvents_Flat {
r := &PB_ResultMsgEvents_Flat{

}
return r
}

func(m *PB_MsgsSeenFromClient)ToFlat() *PB_MsgsSeenFromClient_Flat {
r := &PB_MsgsSeenFromClient_Flat{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_MsgSeen)ToFlat() *PB_MsgSeen_Flat {
r := &PB_MsgSeen_Flat{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_RequestMsgsSeen)ToFlat() *PB_RequestMsgsSeen_Flat {
r := &PB_RequestMsgsSeen_Flat{


}
return r
}

func(m *PB_ResponseMsgsSeen)ToFlat() *PB_ResponseMsgsSeen_Flat {
r := &PB_ResponseMsgsSeen_Flat{

}
return r
}

func(m *PB_ReqRpcAddMsg)ToFlat() *PB_ReqRpcAddMsg_Flat {
r := &PB_ReqRpcAddMsg_Flat{


}
return r
}

func(m *PB_ResRpcAddMsg)ToFlat() *PB_ResRpcAddMsg_Flat {
r := &PB_ResRpcAddMsg_Flat{

    ServerSrc:  m.ServerSrc ,
}
return r
}

func(m *GeoLocation)ToFlat() *GeoLocation_Flat {
r := &GeoLocation_Flat{
    Lat:  float64(m.Lat) ,
    Lon:  float64(m.Lon) ,
}
return r
}

func(m *RoomMessageLog)ToFlat() *RoomMessageLog_Flat {
r := &RoomMessageLog_Flat{

    TargetUserId:  int(m.TargetUserId) ,
    ByUserId:  int(m.ByUserId) ,
}
return r
}

func(m *RoomMessageForwardFrom)ToFlat() *RoomMessageForwardFrom_Flat {
r := &RoomMessageForwardFrom_Flat{
    RoomId:  int(m.RoomId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
}
return r
}

func(m *RoomDraft)ToFlat() *RoomDraft_Flat {
r := &RoomDraft_Flat{
    Message:  m.Message ,
    ReplyTo:  int(m.ReplyTo) ,
}
return r
}

func(m *ChatRoom)ToFlat() *ChatRoom_Flat {
r := &ChatRoom_Flat{
}
return r
}

func(m *Pagination)ToFlat() *Pagination_Flat {
r := &Pagination_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}

func(m *PB_CommandToServer)ToFlat() *PB_CommandToServer_Flat {
r := &PB_CommandToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_CommandToClient)ToFlat() *PB_CommandToClient_Flat {
r := &PB_CommandToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_CommandReachedToServer)ToFlat() *PB_CommandReachedToServer_Flat {
r := &PB_CommandReachedToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
}
return r
}

func(m *PB_CommandReachedToClient)ToFlat() *PB_CommandReachedToClient_Flat {
r := &PB_CommandReachedToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
}
return r
}

func(m *PB_ResponseToClient)ToFlat() *PB_ResponseToClient_Flat {
r := &PB_ResponseToClient_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_UserParam_CheckUserName2)ToFlat() *PB_UserParam_CheckUserName2_Flat {
r := &PB_UserParam_CheckUserName2_Flat{
}
return r
}

func(m *PB_UserResponse_CheckUserName2)ToFlat() *PB_UserResponse_CheckUserName2_Flat {
r := &PB_UserResponse_CheckUserName2_Flat{
}
return r
}

func(m *PB_MsgParam_AddNewTextMessage)ToFlat() *PB_MsgParam_AddNewTextMessage_Flat {
r := &PB_MsgParam_AddNewTextMessage_Flat{
    Text:  m.Text ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,

}
return r
}

func(m *PB_MsgResponse_AddNewTextMessage)ToFlat() *PB_MsgResponse_AddNewTextMessage_Flat {
r := &PB_MsgResponse_AddNewTextMessage_Flat{
}
return r
}

func(m *PB_MsgParam_AddNewMessage)ToFlat() *PB_MsgParam_AddNewMessage_Flat {
r := &PB_MsgParam_AddNewMessage_Flat{
    Text:  m.Text ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,


    Blob:  []byte(m.Blob) ,

}
return r
}

func(m *PB_MsgResponse_AddNewMessage)ToFlat() *PB_MsgResponse_AddNewMessage_Flat {
r := &PB_MsgResponse_AddNewMessage_Flat{
}
return r
}

func(m *PB_MsgParam_SetRoomActionDoing)ToFlat() *PB_MsgParam_SetRoomActionDoing_Flat {
r := &PB_MsgParam_SetRoomActionDoing_Flat{
    GroupId:  int(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,

}
return r
}

func(m *PB_MsgResponse_SetRoomActionDoing)ToFlat() *PB_MsgResponse_SetRoomActionDoing_Flat {
r := &PB_MsgResponse_SetRoomActionDoing_Flat{
}
return r
}

func(m *PB_MsgParam_GetMessagesByIds)ToFlat() *PB_MsgParam_GetMessagesByIds_Flat {
r := &PB_MsgParam_GetMessagesByIds_Flat{

}
return r
}

func(m *PB_MsgResponse_GetMessagesByIds)ToFlat() *PB_MsgResponse_GetMessagesByIds_Flat {
r := &PB_MsgResponse_GetMessagesByIds_Flat{

}
return r
}

func(m *PB_MsgParam_GetMessagesHistory)ToFlat() *PB_MsgParam_GetMessagesHistory_Flat {
r := &PB_MsgParam_GetMessagesHistory_Flat{
    ChatKey:  m.ChatKey ,
    FromSeq:  int(m.FromSeq) ,
    EndSeq:  int(m.EndSeq) ,
}
return r
}

func(m *PB_MsgResponse_GetMessagesHistory)ToFlat() *PB_MsgResponse_GetMessagesHistory_Flat {
r := &PB_MsgResponse_GetMessagesHistory_Flat{

}
return r
}

func(m *PB_MsgParam_SetChatMessagesRangeAsSeen)ToFlat() *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat {
r := &PB_MsgParam_SetChatMessagesRangeAsSeen_Flat{
    ChatKey:  m.ChatKey ,
    FromSeq:  int(m.FromSeq) ,
    EndSeq:  int(m.EndSeq) ,
    SeenTimeMs:  int(m.SeenTimeMs) ,
}
return r
}

func(m *PB_MsgResponse_SetChatMessagesRangeAsSeen)ToFlat() *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat {
r := &PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat{
}
return r
}

func(m *PB_MsgParam_DeleteChatHistory)ToFlat() *PB_MsgParam_DeleteChatHistory_Flat {
r := &PB_MsgParam_DeleteChatHistory_Flat{
    ChatKey:  m.ChatKey ,
    ToSeq:  int(m.ToSeq) ,
}
return r
}

func(m *PB_MsgResponse_DeleteChatHistory)ToFlat() *PB_MsgResponse_DeleteChatHistory_Flat {
r := &PB_MsgResponse_DeleteChatHistory_Flat{
}
return r
}

func(m *PB_MsgParam_DeleteMessagesByIds)ToFlat() *PB_MsgParam_DeleteMessagesByIds_Flat {
r := &PB_MsgParam_DeleteMessagesByIds_Flat{

}
return r
}

func(m *PB_MsgResponse_DeleteMessagesByIds)ToFlat() *PB_MsgResponse_DeleteMessagesByIds_Flat {
r := &PB_MsgResponse_DeleteMessagesByIds_Flat{
}
return r
}

func(m *PB_MsgParam_SetMessagesAsReceived)ToFlat() *PB_MsgParam_SetMessagesAsReceived_Flat {
r := &PB_MsgParam_SetMessagesAsReceived_Flat{

}
return r
}

func(m *PB_MsgResponse_SetMessagesAsReceived)ToFlat() *PB_MsgResponse_SetMessagesAsReceived_Flat {
r := &PB_MsgResponse_SetMessagesAsReceived_Flat{
}
return r
}

func(m *PB_MsgParam_ForwardMessages)ToFlat() *PB_MsgParam_ForwardMessages_Flat {
r := &PB_MsgParam_ForwardMessages_Flat{

    ToDirectChatKeys:  m.ToDirectChatKeys ,
    ToGroupChatIds:  helper.SliceInt64ToInt(m.ToGroupChatIds) ,
}
return r
}

func(m *PB_MsgResponse_ForwardMessages)ToFlat() *PB_MsgResponse_ForwardMessages_Flat {
r := &PB_MsgResponse_ForwardMessages_Flat{
}
return r
}

func(m *PB_MsgParam_EditMessage)ToFlat() *PB_MsgParam_EditMessage_Flat {
r := &PB_MsgParam_EditMessage_Flat{
    ChatKey:  m.ChatKey ,

    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_MsgResponse_EditMessage)ToFlat() *PB_MsgResponse_EditMessage_Flat {
r := &PB_MsgResponse_EditMessage_Flat{
}
return r
}

func(m *PB_MsgParam_BroadcastNewMessage)ToFlat() *PB_MsgParam_BroadcastNewMessage_Flat {
r := &PB_MsgParam_BroadcastNewMessage_Flat{
    Text:  m.Text ,
    PeerId:  int(m.PeerId) ,
    Time:  int(m.Time) ,

}
return r
}

func(m *PB_MsgResponse_BroadcastNewMessage)ToFlat() *PB_MsgResponse_BroadcastNewMessage_Flat {
r := &PB_MsgResponse_BroadcastNewMessage_Flat{
}
return r
}

func(m *PB_MsgParam_GetFreshChatList)ToFlat() *PB_MsgParam_GetFreshChatList_Flat {
r := &PB_MsgParam_GetFreshChatList_Flat{
}
return r
}

func(m *PB_MsgResponse_GetFreshChatList)ToFlat() *PB_MsgResponse_GetFreshChatList_Flat {
r := &PB_MsgResponse_GetFreshChatList_Flat{


}
return r
}

func(m *PB_MsgParam_GetFreshRoomMessagesList)ToFlat() *PB_MsgParam_GetFreshRoomMessagesList_Flat {
r := &PB_MsgParam_GetFreshRoomMessagesList_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    Last:  int(m.Last) ,
}
return r
}

func(m *PB_MsgResponse_GetFreshRoomMessagesList)ToFlat() *PB_MsgResponse_GetFreshRoomMessagesList_Flat {
r := &PB_MsgResponse_GetFreshRoomMessagesList_Flat{

    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MsgParam_Echo)ToFlat() *PB_MsgParam_Echo_Flat {
r := &PB_MsgParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_MsgResponse_PB_MsgParam_Echo)ToFlat() *PB_MsgResponse_PB_MsgParam_Echo_Flat {
r := &PB_MsgResponse_PB_MsgParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_SyncParam_GetDirectUpdates)ToFlat() *PB_SyncParam_GetDirectUpdates_Flat {
r := &PB_SyncParam_GetDirectUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetDirectUpdates)ToFlat() *PB_SyncResponse_GetDirectUpdates_Flat {
r := &PB_SyncResponse_GetDirectUpdates_Flat{












    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetGeneralUpdates)ToFlat() *PB_SyncParam_GetGeneralUpdates_Flat {
r := &PB_SyncParam_GetGeneralUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetGeneralUpdates)ToFlat() *PB_SyncResponse_GetGeneralUpdates_Flat {
r := &PB_SyncResponse_GetGeneralUpdates_Flat{


    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetNotifyUpdates)ToFlat() *PB_SyncParam_GetNotifyUpdates_Flat {
r := &PB_SyncParam_GetNotifyUpdates_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetNotifyUpdates)ToFlat() *PB_SyncResponse_GetNotifyUpdates_Flat {
r := &PB_SyncResponse_GetNotifyUpdates_Flat{

    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_SetLastSyncDirectUpdateId)ToFlat() *PB_SyncParam_SetLastSyncDirectUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncDirectUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncDirectUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncDirectUpdateId_Flat{
}
return r
}

func(m *PB_SyncParam_SetLastSyncGeneralUpdateId)ToFlat() *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncGeneralUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncGeneralUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat{
}
return r
}

func(m *PB_SyncParam_SetLastSyncNotifyUpdateId)ToFlat() *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat {
r := &PB_SyncParam_SetLastSyncNotifyUpdateId_Flat{
    LastId:  int(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncNotifyUpdateId)ToFlat() *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat {
r := &PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat{
}
return r
}

func(m *PB_UserParam_BlockUser)ToFlat() *PB_UserParam_BlockUser_Flat {
r := &PB_UserParam_BlockUser_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserOfflineResponse_BlockUser)ToFlat() *PB_UserOfflineResponse_BlockUser_Flat {
r := &PB_UserOfflineResponse_BlockUser_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UnBlockUser)ToFlat() *PB_UserParam_UnBlockUser_Flat {
r := &PB_UserParam_UnBlockUser_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}

func(m *PB_UserOfflineResponse_UnBlockUser)ToFlat() *PB_UserOfflineResponse_UnBlockUser_Flat {
r := &PB_UserOfflineResponse_UnBlockUser_Flat{

}
return r
}

func(m *PB_UserParam_BlockedList)ToFlat() *PB_UserParam_BlockedList_Flat {
r := &PB_UserParam_BlockedList_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockedList)ToFlat() *PB_UserResponse_BlockedList_Flat {
r := &PB_UserResponse_BlockedList_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UpdateAbout)ToFlat() *PB_UserParam_UpdateAbout_Flat {
r := &PB_UserParam_UpdateAbout_Flat{
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateAbout)ToFlat() *PB_UserOfflineResponse_UpdateAbout_Flat {
r := &PB_UserOfflineResponse_UpdateAbout_Flat{
    UserId:  int(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserParam_UpdateUserName)ToFlat() *PB_UserParam_UpdateUserName_Flat {
r := &PB_UserParam_UpdateUserName_Flat{
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateUserName)ToFlat() *PB_UserOfflineResponse_UpdateUserName_Flat {
r := &PB_UserOfflineResponse_UpdateUserName_Flat{
    UserId:  int(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserParam_ChangeAvatar)ToFlat() *PB_UserParam_ChangeAvatar_Flat {
r := &PB_UserParam_ChangeAvatar_Flat{
    None:  m.None ,
    ImageData2:  []byte(m.ImageData2) ,
}
return r
}

func(m *PB_UserOfflineResponse_ChangeAvatar)ToFlat() *PB_UserOfflineResponse_ChangeAvatar_Flat {
r := &PB_UserOfflineResponse_ChangeAvatar_Flat{
}
return r
}

func(m *PB_UserParam_ChangePrivacy)ToFlat() *PB_UserParam_ChangePrivacy_Flat {
r := &PB_UserParam_ChangePrivacy_Flat{

}
return r
}

func(m *PB_UserResponseOffline_ChangePrivacy)ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
r := &PB_UserResponseOffline_ChangePrivacy_Flat{
}
return r
}

func(m *PB_UserParam_CheckUserName)ToFlat() *PB_UserParam_CheckUserName_Flat {
r := &PB_UserParam_CheckUserName_Flat{

}
return r
}

func(m *PB_UserResponse_CheckUserName)ToFlat() *PB_UserResponse_CheckUserName_Flat {
r := &PB_UserResponse_CheckUserName_Flat{
}
return r
}

func(m *UserView)ToFlat() *UserView_Flat {
r := &UserView_Flat{
}
return r
}

func(m *PB_Chat)ToFlat() *PB_Chat_Flat {
r := &PB_Chat_Flat{
    ChatId:  int(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedTime:  int(m.CreatedTime) ,
    CurrentSeq:  int(m.CurrentSeq) ,
    UpdatedMs:  int(m.UpdatedMs) ,
}
return r
}

func(m *PB_DirectMessage)ToFlat() *PB_DirectMessage_Flat {
r := &PB_DirectMessage_Flat{
    MessageId:  int(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnum:  int(m.MessageTypeEnum) ,
    Text:  m.Text ,
    Time:  int(m.Time) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnum:  int(m.DeliviryStatusEnum) ,
}
return r
}

func(m *PB_MessageFile)ToFlat() *PB_MessageFile_Flat {
r := &PB_MessageFile_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
}
return r
}

func(m *PB_User)ToFlat() *PB_User_Flat {
r := &PB_User_Flat{
    Id:  int(m.Id) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    IsDeleted:  int(m.IsDeleted) ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,
    LikesCount:  int(m.LikesCount) ,
    ResharedCount:  int(m.ResharedCount) ,
    LastActionTime:  int(m.LastActionTime) ,
    LastPostTime:  int(m.LastPostTime) ,
    PrimaryFollowingList:  int(m.PrimaryFollowingList) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    SessionUuid:  m.SessionUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    LastLoginTime:  int(m.LastLoginTime) ,
    LastIpAddress:  m.LastIpAddress ,
}
return r
}

func(m *PB_UpdateGroupParticipants)ToFlat() *PB_UpdateGroupParticipants_Flat {
r := &PB_UpdateGroupParticipants_Flat{
}
return r
}

func(m *PB_UpdateNotifySettings)ToFlat() *PB_UpdateNotifySettings_Flat {
r := &PB_UpdateNotifySettings_Flat{
}
return r
}

func(m *PB_UpdateServiceNotification)ToFlat() *PB_UpdateServiceNotification_Flat {
r := &PB_UpdateServiceNotification_Flat{
}
return r
}

func(m *PB_UpdateMessageMeta)ToFlat() *PB_UpdateMessageMeta_Flat {
r := &PB_UpdateMessageMeta_Flat{
    MessageId:  int(m.MessageId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_UpdateMessageId)ToFlat() *PB_UpdateMessageId_Flat {
r := &PB_UpdateMessageId_Flat{
    OldMessageId:  int(m.OldMessageId) ,
    NewMessageId:  int(m.NewMessageId) ,
}
return r
}

func(m *PB_UpdateMessageToEdit)ToFlat() *PB_UpdateMessageToEdit_Flat {
r := &PB_UpdateMessageToEdit_Flat{
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_UpdateMessageToDelete)ToFlat() *PB_UpdateMessageToDelete_Flat {
r := &PB_UpdateMessageToDelete_Flat{
    MessageId:  int(m.MessageId) ,
}
return r
}

func(m *PB_UpdateRoomActionDoing)ToFlat() *PB_UpdateRoomActionDoing_Flat {
r := &PB_UpdateRoomActionDoing_Flat{
    RoomKey:  m.RoomKey ,

}
return r
}

func(m *PB_UpdateUserBlocked)ToFlat() *PB_UpdateUserBlocked_Flat {
r := &PB_UpdateUserBlocked_Flat{
    UserId:  int(m.UserId) ,
    Blocked:  m.Blocked ,
}
return r
}

func(m *PB_PushHolderView)ToFlat() *PB_PushHolderView_Flat {
r := &PB_PushHolderView_Flat{














}
return r
}

func(m *PB_DirectUpdatesView)ToFlat() *PB_DirectUpdatesView_Flat {
r := &PB_DirectUpdatesView_Flat{












}
return r
}

func(m *PB_GeneralUpdatesView)ToFlat() *PB_GeneralUpdatesView_Flat {
r := &PB_GeneralUpdatesView_Flat{


}
return r
}

func(m *PB_NotifyUpdatesView)ToFlat() *PB_NotifyUpdatesView_Flat {
r := &PB_NotifyUpdatesView_Flat{


}
return r
}

func(m *PB_ChatView)ToFlat() *PB_ChatView_Flat {
r := &PB_ChatView_Flat{
    ChatKey:  m.ChatKey ,
    ChatId:  int(m.ChatId) ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    LastMessageId:  int(m.LastMessageId) ,
    LastDeletedMessageId:  int(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    CurrentSeq:  int(m.CurrentSeq) ,

    SharedMediaCount:  int(m.SharedMediaCount) ,
    UnseenCount:  int(m.UnseenCount) ,


}
return r
}

func(m *PB_MessageView)ToFlat() *PB_MessageView_Flat {
r := &PB_MessageView_Flat{
    MessageId:  int(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    Time:  int(m.Time) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
    ChatId:  int(m.ChatId) ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int(m.RemoteId) ,

}
return r
}

func(m *PB_MessageFileView)ToFlat() *PB_MessageFileView_Flat {
r := &PB_MessageFileView_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    OriginalUserId:  int(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}

func(m *PB_UserView)ToFlat() *PB_UserView_Flat {
r := &PB_UserView_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    IsDeleted:  int(m.IsDeleted) ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    FollowingType:  int(m.FollowingType) ,
}
return r
}



///// from_flat ///

func(m *PB_RoomInlineView_Flat)ToPB() *PB_RoomInlineView {
r := &PB_RoomInlineView{
    RoomId:  int64(m.RoomId) ,

}
return r
}

func(m *PB_MessageForwardedFrom_Flat)ToPB() *PB_MessageForwardedFrom {
r := &PB_MessageForwardedFrom{
    RoomId:  int64(m.RoomId) ,

    MessageId:  int64(m.MessageId) ,
    MessageSeq:  int32(m.MessageSeq) ,
}
return r
}

func(m *PB_GroupView_Flat)ToPB() *PB_GroupView {
r := &PB_GroupView{
    GroupId:  int64(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int32(m.MembersCount) ,
    GroupPrivacyEnum:  int32(m.GroupPrivacyEnum) ,
    CreatorUserId:  int32(m.CreatorUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
}
return r
}

func(m *PB_GroupMemberView_Flat)ToPB() *PB_GroupMemberView {
r := &PB_GroupMemberView{
    Id:  int64(m.Id) ,
    GroupId:  int64(m.GroupId) ,
    UserId:  int32(m.UserId) ,
    ByUserId:  int32(m.ByUserId) ,
    GroupRoleEnum:  int32(m.GroupRoleEnum) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_MessageFileView__DEp_Flat)ToPB() *PB_MessageFileView__DEp {
r := &PB_MessageFileView__DEp{
    MessageFileId:  int64(m.MessageFileId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnum:  int32(m.FileTypeEnum) ,
    MimeType:  m.MimeType ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData64:  m.ThumbData64 ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_ReqLastChangesForTheRoom_Flat)ToPB() *PB_ReqLastChangesForTheRoom {
r := &PB_ReqLastChangesForTheRoom{
    RoomId:  int64(m.RoomId) ,
    LastLogId:  int64(m.LastLogId) ,
    LastHaveSeq:  int32(m.LastHaveSeq) ,
}
return r
}

func(m *PB_ResponseLastChangesForTheRoom_Flat)ToPB() *PB_ResponseLastChangesForTheRoom {
r := &PB_ResponseLastChangesForTheRoom{

}
return r
}

func(m *PB_RequestSetLastSeenMessages_Flat)ToPB() *PB_RequestSetLastSeenMessages {
r := &PB_RequestSetLastSeenMessages{
    RoomId:  int64(m.RoomId) ,
    FromMessageId:  int64(m.FromMessageId) ,
    ToMessageId:  int32(m.ToMessageId) ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}

func(m *PB_ResponseSetLastSeenMessages_Flat)ToPB() *PB_ResponseSetLastSeenMessages {
r := &PB_ResponseSetLastSeenMessages{

}
return r
}

func(m *PB_MessagesCollections_Flat)ToPB() *PB_MessagesCollections {
r := &PB_MessagesCollections{
    DirectMessagesIds:  helper.SliceIntToInt64(m.DirectMessagesIds) ,
    GroupMessagesIds:  helper.SliceIntToInt64(m.GroupMessagesIds) ,
    BroadCatMessagesIds:  helper.SliceIntToInt64(m.BroadCatMessagesIds) ,
}
return r
}

func(m *PB_DirectLogView_Flat)ToPB() *PB_DirectLogView {
r := &PB_DirectLogView{


}
return r
}

func(m *PB_DirectLog_Flat)ToPB() *PB_DirectLog {
r := &PB_DirectLog{
    Id:  int64(m.Id) ,
    ToUserId:  int32(m.ToUserId) ,
    MessageId:  int64(m.MessageId) ,
    ChatId:  int64(m.ChatId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    EventType:  int32(m.EventType) ,
    RoomLogTypeId:  int32(m.RoomLogTypeId) ,
    FromSeq:  int32(m.FromSeq) ,
    ToSeq:  int32(m.ToSeq) ,
    ExtraPB:  m.ExtraPB ,
    ExtraJson:  m.ExtraJson ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}

func(m *PB_PushDirectLogViewsMany_Flat)ToPB() *PB_PushDirectLogViewsMany {
r := &PB_PushDirectLogViewsMany{

}
return r
}

func(m *PB_UserWithMe_Flat)ToPB() *PB_UserWithMe {
r := &PB_UserWithMe{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    AppVersion:  int32(m.AppVersion) ,
    FollowingType:  int32(m.FollowingType) ,
}
return r
}

func(m *PB_Message_Flat)ToPB() *PB_Message {
r := &PB_Message{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int64(m.UserId) ,
    PeerId:  int64(m.PeerId) ,
    RoomTypeId:  int32(m.RoomTypeId) ,
    MessageTypeId:  int32(m.MessageTypeId) ,
    Text:  m.Text ,
    ExtraJson:  m.ExtraJson ,
    IsByMe:  int32(m.IsByMe) ,
    IsStared:  int32(m.IsStared) ,
    CreatedMs:  int64(m.CreatedMs) ,
    CreatedDeviceMs:  int64(m.CreatedDeviceMs) ,
    SortId:  int64(m.SortId) ,
    PeerSeenTime:  int64(m.PeerSeenTime) ,
    ServerReceivedTime:  int64(m.ServerReceivedTime) ,
    ServerDeletedTime:  int64(m.ServerDeletedTime) ,
    ISeenTime:  int64(m.ISeenTime) ,
    ToPush:  int32(m.ToPush) ,
    MsgFile_LocalSrc:  m.MsgFile_LocalSrc ,
    MsgFile_Status:  int32(m.MsgFile_Status) ,

}
return r
}

func(m *PB_MsgFile_Flat)ToPB() *PB_MsgFile {
r := &PB_MsgFile{
    Name:  m.Name ,
    Size:  int64(m.Size) ,
    FileType:  int32(m.FileType) ,
    MimeType:  m.MimeType ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    ThumbData:  m.ThumbData ,
    Data:  m.Data ,
    ServerSrc:  m.ServerSrc ,
}
return r
}

func(m *PB_Response_Flat)ToPB() *PB_Response {
r := &PB_Response{
    Status:  int32(m.Status) ,
}
return r
}

func(m *PB_Request_Flat)ToPB() *PB_Request {
r := &PB_Request{
    Status:  int32(m.Status) ,
}
return r
}

func(m *PB_RequestMsgAddMany_Flat)ToPB() *PB_RequestMsgAddMany {
r := &PB_RequestMsgAddMany{


}
return r
}

func(m *PB_ResponseMsgAddMany_Flat)ToPB() *PB_ResponseMsgAddMany {
r := &PB_ResponseMsgAddMany{

}
return r
}

func(m *PB_Push_Flat)ToPB() *PB_Push {
r := &PB_Push{
    Status:  int32(m.Status) ,
}
return r
}

func(m *PB_Result_Flat)ToPB() *PB_Result {
r := &PB_Result{
    Status:  int32(m.Status) ,
}
return r
}

func(m *PB_PushMsgAddMany_Flat)ToPB() *PB_PushMsgAddMany {
r := &PB_PushMsgAddMany{



}
return r
}

func(m *PB_ResultMsgAddMany_Flat)ToPB() *PB_ResultMsgAddMany {
r := &PB_ResultMsgAddMany{

}
return r
}

func(m *PB_MsgEvent_Flat)ToPB() *PB_MsgEvent {
r := &PB_MsgEvent{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    PeerUserId:  int64(m.PeerUserId) ,
    EventType:  int32(m.EventType) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_PushMsgEvents_Flat)ToPB() *PB_PushMsgEvents {
r := &PB_PushMsgEvents{


}
return r
}

func(m *PB_ResultMsgEvents_Flat)ToPB() *PB_ResultMsgEvents {
r := &PB_ResultMsgEvents{

}
return r
}

func(m *PB_MsgsSeenFromClient_Flat)ToPB() *PB_MsgsSeenFromClient {
r := &PB_MsgsSeenFromClient{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int64(m.UserId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_MsgSeen_Flat)ToPB() *PB_MsgSeen {
r := &PB_MsgSeen{
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int64(m.UserId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_RequestMsgsSeen_Flat)ToPB() *PB_RequestMsgsSeen {
r := &PB_RequestMsgsSeen{


}
return r
}

func(m *PB_ResponseMsgsSeen_Flat)ToPB() *PB_ResponseMsgsSeen {
r := &PB_ResponseMsgsSeen{

}
return r
}

func(m *PB_ReqRpcAddMsg_Flat)ToPB() *PB_ReqRpcAddMsg {
r := &PB_ReqRpcAddMsg{


}
return r
}

func(m *PB_ResRpcAddMsg_Flat)ToPB() *PB_ResRpcAddMsg {
r := &PB_ResRpcAddMsg{

    ServerSrc:  m.ServerSrc ,
}
return r
}

func(m *GeoLocation_Flat)ToPB() *GeoLocation {
r := &GeoLocation{
    Lat:  m.Lat ,
    Lon:  m.Lon ,
}
return r
}

func(m *RoomMessageLog_Flat)ToPB() *RoomMessageLog {
r := &RoomMessageLog{

    TargetUserId:  uint64(m.TargetUserId) ,
    ByUserId:  uint64(m.ByUserId) ,
}
return r
}

func(m *RoomMessageForwardFrom_Flat)ToPB() *RoomMessageForwardFrom {
r := &RoomMessageForwardFrom{
    RoomId:  uint64(m.RoomId) ,
    MessageId:  uint64(m.MessageId) ,
    RoomTypeEnum:  uint32(m.RoomTypeEnum) ,
}
return r
}

func(m *RoomDraft_Flat)ToPB() *RoomDraft {
r := &RoomDraft{
    Message:  m.Message ,
    ReplyTo:  uint64(m.ReplyTo) ,
}
return r
}

func(m *ChatRoom_Flat)ToPB() *ChatRoom {
r := &ChatRoom{
}
return r
}

func(m *Pagination_Flat)ToPB() *Pagination {
r := &Pagination{
    Offset:  uint32(m.Offset) ,
    Limit:  uint32(m.Limit) ,
}
return r
}

func(m *PB_CommandToServer_Flat)ToPB() *PB_CommandToServer {
r := &PB_CommandToServer{
    ClientCallId:  int64(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}

func(m *PB_CommandToClient_Flat)ToPB() *PB_CommandToClient {
r := &PB_CommandToClient{
    ServerCallId:  int64(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}

func(m *PB_CommandReachedToServer_Flat)ToPB() *PB_CommandReachedToServer {
r := &PB_CommandReachedToServer{
    ClientCallId:  int64(m.ClientCallId) ,
}
return r
}

func(m *PB_CommandReachedToClient_Flat)ToPB() *PB_CommandReachedToClient {
r := &PB_CommandReachedToClient{
    ServerCallId:  int64(m.ServerCallId) ,
}
return r
}

func(m *PB_ResponseToClient_Flat)ToPB() *PB_ResponseToClient {
r := &PB_ResponseToClient{
    ClientCallId:  int64(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  m.Data ,
}
return r
}

func(m *PB_UserParam_CheckUserName2_Flat)ToPB() *PB_UserParam_CheckUserName2 {
r := &PB_UserParam_CheckUserName2{
}
return r
}

func(m *PB_UserResponse_CheckUserName2_Flat)ToPB() *PB_UserResponse_CheckUserName2 {
r := &PB_UserResponse_CheckUserName2{
}
return r
}

func(m *PB_MsgParam_AddNewTextMessage_Flat)ToPB() *PB_MsgParam_AddNewTextMessage {
r := &PB_MsgParam_AddNewTextMessage{
    Text:  m.Text ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,

}
return r
}

func(m *PB_MsgResponse_AddNewTextMessage_Flat)ToPB() *PB_MsgResponse_AddNewTextMessage {
r := &PB_MsgResponse_AddNewTextMessage{
}
return r
}

func(m *PB_MsgParam_AddNewMessage_Flat)ToPB() *PB_MsgParam_AddNewMessage {
r := &PB_MsgParam_AddNewMessage{
    Text:  m.Text ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,


    Blob:  m.Blob ,

}
return r
}

func(m *PB_MsgResponse_AddNewMessage_Flat)ToPB() *PB_MsgResponse_AddNewMessage {
r := &PB_MsgResponse_AddNewMessage{
}
return r
}

func(m *PB_MsgParam_SetRoomActionDoing_Flat)ToPB() *PB_MsgParam_SetRoomActionDoing {
r := &PB_MsgParam_SetRoomActionDoing{
    GroupId:  int64(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,

}
return r
}

func(m *PB_MsgResponse_SetRoomActionDoing_Flat)ToPB() *PB_MsgResponse_SetRoomActionDoing {
r := &PB_MsgResponse_SetRoomActionDoing{
}
return r
}

func(m *PB_MsgParam_GetMessagesByIds_Flat)ToPB() *PB_MsgParam_GetMessagesByIds {
r := &PB_MsgParam_GetMessagesByIds{

}
return r
}

func(m *PB_MsgResponse_GetMessagesByIds_Flat)ToPB() *PB_MsgResponse_GetMessagesByIds {
r := &PB_MsgResponse_GetMessagesByIds{

}
return r
}

func(m *PB_MsgParam_GetMessagesHistory_Flat)ToPB() *PB_MsgParam_GetMessagesHistory {
r := &PB_MsgParam_GetMessagesHistory{
    ChatKey:  m.ChatKey ,
    FromSeq:  int32(m.FromSeq) ,
    EndSeq:  int32(m.EndSeq) ,
}
return r
}

func(m *PB_MsgResponse_GetMessagesHistory_Flat)ToPB() *PB_MsgResponse_GetMessagesHistory {
r := &PB_MsgResponse_GetMessagesHistory{

}
return r
}

func(m *PB_MsgParam_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_MsgParam_SetChatMessagesRangeAsSeen {
r := &PB_MsgParam_SetChatMessagesRangeAsSeen{
    ChatKey:  m.ChatKey ,
    FromSeq:  int32(m.FromSeq) ,
    EndSeq:  int32(m.EndSeq) ,
    SeenTimeMs:  int64(m.SeenTimeMs) ,
}
return r
}

func(m *PB_MsgResponse_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_MsgResponse_SetChatMessagesRangeAsSeen {
r := &PB_MsgResponse_SetChatMessagesRangeAsSeen{
}
return r
}

func(m *PB_MsgParam_DeleteChatHistory_Flat)ToPB() *PB_MsgParam_DeleteChatHistory {
r := &PB_MsgParam_DeleteChatHistory{
    ChatKey:  m.ChatKey ,
    ToSeq:  int32(m.ToSeq) ,
}
return r
}

func(m *PB_MsgResponse_DeleteChatHistory_Flat)ToPB() *PB_MsgResponse_DeleteChatHistory {
r := &PB_MsgResponse_DeleteChatHistory{
}
return r
}

func(m *PB_MsgParam_DeleteMessagesByIds_Flat)ToPB() *PB_MsgParam_DeleteMessagesByIds {
r := &PB_MsgParam_DeleteMessagesByIds{

}
return r
}

func(m *PB_MsgResponse_DeleteMessagesByIds_Flat)ToPB() *PB_MsgResponse_DeleteMessagesByIds {
r := &PB_MsgResponse_DeleteMessagesByIds{
}
return r
}

func(m *PB_MsgParam_SetMessagesAsReceived_Flat)ToPB() *PB_MsgParam_SetMessagesAsReceived {
r := &PB_MsgParam_SetMessagesAsReceived{

}
return r
}

func(m *PB_MsgResponse_SetMessagesAsReceived_Flat)ToPB() *PB_MsgResponse_SetMessagesAsReceived {
r := &PB_MsgResponse_SetMessagesAsReceived{
}
return r
}

func(m *PB_MsgParam_ForwardMessages_Flat)ToPB() *PB_MsgParam_ForwardMessages {
r := &PB_MsgParam_ForwardMessages{

    ToDirectChatKeys:  m.ToDirectChatKeys ,
    ToGroupChatIds:  helper.SliceIntToInt64(m.ToGroupChatIds) ,
}
return r
}

func(m *PB_MsgResponse_ForwardMessages_Flat)ToPB() *PB_MsgResponse_ForwardMessages {
r := &PB_MsgResponse_ForwardMessages{
}
return r
}

func(m *PB_MsgParam_EditMessage_Flat)ToPB() *PB_MsgParam_EditMessage {
r := &PB_MsgParam_EditMessage{
    ChatKey:  m.ChatKey ,

    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_MsgResponse_EditMessage_Flat)ToPB() *PB_MsgResponse_EditMessage {
r := &PB_MsgResponse_EditMessage{
}
return r
}

func(m *PB_MsgParam_BroadcastNewMessage_Flat)ToPB() *PB_MsgParam_BroadcastNewMessage {
r := &PB_MsgParam_BroadcastNewMessage{
    Text:  m.Text ,
    PeerId:  int32(m.PeerId) ,
    Time:  int32(m.Time) ,

}
return r
}

func(m *PB_MsgResponse_BroadcastNewMessage_Flat)ToPB() *PB_MsgResponse_BroadcastNewMessage {
r := &PB_MsgResponse_BroadcastNewMessage{
}
return r
}

func(m *PB_MsgParam_GetFreshChatList_Flat)ToPB() *PB_MsgParam_GetFreshChatList {
r := &PB_MsgParam_GetFreshChatList{
}
return r
}

func(m *PB_MsgResponse_GetFreshChatList_Flat)ToPB() *PB_MsgResponse_GetFreshChatList {
r := &PB_MsgResponse_GetFreshChatList{


}
return r
}

func(m *PB_MsgParam_GetFreshRoomMessagesList_Flat)ToPB() *PB_MsgParam_GetFreshRoomMessagesList {
r := &PB_MsgParam_GetFreshRoomMessagesList{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    Last:  int64(m.Last) ,
}
return r
}

func(m *PB_MsgResponse_GetFreshRoomMessagesList_Flat)ToPB() *PB_MsgResponse_GetFreshRoomMessagesList {
r := &PB_MsgResponse_GetFreshRoomMessagesList{

    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_MsgParam_Echo_Flat)ToPB() *PB_MsgParam_Echo {
r := &PB_MsgParam_Echo{
    Text:  m.Text ,
}
return r
}

func(m *PB_MsgResponse_PB_MsgParam_Echo_Flat)ToPB() *PB_MsgResponse_PB_MsgParam_Echo {
r := &PB_MsgResponse_PB_MsgParam_Echo{
    Text:  m.Text ,
}
return r
}

func(m *PB_SyncParam_GetDirectUpdates_Flat)ToPB() *PB_SyncParam_GetDirectUpdates {
r := &PB_SyncParam_GetDirectUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetDirectUpdates_Flat)ToPB() *PB_SyncResponse_GetDirectUpdates {
r := &PB_SyncResponse_GetDirectUpdates{












    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetGeneralUpdates_Flat)ToPB() *PB_SyncParam_GetGeneralUpdates {
r := &PB_SyncParam_GetGeneralUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetGeneralUpdates_Flat)ToPB() *PB_SyncResponse_GetGeneralUpdates {
r := &PB_SyncResponse_GetGeneralUpdates{


    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_GetNotifyUpdates_Flat)ToPB() *PB_SyncParam_GetNotifyUpdates {
r := &PB_SyncParam_GetNotifyUpdates{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_GetNotifyUpdates_Flat)ToPB() *PB_SyncResponse_GetNotifyUpdates {
r := &PB_SyncResponse_GetNotifyUpdates{

    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncParam_SetLastSyncDirectUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncDirectUpdateId {
r := &PB_SyncParam_SetLastSyncDirectUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncDirectUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncDirectUpdateId {
r := &PB_SyncResponse_SetLastSyncDirectUpdateId{
}
return r
}

func(m *PB_SyncParam_SetLastSyncGeneralUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncGeneralUpdateId {
r := &PB_SyncParam_SetLastSyncGeneralUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncGeneralUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncGeneralUpdateId {
r := &PB_SyncResponse_SetLastSyncGeneralUpdateId{
}
return r
}

func(m *PB_SyncParam_SetLastSyncNotifyUpdateId_Flat)ToPB() *PB_SyncParam_SetLastSyncNotifyUpdateId {
r := &PB_SyncParam_SetLastSyncNotifyUpdateId{
    LastId:  int64(m.LastId) ,
}
return r
}

func(m *PB_SyncResponse_SetLastSyncNotifyUpdateId_Flat)ToPB() *PB_SyncResponse_SetLastSyncNotifyUpdateId {
r := &PB_SyncResponse_SetLastSyncNotifyUpdateId{
}
return r
}

func(m *PB_UserParam_BlockUser_Flat)ToPB() *PB_UserParam_BlockUser {
r := &PB_UserParam_BlockUser{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserOfflineResponse_BlockUser_Flat)ToPB() *PB_UserOfflineResponse_BlockUser {
r := &PB_UserOfflineResponse_BlockUser{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UnBlockUser_Flat)ToPB() *PB_UserParam_UnBlockUser {
r := &PB_UserParam_UnBlockUser{
    Offset:  int32(m.Offset) ,
    Limit:  int32(m.Limit) ,
}
return r
}

func(m *PB_UserOfflineResponse_UnBlockUser_Flat)ToPB() *PB_UserOfflineResponse_UnBlockUser {
r := &PB_UserOfflineResponse_UnBlockUser{

}
return r
}

func(m *PB_UserParam_BlockedList_Flat)ToPB() *PB_UserParam_BlockedList {
r := &PB_UserParam_BlockedList{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockedList_Flat)ToPB() *PB_UserResponse_BlockedList {
r := &PB_UserResponse_BlockedList{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UpdateAbout_Flat)ToPB() *PB_UserParam_UpdateAbout {
r := &PB_UserParam_UpdateAbout{
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateAbout_Flat)ToPB() *PB_UserOfflineResponse_UpdateAbout {
r := &PB_UserOfflineResponse_UpdateAbout{
    UserId:  int32(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserParam_UpdateUserName_Flat)ToPB() *PB_UserParam_UpdateUserName {
r := &PB_UserParam_UpdateUserName{
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserOfflineResponse_UpdateUserName_Flat)ToPB() *PB_UserOfflineResponse_UpdateUserName {
r := &PB_UserOfflineResponse_UpdateUserName{
    UserId:  int32(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserParam_ChangeAvatar_Flat)ToPB() *PB_UserParam_ChangeAvatar {
r := &PB_UserParam_ChangeAvatar{
    None:  m.None ,
    ImageData2:  m.ImageData2 ,
}
return r
}

func(m *PB_UserOfflineResponse_ChangeAvatar_Flat)ToPB() *PB_UserOfflineResponse_ChangeAvatar {
r := &PB_UserOfflineResponse_ChangeAvatar{
}
return r
}

func(m *PB_UserParam_ChangePrivacy_Flat)ToPB() *PB_UserParam_ChangePrivacy {
r := &PB_UserParam_ChangePrivacy{

}
return r
}

func(m *PB_UserResponseOffline_ChangePrivacy_Flat)ToPB() *PB_UserResponseOffline_ChangePrivacy {
r := &PB_UserResponseOffline_ChangePrivacy{
}
return r
}

func(m *PB_UserParam_CheckUserName_Flat)ToPB() *PB_UserParam_CheckUserName {
r := &PB_UserParam_CheckUserName{

}
return r
}

func(m *PB_UserResponse_CheckUserName_Flat)ToPB() *PB_UserResponse_CheckUserName {
r := &PB_UserResponse_CheckUserName{
}
return r
}

func(m *UserView_Flat)ToPB() *UserView {
r := &UserView{
}
return r
}

func(m *PB_Chat_Flat)ToPB() *PB_Chat {
r := &PB_Chat{
    ChatId:  int64(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
}
return r
}

func(m *PB_DirectMessage_Flat)ToPB() *PB_DirectMessage {
r := &PB_DirectMessage{
    MessageId:  int64(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnum:  int32(m.MessageTypeEnum) ,
    Text:  m.Text ,
    Time:  int32(m.Time) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnum:  int32(m.DeliviryStatusEnum) ,
}
return r
}

func(m *PB_MessageFile_Flat)ToPB() *PB_MessageFile {
r := &PB_MessageFile{
    MessageFileId:  int64(m.MessageFileId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
}
return r
}

func(m *PB_User_Flat)ToPB() *PB_User {
r := &PB_User{
    Id:  int32(m.Id) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    IsDeleted:  int32(m.IsDeleted) ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ResharedCount:  int32(m.ResharedCount) ,
    LastActionTime:  int32(m.LastActionTime) ,
    LastPostTime:  int32(m.LastPostTime) ,
    PrimaryFollowingList:  int32(m.PrimaryFollowingList) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    SessionUuid:  m.SessionUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    LastLoginTime:  int32(m.LastLoginTime) ,
    LastIpAddress:  m.LastIpAddress ,
}
return r
}

func(m *PB_UpdateGroupParticipants_Flat)ToPB() *PB_UpdateGroupParticipants {
r := &PB_UpdateGroupParticipants{
}
return r
}

func(m *PB_UpdateNotifySettings_Flat)ToPB() *PB_UpdateNotifySettings {
r := &PB_UpdateNotifySettings{
}
return r
}

func(m *PB_UpdateServiceNotification_Flat)ToPB() *PB_UpdateServiceNotification {
r := &PB_UpdateServiceNotification{
}
return r
}

func(m *PB_UpdateMessageMeta_Flat)ToPB() *PB_UpdateMessageMeta {
r := &PB_UpdateMessageMeta{
    MessageId:  int64(m.MessageId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_UpdateMessageId_Flat)ToPB() *PB_UpdateMessageId {
r := &PB_UpdateMessageId{
    OldMessageId:  int64(m.OldMessageId) ,
    NewMessageId:  int64(m.NewMessageId) ,
}
return r
}

func(m *PB_UpdateMessageToEdit_Flat)ToPB() *PB_UpdateMessageToEdit {
r := &PB_UpdateMessageToEdit{
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_UpdateMessageToDelete_Flat)ToPB() *PB_UpdateMessageToDelete {
r := &PB_UpdateMessageToDelete{
    MessageId:  int64(m.MessageId) ,
}
return r
}

func(m *PB_UpdateRoomActionDoing_Flat)ToPB() *PB_UpdateRoomActionDoing {
r := &PB_UpdateRoomActionDoing{
    RoomKey:  m.RoomKey ,

}
return r
}

func(m *PB_UpdateUserBlocked_Flat)ToPB() *PB_UpdateUserBlocked {
r := &PB_UpdateUserBlocked{
    UserId:  int32(m.UserId) ,
    Blocked:  m.Blocked ,
}
return r
}

func(m *PB_PushHolderView_Flat)ToPB() *PB_PushHolderView {
r := &PB_PushHolderView{














}
return r
}

func(m *PB_DirectUpdatesView_Flat)ToPB() *PB_DirectUpdatesView {
r := &PB_DirectUpdatesView{












}
return r
}

func(m *PB_GeneralUpdatesView_Flat)ToPB() *PB_GeneralUpdatesView {
r := &PB_GeneralUpdatesView{


}
return r
}

func(m *PB_NotifyUpdatesView_Flat)ToPB() *PB_NotifyUpdatesView {
r := &PB_NotifyUpdatesView{


}
return r
}

func(m *PB_ChatView_Flat)ToPB() *PB_ChatView {
r := &PB_ChatView{
    ChatKey:  m.ChatKey ,
    ChatId:  int64(m.ChatId) ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    LastMessageId:  int64(m.LastMessageId) ,
    LastDeletedMessageId:  int64(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    CurrentSeq:  int32(m.CurrentSeq) ,

    SharedMediaCount:  int32(m.SharedMediaCount) ,
    UnseenCount:  int32(m.UnseenCount) ,


}
return r
}

func(m *PB_MessageView_Flat)ToPB() *PB_MessageView {
r := &PB_MessageView{
    MessageId:  uint64(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    Time:  int32(m.Time) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
    ChatId:  int64(m.ChatId) ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int64(m.RemoteId) ,

}
return r
}

func(m *PB_MessageFileView_Flat)ToPB() *PB_MessageFileView {
r := &PB_MessageFileView{
    MessageFileId:  int64(m.MessageFileId) ,
    OriginalUserId:  int32(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int64(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}

func(m *PB_UserView_Flat)ToPB() *PB_UserView {
r := &PB_UserView{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    IsDeleted:  int32(m.IsDeleted) ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    FollowingType:  int32(m.FollowingType) ,
}
return r
}



///// folding ///

var PB_RoomInlineView__FOlD = &PB_RoomInlineView{
        RoomId:  0 ,

}


var PB_MessageForwardedFrom__FOlD = &PB_MessageForwardedFrom{
        RoomId:  0 ,

        MessageId:  0 ,
        MessageSeq:  0 ,
}


var PB_GroupView__FOlD = &PB_GroupView{
        GroupId:  0 ,
        GroupName:  "" ,
        MembersCount:  0 ,
        GroupPrivacyEnum:  0 ,
        CreatorUserId:  0 ,
        CreatedTime:  0 ,
        UpdatedMs:  0 ,
        CurrentSeq:  0 ,
}


var PB_GroupMemberView__FOlD = &PB_GroupMemberView{
        Id:  0 ,
        GroupId:  0 ,
        UserId:  0 ,
        ByUserId:  0 ,
        GroupRoleEnum:  0 ,
        CreatedTime:  0 ,
}


var PB_MessageFileView__DEp__FOlD = &PB_MessageFileView__DEp{
        MessageFileId:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileTypeEnum:  0 ,
        MimeType:  "" ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        ThumbData64:  "" ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerThumbPath:  "" ,
        BucketId:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
        CreatedTime:  0 ,
}


var PB_ReqLastChangesForTheRoom__FOlD = &PB_ReqLastChangesForTheRoom{
        RoomId:  0 ,
        LastLogId:  0 ,
        LastHaveSeq:  0 ,
}


var PB_ResponseLastChangesForTheRoom__FOlD = &PB_ResponseLastChangesForTheRoom{

}


var PB_RequestSetLastSeenMessages__FOlD = &PB_RequestSetLastSeenMessages{
        RoomId:  0 ,
        FromMessageId:  0 ,
        ToMessageId:  0 ,
        AtTimeMs:  0 ,
}


var PB_ResponseSetLastSeenMessages__FOlD = &PB_ResponseSetLastSeenMessages{

}


var PB_MessagesCollections__FOlD = &PB_MessagesCollections{
        DirectMessagesIds:  0 ,
        GroupMessagesIds:  0 ,
        BroadCatMessagesIds:  0 ,
}


var PB_DirectLogView__FOlD = &PB_DirectLogView{


}


var PB_DirectLog__FOlD = &PB_DirectLog{
        Id:  0 ,
        ToUserId:  0 ,
        MessageId:  0 ,
        ChatId:  0 ,
        PeerUserId:  0 ,
        EventType:  0 ,
        RoomLogTypeId:  0 ,
        FromSeq:  0 ,
        ToSeq:  0 ,
        ExtraPB:  []byte{} ,
        ExtraJson:  "" ,
        AtTimeMs:  0 ,
}


var PB_PushDirectLogViewsMany__FOlD = &PB_PushDirectLogViewsMany{

}


var PB_UserWithMe__FOlD = &PB_UserWithMe{
        UserId:  0 ,
        UserName:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        About:  "" ,
        FullName:  "" ,
        AvatarUrl:  "" ,
        PrivacyProfile:  0 ,
        Phone:  "" ,
        UpdatedTime:  0 ,
        AppVersion:  0 ,
        FollowingType:  0 ,
}


var PB_Message__FOlD = &PB_Message{
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        PeerId:  0 ,
        RoomTypeId:  0 ,
        MessageTypeId:  0 ,
        Text:  "" ,
        ExtraJson:  "" ,
        IsByMe:  0 ,
        IsStared:  0 ,
        CreatedMs:  0 ,
        CreatedDeviceMs:  0 ,
        SortId:  0 ,
        PeerSeenTime:  0 ,
        ServerReceivedTime:  0 ,
        ServerDeletedTime:  0 ,
        ISeenTime:  0 ,
        ToPush:  0 ,
        MsgFile_LocalSrc:  "" ,
        MsgFile_Status:  0 ,

}


var PB_MsgFile__FOlD = &PB_MsgFile{
        Name:  "" ,
        Size:  0 ,
        FileType:  0 ,
        MimeType:  "" ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        ThumbData:  []byte{} ,
        Data:  []byte{} ,
        ServerSrc:  "" ,
}


var PB_Response__FOlD = &PB_Response{
        Status:  0 ,
}


var PB_Request__FOlD = &PB_Request{
        Status:  0 ,
}


var PB_RequestMsgAddMany__FOlD = &PB_RequestMsgAddMany{


}


var PB_ResponseMsgAddMany__FOlD = &PB_ResponseMsgAddMany{

}


var PB_Push__FOlD = &PB_Push{
        Status:  0 ,
}


var PB_Result__FOlD = &PB_Result{
        Status:  0 ,
}


var PB_PushMsgAddMany__FOlD = &PB_PushMsgAddMany{



}


var PB_ResultMsgAddMany__FOlD = &PB_ResultMsgAddMany{

}


var PB_MsgEvent__FOlD = &PB_MsgEvent{
        MessageKey:  "" ,
        RoomKey:  "" ,
        PeerUserId:  0 ,
        EventType:  0 ,
        AtTime:  0 ,
}


var PB_PushMsgEvents__FOlD = &PB_PushMsgEvents{


}


var PB_ResultMsgEvents__FOlD = &PB_ResultMsgEvents{

}


var PB_MsgsSeenFromClient__FOlD = &PB_MsgsSeenFromClient{
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        AtTime:  0 ,
}


var PB_MsgSeen__FOlD = &PB_MsgSeen{
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        AtTime:  0 ,
}


var PB_RequestMsgsSeen__FOlD = &PB_RequestMsgsSeen{


}


var PB_ResponseMsgsSeen__FOlD = &PB_ResponseMsgsSeen{

}


var PB_ReqRpcAddMsg__FOlD = &PB_ReqRpcAddMsg{


}


var PB_ResRpcAddMsg__FOlD = &PB_ResRpcAddMsg{

        ServerSrc:  "" ,
}


var GeoLocation__FOlD = &GeoLocation{
        Lat:  0.0 ,
        Lon:  0.0 ,
}


var RoomMessageLog__FOlD = &RoomMessageLog{

        TargetUserId:  0 ,
        ByUserId:  0 ,
}


var RoomMessageForwardFrom__FOlD = &RoomMessageForwardFrom{
        RoomId:  0 ,
        MessageId:  0 ,
        RoomTypeEnum:  0 ,
}


var RoomDraft__FOlD = &RoomDraft{
        Message:  "" ,
        ReplyTo:  0 ,
}


var ChatRoom__FOlD = &ChatRoom{
}


var Pagination__FOlD = &Pagination{
        Offset:  0 ,
        Limit:  0 ,
}


var PB_CommandToServer__FOlD = &PB_CommandToServer{
        ClientCallId:  0 ,
        Command:  "" ,
        RespondReached:  false ,
        Data:  []byte{} ,
}


var PB_CommandToClient__FOlD = &PB_CommandToClient{
        ServerCallId:  0 ,
        Command:  "" ,
        RespondReached:  false ,
        Data:  []byte{} ,
}


var PB_CommandReachedToServer__FOlD = &PB_CommandReachedToServer{
        ClientCallId:  0 ,
}


var PB_CommandReachedToClient__FOlD = &PB_CommandReachedToClient{
        ServerCallId:  0 ,
}


var PB_ResponseToClient__FOlD = &PB_ResponseToClient{
        ClientCallId:  0 ,
        PBClass:  "" ,
        RpcFullName:  "" ,
        Data:  []byte{} ,
}


var PB_UserParam_CheckUserName2__FOlD = &PB_UserParam_CheckUserName2{
}


var PB_UserResponse_CheckUserName2__FOlD = &PB_UserResponse_CheckUserName2{
}


var PB_MsgParam_AddNewTextMessage__FOlD = &PB_MsgParam_AddNewTextMessage{
        Text:  "" ,
        PeerId:  0 ,
        Time:  0 ,
        ReplyToMessageId:  0 ,

}


var PB_MsgResponse_AddNewTextMessage__FOlD = &PB_MsgResponse_AddNewTextMessage{
}


var PB_MsgParam_AddNewMessage__FOlD = &PB_MsgParam_AddNewMessage{
        Text:  "" ,
        PeerId:  0 ,
        Time:  0 ,
        ReplyToMessageId:  0 ,


        Blob:  []byte{} ,

}


var PB_MsgResponse_AddNewMessage__FOlD = &PB_MsgResponse_AddNewMessage{
}


var PB_MsgParam_SetRoomActionDoing__FOlD = &PB_MsgParam_SetRoomActionDoing{
        GroupId:  0 ,
        DirectRoomKey:  "" ,

}


var PB_MsgResponse_SetRoomActionDoing__FOlD = &PB_MsgResponse_SetRoomActionDoing{
}


var PB_MsgParam_GetMessagesByIds__FOlD = &PB_MsgParam_GetMessagesByIds{

}


var PB_MsgResponse_GetMessagesByIds__FOlD = &PB_MsgResponse_GetMessagesByIds{

}


var PB_MsgParam_GetMessagesHistory__FOlD = &PB_MsgParam_GetMessagesHistory{
        ChatKey:  "" ,
        FromSeq:  0 ,
        EndSeq:  0 ,
}


var PB_MsgResponse_GetMessagesHistory__FOlD = &PB_MsgResponse_GetMessagesHistory{

}


var PB_MsgParam_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgParam_SetChatMessagesRangeAsSeen{
        ChatKey:  "" ,
        FromSeq:  0 ,
        EndSeq:  0 ,
        SeenTimeMs:  0 ,
}


var PB_MsgResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_MsgResponse_SetChatMessagesRangeAsSeen{
}


var PB_MsgParam_DeleteChatHistory__FOlD = &PB_MsgParam_DeleteChatHistory{
        ChatKey:  "" ,
        ToSeq:  0 ,
}


var PB_MsgResponse_DeleteChatHistory__FOlD = &PB_MsgResponse_DeleteChatHistory{
}


var PB_MsgParam_DeleteMessagesByIds__FOlD = &PB_MsgParam_DeleteMessagesByIds{

}


var PB_MsgResponse_DeleteMessagesByIds__FOlD = &PB_MsgResponse_DeleteMessagesByIds{
}


var PB_MsgParam_SetMessagesAsReceived__FOlD = &PB_MsgParam_SetMessagesAsReceived{

}


var PB_MsgResponse_SetMessagesAsReceived__FOlD = &PB_MsgResponse_SetMessagesAsReceived{
}


var PB_MsgParam_ForwardMessages__FOlD = &PB_MsgParam_ForwardMessages{

        ToDirectChatKeys:  "" ,
        ToGroupChatIds:  0 ,
}


var PB_MsgResponse_ForwardMessages__FOlD = &PB_MsgResponse_ForwardMessages{
}


var PB_MsgParam_EditMessage__FOlD = &PB_MsgParam_EditMessage{
        ChatKey:  "" ,

        MessageId:  0 ,
        NewText:  "" ,
}


var PB_MsgResponse_EditMessage__FOlD = &PB_MsgResponse_EditMessage{
}


var PB_MsgParam_BroadcastNewMessage__FOlD = &PB_MsgParam_BroadcastNewMessage{
        Text:  "" ,
        PeerId:  0 ,
        Time:  0 ,

}


var PB_MsgResponse_BroadcastNewMessage__FOlD = &PB_MsgResponse_BroadcastNewMessage{
}


var PB_MsgParam_GetFreshChatList__FOlD = &PB_MsgParam_GetFreshChatList{
}


var PB_MsgResponse_GetFreshChatList__FOlD = &PB_MsgResponse_GetFreshChatList{


}


var PB_MsgParam_GetFreshRoomMessagesList__FOlD = &PB_MsgParam_GetFreshRoomMessagesList{
        ChatKey:  "" ,
        RoomKey:  "" ,
        Last:  0 ,
}


var PB_MsgResponse_GetFreshRoomMessagesList__FOlD = &PB_MsgResponse_GetFreshRoomMessagesList{

        HasMore:  false ,
}


var PB_MsgParam_Echo__FOlD = &PB_MsgParam_Echo{
        Text:  "" ,
}


var PB_MsgResponse_PB_MsgParam_Echo__FOlD = &PB_MsgResponse_PB_MsgParam_Echo{
        Text:  "" ,
}


var PB_SyncParam_GetDirectUpdates__FOlD = &PB_SyncParam_GetDirectUpdates{
        LastId:  0 ,
}


var PB_SyncResponse_GetDirectUpdates__FOlD = &PB_SyncResponse_GetDirectUpdates{












        LastId:  0 ,
}


var PB_SyncParam_GetGeneralUpdates__FOlD = &PB_SyncParam_GetGeneralUpdates{
        LastId:  0 ,
}


var PB_SyncResponse_GetGeneralUpdates__FOlD = &PB_SyncResponse_GetGeneralUpdates{


        LastId:  0 ,
}


var PB_SyncParam_GetNotifyUpdates__FOlD = &PB_SyncParam_GetNotifyUpdates{
        LastId:  0 ,
}


var PB_SyncResponse_GetNotifyUpdates__FOlD = &PB_SyncResponse_GetNotifyUpdates{

        LastId:  0 ,
}


var PB_SyncParam_SetLastSyncDirectUpdateId__FOlD = &PB_SyncParam_SetLastSyncDirectUpdateId{
        LastId:  0 ,
}


var PB_SyncResponse_SetLastSyncDirectUpdateId__FOlD = &PB_SyncResponse_SetLastSyncDirectUpdateId{
}


var PB_SyncParam_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncParam_SetLastSyncGeneralUpdateId{
        LastId:  0 ,
}


var PB_SyncResponse_SetLastSyncGeneralUpdateId__FOlD = &PB_SyncResponse_SetLastSyncGeneralUpdateId{
}


var PB_SyncParam_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncParam_SetLastSyncNotifyUpdateId{
        LastId:  0 ,
}


var PB_SyncResponse_SetLastSyncNotifyUpdateId__FOlD = &PB_SyncResponse_SetLastSyncNotifyUpdateId{
}


var PB_UserParam_BlockUser__FOlD = &PB_UserParam_BlockUser{
        UserId:  0 ,
        UserName:  "" ,
}


var PB_UserOfflineResponse_BlockUser__FOlD = &PB_UserOfflineResponse_BlockUser{
        ByUserId:  0 ,
        TargetUserId:  0 ,
        TargetUserName:  "" ,
}


var PB_UserParam_UnBlockUser__FOlD = &PB_UserParam_UnBlockUser{
        Offset:  0 ,
        Limit:  0 ,
}


var PB_UserOfflineResponse_UnBlockUser__FOlD = &PB_UserOfflineResponse_UnBlockUser{

}


var PB_UserParam_BlockedList__FOlD = &PB_UserParam_BlockedList{
        UserId:  0 ,
        UserName:  "" ,
}


var PB_UserResponse_BlockedList__FOlD = &PB_UserResponse_BlockedList{
        ByUserId:  0 ,
        TargetUserId:  0 ,
        TargetUserName:  "" ,
}


var PB_UserParam_UpdateAbout__FOlD = &PB_UserParam_UpdateAbout{
        NewAbout:  "" ,
}


var PB_UserOfflineResponse_UpdateAbout__FOlD = &PB_UserOfflineResponse_UpdateAbout{
        UserId:  0 ,
        NewAbout:  "" ,
}


var PB_UserParam_UpdateUserName__FOlD = &PB_UserParam_UpdateUserName{
        NewUserName:  "" ,
}


var PB_UserOfflineResponse_UpdateUserName__FOlD = &PB_UserOfflineResponse_UpdateUserName{
        UserId:  0 ,
        NewUserName:  "" ,
}


var PB_UserParam_ChangeAvatar__FOlD = &PB_UserParam_ChangeAvatar{
        None:  false ,
        ImageData2:  []byte{} ,
}


var PB_UserOfflineResponse_ChangeAvatar__FOlD = &PB_UserOfflineResponse_ChangeAvatar{
}


var PB_UserParam_ChangePrivacy__FOlD = &PB_UserParam_ChangePrivacy{

}


var PB_UserResponseOffline_ChangePrivacy__FOlD = &PB_UserResponseOffline_ChangePrivacy{
}


var PB_UserParam_CheckUserName__FOlD = &PB_UserParam_CheckUserName{

}


var PB_UserResponse_CheckUserName__FOlD = &PB_UserResponse_CheckUserName{
}


var UserView__FOlD = &UserView{
}


var PB_Chat__FOlD = &PB_Chat{
        ChatId:  0 ,
        ChatKey:  "" ,
        RoomTypeEnumId:  0 ,
        UserId:  0 ,
        LastSeqSeen:  0 ,
        LastSeqDelete:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedTime:  0 ,
        CurrentSeq:  0 ,
        UpdatedMs:  0 ,
}


var PB_DirectMessage__FOlD = &PB_DirectMessage{
        MessageId:  0 ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnum:  0 ,
        Text:  "" ,
        Time:  0 ,
        PeerReceivedTime:  0 ,
        PeerSeenTime:  0 ,
        DeliviryStatusEnum:  0 ,
}


var PB_MessageFile__FOlD = &PB_MessageFile{
        MessageFileId:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileTypeEnumId:  0 ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        HashMd5:  "" ,
        HashAccess:  0 ,
        CreatedSe:  0 ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerThumbPath:  "" ,
        BucketId:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
}


var PB_User__FOlD = &PB_User{
        Id:  0 ,
        UserName:  "" ,
        UserNameLower:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        About:  "" ,
        FullName:  "" ,
        AvatarUrl:  "" ,
        PrivacyProfile:  0 ,
        Phone:  "" ,
        Email:  "" ,
        IsDeleted:  0 ,
        PasswordHash:  "" ,
        PasswordSalt:  "" ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        MediaCount:  0 ,
        LikesCount:  0 ,
        ResharedCount:  0 ,
        LastActionTime:  0 ,
        LastPostTime:  0 ,
        PrimaryFollowingList:  0 ,
        CreatedTime:  0 ,
        UpdatedTime:  0 ,
        SessionUuid:  "" ,
        DeviceUuid:  "" ,
        LastWifiMacAddress:  "" ,
        LastNetworkType:  "" ,
        AppVersion:  0 ,
        LastActivityTime:  0 ,
        LastLoginTime:  0 ,
        LastIpAddress:  "" ,
}


var PB_UpdateGroupParticipants__FOlD = &PB_UpdateGroupParticipants{
}


var PB_UpdateNotifySettings__FOlD = &PB_UpdateNotifySettings{
}


var PB_UpdateServiceNotification__FOlD = &PB_UpdateServiceNotification{
}


var PB_UpdateMessageMeta__FOlD = &PB_UpdateMessageMeta{
        MessageId:  0 ,
        AtTime:  0 ,
}


var PB_UpdateMessageId__FOlD = &PB_UpdateMessageId{
        OldMessageId:  0 ,
        NewMessageId:  0 ,
}


var PB_UpdateMessageToEdit__FOlD = &PB_UpdateMessageToEdit{
        MessageId:  0 ,
        NewText:  "" ,
}


var PB_UpdateMessageToDelete__FOlD = &PB_UpdateMessageToDelete{
        MessageId:  0 ,
}


var PB_UpdateRoomActionDoing__FOlD = &PB_UpdateRoomActionDoing{
        RoomKey:  "" ,

}


var PB_UpdateUserBlocked__FOlD = &PB_UpdateUserBlocked{
        UserId:  0 ,
        Blocked:  false ,
}


var PB_PushHolderView__FOlD = &PB_PushHolderView{














}


var PB_DirectUpdatesView__FOlD = &PB_DirectUpdatesView{












}


var PB_GeneralUpdatesView__FOlD = &PB_GeneralUpdatesView{


}


var PB_NotifyUpdatesView__FOlD = &PB_NotifyUpdatesView{


}


var PB_ChatView__FOlD = &PB_ChatView{
        ChatKey:  "" ,
        ChatId:  0 ,
        RoomTypeEnumId:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedTime:  0 ,
        UpdatedMs:  0 ,
        LastMessageId:  0 ,
        LastDeletedMessageId:  0 ,
        LastSeenMessageId:  0 ,
        LastSeqSeen:  0 ,
        LastSeqDelete:  0 ,
        CurrentSeq:  0 ,

        SharedMediaCount:  0 ,
        UnseenCount:  0 ,


}


var PB_MessageView__FOlD = &PB_MessageView{
        MessageId:  0 ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnumId:  0 ,
        Text:  "" ,
        Time:  0 ,
        PeerReceivedTime:  0 ,
        PeerSeenTime:  0 ,
        DeliviryStatusEnumId:  0 ,
        ChatId:  0 ,
        RoomTypeEnumId:  0 ,
        IsByMe:  false ,
        RemoteId:  0 ,

}


var PB_MessageFileView__FOlD = &PB_MessageFileView{
        MessageFileId:  0 ,
        OriginalUserId:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileTypeEnumId:  0 ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        HashMd5:  "" ,
        HashAccess:  0 ,
        CreatedSe:  0 ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerThumbPath:  "" ,
        BucketId:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
        ServerThumbLocalSrc:  "" ,
        RemoteMessageFileId:  0 ,
        LocalSrc:  "" ,
        ThumbLocalSrc:  "" ,
        MessageFileStatusId:  "" ,
}


var PB_UserView__FOlD = &PB_UserView{
        UserId:  0 ,
        UserName:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        About:  "" ,
        FullName:  "" ,
        AvatarUrl:  "" ,
        PrivacyProfile:  0 ,
        IsDeleted:  0 ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        UpdatedTime:  0 ,
        AppVersion:  0 ,
        LastActivityTime:  0 ,
        FollowingType:  0 ,
}



*/
