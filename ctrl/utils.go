package ctrl

import (
	"ms/sun/base"
	"ms/sun/constants"
	"ms/sun/models"
)

//////////////// Sessions utils //////////////////////
func MustBeUserAndUpdate(c *base.Action) {
	uid := c.GetParamInt("user_id", -1)
	session := c.Req.Form.Get("session")

	isLoged := models.UserMemoryStore.IsUserSessionAndUpdateActivity(uid, session)
	if isLoged == false {
		print("NOT LOGED")
		panic(constants.HttpIsNotUser)
	}
	c.SetUserId(uid)
}

func MustBeUser(c *base.Action) {
	uid := c.GetParamInt("user_id", -1)
	session := c.Req.Form.Get("session")

	isLoged := models.UserMemoryStore.IsUserSession(uid, session)
	if isLoged == false {
		panic(constants.HttpIsNotUser)
	}
	c.SetUserId(uid)
}

func UpdateSessionActivityIfUser(c *base.Action) {
	uid := c.GetParamInt("user_id", -1)
	session := c.Req.Form.Get("session")

	isLoged := models.UserMemoryStore.IsUserSessionAndUpdateActivity(uid, session)
	if isLoged == false {

	}
	c.SetUserId(uid)
}

func GetSessionUserId(c *base.Action) int {
	if c.UserId() > 0 {
		return c.UserId()
	}

	uid := c.GetParamInt("user_id", -1)
	session := c.Req.Form.Get("session")

	isLoged := models.UserMemoryStore.IsUserSession(uid, session)
	if isLoged == false {
		return 0
	}
	c.SetUserId(uid)
	return uid
}

///////////////////////////////////////////////////
