package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func DebugMemUser_ctrl(c *base.Action) base.AppErr {

	_, _ = models.MemoryStore_User.GetUser(10)
	c.SendJson(models.MemoryStore_User)
	return nil
}
