package ctrl

import (
	// "encoding/json"
	// "strings"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

const LIKES_LIST_LIMIT = 20

//todo: suport pagination + AmIFollowing
func GetLikesAction(c *base.Action) base.AppErr {
	pidStr := c.Req.Form.Get("post_id")
	pageStr := c.Req.Form.Get("page")
	last_id := c.GetParamInt("last", 0)
	limit := c.GetParamInt("limit", LIKES_LIST_LIMIT)
	post_id := helper.StrToInt(pidStr, 0)
	page := helper.StrToInt(pageStr, 0)

	_ = last_id
	if post_id < 1 {
		return nil
	}
	//page indexed : >= 1
	offset := (page - 1) * LIKES_LIST_LIMIT

	selector := models.NewLike_Selector().Select_UserId().PostId_EQ(post_id)

	if last_id > 0 {
		selector.Id_GT(last_id)
	} else {
		selector.Offset(offset)
	}

	UserIds, err := selector.OrderBy_Id_Desc().Limit(limit).GetIntSlice(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return nil
	}

	//usersFollow := models.UsersToInlineFollowView(UserIds, c.UserId())
	usersFollow := models.Views.UserBasicAndMeForUsers(c.UserId(), UserIds)
	c.SendJson(usersFollow)
	return nil
}

func PostAddLikeAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	pids := c.Req.Form.Get("post_id")
	pid := helper.StrToInt(pids, 0)
	if pid < 1 {
		return nil
	}
	//models.UserMemoryStore.AddPostLike(c.UserId(), pid)
	models.CreatePostLike(c.UserId(), pid)
	c.SendText("OK")
	return nil
}

func PostRemoveLikeAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	pids := c.Req.Form.Get("post_id")
	pid := helper.StrToInt(pids, 0)
	if pid < 1 {
		return nil
	}
	//models.UserMemoryStore.RemovePostLike(c.UserId(), pid)
	models.DeletePostLike(c.UserId(), pid)
	c.SendText("OK")
	return nil
}

/*var UserIds []int
err := base.DB.Select(&UserIds, "select UserId from likes where PostId = ? order by Id DESC limit ? offset ? ", post_id, LIKES_LIST_LIMIT, offset)
*/
