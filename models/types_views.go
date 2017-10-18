package models

import "ms/sun/models/x"

/////////////// Activity ///////////////
type ActivityView struct {
	*x.Activity
	Load interface{}
	//Actor   UserBasicAndMe
}

type ActivityPayload_DEP struct {
	Actor   *UserBasicAndMe
	Post    *x.Post
	Comment *x.Comments
}

type ActivityPayload struct {
	Actor    *UserInlineWithMeView
	Post     *PostView
	Comment  *x.Comments
	Followed *UserInlineWithMeView
}

//////////// Notifications ////////////
type NotificationView struct {
	*x.Notification
	Load interface{}
	//Actor   UserBasicAndMe
}

type NotifPayload struct {
	Actor    *UserInlineWithMeView
	Post     *PostView
	Comment  *x.Comments
	Followed *UserInlineWithMeView
}

type NotifPayload_DEP struct {
	Actor   *UserBasicAndMe
	Post    *x.Post
	Comment *x.Comments
}

//////////// Tags /////////////
type TopTagsWithPostsView struct {
	Tag   *x.Tag
	Posts []*PostView
}

///////////// Others //////////
type UserInlineView struct {
	// UserView
	UserId    int
	UserName  string
	FullName  string
	AvatarUrl string
}

type UserInlineWithMeView struct {
	UserInlineView
	FollowingType int
}

type PhotoView struct {
	PhotoId     int
	UserId      int
	PostId      int
	AlbumId     int
	ImageTypeId int
	Width       int
	Height      int
	Ratio       float32
	Color       string
	UrlFormat   string
	Sizes       []int //80,160,360
}

type PostView struct {
	*x.Post
	TypeName string //for text, photo video
	//Comments []CommentInlineInfo
	//Likes    []Likes
	PhotoView *PhotoView
	//Images   *ImageResHolder
	//AmIlike bool //dep
	MyLike int //type of like
	Sender UserInlineView
}

type CommentInlineInfoView struct {
	x.Comments
	Sender UserInlineView
}

////////////// USer ////////////
/////////// for Responses //////////////////////////////////////
type UserSyncAndMeView struct {
	//x.UserBasic
	//UserId        int
	//FollowingType int
	UserBasicAndMe
	AppVersion int
	Phone      string
	//UpdatedTime   int
}

//todo clean this 2 struct
type UserBasicAndMe struct { //legacy switch to UserTable
	x.UserBasic
	UserId        int
	UpdatedTime   int //rm ??
	FollowingType int /// 0: not_following  1: following  2: follow_requested
	//FollowingLists int
}

type ProfileInfo struct {
	x.UserCounts
	UserSyncAndMeView
}

/*type UserAndMe struct{
    UserSyncAndMeView
    x.UserCounts
} */

/////////////////////////////////////////////
