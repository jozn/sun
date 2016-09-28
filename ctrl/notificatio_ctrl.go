package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

func VNoftificationsCtrl(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	helper.DebugPrintln("CCCCCCCCC")

	c.SendJson(models.Notification_GetLastsViews(c.UserId()))
	return nil
}

func VNoftificationsCtrl2(c *base.Action) {
	//MustBeUserAndUpdate(c)

	//_ =(models.Notification_GetLastsViews(c.UserId()))
	c.SendJson("asdsa")
}
