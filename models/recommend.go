package models

import (
	"math"
	"ms/sun/base"
	"ms/sun/ds"
	"ms/sun/helper"
	"time"
)

var TopUsers []int
var TopPosts []int

//////////////// Jobs /////////////////////
func Recommend_Jobs() {
	go Recommend_Job_TopUsers_Infinite()
	go Recommend_Job_TopPosts_Infinite()

}

func Recommend_Job_TopUsers_Infinite() {
	helper.JustRecover()

	helper.SleepForDebugDelay(5)

	for {
		TopUsers = Recommend_genTopUsers(50)
		time.Sleep(time.Minute * 30)
	}
}

func Recommend_Job_TopPosts_Infinite() {
	helper.JustRecover()

	helper.SleepForDebugDelay(-1)

	for {
		TopPosts = Recommend_GenTopPosts(500)
		time.Sleep(time.Minute * 30)
	}
}

//////////////// User //////////////////
func Recommend_Scheduler_GenUserRecomendations(ForUserId int) {
	m, ok := Store.GetUserMetaInfoByUserId(ForUserId)
	if ok {
		if m.LastUserRecGen+4*3600 < helper.TimeNow() {
			Recommend_GenPostsForUser_BG(ForUserId)
		}
	}

}

//////////////// Posts ///////////////////
func Recommend_GenPostsForUser_BG(ForUserId int) {
	go func() {
		defer helper.JustRecover()
		Recommend_ReGenPostsForUser(ForUserId)
	}()

}

func Recommend_ReGenPostsForUser(ForUserId int) {

}

func Recommend_GenTopPosts(limit int) []int {
	//EXPLAIN SELECT l.*, p.TypeId,COUNT(p.Id) AS Cnt FROM likes l JOIN post p ON p.Id = l.PostId  WHERE p.CreatedTime > 1477914190 GROUP BY p.Id ORDER BY cnt DESC LIMIT 500
	var ids []int

	last, err := NewPost_Selector().Select_Id().OrderBy_Id_Desc().Limit(1).GetInt(base.DB)
	XOLogErr(err)
	if err == nil {
		q := `SELECT p.Id FROM likes l JOIN post p ON p.Id = l.PostId  WHERE p.Id > ? AND p.TypeId = ? GROUP BY p.Id  ORDER BY COUNT(p.Id) DESC ,p.Id DESC LIMIT 500`
		err = base.DB.Select(&ids, q, last-20000, POST_TYPE_PHOTO)
		XOLogErr(err)
	}

	return ids
}

////////////////// Recom Users //////////////////

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
		Limit(200).
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

func Recommend_GenTopPosts__LEGACY(limit int) []int {
	rowsId, err := NewLike_Selector().Select_PostId().Limit(100000).OrderBy_Id_Desc().GetIntSlice(base.DB)
	if err != nil {
		return []int{}
	}

	rowsId, err = NewPost_Selector().Select_Id().Id_In(rowsId).TypeId_Eq(POST_TYPE_PHOTO).GetIntSlice(base.DB)
	if err != nil {
		return []int{}
	}

	mp := make(map[int]int, 100000) // map[FollowedUserId] = count
	for _, r := range rowsId {
		mp[r] += 1
	}

	coll := ds.New()
	for user, _ := range mp {
		coll.Add(user)
	}
	coll.SortDesc()

	min := math.Min(float64(coll.Size()), float64(limit))
	_ = min

	return coll.Values()
}
