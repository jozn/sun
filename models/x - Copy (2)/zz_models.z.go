// Package x contains the types for schema 'ms'.
package x

// GENERATED BY XO. DO NOT EDIT.

// Activity 'ms.activity'.
type Activity struct {
	Id           int
	ActorUserId  int
	ActionTypeId int
	RowId        int
	RootId       int
	RefId        int
	CreatedAt    int

	_exists, _deleted bool
}

/*
:= &Activity {
	Id: 0,
	ActorUserId: 0,
	ActionTypeId: 0,
	RowId: 0,
	RootId: 0,
	RefId: 0,
	CreatedAt: 0,
*/

// Bucket 'ms.bucket'.
type Bucket struct {
	BucketId            int
	BucketName          string
	Server1Id           int
	Server2Id           int
	BackupServerId      int
	ContentObjectTypeId int
	ContentObjectCount  int
	CreatedTime         int

	_exists, _deleted bool
}

/*
:= &Bucket {
	BucketId: 0,
	BucketName: "",
	Server1Id: 0,
	Server2Id: 0,
	BackupServerId: 0,
	ContentObjectTypeId: 0,
	ContentObjectCount: 0,
	CreatedTime: 0,
*/

// Comment 'ms.comments'.
type Comment struct {
	Id          int
	UserId      int
	PostId      int
	Text        string
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &Comment {
	Id: 0,
	UserId: 0,
	PostId: 0,
	Text: "",
	CreatedTime: 0,
*/

// DirectMessage 'ms.direct_message'.
type DirectMessage struct {
	MessageId          int
	RoomKey            string
	UserId             int
	MessageFileId      int
	MessageTypeEnum    int
	Text               string
	CreatedMs          int
	PeerReceivedTime   int
	PeerSeenTime       int
	DeliviryStatusEnum int

	_exists, _deleted bool
}

/*
:= &DirectMessage {
	MessageId: 0,
	RoomKey: "",
	UserId: 0,
	MessageFileId: 0,
	MessageTypeEnum: 0,
	Text: "",
	CreatedMs: 0,
	PeerReceivedTime: 0,
	PeerSeenTime: 0,
	DeliviryStatusEnum: 0,
*/

// DirectToMessage 'ms.direct_to_message'.
type DirectToMessage struct {
	Id         int
	RoomId     int
	MessageId  int
	Seq        int
	SourceEnum int

	_exists, _deleted bool
}

/*
:= &DirectToMessage {
	Id: 0,
	RoomId: 0,
	MessageId: 0,
	Seq: 0,
	SourceEnum: 0,
*/

// FollowingList 'ms.following_list'.
type FollowingList struct {
	Id          int
	UserId      int
	ListType    int
	Name        string
	Count       int
	IsAuto      int
	IsPimiry    int
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &FollowingList {
	Id: 0,
	UserId: 0,
	ListType: 0,
	Name: "",
	Count: 0,
	IsAuto: 0,
	IsPimiry: 0,
	CreatedTime: 0,
*/

// FollowingListMember 'ms.following_list_member'.
type FollowingListMember struct {
	Id             int
	ListId         int
	UserId         int
	FollowedUserId int
	FollowType     int
	UpdatedTimeMs  int

	_exists, _deleted bool
}

/*
:= &FollowingListMember {
	Id: 0,
	ListId: 0,
	UserId: 0,
	FollowedUserId: 0,
	FollowType: 0,
	UpdatedTimeMs: 0,
*/

// FollowingListMemberHistory 'ms.following_list_member_history'.
type FollowingListMemberHistory struct {
	Id             int
	ListId         int
	UserId         int
	FollowedUserId int
	FollowType     int
	UpdatedTimeMs  int
	FollowId       int

	_exists, _deleted bool
}

/*
:= &FollowingListMemberHistory {
	Id: 0,
	ListId: 0,
	UserId: 0,
	FollowedUserId: 0,
	FollowType: 0,
	UpdatedTimeMs: 0,
	FollowId: 0,
*/

// Group 'ms.group'.
type Group struct {
	GroupId          int
	GroupName        string
	MembersCount     int
	GroupPrivacyEnum int
	CreatorUserId    int
	CreatedTime      int
	UpdatedMs        int
	CurrentSeq       int

	_exists, _deleted bool
}

/*
:= &Group {
	GroupId: 0,
	GroupName: "",
	MembersCount: 0,
	GroupPrivacyEnum: 0,
	CreatorUserId: 0,
	CreatedTime: 0,
	UpdatedMs: 0,
	CurrentSeq: 0,
*/

// GroupMember 'ms.group_member'.
type GroupMember struct {
	Id            int
	GroupId       int
	GroupKey      string
	UserId        int
	ByUserId      int
	GroupRoleEnum int
	CreatedTime   int

	_exists, _deleted bool
}

/*
:= &GroupMember {
	Id: 0,
	GroupId: 0,
	GroupKey: "",
	UserId: 0,
	ByUserId: 0,
	GroupRoleEnum: 0,
	CreatedTime: 0,
*/

// GroupMessage 'ms.group_message'.
type GroupMessage struct {
	MessageId          int
	RoomKey            string
	UserId             int
	MessageFileId      int
	MessageTypeEnum    int
	Text               string
	CreatedMs          int
	DeliveryStatusEnum int

	_exists, _deleted bool
}

/*
:= &GroupMessage {
	MessageId: 0,
	RoomKey: "",
	UserId: 0,
	MessageFileId: 0,
	MessageTypeEnum: 0,
	Text: "",
	CreatedMs: 0,
	DeliveryStatusEnum: 0,
*/

// GroupToMessage 'ms.group_to_message'.
type GroupToMessage struct {
	Id         int
	GroupId    int
	MessageId  int
	CreatedMs  int
	StatusEnum int

	_exists, _deleted bool
}

/*
:= &GroupToMessage {
	Id: 0,
	GroupId: 0,
	MessageId: 0,
	CreatedMs: 0,
	StatusEnum: 0,
*/

// Like 'ms.likes'.
type Like struct {
	Id          int
	PostId      int
	PostTypeId  int
	UserId      int
	TypeId      int
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &Like {
	Id: 0,
	PostId: 0,
	PostTypeId: 0,
	UserId: 0,
	TypeId: 0,
	CreatedTime: 0,
*/

// Media 'ms.media'.
type Media struct {
	Id          int
	UserId      int
	PostId      int
	AlbumId     int
	TypeId      int
	CreatedTime int
	Src         string //json of sizes

	_exists, _deleted bool
}

/*
:= &Media {
	Id: 0,
	UserId: 0,
	PostId: 0,
	AlbumId: 0,
	TypeId: 0,
	CreatedTime: 0,
	Src: "",
*/

// MessageFile 'ms.message_file'.
type MessageFile struct {
	MessageFileId   int
	Name            string
	Size            int
	FileTypeEnum    int
	MimeType        string
	Width           int
	Height          int
	Duration        int
	Extension       string
	ThumbData       []byte
	ThumbData64     string
	ServerSrc       string
	ServerPath      string
	ServerThumbPath string
	BucketId        string
	ServerId        int
	CanDel          int
	CreatedTime     int

	_exists, _deleted bool
}

/*
:= &MessageFile {
	MessageFileId: 0,
	Name: "",
	Size: 0,
	FileTypeEnum: 0,
	MimeType: "",
	Width: 0,
	Height: 0,
	Duration: 0,
	Extension: "",
	ThumbData: UNKNOWN,
	ThumbData64: "",
	ServerSrc: "",
	ServerPath: "",
	ServerThumbPath: "",
	BucketId: "",
	ServerId: 0,
	CanDel: 0,
	CreatedTime: 0,
*/

// Notification 'ms.notification'.
type Notification struct {
	Id           int
	ForUserId    int
	ActorUserId  int
	ActionTypeId int
	ObjectTypeId int
	RowId        int
	RootId       int
	RefId        int
	SeenStatus   int
	CreatedTime  int

	_exists, _deleted bool
}

/*
:= &Notification {
	Id: 0,
	ForUserId: 0,
	ActorUserId: 0,
	ActionTypeId: 0,
	ObjectTypeId: 0,
	RowId: 0,
	RootId: 0,
	RefId: 0,
	SeenStatus: 0,
	CreatedTime: 0,
*/

// NotificationRemoved 'ms.notification_removed'.
type NotificationRemoved struct {
	NotificationId int
	ForUserId      int

	_exists, _deleted bool
}

/*
:= &NotificationRemoved {
	NotificationId: 0,
	ForUserId: 0,
*/

// OldMessage 'ms.old_messages'.
type OldMessage struct {
	Id            int
	Uid           int
	UserId        int
	MessageKey    string
	RoomKey       string
	MessageType   int
	RoomType      int
	MsgFileId     int
	DataPB        []byte
	Data64        string
	DataJson      string
	CreatedTimeMs int

	_exists, _deleted bool
}

/*
:= &OldMessage {
	Id: 0,
	Uid: 0,
	UserId: 0,
	MessageKey: "",
	RoomKey: "",
	MessageType: 0,
	RoomType: 0,
	MsgFileId: 0,
	DataPB: UNKNOWN,
	Data64: "",
	DataJson: "",
	CreatedTimeMs: 0,
*/

// OldMsgFile 'ms.old_msg_file'.
type OldMsgFile struct {
	Id          int
	Name        string
	Size        int
	FileType    int
	MimeType    string
	Width       int
	Height      int
	Duration    int
	Extension   string
	ThumbData   []byte
	ThumbData64 string
	ServerSrc   string
	ServerPath  string
	ServerId    int
	CanDel      int

	_exists, _deleted bool
}

/*
:= &OldMsgFile {
	Id: 0,
	Name: "",
	Size: 0,
	FileType: 0,
	MimeType: "",
	Width: 0,
	Height: 0,
	Duration: 0,
	Extension: "",
	ThumbData: UNKNOWN,
	ThumbData64: "",
	ServerSrc: "",
	ServerPath: "",
	ServerId: 0,
	CanDel: 0,
*/

// OldMsgPush 'ms.old_msg_push'.
type OldMsgPush struct {
	Id            uint
	Uid           uint
	ToUser        int
	MsgUid        int
	CreatedTimeMs int

	_exists, _deleted bool
}

/*
:= &OldMsgPush {
	Id: UNKNOWN,
	Uid: UNKNOWN,
	ToUser: 0,
	MsgUid: 0,
	CreatedTimeMs: 0,
*/

// OldMsgPushEvent 'ms.old_msg_push_event'.
type OldMsgPushEvent struct {
	Id         int
	Uid        int
	ToUserId   int
	MsgUid     int
	MsgKey     string
	RoomKey    string
	PeerUserId int
	EventType  int
	AtTime     int

	_exists, _deleted bool
}

/*
:= &OldMsgPushEvent {
	Id: 0,
	Uid: 0,
	ToUserId: 0,
	MsgUid: 0,
	MsgKey: "",
	RoomKey: "",
	PeerUserId: 0,
	EventType: 0,
	AtTime: 0,
*/

// PhoneContact 'ms.phone_contacts'.
type PhoneContact struct {
	Id                    int
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     int
	UserId                int
	DeviceUuidId          int
	CreatedTime           int
	UpdatedTime           int

	_exists, _deleted bool
}

/*
:= &PhoneContact {
	Id: 0,
	PhoneDisplayName: "",
	PhoneFamilyName: "",
	PhoneNumber: "",
	PhoneNormalizedNumber: "",
	PhoneContactRowId: 0,
	UserId: 0,
	DeviceUuidId: 0,
	CreatedTime: 0,
	UpdatedTime: 0,
*/

// Photo 'ms.photo'.
type Photo struct {
	PhotoId     int
	UserId      int
	PostId      int
	AlbumId     int
	ImageTypeId int
	Title       string
	Src         string //json of sizes
	PathSrc     string
	BucketId    int
	Width       int
	Height      int
	Ratio       float32
	HashMd5     string
	Color       string
	CreatedTime int
	W1080       int
	W720        int
	W480        int
	W320        int
	W160        int
	W80         int `json:"-"` //nojson

	_exists, _deleted bool
}

/*
:= &Photo {
	PhotoId: 0,
	UserId: 0,
	PostId: 0,
	AlbumId: 0,
	ImageTypeId: 0,
	Title: "",
	Src: "",
	PathSrc: "",
	BucketId: 0,
	Width: 0,
	Height: 0,
	Ratio: float32(0),
	HashMd5: "",
	Color: "",
	CreatedTime: 0,
	W1080: 0,
	W720: 0,
	W480: 0,
	W320: 0,
	W160: 0,
	W80: 0,
*/

// Post 'ms.post'.
type Post struct {
	Id             int
	UserId         int
	TypeId         int
	Text           string
	FormatedText   string
	MediaCount     int
	SharedTo       int
	DisableComment int
	HasTag         int
	LikesCount     int
	CommentsCount  int
	EditedTime     int
	CreatedTime    int

	_exists, _deleted bool
}

/*
:= &Post {
	Id: 0,
	UserId: 0,
	TypeId: 0,
	Text: "",
	FormatedText: "",
	MediaCount: 0,
	SharedTo: 0,
	DisableComment: 0,
	HasTag: 0,
	LikesCount: 0,
	CommentsCount: 0,
	EditedTime: 0,
	CreatedTime: 0,
*/

// PushEvent 'ms.push_event'.
type PushEvent struct {
	PushEventId       int
	ToUserId          int
	ToDeviceId        int
	MessageId         int
	RoomTypeEnum      int
	RoomId            int
	PeerUserId        int
	PushEventTypeEnum int
	AtTime            int

	_exists, _deleted bool
}

/*
:= &PushEvent {
	PushEventId: 0,
	ToUserId: 0,
	ToDeviceId: 0,
	MessageId: 0,
	RoomTypeEnum: 0,
	RoomId: 0,
	PeerUserId: 0,
	PushEventTypeEnum: 0,
	AtTime: 0,
*/

// PushMessage 'ms.push_message'.
type PushMessage struct {
	PushMessageId uint
	ToUserId      int
	ToDeviceId    int
	MessageId     int
	RoomTypeEnum  int
	CreatedMs     int

	_exists, _deleted bool
}

/*
:= &PushMessage {
	PushMessageId: UNKNOWN,
	ToUserId: 0,
	ToDeviceId: 0,
	MessageId: 0,
	RoomTypeEnum: 0,
	CreatedMs: 0,
*/

// RecommendUser 'ms.recommend_user'.
type RecommendUser struct {
	Id          int
	UserId      int
	TargetId    int
	Weight      float32
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &RecommendUser {
	Id: 0,
	UserId: 0,
	TargetId: 0,
	Weight: float32(0),
	CreatedTime: 0,
*/

// Room 'ms.room'.
type Room struct {
	RoomId        int
	RoomKey       string
	RoomTypeEnum  int
	UserId        int
	LastSeqSeen   int
	LastSeqDelete int
	PeerUserId    int
	GroupId       int
	CreatedTime   int
	CurrentSeq    int //just for peer to peer

	_exists, _deleted bool
}

/*
:= &Room {
	RoomId: 0,
	RoomKey: "",
	RoomTypeEnum: 0,
	UserId: 0,
	LastSeqSeen: 0,
	LastSeqDelete: 0,
	PeerUserId: 0,
	GroupId: 0,
	CreatedTime: 0,
	CurrentSeq: 0,
*/

// SearchClicked 'ms.search_clicked'.
type SearchClicked struct {
	Id        int
	Query     string
	ClickType int
	TargetId  int
	UserId    int
	CreatedAt int

	_exists, _deleted bool
}

/*
:= &SearchClicked {
	Id: 0,
	Query: "",
	ClickType: 0,
	TargetId: 0,
	UserId: 0,
	CreatedAt: 0,
*/

// Session 'ms.session'.
type Session struct {
	Id                 int
	UserId             int
	SessionUuid        string
	ClientUuid         string
	DeviceUuid         string
	LastActivityTime   int
	LastIpAddress      string
	LastWifiMacAddress string
	LastNetworkType    string
	LastNetworkTypeId  int
	AppVersion         int
	UpdatedTime        int
	CreatedTime        int

	_exists, _deleted bool
}

/*
:= &Session {
	Id: 0,
	UserId: 0,
	SessionUuid: "",
	ClientUuid: "",
	DeviceUuid: "",
	LastActivityTime: 0,
	LastIpAddress: "",
	LastWifiMacAddress: "",
	LastNetworkType: "",
	LastNetworkTypeId: 0,
	AppVersion: 0,
	UpdatedTime: 0,
	CreatedTime: 0,
*/

// Tag 'ms.tags'.
type Tag struct {
	Id          int
	Name        string
	Count       int
	IsBlocked   int
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &Tag {
	Id: 0,
	Name: "",
	Count: 0,
	IsBlocked: 0,
	CreatedTime: 0,
*/

// TagsPost 'ms.tags_posts'.
type TagsPost struct {
	Id          int
	TagId       int
	PostId      int
	TypeId      int
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &TagsPost {
	Id: 0,
	TagId: 0,
	PostId: 0,
	TypeId: 0,
	CreatedTime: 0,
*/

// UserMetaInfo 'ms.user_meta_info'.
type UserMetaInfo struct {
	Id                  int
	UserId              int
	IsNotificationDirty int
	LastUserRecGen      int

	_exists, _deleted bool
}

/*
:= &UserMetaInfo {
	Id: 0,
	UserId: 0,
	IsNotificationDirty: 0,
	LastUserRecGen: 0,
*/

// UserPassword 'ms.user_password'.
type UserPassword struct {
	UserId      int
	Password    string
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &UserPassword {
	UserId: 0,
	Password: "",
	CreatedTime: 0,
*/
