package models

/////////////// Activity ///////////////
type ActivityView struct {
	*Activity
	Load interface{}
	//Actor   UserBasicAndMe
}

type ActivityPayload struct {
	Actor   *UserBasicAndMe
	Post    *Post
	Comment *Comment
}

//////////// Notifications ////////////
type NotificationView struct {
	*Notification
	Load interface{}
	//Actor   UserBasicAndMe
}

type NotifPayload struct {
	Actor   *UserBasicAndMe
	Post    *Post
	Comment *Comment
}

//////////// Tags /////////////
type TopTagsWithPostsView struct {
	Tag   *Tag
	Posts []*PostView
}

///////////// Others //////////
type UserInlineView struct {
	// User
	UserId    int
	UserName  string
	FullName  string
	AvatarUrl string
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
	UrlFormat   string
	Sizes       []int //80,160,360
}

type PostView struct {
	*Post
	TypeName string //for text, photo video
	//Comments []CommentInlineInfo
	//Likes    []Like
	PhotoView *PhotoView
	//Images   *ImageResHolder
	AmIlike bool //dep
	MyLike  int  //type of like
	Sender  UserInlineView
}

type CommentInlineInfoView struct {
	Comment
	Sender UserInlineView
}
