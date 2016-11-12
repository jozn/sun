package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func SessionGetUserInfo(c *base.Action) base.AppErr {
	MustBeUser(c)
	mem, _ := models.MemoryStore_User.GetUser(c.UserId())
	c.SendJson(mem)
	return nil
}
