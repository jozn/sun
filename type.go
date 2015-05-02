package main

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

type FollowingListMemeber struct {
	Id               int
	ListId           int
	FollowdUserId    int
	CreatedTimestamp int
}

type Like struct {
	Id               int
	PostId           int
	UserId           int
	TypeID           int
	CreatedTimestamp int
}

type Media struct {
	Id               int
	UserId           int
	PostId           int
	AlbumId          int
	TypeID           int
	CreatedTimestamp int
	Src              string
}

type Post struct {
	Id               int
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

type User struct {
	Id int
	// UserId int
	UserName         string
	FirstName        string
	LastName         string
	FullName         string
	Email            string `json:"-"`
	PasswordHash     string `json:"-"`
	PasswordSalt     string `json:"-"`
	IsProfilePrivate int
	CreatedTimestamp int `json:"-"`
}

type UserInfo struct {
	UserId             int
	FollowersCount     int
	FollowingCount     int
	PostsCount         int
	MediaCount         int
	LastLoginTimestamp int
}

type UserPassword struct {
	UserId           int
	Password         int
	CreatedTimestamp int
}

///////views for json///////

type ImageResHolder struct {
	LowRes      *ImageResRow //320
	StandardRes *ImageResRow //640
	High        *ImageResRow //? 960 ?
	ThumbRes    *ImageResRow //150
	Tiny150Res  *ImageResRow
	Tiny100Res  *ImageResRow
	Tiny50Res   *ImageResRow
}

type ImageResRow struct {
	Width  int
	Height int
	Url    string
}

type PostAndDetailes struct {
	*Post
	Comments []Comment
	Likes    []Like
	Images   ImageResHolder
	AmIlike  bool
	Sender
}

//for comments and mini-likes
type UserTinyInfoView struct {
	UserId   int
	UserName string
	FullName string
	// Avatars Avatars
}

////play//////////

type ViewUser struct {
	*User
	Counts   *UserInfo
	PlayPLAY Play1
}

type Play1 struct {
	*User
	Counts *UserInfo
	Session
	SEEEE *Session
}
