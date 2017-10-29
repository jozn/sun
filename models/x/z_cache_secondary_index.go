package x

import (
	"fmt"
	"ms/sun/base"
)

// Activity - PRIMARY

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

// Activity - ActorUserId

// Bucket - PRIMARY

// Chat - PRIMARY

// Comment - PRIMARY

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

// DirectMessage - PRIMARY

// DirectOffline - PRIMARY

//field//field//field

///// Generated from index 'ToUserId'.
func (c _StoreImpl) DirectOffline_ByToUserId(ToUserId int) (*DirectOffline, bool) {
	o, ok := RowCacheIndex.Get("DirectOffline_ToUserId:" + fmt.Sprintf("%v", ToUserId))
	if ok {
		if obj, ok := o.(*DirectOffline); ok {
			return obj, true
		}
	}

	row, err := NewDirectOffline_Selector().ToUserId_Eq(ToUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("DirectOffline_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectOffline_ByToUserIds(ToUserIds []int) {
	not_cached := make([]int, 0, len(ToUserIds))

	for _, id := range ToUserIds {
		_, ok := RowCacheIndex.Get("DirectOffline_ToUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewDirectOffline_Selector().ToUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("DirectOffline_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
			}
		}
	}
}

// DirectOfflineDep - PRIMARY

//field//field//field

///// Generated from index 'ToUserId'.
func (c _StoreImpl) DirectOfflineDep_ByToUserId(ToUserId int) (*DirectOfflineDep, bool) {
	o, ok := RowCacheIndex.Get("DirectOfflineDep_ToUserId:" + fmt.Sprintf("%v", ToUserId))
	if ok {
		if obj, ok := o.(*DirectOfflineDep); ok {
			return obj, true
		}
	}

	row, err := NewDirectOfflineDep_Selector().ToUserId_Eq(ToUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("DirectOfflineDep_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectOfflineDep_ByToUserIds(ToUserIds []int) {
	not_cached := make([]int, 0, len(ToUserIds))

	for _, id := range ToUserIds {
		_, ok := RowCacheIndex.Get("DirectOfflineDep_ToUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewDirectOfflineDep_Selector().ToUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("DirectOfflineDep_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
			}
		}
	}
}

// DirectToMessage - PRIMARY

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) DirectToMessage_ById(Id int) (*DirectToMessage, bool) {
	o, ok := RowCacheIndex.Get("DirectToMessage_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*DirectToMessage); ok {
			return obj, true
		}
	}

	row, err := NewDirectToMessage_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("DirectToMessage_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectToMessage_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("DirectToMessage_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewDirectToMessage_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("DirectToMessage_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

// DirectUpdate - PRIMARY

//field//field//field

///// Generated from index 'ToUserId'.
func (c _StoreImpl) DirectUpdate_ByToUserId(ToUserId int) (*DirectUpdate, bool) {
	o, ok := RowCacheIndex.Get("DirectUpdate_ToUserId:" + fmt.Sprintf("%v", ToUserId))
	if ok {
		if obj, ok := o.(*DirectUpdate); ok {
			return obj, true
		}
	}

	row, err := NewDirectUpdate_Selector().ToUserId_Eq(ToUserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("DirectUpdate_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadDirectUpdate_ByToUserIds(ToUserIds []int) {
	not_cached := make([]int, 0, len(ToUserIds))

	for _, id := range ToUserIds {
		_, ok := RowCacheIndex.Get("DirectUpdate_ToUserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewDirectUpdate_Selector().ToUserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("DirectUpdate_ToUserId:"+fmt.Sprintf("%v", row.ToUserId), row, 0)
			}
		}
	}
}

// FollowingList - PRIMARY

// FollowingListMember - PRIMARY

// FollowingListMember - UserId

// FollowingListMember - FollowedUserId

// FollowingListMember - UserId_2

// FollowingListMemberHistory - PRIMARY

// FollowingListMemberHistory - UserId

// GeneralLog - PRIMARY

// Group - PRIMARY

// GroupMember - PRIMARY

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) GroupMember_ById(Id int) (*GroupMember, bool) {
	o, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*GroupMember); ok {
			return obj, true
		}
	}

	row, err := NewGroupMember_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("GroupMember_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadGroupMember_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("GroupMember_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewGroupMember_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("GroupMember_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

// GroupMessage - PRIMARY

// GroupToMessage - PRIMARY

// Like - PRIMARY

// Like - PostId

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) Like_ById(Id int) (*Like, bool) {
	o, ok := RowCacheIndex.Get("Like_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*Like); ok {
			return obj, true
		}
	}

	row, err := NewLike_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Like_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadLike_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("Like_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewLike_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Like_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

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

// LogChange - PRIMARY

// Media - PRIMARY

// MessageFile - PRIMARY

// Msg - PRIMARY

// Notification - PRIMARY

// Notification - ForUserId

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

// NotificationRemoved - PRIMARY

// Offline - PRIMARY

// OldMessage - PRIMARY

//field//field//field

///// Generated from index 'Uid'.
func (c _StoreImpl) OldMessage_ByUid(Uid int) (*OldMessage, bool) {
	o, ok := RowCacheIndex.Get("OldMessage_Uid:" + fmt.Sprintf("%v", Uid))
	if ok {
		if obj, ok := o.(*OldMessage); ok {
			return obj, true
		}
	}

	row, err := NewOldMessage_Selector().Uid_Eq(Uid).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("OldMessage_Uid:"+fmt.Sprintf("%v", row.Uid), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadOldMessage_ByUids(Uids []int) {
	not_cached := make([]int, 0, len(Uids))

	for _, id := range Uids {
		_, ok := RowCacheIndex.Get("OldMessage_Uid:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewOldMessage_Selector().Uid_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("OldMessage_Uid:"+fmt.Sprintf("%v", row.Uid), row, 0)
			}
		}
	}
}

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

// OldMessage - UserId

// OldMsgFile - PRIMARY

// OldMsgPush - PRIMARY

// OldMsgPush - ToUser

// OldMsgPushEvent - PRIMARY

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

// PhoneContact - PRIMARY

// PhoneContact - PhoneContactRowId

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

//field//field//field

///// Generated from index 'UserId'.
func (c _StoreImpl) PhoneContact_ByUserId(UserId int) (*PhoneContact, bool) {
	o, ok := RowCacheIndex.Get("PhoneContact_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*PhoneContact); ok {
			return obj, true
		}
	}

	row, err := NewPhoneContact_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PhoneContact_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPhoneContact_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("PhoneContact_UserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPhoneContact_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PhoneContact_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// PhoneContact - UserId_Time

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

// Photo - PRIMARY

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

// Post - PRIMARY

//field//field//field

///// Generated from index 'UserId'.
func (c _StoreImpl) Post_ByUserId(UserId int) (*Post, bool) {
	o, ok := RowCacheIndex.Get("Post_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*Post); ok {
			return obj, true
		}
	}

	row, err := NewPost_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Post_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPost_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("Post_UserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPost_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Post_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// PushEvent - PRIMARY

// PushMessage - PRIMARY

// PushMessage - ToUser

//field//field//field

///// Generated from index 'Uid'.
func (c _StoreImpl) PushMessage_ByPushMessageId(PushMessageId int) (*PushMessage, bool) {
	o, ok := RowCacheIndex.Get("PushMessage_Uid:" + fmt.Sprintf("%v", PushMessageId))
	if ok {
		if obj, ok := o.(*PushMessage); ok {
			return obj, true
		}
	}

	row, err := NewPushMessage_Selector().PushMessageId_Eq(PushMessageId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("PushMessage_Uid:"+fmt.Sprintf("%v", row.PushMessageId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadPushMessage_ByPushMessageIds(PushMessageIds []int) {
	not_cached := make([]int, 0, len(PushMessageIds))

	for _, id := range PushMessageIds {
		_, ok := RowCacheIndex.Get("PushMessage_Uid:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewPushMessage_Selector().PushMessageId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("PushMessage_Uid:"+fmt.Sprintf("%v", row.PushMessageId), row, 0)
			}
		}
	}
}

// RecommendUser - PRIMARY

// Room - PRIMARY

// SearchClicked - PRIMARY

// Session - PRIMARY

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

//field//field//field

///// Generated from index 'UserId'.
func (c _StoreImpl) Session_ByUserId(UserId int) (*Session, bool) {
	o, ok := RowCacheIndex.Get("Session_UserId:" + fmt.Sprintf("%v", UserId))
	if ok {
		if obj, ok := o.(*Session); ok {
			return obj, true
		}
	}

	row, err := NewSession_Selector().UserId_Eq(UserId).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("Session_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadSession_ByUserIds(UserIds []int) {
	not_cached := make([]int, 0, len(UserIds))

	for _, id := range UserIds {
		_, ok := RowCacheIndex.Get("Session_UserId:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewSession_Selector().UserId_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("Session_UserId:"+fmt.Sprintf("%v", row.UserId), row, 0)
			}
		}
	}
}

// SettingClient - PRIMARY

// SettingNotification - PRIMARY

// Tag - PRIMARY

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

// TagsPost - PRIMARY

// TagsPost - TagId

// TestChat - PRIMARY

// TestChat - UserId

// TriggerLog - PRIMARY

//field//field//field

///// Generated from index 'Id'.
func (c _StoreImpl) TriggerLog_ById(Id int) (*TriggerLog, bool) {
	o, ok := RowCacheIndex.Get("TriggerLog_Id:" + fmt.Sprintf("%v", Id))
	if ok {
		if obj, ok := o.(*TriggerLog); ok {
			return obj, true
		}
	}

	row, err := NewTriggerLog_Selector().Id_Eq(Id).GetRow(base.DB)
	if err == nil {
		RowCacheIndex.Set("TriggerLog_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
		return row, true
	}

	XOLogErr(err)
	return nil, false
}

func (c _StoreImpl) PreLoadTriggerLog_ByIds(Ids []int) {
	not_cached := make([]int, 0, len(Ids))

	for _, id := range Ids {
		_, ok := RowCacheIndex.Get("TriggerLog_Id:" + fmt.Sprintf("%v", id))
		if !ok {
			not_cached = append(not_cached, id)
		}
	}

	if len(not_cached) > 0 {
		rows, err := NewTriggerLog_Selector().Id_In(not_cached).GetRows(base.DB)
		if err == nil {
			for _, row := range rows {
				RowCacheIndex.Set("TriggerLog_Id:"+fmt.Sprintf("%v", row.Id), row, 0)
			}
		}
	}
}

// User - PRIMARY

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

// User - Email_2

// UserMetaInfo - PRIMARY

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

// UserPassword - PRIMARY