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

/////////////////////////// Phone ////////////////////////////////////

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
		u,ok  := MemoryStore_User.GetUser(UserId)
        if ok {
            return u.Phone
        }
	}
	return ""
}

////////////////////////// UserLikedPostsList ////////////////////////////

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

func (e _memoryStoreImpl) UserLikedPostsList_Invalidate(UserId int) {
	key := fmt.Sprintf("UserLikePosts:%d", UserId)
	Cacher.Delete(key)
}

////////////////////////// UserFollowingList ////////////////////////////
func (e _memoryStoreImpl) UserFollowingList_Get(UserId int) *ds.IntList {
	key := fmt.Sprintf("UserFollowing:%d", UserId)
	var collection *ds.IntList

	val, ok := Cacher.Get(key)
	if !ok {
		collection = ds.New()

		followed, err := NewFollowingListMember_Selector().Select_FollowedUserId().UserId_EQ(UserId).OrderBy_FollowedUserId_Desc().GetIntSlice(base.DB)
		if err != nil {
			return collection
		}

		/*for _, pid := range followed {
		    collection.Add(pid)
		}*/
		collection.Add(followed...)
		collection.TrySortDesc()

		t := time.Hour * 4
		Cacher.Set(key, collection, t)
	} else {
		collection = val.(*ds.IntList)
	}

	return collection
}

func (e _memoryStoreImpl) UserFollowingList_Add(UserId int, PostId int) {
	e.UserFollowingList_Get(UserId).AddAndSort(PostId)
}

func (e _memoryStoreImpl) UserFollowingList_Remove(UserId int, FollowedUserId int) {
	e.UserFollowingList_Get(UserId).RemoveAndSort(FollowedUserId)
}

func (e _memoryStoreImpl) UserFollowingList_GetFollowingTypeForUsers(UserId int, PostId int) int {
	l := e.UserFollowingList_Get(UserId)
	ftype := 0
	if l.BinaryContains(PostId) {
		//if s.Followings.Contains(ReqFollowedUserId){
		ftype = 1
	}
	//todo add FollowingsRequested funcs
	/*else if l.FollowingsRequested.BinaryContains(ReqFollowedUserId) {
	    ftype = 2
	}*/

	return ftype
}

func (e _memoryStoreImpl) UserFollowingList_Invalidate(UserId int) {
	key := fmt.Sprintf("UserFollowing:%d", UserId)
	Cacher.Delete(key)
}

////////////////////////////////////////////////////////////////
