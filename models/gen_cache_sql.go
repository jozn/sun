// Package models contains the types for schema 'ms'.
package models

import (
	"ms/sun/base"
	"strconv"
)

// GENERATED BY XO. DO NOT EDIT.

func (c _StoreImpl) GetActivityById__(Id int) (*Activity, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadActivityById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetCommentById__(Id int) (*Comment, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadCommentById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetFollowingListByUserId__(UserId int) (*FollowingList, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListByUserId__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetFollowingListMemberById__(Id int) (*FollowingListMember, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetFollowingListMemberHistoryById__(Id int) (*FollowingListMemberHistory, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadFollowingListMemberHistoryById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetLikeById__(Id int) (*Like, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadLikeById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetMediaById__(Id int) (*Media, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadMediaById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetMessageById__(Id int) (*Message, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadMessageById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetMsgDeletedFromServerById__(Id int) (*MsgDeletedFromServer, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadMsgDeletedFromServerById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetMsgReceivedToPeerById__(Id int) (*MsgReceivedToPeer, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadMsgReceivedToPeerById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetMsgSeenByPeerById__(Id int) (*MsgSeenByPeer, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadMsgSeenByPeerById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetNotificationById__(Id int) (*Notification, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadNotificationById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetNotificationRemovedByNotificationId__(NotificationId int) (*NotificationRemoved, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadNotificationRemovedByNotificationId__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetPhoneContactById__(Id int) (*PhoneContact, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContactById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetPostById__(Id int) (*Post, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadPostById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetRecommendUserById__(Id int) (*RecommendUser, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadRecommendUserById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetSearchClickedById__(Id int) (*SearchClicked, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadSearchClickedById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetSessionById__(Id int) (*Session, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadSessionById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetTagById__(Id int) (*Tag, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadTagById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetTagsPostById__(Id int) (*TagsPost, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadTagsPostById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetUserById__(Id int) (*User, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadUserById__(ids []int) {
	var not_cached []int

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

func (c _StoreImpl) GetUserMetaInfoByUserId__(UserId int) (*UserMetaInfo, bool) {
	o, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(UserId))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}
	obj2, err := UserMetaInfoByUserId(base.DB, UserId)
	if err == nil {
		return obj2, true
	}
	return nil, false
}

func (c _StoreImpl) PreLoadUserMetaInfoByUserId__(ids []int) {
	var not_cached []int

	for _, id := range ids {
		_, ok := RowCache.Get("UserMetaInfo:" + strconv.Itoa(id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		NewUserMetaInfo_Selector().UserId_In(not_cached).GetRows(base.DB)
	}
}

func (c _StoreImpl) GetUserPasswordByUserId__(UserId int) (*UserPassword, bool) {
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
	return nil, false
}

func (c _StoreImpl) PreLoadUserPasswordByUserId__(ids []int) {
	var not_cached []int

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
