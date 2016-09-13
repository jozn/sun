package pipes

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/keygen"
	"ms/sun/store"
	//"ms/sun/cmd"
	"encoding/json"
)

func SaveCmdToRedis(UserId int, cmd *base.Command) {
	p := store.GetRedisPool()
	gen := keygen.NewForUser(UserId)
	key := gen.RedisMsgsAllKey() // "user_msgs:156"
	//con,_ := p.Get()
	r := p.Cmd(store.REDIS_SORTED_LIST_ADD, key, cmd.CmdId, helper.ToJson(cmd))
	helper.Debug("SaveCmdToRedis()  ", r.Err)
}

func RemoveCmdsFromRedis(UserId int, minNano, maxNano int64) {
	p := store.GetRedisPool()
	gen := keygen.NewForUser(UserId)
	key := gen.RedisMsgsAllKey() // "user_msgs:156"
	//con,_ := p.Get()
	r := p.Cmd(
		store.REDIS_SORTED_LIST_REMOVE_RANGE_SCORE,
		key,
		minNano-100,
		maxNano+100)
	helper.Debug("RemoveCmdsFromRedis()  ", r.Err, key, minNano, maxNano)
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
	//fmt.Println("ZRANGE()  ",r.Err, key, minNano, maxNano)
	if r.Err != nil {
		res, err := r.List()
		if err != nil {
			for _, c := range res {
				cmd := base.Command{}
				json.Unmarshal([]byte(c), &cmd)
				cmds = append(cmds, &cmd)
			}
		}
	}
	return nil
}
