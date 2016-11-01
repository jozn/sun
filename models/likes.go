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

func AmILikePost(UserId, PostId int) bool {
	return MemoryStore.UserLikedPostsList_IsLiked(UserId, PostId)
}

func DeletePostLike(UserId, PostId int) {
	l, err := GetLikeOf(UserId, PostId)
	if err != nil {
		return
	}

	q := "DELETE FROM likes WHERE UserId = ? AND PostId = ?"
	res, err := base.DB.Exec(q, UserId, PostId)
	if err == nil {
		if n, _ := res.RowsAffected(); n > 0 {
			NewPost_Updater().LikesCount_Increment(-1).Id_EQ(PostId).Update(base.DB)
		}
		//UserMemoryStore.RemovePostLike(UserId, PostId)
		MemoryStore.UserLikedPostsList_Remove(UserId, PostId)
		OnPostUnLiked(l)
	} else {
		helper.DebugPrintln(err)
	}
}

//todo: move to CacheModel
func GetLikeOf(UserId, PostId int) (*Like, error) {
	l := new(Like)
	q := "select * FROM likes WHERE UserId = ? AND PostId = ?"
	err := base.DB.Get(l, q, UserId, PostId)
	if err == nil {
		return l, nil
	}
	return nil, err
}
