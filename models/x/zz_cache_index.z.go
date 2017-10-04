// Package x contains the types for schema 'ms'.
package x

// GENERATED BY XO. DO NOT EDIT.

import (
	"fmt"
	"ms/sun/base"
)

// Comment - comments_Id_pkey

// Bucket - bucket_BucketId_pkey

// FollowingList - following_list_UserId_pkey

// RecommendUser - recommend_user_Id_pkey

// GroupMember - group_member_Id_pkey

// OldMsgPush - ToUser

// FollowingListMember - FollowedUserId

// Tag - tags_Id_pkey

// PhoneContact - PhoneContactRowId

// OldMessage - old_messages_Id_pkey

//field//field//field

///// Generated from index 'ToUserId'.
func (c _StoreImpl) OldMsgPushEvent_ByToUserId(ToUserId int) (*OldMsgPushEvent, bool) {
	o, ok := RowCacheIndex.Get("OldMsgPushEvent_ToUserId:" + fmt.Sprintf("%v", ToUserId))
	if ok {
		if obj, ok := o.(*OldMsgPushEvent); ok {
			return obj, true
		}
	}

	row, err := NewOldMsgPushEvent_Selector().ToUserId_Eq(ToUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("OldMsgPushEvent_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMsgPushEvent_ByToUserIds(ToUserIds []int) {
	not_cached := make([]int, 0, len(ToUserIds))

	for _, id := range ToUserIds {
		_, ok := RowCacheIndex.Get("OldMsgPushEvent_ToUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewOldMsgPushEvent_Selector().ToUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("OldMsgPushEvent_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'SessionUuid2'.
func (c _StoreImpl) Session_BySessionUuid(SessionUuid string) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_SessionUuid2:" + fmt.Sprintf("%v", SessionUuid))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().SessionUuid_Eq(SessionUuid).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_SessionUuid2:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSession_BySessionUuids(SessionUuids []string) {
	not_cached := make([]string, 0, len(SessionUuids))

	for _, id := range SessionUuids {
		_, ok := RowCacheIndex.Get("Session_SessionUuid2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().SessionUuid_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_SessionUuid2:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
			}
		}
	}
}

// TestChat - test_chat_Id4_pkey

//field//field//field

///// Generated from index 'Uid'.
func (c _StoreImpl) OldMsgPushEvent_ByUid(Uid int) (*OldMsgPushEvent, bool) {
	o, ok := RowCacheIndex.Get("OldMsgPushEvent_Uid:" + fmt.Sprintf("%v", Uid))
	if ok {
		if obj, ok := o.(*OldMsgPushEvent); ok {
			return obj, true
		}
	}

	row, err := NewOldMsgPushEvent_Selector().Uid_Eq(Uid).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("OldMsgPushEvent_Uid:"+fmt.Sprintf("%v", row.Uid), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMsgPushEvent_ByUids(Uids []int) {
	not_cached := make([]int, 0, len(Uids))

	for _, id := range Uids {
		_, ok := RowCacheIndex.Get("OldMsgPushEvent_Uid:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewOldMsgPushEvent_Selector().Uid_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("OldMsgPushEvent_Uid:"+fmt.Sprintf("%v", row.Uid), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'PostId'.
func (c _StoreImpl) Comment_ByPostId(PostId int) (*Comment, bool) {
	o, ok := RowCacheIndex.Get("Comment_PostId:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Comment); ok {
			return obj, true
		}
	}

	row, err := NewComment_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Comment_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadComment_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Comment_PostId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewComment_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Comment_PostId:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// TagsPost - tags_posts_Id_pkey

// DirectToMessage - direct_to_message_Id_pkey

//field//field//field

///// Generated from index 'PostId_2'.
func (c _StoreImpl) Like_ByPostId(PostId int) (*Like, bool) {
	o, ok := RowCacheIndex.Get("Like_PostId_2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	row, err := NewLike_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Like_PostId_2:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLike_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Like_PostId_2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewLike_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Like_PostId_2:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

// PhoneContact - phone_contacts_Id_pkey

// GeneralLog - general_log_Id_pkey

// FollowingListMember - following_list_member_Id_pkey

//field//field//field

///// Generated from index 'RefId'.
func (c _StoreImpl) Activity_ByRefId(RefId int) (*Activity, bool) {
	o, ok := RowCacheIndex.Get("Activity_RefId:" + fmt.Sprintf("%v", RefId))
	if ok {
		if obj, ok := o.(*Activity); ok {
			return obj, true
		}
	}

	row, err := NewActivity_Selector().RefId_Eq(RefId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Activity_RefId:"+fmt.Sprintf("%v", row.RefId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadActivity_ByRefIds(RefIds []int) {
	not_cached := make([]int, 0, len(RefIds))

	for _, id := range RefIds {
		_, ok := RowCacheIndex.Get("Activity_RefId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewActivity_Selector().RefId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Activity_RefId:"+fmt.Sprintf("%v", row.RefId), row, 0)
			}
		}
	}
}

// Media - media_Id_pkey

//field//field//field

///// Generated from index 'TargetId'.
func (c _StoreImpl) Notification_ByRowId(RowId int) (*Notification, bool) {
	o, ok := RowCacheIndex.Get("Notification_TargetId:" + fmt.Sprintf("%v", RowId))
	if ok {
		if obj, ok := o.(*Notification); ok {
			return obj, true
		}
	}

	row, err := NewNotification_Selector().RowId_Eq(RowId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Notification_TargetId:"+fmt.Sprintf("%v", row.RowId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadNotification_ByRowIds(RowIds []int) {
	not_cached := make([]int, 0, len(RowIds))

	for _, id := range RowIds {
		_, ok := RowCacheIndex.Get("Notification_TargetId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewNotification_Selector().RowId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Notification_TargetId:"+fmt.Sprintf("%v", row.RowId), row, 0)
			}
		}
	}
}

// PushMessage - push_message_PushMessageId_pkey

//field//field//field

///// Generated from index 'PhoneNormalizedNumber'.
func (c _StoreImpl) PhoneContact_ByPhoneNormalizedNumber(PhoneNormalizedNumber string) (*PhoneContact, bool) {
	o, ok := RowCacheIndex.Get("PhoneContact_PhoneNormalizedNumber:" + fmt.Sprintf("%v", PhoneNormalizedNumber))
	if ok {
		if obj, ok := o.(*PhoneContact); ok {
			return obj, true
		}
	}

	row, err := NewPhoneContact_Selector().PhoneNormalizedNumber_Eq(PhoneNormalizedNumber).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PhoneContact_PhoneNormalizedNumber:"+fmt.Sprintf("%v", row.PhoneNormalizedNumber), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContact_ByPhoneNormalizedNumbers(PhoneNormalizedNumbers []string) {
	not_cached := make([]string, 0, len(PhoneNormalizedNumbers))

	for _, id := range PhoneNormalizedNumbers {
		_, ok := RowCacheIndex.Get("PhoneContact_PhoneNormalizedNumber:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoneContact_Selector().PhoneNormalizedNumber_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PhoneContact_PhoneNormalizedNumber:"+fmt.Sprintf("%v", row.PhoneNormalizedNumber), row, 0)
			}
		}
	}
}

// MessageFile - message_file_MessageFileId_pkey

//field//field//field

///// Generated from index 'SessionUuid'.
func (c _StoreImpl) User_BySessionUuid(SessionUuid string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_SessionUuid:" + fmt.Sprintf("%v", SessionUuid))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().SessionUuid_Eq(SessionUuid).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_SessionUuid:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_BySessionUuids(SessionUuids []string) {
	not_cached := make([]string, 0, len(SessionUuids))

	for _, id := range SessionUuids {
		_, ok := RowCacheIndex.Get("User_SessionUuid:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().SessionUuid_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_SessionUuid:"+fmt.Sprintf("%v", row.SessionUuid), row, 0)
			}
		}
	}
}

// FollowingListMember - UserId_2

// Activity - activity_Id_pkey

// Room - room_RoomId_pkey

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) Session_ById(Id int) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSession_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("Session_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

// Post - post_Id_pkey

//field//field//field

///// Generated from index 'Phone'.
func (c _StoreImpl) User_ByPhone(Phone string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", Phone))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().Phone_Eq(Phone).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_Phone:"+fmt.Sprintf("%v", row.Phone), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByPhones(Phones []string) {
	not_cached := make([]string, 0, len(Phones))

	for _, id := range Phones {
		_, ok := RowCacheIndex.Get("User_Phone:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().Phone_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_Phone:"+fmt.Sprintf("%v", row.Phone), row, 0)
			}
		}
	}
}

// OldMsgPushEvent - old_msg_push_event_Id_pkey

// TagsPost - TagId

// Session - session_Id_pkey

// UserPassword - user_password_UserId_pkey

//field//field//field

///// Generated from index 'PhoneNumber'.
func (c _StoreImpl) PhoneContact_ByPhoneNumber(PhoneNumber string) (*PhoneContact, bool) {
	o, ok := RowCacheIndex.Get("PhoneContact_PhoneNumber:" + fmt.Sprintf("%v", PhoneNumber))
	if ok {
		if obj, ok := o.(*PhoneContact); ok {
			return obj, true
		}
	}

	row, err := NewPhoneContact_Selector().PhoneNumber_Eq(PhoneNumber).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PhoneContact_PhoneNumber:"+fmt.Sprintf("%v", row.PhoneNumber), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContact_ByPhoneNumbers(PhoneNumbers []string) {
	not_cached := make([]string, 0, len(PhoneNumbers))

	for _, id := range PhoneNumbers {
		_, ok := RowCacheIndex.Get("PhoneContact_PhoneNumber:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoneContact_Selector().PhoneNumber_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PhoneContact_PhoneNumber:"+fmt.Sprintf("%v", row.PhoneNumber), row, 0)
			}
		}
	}
}

// TestChat - UserId

// SearchClicked - search_clicked_Id_pkey

// DirectUpdate - direct_update_Id_pkey

// Group - group_GroupId_pkey

// LogChange - log_changes_Id_pkey

// OldMsgFile - old_msg_file_Id_pkey

// GroupToMessage - group_to_message_Id_pkey

//field//field//field

///// Generated from index 'UserId2'.
func (c _StoreImpl) UserMetaInfo_ByUserId(UserId int) (*UserMetaInfo, bool) {
	o, ok := RowCacheIndex.Get("UserMetaInfo_UserId2:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*UserMetaInfo); ok {
			return obj, true
		}
	}

	row, err := NewUserMetaInfo_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("UserMetaInfo_UserId2:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUserMetaInfo_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("UserMetaInfo_UserId2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUserMetaInfo_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("UserMetaInfo_UserId2:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// NotificationRemoved - notification_removed_NotificationId_pkey

// Chat - chat_ChatKey_pkey

//field//field//field

///// Generated from index 'PostId2'.
func (c _StoreImpl) Photo_ByPostId(PostId int) (*Photo, bool) {
	o, ok := RowCacheIndex.Get("Photo_PostId2:" + fmt.Sprintf("%v", PostId))
	if ok {
		if obj, ok := o.(*Photo); ok {
			return obj, true
		}
	}

	row, err := NewPhoto_Selector().PostId_Eq(PostId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Photo_PostId2:"+fmt.Sprintf("%v", row.PostId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoto_ByPostIds(PostIds []int) {
	not_cached := make([]int, 0, len(PostIds))

	for _, id := range PostIds {
		_, ok := RowCacheIndex.Get("Photo_PostId2:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoto_Selector().PostId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Photo_PostId2:"+fmt.Sprintf("%v", row.PostId), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'Name'.
func (c _StoreImpl) Tag_ByName(Name string) (*Tag, bool) {
	o, ok := RowCacheIndex.Get("Tag_Name:" + fmt.Sprintf("%v", Name))
	if ok {
		if obj, ok := o.(*Tag); ok {
			return obj, true
		}
	}

	row, err := NewTag_Selector().Name_Eq(Name).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Tag_Name:"+fmt.Sprintf("%v", row.Name), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTag_ByNames(Names []string) {
	not_cached := make([]string, 0, len(Names))

	for _, id := range Names {
		_, ok := RowCacheIndex.Get("Tag_Name:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewTag_Selector().Name_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Tag_Name:"+fmt.Sprintf("%v", row.Name), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'CreatedTime'.
func (c _StoreImpl) Photo_ByCreatedTime(CreatedTime int) (*Photo, bool) {
	o, ok := RowCacheIndex.Get("Photo_CreatedTime:" + fmt.Sprintf("%v", CreatedTime))
	if ok {
		if obj, ok := o.(*Photo); ok {
			return obj, true
		}
	}

	row, err := NewPhoto_Selector().CreatedTime_Eq(CreatedTime).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Photo_CreatedTime:"+fmt.Sprintf("%v", row.CreatedTime), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoto_ByCreatedTimes(CreatedTimes []int) {
	not_cached := make([]int, 0, len(CreatedTimes))

	for _, id := range CreatedTimes {
		_, ok := RowCacheIndex.Get("Photo_CreatedTime:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoto_Selector().CreatedTime_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Photo_CreatedTime:"+fmt.Sprintf("%v", row.CreatedTime), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'AlbumId'.
func (c _StoreImpl) Photo_ByAlbumId(AlbumId int) (*Photo, bool) {
	o, ok := RowCacheIndex.Get("Photo_AlbumId:" + fmt.Sprintf("%v", AlbumId))
	if ok {
		if obj, ok := o.(*Photo); ok {
			return obj, true
		}
	}

	row, err := NewPhoto_Selector().AlbumId_Eq(AlbumId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Photo_AlbumId:"+fmt.Sprintf("%v", row.AlbumId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoto_ByAlbumIds(AlbumIds []int) {
	not_cached := make([]int, 0, len(AlbumIds))

	for _, id := range AlbumIds {
		_, ok := RowCacheIndex.Get("Photo_AlbumId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoto_Selector().AlbumId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Photo_AlbumId:"+fmt.Sprintf("%v", row.AlbumId), row, 0)
			}
		}
	}
}

// User - user_Id_pkey

// GroupMessage - group_message_MessageId_pkey

// Activity - ActorUserId

// DirectMessage - direct_message_MessageId_pkey

//field//field//field

///// Generated from index 'MessageKey'.
func (c _StoreImpl) OldMessage_ByMessageKey(MessageKey string) (*OldMessage, bool) {
	o, ok := RowCacheIndex.Get("OldMessage_MessageKey:" + fmt.Sprintf("%v", MessageKey))
	if ok {
		if obj, ok := o.(*OldMessage); ok {
			return obj, true
		}
	}

	row, err := NewOldMessage_Selector().MessageKey_Eq(MessageKey).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("OldMessage_MessageKey:"+fmt.Sprintf("%v", row.MessageKey), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMessage_ByMessageKeys(MessageKeys []string) {
	not_cached := make([]string, 0, len(MessageKeys))

	for _, id := range MessageKeys {
		_, ok := RowCacheIndex.Get("OldMessage_MessageKey:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewOldMessage_Selector().MessageKey_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("OldMessage_MessageKey:"+fmt.Sprintf("%v", row.MessageKey), row, 0)
			}
		}
	}
}

//field//field//field

///// Generated from index 'UserName'.
func (c _StoreImpl) User_ByUserName(UserName string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_UserName:" + fmt.Sprintf("%v", UserName))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().UserName_Eq(UserName).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_UserName:"+fmt.Sprintf("%v", row.UserName), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByUserNames(UserNames []string) {
	not_cached := make([]string, 0, len(UserNames))

	for _, id := range UserNames {
		_, ok := RowCacheIndex.Get("User_UserName:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().UserName_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_UserName:"+fmt.Sprintf("%v", row.UserName), row, 0)
			}
		}
	}
}

// OldMsgPush - old_msg_push_Id_pkey

// Notification - ForUserId

// Like - likes_Id_pkey

// FollowingListMemberHistory - following_list_member_history_Id_pkey

// PushEvent - push_event_PushEventId_pkey

// Notification - notification_Id_pkey

//field//field//field

///// Generated from index 'HashMd5'.
func (c _StoreImpl) Photo_ByHashMd5(HashMd5 string) (*Photo, bool) {
	o, ok := RowCacheIndex.Get("Photo_HashMd5:" + fmt.Sprintf("%v", HashMd5))
	if ok {
		if obj, ok := o.(*Photo); ok {
			return obj, true
		}
	}

	row, err := NewPhoto_Selector().HashMd5_Eq(HashMd5).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Photo_HashMd5:"+fmt.Sprintf("%v", row.HashMd5), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoto_ByHashMd5s(HashMd5s []string) {
	not_cached := make([]string, 0, len(HashMd5s))

	for _, id := range HashMd5s {
		_, ok := RowCacheIndex.Get("Photo_HashMd5:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoto_Selector().HashMd5_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Photo_HashMd5:"+fmt.Sprintf("%v", row.HashMd5), row, 0)
			}
		}
	}
}

// Photo - photo_PhotoId_pkey

// UserMetaInfo - user_meta_info_Id_pkey

// PhoneContact - UserId_Time

//field//field//field

///// Generated from index 'Email'.
func (c _StoreImpl) User_ByEmail(Email string) (*User, bool) {
	o, ok := RowCacheIndex.Get("User_Email:" + fmt.Sprintf("%v", Email))
	if ok {
		if obj, ok := o.(*User); ok {
			return obj, true
		}
	}

	row, err := NewUser_Selector().Email_Eq(Email).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("User_Email:"+fmt.Sprintf("%v", row.Email), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadUser_ByEmails(Emails []string) {
	not_cached := make([]string, 0, len(Emails))

	for _, id := range Emails {
		_, ok := RowCacheIndex.Get("User_Email:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewUser_Selector().Email_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("User_Email:"+fmt.Sprintf("%v", row.Email), row, 0)
			}
		}
	}
}
