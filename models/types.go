package models

type Activity struct {
    Id             int
    ActorUserId    int
    ActivityType   int
    TargetActionId int
    Privacy        int
    CreatedTime    int
}

type Blocked struct {
    Id            int
    UserId        int
    BlockedUserID int
    CreatedTime   int
}

type Comment struct {
    Id          int
    UserId      int
    PostId      int
    Text        string
    CreatedTime int
}


type Like struct {
    Id          int
    PostId      int
    UserId      int
    TypeId      int
    CreatedTime int
}

type Post struct {
	Id            int
	UserId        int
	TypeId        int
	Text          string

	MediaUrl      string
	MediaServerId int `json:"-"`
    Width int
    Height int
	FormatedText  string
	SharedTo      int
	HasTag        bool
	LikesCount    int
	CommentsCount int
	CreatedTime   int
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
}
