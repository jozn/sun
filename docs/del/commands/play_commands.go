package commands

import (
	"encoding/json"
	"ms/sun/base"
	"ms/sun/models"
	"time"
)

func HelloCommand(c *base.WSAction) {
	//	_ = User{}
	time.Sleep(time.Microsecond * 10)
	//	c.Req.Params
	c.Res.Payload = models.User{}
	c.SendRes()
}

func TimeCommand(c *base.WSAction) {
	time.Sleep(time.Microsecond * 10)
	res := base.WSRes{}
	res.ResTime = time.Now().Unix()
	res.Status = "sayHi"
	c.Send(res)
}

func TimeCommand2(c *base.WSAction) {
	time.Sleep(time.Microsecond * 10)
	res := base.WSRes{}
	bu := models.UserInlineFollowView{}
	bu.UserId = 150
	bu.IFollowType = int(time.Now().Unix())
	bu.UserName = "ssssssssssss7777sssssss7"
	bu.AvatarSrc = "avatarataratar"
	s, _ := json.Marshal(bu)
	com := base.Command{Name: "time", Data: string(s)}
	res.Commands = append(res.Commands, &com)
	//res. = append(res.Commands, com)

	res.ResTime = time.Now().Unix()
	res.Status = "sayHi"
	c.Send(res)
}
