package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func SessionGetUserInfo(c *base.Action) base.AppErr {
	MustBeUser(c)
	mem := models.UserMemoryStore.GetForUser(c.UserId())
	c.SendJson(mem.UserTable)
	return nil
}
