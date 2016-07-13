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

type FollowingList struct {
    Id          int
    UserId      int
    ListType    int
    Name        string
    Count       int
    IsAuto      int
    IsPimiry    int
    CreatedTime int
}

type FollowingListMember struct {
    Id             int
    ListId         int
    UserId         int
    FollowedUserId int
    FollowType     int
    UpdatedTimeMs  int
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
	FormatedText  string
	SharedTo      int
	HasTag        bool
	LikesCount    int
	CommentsCount int
	CreatedTime   int
}

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

type Tag struct {
    Id          int
    Name        string
    Count       int
    IsBlocked   bool
    CreatedTime int
}

type TagPost struct {
    Id          int
    TagId       int
    PostId      int
    TypeId      int // text? photo? video?
    CreatedTime int
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
