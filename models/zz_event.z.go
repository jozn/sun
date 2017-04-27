// Package models contains the types for schema 'ms'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"strconv"
	"time"
)

//UserPassword Events
func OnUserPassword_AfterInsert(row *UserPassword) {
	RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_AfterUpdate(row *UserPassword) {
	RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_AfterDelete(row *UserPassword) {
	RowCache.Delete("UserPassword:" + strconv.Itoa(row.UserId))
}

func OnUserPassword_LoadOne(row *UserPassword) {
	RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserPassword_LoadMany(rows []*UserPassword) {
	for _, row := range rows {
		RowCache.Set("UserPassword:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//FollowingListMember Events
func OnFollowingListMember_AfterInsert(row *FollowingListMember) {
	RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_AfterUpdate(row *FollowingListMember) {
	RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_AfterDelete(row *FollowingListMember) {
	RowCache.Delete("FollowingListMember:" + strconv.Itoa(row.Id))
}

func OnFollowingListMember_LoadOne(row *FollowingListMember) {
	RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMember_LoadMany(rows []*FollowingListMember) {
	for _, row := range rows {
		RowCache.Set("FollowingListMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Media Events
func OnMedia_AfterInsert(row *Media) {
	RowCache.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMedia_AfterUpdate(row *Media) {
	RowCache.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMedia_AfterDelete(row *Media) {
	RowCache.Delete("Media:" + strconv.Itoa(row.Id))
}

func OnMedia_LoadOne(row *Media) {
	RowCache.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMedia_LoadMany(rows []*Media) {
	for _, row := range rows {
		RowCache.Set("Media:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Message Events
func OnMessage_AfterInsert(row *Message) {
	RowCache.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMessage_AfterUpdate(row *Message) {
	RowCache.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMessage_AfterDelete(row *Message) {
	RowCache.Delete("Message:" + strconv.Itoa(row.Id))
}

func OnMessage_LoadOne(row *Message) {
	RowCache.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMessage_LoadMany(rows []*Message) {
	for _, row := range rows {
		RowCache.Set("Message:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Notification Events
func OnNotification_AfterInsert(row *Notification) {
	RowCache.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotification_AfterUpdate(row *Notification) {
	RowCache.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotification_AfterDelete(row *Notification) {
	RowCache.Delete("Notification:" + strconv.Itoa(row.Id))
}

func OnNotification_LoadOne(row *Notification) {
	RowCache.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnNotification_LoadMany(rows []*Notification) {
	for _, row := range rows {
		RowCache.Set("Notification:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//RecommendUser Events
func OnRecommendUser_AfterInsert(row *RecommendUser) {
	RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_AfterUpdate(row *RecommendUser) {
	RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_AfterDelete(row *RecommendUser) {
	RowCache.Delete("RecommendUser:" + strconv.Itoa(row.Id))
}

func OnRecommendUser_LoadOne(row *RecommendUser) {
	RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnRecommendUser_LoadMany(rows []*RecommendUser) {
	for _, row := range rows {
		RowCache.Set("RecommendUser:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//SearchClicked Events
func OnSearchClicked_AfterInsert(row *SearchClicked) {
	RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_AfterUpdate(row *SearchClicked) {
	RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_AfterDelete(row *SearchClicked) {
	RowCache.Delete("SearchClicked:" + strconv.Itoa(row.Id))
}

func OnSearchClicked_LoadOne(row *SearchClicked) {
	RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSearchClicked_LoadMany(rows []*SearchClicked) {
	for _, row := range rows {
		RowCache.Set("SearchClicked:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//TagsPost Events
func OnTagsPost_AfterInsert(row *TagsPost) {
	RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_AfterUpdate(row *TagsPost) {
	RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_AfterDelete(row *TagsPost) {
	RowCache.Delete("TagsPost:" + strconv.Itoa(row.Id))
}

func OnTagsPost_LoadOne(row *TagsPost) {
	RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTagsPost_LoadMany(rows []*TagsPost) {
	for _, row := range rows {
		RowCache.Set("TagsPost:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Activity Events
func OnActivity_AfterInsert(row *Activity) {
	RowCache.Set("Activity:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnActivity_AfterUpdate(row *Activity) {
	RowCache.Set("Activity:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnActivity_AfterDelete(row *Activity) {
	RowCache.Delete("Activity:" + strconv.Itoa(row.Id))
}

func OnActivity_LoadOne(row *Activity) {
	RowCache.Set("Activity:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnActivity_LoadMany(rows []*Activity) {
	for _, row := range rows {
		RowCache.Set("Activity:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Bucket Events
func OnBucket_AfterInsert(row *Bucket) {
	RowCache.Set("Bucket:"+strconv.Itoa(row.BucketId), row, time.Hour*0)
}

func OnBucket_AfterUpdate(row *Bucket) {
	RowCache.Set("Bucket:"+strconv.Itoa(row.BucketId), row, time.Hour*0)
}

func OnBucket_AfterDelete(row *Bucket) {
	RowCache.Delete("Bucket:" + strconv.Itoa(row.BucketId))
}

func OnBucket_LoadOne(row *Bucket) {
	RowCache.Set("Bucket:"+strconv.Itoa(row.BucketId), row, time.Hour*0)
}

func OnBucket_LoadMany(rows []*Bucket) {
	for _, row := range rows {
		RowCache.Set("Bucket:"+strconv.Itoa(row.BucketId), row, time.Hour*0)
	}
}

//FollowingList Events
func OnFollowingList_AfterInsert(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnFollowingList_AfterUpdate(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnFollowingList_AfterDelete(row *FollowingList) {
	RowCache.Delete("FollowingList:" + strconv.Itoa(row.UserId))
}

func OnFollowingList_LoadOne(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnFollowingList_LoadMany(rows []*FollowingList) {
	for _, row := range rows {
		RowCache.Set("FollowingList:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//MsgDeletedFromServer Events
func OnMsgDeletedFromServer_AfterInsert(row *MsgDeletedFromServer) {
	RowCache.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgDeletedFromServer_AfterUpdate(row *MsgDeletedFromServer) {
	RowCache.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgDeletedFromServer_AfterDelete(row *MsgDeletedFromServer) {
	RowCache.Delete("MsgDeletedFromServer:" + strconv.Itoa(row.Id))
}

func OnMsgDeletedFromServer_LoadOne(row *MsgDeletedFromServer) {
	RowCache.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgDeletedFromServer_LoadMany(rows []*MsgDeletedFromServer) {
	for _, row := range rows {
		RowCache.Set("MsgDeletedFromServer:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//NotificationRemoved Events
func OnNotificationRemoved_AfterInsert(row *NotificationRemoved) {
	RowCache.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
}

func OnNotificationRemoved_AfterUpdate(row *NotificationRemoved) {
	RowCache.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
}

func OnNotificationRemoved_AfterDelete(row *NotificationRemoved) {
	RowCache.Delete("NotificationRemoved:" + strconv.Itoa(row.NotificationId))
}

func OnNotificationRemoved_LoadOne(row *NotificationRemoved) {
	RowCache.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
}

func OnNotificationRemoved_LoadMany(rows []*NotificationRemoved) {
	for _, row := range rows {
		RowCache.Set("NotificationRemoved:"+strconv.Itoa(row.NotificationId), row, time.Hour*0)
	}
}

//PhoneContact Events
func OnPhoneContact_AfterInsert(row *PhoneContact) {
	RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_AfterUpdate(row *PhoneContact) {
	RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_AfterDelete(row *PhoneContact) {
	RowCache.Delete("PhoneContact:" + strconv.Itoa(row.Id))
}

func OnPhoneContact_LoadOne(row *PhoneContact) {
	RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPhoneContact_LoadMany(rows []*PhoneContact) {
	for _, row := range rows {
		RowCache.Set("PhoneContact:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Photo Events
func OnPhoto_AfterInsert(row *Photo) {
	RowCache.Set("Photo:"+strconv.Itoa(row.PhotoId), row, time.Hour*0)
}

func OnPhoto_AfterUpdate(row *Photo) {
	RowCache.Set("Photo:"+strconv.Itoa(row.PhotoId), row, time.Hour*0)
}

func OnPhoto_AfterDelete(row *Photo) {
	RowCache.Delete("Photo:" + strconv.Itoa(row.PhotoId))
}

func OnPhoto_LoadOne(row *Photo) {
	RowCache.Set("Photo:"+strconv.Itoa(row.PhotoId), row, time.Hour*0)
}

func OnPhoto_LoadMany(rows []*Photo) {
	for _, row := range rows {
		RowCache.Set("Photo:"+strconv.Itoa(row.PhotoId), row, time.Hour*0)
	}
}

//FollowingListMemberHistory Events
func OnFollowingListMemberHistory_AfterInsert(row *FollowingListMemberHistory) {
	RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_AfterUpdate(row *FollowingListMemberHistory) {
	RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_AfterDelete(row *FollowingListMemberHistory) {
	RowCache.Delete("FollowingListMemberHistory:" + strconv.Itoa(row.Id))
}

func OnFollowingListMemberHistory_LoadOne(row *FollowingListMemberHistory) {
	RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingListMemberHistory_LoadMany(rows []*FollowingListMemberHistory) {
	for _, row := range rows {
		RowCache.Set("FollowingListMemberHistory:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Post Events
func OnPost_AfterInsert(row *Post) {
	RowCache.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPost_AfterUpdate(row *Post) {
	RowCache.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPost_AfterDelete(row *Post) {
	RowCache.Delete("Post:" + strconv.Itoa(row.Id))
}

func OnPost_LoadOne(row *Post) {
	RowCache.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnPost_LoadMany(rows []*Post) {
	for _, row := range rows {
		RowCache.Set("Post:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Tag Events
func OnTag_AfterInsert(row *Tag) {
	RowCache.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTag_AfterUpdate(row *Tag) {
	RowCache.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTag_AfterDelete(row *Tag) {
	RowCache.Delete("Tag:" + strconv.Itoa(row.Id))
}

func OnTag_LoadOne(row *Tag) {
	RowCache.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTag_LoadMany(rows []*Tag) {
	for _, row := range rows {
		RowCache.Set("Tag:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//User Events
func OnUser_AfterInsert(row *User) {
	RowCache.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUser_AfterUpdate(row *User) {
	RowCache.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUser_AfterDelete(row *User) {
	RowCache.Delete("User:" + strconv.Itoa(row.Id))
}

func OnUser_LoadOne(row *User) {
	RowCache.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUser_LoadMany(rows []*User) {
	for _, row := range rows {
		RowCache.Set("User:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//UserMetaInfo Events
func OnUserMetaInfo_AfterInsert(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserMetaInfo_AfterUpdate(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserMetaInfo_AfterDelete(row *UserMetaInfo) {
	RowCache.Delete("UserMetaInfo:" + strconv.Itoa(row.UserId))
}

func OnUserMetaInfo_LoadOne(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnUserMetaInfo_LoadMany(rows []*UserMetaInfo) {
	for _, row := range rows {
		RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//Comment Events
func OnComment_AfterInsert(row *Comment) {
	RowCache.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnComment_AfterUpdate(row *Comment) {
	RowCache.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnComment_AfterDelete(row *Comment) {
	RowCache.Delete("Comment:" + strconv.Itoa(row.Id))
}

func OnComment_LoadOne(row *Comment) {
	RowCache.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnComment_LoadMany(rows []*Comment) {
	for _, row := range rows {
		RowCache.Set("Comment:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Like Events
func OnLike_AfterInsert(row *Like) {
	RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_AfterUpdate(row *Like) {
	RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_AfterDelete(row *Like) {
	RowCache.Delete("Like:" + strconv.Itoa(row.Id))
}

func OnLike_LoadOne(row *Like) {
	RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLike_LoadMany(rows []*Like) {
	for _, row := range rows {
		RowCache.Set("Like:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//MsgReceivedToPeer Events
func OnMsgReceivedToPeer_AfterInsert(row *MsgReceivedToPeer) {
	RowCache.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgReceivedToPeer_AfterUpdate(row *MsgReceivedToPeer) {
	RowCache.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgReceivedToPeer_AfterDelete(row *MsgReceivedToPeer) {
	RowCache.Delete("MsgReceivedToPeer:" + strconv.Itoa(row.Id))
}

func OnMsgReceivedToPeer_LoadOne(row *MsgReceivedToPeer) {
	RowCache.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgReceivedToPeer_LoadMany(rows []*MsgReceivedToPeer) {
	for _, row := range rows {
		RowCache.Set("MsgReceivedToPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//MsgSeenByPeer Events
func OnMsgSeenByPeer_AfterInsert(row *MsgSeenByPeer) {
	RowCache.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgSeenByPeer_AfterUpdate(row *MsgSeenByPeer) {
	RowCache.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgSeenByPeer_AfterDelete(row *MsgSeenByPeer) {
	RowCache.Delete("MsgSeenByPeer:" + strconv.Itoa(row.Id))
}

func OnMsgSeenByPeer_LoadOne(row *MsgSeenByPeer) {
	RowCache.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnMsgSeenByPeer_LoadMany(rows []*MsgSeenByPeer) {
	for _, row := range rows {
		RowCache.Set("MsgSeenByPeer:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Session Events
func OnSession_AfterInsert(row *Session) {
	RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_AfterUpdate(row *Session) {
	RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_AfterDelete(row *Session) {
	RowCache.Delete("Session:" + strconv.Itoa(row.Id))
}

func OnSession_LoadOne(row *Session) {
	RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnSession_LoadMany(rows []*Session) {
	for _, row := range rows {
		RowCache.Set("Session:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}
