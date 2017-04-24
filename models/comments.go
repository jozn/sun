package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func Comment_Add(UserId, PostId int, Text string) Comment {
	cmt := Comment{
		UserId:      UserId,
		PostId:      PostId,
		CreatedTime: helper.TimeNow(),
		Text:        Text,
	}

	err := cmt.Insert(base.DB)

	if err == nil {
		Counter.IncerPostCommentsCount(PostId, 1)
		post, ok := Store.GetPostById(PostId)
		if ok {
			Notification_OnPostCommented(&cmt, post)
			Activity_OnPostCommentAdd(&cmt, post)
		}
	}

	return cmt
}

func Comment_Delete(UserId, PostId, CommentId int) bool {
	post, _ := Store.GetPostById(PostId)

	com, err := NewComment_Selector().Id_EQ(CommentId).UserId_EQ(UserId).PostId_EQ(PostId).GetRow(base.DB)
	if err != nil {
		com.Delete(base.DB)
		Counter.IncerPostCommentsCount(PostId, -1)
		Notification_OnPostCommentedDeleted(com, post)
		Activity_OnPostCommentDeleted(com, post)
		return true
	}

	return false
}
