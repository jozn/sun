package pipesold

import "ms/sun/helper"

func OnNewUserWsConnected(UserId int) {
	cmds := GetEarlistCmdsFromRedis(UserId)
	helper.Debug("OnNewUserWsConnected: len=", len(cmds))

	AllPipesMap.SendCmdsToUser(UserId, cmds)
}
