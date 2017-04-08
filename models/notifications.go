package models

import (
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

type NotificationRemoved struct {
	NotificationId int
	ForUserId      int

	_exists, _deleted bool
}

//////// For Comments //////////
func Notification_OnPostCommented(comment *Comment, post *Post) {
	if comment == nil || post == nil {
		return
	}

	if comment.UserId == post.UserId {
		return
	}

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
	not.Save(base.DB)

	Notification_PushToUserPipe(not)
}

func Notification_OnPostCommentedDeleted(comment *Comment, post *Post) {
	if comment == nil || post == nil {
		return
	}

	row, err := NewNotification_Selector().
		ForUserId_EQ(post.UserId).
		ActorUserId_EQ(comment.UserId).
		TargetId_EQ(comment.Id).
		ActionTypeId_EQ(ACTION_TYPE_POST_COMMENTED).
		GetRow(base.DB)

	if err == nil {
		nr := NotificationRemoved{
			NotificationId: row.Id,
			ForUserId:      comment.UserId,
		}

		row.Delete(base.DB)
		nr.Save(base.DB)

		Notification_PushToUserPipeRemoved(row.Id)
	}
}

////////// For Follows ///////////
func Notification_OnFollowed(UserId, FollowedPeerUserId int) {
	if UserId == FollowedPeerUserId { //must never reach here at all
		return
	}
	nf := Notification{
		Id:           0,
		ForUserId:    FollowedPeerUserId,
		ActorUserId:  UserId,
		ActionTypeId: ACTION_TYPE_FOLLOWED_USER,
		ObjectTypeId: OBJECT_FOLLOWING,
		TargetId:     UserId,
		ObjectId:     0,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}

	nf.Save(base.DB)

	Notification_PushToUserPipe(nf)
}

func Notification_OnUnFollowed(UserId, FollowedPeerUserId int) {
	row, err := NewNotification_Selector().
		ForUserId_EQ(FollowedPeerUserId).
		ActorUserId_EQ(UserId).
		ActionTypeId_EQ(ACTION_TYPE_FOLLOWED_USER).
		GetRow(base.DB)

	if err == nil {
		nr := NotificationRemoved{
			NotificationId: row.Id,
			ForUserId:      UserId,
		}

		row.Delete(base.DB)
		nr.Save(base.DB)

		Notification_PushToUserPipeRemoved(row.Id)
	}
}

////////////// For Likes ///////////////
func Notification_OnPostLiked(lk *Like) {
	post, err := CacheModels.GetPostById(lk.PostId)
	if err != nil {
		return
	}

	if lk.UserId == post.UserId {
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

	nf.Save(base.DB)
}

func Notification_OnPostUnLiked(lk *Like) {
	post, err := CacheModels.GetPostById(lk.PostId)
	if err != nil {
		return
	}

	row, err := NewNotification_Selector().
		ForUserId_EQ(post.UserId).
		ActorUserId_EQ(lk.UserId).
		TargetId_EQ(lk.PostId).
		ActionTypeId_EQ(ACTION_TYPE_POST_LIKED).
		GetRow(base.DB)

	if err == nil {
		nr := NotificationRemoved{
			NotificationId: row.Id,
			ForUserId:      post.UserId,
		}

		row.Delete(base.DB)
		nr.Save(base.DB)
	}
}

//fix: must be NotificationView
func Notification_PushToUserPipe(nf Notification) {
	nv := Notification_notifyToView(&nf, nf.ForUserId)
	call := base.NewCallWithData("NotifyAddOne", nv)
	AllPipesMap.SendToUser(nf.ForUserId, call)
}

func Notification_PushToUserPipeRemoved(id int) {

}

func Notification_ListOfRemovedAndEmptyIt(UserId int) []int {
	res, _ := NewNotificationRemoved_Selector().Select_NotificationId().ForUserId_EQ(UserId).GetIntSlice(base.DB)
	if res != nil && len(res) > 0 {
		NewNotificationRemoved_Deleter().ForUserId_EQ(UserId).Delete(base.DB)
	}
	return res
}

////////////////////// Views ////////////////////////////
type NotificationView struct {
	*Notification
	Load interface{}
	//Actor   UserBasicAndMe
}

type NotifPayload struct {
	Actor   *UserBasicAndMe
	Post    *Post
	Comment *Comment
}

func Notification_GetLastsViews(UserId, last int) []NotificationView {
	selector := NewNotification_Selector().ForUserId_EQ(UserId).Limit(100).OrderBy_Id_Desc()
	if last > 0 {
		selector.Id_GT(last)
	}

	nots, err := selector.GetRows(base.DB)

	if err != nil {
		helper.DebugPrintln(err)
	}

	res := make([]NotificationView, 0, len(nots))

	Notification_fillCaches(nots)

	for _, nf := range nots {
		/*nv := NotificationView{}
		nv.Notification = nf

		load := NotifPayload{}
		nv.Load = &load

		if nf.ActionTypeId > 0 { //old check not need anymore (it was for when ActionTypedId could be negative)
			load.Actor = GetUserBasicAndMe(nf.ActorUserId, UserId)

			switch nf.ActionTypeId {
			case ACTION_TYPE_FOLLOWED_USER:

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

		}*/
		res = append(res, Notification_notifyToView(nf, UserId))

	}

	return res

}

func Notification_notifyToView(nf *Notification, UserId int) NotificationView {
	nv := NotificationView{}
	nv.Notification = nf

	load := NotifPayload{}
	nv.Load = &load

	if nf.ActionTypeId > 0 { //old check not need anymore (it was for when ActionTypedId could be negative)
		load.Actor = GetUserBasicAndMe(nf.ActorUserId, UserId)

		switch nf.ActionTypeId {
		case ACTION_TYPE_FOLLOWED_USER:

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
	return nv
}

//copy of Activity_fillCaches with modificaion - make in sync
func Notification_fillCaches(nots []*Notification) {
	//preload start
	var pre_posts []int
	var pre_comments []int

	for _, nf := range nots {
		switch nf.ActionTypeId {
		case ACTION_TYPE_FOLLOWED_USER:

		case ACTION_TYPE_POST_LIKED:
			pre_posts = append(pre_posts, nf.TargetId)

		case ACTION_TYPE_POST_COMMENTED:
			pre_comments = append(pre_comments, nf.TargetId)
		}
	}

	Store.PreLoadCommentByIds(pre_comments)

	for _, commentId := range pre_comments {
		com, ok := Store.GetCommentById(commentId)
		if ok {
			pre_posts = append(pre_posts, com.PostId)
		}
	}

	Store.PreLoadPostByIds(pre_posts)
}
