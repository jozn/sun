package models

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
