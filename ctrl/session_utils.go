package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

//////////////// Sessions utils //////////////////////

func MustBeUserAndUpdate(c *base.Action) *models.ReqParams {
	req := MustBeUser(c)

	return req
}

func MustBeUser(c *base.Action) *models.ReqParams {
	req := models.Session_ProcessHttpReq(c.Req)
	if req.UserId < 1 {
		print("NOT LOGED")
		//panic(constants.HttpIsNotUser)
	}
	c.SetUserId(req.UserId)
	return req
}

func UpdateSessionActivityIfUser(c *base.Action) *models.ReqParams {
	req := models.Session_ProcessHttpReq(c.Req)

	c.SetUserId(req.UserId)
	return req
}
