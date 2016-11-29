package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func CreateNewComment(UserId, PostId int, Text string) Comment {
	cmt := Comment{
		UserId:      UserId,
		PostId:      PostId,
		CreatedTime: helper.TimeNow(),
		Text:        Text,
	}

	err := cmt.Insert(base.DB)

	if err == nil {
		Counter.IncerPostCommentsCount(PostId, 1)
		post, _ := CacheModels.GetPostById(PostId)
        Notification_OnPostCommented(&cmt, post)
	}

	return cmt
}

func RemoveComment(UserId, PostId, CommentId int) bool {
	post, _ := CacheModels.GetPostById(PostId)

	com, err := NewComment_Selector().Id_EQ(CommentId).UserId_EQ(UserId).PostId_EQ(PostId).GetRow(base.DB)
	if err != nil {
		com.Delete(base.DB)
		Counter.IncerPostCommentsCount(PostId, -1)
		Notification_OnPostCommentedDelted(com, post)
		return true
	}

	return false
}
