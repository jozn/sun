package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func CreatePostLike(UserId, PostId int) {
	l := new(Like)
	l.UserId = UserId
	l.PostId = PostId
	l.CreatedTime = helper.TimeNow()

	//_, err := base.DbInsertStruct(l, "likes")

	err := l.Insert(base.DB)
	if err == nil {
		//UserMemoryStore.AddPostLike(UserId, PostId)
		NewPost_Updater().LikesCount_Increment(1).Id_EQ(PostId).Update(base.DB)
		MemoryStore.UserLikedPostsList_Add(UserId, PostId)
		OnPostLiked(l)
	} else {
		helper.DebugPrintln(err)
	}
}

func DeletePostLike(UserId, PostId int) {
	l, err := NewLike_Selector().UserId_EQ(UserId).PostId_EQ(PostId).GetRow(base.DB)
	if err != nil {
		return
	}

	q := "DELETE FROM likes WHERE UserId = ? AND PostId = ?"
	res, err := base.DB.Exec(q, UserId, PostId)
	if err == nil {
		if n, _ := res.RowsAffected(); n > 0 {
			NewPost_Updater().LikesCount_Increment(-1).Id_EQ(PostId).Update(base.DB)
		}
		MemoryStore.UserLikedPostsList_Remove(UserId, PostId)
		OnPostUnLiked(l)
	} else {
		helper.DebugPrintln(err)
	}
}


