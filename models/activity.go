package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

func Activity_OnPostCommentAdd(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	refId := Ref_CommentAdd(comment.Id) //comment.Id*1000 + ACTION_TYPE_POST_COMMENTED
	not := x.Activity{
		Id:           0,
		ActorUserId:  comment.UserId,
		ActionTypeId: ACTION_TYPE_POST_COMMENTED,
		RowId:        comment.Id,
		RootId:       post.Id,
		RefId:        refId,
		CreatedAt:    helper.TimeNow(),
	}
	not.Save(base.DB)
}

func Activity_OnPostCommentDeleted(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	refId := Ref_CommentAdd(comment.Id) //comment.Id*1000 + ACTION_TYPE_POST_COMMENTED
	x.NewActivity_Deleter().
		ActorUserId_Eq(comment.UserId).
		RefId_Eq(refId).
		Delete(base.DB)

}

////////// For Follows ///////////
func Activity_OnFollowed(UserId, FollowedPeerUserId, FLId int) {

	refId := Ref_FollowAdd(FLId) //FLId*1000 + ACTION_TYPE_FOLLOWED_USER
	not := x.Activity{
		ActorUserId:  UserId,
		ActionTypeId: ACTION_TYPE_FOLLOWED_USER,
		RowId:        FLId,
		RootId:       FollowedPeerUserId,
		RefId:        refId,
		CreatedAt:    helper.TimeNow(),
	}

	not.Save(base.DB)
}

func Activity_OnUnFollowed(UserId, FollowedPeerUserId, FLId int) {
	refId := Ref_FollowAdd(FLId) //FLId*1000 + ACTION_TYPE_FOLLOWED_USER

	x.NewActivity_Deleter().
		ActorUserId_Eq(UserId).
		RefId_Eq(refId).
		Delete(base.DB)
}

////////////// For Likes ///////////////
func Activity_OnPostLiked(lk *x.Like) {
	if lk == nil {
		return
	}
	refId := Ref_LikeAdd(lk.Id) //lk.Id*1000 + ACTION_TYPE_POST_LIKED
	not := x.Activity{
		ActorUserId:  lk.UserId,
		ActionTypeId: ACTION_TYPE_POST_LIKED,
		RowId:        lk.Id,
		RootId:       lk.PostId,
		RefId:        refId,
		CreatedAt:    helper.TimeNow(),
	}

	not.Save(base.DB)
}

func Activity_OnPostUnLiked(lk *x.Like) {
	if lk == nil {
		return
	}
	refId := Ref_LikeAdd(lk.Id) //lk.Id*1000 + ACTION_TYPE_POST_LIKED

	x.NewActivity_Deleter().
		ActorUserId_Eq(lk.UserId).
		RefId_Eq(refId).
		Delete(base.DB)
}

////////////////////// Views ////////////////////////////

//page >= 1
func Activity_GetLastsViews(UserId, Page, Limit, Last int) []ActivityView {
	uids := MemoryStore.UserFollowingList_Get(UserId).Values()

	selector := x.NewActivity_Selector().ActorUserId_In(uids).OrderBy_Id_Desc().Limit(Limit)
	if Last > 0 {
		selector.Id_LT(Last)
	} else if Page >= 1 {
		selector.Offset((Page - 1) * Limit)
	}

	nots, err := selector.GetRows(base.DB)

	res := make([]ActivityView, 0, len(nots))
	if err != nil {
		return res
	}

	//fill caches
	Activity_fillCaches(nots)

	for _, act := range nots {
		av := ActivityView{}
		av.Activity = act

		load := ActivityPayload{}
		av.Load = &load

		load.Actor, _ = Views.GetUserInlineWithMeView(UserId, act.ActorUserId)
		switch act.ActionTypeId {
		case ACTION_TYPE_FOLLOWED_USER:
			load.Followed, _ = Views.GetUserInlineWithMeView(UserId, act.RootId)

		case ACTION_TYPE_POST_LIKED:
			post, ok := x.Store.GetPostById(act.RootId)
			if ok {
				load.Post = Views.PostSingleView(post, UserId)
			} else {
				helper.DebugPrintln(err)
			}

		case ACTION_TYPE_POST_COMMENTED:
			com, ok := x.Store.GetCommentById(act.RowId)
			if ok {
				load.Comment = com

				post, ok := x.Store.GetPostById(com.PostId)
				if ok {
					load.Post = Views.PostSingleView(post, UserId)
				}
			} else {
				helper.DebugPrintln(err)
			}
		}

		res = append(res, av)
	}

	return res

}

func Activity_fillCaches(nots []*x.Activity) {
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

	//This is now unneccsssory we can use RootId
	/*for _, commentId := range pre_comments {
		com, ok := x.Store.GetCommentById(commentId)
		if ok {
			pre_posts = append(pre_posts, com.PostId)
		}
	}*/

	x.Store.PreLoadPostByIds(pre_posts)

	/*if len(pre_posts) >0 {
	    NewPost_Selector().Id_In(pre_posts).GetRows(base.DB)
	}*/

}
