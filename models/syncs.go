package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/ds"
	"time"
)

//TODO: must add update users as well
func SyncGetAllChangedUser(CurrentUserId, LastTime int) (res []*UserSyncAndMeView) {
	var contactsUsers []int
	var followedUsers []int

	sel := NewPhoneContact_Selector().Select_PhoneNormalizedNumber().
		UserId_Eq(CurrentUserId).
		PhoneNormalizedNumber_NotEq("")

	if LastTime > 0 {
		sel.CreatedTime_GE(LastTime)
	}

	phones, err := sel.GetStringSlice(base.DB)
	if err != nil {
		return
	}

	if len(phones) > 0 {
		contactsUsers, err = NewUser_Selector().Select_Id().Phone_In(phones).OrderBy_Id_Desc().GetIntSlice(base.DB)
		if err != nil {
			return
		}
		collection := ds.New()
		collection.AddAndSort(contactsUsers...)
		//this is used in MemoryStore.GetPhoneForUserIfIsContact()
		Cacher.Set(fmt.Sprintf("UserContacts:%d", CurrentUserId), collection, time.Hour*4)
	}

	sel2 := NewFollowingListMember_Selector().Select_FollowedUserId().
		UserId_Eq(CurrentUserId)

	if LastTime > 0 {
		sel2.UpdatedTimeMs_GE(LastTime)
	}
	followedUsers, err = sel2.GetIntSlice(base.DB)

	//can improve use a map to not duplicate res
	for _, u := range contactsUsers {
		res = append(res, Views.UserViewSync(CurrentUserId, u))
	}

	for _, u := range followedUsers {
		res = append(res, Views.UserViewSync(CurrentUserId, u))
	}
	return
}
