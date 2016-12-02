// Package models contains the types for schema 'ms'.
package models

import (
    "strconv"
    "time"
)

//Comment Events
func OnComment_AfterInsert(row *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnComment_AfterUpdate(row *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnComment_AfterDelete(row *Comment) {
    Cacher.Delete("Comment:" + strconv.Itoa(row.Id))
}

func OnComment_LoadOne(row *Comment) {
    Cacher.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnComment_LoadMany(rows []*Comment) {
    for _, row := range rows {
        Cacher.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//FollowingList Events
func OnFollowingList_AfterInsert(row *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnFollowingList_AfterUpdate(row *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnFollowingList_AfterDelete(row *FollowingList) {
    Cacher.Delete("FollowingList:" + strconv.Itoa(row.UserId))
}

func OnFollowingList_LoadOne(row *FollowingList) {
    Cacher.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnFollowingList_LoadMany(rows []*FollowingList) {
    for _, row := range rows {
        Cacher.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
    }
}

//FollowingListMember Events
func OnFollowingListMember_AfterInsert(row *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_AfterUpdate(row *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_AfterDelete(row *FollowingListMember) {
    Cacher.Delete("FollowingListMember:" + strconv.Itoa(row.Id))
}

func OnFollowingListMember_LoadOne(row *FollowingListMember) {
    Cacher.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_LoadMany(rows []*FollowingListMember) {
    for _, row := range rows {
        Cacher.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//FollowingListMemberHistory Events
func OnFollowingListMemberHistory_AfterInsert(row *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_AfterUpdate(row *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_AfterDelete(row *FollowingListMemberHistory) {
    Cacher.Delete("FollowingListMemberHistory:" + strconv.Itoa(row.Id))
}

func OnFollowingListMemberHistory_LoadOne(row *FollowingListMemberHistory) {
    Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_LoadMany(rows []*FollowingListMemberHistory) {
    for _, row := range rows {
        Cacher.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Like Events
func OnLike_AfterInsert(row *Like) {
    Cacher.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_AfterUpdate(row *Like) {
    Cacher.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_AfterDelete(row *Like) {
    Cacher.Delete("Like:" + strconv.Itoa(row.Id))
}

func OnLike_LoadOne(row *Like) {
    Cacher.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_LoadMany(rows []*Like) {
    for _, row := range rows {
        Cacher.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Media Events
func OnMedia_AfterInsert(row *Media) {
    Cacher.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMedia_AfterUpdate(row *Media) {
    Cacher.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMedia_AfterDelete(row *Media) {
    Cacher.Delete("Media:" + strconv.Itoa(row.Id))
}

func OnMedia_LoadOne(row *Media) {
    Cacher.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMedia_LoadMany(rows []*Media) {
    for _, row := range rows {
        Cacher.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Message Events
func OnMessage_AfterInsert(row *Message) {
    Cacher.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMessage_AfterUpdate(row *Message) {
    Cacher.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMessage_AfterDelete(row *Message) {
    Cacher.Delete("Message:" + strconv.Itoa(row.Id))
}

func OnMessage_LoadOne(row *Message) {
    Cacher.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMessage_LoadMany(rows []*Message) {
    for _, row := range rows {
        Cacher.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//MsgDeletedFromServer Events
func OnMsgDeletedFromServer_AfterInsert(row *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgDeletedFromServer_AfterUpdate(row *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgDeletedFromServer_AfterDelete(row *MsgDeletedFromServer) {
    Cacher.Delete("MsgDeletedFromServer:" + strconv.Itoa(row.Id))
}

func OnMsgDeletedFromServer_LoadOne(row *MsgDeletedFromServer) {
    Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgDeletedFromServer_LoadMany(rows []*MsgDeletedFromServer) {
    for _, row := range rows {
        Cacher.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//MsgReceivedToPeer Events
func OnMsgReceivedToPeer_AfterInsert(row *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgReceivedToPeer_AfterUpdate(row *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgReceivedToPeer_AfterDelete(row *MsgReceivedToPeer) {
    Cacher.Delete("MsgReceivedToPeer:" + strconv.Itoa(row.Id))
}

func OnMsgReceivedToPeer_LoadOne(row *MsgReceivedToPeer) {
    Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgReceivedToPeer_LoadMany(rows []*MsgReceivedToPeer) {
    for _, row := range rows {
        Cacher.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//MsgSeenByPeer Events
func OnMsgSeenByPeer_AfterInsert(row *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgSeenByPeer_AfterUpdate(row *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgSeenByPeer_AfterDelete(row *MsgSeenByPeer) {
    Cacher.Delete("MsgSeenByPeer:" + strconv.Itoa(row.Id))
}

func OnMsgSeenByPeer_LoadOne(row *MsgSeenByPeer) {
    Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgSeenByPeer_LoadMany(rows []*MsgSeenByPeer) {
    for _, row := range rows {
        Cacher.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Notification Events
func OnNotification_AfterInsert(row *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotification_AfterUpdate(row *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotification_AfterDelete(row *Notification) {
    Cacher.Delete("Notification:" + strconv.Itoa(row.Id))
}

func OnNotification_LoadOne(row *Notification) {
    Cacher.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotification_LoadMany(rows []*Notification) {
    for _, row := range rows {
        Cacher.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//NotificationRemoved Events
func OnNotificationRemoved_AfterInsert(row *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
}

func OnNotificationRemoved_AfterUpdate(row *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
}

func OnNotificationRemoved_AfterDelete(row *NotificationRemoved) {
    Cacher.Delete("NotificationRemoved:" + strconv.Itoa(row.NotificationId))
}

func OnNotificationRemoved_LoadOne(row *NotificationRemoved) {
    Cacher.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
}

func OnNotificationRemoved_LoadMany(rows []*NotificationRemoved) {
    for _, row := range rows {
        Cacher.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
    }
}

//PhoneContact Events
func OnPhoneContact_AfterInsert(row *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_AfterUpdate(row *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_AfterDelete(row *PhoneContact) {
    Cacher.Delete("PhoneContact:" + strconv.Itoa(row.Id))
}

func OnPhoneContact_LoadOne(row *PhoneContact) {
    Cacher.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_LoadMany(rows []*PhoneContact) {
    for _, row := range rows {
        Cacher.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Post Events
func OnPost_AfterInsert(row *Post) {
    Cacher.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPost_AfterUpdate(row *Post) {
    Cacher.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPost_AfterDelete(row *Post) {
    Cacher.Delete("Post:" + strconv.Itoa(row.Id))
}

func OnPost_LoadOne(row *Post) {
    Cacher.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPost_LoadMany(rows []*Post) {
    for _, row := range rows {
        Cacher.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//RecommendUser Events
func OnRecommendUser_AfterInsert(row *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_AfterUpdate(row *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_AfterDelete(row *RecommendUser) {
    Cacher.Delete("RecommendUser:" + strconv.Itoa(row.Id))
}

func OnRecommendUser_LoadOne(row *RecommendUser) {
    Cacher.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_LoadMany(rows []*RecommendUser) {
    for _, row := range rows {
        Cacher.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//SearchClicked Events
func OnSearchClicked_AfterInsert(row *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_AfterUpdate(row *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_AfterDelete(row *SearchClicked) {
    Cacher.Delete("SearchClicked:" + strconv.Itoa(row.Id))
}

func OnSearchClicked_LoadOne(row *SearchClicked) {
    Cacher.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_LoadMany(rows []*SearchClicked) {
    for _, row := range rows {
        Cacher.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Session Events
func OnSession_AfterInsert(row *Session) {
    Cacher.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_AfterUpdate(row *Session) {
    Cacher.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_AfterDelete(row *Session) {
    Cacher.Delete("Session:" + strconv.Itoa(row.Id))
}

func OnSession_LoadOne(row *Session) {
    Cacher.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_LoadMany(rows []*Session) {
    for _, row := range rows {
        Cacher.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//Tag Events
func OnTag_AfterInsert(row *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTag_AfterUpdate(row *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTag_AfterDelete(row *Tag) {
    Cacher.Delete("Tag:" + strconv.Itoa(row.Id))
}

func OnTag_LoadOne(row *Tag) {
    Cacher.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTag_LoadMany(rows []*Tag) {
    for _, row := range rows {
        Cacher.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//TagsPost Events
func OnTagsPost_AfterInsert(row *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_AfterUpdate(row *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_AfterDelete(row *TagsPost) {
    Cacher.Delete("TagsPost:" + strconv.Itoa(row.Id))
}

func OnTagsPost_LoadOne(row *TagsPost) {
    Cacher.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_LoadMany(rows []*TagsPost) {
    for _, row := range rows {
        Cacher.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//User Events
func OnUser_AfterInsert(row *User) {
    Cacher.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUser_AfterUpdate(row *User) {
    Cacher.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUser_AfterDelete(row *User) {
    Cacher.Delete("User:" + strconv.Itoa(row.Id))
}

func OnUser_LoadOne(row *User) {
    Cacher.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUser_LoadMany(rows []*User) {
    for _, row := range rows {
        Cacher.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
    }
}

//UserMetaInfo Events
func OnUserMetaInfo_AfterInsert(row *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserMetaInfo_AfterUpdate(row *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserMetaInfo_AfterDelete(row *UserMetaInfo) {
    Cacher.Delete("UserMetaInfo:" + strconv.Itoa(row.UserId))
}

func OnUserMetaInfo_LoadOne(row *UserMetaInfo) {
    Cacher.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserMetaInfo_LoadMany(rows []*UserMetaInfo) {
    for _, row := range rows {
        Cacher.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
    }
}

//UserPassword Events
func OnUserPassword_AfterInsert(row *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_AfterUpdate(row *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_AfterDelete(row *UserPassword) {
    Cacher.Delete("UserPassword:" + strconv.Itoa(row.UserId))
}

func OnUserPassword_LoadOne(row *UserPassword) {
    Cacher.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_LoadMany(rows []*UserPassword) {
    for _, row := range rows {
        Cacher.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
    }
}