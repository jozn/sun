package x

// activity 'Activity'.
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
// bucket 'Bucket'.
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
// chat 'Chat'.
type Chat struct {
	ChatKey              string
	RoomKey              string
	RoomTypeEnumId       int
	UserId               int
	PeerUserId           int
	GroupId              int
	CreatedSe            int
	StartMessageIdFrom   int
	LastDeletedMessageId int
	LastSeenMessageId    int
	LastMessageId        int //just direct-not reliable- just increment not decremetnt
	UpdatedMs            int

	_exists, _deleted bool
}

/*
:= &Chat {
	ChatKey: "",
	RoomKey: "",
	RoomTypeEnumId: 0,
	UserId: 0,
	PeerUserId: 0,
	GroupId: 0,
	CreatedSe: 0,
	StartMessageIdFrom: 0,
	LastDeletedMessageId: 0,
	LastSeenMessageId: 0,
	LastMessageId: 0,
	UpdatedMs: 0,
*/
// comment 'Comment'.
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
// direct_message 'DirectMessage'.
type DirectMessage struct {
	MessageId            int
	MessageKey           string
	RoomKey              string
	UserId               int
	MessageFileId        int
	MessageTypeEnumId    int
	Text                 string
	CreatedSe            int
	PeerReceivedTime     int
	PeerSeenTime         int
	DeliviryStatusEnumId int

	_exists, _deleted bool
}

/*
:= &DirectMessage {
	MessageId: 0,
	MessageKey: "",
	RoomKey: "",
	UserId: 0,
	MessageFileId: 0,
	MessageTypeEnumId: 0,
	Text: "",
	CreatedSe: 0,
	PeerReceivedTime: 0,
	PeerSeenTime: 0,
	DeliviryStatusEnumId: 0,
*/
// direct_offline 'DirectOffline'.
type DirectOffline struct {
	DirectOfflineId int
	ToUserId        int
	ChatKey         string
	MessageId       int
	MessageFileId   int
	PBClass         string
	DataPB          []byte
	DataJson        string
	DataTemp        string //not real data
	AtTimeMs        int

	_exists, _deleted bool
}

/*
:= &DirectOffline {
	DirectOfflineId: 0,
	ToUserId: 0,
	ChatKey: "",
	MessageId: 0,
	MessageFileId: 0,
	PBClass: "",
	DataPB: []byte{},
	DataJson: "",
	DataTemp: "",
	AtTimeMs: 0,
*/
// direct_offline_dep 'DirectOfflineDep'.
type DirectOfflineDep struct {
	DirectOfflineId int
	ToUserId        int
	MessageId       int
	MessageFileId   int
	OtherId         int
	ChatKey         string
	PeerUserId      int
	RoomLogTypeId   int
	PBClass         string
	DataPB          []byte
	DataJson        string
	DataTemp        string //not real data
	AtTimeMs        int

	_exists, _deleted bool
}

/*
:= &DirectOfflineDep {
	DirectOfflineId: 0,
	ToUserId: 0,
	MessageId: 0,
	MessageFileId: 0,
	OtherId: 0,
	ChatKey: "",
	PeerUserId: 0,
	RoomLogTypeId: 0,
	PBClass: "",
	DataPB: []byte{},
	DataJson: "",
	DataTemp: "",
	AtTimeMs: 0,
*/
// direct_to_message 'DirectToMessage'.
type DirectToMessage struct {
	Id           int
	ChatKey      string
	MessageId    int
	SourceEnumId int

	_exists, _deleted bool
}

/*
:= &DirectToMessage {
	Id: 0,
	ChatKey: "",
	MessageId: 0,
	SourceEnumId: 0,
*/
// direct_update 'DirectUpdate'.
type DirectUpdate struct {
	DirectUpdateId int
	ToUserId       int
	MessageId      int
	MessageFileId  int
	OtherId        int
	ChatKey        string
	PeerUserId     int
	EventType      int
	RoomLogTypeId  int
	FromSeq        int
	ToSeq          int
	ExtraPB        []byte
	ExtraJson      string
	AtTimeMs       int

	_exists, _deleted bool
}

/*
:= &DirectUpdate {
	DirectUpdateId: 0,
	ToUserId: 0,
	MessageId: 0,
	MessageFileId: 0,
	OtherId: 0,
	ChatKey: "",
	PeerUserId: 0,
	EventType: 0,
	RoomLogTypeId: 0,
	FromSeq: 0,
	ToSeq: 0,
	ExtraPB: []byte{},
	ExtraJson: "",
	AtTimeMs: 0,
*/
// following_list 'FollowingList'.
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
// following_list_member 'FollowingListMember'.
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
// following_list_member_history 'FollowingListMemberHistory'.
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
// general_log 'GeneralLog'.
type GeneralLog struct {
	Id        int
	ToUserId  int
	TargetId  int
	LogTypeId int
	ExtraPB   []byte
	ExtraJson string
	CreatedMs int

	_exists, _deleted bool
}

/*
:= &GeneralLog {
	Id: 0,
	ToUserId: 0,
	TargetId: 0,
	LogTypeId: 0,
	ExtraPB: []byte{},
	ExtraJson: "",
	CreatedMs: 0,
*/
// group 'Group'.
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
// group_member 'GroupMember'.
type GroupMember struct {
	Id              int
	GroupId         int
	GroupKey        string
	UserId          int
	ByUserId        int
	GroupRoleEnumId int
	CreatedTime     int

	_exists, _deleted bool
}

/*
:= &GroupMember {
	Id: 0,
	GroupId: 0,
	GroupKey: "",
	UserId: 0,
	ByUserId: 0,
	GroupRoleEnumId: 0,
	CreatedTime: 0,
*/
// group_message 'GroupMessage'.
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
// group_to_message 'GroupToMessage'.
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
// likes 'Like'.
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
// log_changes 'LogChange'.
type LogChange struct {
	Id int
	T  string

	_exists, _deleted bool
}

/*
:= &LogChange {
	Id: 0,
	T: "",
*/
// media 'Media'.
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
// message_file 'MessageFile'.
type MessageFile struct {
	MessageFileId   int
	MessageFileKey  string
	OriginalUserId  int
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

	_exists, _deleted bool
}

/*
:= &MessageFile {
	MessageFileId: 0,
	MessageFileKey: "",
	OriginalUserId: 0,
	Name: "",
	Size: 0,
	FileTypeEnumId: 0,
	Width: 0,
	Height: 0,
	Duration: 0,
	Extension: "",
	HashMd5: "",
	HashAccess: 0,
	CreatedSe: 0,
	ServerSrc: "",
	ServerPath: "",
	ServerThumbPath: "",
	BucketId: "",
	ServerId: 0,
	CanDel: 0,
*/
// msg 'Msg'.
type Msg struct {
	Key  string
	Name string
	Id   int

	_exists, _deleted bool
}

/*
:= &Msg {
	Key: "",
	Name: "",
	Id: 0,
*/
// notification 'Notification'.
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
// notification_removed 'NotificationRemoved'.
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
// offline 'Offline'.
type Offline struct {
	Id          int
	FromUserId  int
	ToUserId    int
	RpcName     string
	PBClassName string
	Key         string
	DataJson    string
	DataBlob    []byte
	DataLength  int
	CreatedMs   int

	_exists, _deleted bool
}

/*
:= &Offline {
	Id: 0,
	FromUserId: 0,
	ToUserId: 0,
	RpcName: "",
	PBClassName: "",
	Key: "",
	DataJson: "",
	DataBlob: []byte{},
	DataLength: 0,
	CreatedMs: 0,
*/
// old_messages 'OldMessage'.
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
	DataPB: []byte{},
	Data64: "",
	DataJson: "",
	CreatedTimeMs: 0,
*/
// old_msg_file 'OldMsgFile'.
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
	ThumbData: []byte{},
	ThumbData64: "",
	ServerSrc: "",
	ServerPath: "",
	ServerId: 0,
	CanDel: 0,
*/
// old_msg_push 'OldMsgPush'.
type OldMsgPush struct {
	Id            int
	Uid           int
	ToUser        int
	MsgUid        int
	CreatedTimeMs int

	_exists, _deleted bool
}

/*
:= &OldMsgPush {
	Id: 0,
	Uid: 0,
	ToUser: 0,
	MsgUid: 0,
	CreatedTimeMs: 0,
*/
// old_msg_push_event 'OldMsgPushEvent'.
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
// phone_contacts 'PhoneContact'.
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
// photo 'Photo'.
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
// post 'Post'.
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
// push_event 'PushEvent'.
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
// push_message 'PushMessage'.
type PushMessage struct {
	PushMessageId int
	ToUserId      int
	ToDeviceId    int
	MessageId     int
	RoomTypeEnum  int
	CreatedMs     int

	_exists, _deleted bool
}

/*
:= &PushMessage {
	PushMessageId: 0,
	ToUserId: 0,
	ToDeviceId: 0,
	MessageId: 0,
	RoomTypeEnum: 0,
	CreatedMs: 0,
*/
// recommend_user 'RecommendUser'.
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
// room 'Room'.
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
// search_clicked 'SearchClicked'.
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
// session 'Session'.
type Session struct {
	Id                      int
	UserId                  int
	SessionUuid             string
	ClientUuid              string
	DeviceUuid              string
	LastActivityTime        int
	LastIpAddress           string
	LastWifiMacAddress      string
	LastNetworkType         string
	LastNetworkTypeId       int
	AppVersion              int
	UpdatedTime             int
	CreatedTime             int
	LastSyncDirectUpdateId  int
	LastSyncGeneralUpdateId int
	LastSyncNotifyUpdateId  int

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
	LastSyncDirectUpdateId: 0,
	LastSyncGeneralUpdateId: 0,
	LastSyncNotifyUpdateId: 0,
*/
// setting_client 'SettingClient'.
type SettingClient struct {
	UserId                    int
	AutoDownloadWifiVoice     int
	AutoDownloadWifiImage     int
	AutoDownloadWifiVideo     int
	AutoDownloadWifiFile      int
	AutoDownloadWifiMusic     int
	AutoDownloadWifiGif       int
	AutoDownloadCellularVoice int
	AutoDownloadCellularImage int
	AutoDownloadCellularVideo int
	AutoDownloadCellularFile  int
	AutoDownloadCellularMusic int
	AutoDownloadCellularGif   int
	AutoDownloadRoamingVoice  int
	AutoDownloadRoamingImage  int
	AutoDownloadRoamingVideo  int
	AutoDownloadRoamingFile   int
	AutoDownloadRoamingMusic  int
	AutoDownloadRoamingGif    int
	SaveToGallery             int

	_exists, _deleted bool
}

/*
:= &SettingClient {
	UserId: 0,
	AutoDownloadWifiVoice: 0,
	AutoDownloadWifiImage: 0,
	AutoDownloadWifiVideo: 0,
	AutoDownloadWifiFile: 0,
	AutoDownloadWifiMusic: 0,
	AutoDownloadWifiGif: 0,
	AutoDownloadCellularVoice: 0,
	AutoDownloadCellularImage: 0,
	AutoDownloadCellularVideo: 0,
	AutoDownloadCellularFile: 0,
	AutoDownloadCellularMusic: 0,
	AutoDownloadCellularGif: 0,
	AutoDownloadRoamingVoice: 0,
	AutoDownloadRoamingImage: 0,
	AutoDownloadRoamingVideo: 0,
	AutoDownloadRoamingFile: 0,
	AutoDownloadRoamingMusic: 0,
	AutoDownloadRoamingGif: 0,
	SaveToGallery: 0,
*/
// setting_notifications 'SettingNotification'.
type SettingNotification struct {
	UserId                   int
	SocialLedOn              int
	SocialLedColor           string
	ReqestToFollowYou        int
	FollowedYou              int
	AccptedYourFollowRequest int
	YourPostLiked            int
	YourPostCommented        int
	MenthenedYouInPost       int
	MenthenedYouInComment    int
	YourContactsJoined       int
	DirectMessage            int
	DirectAlert              int
	DirectPerview            int
	DirectLedOn              int
	DirectLedColor           int
	DirectVibrate            int
	DirectPopup              int
	DirectSound              int
	DirectPriority           int

	_exists, _deleted bool
}

/*
:= &SettingNotification {
	UserId: 0,
	SocialLedOn: 0,
	SocialLedColor: "",
	ReqestToFollowYou: 0,
	FollowedYou: 0,
	AccptedYourFollowRequest: 0,
	YourPostLiked: 0,
	YourPostCommented: 0,
	MenthenedYouInPost: 0,
	MenthenedYouInComment: 0,
	YourContactsJoined: 0,
	DirectMessage: 0,
	DirectAlert: 0,
	DirectPerview: 0,
	DirectLedOn: 0,
	DirectLedColor: 0,
	DirectVibrate: 0,
	DirectPopup: 0,
	DirectSound: 0,
	DirectPriority: 0,
*/
// tag 'Tag'.
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
// tags_posts 'TagsPost'.
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
// test_chat 'TestChat'.
type TestChat struct {
	Id     int
	Id4    int
	TimeMs int
	Text   string
	Name   string
	UserId int
	C2     int
	C3     int
	C4     int
	C5     int

	_exists, _deleted bool
}

/*
:= &TestChat {
	Id: 0,
	Id4: 0,
	TimeMs: 0,
	Text: "",
	Name: "",
	UserId: 0,
	C2: 0,
	C3: 0,
	C4: 0,
	C5: 0,
*/
// trigger_log 'TriggerLog'.
type TriggerLog struct {
	Id         int
	ModelName  string
	ChangeType string
	TargetId   int
	TargetStr  string
	CreatedSe  int

	_exists, _deleted bool
}

/*
:= &TriggerLog {
	Id: 0,
	ModelName: "",
	ChangeType: "",
	TargetId: 0,
	TargetStr: "",
	CreatedSe: 0,
*/
// user_meta_info 'UserMetaInfo'.
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
// user_password 'UserPassword'.
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