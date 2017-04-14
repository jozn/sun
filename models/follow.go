package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

/*
// In Orma-gen
type FollowingList struct {
	Id          int
	UserId      int
	ListType    int
	Name        string
	Count       int
	IsAuto      int
	IsPimiry    int
	CreatedTime int
	// xo fields
	_exists, _deleted bool
}

//following_list_member
type FollowingListMember struct {
	Id             int
	ListId         int
	UserId         int
	FollowedUserId int
	FollowType     int
	UpdatedTimeMs  int
	// xo fields
	_exists, _deleted bool
}
*/

/* ?
//table: following_list_member_history
type FollowingListMemberHistory struct {
	FollowingListMember
	FollowId int
	// xo fields
	_exists, _deleted bool
}
*/

func Follow(UserId, FollowedPeerUserId int) int {
	if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
		return 0
	}

	flm := FollowingListMember{
		ListId:         UserId,
		UserId:         UserId,
		FollowedUserId: FollowedPeerUserId,
		FollowType:     1, //remove
		UpdatedTimeMs:  helper.TimeNowMs(),
	}

	err := flm.Insert(base.DB)

	if err == nil {
		MemoryStore.UserFollowingList_Add(UserId, FollowedPeerUserId)
		Counter.UpdateUserFollowingCounts(UserId, 1)
		Counter.UpdateUserFollowersCounts(FollowedPeerUserId, 1)

		Notification_OnFollowed(UserId, FollowedPeerUserId)
		Activity_OnFollowed(UserId, FollowedPeerUserId, flm.Id)
		return 1
	}
	return 0
}

func UnFollow(UserId, FollowedPeerUserId int) {
	if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
		return
	}

	MemoryStore.UserFollowingList_Remove(UserId, FollowedPeerUserId)
	flm, err := NewFollowingListMember_Selector().UserId_EQ(UserId).FollowedUserId_EQ(FollowedPeerUserId).GetRow(base.DB)
	if err != nil {
		return
	}

	err = flm.Delete(base.DB)

	if err == nil {
		Counter.UpdateUserFollowingCounts(UserId, -1)
		Counter.UpdateUserFollowersCounts(FollowedPeerUserId, -1)

		Notification_OnUnFollowed(UserId, FollowedPeerUserId)
		Activity_OnUnFollowed(UserId, FollowedPeerUserId, flm.Id)
	}
	//n, err := NewFollowingListMember_Deleter().UserId_EQ(UserId).FollowedUserId_EQ(FollowedPeerUserId).Delete(base.DB)
}
