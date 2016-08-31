package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
)


func NoftificationsCtrl(c *base.Action) base.AppErr  {
    //MustBeUserAndUpdate(c)

    _ =(models.Notification_GetLastsViews(c.UserId()))
    c.SendJson("asdsa")
    return nil
}
