package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"

	"ms/sun/helper"
	//"ms/sun/docs/del"
)

const LIST_TYPE_FOLLOWINGS = 1
const LIST_TYPE_FOLLOWERS = 2
const LIST_TYPE_LIKES = 3

func PingAction(c *base.Action) base.AppErr {
	c.SendText("PONG")
	return nil
}

func FollowAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	cuid := c.UserId()

	//fid := c.Req.Form.Get("followed_user_id")
	fid := c.Req.Form.Get("followed_user_id")
	fUserId := helper.StrToInt(fid, -1)
	if fUserId < 1 {
		return nil
	}
	models.Follow(cuid, fUserId)
	//c.SendJson(models.MemoryStore.UserFollowingList_GetFollowingTypeForUsers(cuid, fUserId))
	c.SendJson(models.MemoryStore.UserFollowingList_GetFollowingTypeForUsers(cuid, fUserId))
	return nil
}

func UnfollowAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	cuid := c.UserId()

	fid := c.Req.Form.Get("followed_user_id")
	fUserId := helper.StrToInt(fid, -1)
	if fUserId < 1 {
		return nil
	}
	models.UnFollow(cuid, fUserId)
	c.SendText("ok")
	return nil
}

const FOLLOW_LIST_SHOW_LIMIT = 30

//followers of a user
func GetFollowersListAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	cuid := c.UserId()

	username := c.Req.Form.Get("username")
	pageStr := c.Req.Form.Get("page")
	page := helper.StrToInt(pageStr, 0)
	limit := c.GetParamInt("limit", FOLLOW_LIST_SHOW_LIMIT)
	offset := limit * (page - 1)
	last := c.GetParamInt("last", 1000000000)
	peer_id := c.GetParamInt("peer_id", 0)

	_ = last
	var user models.UserTable
	var ok bool
	if peer_id < 1 {
		//user, err = models.GetUserByUsername2(username)
		user, ok = models.MemoryStore_User.GetUserByUserName(username)
	} else {
		user,ok = models.MemoryStore_User.GetUser(peer_id)
        peer_id = user.Id
	}

	if !ok { //|| user == nil{
		c.Protocol.Status = "ERR"
		c.Protocol.Error = " کاربر پیدا نشد "
		return nil
	}

	selector := models.NewFollowingListMember_Selector().
		Select_UserId().
		FollowedUserId_EQ(peer_id).
		OrderBy_Id_Desc().
		Limit(limit)

	if offset > 0 {
		selector.Offset(offset)
	}

	userIds, err := selector.GetIntSlice(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return nil
	}

	//usersFollow := models.GetListOfUserForFollowType(userIds, cuid)
	usersFollow := models.Views.UserBasicAndMeForUsers(cuid, userIds)

	c.SendJson(usersFollow)
	return nil
}

func GetFollowingsListAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	cuid := c.UserId()

	username := c.Req.Form.Get("username")
	pageStr := c.Req.Form.Get("page")
	page := helper.StrToInt(pageStr, 0)
	limit := c.GetParamInt("limit", FOLLOW_LIST_SHOW_LIMIT)
	offset := limit * (page - 1)
	last := c.GetParamInt("last", 1000000000)
	peer_id := c.GetParamInt("peer_id", 0)

	_ = last
	var user models.UserTable
    var ok bool
    if peer_id < 1 {
        //user, err = models.GetUserByUsername2(username)
        user, ok = models.MemoryStore_User.GetUserByUserName(username)
    } else {
        user,ok = models.MemoryStore_User.GetUser(peer_id)
        peer_id = user.Id
    }

    if !ok {
		c.Protocol.Status = "ERR"
		c.Protocol.Error = " کاربر پیدا نشد "
		return nil
	}

	selector := models.NewFollowingListMember_Selector().
		Select_FollowedUserId().
		UserId_EQ(peer_id).
		OrderBy_Id_Desc().
		Limit(limit)

	if offset > 0 {
		selector.Offset(offset)
	}

	userIds, err := selector.GetIntSlice(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return nil
	}

	//usersFollow := models.GetListOfUserForFollowType(userIds, cuid)
	usersFollow := models.Views.UserBasicAndMeForUsers(cuid, userIds)

	c.SendJson(usersFollow)
	return nil
}

/*
type SyncFollowings struct {
	Add    []models.UserBasicAndMe
	Remove []int
}
*/

/*
func SyncFollowingsAction(a *base.Action) base.AppErr {
	MustBeUserAndUpdate(a)
	last_str := a.Req.Form.Get("last") //last TimeStamp
	last := helper.StrToInt(last_str, 0)

	users := models.GetAllFollowingsUser(a.UserId(), last)

	list := models.PreloadFollowingsListTypesForUser(a.UserId())
	followings_res := []models.UserBasicAndMe{}
	for _, u := range users {
		cu := models.UserBasicAndMe{}
		cu.FromUser(u, list)
		//cu.UserBasic = u.UserBasic
		//cu.UpdatedTimestamp = u.UpdatedTimestamp
		//cu.FollowingType = list.FollowingType(u.Id)
		followings_res = append(followings_res, cu)
	}

	sync_res := SyncFollowings{
		Add:    followings_res,
		Remove: models.GetAllUnFollowedUserIds(a.UserId(), last),
	}
	a.SendJson(sync_res)
	return nil
}
*/
