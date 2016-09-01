package models

import (
    "ms/sun/base"
    "ms/sun/helper"
    "math"
)

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
}

func (n *Notification) InsertToDb() {
	res ,err := base.DbInsertStruct(n, "notification")
    if err !=nil {
        helper.DebugPrintln(res, err)
    }
}

func (n *Notification) Delete() {
}


func LoadLastNotificationsForUser(UserId, FromTime int) []Notification {
	var res []Notification
	q := "select * from notification where UserId = ? And CreatedTime > ? order by CreatedTime desc limit 200"
	base.DB.Select(res, q, UserId, FromTime)
	return res
}

////////////////////////////////////////
//////////////// Events -Notifiactions ///////////////

//////// Comments //////////
func Notification_OnPostCommented(comment *Comment, post *Post, toAdd bool) {
    objId := post.Id*1000+ACTION_TYPE_POST_COMMENTED

    not := Notification{
        Id:0,
        ForUserId: post.UserId,
        ActorUserId: comment.UserId,
        ActionTypeId: ACTION_TYPE_POST_COMMENTED,
        ObjectTypeId: OBJECT_COMMENT,
        TargetId: comment.Id,
        ObjectId: objId,
        SeenStatus: 0,
        CreatedTime: helper.TimeNow(),
    }
    if toAdd == false {
        not.ObjectId = -objId
        not.ActionTypeId = -ACTION_TYPE_POST_COMMENTED
    }
    not.InsertToDb()
}

func Notification_OnPostCommentedDelted(comment *Comment, post *Post) {
    q:="delete from notification where ForUserId = ? and ActorUserId = ? and ActionTypeId = ? and TargetId = ?"
    base.DbExecute(q, post.UserId, comment.UserId , ACTION_TYPE_POST_COMMENTED ,post.Id)
    Notification_OnPostCommented(comment , post ,false )
}

////////// Follows ///////////
func Notification_OnFollowed(UserId, FollowedPeerUserId int) {
    Notification_OnFollowing_Imple(UserId, FollowedPeerUserId,true)
}

func Notification_OnUnFollowed(UserId, FollowedPeerUserId int) {
    Notification_OnFollowing_Imple(UserId, FollowedPeerUserId,false)
}

func Notification_OnFollowing_Imple(UserId, FollowedPeerUserId int , added bool) {
    nf := Notification{
        Id:0,
        ForUserId: FollowedPeerUserId,
        ActorUserId: UserId,
        ActionTypeId: ACTION_TYPE_FOLLOWED_YOU,
        ObjectTypeId: OBJECT_FOLLOWING,
        TargetId: UserId,
        ObjectId: 0,
        SeenStatus: 0,
        CreatedTime: helper.TimeNow(),
    }

    if added {
        nf.InsertToDb()
    }else{
        nf.ActionTypeId = -ACTION_TYPE_FOLLOWED_YOU
        nf.InsertToDb()
        Notification_Delete(nf)
    }
}

///////// Likes /////////////
func Notification_OnPostLiked(lk *Like) {
    Notification_OnPostLikeing_Imple(lk,true)
}

func Notification_OnPostUnLiked(lk *Like) {
    Notification_OnPostLikeing_Imple(lk,false)
}

func Notification_OnPostLikeing_Imple(lk *Like , added bool) {
    post,err := CacheModels.GetPostById(lk.PostId)
    if err != nil {
        return
    }

    nf := Notification{
        Id:0,
        ForUserId: post.UserId,
        ActorUserId: lk.UserId,
        ActionTypeId: ACTION_TYPE_POST_LIKED,
        ObjectTypeId: OBJECT_LIKE,
        TargetId: post.Id,
        ObjectId: 0,
        SeenStatus: 0,
        CreatedTime: helper.TimeNow(),
    }

    if added {
        nf.InsertToDb()
    }else{
        nf.ActionTypeId = - nf.ActionTypeId
        nf.InsertToDb()
        Notification_Delete(nf)
    }
}

//////////////////////////////////
func Notification_Delete(nf Notification) {
    q:="delete from notification where ForUserId = ? and ActorUserId = ? and ActionTypeId = ? and TargetId = ?"
    aid := math.Abs(float64(nf.ActionTypeId)) // alyase +
    base.DbExecute(q, nf.ForUserId, nf.ActorUserId , aid , nf.TargetId)
    //base.DbExecute(q, post.UserId, comment.UserId , ACTION_TYPE_POST_COMMENTED ,post.Id)
}

//////////////////////////////////////////////////
type NotificationView struct {
    Notification
    Load interface{}
	//Actor   UserBasicAndMe
}

type NotifPayload struct  {
    Actor   *UserBasicAndMe
    Post *Post
    Comment *Comment
}

func Notification_GetLastsViews(UserId int) ([]NotificationView) {
    q:= "select * from notification where ForUserId = ? order by Id desc limit 200 "

    var nots []Notification
    err := base.DB.Select(&nots, q, UserId)
    if err != nil {
        helper.DebugPrintln(err)
    }
    res := make([]NotificationView,0, len(nots))

    for _, nf := range nots {
        nv :=NotificationView{}
        nv.Notification = nf

        load := NotifPayload{}
        nv.Load = &load

        if (nf.ActionTypeId > 0 ){
            load.Actor = GetUserBasicAndMe(nf.ActorUserId ,UserId)

            switch nf.ActionTypeId {
            case ACTION_TYPE_FOLLOWED_YOU:

            case ACTION_TYPE_POST_LIKED:
                post,err := CacheModels.GetPostById(nf.TargetId)
                if err == nil {
                    load.Post = post
                }else {
                    helper.DebugPrintln(err)
                }

            case ACTION_TYPE_POST_COMMENTED:
                com,err := CacheModels.GetCommentById(nf.TargetId)
                if err == nil {
                    load.Comment = com
                    post,_ := CacheModels.GetPostById(com.PostId)
                    load.Post = post
                }else {
                    helper.DebugPrintln(err)
                }
            }

        }
        res = append(res,nv)

    }

    return res

}

/*

func Notification_GetLastsViews_BK(UserId int) ([]NotificationView) {
    q:= "select * from notification where ForUserId = ? order by Id desc limit 200 "

    var nots []Notification
    err := base.DB.Select(&nots, q, UserId)
    if err != nil {
        helper.DebugPrintln(err)
    }
    res := make([]NotificationView,0, len(nots))

    for _, nf := range nots {
        nv :=NotificationView{}
        nv.Notification = nf

        if (nf.ActionTypeId > 0 ){
            nv.Actor = *GetUserBasicAndMe(nf.ActorUserId ,UserId)

            switch nf.ActionTypeId {
            case ACTION_TYPE_FOLLOWED_YOU:
                nv.PayLoad = nv.Actor
            case ACTION_TYPE_POST_LIKED:
                post,err := CacheModels.GetPostById(nf.TargetId)
                helper.DebugPrintln(err)
                if err == nil {
                    nv.PayLoad = *post
                }else {
                    helper.DebugPrintln(err)
                }

            case ACTION_TYPE_POST_COMMENTED:
                com,err := CacheModels.GetCommentById(nf.TargetId)
                if err == nil {
                    nv.PayLoad = *com
                }else {
                    helper.DebugPrintln(err)
                }
            }

        }
        res = append(res,nv)

    }

    return res

}
*/
