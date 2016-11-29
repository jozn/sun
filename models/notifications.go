package models

import (
	"math"
	"ms/sun/base"
	"ms/sun/helper"
)

// In Orma-gen
type Notification struct {
	Id           int
	ForUserId    int
	ActorUserId  int
	ActionTypeId int //added_post ,added_comment == NotifyTypeId
	ObjectTypeId int //post , comment
	TargetId     int // PostId, CommentId
	ObjectId     int //dep == TargetId * 1000 + ObjectTypeId  ---- eg: 1256*1000 + 3 = 1256003
	SeenStatus   int
	CreatedTime  int
	// xo fields
	_exists, _deleted bool
}

//dep
func (n *Notification) InsertToDb() {
	res, err := base.DbInsertStruct(n, "notification")
	if err != nil {
		helper.DebugPrintln(res, err)
	}

}

//////////////////////////////////////////////////////
//////////////// Events -Notifiactions ///////////////

//////// Comments //////////
func Notification_OnPostCommented(comment *Comment, post *Post, toAdd bool) {
	objId := post.Id*1000 + ACTION_TYPE_POST_COMMENTED

	not := Notification{
		Id:           0,
		ForUserId:    post.UserId,
		ActorUserId:  comment.UserId,
		ActionTypeId: ACTION_TYPE_POST_COMMENTED,
		ObjectTypeId: OBJECT_COMMENT,
		TargetId:     comment.Id,
		ObjectId:     objId,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}
	if toAdd == false {
		not.ObjectId = -objId
		not.ActionTypeId = -ACTION_TYPE_POST_COMMENTED
	}
	not.InsertToDb()
}

func Notification_OnPostCommentedDelted(comment *Comment, post *Post) {
	q := "delete from notification where ForUserId = ? and ActorUserId = ? and ActionTypeId = ? and TargetId = ?"
	base.DbExecute(q, post.UserId, comment.UserId, ACTION_TYPE_POST_COMMENTED, post.Id)
	Notification_OnPostCommented(comment, post, false)
}

////////// Follows ///////////
func Notification_OnFollowed(UserId, FollowedPeerUserId int) {
	Notification_OnFollowing_Imple(UserId, FollowedPeerUserId, true)
}

func Notification_OnUnFollowed(UserId, FollowedPeerUserId int) {
	Notification_OnFollowing_Imple(UserId, FollowedPeerUserId, false)
}

func Notification_OnFollowing_Imple(UserId, FollowedPeerUserId int, added bool) {
	nf := Notification{
		Id:           0,
		ForUserId:    FollowedPeerUserId,
		ActorUserId:  UserId,
		ActionTypeId: ACTION_TYPE_FOLLOWED_YOU,
		ObjectTypeId: OBJECT_FOLLOWING,
		TargetId:     UserId,
		ObjectId:     0,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}

	if added {
		nf.Save(base.DB)
	} else {
		nf.ActionTypeId = -ACTION_TYPE_FOLLOWED_YOU
        nf.Save(base.DB)
		Notification_Delete(nf)
	}

    Notification_PushToUserPipe(nf)
}

////////////// Likes ///////////////
func Notification_OnPostLiked(lk *Like) {
	Notification_OnPostLikeing_Imple(lk, true)
}

func Notification_OnPostUnLiked(lk *Like) {
	Notification_OnPostLikeing_Imple(lk, false)
}

func Notification_OnPostLikeing_Imple(lk *Like, added bool) {
	post, err := CacheModels.GetPostById(lk.PostId)
	if err != nil {
		return
	}

	nf := Notification{
		Id:           0,
		ForUserId:    post.UserId,
		ActorUserId:  lk.UserId,
		ActionTypeId: ACTION_TYPE_POST_LIKED,
		ObjectTypeId: OBJECT_LIKE,
		TargetId:     post.Id,
		ObjectId:     0,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}

	if added {
		nf.InsertToDb()
	} else {
		nf.ActionTypeId = -nf.ActionTypeId
		nf.InsertToDb()
		Notification_Delete(nf)
	}
}

//////////////////////////////////
func Notification_Delete(nf Notification) {
	aid := int(math.Abs(float64(nf.ActionTypeId))) // alwayse +
    NewNotification_Deleter().ForUserId_EQ(nf.ForUserId).ActorUserId_EQ(nf.ActorUserId).ActionTypeId_EQ(aid).Delete(base.DB)
}

func Notification_PushToUserPipe(nf Notification) {
    call := base.NewCallWithData("Notification",nf)
    AllPipesMap.SendToUser(nf.ForUserId,call)
}

//////////////////////////////////////////////////
type NotificationView struct {
	Notification
	Load interface{}
	//Actor   UserBasicAndMe
}

type NotifPayload struct {
	Actor   *UserBasicAndMe
	Post    *Post
	Comment *Comment
}

func Notification_GetLastsViews(UserId int) []NotificationView {
	q := "select * from notification where ForUserId = ? order by Id desc limit 200 "

	var nots []Notification
	err := base.DB.Select(&nots, q, UserId)
	if err != nil {
		helper.DebugPrintln(err)
	}
	res := make([]NotificationView, 0, len(nots))

	for _, nf := range nots {
		nv := NotificationView{}
		nv.Notification = nf

		load := NotifPayload{}
		nv.Load = &load

		if nf.ActionTypeId > 0 {
			load.Actor = GetUserBasicAndMe(nf.ActorUserId, UserId)

			switch nf.ActionTypeId {
			case ACTION_TYPE_FOLLOWED_YOU:

			case ACTION_TYPE_POST_LIKED:
				post, err := CacheModels.GetPostById(nf.TargetId)
				if err == nil {
					load.Post = post
				} else {
					helper.DebugPrintln(err)
				}

			case ACTION_TYPE_POST_COMMENTED:
				com, err := CacheModels.GetCommentById(nf.TargetId)
				if err == nil {
					load.Comment = com
					post, _ := CacheModels.GetPostById(com.PostId)
					load.Post = post
				} else {
					helper.DebugPrintln(err)
				}
			}

		}
		res = append(res, nv)

	}

	return res

}

