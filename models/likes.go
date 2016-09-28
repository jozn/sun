package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func CreateLike(UserId, PostId int) {
	l := new(Like)
	l.UserId = UserId
	l.PostId = PostId
	l.CreatedTime = helper.TimeNow()

	_, err := base.DbInsertStruct(l, "likes")
	if err == nil {
		UserMemoryStore.AddPostLike(UserId, PostId)
		OnPostLiked(l)
	} else {
		helper.DebugPrintln(err)
	}
}

func AmILikePost(UserId, PostId int) bool {
	return false
}

func DeleteLike(UserId, PostId int) {
	l, err := GetLikeOf(UserId, PostId)
	if err != nil {
		return
	}

	q := "DELETE FROM likes WHERE UserId = ? AND PostId = ?"
	_, err = base.DB.Exec(q, UserId, PostId)
	if err == nil {
		UserMemoryStore.RemovePostLike(UserId, PostId)
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
