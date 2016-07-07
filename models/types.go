package models

type Activity struct {
	Id               int
	ActorUserId      int
	ActivityType     int
	TargetActionId   int
	Privacy          int
	CreatedTimestamp int
}

type Blocked struct {
	Id               int
	UserId           int
	BlockedUserID    int
	CreatedTimestamp int
}

type Comment struct {
	Id               int
	UserId           int
	PostId           int
	Text             string
	CreatedTimestamp int
}

type FollowingList struct {
	Id               int
	UserId           int
	ListType         int
	Name             string
	Count            int
	IsAuto           int
	IsPimiry         int
	CreatedTimestamp int
}

type FollowingListMember struct {
	Id               int
	ListId           int
	UserId           int
	FollowedUserId   int
	FollowType   int
	UpdatedTimestampMs int
}

type Like struct {
	Id               int
	PostId           int
	UserId           int
	TypeId           int
	CreatedTimestamp int
}

type Media struct {
	Id               int
	UserId           int
	PostId           int
	AlbumId          int
	TypeId           int
	CreatedTimestamp int
	Src              string
}

type Post struct {
	Id int
	// PostId           int
	UserId           int
	TypeId           int
	Text             string
	FormatedText     string
	SharedTo         int
	HasTag           bool
	LikesCount       int
	CommentsCount    int
	CreatedTimestamp int
}

type Session struct {
	Id                    int
	UserId                int
	SessionUuid           string
	ClientUuid            string
	DeviceUuid            string
	LastActivityTimestamp int
	LastIpAddress         string
	LastWifiMacAddress    string
	LastNetworkType       string
	CreatedTimestamp      int
}

type Tag struct {
	Id               int
	Name             string
	Count            int
	IsBlocked        bool
	CreatedTimestamp int
}

type TagPost struct {
	Id               int
	TagId            int
	PostId           int
	TypeId           int // text? photo? video?
	CreatedTimestamp int
}
