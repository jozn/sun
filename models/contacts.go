package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/ds"
	"ms/sun/models/x"
	"time"
)

func Contacts_GetChachesContactsUserIdsForUserId(UserId, LastTime int) *ds.IntList {
	key := fmt.Sprintf("UserContacts:%d", UserId) //in Views.GetPhoneForUserIfIsContact()

	val, ok := Cacher.Get(key)
	if ok {
		collection, ok := val.(*ds.IntList)
		if ok {
			return collection
		}
	}

	var contactsUsers []int
	collection := ds.New()

	sel := x.NewPhoneContact_Selector().Select_PhoneNormalizedNumber().
		UserId_Eq(UserId).
		PhoneNormalizedNumber_NotEq("")

	if LastTime > 0 {
		sel.CreatedTime_GE(LastTime)
	}

	phones, err := sel.GetStringSlice(base.DB)
	if err != nil {
		return collection
	}

	if len(phones) > 0 {
		contactsUsers, err = x.NewUser_Selector().Select_Id().Phone_In(phones).OrderBy_Id_Desc().GetIntSlice(base.DB)
		if err != nil {
			return collection
		}

		collection.AddAndSort(contactsUsers...)
		//this is used in MemoryStore.GetPhoneForUserIfIsContact()
		Cacher.Set(key, collection, time.Hour*4)
	}
	return collection
}
