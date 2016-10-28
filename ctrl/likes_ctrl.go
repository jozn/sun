package ctrl

import (
	// "encoding/json"
	// "strings"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
    "fmt"
)

const LIKES_LIST_LIMIT = 50

//todo: suport pagination + AmIFollowing
func GetLikesAction(c *base.Action) base.AppErr {
	pidStr := c.Req.Form.Get("post_id")
	pageStr := c.Req.Form.Get("page")
	post_id := helper.StrToInt(pidStr, 0)
	page := helper.StrToInt(pageStr, 0)

	if post_id < 1 {
		return nil
	}
	offset := page * LIKES_LIST_LIMIT

	//var likes []models.Like
	var UserIds []int
	err := base.DB.Select(&UserIds, "select UserId from likes where PostId = ? order by Id DESC limit ? offset ? ", post_id, LIKES_LIST_LIMIT, offset)
	if err != nil {
		helper.DebugErr(err)
		return nil
	}

	usersFollow := models.UsersToInlineFollowView(UserIds, c.UserId())
	c.SendJson(usersFollow)
	return nil
}

func PostAddLikeAction(c *base.Action) base.AppErr {
    defer func() {
        e := recover()
        if e != nil {

            fmt.Println("PANIC 5: ",e)
            c.SendJson(c)
        }

    }()
	pids := c.Req.Form.Get("post_id")
	pid := helper.StrToInt(pids, 0)
	if pid < 1 {
		return nil
	}
	models.UserMemoryStore.AddPostLike(c.UserId(), pid)
	c.SendText("OK")
	return nil
}

func PostRemoveLikeAction(c *base.Action) base.AppErr {
	pids := c.Req.Form.Get("post_id")
	pid := helper.StrToInt(pids, 0)
	if pid < 1 {
		return nil
	}
	models.UserMemoryStore.RemovePostLike(c.UserId(), pid)
	c.SendText("OK")
	return nil
}
