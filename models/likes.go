package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func Like_LikePost(UserId, PostId int) {
	p, ok := Store.GetPostById(PostId)
	if !ok {
		return
	}
	l := &Like{
		UserId:      UserId,
		PostTypeId:  p.TypeId,
		PostId:      PostId,
		TypeId:      0, //emotions
		CreatedTime: helper.TimeNow(),
	}

	err := l.Insert(base.DB)
	if err == nil {
		NewPost_Updater().LikesCount_Increment(1).Id_Eq(PostId).Update(base.DB)
		MemoryStore.UserLikedPostsList_Add(UserId, PostId)
		Notification_OnPostLiked(l)
		Activity_OnPostLiked(l)
	}
}

func Like_UnlikePost(UserId, PostId int) {
	l, err := NewLike_Selector().UserId_Eq(UserId).PostId_Eq(PostId).GetRow(base.DB)
	if err != nil {
		return
	}

	err = l.Delete(base.DB)
	if err == nil {
		MemoryStore.UserLikedPostsList_Remove(UserId, PostId)
        NewPost_Updater().LikesCount_Increment(-1).Id_Eq(PostId).Update(base.DB)
		Notification_OnPostUnLiked(l)
		Activity_OnPostUnLiked(l)
	}
}
