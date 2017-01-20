package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

func NotifyCtrl(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

    last := c.GetParamInt("last",0)

	c.SendJson(models.Notification_GetLastsViews(c.UserId() ,last))
	return nil
}
