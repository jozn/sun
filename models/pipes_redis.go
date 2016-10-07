package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/keygen"
	"ms/sun/store"
	//"ms/sun/cmd"
	"encoding/json"
)

func StoreCommandsToRedis(UserId int, cmd *base.Command) {
	p := store.GetRedisPool()
	gen := keygen.NewForUser(UserId)
	key := gen.RedisMsgsAllKey() // "user_msgs:156"
	//con,_ := p.Get()
	score := helper.TimeNowNano()
	if cmd.ServerNanoId < 1 {
		cmd.ServerNanoId = score
	}
	score = cmd.ServerNanoId
	r := p.Cmd(store.REDIS_SORTED_LIST_ADD, key, score, helper.ToJson(cmd))
	helper.Debug("SaveCmdToRedis() Err: ", r.Err)
}

func RemoveCommandsFromRedis(UserId int, minNano, maxNano int64) {
	p := store.GetRedisPool()
	gen := keygen.NewForUser(UserId)
	key := gen.RedisMsgsAllKey() // "user_msgs:156"
	//con,_ := p.Get()
	r := p.Cmd(
		store.REDIS_SORTED_LIST_REMOVE_RANGE_SCORE,
		key,
		minNano-10,
		maxNano+10)
	helper.Debug("RemoveCmdsFromRedis() Err: ", r.Err, key, minNano, maxNano)
}

func GetEarlistCmdsFromRedis(UserId int) (cmds []*base.Command) {
	p := store.GetRedisPool()
	gen := keygen.NewForUser(UserId)
	key := gen.RedisMsgsAllKey() // "user_msgs:156"
	//con,_ := p.Get()
	r := p.Cmd(
		store.REDIS_SORTED_LIST_GET_RANGE_INDEX,
		key,
		0,
		-1)
	//helper.Debugf("ZRANGE()  ",r.Err, key)
	if r.Err == nil {
		res, err := r.List()
		//helper.Debugf("ZRANGE()  ",len(res), err, key)
		if err == nil {
			for _, c := range res {
				cmd := base.Command{}
				//helper.Debugf("ZRANGE()  ",len(res), err, key)
				err = json.Unmarshal([]byte(c), &cmd)
				//helper.Debugf("ZRANGE() Unamrshal err  ", err)
				cmds = append(cmds, &cmd)
			}
		}
	}
	return
}
