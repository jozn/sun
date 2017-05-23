// Package x contains the types for schema 'ms'.
package x

import (
	"ms/sun/base"
	"strconv"
)

// GENERATED BY XO. DO NOT EDIT.

func (c _StoreImpl) GetMessageById(Id int) (*Message, bool) {
	o, ok := RowCache.Get("Message:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Message); ok {
			return obj, true
		}
	}
	obj2, err := MessageById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMessageByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Message:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMessage_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

func (c _StoreImpl) GetMsgDeletedFromServerById(Id int) (*MsgDeletedFromServer, bool) {
	o, ok := RowCache.Get("MsgDeletedFromServer:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*MsgDeletedFromServer); ok {
			return obj, true
		}
	}
	obj2, err := MsgDeletedFromServerById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMsgDeletedFromServerByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("MsgDeletedFromServer:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMsgDeletedFromServer_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

func (c _StoreImpl) GetTagsPostById(Id int) (*TagsPost, bool) {
	o, ok := RowCache.Get("TagsPost:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*TagsPost); ok {
			return obj, true
		}
	}
	obj2, err := TagsPostById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTagsPostByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("TagsPost:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTagsPost_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

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

func (c _StoreImpl) GetTagById(Id int) (*Tag, bool) {
	o, ok := RowCache.Get("Tag:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Tag); ok {
			return obj, true
		}
	}
	obj2, err := TagById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTagByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Tag:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewTag_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

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

func (c _StoreImpl) GetCommentById(Id int) (*Comment, bool) {
	o, ok := RowCache.Get("Comment:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}
	obj2, err := CommentById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadCommentByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Comment:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewComment_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

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

func (c _StoreImpl) GetMsgReceivedToPeerById(Id int) (*MsgReceivedToPeer, bool) {
	o, ok := RowCache.Get("MsgReceivedToPeer:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*MsgReceivedToPeer); ok {
			return obj, true
		}
	}
	obj2, err := MsgReceivedToPeerById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMsgReceivedToPeerByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("MsgReceivedToPeer:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMsgReceivedToPeer_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

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

func (c _StoreImpl) GetLikeById(Id int) (*Like, bool) {
	o, ok := RowCache.Get("Like:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}
	obj2, err := LikeById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLikeByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("Like:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewLike_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

func (c _StoreImpl) GetMsgSeenByPeerById(Id int) (*MsgSeenByPeer, bool) {
	o, ok := RowCache.Get("MsgSeenByPeer:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*MsgSeenByPeer); ok {
			return obj, true
		}
	}
	obj2, err := MsgSeenByPeerById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadMsgSeenByPeerByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("MsgSeenByPeer:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewMsgSeenByPeer_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

func (c _StoreImpl) GetPhoneContactById(Id int) (*PhoneContact, bool) {
	o, ok := RowCache.Get("PhoneContact:" + strconv.Itoa(Id))
	if ok {
		if obj, ok := o.(*PhoneContact); ok {
			return obj, true
		}
	}
	obj2, err := PhoneContactById(base.DB, Id)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContactByIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("PhoneContact:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewPhoneContact_Selector().Id_In(not_cached).GetRows(base.DB)
	}
}

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

func (c _StoreImpl) GetFollowingListByUserId(UserId int) (*FollowingList, bool) {
	o, ok := RowCache.Get("FollowingList:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*FollowingList); ok {
			return obj, true
		}
	}
	obj2, err := FollowingListByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListByUserIds(ids []int) {
	not_cached := make([]int, 0, len(ids))

	for _, id := range ids {
		_, ok := RowCache.Get("FollowingList:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewFollowingList_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

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
