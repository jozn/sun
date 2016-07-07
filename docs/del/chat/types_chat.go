package chat

//Sole cilent resposibility -- no sereverside fields
//build on data at hand in cilents(android)
type RoomListTableSpec struct {
	//	Id                        int
	RoomKey                   string
	RoomTypeId                int
	RoomName                  string
	UseCustomRoomSettings     int
	LastMessageType           int
	LastMessageText           string
	LastMessageDeliveryStatus int
	LastRoomOpenedTimestamp   int
	UnseenMessageCount        int
	LastSeenMessageKey        string
	MessageDraftText          string
	UpdatedTimestampMs        int
	CreatedTimestamp          int
}

//equal to client
type GroupMember struct {
	Id               int `json:"-"`
	RoomKey          string
	UserId           int
	ByUserId         int
	RoleId           int
	IsAdmin          int //???
	CreatedTimestamp int
}

type GroupMembersTableSpec struct {
	GroupMember // - Id
}

//equal to client - beacuse of many users
type GroupInfo struct {
	RoomKey          string
	RoomName         string
	MembersCount     int
	GroupPrivacyId   int
	CreatorUserId    int
	CreatedTimestamp int
	UpdatedTimestamp int
}

type GroupInfoTableSpec struct {
	GroupInfo
}

//for now sole client
type CustomRoomSettingsTableSpec struct {
	RoomKey        string
	BackgroundId   int
	BackgroundPath string
}

//equal to client
type BlockedUser struct {
	Id               int `json:"-"`
	UserId           int
	BlockedUserId    int
	CreatedTimestamp int
}

type BlockedUsersTableSpec struct {
	BlockedUser
}

//contacts is based on unique phone number
type ContactsTableSpec struct {
	Id                    int `json:"-"`
	PhoneContactRowId     int
	IsPhoneContact        int
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneDisplayName      string
	PhoneFamilyName       string
	IsFollowing           int
	MsUserId              int
	UserId                int
	AvatarUrl             string
	StatusText            string
	StatusId              string
	UpdateTimestamp       string
	UnseenMessageCount    int
}

type ContactUserInfo struct {
	PhoneNormalizedNumberInt int //TODO:add to user table
	IsFollowing              int
	UserId                   int
	AvatarUrl                string
	StatusText               string
	StatusId                 string
	UpdateTimestamp          string
	//	UnseenMessageCount    int
}

type MessagesTableSpec struct {
	MessagesTable
	RoomTypeId              int
	MediaLocalPath          string
	IsByMe                  int
	IsSeen                  int
	IsStared                int
	DeliveryStatus          int
	SeenTimestamp           int//did changed
	ServerReceivedTimestamp int//did changed
	ClientReceivedTimestamp int//did changed
	ServerDeletedTimestamp  int//did changed
	//MessageKey              string
	//RoomKey                 string
	//UserId                  int
	//MessageTypeId           int
	//Text                    string
	//MediaUrl                string
	//ExtraJson               string
	//CreatedTimestampMs      int
}

type MessagesTable struct {
	MessageKey         string
	RoomKey            string
	UserId             int
	MessageTypeId      int
	Text               string
	MediaUrl           string
	ExtraJson          string
	CreatedMs int
	//CreatedTimestampMs int
}



/////////////////////////

type PhoneContact struct {
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     int
}

type PhoneContactTable struct {
	//	PhoneContact
	Id int
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     int
	UserId                int
	DeviceUuidId          int
	CreatedTimeStamp      int
	UpdatedTimeStamp      int
}



//deprectaed
type ContactPhone struct {
	PhoneContactRowId        int
	IsPhoneContact           int
	PhoneNumber              string
	PhoneNormalizedNumber    string
	PhoneNormalizedNumberInt int
	PhoneDisplayName         string
	PhoneFamilyName          string
}
//dep
/////////////////from Android client
type ClientPhoneContact struct {
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     int
}




