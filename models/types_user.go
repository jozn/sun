package models

type UserBasic struct {
	Id        int
	UserName  string
	FirstName string
	LastName  string
	FullName  string //deprecated
	AvatarSrc string //dep
	AvatarUrl string
	IsProfilePrivate     int
}

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
	LikesCount     int
	ResharedCount  int
}

type UserExtra struct {
	Email                string `json:"-"`
	PrimaryFollowingList int
	CreatedTimestamp     int `json:"-"`
	UpdatedTimestamp     int //`json:"-"`
	LastLoginTimestamp   int
}

// +gen slice:"Where,GroupBy[string],DistinctBy,SortBy,Select[string]"
type User struct { //legacy switch to UserTable
	UserTable
}
type UserTable struct {
	UserBasic
	UserPhone
	UserSecurity
	UserCounts
	UserExtra
}

/////////// for Responses ///////////////////////////////////////

type UserBasicAndMe struct { //legacy switch to UserTable
	UserBasic
	UserId int
	UpdatedTimestamp     int
	AmIFollowing int//dep
	FollowingType  int /// 0: not_following  1: following  2: follow_requested
	//FollowingLists int
}

type UserBasicPhoneAndMe struct {
	UserBasicAndMe
	UserPhone
}

//type UserBasicContactsAndMe struct {
//	UserBasicAndMe
//	PhoneContact
//}

func (t *UserBasicAndMe) FromUser(u UserTable,list FollowChecker )  {
	t.UserBasic = u.UserBasic
	t.UserId = u.Id
	t.UpdatedTimestamp = u.UpdatedTimestamp
	t.FollowingType = list.FollowingType(u.Id)
}

func (t *UserBasicAndMe) FormUserAndMe(u UserTable, meId int )  {
	t.UserBasic = u.UserBasic
	t.UserId = u.Id
	t.UpdatedTimestamp = u.UpdatedTimestamp
	t.FollowingType = GetFollowingType(meId,u.Id)
}

/////////// Deprecated //////////////////////////////

type UserInfo struct { //deprecated
	UserId             int
	FollowersCount     int
	FollowingCount     int
	PostsCount         int
	MediaCount         int
	LikesCount         int
	ResharedCount      int
	LastLoginTimestamp int
}

type UserPassword struct {
	UserId           int
	Password         int
	CreatedTimestamp int
}

/////////////////////////////////////////////////////////

//type User struct {
//	Id int
//				    // UserId int
//	UserName             string
//	FirstName            string
//	LastName             string
//	FullName             string //deprecated
//	Email                string `json:"-"`
//	Phone                string //`json:"-"`
//	PrimaryFollowingList int
//	PasswordHash         string `json:"-"`
//	PasswordSalt         string `json:"-"`
//	IsProfilePrivate     int
//	AvatarSrc            string
//	CreatedTimestamp     int `json:"-"`
//	UpdatedTimestamp     int //`json:"-"`
//	FollowersCount     int
//	FollowingCount     int
//	PostsCount         int
//	MediaCount         int
//	LikesCount         int
//	ResharedCount      int
//	LastLoginTimestamp int
//	IsDeleted int
//}

//type User struct {
//	Email                string `json:"-"`
//	PrimaryFollowingList int
//	IsProfilePrivate     int
//	CreatedTimestamp     int `json:"-"`
//	UpdatedTimestamp     int //`json:"-"`
//	LastLoginTimestamp int
//}
