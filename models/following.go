package models

import (
	// "sort"
	"ms/sun/base"
	"ms/sun/helper"
)

func GetAllFollowingsUserIds(userId int, lastTimestamp int) (ids []int) {
	userIds := []int{}
	q := "select FollowedUserId from following_list_member where UserId = ? and FollowType = 1 "
	if lastTimestamp > 0 {
		q = q + " And UpdatedTimestampMs >= " + helper.IntToStr(lastTimestamp)
	}
	base.DB.Select(&userIds, q, userId)
	return userIds
}

func GetAllFollowingsUser(userId int, lastTimestamp int) []UserTable {
	userIds := GetAllFollowingsUserIds(userId, lastTimestamp)
	users := []UserTable{}
	q := "select * from user where Id in(" + helper.IntsToSqlIn(userIds) + ")"
	base.DB.Select(&users, q)
	return users
}

func GetAllUnFollowedUserIds(userId int, lastTimestamp int) []int {
	userIds := []int{}
	q := "select FollowedUserId from following_list_member where UserId = ? and FollowType = 0 "
	if lastTimestamp > 0 {
		q = q + " And UpdatedTimestampMs >= " + helper.IntToStr(lastTimestamp)
	}
	base.DB.Select(&userIds, q, userId)
	return userIds
}

//deprecated
func GetFollowingType(iUserId, peerUserId int) int {
	var typ int
	q := "select FollowType from following_list_member where UserId = ? and FollowedUserId = ? "
	base.DB.Get(&typ, q, iUserId, peerUserId)
	return typ
}

type userfollowingsList map[int]int // map[user_id] follow_type

func PreloadFollowingsListTypesForUser(userId int) (list userfollowingsList) {
	//DB.Select
	list = make(map[int]int)
	//var userIds []int
	//userIds := make([]int,0)
	userIds := []int{}
	q := "select FollowedUserId from following_list_member where UserId = ? " //and FollowType = 1 "
	//println(userIds)
	base.DB.Select(&userIds, q, userId)

	for _, uid := range userIds {
		list[uid] = 1
		//list = append(list,uid)
	}
	return
}

func (list userfollowingsList) IsFollowing(peerUserId int) bool {
	u := list[peerUserId]
	if u == 1 {
		return true
	}
	return false
}

//TODO: implement 2:follow_requested
/// 0: not_following  1: following  2: follow_requested
func (list userfollowingsList) FollowingType(peerUserId int) int {
	u := list[peerUserId]
	if u == 1 {
		return 1 //
	}
	return 0
}

type FollowChecker interface {
	FollowingType(int) int
	IsFollowing(int) bool
}

////////////////////////////////// new Apis 0.4 ///////////////////////

func UsersToInlineFollowView(userIds []int, cuid int) []UserInlineFollowView {
	var userListView []UserInlineFollowView
	for _, uid := range userIds {
		userView := UserInlineFollowView{}
		userView.UserInlineView = GetUserView(uid)
		if cuid > 0 {
			//userView.AmIFollowing = IsUserFollowing(cuid, userView.UserId)
			userView.IFollowType = MemoryStore.UserFollowingList_GetFollowingTypeForUsers(cuid, uid)
			//debug("GetFollowingTypeForUsers: ", MemoryStore.UserFollowingList_GetFollowingTypeForUsers(cuid, uid))
		}
		userListView = append(userListView, userView)
	}
	return userListView
}

