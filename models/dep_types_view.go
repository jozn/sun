package models

///////views for json////////

type ImageResHolder struct {
	LowRes      *ImageResRow //320
	StandardRes *ImageResRow //640
	HighRes     *ImageResRow //? 960 ?
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
	// *Post
	Post
	TypeName string //for text, photo video
	Comments []CommentInlineInfo
	Likes    []Like
	Images   *ImageResHolder
	AmIlike  bool
	Sender   UserInlineView
}

type CommentInlineInfo struct {
	Comment
	Sender UserInlineView
}

/*
type LikelineInfo struct {
	UserInlineInfo
}
*/

/*
type UserMajorInfoMe struct {
	User
	WithMe
}
*/

/*
type WithMe struct {
	User
	AmIFollowing bool
}
*/

type UserInlineFollowView struct {
	UserInlineView
	IsProfilePrivate bool
	AmIFollowing     bool
	IFollowType      int
	FollowingLists   []int
}

////play and deprecated//////////

//for comments and mini-likes
//dep :use UserView
//dep
/*type Avatar struct {
	Size int
	Src  string
}*/

/*
type UserInlineInfo struct {
	UserId   int
	UserName string
	FullName string
	// Avatars Avatars
}
*/

/*

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
*/

/*
type ProfileView struct {
	User         User
	UserInfo     UserInfo //usercounts
	Images       []ImageResHolder
	SimilerUsers []UserInlineFollowView
	Posts        []PostAndDetailes
	OtherInfo    User
}
*/
