package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func Like_LikePost(UserId, PostId int) {
    l :=Like{
        UserId: UserId,
        PostId: PostId,
        TypeId:0,
        CreatedTime:helper.TimeNow(),
    }

	err := l.Insert(base.DB)
	if err == nil {
		NewPost_Updater().LikesCount_Increment(1).Id_EQ(PostId).Update(base.DB)
		MemoryStore.UserLikedPostsList_Add(UserId, PostId)
        Notification_OnPostLiked(l)
	}
}

func Like_UnlikePost(UserId, PostId int) {
	l, err := NewLike_Selector().UserId_EQ(UserId).PostId_EQ(PostId).GetRow(base.DB)
	if err != nil {
		return
	}

    err= l.Delete(base.DB)
	if err == nil {
		MemoryStore.UserLikedPostsList_Remove(UserId, PostId)
		Notification_OnPostUnLiked(l)
	}
}
