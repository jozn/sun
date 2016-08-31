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
	ObjectId     int // == TargetId * 1000 + ObjectTypeId  ---- eg: 1256*1000 + 3 = 1256003
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
        TargetId: post.Id,
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
	PayLoad interface{}
	Actor   UserBasicAndMe
}
