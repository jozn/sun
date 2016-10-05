package pipes

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

func GetUserForTable(c *base.CmdAction) {
	uid := helper.StrToInt(c.Cmd.Data, 0)
	if uid == 0 {
		return
	}
	print("data: ", uid)

	u := models.GetUserById2(uid)
	res := models.UserBasicAndMe{}
	res.FormUserAndMe(u, c.UserId)
	cmd := base.Command{
		Name: "SetUserForTable",
		Data: helper.ToJson(res),
	}

	AllPipesMap.SendCmdToUser(c.UserId, &cmd)
}
