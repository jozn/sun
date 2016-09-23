package ctrl

import (
    "ms/sun/models"
    "ms/sun/base"
)

func SyncUsersCtrl(c *base.Action) base.AppErr  {
    MustBeUserAndUpdate(c)

    lastTime := c.GetParamInt("last",-1);

    c.SendJson(models.SyncGetAllChangedUser(c.UserId(),lastTime))

    return nil
}
