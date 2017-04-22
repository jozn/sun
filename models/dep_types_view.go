package models

type PostAndDetailes struct {
	// *Post
	Post
	TypeName string //for text, photo video
	Comments []CommentInlineInfo
	Likes    []Like
	AmIlike  bool
	Sender   UserInlineView
}

type CommentInlineInfo struct {
	Comment
	Sender UserInlineView
}
