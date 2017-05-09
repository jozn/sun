package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"ms/sun/models/x"
)

func GetPostsForProfileAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	const LIMIT = 30

	laststr := c.Req.Form.Get("last") //last that have
	pagestr := c.Req.Form.Get("page")
	last := helper.StrToInt(laststr, 0)
	page := helper.StrToInt(pagestr, 0)
	limit := c.GetParamInt("limit", LIMIT)

	uid := c.UserId()
	profileId := c.GetParamInt("profile_id", 0)
	userName := c.Req.Form.Get("user_name")

	var u x.User
	var ok bool

	if profileId > 0 {
		u, ok = models.MemoryStore_User.GetUser(profileId)

	} else if len(userName) > 0 {
		u, ok = models.MemoryStore_User.GetUserByUserName(userName)
	}

	if !ok {
		c.Protocol.Error = "NOT FOUND"
		c.Protocol.Status = "ERROR"
		return nil
	}

	profileId = u.Id

	selctor := x.NewPost_Selector().UserId_Eq(profileId).OrderBy_Id_Desc().Limit(limit)

	if last > 0 {
		selctor.Id_LT(last)
	} else if page > 0 {
		selctor.Offset((page - 1) * limit)
	}

	posts, err := selctor.GetRows(base.DB)
	if err != nil {
		helper.DebugPrintln(err)
		c.SendJson(nil)
		return err
	}

	view := models.Views.PostsViews(posts, uid)
	c.SendJson(view)

	return nil
}

func GetProfileInfoAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	profileId := c.GetParamInt("profile_id", 0)
	userName := c.Req.Form.Get("user_name")

	var u x.User
	var ok bool

	if profileId > 0 {
		u, ok = x.Store.GetUserById(profileId)

	} else if len(userName) > 0 {
		u, ok = models.MemoryStore_User.GetUserByUserName(userName)
	}

	if !ok {
		c.Protocol.Error = "NOT FOUND"
		c.Protocol.Status = "ERROR"
		return nil
	}

	res := models.ProfileInfo{
		u.UserCounts,
		*models.Views.UserViewSync(c.UserId(), u.Id),
	}

	//res := models.Views.UserViewSync(c.UserId(),u.Id)

	c.SendJson(res)
	return nil

	return nil
}
