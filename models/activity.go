package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

type Activity struct {
	Id           int
	ActorUserId  int
	ActionTypeId int
	TargetId     int
	RefId        int
	CreatedAt    int

	_exists, _deleted bool
}

func Activity_OnPostCommented(comment *Comment, post *Post) {
	if comment == nil || post == nil {
		return
	}

	refId := comment.Id*1000 + ACTION_TYPE_POST_COMMENTED
	not := Activity{
		Id:           0,
		ActorUserId:  comment.UserId,
		ActionTypeId: ACTION_TYPE_POST_COMMENTED,
		TargetId:     comment.Id,
		RefId:        refId,
		CreatedAt:    helper.TimeNow(),
	}
	not.Save(base.DB)
}

func Activity_OnPostCommentedDeleted(comment *Comment, post *Post) {
	if comment == nil || post == nil {
		return
	}

	refId := comment.Id*1000 + ACTION_TYPE_POST_COMMENTED
	NewActivity_Deleter().
		ActorUserId_EQ(comment.UserId).
		RefId_EQ(refId).
		Delete(base.DB)

}

////////// For Follows ///////////
func Activity_OnFollowed(UserId, FollowedPeerUserId, FLId int) {

	refId := FLId*1000 + ACTION_TYPE_FOLLOWED_USER
	not := Activity{
		ActorUserId:  UserId,
		ActionTypeId: ACTION_TYPE_FOLLOWED_USER,
		TargetId:     FollowedPeerUserId,
		RefId:        refId,
		CreatedAt:    helper.TimeNow(),
	}

	not.Save(base.DB)
}

func Activity_OnUnFollowed(UserId, FollowedPeerUserId, FLId int) {
	refId := FLId*1000 + ACTION_TYPE_FOLLOWED_USER

	NewActivity_Deleter().
		ActorUserId_EQ(UserId).
		RefId_EQ(refId).
		Delete(base.DB)
}

////////////// For Likes ///////////////
func Activity_OnPostLiked(lk *Like) {
	if lk == nil {
		return
	}
	refId := lk.Id*1000 + ACTION_TYPE_POST_LIKED
	not := Activity{
		ActorUserId:  lk.UserId,
		ActionTypeId: ACTION_TYPE_POST_LIKED,
		TargetId:     lk.PostId,
		RefId:        refId,
		CreatedAt:    helper.TimeNow(),
	}

	not.Save(base.DB)
}

func Activity_OnPostUnLiked(lk *Like) {
	if lk == nil {
		return
	}
	refId := lk.Id*1000 + ACTION_TYPE_POST_LIKED

	NewActivity_Deleter().
		ActorUserId_EQ(lk.UserId).
		RefId_EQ(refId).
		Delete(base.DB)
}

////////////////////// Views ////////////////////////////
type ActivityView struct {
	Activity
	Load interface{}
	//Actor   UserBasicAndMe
}

type ActivityPayload struct {
	Actor   *UserBasicAndMe
	Post    *Post
	Comment *Comment
}

//page >= 1
func Activity_GetLastsViews(UserId, Page, Limit, Last int) []ActivityView {
	uids := MemoryStore.UserFollowingList_Get(UserId).Elements

	selector := NewActivity_Selector().ActorUserId_In(uids).OrderBy_Id_Desc().Limit(Limit)
	if Last > 0 {
		selector.Id_LT(Last)
	} else if Page >= 1 {
		selector.Offset((Page - 1) * Limit)
	}

	nots, err := selector.GetRows2(base.DB)

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

		load.Actor = Views.UserBasicAndMeView(UserId, act.ActorUserId)
		switch act.ActionTypeId {
		case ACTION_TYPE_FOLLOWED_USER:
			//no load data

		case ACTION_TYPE_POST_LIKED:
			post, ok := Store.GetPostById(act.TargetId)
			if ok {
				load.Post = post
			} else {
				helper.DebugPrintln(err)
			}

		case ACTION_TYPE_POST_COMMENTED:
			com, ok := Store.GetCommentById(act.TargetId)
			if ok {
				load.Comment = com
				post, _ := Store.GetPostById(com.PostId)
				load.Post = post
			} else {
				helper.DebugPrintln(err)
			}
		}

		res = append(res, av)

	}

	return res

}

func Activity_fillCaches(nots []Activity) {
	//preload start
	var pre_posts []int
	var pre_comments []int

	for _, nf := range nots {
		switch nf.ActionTypeId {
		case ACTION_TYPE_FOLLOWED_USER:

		case ACTION_TYPE_POST_LIKED:
			pre_posts = append(pre_posts, nf.TargetId)

		case ACTION_TYPE_POST_COMMENTED:
			//pre_posts = append(pre_posts, nf.TargetId)
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

	/*if len(pre_posts) >0 {
	    NewPost_Selector().Id_In(pre_posts).GetRows(base.DB)
	}*/

}
