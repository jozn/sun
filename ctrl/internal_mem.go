package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func DebugMemUser_ctrl(c *base.Action) base.AppErr {
    c.SendJson(models.MemoryStore_User)
	return nil
}

func DebugMemUser_ctrl2(c *base.Action) base.AppErr {
    u, _ := models.MemoryStore_User.GetUser(10)
    c.SendJson(u)
    return nil
}
