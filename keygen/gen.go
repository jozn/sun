package keygen

import "ms/sun/helper"

type gen struct {
	uid int
}

func NewForUser(UserId int) gen {
	return gen{uid: UserId}
}

func (g *gen) RedisMsgsAllKey() string {
	return "user_msgs:" + helper.IntToStr(g.uid)
}
