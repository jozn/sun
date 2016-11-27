package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func CreateNewComment(UserId, PostId int, Text string) Comment {
	cmt := Comment{
        UserId: UserId,
        PostId : PostId,
        CreatedTime:helper.TimeNow(),
        Text : Text,

    }

    err:=cmt.Insert(base.DB)

	if err == nil {
		Counter.IncerPostCommentsCount(PostId, 1)
        //Eventing
        //todo clean this up: post alwayse load
        post, _ := CacheModels.GetPostById(PostId)
        OnPostCommented(&cmt, post)
    }

	return cmt
}

func RemoveComment(UserId, PostId, CommentId int) bool {
	//Eventing
	//todo clean this up: post alwayse load
	post, _ := CacheModels.GetPostById(PostId)
	com, _ := CacheModels.GetCommentById(CommentId)

    com,err:= NewComment_Selector().Id_EQ(CommentId).UserId_EQ(UserId).PostId_EQ(PostId).GetRow(base.DB)
    if err!= nil{
        com.Delete(base.DB)
        Counter.IncerPostCommentsCount(PostId,-1)
        OnPostCommentedDelted(com, post)
        return true
    }

	return false
}
