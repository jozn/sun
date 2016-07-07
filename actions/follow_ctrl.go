package actions

import (
	. "ms/sun/base"
	. "ms/sun/models"

	"strconv"
	"ms/sun/models"
	"ms/sun/helper"
)

func PingAction(c *Action) {
	c.SendText("PONG")
}
func FollowAction(c *Action) {
	devPrintn("Action follow")
	c.MustUser()
	c.MustPost()
	cuid := c.UserId()

	id := (c.Req.Form["user_id"])
	userId, err := strconv.Atoi(id[0])
	noErr(err)
	user := GetUserById(cuid)
	fm := FollowingListMember{}
	fm.ListId = user.PrimaryFollowingList
	fm.UserId = user.Id //cuid
	fm.FollowedUserId = userId
	fm.UpdatedTimestampMs = helper.TimeNowMs()
	//	fm.
	DbInsertStruct(&fm, "following_list_member")
	// c.SendText("ok")

	c.SendJson(fm)

}

func UnfollowAction(c *Action) {
	devPrintn("Action follow")
	c.MustUser()
	c.MustPost()
	cuid := c.UserId()

	id := (c.Req.Form["user_id"])
	userId, err := strconv.Atoi(id[0])
	noErr(err)
	user := GetUserById(cuid)

	DB.MustExec("delete from following_list_member where ListId = ? and FollowedUserId = ?", user.PrimaryFollowingList, userId)

	c.SendText("ok")

}

//followers of a user
func GetFollowersAction(c *Action) {
	devPrintn("getFollowersAction")
	//c.MustUser()
	//c.MustPost()
	cuid := c.UserId()

	username := (c.Req.Form["username"])[0]
	//userId, err := strconv.Atoi(id[0])
	//noErr(err)
	user := GetUserByUsername(username)

	var followedIds []int
	//todo make mysql uniqe select
	DB.Select(&followedIds, "select UserId from following_list_member where FollowedUserId = ?", user.Id)

	usersFollow := usersToInlineFollowView(followedIds, cuid)

	c.SendJson(usersFollow)
}

//followers of a user
func GetFollowingsAction(c *Action) {
	devPrintn("getFollowersAction")
	cuid := c.UserId()

	username := (c.Req.Form["username"])[0]
	user := GetUserByUsername(username)

	//FOLLOWINGS
	var followedIds []int
	//todo make mysql uniqe select
	DB.Select(&followedIds, "select FollowedUserId from following_list_member where UserId = ?", user.Id)

	usersFollow := usersToInlineFollowView(followedIds, cuid)

	c.SendJson(usersFollow)
}

func usersToInlineFollowView(userIds []int, cuid int) []UserInlineFollowView {
	var userListView []UserInlineFollowView
	for _, uid := range userIds {
		userView := UserInlineFollowView{}
		userView.UserInlineView = GetUserView(uid)
		if cuid > 0 {
			userView.AmIFollowing = IsUserFollowing(cuid, userView.UserId)
		}
		userListView = append(userListView, userView)
	}
	return userListView
}

type SyncFollowings struct {
	Add []UserBasicAndMe
	Remove []int
}

func SyncFollowingsAction(a *Action) AppErr  {
	last_str := a.Req.Form.Get("last")//last TimeStamp
	last := helper.StrToInt(last_str,0)

	users:=models.GetAllFollowingsUser(a.UserId(),last)

	list := models.PreloadFollowingsListTypesForUser(a.UserId())
	followings_res := []UserBasicAndMe{}
	for _,u := range users {
		cu := models.UserBasicAndMe{}
		cu.FromUser(u,list)
		//cu.UserBasic = u.UserBasic
		//cu.UpdatedTimestamp = u.UpdatedTimestamp
		//cu.FollowingType = list.FollowingType(u.Id)
		followings_res = append(followings_res, cu)
	}

	sync_res := SyncFollowings{
		Add: followings_res,
		Remove: models.GetAllUnFollowedUserIds(a.UserId(),last),
	}
	a.SendJson(sync_res)
	return nil
}
