package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

func PingAction(c *base.Action) base.AppErr {
	c.SendText("PONG")
	return nil
}

func PushViewAction(c *base.Action) base.AppErr {
	rows, err := x.NewDirectUpdate_Selector().ToUserId_Eq(6).GetRows(base.DB)
	helper.NoErr(err)
	_ = rows
	res := 1//models.ViewPush_DirectUpdatesList_To_GetDirectUpdatesView(3, rows)
	c.SendJson(res)
	return nil
}
