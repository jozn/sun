package models

import "ms/sun/base"

//todo upadte
func OnNewUserWsConnected(UserId int) {
	call := base.NewCall("TimeMs")
	AllPipesMap.SendToUser(UserId, call)

	/*cmds := GetEarlistCmdsFromRedis(UserId)
	helper.Debug("OnNewUserWsConnected: len=", len(cmds))

	AllPipesMap.SendCmdsToUser(UserId, cmds)*/
}
