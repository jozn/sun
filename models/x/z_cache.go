package x

import (
	"ms/sun/base"
	"strconv"
)

func (c _StoreImpl) GetActivityById(Id int) (*Activity, bool) {
	o, ok := RowCache.Get("Activity:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Activity); ok {
			return obj, true
		}
	}
	obj2, err := ActivityById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadActivityByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Activity:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewActivity_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetBucketByBucketId(BucketId int) (*Bucket, bool) {
	o, ok := RowCache.Get("Bucket:" + strconv.Itoa(BucketId))
	if ok {
		if obj, ok := o.(*Bucket); ok {
			return obj, true
		}
	}
	obj2, err := BucketByBucketId(base.DB, BucketId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadBucketByBucketIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Bucket:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewBucket_Selector().BucketId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetChatByChatKey(ChatKey string) (*Chat, bool) {
	o, ok := RowCache.Get("Chat:" + ChatKey)
	if ok {
		if obj, ok := o.(*Chat); ok {
			return obj, true
		}
	}
	obj2, err := ChatByChatKey(base.DB, ChatKey)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadChatByChatKeys(ids []string) {
	not_cached := make([]string, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Chat:" + id)
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewChat_Selector().ChatKey_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 string

func (c _StoreImpl) GetCommentsById(Id int) (*Comments, bool) {
	o, ok := RowCache.Get("Comments:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Comments); ok {
			return obj, true
		}
	}
	obj2, err := CommentsById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadCommentsByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Comments:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewComments_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetDirectMessageByMessageId(MessageId int) (*DirectMessage, bool) {
	o, ok := RowCache.Get("DirectMessage:" + strconv.Itoa(MessageId))
	if ok {
		if obj, ok := o.(*DirectMessage); ok {
			return obj, true
		}
	}
	obj2, err := DirectMessageByMessageId(base.DB, MessageId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectMessageByMessageIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("DirectMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewDirectMessage_Selector().MessageId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetDirectToMessageById(Id int) (*DirectToMessage, bool) {
	o, ok := RowCache.Get("DirectToMessage:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*DirectToMessage); ok {
			return obj, true
		}
	}
	obj2, err := DirectToMessageById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectToMessageByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("DirectToMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewDirectToMessage_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetDirectUpdateByDirectUpdateId(DirectUpdateId int) (*DirectUpdate, bool) {
	o, ok := RowCache.Get("DirectUpdate:" + strconv.Itoa(DirectUpdateId))
	if ok {
		if obj, ok := o.(*DirectUpdate); ok {
			return obj, true
		}
	}
	obj2, err := DirectUpdateByDirectUpdateId(base.DB, DirectUpdateId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectUpdateByDirectUpdateIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("DirectUpdate:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewDirectUpdate_Selector().DirectUpdateId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFollowingListById(Id int) (*FollowingList, bool) {
	o, ok := RowCache.Get("FollowingList:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingList); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingList:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingList_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFollowingListMemberById(Id int) (*FollowingListMember, bool) {
	o, ok := RowCache.Get("FollowingListMember:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMember); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListMemberById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingListMember:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingListMember_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetFollowingListMemberHistoryById(Id int) (*FollowingListMemberHistory, bool) {
	o, ok := RowCache.Get("FollowingListMemberHistory:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*FollowingListMemberHistory); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListMemberHistoryById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberHistoryByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingListMemberHistory:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingListMemberHistory_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGeneralLogById(Id int) (*GeneralLog, bool) {
	o, ok := RowCache.Get("GeneralLog:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*GeneralLog); ok {
			return obj, true
		}
	}
	obj2, err := GeneralLogById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGeneralLogByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GeneralLog:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGeneralLog_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupByGroupId(GroupId int) (*Group, bool) {
	o, ok := RowCache.Get("Group:" + strconv.Itoa(GroupId))
	if ok {
		if obj, ok := o.(*Group); ok {
			return obj, true
		}
	}
	obj2, err := GroupByGroupId(base.DB, GroupId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupByGroupIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Group:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroup_Selector().GroupId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupMemberById(Id int) (*GroupMember, bool) {
	o, ok := RowCache.Get("GroupMember:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}
	obj2, err := GroupMemberById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMemberByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupMember:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupMember_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupMessageByMessageId(MessageId int) (*GroupMessage, bool) {
	o, ok := RowCache.Get("GroupMessage:" + strconv.Itoa(MessageId))
	if ok {
		if obj, ok := o.(*GroupMessage); ok {
			return obj, true
		}
	}
	obj2, err := GroupMessageByMessageId(base.DB, MessageId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMessageByMessageIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupMessage_Selector().MessageId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetGroupToMessageById(Id int) (*GroupToMessage, bool) {
	o, ok := RowCache.Get("GroupToMessage:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*GroupToMessage); ok {
			return obj, true
		}
	}
	obj2, err := GroupToMessageById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupToMessageByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("GroupToMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewGroupToMessage_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetLikesById(Id int) (*Likes, bool) {
	o, ok := RowCache.Get("Likes:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Likes); ok {
			return obj, true
		}
	}
	obj2, err := LikesById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLikesByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Likes:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewLikes_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetLogChangesById(Id int) (*LogChanges, bool) {
	o, ok := RowCache.Get("LogChanges:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*LogChanges); ok {
			return obj, true
		}
	}
	obj2, err := LogChangesById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLogChangesByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("LogChanges:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewLogChanges_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetMediaById(Id int) (*Media, bool) {
	o, ok := RowCache.Get("Media:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Media); ok {
			return obj, true
		}
	}
	obj2, err := MediaById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMediaByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Media:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMedia_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetMessageFileByMessageFileId(MessageFileId int) (*MessageFile, bool) {
	o, ok := RowCache.Get("MessageFile:" + strconv.Itoa(MessageFileId))
	if ok {
		if obj, ok := o.(*MessageFile); ok {
			return obj, true
		}
	}
	obj2, err := MessageFileByMessageFileId(base.DB, MessageFileId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMessageFileByMessageFileIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("MessageFile:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMessageFile_Selector().MessageFileId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetNotificationById(Id int) (*Notification, bool) {
	o, ok := RowCache.Get("Notification:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Notification); ok {
			return obj, true
		}
	}
	obj2, err := NotificationById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadNotificationByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Notification:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewNotification_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetNotificationRemovedByNotificationId(NotificationId int) (*NotificationRemoved, bool) {
	o, ok := RowCache.Get("NotificationRemoved:" + strconv.Itoa(NotificationId))
	if ok {
		if obj, ok := o.(*NotificationRemoved); ok {
			return obj, true
		}
	}
	obj2, err := NotificationRemovedByNotificationId(base.DB, NotificationId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadNotificationRemovedByNotificationIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("NotificationRemoved:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewNotificationRemoved_Selector().NotificationId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetOldMessagesById(Id int) (*OldMessages, bool) {
	o, ok := RowCache.Get("OldMessages:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*OldMessages); ok {
			return obj, true
		}
	}
	obj2, err := OldMessagesById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMessagesByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("OldMessages:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewOldMessages_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetOldMsgFileById(Id int) (*OldMsgFile, bool) {
	o, ok := RowCache.Get("OldMsgFile:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*OldMsgFile); ok {
			return obj, true
		}
	}
	obj2, err := OldMsgFileById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMsgFileByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("OldMsgFile:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewOldMsgFile_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetOldMsgPushById(Id int) (*OldMsgPush, bool) {
	o, ok := RowCache.Get("OldMsgPush:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*OldMsgPush); ok {
			return obj, true
		}
	}
	obj2, err := OldMsgPushById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMsgPushByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("OldMsgPush:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewOldMsgPush_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetOldMsgPushEventById(Id int) (*OldMsgPushEvent, bool) {
	o, ok := RowCache.Get("OldMsgPushEvent:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*OldMsgPushEvent); ok {
			return obj, true
		}
	}
	obj2, err := OldMsgPushEventById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMsgPushEventByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("OldMsgPushEvent:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewOldMsgPushEvent_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPhoneContactsById(Id int) (*PhoneContacts, bool) {
	o, ok := RowCache.Get("PhoneContacts:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PhoneContacts); ok {
			return obj, true
		}
	}
	obj2, err := PhoneContactsById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContactsByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PhoneContacts:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPhoneContacts_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPhotoByPhotoId(PhotoId int) (*Photo, bool) {
	o, ok := RowCache.Get("Photo:" + strconv.Itoa(PhotoId))
	if ok {
		if obj, ok := o.(*Photo); ok {
			return obj, true
		}
	}
	obj2, err := PhotoByPhotoId(base.DB, PhotoId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhotoByPhotoIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Photo:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPhoto_Selector().PhotoId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPostById(Id int) (*Post, bool) {
	o, ok := RowCache.Get("Post:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Post); ok {
			return obj, true
		}
	}
	obj2, err := PostById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPostByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Post:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPost_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPushEventByPushEventId(PushEventId int) (*PushEvent, bool) {
	o, ok := RowCache.Get("PushEvent:" + strconv.Itoa(PushEventId))
	if ok {
		if obj, ok := o.(*PushEvent); ok {
			return obj, true
		}
	}
	obj2, err := PushEventByPushEventId(base.DB, PushEventId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPushEventByPushEventIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PushEvent:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPushEvent_Selector().PushEventId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetPushMessageByPushMessageId(PushMessageId int) (*PushMessage, bool) {
	o, ok := RowCache.Get("PushMessage:" + strconv.Itoa(PushMessageId))
	if ok {
		if obj, ok := o.(*PushMessage); ok {
			return obj, true
		}
	}
	obj2, err := PushMessageByPushMessageId(base.DB, PushMessageId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPushMessageByPushMessageIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PushMessage:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPushMessage_Selector().PushMessageId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetRecommendUserById(Id int) (*RecommendUser, bool) {
	o, ok := RowCache.Get("RecommendUser:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*RecommendUser); ok {
			return obj, true
		}
	}
	obj2, err := RecommendUserById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadRecommendUserByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("RecommendUser:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewRecommendUser_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetRoomByRoomId(RoomId int) (*Room, bool) {
	o, ok := RowCache.Get("Room:" + strconv.Itoa(RoomId))
	if ok {
		if obj, ok := o.(*Room); ok {
			return obj, true
		}
	}
	obj2, err := RoomByRoomId(base.DB, RoomId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadRoomByRoomIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Room:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewRoom_Selector().RoomId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSearchClickedById(Id int) (*SearchClicked, bool) {
	o, ok := RowCache.Get("SearchClicked:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*SearchClicked); ok {
			return obj, true
		}
	}
	obj2, err := SearchClickedById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSearchClickedByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SearchClicked:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSearchClicked_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSessionById(Id int) (*Session, bool) {
	o, ok := RowCache.Get("Session:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}
	obj2, err := SessionById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSessionByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Session:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSession_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSettingClientByUserId(UserId int) (*SettingClient, bool) {
	o, ok := RowCache.Get("SettingClient:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*SettingClient); ok {
			return obj, true
		}
	}
	obj2, err := SettingClientByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSettingClientByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SettingClient:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSettingClient_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetSettingNotificationsByUserId(UserId int) (*SettingNotifications, bool) {
	o, ok := RowCache.Get("SettingNotifications:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*SettingNotifications); ok {
			return obj, true
		}
	}
	obj2, err := SettingNotificationsByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSettingNotificationsByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("SettingNotifications:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewSettingNotifications_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetTagsById(Id int) (*Tags, bool) {
	o, ok := RowCache.Get("Tags:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Tags); ok {
			return obj, true
		}
	}
	obj2, err := TagsById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTagsByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Tags:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTags_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetTagsPostsById(Id int) (*TagsPosts, bool) {
	o, ok := RowCache.Get("TagsPosts:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TagsPosts); ok {
			return obj, true
		}
	}
	obj2, err := TagsPostsById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTagsPostsByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("TagsPosts:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTagsPosts_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetTestChatById4(Id4 int) (*TestChat, bool) {
	o, ok := RowCache.Get("TestChat:" + strconv.Itoa(Id4))
	if ok {
		if obj, ok := o.(*TestChat); ok {
			return obj, true
		}
	}
	obj2, err := TestChatById4(base.DB, Id4)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTestChatById4s(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("TestChat:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTestChat_Selector().Id4_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetUserById(Id int) (*User, bool) {
	o, ok := RowCache.Get("User:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}
	obj2, err := UserById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("User:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUser_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetUserMetaInfoById(Id int) (*UserMetaInfo, bool) {
	o, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}
	obj2, err := UserMetaInfoById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserMetaInfoByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUserMetaInfo_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int

func (c _StoreImpl) GetUserPasswordByUserId(UserId int) (*UserPassword, bool) {
	o, ok := RowCache.Get("UserPassword:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*UserPassword); ok {
			return obj, true
		}
	}
	obj2, err := UserPasswordByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserPasswordByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("UserPassword:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUserPassword_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

// yes 222 int
