package models

import (
	"math"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/ds"
	"ms/sun/helper"
	"time"
)

func Recommend_Jobs() {
	go Recommend_Job_TopUsers_Infinit()

}

func Recommend_GenPostsForUser_BG(ForUserId int) {
	go func() {
		defer helper.JustRecover()
		Recommend_ReGenPostsForUser(ForUserId)
	}()

}

func Recommend_ReGenPostsForUser(ForUserId int) {

}

func Recommend_GenUsersForUser_BG(ForUserId int) {
	go func() {
		defer helper.JustRecover()
		Recommend_ReGenUsersForUser(ForUserId)
	}()

}

func Recommend_ReGenUsersForUser(ForUserId int) {
	contacts_coll := Contacts_GetChachesContactsUserIdsForUserId(ForUserId, 0) //contacts UserIds list
	following_coll := MemoryStore.UserFollowingList_Get(ForUserId)

	notUserIds := following_coll.Values() //make([]int,0,len(following_coll) + len(contacts_coll))
	for _, id := range contacts_coll.Values() {
		if !following_coll.Contains(id) {
			notUserIds = append(notUserIds, id)
		}
	}

	//select those that they follows me but i don'f follow them
	toFollowRows, err := NewFollowingListMember_Selector().
		Select_UserId().
		FollowedUserId_Eq(ForUserId).
		UserId_NotIn(notUserIds).
		GetIntSlice(base.DB)

	if err != nil {
		return
	}

	toFollow := ds.New()
	toFollow.AddAndSort(toFollowRows...)

	for _, id := range contacts_coll.Values() {
		if !following_coll.Contains(id) { //if we don't follow this contacts already
			toFollow.Add(id)
		}
	}
	toFollow.SortDesc()

	var rows []RecommendUser
	for _, id := range toFollow.Values() {
		rows = append(rows, RecommendUser{
			UserId:      ForUserId,
			TargetId:    id,
			Weight:      0.5,
			CreatedTime: helper.TimeNow(),
		})
	}

	for _, id := range TopUsers {
		rows = append(rows, RecommendUser{
			UserId:      ForUserId,
			TargetId:    id,
			Weight:      0.2,
			CreatedTime: helper.TimeNow(),
		})
	}

	MassInsert_RecommendUser(rows, base.DB)

}

func Recommend_GetTopPosts(ForUserId int) {

}

var TopUsers []int

func Recommend_Job_TopUsers_Infinit() {
	helper.JustRecover()
	for {
		if config.DEBUG_DELAY_RUN_STARTUPS { //just don't make the log files messy for this at each startups
			time.Sleep(time.Minute * 5)
		}
		TopUsers = Recommend_genTopUsers(50)
	}
}

func Recommend_genTopUsers(cnt int) []int {
	//"select FollowedUserId as UserId, Count(*) AS Cnt form following_list_member where UpdatedTimeMs < ? group by FollowedUserId order by Cnt desc  limit 50  "
	rows, err := NewFollowingListMember_Selector().Limit(20000).OrderBy_Id_Desc().GetRows(base.DB)
	if err != nil {
		return []int{}
	}

	mp := make(map[int]int, 20000) // map[FollowedUserId] = count
	for _, r := range rows {
		mp[r.FollowedUserId] += 1
	}

	coll := ds.New()
	for user, _ := range mp {
		coll.Add(user)
	}
	coll.SortDesc()

	min := math.Min(float64(coll.Size()), 50)

	return coll.Values()[:int(min)]

}
