package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func GetUserForTable(c *base.CmdAction) {
	uid := helper.StrToInt(c.Cmd.Data, 0)
	if uid == 0 {
		return
	}
	print("data: ", uid)

	u := GetUserById2(uid)
	res := UserBasicAndMe{}
	res.FormUserAndMe(u, c.UserId)
	_ = base.Command{
		Name: "SetUserForTable",
		Data: helper.ToJson(res),
	}

	//AllPipesMap.SendCmdToUser(c.UserId, &cmd)
}
