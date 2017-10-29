package xconst

const (
	Activity_Table                     = "activity"
	Activity_TableGo                   = "Activity"
	Bucket_Table                       = "bucket"
	Bucket_TableGo                     = "Bucket"
	Chat_Table                         = "chat"
	Chat_TableGo                       = "Chat"
	Comment_Table                      = "comment"
	Comment_TableGo                    = "Comment"
	DirectMessage_Table                = "direct_message"
	DirectMessage_TableGo              = "DirectMessage"
	DirectOffline_Table                = "direct_offline"
	DirectOffline_TableGo              = "DirectOffline"
	DirectOfflineDep_Table             = "direct_offline_dep"
	DirectOfflineDep_TableGo           = "DirectOfflineDep"
	DirectToMessage_Table              = "direct_to_message"
	DirectToMessage_TableGo            = "DirectToMessage"
	DirectUpdate_Table                 = "direct_update"
	DirectUpdate_TableGo               = "DirectUpdate"
	FollowingList_Table                = "following_list"
	FollowingList_TableGo              = "FollowingList"
	FollowingListMember_Table          = "following_list_member"
	FollowingListMember_TableGo        = "FollowingListMember"
	FollowingListMemberHistory_Table   = "following_list_member_history"
	FollowingListMemberHistory_TableGo = "FollowingListMemberHistory"
	GeneralLog_Table                   = "general_log"
	GeneralLog_TableGo                 = "GeneralLog"
	Group_Table                        = "group"
	Group_TableGo                      = "Group"
	GroupMember_Table                  = "group_member"
	GroupMember_TableGo                = "GroupMember"
	GroupMessage_Table                 = "group_message"
	GroupMessage_TableGo               = "GroupMessage"
	GroupToMessage_Table               = "group_to_message"
	GroupToMessage_TableGo             = "GroupToMessage"
	Like_Table                         = "likes"
	Like_TableGo                       = "Like"
	LogChange_Table                    = "log_changes"
	LogChange_TableGo                  = "LogChange"
	Media_Table                        = "media"
	Media_TableGo                      = "Media"
	MessageFile_Table                  = "message_file"
	MessageFile_TableGo                = "MessageFile"
	Msg_Table                          = "msg"
	Msg_TableGo                        = "Msg"
	Notification_Table                 = "notification"
	Notification_TableGo               = "Notification"
	NotificationRemoved_Table          = "notification_removed"
	NotificationRemoved_TableGo        = "NotificationRemoved"
	Offline_Table                      = "offline"
	Offline_TableGo                    = "Offline"
	OldMessage_Table                   = "old_messages"
	OldMessage_TableGo                 = "OldMessage"
	OldMsgFile_Table                   = "old_msg_file"
	OldMsgFile_TableGo                 = "OldMsgFile"
	OldMsgPush_Table                   = "old_msg_push"
	OldMsgPush_TableGo                 = "OldMsgPush"
	OldMsgPushEvent_Table              = "old_msg_push_event"
	OldMsgPushEvent_TableGo            = "OldMsgPushEvent"
	PhoneContact_Table                 = "phone_contacts"
	PhoneContact_TableGo               = "PhoneContact"
	Photo_Table                        = "photo"
	Photo_TableGo                      = "Photo"
	Post_Table                         = "post"
	Post_TableGo                       = "Post"
	PushEvent_Table                    = "push_event"
	PushEvent_TableGo                  = "PushEvent"
	PushMessage_Table                  = "push_message"
	PushMessage_TableGo                = "PushMessage"
	RecommendUser_Table                = "recommend_user"
	RecommendUser_TableGo              = "RecommendUser"
	Room_Table                         = "room"
	Room_TableGo                       = "Room"
	SearchClicked_Table                = "search_clicked"
	SearchClicked_TableGo              = "SearchClicked"
	Session_Table                      = "session"
	Session_TableGo                    = "Session"
	SettingClient_Table                = "setting_client"
	SettingClient_TableGo              = "SettingClient"
	SettingNotification_Table          = "setting_notifications"
	SettingNotification_TableGo        = "SettingNotification"
	Tag_Table                          = "tag"
	Tag_TableGo                        = "Tag"
	TagsPost_Table                     = "tags_posts"
	TagsPost_TableGo                   = "TagsPost"
	TestChat_Table                     = "test_chat"
	TestChat_TableGo                   = "TestChat"
	TriggerLog_Table                   = "trigger_log"
	TriggerLog_TableGo                 = "TriggerLog"
	User_Table                         = "user"
	User_TableGo                       = "User"
	UserMetaInfo_Table                 = "user_meta_info"
	UserMetaInfo_TableGo               = "UserMetaInfo"
	UserPassword_Table                 = "user_password"
	UserPassword_TableGo               = "UserPassword"
)

var Activity = struct {
	Id           string
	ActorUserId  string
	ActionTypeId string
	RowId        string
	RootId       string
	RefId        string
	CreatedAt    string
}{

	Id:           "Id",
	ActorUserId:  "ActorUserId",
	ActionTypeId: "ActionTypeId",
	RowId:        "RowId",
	RootId:       "RootId",
	RefId:        "RefId",
	CreatedAt:    "CreatedAt",
}

var Bucket = struct {
	BucketId            string
	BucketName          string
	Server1Id           string
	Server2Id           string
	BackupServerId      string
	ContentObjectTypeId string
	ContentObjectCount  string
	CreatedTime         string
}{

	BucketId:            "BucketId",
	BucketName:          "BucketName",
	Server1Id:           "Server1Id",
	Server2Id:           "Server2Id",
	BackupServerId:      "BackupServerId",
	ContentObjectTypeId: "ContentObjectTypeId",
	ContentObjectCount:  "ContentObjectCount",
	CreatedTime:         "CreatedTime",
}

var Chat = struct {
	ChatKey              string
	RoomKey              string
	RoomTypeEnumId       string
	UserId               string
	PeerUserId           string
	GroupId              string
	CreatedSe            string
	StartMessageIdFrom   string
	LastDeletedMessageId string
	LastSeenMessageId    string
	LastMessageId        string
	UpdatedMs            string
}{

	ChatKey:              "ChatKey",
	RoomKey:              "RoomKey",
	RoomTypeEnumId:       "RoomTypeEnumId",
	UserId:               "UserId",
	PeerUserId:           "PeerUserId",
	GroupId:              "GroupId",
	CreatedSe:            "CreatedSe",
	StartMessageIdFrom:   "StartMessageIdFrom",
	LastDeletedMessageId: "LastDeletedMessageId",
	LastSeenMessageId:    "LastSeenMessageId",
	LastMessageId:        "LastMessageId",
	UpdatedMs:            "UpdatedMs",
}

var Comment = struct {
	Id          string
	UserId      string
	PostId      string
	Text        string
	CreatedTime string
}{

	Id:          "Id",
	UserId:      "UserId",
	PostId:      "PostId",
	Text:        "Text",
	CreatedTime: "CreatedTime",
}

var DirectMessage = struct {
	MessageId            string
	MessageKey           string
	RoomKey              string
	UserId               string
	MessageFileId        string
	MessageTypeEnumId    string
	Text                 string
	CreatedSe            string
	PeerReceivedTime     string
	PeerSeenTime         string
	DeliviryStatusEnumId string
}{

	MessageId:            "MessageId",
	MessageKey:           "MessageKey",
	RoomKey:              "RoomKey",
	UserId:               "UserId",
	MessageFileId:        "MessageFileId",
	MessageTypeEnumId:    "MessageTypeEnumId",
	Text:                 "Text",
	CreatedSe:            "CreatedSe",
	PeerReceivedTime:     "PeerReceivedTime",
	PeerSeenTime:         "PeerSeenTime",
	DeliviryStatusEnumId: "DeliviryStatusEnumId",
}

var DirectOffline = struct {
	DirectOfflineId string
	ToUserId        string
	ChatKey         string
	MessageId       string
	MessageFileId   string
	PBClass         string
	DataPB          string
	DataJson        string
	DataTemp        string
	AtTimeMs        string
}{

	DirectOfflineId: "DirectOfflineId",
	ToUserId:        "ToUserId",
	ChatKey:         "ChatKey",
	MessageId:       "MessageId",
	MessageFileId:   "MessageFileId",
	PBClass:         "PBClass",
	DataPB:          "DataPB",
	DataJson:        "DataJson",
	DataTemp:        "DataTemp",
	AtTimeMs:        "AtTimeMs",
}

var DirectOfflineDep = struct {
	DirectOfflineId string
	ToUserId        string
	MessageId       string
	MessageFileId   string
	OtherId         string
	ChatKey         string
	PeerUserId      string
	RoomLogTypeId   string
	PBClass         string
	DataPB          string
	DataJson        string
	DataTemp        string
	AtTimeMs        string
}{

	DirectOfflineId: "DirectOfflineId",
	ToUserId:        "ToUserId",
	MessageId:       "MessageId",
	MessageFileId:   "MessageFileId",
	OtherId:         "OtherId",
	ChatKey:         "ChatKey",
	PeerUserId:      "PeerUserId",
	RoomLogTypeId:   "RoomLogTypeId",
	PBClass:         "PBClass",
	DataPB:          "DataPB",
	DataJson:        "DataJson",
	DataTemp:        "DataTemp",
	AtTimeMs:        "AtTimeMs",
}

var DirectToMessage = struct {
	Id           string
	ChatKey      string
	MessageId    string
	SourceEnumId string
}{

	Id:           "Id",
	ChatKey:      "ChatKey",
	MessageId:    "MessageId",
	SourceEnumId: "SourceEnumId",
}

var DirectUpdate = struct {
	DirectUpdateId string
	ToUserId       string
	MessageId      string
	MessageFileId  string
	OtherId        string
	ChatKey        string
	PeerUserId     string
	EventType      string
	RoomLogTypeId  string
	FromSeq        string
	ToSeq          string
	ExtraPB        string
	ExtraJson      string
	AtTimeMs       string
}{

	DirectUpdateId: "DirectUpdateId",
	ToUserId:       "ToUserId",
	MessageId:      "MessageId",
	MessageFileId:  "MessageFileId",
	OtherId:        "OtherId",
	ChatKey:        "ChatKey",
	PeerUserId:     "PeerUserId",
	EventType:      "EventType",
	RoomLogTypeId:  "RoomLogTypeId",
	FromSeq:        "FromSeq",
	ToSeq:          "ToSeq",
	ExtraPB:        "ExtraPB",
	ExtraJson:      "ExtraJson",
	AtTimeMs:       "AtTimeMs",
}

var FollowingList = struct {
	Id          string
	UserId      string
	ListType    string
	Name        string
	Count       string
	IsAuto      string
	IsPimiry    string
	CreatedTime string
}{

	Id:          "Id",
	UserId:      "UserId",
	ListType:    "ListType",
	Name:        "Name",
	Count:       "Count",
	IsAuto:      "IsAuto",
	IsPimiry:    "IsPimiry",
	CreatedTime: "CreatedTime",
}

var FollowingListMember = struct {
	Id             string
	ListId         string
	UserId         string
	FollowedUserId string
	FollowType     string
	UpdatedTimeMs  string
}{

	Id:             "Id",
	ListId:         "ListId",
	UserId:         "UserId",
	FollowedUserId: "FollowedUserId",
	FollowType:     "FollowType",
	UpdatedTimeMs:  "UpdatedTimeMs",
}

var FollowingListMemberHistory = struct {
	Id             string
	ListId         string
	UserId         string
	FollowedUserId string
	FollowType     string
	UpdatedTimeMs  string
	FollowId       string
}{

	Id:             "Id",
	ListId:         "ListId",
	UserId:         "UserId",
	FollowedUserId: "FollowedUserId",
	FollowType:     "FollowType",
	UpdatedTimeMs:  "UpdatedTimeMs",
	FollowId:       "FollowId",
}

var GeneralLog = struct {
	Id        string
	ToUserId  string
	TargetId  string
	LogTypeId string
	ExtraPB   string
	ExtraJson string
	CreatedMs string
}{

	Id:        "Id",
	ToUserId:  "ToUserId",
	TargetId:  "TargetId",
	LogTypeId: "LogTypeId",
	ExtraPB:   "ExtraPB",
	ExtraJson: "ExtraJson",
	CreatedMs: "CreatedMs",
}

var Group = struct {
	GroupId          string
	GroupName        string
	MembersCount     string
	GroupPrivacyEnum string
	CreatorUserId    string
	CreatedTime      string
	UpdatedMs        string
	CurrentSeq       string
}{

	GroupId:          "GroupId",
	GroupName:        "GroupName",
	MembersCount:     "MembersCount",
	GroupPrivacyEnum: "GroupPrivacyEnum",
	CreatorUserId:    "CreatorUserId",
	CreatedTime:      "CreatedTime",
	UpdatedMs:        "UpdatedMs",
	CurrentSeq:       "CurrentSeq",
}

var GroupMember = struct {
	Id              string
	GroupId         string
	GroupKey        string
	UserId          string
	ByUserId        string
	GroupRoleEnumId string
	CreatedTime     string
}{

	Id:              "Id",
	GroupId:         "GroupId",
	GroupKey:        "GroupKey",
	UserId:          "UserId",
	ByUserId:        "ByUserId",
	GroupRoleEnumId: "GroupRoleEnumId",
	CreatedTime:     "CreatedTime",
}

var GroupMessage = struct {
	MessageId          string
	RoomKey            string
	UserId             string
	MessageFileId      string
	MessageTypeEnum    string
	Text               string
	CreatedMs          string
	DeliveryStatusEnum string
}{

	MessageId:          "MessageId",
	RoomKey:            "RoomKey",
	UserId:             "UserId",
	MessageFileId:      "MessageFileId",
	MessageTypeEnum:    "MessageTypeEnum",
	Text:               "Text",
	CreatedMs:          "CreatedMs",
	DeliveryStatusEnum: "DeliveryStatusEnum",
}

var GroupToMessage = struct {
	Id         string
	GroupId    string
	MessageId  string
	CreatedMs  string
	StatusEnum string
}{

	Id:         "Id",
	GroupId:    "GroupId",
	MessageId:  "MessageId",
	CreatedMs:  "CreatedMs",
	StatusEnum: "StatusEnum",
}

var Like = struct {
	Id          string
	PostId      string
	PostTypeId  string
	UserId      string
	TypeId      string
	CreatedTime string
}{

	Id:          "Id",
	PostId:      "PostId",
	PostTypeId:  "PostTypeId",
	UserId:      "UserId",
	TypeId:      "TypeId",
	CreatedTime: "CreatedTime",
}

var LogChange = struct {
	Id string
	T  string
}{

	Id: "Id",
	T:  "T",
}

var Media = struct {
	Id          string
	UserId      string
	PostId      string
	AlbumId     string
	TypeId      string
	CreatedTime string
	Src         string
}{

	Id:          "Id",
	UserId:      "UserId",
	PostId:      "PostId",
	AlbumId:     "AlbumId",
	TypeId:      "TypeId",
	CreatedTime: "CreatedTime",
	Src:         "Src",
}

var MessageFile = struct {
	MessageFileId   string
	MessageFileKey  string
	OriginalUserId  string
	Name            string
	Size            string
	FileTypeEnumId  string
	Width           string
	Height          string
	Duration        string
	Extension       string
	HashMd5         string
	HashAccess      string
	CreatedSe       string
	ServerSrc       string
	ServerPath      string
	ServerThumbPath string
	BucketId        string
	ServerId        string
	CanDel          string
}{

	MessageFileId:   "MessageFileId",
	MessageFileKey:  "MessageFileKey",
	OriginalUserId:  "OriginalUserId",
	Name:            "Name",
	Size:            "Size",
	FileTypeEnumId:  "FileTypeEnumId",
	Width:           "Width",
	Height:          "Height",
	Duration:        "Duration",
	Extension:       "Extension",
	HashMd5:         "HashMd5",
	HashAccess:      "HashAccess",
	CreatedSe:       "CreatedSe",
	ServerSrc:       "ServerSrc",
	ServerPath:      "ServerPath",
	ServerThumbPath: "ServerThumbPath",
	BucketId:        "BucketId",
	ServerId:        "ServerId",
	CanDel:          "CanDel",
}

var Msg = struct {
	Key  string
	Name string
	Id   string
}{

	Key:  "Key",
	Name: "Name",
	Id:   "Id",
}

var Notification = struct {
	Id           string
	ForUserId    string
	ActorUserId  string
	ActionTypeId string
	ObjectTypeId string
	RowId        string
	RootId       string
	RefId        string
	SeenStatus   string
	CreatedTime  string
}{

	Id:           "Id",
	ForUserId:    "ForUserId",
	ActorUserId:  "ActorUserId",
	ActionTypeId: "ActionTypeId",
	ObjectTypeId: "ObjectTypeId",
	RowId:        "RowId",
	RootId:       "RootId",
	RefId:        "RefId",
	SeenStatus:   "SeenStatus",
	CreatedTime:  "CreatedTime",
}

var NotificationRemoved = struct {
	NotificationId string
	ForUserId      string
}{

	NotificationId: "NotificationId",
	ForUserId:      "ForUserId",
}

var Offline = struct {
	Id          string
	FromUserId  string
	ToUserId    string
	RpcName     string
	PBClassName string
	Key         string
	DataJson    string
	DataBlob    string
	DataLength  string
	CreatedMs   string
}{

	Id:          "Id",
	FromUserId:  "FromUserId",
	ToUserId:    "ToUserId",
	RpcName:     "RpcName",
	PBClassName: "PBClassName",
	Key:         "Key",
	DataJson:    "DataJson",
	DataBlob:    "DataBlob",
	DataLength:  "DataLength",
	CreatedMs:   "CreatedMs",
}

var OldMessage = struct {
	Id            string
	Uid           string
	UserId        string
	MessageKey    string
	RoomKey       string
	MessageType   string
	RoomType      string
	MsgFileId     string
	DataPB        string
	Data64        string
	DataJson      string
	CreatedTimeMs string
}{

	Id:            "Id",
	Uid:           "Uid",
	UserId:        "UserId",
	MessageKey:    "MessageKey",
	RoomKey:       "RoomKey",
	MessageType:   "MessageType",
	RoomType:      "RoomType",
	MsgFileId:     "MsgFileId",
	DataPB:        "DataPB",
	Data64:        "Data64",
	DataJson:      "DataJson",
	CreatedTimeMs: "CreatedTimeMs",
}

var OldMsgFile = struct {
	Id          string
	Name        string
	Size        string
	FileType    string
	MimeType    string
	Width       string
	Height      string
	Duration    string
	Extension   string
	ThumbData   string
	ThumbData64 string
	ServerSrc   string
	ServerPath  string
	ServerId    string
	CanDel      string
}{

	Id:          "Id",
	Name:        "Name",
	Size:        "Size",
	FileType:    "FileType",
	MimeType:    "MimeType",
	Width:       "Width",
	Height:      "Height",
	Duration:    "Duration",
	Extension:   "Extension",
	ThumbData:   "ThumbData",
	ThumbData64: "ThumbData64",
	ServerSrc:   "ServerSrc",
	ServerPath:  "ServerPath",
	ServerId:    "ServerId",
	CanDel:      "CanDel",
}

var OldMsgPush = struct {
	Id            string
	Uid           string
	ToUser        string
	MsgUid        string
	CreatedTimeMs string
}{

	Id:            "Id",
	Uid:           "Uid",
	ToUser:        "ToUser",
	MsgUid:        "MsgUid",
	CreatedTimeMs: "CreatedTimeMs",
}

var OldMsgPushEvent = struct {
	Id         string
	Uid        string
	ToUserId   string
	MsgUid     string
	MsgKey     string
	RoomKey    string
	PeerUserId string
	EventType  string
	AtTime     string
}{

	Id:         "Id",
	Uid:        "Uid",
	ToUserId:   "ToUserId",
	MsgUid:     "MsgUid",
	MsgKey:     "MsgKey",
	RoomKey:    "RoomKey",
	PeerUserId: "PeerUserId",
	EventType:  "EventType",
	AtTime:     "AtTime",
}

var PhoneContact = struct {
	Id                    string
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     string
	UserId                string
	DeviceUuidId          string
	CreatedTime           string
	UpdatedTime           string
}{

	Id:                    "Id",
	PhoneDisplayName:      "PhoneDisplayName",
	PhoneFamilyName:       "PhoneFamilyName",
	PhoneNumber:           "PhoneNumber",
	PhoneNormalizedNumber: "PhoneNormalizedNumber",
	PhoneContactRowId:     "PhoneContactRowId",
	UserId:                "UserId",
	DeviceUuidId:          "DeviceUuidId",
	CreatedTime:           "CreatedTime",
	UpdatedTime:           "UpdatedTime",
}

var Photo = struct {
	PhotoId     string
	UserId      string
	PostId      string
	AlbumId     string
	ImageTypeId string
	Title       string
	Src         string
	PathSrc     string
	BucketId    string
	Width       string
	Height      string
	Ratio       string
	HashMd5     string
	Color       string
	CreatedTime string
	W1080       string
	W720        string
	W480        string
	W320        string
	W160        string
	W80         string
}{

	PhotoId:     "PhotoId",
	UserId:      "UserId",
	PostId:      "PostId",
	AlbumId:     "AlbumId",
	ImageTypeId: "ImageTypeId",
	Title:       "Title",
	Src:         "Src",
	PathSrc:     "PathSrc",
	BucketId:    "BucketId",
	Width:       "Width",
	Height:      "Height",
	Ratio:       "Ratio",
	HashMd5:     "HashMd5",
	Color:       "Color",
	CreatedTime: "CreatedTime",
	W1080:       "W1080",
	W720:        "W720",
	W480:        "W480",
	W320:        "W320",
	W160:        "W160",
	W80:         "W80",
}

var Post = struct {
	Id             string
	UserId         string
	TypeId         string
	Text           string
	FormatedText   string
	MediaCount     string
	SharedTo       string
	DisableComment string
	HasTag         string
	LikesCount     string
	CommentsCount  string
	EditedTime     string
	CreatedTime    string
}{

	Id:             "Id",
	UserId:         "UserId",
	TypeId:         "TypeId",
	Text:           "Text",
	FormatedText:   "FormatedText",
	MediaCount:     "MediaCount",
	SharedTo:       "SharedTo",
	DisableComment: "DisableComment",
	HasTag:         "HasTag",
	LikesCount:     "LikesCount",
	CommentsCount:  "CommentsCount",
	EditedTime:     "EditedTime",
	CreatedTime:    "CreatedTime",
}

var PushEvent = struct {
	PushEventId       string
	ToUserId          string
	ToDeviceId        string
	MessageId         string
	RoomTypeEnum      string
	RoomId            string
	PeerUserId        string
	PushEventTypeEnum string
	AtTime            string
}{

	PushEventId:       "PushEventId",
	ToUserId:          "ToUserId",
	ToDeviceId:        "ToDeviceId",
	MessageId:         "MessageId",
	RoomTypeEnum:      "RoomTypeEnum",
	RoomId:            "RoomId",
	PeerUserId:        "PeerUserId",
	PushEventTypeEnum: "PushEventTypeEnum",
	AtTime:            "AtTime",
}

var PushMessage = struct {
	PushMessageId string
	ToUserId      string
	ToDeviceId    string
	MessageId     string
	RoomTypeEnum  string
	CreatedMs     string
}{

	PushMessageId: "PushMessageId",
	ToUserId:      "ToUserId",
	ToDeviceId:    "ToDeviceId",
	MessageId:     "MessageId",
	RoomTypeEnum:  "RoomTypeEnum",
	CreatedMs:     "CreatedMs",
}

var RecommendUser = struct {
	Id          string
	UserId      string
	TargetId    string
	Weight      string
	CreatedTime string
}{

	Id:          "Id",
	UserId:      "UserId",
	TargetId:    "TargetId",
	Weight:      "Weight",
	CreatedTime: "CreatedTime",
}

var Room = struct {
	RoomId        string
	RoomKey       string
	RoomTypeEnum  string
	UserId        string
	LastSeqSeen   string
	LastSeqDelete string
	PeerUserId    string
	GroupId       string
	CreatedTime   string
	CurrentSeq    string
}{

	RoomId:        "RoomId",
	RoomKey:       "RoomKey",
	RoomTypeEnum:  "RoomTypeEnum",
	UserId:        "UserId",
	LastSeqSeen:   "LastSeqSeen",
	LastSeqDelete: "LastSeqDelete",
	PeerUserId:    "PeerUserId",
	GroupId:       "GroupId",
	CreatedTime:   "CreatedTime",
	CurrentSeq:    "CurrentSeq",
}

var SearchClicked = struct {
	Id        string
	Query     string
	ClickType string
	TargetId  string
	UserId    string
	CreatedAt string
}{

	Id:        "Id",
	Query:     "Query",
	ClickType: "ClickType",
	TargetId:  "TargetId",
	UserId:    "UserId",
	CreatedAt: "CreatedAt",
}

var Session = struct {
	Id                      string
	UserId                  string
	SessionUuid             string
	ClientUuid              string
	DeviceUuid              string
	LastActivityTime        string
	LastIpAddress           string
	LastWifiMacAddress      string
	LastNetworkType         string
	LastNetworkTypeId       string
	AppVersion              string
	UpdatedTime             string
	CreatedTime             string
	LastSyncDirectUpdateId  string
	LastSyncGeneralUpdateId string
	LastSyncNotifyUpdateId  string
}{

	Id:                      "Id",
	UserId:                  "UserId",
	SessionUuid:             "SessionUuid",
	ClientUuid:              "ClientUuid",
	DeviceUuid:              "DeviceUuid",
	LastActivityTime:        "LastActivityTime",
	LastIpAddress:           "LastIpAddress",
	LastWifiMacAddress:      "LastWifiMacAddress",
	LastNetworkType:         "LastNetworkType",
	LastNetworkTypeId:       "LastNetworkTypeId",
	AppVersion:              "AppVersion",
	UpdatedTime:             "UpdatedTime",
	CreatedTime:             "CreatedTime",
	LastSyncDirectUpdateId:  "LastSyncDirectUpdateId",
	LastSyncGeneralUpdateId: "LastSyncGeneralUpdateId",
	LastSyncNotifyUpdateId:  "LastSyncNotifyUpdateId",
}

var SettingClient = struct {
	UserId                    string
	AutoDownloadWifiVoice     string
	AutoDownloadWifiImage     string
	AutoDownloadWifiVideo     string
	AutoDownloadWifiFile      string
	AutoDownloadWifiMusic     string
	AutoDownloadWifiGif       string
	AutoDownloadCellularVoice string
	AutoDownloadCellularImage string
	AutoDownloadCellularVideo string
	AutoDownloadCellularFile  string
	AutoDownloadCellularMusic string
	AutoDownloadCellularGif   string
	AutoDownloadRoamingVoice  string
	AutoDownloadRoamingImage  string
	AutoDownloadRoamingVideo  string
	AutoDownloadRoamingFile   string
	AutoDownloadRoamingMusic  string
	AutoDownloadRoamingGif    string
	SaveToGallery             string
}{

	UserId:                    "UserId",
	AutoDownloadWifiVoice:     "AutoDownloadWifiVoice",
	AutoDownloadWifiImage:     "AutoDownloadWifiImage",
	AutoDownloadWifiVideo:     "AutoDownloadWifiVideo",
	AutoDownloadWifiFile:      "AutoDownloadWifiFile",
	AutoDownloadWifiMusic:     "AutoDownloadWifiMusic",
	AutoDownloadWifiGif:       "AutoDownloadWifiGif",
	AutoDownloadCellularVoice: "AutoDownloadCellularVoice",
	AutoDownloadCellularImage: "AutoDownloadCellularImage",
	AutoDownloadCellularVideo: "AutoDownloadCellularVideo",
	AutoDownloadCellularFile:  "AutoDownloadCellularFile",
	AutoDownloadCellularMusic: "AutoDownloadCellularMusic",
	AutoDownloadCellularGif:   "AutoDownloadCellularGif",
	AutoDownloadRoamingVoice:  "AutoDownloadRoamingVoice",
	AutoDownloadRoamingImage:  "AutoDownloadRoamingImage",
	AutoDownloadRoamingVideo:  "AutoDownloadRoamingVideo",
	AutoDownloadRoamingFile:   "AutoDownloadRoamingFile",
	AutoDownloadRoamingMusic:  "AutoDownloadRoamingMusic",
	AutoDownloadRoamingGif:    "AutoDownloadRoamingGif",
	SaveToGallery:             "SaveToGallery",
}

var SettingNotification = struct {
	UserId                   string
	SocialLedOn              string
	SocialLedColor           string
	ReqestToFollowYou        string
	FollowedYou              string
	AccptedYourFollowRequest string
	YourPostLiked            string
	YourPostCommented        string
	MenthenedYouInPost       string
	MenthenedYouInComment    string
	YourContactsJoined       string
	DirectMessage            string
	DirectAlert              string
	DirectPerview            string
	DirectLedOn              string
	DirectLedColor           string
	DirectVibrate            string
	DirectPopup              string
	DirectSound              string
	DirectPriority           string
}{

	UserId:                   "UserId",
	SocialLedOn:              "SocialLedOn",
	SocialLedColor:           "SocialLedColor",
	ReqestToFollowYou:        "ReqestToFollowYou",
	FollowedYou:              "FollowedYou",
	AccptedYourFollowRequest: "AccptedYourFollowRequest",
	YourPostLiked:            "YourPostLiked",
	YourPostCommented:        "YourPostCommented",
	MenthenedYouInPost:       "MenthenedYouInPost",
	MenthenedYouInComment:    "MenthenedYouInComment",
	YourContactsJoined:       "YourContactsJoined",
	DirectMessage:            "DirectMessage",
	DirectAlert:              "DirectAlert",
	DirectPerview:            "DirectPerview",
	DirectLedOn:              "DirectLedOn",
	DirectLedColor:           "DirectLedColor",
	DirectVibrate:            "DirectVibrate",
	DirectPopup:              "DirectPopup",
	DirectSound:              "DirectSound",
	DirectPriority:           "DirectPriority",
}

var Tag = struct {
	Id          string
	Name        string
	Count       string
	IsBlocked   string
	CreatedTime string
}{

	Id:          "Id",
	Name:        "Name",
	Count:       "Count",
	IsBlocked:   "IsBlocked",
	CreatedTime: "CreatedTime",
}

var TagsPost = struct {
	Id          string
	TagId       string
	PostId      string
	TypeId      string
	CreatedTime string
}{

	Id:          "Id",
	TagId:       "TagId",
	PostId:      "PostId",
	TypeId:      "TypeId",
	CreatedTime: "CreatedTime",
}

var TestChat = struct {
	Id     string
	Id4    string
	TimeMs string
	Text   string
	Name   string
	UserId string
	C2     string
	C3     string
	C4     string
	C5     string
}{

	Id:     "Id",
	Id4:    "Id4",
	TimeMs: "TimeMs",
	Text:   "Text",
	Name:   "Name",
	UserId: "UserId",
	C2:     "C2",
	C3:     "C3",
	C4:     "C4",
	C5:     "C5",
}

var TriggerLog = struct {
	Id         string
	ModelName  string
	ChangeType string
	TargetId   string
	TargetStr  string
	CreatedSe  string
}{

	Id:         "Id",
	ModelName:  "ModelName",
	ChangeType: "ChangeType",
	TargetId:   "TargetId",
	TargetStr:  "TargetStr",
	CreatedSe:  "CreatedSe",
}

var User = struct {
	Id                   string
	UserName             string
	UserNameLower        string
	FirstName            string
	LastName             string
	About                string
	FullName             string
	AvatarUrl            string
	PrivacyProfile       string
	Phone                string
	Email                string
	IsDeleted            string
	PasswordHash         string
	PasswordSalt         string
	FollowersCount       string
	FollowingCount       string
	PostsCount           string
	MediaCount           string
	LikesCount           string
	ResharedCount        string
	LastActionTime       string
	LastPostTime         string
	PrimaryFollowingList string
	CreatedTime          string
	UpdatedTime          string
	SessionUuid          string
	DeviceUuid           string
	LastWifiMacAddress   string
	LastNetworkType      string
	AppVersion           string
	LastActivityTime     string
	LastLoginTime        string
	LastIpAddress        string
}{

	Id:                   "Id",
	UserName:             "UserName",
	UserNameLower:        "UserNameLower",
	FirstName:            "FirstName",
	LastName:             "LastName",
	About:                "About",
	FullName:             "FullName",
	AvatarUrl:            "AvatarUrl",
	PrivacyProfile:       "PrivacyProfile",
	Phone:                "Phone",
	Email:                "Email",
	IsDeleted:            "IsDeleted",
	PasswordHash:         "PasswordHash",
	PasswordSalt:         "PasswordSalt",
	FollowersCount:       "FollowersCount",
	FollowingCount:       "FollowingCount",
	PostsCount:           "PostsCount",
	MediaCount:           "MediaCount",
	LikesCount:           "LikesCount",
	ResharedCount:        "ResharedCount",
	LastActionTime:       "LastActionTime",
	LastPostTime:         "LastPostTime",
	PrimaryFollowingList: "PrimaryFollowingList",
	CreatedTime:          "CreatedTime",
	UpdatedTime:          "UpdatedTime",
	SessionUuid:          "SessionUuid",
	DeviceUuid:           "DeviceUuid",
	LastWifiMacAddress:   "LastWifiMacAddress",
	LastNetworkType:      "LastNetworkType",
	AppVersion:           "AppVersion",
	LastActivityTime:     "LastActivityTime",
	LastLoginTime:        "LastLoginTime",
	LastIpAddress:        "LastIpAddress",
}

var UserMetaInfo = struct {
	Id                  string
	UserId              string
	IsNotificationDirty string
	LastUserRecGen      string
}{

	Id:                  "Id",
	UserId:              "UserId",
	IsNotificationDirty: "IsNotificationDirty",
	LastUserRecGen:      "LastUserRecGen",
}

var UserPassword = struct {
	UserId      string
	Password    string
	CreatedTime string
}{

	UserId:      "UserId",
	Password:    "Password",
	CreatedTime: "CreatedTime",
}
