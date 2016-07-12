package ctrl

import (
	 "ms/sun/base"
	"ms/sun/models"

	"ms/sun/helper"
	//"ms/sun/docs/del"
)

func PingAction(c *base.Action) base.AppErr {
	c.SendText("PONG")
	return nil
}

func FollowAction(c *base.Action)  base.AppErr{
	cuid := c.UserId()

	//fid := c.Req.Form.Get("followed_user_id")
	fid := c.Req.Form.Get("followed_user_id")
	fUserId := helper.StrToInt(fid,-1)
	if fUserId < 1 {
		return nil
	}
	models.UserMemoryStore.AddFollow(cuid, fUserId)
	models.UserMemoryStore.UpdateUserFollowingCounts(cuid,1)
	return nil
}


func UnfollowAction(c *base.Action) base.AppErr {
	cuid := c.UserId()

	fid := c.Req.Form.Get("followed_user_id")
	fUserId := helper.StrToInt(fid,-1)
	if fUserId < 1 {
		return nil
	}
	models.UserMemoryStore.RemoveFollow(cuid, fUserId)
	models.UserMemoryStore.UpdateUserFollowingCounts(cuid,-1)
	c.SendText("ok")
	return nil
}

const FOLLOW_LIST_SHOW_LIMIT = 50
//followers of a user
func GetFollowersListAction(c *base.Action) base.AppErr{
	//devPrintn("getFollowersAction")
	//c.MustUser()
	//c.MustPost()
	cuid := c.UserId()

	username := c.Req.Form.Get("username")
	pageStr := c.Req.Form.Get("page")
	page := helper.StrToInt(pageStr,0)
	offset := FOLLOW_LIST_SHOW_LIMIT * page
	//userId, err := strconv.Atoi(id[0])
	//noErr(err)
	user,err := models.GetUserByUsername2(username)
	if err != nil{
		c.Protocol.Status = "ERR"
		c.Protocol.Error = " کاربر پیدا نشد "
		return nil
	}
	var followedIds []int
	//todo make mysql uniqe select
	err=base.DB.Select(&followedIds, "select UserId from following_list_member where FollowedUserId = ? order by Id DESC limit ? offset ? ", user.Id, FOLLOW_LIST_SHOW_LIMIT, offset)
	if err != nil {
		helper.DebugErr(err)
		return nil
	}
	usersFollow := models.UsersToInlineFollowView(followedIds, cuid)

	c.SendJson(usersFollow)
	return nil
}

func GetFollowingsListAction(c *base.Action) base.AppErr{
	//devPrintn("getFollowersAction")
	//c.MustUser()
	//c.MustPost()
	cuid := c.UserId()

	username := c.Req.Form.Get("username")
	pageStr := c.Req.Form.Get("page")
	page := helper.StrToInt(pageStr,0)
	offset := FOLLOW_LIST_SHOW_LIMIT * page
	//userId, err := strconv.Atoi(id[0])
	//noErr(err)
	user,err := models.GetUserByUsername2(username)

	if err != nil{
		c.Protocol.Status = "ERR"
		c.Protocol.Error = " کاربر پیدا نشد "
		return nil
	}

	var followedIds []int
	//todo make mysql uniqe select
	err=base.DB.Select(&followedIds, "select FollowedUserId from following_list_member where UserId = ? order by Id DESC limit ? offset ? ", user.Id, FOLLOW_LIST_SHOW_LIMIT, offset)
	if err != nil {
		helper.DebugErr(err)
		return nil
	}
	usersFollow := models.UsersToInlineFollowView(followedIds, cuid)

	c.SendJson(usersFollow)
	return nil
}

type SyncFollowings struct {
	Add []models.UserBasicAndMe
	Remove []int
}

func SyncFollowingsAction(a *base.Action) base.AppErr  {
	last_str := a.Req.Form.Get("last")//last TimeStamp
	last := helper.StrToInt(last_str,0)

	users:=models.GetAllFollowingsUser(a.UserId(),last)

	list := models.PreloadFollowingsListTypesForUser(a.UserId())
	followings_res := []models.UserBasicAndMe{}
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

/*
func UsersToInlineFollowView(userIds []int, cuid int) []models.UserInlineFollowView {
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
/*

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
*/
