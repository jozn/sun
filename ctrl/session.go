package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func SessionGetUserInfo(c *base.Action) base.AppErr {
	MustBeUser(c)
	mem, ok := models.MemoryStore_User.GetUser(c.UserId())
	if ok {
		c.SendJson(mem)
	}
	c.SendJson(nil)
	return nil
}
