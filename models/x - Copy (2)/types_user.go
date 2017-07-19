package x

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
	UserNameLower        string `json:"-"`
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

/// User functions
func (ub *UserBasic) GetFullName() string {
	return ub.FirstName + " " + ub.LastName
}

/*
func (ub *UserBasic) ToUserInlineView() *UserInlineView {
    v := UserInlineView{}
    v.FullName = ub.GetFullName()
    v.UserId = ub.Id
    v.UserName = ub.UserName
    v.AvatarUrl = ub.AvatarUrl
    return &v
}
*/
