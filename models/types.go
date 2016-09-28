package models

type Activity struct {
	Id             int
	ActorUserId    int
	ActivityType   int
	TargetActionId int
	Privacy        int
	CreatedTime    int
    // xo fields
    _exists, _deleted bool
}

type Blocked struct {
	Id            int
	UserId        int
	BlockedUserID int
	CreatedTime   int
    // xo fields
    _exists, _deleted bool
}

type Comment struct {
	Id          int
	UserId      int
	PostId      int
	Text        string
	CreatedTime int
    // xo fields
    _exists, _deleted bool
}

type Like struct {
	Id          int
	PostId      int
	UserId      int
	TypeId      int
	CreatedTime int
    // xo fields
    _exists, _deleted bool
}

type Post struct {
	Id     int
	UserId int
	TypeId int
	Text   string

	MediaUrl      string
	MediaServerId int `json:"-"`
	Width         int
	Height        int
	FormatedText  string
	SharedTo      int
	HasTag        bool
	LikesCount    int
	CommentsCount int
	CreatedTime   int
    // xo fields
    _exists, _deleted bool
}

//not used in this version but maybe for futuer versions
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
	CreatedTime        int
    // xo fields
    _exists, _deleted bool
}

//Deprecated
type Media struct {
	Id          int
	UserId      int
	PostId      int
	AlbumId     int
	TypeId      int
	CreatedTime int
	Src         string
    // xo fields
    _exists, _deleted bool
}

