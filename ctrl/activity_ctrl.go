package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func ActivityListCtrl(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	limit := c.GetParamInt("limit", 25)
	last := c.GetParamInt("last", 0)
	page := c.GetParamInt("page", 1)

	res := models.Activity_GetLastsViews(c.UserId(), page, limit, last)

	c.SendJson(res)
	return nil
}
