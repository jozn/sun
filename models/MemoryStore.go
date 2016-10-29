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
	key := fmt.Sprintf("UserContacts:%d", CurrentUserId)
	var collection *ds.IntList

	val, ok := Cacher.Get(key)
	if !ok {
		phones, err := NewPhoneContact_Selector().Select_PhoneNormalizedNumber().UserId_EQ(CurrentUserId).GetStringSlice(base.DB)
		if err != nil {
			return ""
		}

		collection = ds.New()
		if len(phones) > 0 {
			users, err := NewUser_Selector().Select_Id().Phone_In(phones).OrderBy_Id_Desc().GetIntSlice(base.DB)
			if err != nil {
				return ""
			}
			collection.AddAndSort(users...)
		}
		t := time.Hour * 4
		Cacher.Set(key, collection, t)
	} else {
		collection = val.(*ds.IntList)
	}

	//vv := *ds.IntList(val)
	if collection.Contains(UserId) {
		u := UserMemoryStore.GetForUser(UserId)
		if u != nil {
			return u.Phone
		}
	}
	return ""
}

func (e _memoryStoreImpl) UserLikedPostsList_Get(UserId int) *ds.IntList {
	key := fmt.Sprintf("UserLikePosts:%d", UserId)
	var collection *ds.IntList

	val, ok := Cacher.Get(key)
	if !ok {
		collection = ds.New()

		postsLiked, err := NewLike_Selector().Select_PostId().UserId_EQ(UserId).GetIntSlice(base.DB)
		if err != nil {
			return collection
		}

		for _, pid := range postsLiked {
			collection.Add(pid)
		}
		collection.TrySortDesc()

		t := time.Hour * 4
		Cacher.Set(key, collection, t)
	} else {
		collection = val.(*ds.IntList)
	}

	return collection
}

func (e _memoryStoreImpl) UserLikedPostsList_Add(UserId int, PostId int) {
	e.UserLikedPostsList_Get(UserId).AddAndSort(PostId)
}

func (e _memoryStoreImpl) UserLikedPostsList_Remove(UserId int, PostId int) {
	e.UserLikedPostsList_Get(UserId).RemoveAndSort(PostId)
}

func (e _memoryStoreImpl) UserLikedPostsList_IsLiked(UserId int, PostId int) bool {
	return e.UserLikedPostsList_Get(UserId).Contains(PostId)
}
