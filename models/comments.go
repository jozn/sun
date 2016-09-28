package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func CreateNewComment(UserId, PostId int, Text string) Comment {
	cmt := Comment{}
	cmt.UserId = UserId

	cmt.PostId = PostId
	cmt.CreatedTime = helper.TimeNow()
	//todo format text for security and size ,parsing ,tages,...
	cmt.Text = Text

	res, err := base.DbInsertStruct(&cmt, "comments")
	if err == nil {
		id, _ := res.LastInsertId()
		cmt.Id = int(id)
		QueryIncerPostCommentsCount(PostId, 1)
	}

	//Eventing
	//todo clean this up: post alwayse load
	post, _ := CacheModels.GetPostById(PostId)
	OnPostCommented(&cmt, post)

	return cmt
}

func RemoveComment(UserId, PostId, CommentId int) bool {
	//Eventing
	//todo clean this up: post alwayse load
	post, _ := CacheModels.GetPostById(PostId)
	com, _ := CacheModels.GetCommentById(CommentId)

	q := "DELETE FROM comments WHERE UserId = ? AND PostId = ? AND Id = ?"
	res, err := base.DbExecute(q, UserId, PostId, CommentId)

	removed := false
	if err == nil {
		cnt, err := res.RowsAffected()
		if cnt == 1 {
			QueryDecerPostCommentsCount(PostId, 1)
			removed = true
		}
		helper.DebugPrintln(err)

		OnPostCommentedDelted(com, post)
	}
	helper.DebugPrintln(err)
	return removed
}
