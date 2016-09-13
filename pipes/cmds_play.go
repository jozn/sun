package pipes

import (
    "ms/sun/base"
    "fmt"
    "encoding/json"
    "math/rand"
    "ms/sun/helper"
//"time"
    "ms/sun/models"
    "ms/sun/sync"
)
func AddNewMsg(c *base.CmdAction){
    fmt.Println("called AddNewMsg :",*c)

    m := models.MessagesTable{}
    json.Unmarshal([]byte(c.Cmd.Data),&m)

    uid := helper.StrToInt(m.RoomKey[1:],rand.Intn(80))
    msg := models.MessagesTable{}
    msg.RoomKey ="u"+ helper.IntToStr(uid)
    msg.UserId = uid
    msg.MessageKey = helper.RandString(10)
    msg.CreatedTimestampMs = helper.TimeNowMs()+100000
    msg.MessageTypeId = 10
    msg.Text = helper.FactRandStr(15)+ " ========== "+ m.RoomKey + " =========\n "+ m.Text

    cmd := base.NewCommand("addMsg")
    cmd.SetData(msg)
    //res :=base.WSRes{
    //}
    //res.Commands = []*base.Command{&cmd}
    sync.AllPipesMap.SendAndStoreCmdToUser(6,cmd)
    //sync.AllPipesMap.SendToUser(6,res)
    //time.Sleep(time.Millisecond * time.Duration(dInt))


}

func HelloCmd(c *base.CmdAction) {
    fmt.Println("called HelloCmd :",*c)
    //time.Sleep(time.Microsecond * 10)
    //res := base.WSRes{}
    //bu := models.UserInlineFollowView{}
    //bu.UserId = 150
    //bu.IFollowType = int(time.Now().Unix())
    //bu.UserName = "ssssssssssss7777sssssss7"
    //bu.AvatarUrl = "avatarataratar"
    //s, _ := json.Marshal(bu)
    //com := base.Command{Name: "time", Data: string(s)}
    //res.Commands = append(res.Commands, com)
    ////res. = append(res.Commands, com)
    //
    //res.ResTime = time.Now().Unix()
    //res.Status = "sayHi"
    //c.Send(res)
}

func EchoRes(c *base.CmdAction) {
    cmd:=base.NewResponseCommand(c.Cmd.ResId)

    cmd.SetData(models.UserBasic{FirstName:helper.RandString(20)})
    sync.AllPipesMap.SendCmdToUser(6,cmd)
}

func EchoCmd(c *base.CmdAction) {
    b,_:=json.Marshal(c.Cmd)
    r :=base.WSRes{Status:"BB",ReqKey: string(b)}

    sync.AllPipesMap.SendToUser(c.UserId,r)
}

