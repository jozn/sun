package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
    "ms/sun/models/x"
    "ms/sun/helper"
)

func PingAction(c *base.Action) base.AppErr {
	c.SendText("PONG")
	return nil
}

func PushViewAction(c *base.Action) base.AppErr {
    rows,err := x.NewDirectLog_Selector().ToUserId_Eq(6).GetRows(base.DB)
    helper.NoErr(err)
    res := models.PushView_directLogsTo_PB_ChangesHolderView(3,rows)
    c.SendJson(res)
    return nil
}

