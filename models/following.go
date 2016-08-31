package models

import (
	// "sort"
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"strings"
)

func GetAllPrimiryFollowingIds(uid int) (ids []int) {
	// cashe.
	u := GetUserById(uid)
	key := "following-list-ids_" + intToStr(u.PrimaryFollowingList)
	if v, ok := cashe.Get(key); ok {
		return v.([]int)
	}
	base.DB.Select(&ids, "select FollowedUserId from following_list_member where ListId =?", u.PrimaryFollowingList)
	cashe.Set(key, ids, minutes(5))
	return
}

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

func FollowUser(uid int, followedId int) {
	if uid == followedId {
		return
	}
	u := GetUserById(uid)
	flm := FollowingListMember{}
	flm.FollowedUserId = followedId
	flm.UserId = uid
	//flm.FollowType = 1
	flm.ListId = u.PrimaryFollowingList
	flm.UpdatedTimeMs = helper.TimeNowMs()

	//todo check for duplicate
	changeListCountBy(u.PrimaryFollowingList, 1)
	//base.DbInsertStruct(&flm, "following_list_member")
	keys, values := helper.StructToFiledsRejectsEscape(&flm, "Id")
	q := "replace into following_list_member (" + strings.Join(keys, ",") + ") values (" + strings.Join(values, ",") + ")"
	fmt.Println(q)
	_, err := base.DB.Exec(q)
	fmt.Println(err)
	// cls, vals := structSqlNamesValues(flm)
	// DB.MustExec("insert ignore into following_list_member ("+cls+") "+ () + structSqlNamesValues(flm))
}

func UnfollowUser(uid int, followedId int) {
	u := GetUserById(uid)
	flm := FollowingListMember{}
	flm.FollowedUserId = followedId
	flm.UserId = uid
	//flm.FollowType = 0 //update
	flm.ListId = u.PrimaryFollowingList
	flm.UpdatedTimeMs = helper.TimeNowMs()
	//todo check for duplicate
	changeListCountBy(u.PrimaryFollowingList, 1)
	keys, values := helper.StructToFiledsRejectsEscape(&flm, "Id")
	q := "replace into following_list_member (" + strings.Join(keys, ",") + ") values (" + strings.Join(values, ",") + ")"
	//println(q)
	base.DB.Exec(q)
}

func UnfollowUser_dep(uid int, followedId int) {
	u := GetUserById(uid)
	base.DB.Exec("delete * from following_list_member where ListId=? AND FollowedUserId = ? ", u.PrimaryFollowingList, followedId)
	changeListCountBy(u.PrimaryFollowingList, -1)

	// flm := FollowingListMember{}
	// flm.FollowedUserId = followedId
	// flm.ListId = u.PrimaryFollowingList

	//todo check for duplicate
	// DbInsertStruct(&flm, "following_list_member")
}

//todo to implment
func changeListCountBy(listId, counter int) {
	//to
}

func IsUserFollowing(userId, followedUserId int) bool {
	user := GetUserById(userId)
	var allFollowedUserIds []int

	key := "all-following-user-ids-of-" + intToStr(userId)
	if v, ok := cashe.Get(key); ok {
		allFollowedUserIds = v.([]int)
	} else {
		base.DB.Select(&allFollowedUserIds, "select FollowedUserId from following_list_member where ListId = ? order by FollowedUserId ASC ", user.PrimaryFollowingList)
		cashe.Set(key, allFollowedUserIds, 0)
	}
	devPrintn("allFollowedUserIds: ", allFollowedUserIds)

	// index := sort.SearchInts(allFollowedUserIds, followedUserId)
	// if index >= 2

	//todo use binary search
	for _, v := range allFollowedUserIds {
		if followedUserId == v {
			return true
		}
	}

	return false
}

//deprecated
func GetFollowingType(iUserId, peerUserId int) int {
	var typ int
	q := "select FollowType from following_list_member where UserId = ? and FollowedUserId = ? "
	base.DB.Get(&typ, q, iUserId, peerUserId)
	return typ
}

func AddFollowingsToUserList(user []User, iUserId int) {

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
			userView.IFollowType = UserMemoryStore.GetFollowingTypeForUsers(cuid, uid)
			//debug("GetFollowingTypeForUsers: ", UserMemoryStore.GetFollowingTypeForUsers(cuid, uid))
		}
		userListView = append(userListView, userView)
	}
	return userListView
}

/*
func UsersListForFollowView(userIds []int, cuid int) []UserInlineFollowView {
    var userListView []UserInlineFollowView
    for _, uid := range userIds {
        userView := UserInlineFollowView{}
        userView.UserInlineView = GetUserView(uid)
        if cuid > 0 {
            //userView.AmIFollowing = IsUserFollowing(cuid, userView.UserId)
            userView.IFollowType = UserMemoryStore.GetFollowingTypeForUsers(cuid, uid)
            //debug("GetFollowingTypeForUsers: ", UserMemoryStore.GetFollowingTypeForUsers(cuid, uid))
        }
        userListView = append(userListView, userView)
    }
    return userListView
}*/
