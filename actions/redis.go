package actions

import (
	//"ms/sun/base"
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"ms/sun/store"
)

func RedisSavePlay(c *base.Action) base.AppErr {
	p := store.GetRedisPool()
	U := x.UserTable{}
	key := "UserView:" + helper.RandString(1)
	con, _ := p.Get()
	con.Cmd("set", key, helper.ToJson(U))

	//fmt.Println("redis redult ",r.String())

	res := con.Cmd("GET", key)
	b, err := res.Str()
	b2, err := res.Bytes()
	fmt.Println("redis result ", key, " : ", res.Err, err, b)
	u2 := x.UserTable{}
	_ = b2
	_ = u2
	us := make([]x.UserTable, 10)
	//json.Unmarshal(b2,&u2)
	c.SendJson(us)

	return nil
}

func PlaySomething(c *base.Action) base.AppErr {
	__sorted_redis()
	return nil
}

func __sorted_redis() {
	//p := base.GetRedisPool()
	//U := models.UserTable{}
	//U.Id = rand.Intn(1000)
	//key := "UserView:6"
	//con,_ := p.get()
	//r := con.Cmd("zadd",key , U.Id ,helper.ToJson(U))
	//fmt.Println("__sorted_redis  ",r.Err,U)

}
