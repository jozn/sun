// Package models contains the types for schema 'ms'.
package models

// copy of genretaed

import (
    "strconv"
    "time"
)

//Comment Events
func OnComment_Insert(o *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnComment_Update(o *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnComment_Delete(o *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnComment_LoadOne(o *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnComment_LoadMany(o []*Comment) {
    for _, v := range o {
        Cacher.Set("Comment:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//FollowingList Events
func OnFollowingList_Insert(o *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnFollowingList_Update(o *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnFollowingList_Delete(o *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnFollowingList_LoadOne(o *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnFollowingList_LoadMany(o []*FollowingList) {
    for _, v := range o {
        Cacher.Set("FollowingList:"+strconv.Itoa(v.UserId), o, time.Hour*0)
    }
}

//FollowingListMember Events
func OnFollowingListMember_Insert(o *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMember_Update(o *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMember_Delete(o *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMember_LoadOne(o *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMember_LoadMany(o []*FollowingListMember) {
    for _, v := range o {
        Cacher.Set("FollowingListMember:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//FollowingListMemberHistory Events
func OnFollowingListMemberHistory_Insert(o *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMemberHistory_Update(o *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMemberHistory_Delete(o *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMemberHistory_LoadOne(o *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnFollowingListMemberHistory_LoadMany(o []*FollowingListMemberHistory) {
    for _, v := range o {
        Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//Like Events
func OnLike_Insert(o *Like) {
    Cacher.Set("Like:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnLike_Update(o *Like) {
    Cacher.Set("Like:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnLike_Delete(o *Like) {
    Cacher.Set("Like:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnLike_LoadOne(o *Like) {
    Cacher.Set("Like:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnLike_LoadMany(o []*Like) {
    for _, v := range o {
        Cacher.Set("Like:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//Media Events
func OnMedia_Insert(o *Media) {
    Cacher.Set("Media:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMedia_Update(o *Media) {
    Cacher.Set("Media:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMedia_Delete(o *Media) {
    Cacher.Set("Media:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMedia_LoadOne(o *Media) {
    Cacher.Set("Media:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMedia_LoadMany(o []*Media) {
    for _, v := range o {
        Cacher.Set("Media:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//Message Events
func OnMessage_Insert(o *Message) {
    Cacher.Set("Message:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMessage_Update(o *Message) {
    Cacher.Set("Message:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMessage_Delete(o *Message) {
    Cacher.Set("Message:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMessage_LoadOne(o *Message) {
    Cacher.Set("Message:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMessage_LoadMany(o []*Message) {
    for _, v := range o {
        Cacher.Set("Message:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//MsgDeletedFromServer Events
func OnMsgDeletedFromServer_Insert(o *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgDeletedFromServer_Update(o *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgDeletedFromServer_Delete(o *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgDeletedFromServer_LoadOne(o *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgDeletedFromServer_LoadMany(o []*MsgDeletedFromServer) {
    for _, v := range o {
        Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//MsgReceivedToPeer Events
func OnMsgReceivedToPeer_Insert(o *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgReceivedToPeer_Update(o *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgReceivedToPeer_Delete(o *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgReceivedToPeer_LoadOne(o *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgReceivedToPeer_LoadMany(o []*MsgReceivedToPeer) {
    for _, v := range o {
        Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//MsgSeenByPeer Events
func OnMsgSeenByPeer_Insert(o *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgSeenByPeer_Update(o *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgSeenByPeer_Delete(o *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgSeenByPeer_LoadOne(o *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnMsgSeenByPeer_LoadMany(o []*MsgSeenByPeer) {
    for _, v := range o {
        Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//Notification Events
func OnNotification_Insert(o *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnNotification_Update(o *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnNotification_Delete(o *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnNotification_LoadOne(o *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnNotification_LoadMany(o []*Notification) {
    for _, v := range o {
        Cacher.Set("Notification:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//NotificationRemoved Events
func OnNotificationRemoved_Insert(o *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(o.NotificationId), o, time.Hour*0)
}

func OnNotificationRemoved_Update(o *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(o.NotificationId), o, time.Hour*0)
}

func OnNotificationRemoved_Delete(o *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(o.NotificationId), o, time.Hour*0)
}

func OnNotificationRemoved_LoadOne(o *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(o.NotificationId), o, time.Hour*0)
}

func OnNotificationRemoved_LoadMany(o []*NotificationRemoved) {
    for _, v := range o {
        Cacher.Set("NotificationRemoved:"+strconv.Itoa(v.NotificationId), o, time.Hour*0)
    }
}

//PhoneContact Events
func OnPhoneContact_Insert(o *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPhoneContact_Update(o *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPhoneContact_Delete(o *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPhoneContact_LoadOne(o *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPhoneContact_LoadMany(o []*PhoneContact) {
    for _, v := range o {
        Cacher.Set("PhoneContact:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//Post Events
func OnPost_Insert(o *Post) {
    Cacher.Set("Post:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPost_Update(o *Post) {
    Cacher.Set("Post:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPost_Delete(o *Post) {
    Cacher.Set("Post:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPost_LoadOne(o *Post) {
    Cacher.Set("Post:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnPost_LoadMany(o []*Post) {
    for _, v := range o {
        Cacher.Set("Post:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//RecommendUser Events
func OnRecommendUser_Insert(o *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnRecommendUser_Update(o *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnRecommendUser_Delete(o *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnRecommendUser_LoadOne(o *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnRecommendUser_LoadMany(o []*RecommendUser) {
    for _, v := range o {
        Cacher.Set("RecommendUser:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//SearchClicked Events
func OnSearchClicked_Insert(o *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnSearchClicked_Update(o *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnSearchClicked_Delete(o *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnSearchClicked_LoadOne(o *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnSearchClicked_LoadMany(o []*SearchClicked) {
    for _, v := range o {
        Cacher.Set("SearchClicked:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//Session Events
func OnSession_Insert(o *Session) {
    Cacher.Set("Session:"+strconv.Itoa(o.SessionUuid), o, time.Hour*0)
}

func OnSession_Update(o *Session) {
    Cacher.Set("Session:"+strconv.Itoa(o.SessionUuid), o, time.Hour*0)
}

func OnSession_Delete(o *Session) {
    Cacher.Set("Session:"+strconv.Itoa(o.SessionUuid), o, time.Hour*0)
}

func OnSession_LoadOne(o *Session) {
    Cacher.Set("Session:"+strconv.Itoa(o.SessionUuid), o, time.Hour*0)
}

func OnSession_LoadMany(o []*Session) {
    for _, v := range o {
        Cacher.Set("Session:"+strconv.Itoa(v.SessionUuid), o, time.Hour*0)
    }
}

//Tag Events
func OnTag_Insert(o *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTag_Update(o *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTag_Delete(o *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTag_LoadOne(o *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTag_LoadMany(o []*Tag) {
    for _, v := range o {
        Cacher.Set("Tag:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//TagsPost Events
func OnTagsPost_Insert(o *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTagsPost_Update(o *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTagsPost_Delete(o *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTagsPost_LoadOne(o *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnTagsPost_LoadMany(o []*TagsPost) {
    for _, v := range o {
        Cacher.Set("TagsPost:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//User Events
func OnUser_Insert(o *User) {
    Cacher.Set("User:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnUser_Update(o *User) {
    Cacher.Set("User:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnUser_Delete(o *User) {
    Cacher.Set("User:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnUser_LoadOne(o *User) {
    Cacher.Set("User:"+strconv.Itoa(o.Id), o, time.Hour*0)
}

func OnUser_LoadMany(o []*User) {
    for _, v := range o {
        Cacher.Set("User:"+strconv.Itoa(v.Id), o, time.Hour*0)
    }
}

//UserMetaInfo Events
func OnUserMetaInfo_Insert(o *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserMetaInfo_Update(o *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserMetaInfo_Delete(o *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserMetaInfo_LoadOne(o *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserMetaInfo_LoadMany(o []*UserMetaInfo) {
    for _, v := range o {
        Cacher.Set("UserMetaInfo:"+strconv.Itoa(v.UserId), o, time.Hour*0)
    }
}

//UserPassword Events
func OnUserPassword_Insert(o *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserPassword_Update(o *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserPassword_Delete(o *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserPassword_LoadOne(o *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(o.UserId), o, time.Hour*0)
}

func OnUserPassword_LoadMany(o []*UserPassword) {
    for _, v := range o {
        Cacher.Set("UserPassword:"+strconv.Itoa(v.UserId), o, time.Hour*0)
    }
}
