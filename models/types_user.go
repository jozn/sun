package models

// +gen slice:"Where,Count,GroupBy[string]"
type UserBasic struct {
	Id        int
	UserName  string
	FirstName string
	LastName  string
	FullName  string //deprecated
	//AvatarSrc      string //dep
	AvatarUrl      string
	PrivacyProfile int
	About          string
}

// +gen * slice:"Where,Count,GroupBy[int],Any,Average[int],Average[int8]"
type UserPhone struct {
	Phone string //`json:"-"`
}

type UserSecurity struct {
	IsDeleted    int
	PasswordHash string `json:"-"`
	PasswordSalt string `json:"-"`
}

type UserCounts struct {
	FollowersCount int
	FollowingCount int
	PostsCount     int
	MediaCount     int
	LikesCount     int //unrelable
	ResharedCount  int
	CommentsCount  int //unrelable
}

type UserExtra struct {
	Email                string `json:"-"`
	PrimaryFollowingList int
	CreatedTime          int `json:"-"`
	UpdatedTime          int //`json:"-"`

	LastPostTime   int
	LastActionTime int //post ,comment, like ,... === for activity in memmoery
}

type UserSession struct {
	SessionUuid        string
	DeviceUuid         string
	AppVersion         int
	LastLoginTime      int
	LastActivityTime   int
	LastIpAddress      string
	LastWifiMacAddress string
	LastNetworkType    string
}

// +gen slice:"Where,GroupBy[string],DistinctBy,SortBy,Select[string]"
type User struct { //legacy switch to UserTable
	UserTable
	// xo fields
	_exists, _deleted bool
}

type UserTable struct {
	UserBasic
	UserPhone
	UserSecurity
	UserCounts
	UserSession
	UserExtra
}

/////////// for Responses //////////////////////////////////////
type UserViewSyncAndMe struct {
	UserBasic
	UserId        int
	FollowingType int
	AppVersion    int
	Phone         string
	UpdatedTime   int
}

//todo clean this 2 struct
type UserBasicAndMe struct { //legacy switch to UserTable
	UserBasic
	UserId        int
	UpdatedTime   int //rm ??
	FollowingType int /// 0: not_following  1: following  2: follow_requested
	//FollowingLists int
}

/////////////////////////////////////////////
/// User functions
func (ub *UserBasic) GetFullName() string {
	return ub.FirstName + " " + ub.LastName
}

func (ub *UserBasic) ToUserInlineView() *UserInlineView {
	v := UserInlineView{}
	v.FullName = ub.GetFullName()
	v.UserId = ub.Id
	v.UserName = ub.UserName
	v.AvatarUrl = ub.AvatarUrl
	return &v
}
