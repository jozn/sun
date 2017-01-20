package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func NotifyCtrl(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

    last := c.GetParamInt("last",0)

	c.SendJson(models.Notification_GetLastsViews(c.UserId() ,last))
	return nil
}

func NotifyAddRemoveCtrl(c *base.Action) base.AppErr {
    MustBeUserAndUpdate(c)

    last := c.GetParamInt("last",0)
    res := NotifyAddRemoveView{
        Add: models.Notification_GetLastsViews(c.UserId() ,last),
        Remove: models.Notification_ListOfRemoved(c.UserId()),
    }

    c.SendJson(res)
    return nil
}

type NotifyAddRemoveView struct {
    Add []models.NotificationView
    Remove []int
}