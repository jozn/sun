package ctrl

import (
	"ms/sun/base"
    "ms/sun/models/x"
)

func SessionGetUserInfo(c *base.Action) base.AppErr {
	MustBeUser(c)
	//mem, ok := models.MemoryStore_User.GetUser(c.UserId())
	mem, ok := x.Store.GetUserById(c.UserId())
	if ok {
        mem.FullName = mem.FirstName + " " +mem.LastName
		c.SendJson(mem)
        return nil
	}
	c.SendJson(nil)
	return nil
}
