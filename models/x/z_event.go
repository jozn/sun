package x

import (
	"strconv"
	"time"
)

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

//Chat Events

func OnChat_AfterInsert(row *Chat) {
	RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
}

func OnChat_AfterUpdate(row *Chat) {
	RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
}

func OnChat_AfterDelete(row *Chat) {
	RowCache.Delete("Chat:" + row.ChatKey)
}

func OnChat_LoadOne(row *Chat) {
	RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
}

func OnChat_LoadMany(rows []*Chat) {
	for _, row := range rows {
		RowCache.Set("Chat:"+row.ChatKey, row, time.Hour*0)
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

//DirectMessage Events

func OnDirectMessage_AfterInsert(row *DirectMessage) {
	RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnDirectMessage_AfterUpdate(row *DirectMessage) {
	RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnDirectMessage_AfterDelete(row *DirectMessage) {
	RowCache.Delete("DirectMessage:" + strconv.Itoa(row.MessageId))
}

func OnDirectMessage_LoadOne(row *DirectMessage) {
	RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnDirectMessage_LoadMany(rows []*DirectMessage) {
	for _, row := range rows {
		RowCache.Set("DirectMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
	}
}

//DirectOffline Events

func OnDirectOffline_AfterInsert(row *DirectOffline) {
	RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOffline_AfterUpdate(row *DirectOffline) {
	RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOffline_AfterDelete(row *DirectOffline) {
	RowCache.Delete("DirectOffline:" + strconv.Itoa(row.DirectOfflineId))
}

func OnDirectOffline_LoadOne(row *DirectOffline) {
	RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOffline_LoadMany(rows []*DirectOffline) {
	for _, row := range rows {
		RowCache.Set("DirectOffline:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
	}
}

//DirectOfflineDep Events

func OnDirectOfflineDep_AfterInsert(row *DirectOfflineDep) {
	RowCache.Set("DirectOfflineDep:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOfflineDep_AfterUpdate(row *DirectOfflineDep) {
	RowCache.Set("DirectOfflineDep:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOfflineDep_AfterDelete(row *DirectOfflineDep) {
	RowCache.Delete("DirectOfflineDep:" + strconv.Itoa(row.DirectOfflineId))
}

func OnDirectOfflineDep_LoadOne(row *DirectOfflineDep) {
	RowCache.Set("DirectOfflineDep:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
}

func OnDirectOfflineDep_LoadMany(rows []*DirectOfflineDep) {
	for _, row := range rows {
		RowCache.Set("DirectOfflineDep:"+strconv.Itoa(row.DirectOfflineId), row, time.Hour*0)
	}
}

//DirectToMessage Events

func OnDirectToMessage_AfterInsert(row *DirectToMessage) {
	RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnDirectToMessage_AfterUpdate(row *DirectToMessage) {
	RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnDirectToMessage_AfterDelete(row *DirectToMessage) {
	RowCache.Delete("DirectToMessage:" + strconv.Itoa(row.Id))
}

func OnDirectToMessage_LoadOne(row *DirectToMessage) {
	RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnDirectToMessage_LoadMany(rows []*DirectToMessage) {
	for _, row := range rows {
		RowCache.Set("DirectToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//DirectUpdate Events

func OnDirectUpdate_AfterInsert(row *DirectUpdate) {
	RowCache.Set("DirectUpdate:"+strconv.Itoa(row.DirectUpdateId), row, time.Hour*0)
}

func OnDirectUpdate_AfterUpdate(row *DirectUpdate) {
	RowCache.Set("DirectUpdate:"+strconv.Itoa(row.DirectUpdateId), row, time.Hour*0)
}

func OnDirectUpdate_AfterDelete(row *DirectUpdate) {
	RowCache.Delete("DirectUpdate:" + strconv.Itoa(row.DirectUpdateId))
}

func OnDirectUpdate_LoadOne(row *DirectUpdate) {
	RowCache.Set("DirectUpdate:"+strconv.Itoa(row.DirectUpdateId), row, time.Hour*0)
}

func OnDirectUpdate_LoadMany(rows []*DirectUpdate) {
	for _, row := range rows {
		RowCache.Set("DirectUpdate:"+strconv.Itoa(row.DirectUpdateId), row, time.Hour*0)
	}
}

//FollowingList Events

func OnFollowingList_AfterInsert(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingList_AfterUpdate(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingList_AfterDelete(row *FollowingList) {
	RowCache.Delete("FollowingList:" + strconv.Itoa(row.Id))
}

func OnFollowingList_LoadOne(row *FollowingList) {
	RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnFollowingList_LoadMany(rows []*FollowingList) {
	for _, row := range rows {
		RowCache.Set("FollowingList:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//GeneralLog Events

func OnGeneralLog_AfterInsert(row *GeneralLog) {
	RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGeneralLog_AfterUpdate(row *GeneralLog) {
	RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGeneralLog_AfterDelete(row *GeneralLog) {
	RowCache.Delete("GeneralLog:" + strconv.Itoa(row.Id))
}

func OnGeneralLog_LoadOne(row *GeneralLog) {
	RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGeneralLog_LoadMany(rows []*GeneralLog) {
	for _, row := range rows {
		RowCache.Set("GeneralLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//Group Events

func OnGroup_AfterInsert(row *Group) {
	RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
}

func OnGroup_AfterUpdate(row *Group) {
	RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
}

func OnGroup_AfterDelete(row *Group) {
	RowCache.Delete("Group:" + strconv.Itoa(row.GroupId))
}

func OnGroup_LoadOne(row *Group) {
	RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
}

func OnGroup_LoadMany(rows []*Group) {
	for _, row := range rows {
		RowCache.Set("Group:"+strconv.Itoa(row.GroupId), row, time.Hour*0)
	}
}

//GroupMember Events

func OnGroupMember_AfterInsert(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupMember_AfterUpdate(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupMember_AfterDelete(row *GroupMember) {
	RowCache.Delete("GroupMember:" + strconv.Itoa(row.Id))
}

func OnGroupMember_LoadOne(row *GroupMember) {
	RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupMember_LoadMany(rows []*GroupMember) {
	for _, row := range rows {
		RowCache.Set("GroupMember:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//GroupMessage Events

func OnGroupMessage_AfterInsert(row *GroupMessage) {
	RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnGroupMessage_AfterUpdate(row *GroupMessage) {
	RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnGroupMessage_AfterDelete(row *GroupMessage) {
	RowCache.Delete("GroupMessage:" + strconv.Itoa(row.MessageId))
}

func OnGroupMessage_LoadOne(row *GroupMessage) {
	RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
}

func OnGroupMessage_LoadMany(rows []*GroupMessage) {
	for _, row := range rows {
		RowCache.Set("GroupMessage:"+strconv.Itoa(row.MessageId), row, time.Hour*0)
	}
}

//GroupToMessage Events

func OnGroupToMessage_AfterInsert(row *GroupToMessage) {
	RowCache.Set("GroupToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupToMessage_AfterUpdate(row *GroupToMessage) {
	RowCache.Set("GroupToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupToMessage_AfterDelete(row *GroupToMessage) {
	RowCache.Delete("GroupToMessage:" + strconv.Itoa(row.Id))
}

func OnGroupToMessage_LoadOne(row *GroupToMessage) {
	RowCache.Set("GroupToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnGroupToMessage_LoadMany(rows []*GroupToMessage) {
	for _, row := range rows {
		RowCache.Set("GroupToMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//LogChange Events

func OnLogChange_AfterInsert(row *LogChange) {
	RowCache.Set("LogChange:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLogChange_AfterUpdate(row *LogChange) {
	RowCache.Set("LogChange:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLogChange_AfterDelete(row *LogChange) {
	RowCache.Delete("LogChange:" + strconv.Itoa(row.Id))
}

func OnLogChange_LoadOne(row *LogChange) {
	RowCache.Set("LogChange:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnLogChange_LoadMany(rows []*LogChange) {
	for _, row := range rows {
		RowCache.Set("LogChange:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//MessageFile Events

func OnMessageFile_AfterInsert(row *MessageFile) {
	RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
}

func OnMessageFile_AfterUpdate(row *MessageFile) {
	RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
}

func OnMessageFile_AfterDelete(row *MessageFile) {
	RowCache.Delete("MessageFile:" + strconv.Itoa(row.MessageFileId))
}

func OnMessageFile_LoadOne(row *MessageFile) {
	RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
}

func OnMessageFile_LoadMany(rows []*MessageFile) {
	for _, row := range rows {
		RowCache.Set("MessageFile:"+strconv.Itoa(row.MessageFileId), row, time.Hour*0)
	}
}

//Msg Events

func OnMsg_AfterInsert(row *Msg) {
	RowCache.Set("Msg:"+row.Key, row, time.Hour*0)
}

func OnMsg_AfterUpdate(row *Msg) {
	RowCache.Set("Msg:"+row.Key, row, time.Hour*0)
}

func OnMsg_AfterDelete(row *Msg) {
	RowCache.Delete("Msg:" + row.Key)
}

func OnMsg_LoadOne(row *Msg) {
	RowCache.Set("Msg:"+row.Key, row, time.Hour*0)
}

func OnMsg_LoadMany(rows []*Msg) {
	for _, row := range rows {
		RowCache.Set("Msg:"+row.Key, row, time.Hour*0)
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

//Offline Events

func OnOffline_AfterInsert(row *Offline) {
	RowCache.Set("Offline:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOffline_AfterUpdate(row *Offline) {
	RowCache.Set("Offline:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOffline_AfterDelete(row *Offline) {
	RowCache.Delete("Offline:" + strconv.Itoa(row.Id))
}

func OnOffline_LoadOne(row *Offline) {
	RowCache.Set("Offline:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOffline_LoadMany(rows []*Offline) {
	for _, row := range rows {
		RowCache.Set("Offline:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//OldMessage Events

func OnOldMessage_AfterInsert(row *OldMessage) {
	RowCache.Set("OldMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMessage_AfterUpdate(row *OldMessage) {
	RowCache.Set("OldMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMessage_AfterDelete(row *OldMessage) {
	RowCache.Delete("OldMessage:" + strconv.Itoa(row.Id))
}

func OnOldMessage_LoadOne(row *OldMessage) {
	RowCache.Set("OldMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMessage_LoadMany(rows []*OldMessage) {
	for _, row := range rows {
		RowCache.Set("OldMessage:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//OldMsgFile Events

func OnOldMsgFile_AfterInsert(row *OldMsgFile) {
	RowCache.Set("OldMsgFile:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgFile_AfterUpdate(row *OldMsgFile) {
	RowCache.Set("OldMsgFile:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgFile_AfterDelete(row *OldMsgFile) {
	RowCache.Delete("OldMsgFile:" + strconv.Itoa(row.Id))
}

func OnOldMsgFile_LoadOne(row *OldMsgFile) {
	RowCache.Set("OldMsgFile:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgFile_LoadMany(rows []*OldMsgFile) {
	for _, row := range rows {
		RowCache.Set("OldMsgFile:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//OldMsgPush Events

func OnOldMsgPush_AfterInsert(row *OldMsgPush) {
	RowCache.Set("OldMsgPush:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgPush_AfterUpdate(row *OldMsgPush) {
	RowCache.Set("OldMsgPush:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgPush_AfterDelete(row *OldMsgPush) {
	RowCache.Delete("OldMsgPush:" + strconv.Itoa(row.Id))
}

func OnOldMsgPush_LoadOne(row *OldMsgPush) {
	RowCache.Set("OldMsgPush:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgPush_LoadMany(rows []*OldMsgPush) {
	for _, row := range rows {
		RowCache.Set("OldMsgPush:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

//OldMsgPushEvent Events

func OnOldMsgPushEvent_AfterInsert(row *OldMsgPushEvent) {
	RowCache.Set("OldMsgPushEvent:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgPushEvent_AfterUpdate(row *OldMsgPushEvent) {
	RowCache.Set("OldMsgPushEvent:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgPushEvent_AfterDelete(row *OldMsgPushEvent) {
	RowCache.Delete("OldMsgPushEvent:" + strconv.Itoa(row.Id))
}

func OnOldMsgPushEvent_LoadOne(row *OldMsgPushEvent) {
	RowCache.Set("OldMsgPushEvent:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnOldMsgPushEvent_LoadMany(rows []*OldMsgPushEvent) {
	for _, row := range rows {
		RowCache.Set("OldMsgPushEvent:"+strconv.Itoa(row.Id), row, time.Hour*0)
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

//PushEvent Events

func OnPushEvent_AfterInsert(row *PushEvent) {
	RowCache.Set("PushEvent:"+strconv.Itoa(row.PushEventId), row, time.Hour*0)
}

func OnPushEvent_AfterUpdate(row *PushEvent) {
	RowCache.Set("PushEvent:"+strconv.Itoa(row.PushEventId), row, time.Hour*0)
}

func OnPushEvent_AfterDelete(row *PushEvent) {
	RowCache.Delete("PushEvent:" + strconv.Itoa(row.PushEventId))
}

func OnPushEvent_LoadOne(row *PushEvent) {
	RowCache.Set("PushEvent:"+strconv.Itoa(row.PushEventId), row, time.Hour*0)
}

func OnPushEvent_LoadMany(rows []*PushEvent) {
	for _, row := range rows {
		RowCache.Set("PushEvent:"+strconv.Itoa(row.PushEventId), row, time.Hour*0)
	}
}

//PushMessage Events

func OnPushMessage_AfterInsert(row *PushMessage) {
	RowCache.Set("PushMessage:"+strconv.Itoa(row.PushMessageId), row, time.Hour*0)
}

func OnPushMessage_AfterUpdate(row *PushMessage) {
	RowCache.Set("PushMessage:"+strconv.Itoa(row.PushMessageId), row, time.Hour*0)
}

func OnPushMessage_AfterDelete(row *PushMessage) {
	RowCache.Delete("PushMessage:" + strconv.Itoa(row.PushMessageId))
}

func OnPushMessage_LoadOne(row *PushMessage) {
	RowCache.Set("PushMessage:"+strconv.Itoa(row.PushMessageId), row, time.Hour*0)
}

func OnPushMessage_LoadMany(rows []*PushMessage) {
	for _, row := range rows {
		RowCache.Set("PushMessage:"+strconv.Itoa(row.PushMessageId), row, time.Hour*0)
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

//Room Events

func OnRoom_AfterInsert(row *Room) {
	RowCache.Set("Room:"+strconv.Itoa(row.RoomId), row, time.Hour*0)
}

func OnRoom_AfterUpdate(row *Room) {
	RowCache.Set("Room:"+strconv.Itoa(row.RoomId), row, time.Hour*0)
}

func OnRoom_AfterDelete(row *Room) {
	RowCache.Delete("Room:" + strconv.Itoa(row.RoomId))
}

func OnRoom_LoadOne(row *Room) {
	RowCache.Set("Room:"+strconv.Itoa(row.RoomId), row, time.Hour*0)
}

func OnRoom_LoadMany(rows []*Room) {
	for _, row := range rows {
		RowCache.Set("Room:"+strconv.Itoa(row.RoomId), row, time.Hour*0)
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

//SettingClient Events

func OnSettingClient_AfterInsert(row *SettingClient) {
	RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingClient_AfterUpdate(row *SettingClient) {
	RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingClient_AfterDelete(row *SettingClient) {
	RowCache.Delete("SettingClient:" + strconv.Itoa(row.UserId))
}

func OnSettingClient_LoadOne(row *SettingClient) {
	RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingClient_LoadMany(rows []*SettingClient) {
	for _, row := range rows {
		RowCache.Set("SettingClient:"+strconv.Itoa(row.UserId), row, time.Hour*0)
	}
}

//SettingNotification Events

func OnSettingNotification_AfterInsert(row *SettingNotification) {
	RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotification_AfterUpdate(row *SettingNotification) {
	RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotification_AfterDelete(row *SettingNotification) {
	RowCache.Delete("SettingNotification:" + strconv.Itoa(row.UserId))
}

func OnSettingNotification_LoadOne(row *SettingNotification) {
	RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
}

func OnSettingNotification_LoadMany(rows []*SettingNotification) {
	for _, row := range rows {
		RowCache.Set("SettingNotification:"+strconv.Itoa(row.UserId), row, time.Hour*0)
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

//TestChat Events

func OnTestChat_AfterInsert(row *TestChat) {
	RowCache.Set("TestChat:"+strconv.Itoa(row.Id4), row, time.Hour*0)
}

func OnTestChat_AfterUpdate(row *TestChat) {
	RowCache.Set("TestChat:"+strconv.Itoa(row.Id4), row, time.Hour*0)
}

func OnTestChat_AfterDelete(row *TestChat) {
	RowCache.Delete("TestChat:" + strconv.Itoa(row.Id4))
}

func OnTestChat_LoadOne(row *TestChat) {
	RowCache.Set("TestChat:"+strconv.Itoa(row.Id4), row, time.Hour*0)
}

func OnTestChat_LoadMany(rows []*TestChat) {
	for _, row := range rows {
		RowCache.Set("TestChat:"+strconv.Itoa(row.Id4), row, time.Hour*0)
	}
}

//TriggerLog Events

func OnTriggerLog_AfterInsert(row *TriggerLog) {
	RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTriggerLog_AfterUpdate(row *TriggerLog) {
	RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTriggerLog_AfterDelete(row *TriggerLog) {
	RowCache.Delete("TriggerLog:" + strconv.Itoa(row.Id))
}

func OnTriggerLog_LoadOne(row *TriggerLog) {
	RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnTriggerLog_LoadMany(rows []*TriggerLog) {
	for _, row := range rows {
		RowCache.Set("TriggerLog:"+strconv.Itoa(row.Id), row, time.Hour*0)
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
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUserMetaInfo_AfterUpdate(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUserMetaInfo_AfterDelete(row *UserMetaInfo) {
	RowCache.Delete("UserMetaInfo:" + strconv.Itoa(row.Id))
}

func OnUserMetaInfo_LoadOne(row *UserMetaInfo) {
	RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
}

func OnUserMetaInfo_LoadMany(rows []*UserMetaInfo) {
	for _, row := range rows {
		RowCache.Set("UserMetaInfo:"+strconv.Itoa(row.Id), row, time.Hour*0)
	}
}

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