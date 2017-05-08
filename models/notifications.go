package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

//////// For Comments //////////
func Notification_OnPostCommented(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	if comment.UserId == post.UserId {
		return
	}

	refId := Ref_CommentAdd(comment.Id) //post.Id*1000 + ACTION_TYPE_POST_COMMENTED
	not := x.Notification{
		Id:           0,
		ForUserId:    post.UserId,
		ActorUserId:  comment.UserId,
		ActionTypeId: ACTION_TYPE_POST_COMMENTED,
		ObjectTypeId: OBJECT_COMMENT,
		RowId:        comment.Id,
		RootId:       comment.PostId,
		RefId:        refId,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}
	not.Save(base.DB)

	Notification_PushToUserPipe(not)
}

func Notification_OnPostCommentedDeleted(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	row, err := x.NewNotification_Selector().
		ForUserId_Eq(post.UserId).
		ActorUserId_Eq(comment.UserId).
		RowId_Eq(comment.Id).
		ActionTypeId_Eq(ACTION_TYPE_POST_COMMENTED).
		GetRow(base.DB)

	if err == nil {
		nr := x.NotificationRemoved{
			NotificationId: row.Id,
			ForUserId:      comment.UserId,
		}

		row.Delete(base.DB)
		nr.Save(base.DB)

		Notification_PushToUserPipeRemoved(row.Id)
	}
}

////////// For Follows ///////////
func Notification_OnFollowed(UserId, FollowedPeerUserId, FLId int) {
	if UserId == FollowedPeerUserId { //must never reach here at all
		return
	}

	ref := Ref_FollowAdd(FLId)
	nf := x.Notification{
		Id:           0,
		ForUserId:    FollowedPeerUserId,
		ActorUserId:  UserId,
		ActionTypeId: ACTION_TYPE_FOLLOWED_USER,
		ObjectTypeId: OBJECT_FOLLOWING,
		RowId:        FLId,
		RootId:       FollowedPeerUserId,
		RefId:        ref,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}

	nf.Save(base.DB)

	Notification_PushToUserPipe(nf)
}

func Notification_OnUnFollowed(UserId, FollowedPeerUserId int) {
	row, err := x.NewNotification_Selector().
		ForUserId_Eq(FollowedPeerUserId).
		ActorUserId_Eq(UserId).
		ActionTypeId_Eq(ACTION_TYPE_FOLLOWED_USER).
		GetRow(base.DB)

	if err == nil {
		nr := x.NotificationRemoved{
			NotificationId: row.Id,
			ForUserId:      UserId,
		}

		row.Delete(base.DB)
		nr.Save(base.DB)

		Notification_PushToUserPipeRemoved(row.Id)
	}
}

////////////// For Likes ///////////////
func Notification_OnPostLiked(lk *x.Like) {
	post, ok := x.Store.GetPostById(lk.PostId)
	if !ok {
		return
	}

	if lk.UserId == post.UserId {
		return
	}

	refId := Ref_LikeAdd(lk.Id)
	nf := x.Notification{
		Id:           0,
		ForUserId:    post.UserId,
		ActorUserId:  lk.UserId,
		ActionTypeId: ACTION_TYPE_POST_LIKED,
		ObjectTypeId: OBJECT_LIKE,
		RowId:        post.Id,
		RootId:       lk.PostId,
		RefId:        refId,
		SeenStatus:   0,
		CreatedTime:  helper.TimeNow(),
	}

	nf.Save(base.DB)
}

func Notification_OnPostUnLiked(lk *x.Like) {
	post, ok := x.Store.GetPostById(lk.PostId)
	if !ok {
		return
	}

	row, err := x.NewNotification_Selector().
		ForUserId_Eq(post.UserId).
		ActorUserId_Eq(lk.UserId).
		RootId_Eq(lk.PostId).
		ActionTypeId_Eq(ACTION_TYPE_POST_LIKED).
		GetRow(base.DB)

	if err == nil {
		nr := x.NotificationRemoved{
			NotificationId: row.Id,
			ForUserId:      post.UserId,
		}

		row.Delete(base.DB)
		nr.Save(base.DB)
	}
}

//fix: must be NotificationView
func Notification_PushToUserPipe(nf x.Notification) {
	nv := Notification_notifyToView(&nf, nf.ForUserId)
	call := base.NewCallWithData("NotifyAddOne", nv)
	AllPipesMap.SendToUser(nf.ForUserId, call)
}

func Notification_PushToUserPipeRemoved(id int) {

}

func Notification_ListOfRemovedAndEmptyIt(UserId int) []int {
	res, _ := x.NewNotificationRemoved_Selector().Select_NotificationId().ForUserId_Eq(UserId).GetIntSlice(base.DB)
	if res != nil && len(res) > 0 {
		x.NewNotificationRemoved_Deleter().ForUserId_Eq(UserId).Delete(base.DB)
	}
	return res
}

////////////////////// Views ////////////////////////////

func Notification_GetLastsViews(UserId, last int) (res []NotificationView) {
	selector := x.NewNotification_Selector().ForUserId_Eq(UserId).Limit(100).OrderBy_Id_Desc()
	if last > 0 {
		selector.Id_GT(last)
	}

	nots, err := selector.GetRows(base.DB)

	if err != nil {
		helper.DebugPrintln(err)
		return
	}

	res = make([]NotificationView, 0, len(nots))

	Notification_fillCaches(nots)

	for _, nf := range nots {
		res = append(res, Notification_notifyToView(nf, UserId))
	}

	return res
}

func Notification_notifyToView(nf *x.Notification, UserId int) NotificationView {
	nv := NotificationView{}
	nv.Notification = nf

	load := NotifPayload{}
	nv.Load = &load

	if nf.ActionTypeId > 0 { //old check not need anymore (it was for when ActionTypedId could be negative)
		load.Actor, _ = Views.GetUserInlineWithMeView(UserId,nf.ActorUserId)

		switch nf.ActionTypeId {

		case ACTION_TYPE_FOLLOWED_USER:
			load.Followed, _ = Views.GetUserInlineWithMeView(UserId, nf.RootId)
		case ACTION_TYPE_POST_LIKED:
			//post, ok := x.Store.GetPostById(nf.RowId)
			post, ok := Views.PostSingleViewForPostId(nf.RootId, UserId)
			if ok {
				load.Post = post
			}

		case ACTION_TYPE_POST_COMMENTED:
			com, ok := x.Store.GetCommentById(nf.RowId)
			if ok {
				load.Comment = com
				post, _ := Views.PostSingleViewForPostId(com.PostId, UserId) //x.Store.GetPostById(com.PostId)
				load.Post = post
			}
		}

	}
	return nv
}

//copy of Activity_fillCaches with modificaion - make in sync
func Notification_fillCaches(nots []*x.Notification) {
	//preload start
	var pre_posts []int
	var pre_comments []int

	for _, nf := range nots {
		switch nf.ActionTypeId {
		case ACTION_TYPE_FOLLOWED_USER:

		case ACTION_TYPE_POST_LIKED:
			pre_posts = append(pre_posts, nf.RootId)

		case ACTION_TYPE_POST_COMMENTED:
			pre_comments = append(pre_comments, nf.RowId)
            pre_posts = append(pre_posts, nf.RootId)

        }
	}

	x.Store.PreLoadCommentByIds(pre_comments)

	/*for _, commentId := range pre_comments {
		com, ok := x.Store.GetCommentById(commentId)
		if ok {
			pre_posts = append(pre_posts, com.PostId)
		}
	}*/

	x.Store.PreLoadPostByIds(pre_posts)
}
