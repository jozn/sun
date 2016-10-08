package models

import (
	"fmt"
	c "github.com/patrickmn/go-cache"
	"ms/sun/base"
	"ms/sun/ds"
	"time"
)

type _memoryStoreImpl int

var MemoryStore _memoryStoreImpl
var Cacher *c.Cache

func init() {
	Cacher = c.New(time.Second*4*3600, time.Second*60)
}

func (e _memoryStoreImpl) GetPhoneForUserIfIsContact(CurrentUserId, UserId int) string {
	phones, err := NewPhoneContact_Selector().Select_PhoneNormalizedNumber().UserId_EQ(CurrentUserId).GetStringSlice(base.DB)
	if err != nil {
		return ""
	}
	users, err := NewUser_Selector().Select_Id().Phone_In(phones).OrderBy_Id_Desc().GetIntSlice(base.DB)
	if err != nil {
		return ""
	}
	collection := ds.New()
	collection.AddAndSort(users...)
	if collection.Contains(UserId) {
		u := UserMemoryStore.GetForUser(UserId)
		if u != nil {
			return u.Phone
		}
	}
	key := fmt.Sprintf("UserContacts:%d", CurrentUserId)
	t := time.Hour * 4
	Cacher.Replace(key, collection, t)

	return ""
}
