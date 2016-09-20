package actions

import (
    "ms/sun/base"
    "ms/sun/helper"
    "time"
    "math/rand"
"ms/sun/sync"
    "ms/sun/constants"
    "ms/sun/models"
    "ms/sun/pipes"
)

func SendSampleMesgTable(a *base.Action) base.AppErr {
    ds:=a.Req.Form.Get("delay")
    us :=a.Req.Form.Get("user")
    tus :=a.Req.Form.Get("from")
    l :=a.Req.Form.Get("limit")
    text :=a.Req.Form.Get("text")
    img :=a.Req.Form.Get("img")
    e :=a.Req.Form.Get("emoji")

    dInt := helper.StrToInt(ds,2000)
    user := helper.StrToInt(us,6)
    limit := helper.StrToInt(l,10)

    emoji := true
    if(e!=""){
        emoji = false
    }

    is_image := false
    if len(img)>0{
        is_image = true
    }

    go func() {
        for i:=0 ; i< limit ; i++ {
            txt:=helper.FactRandStr(15)
            rnd:=rand.Intn(10)
            if  rnd== 5 {//big text 10%
                txt = helper.FactRandStrEmoji(150,emoji)
            }else if rnd== 6 {
                txt = helper.FactRandStrEmoji(500,emoji)
            }else if rnd== 7 {
                txt = helper.FactRandStrEmoji(2500,emoji)
            }
            if text != "" {
                txt = text
            }

            uid := helper.StrToInt(tus,rand.Intn(80)+1)
            //msg := chat.MessagesTable{}
            msg := models.MessagesTableFromClient{}
            msg.RoomKey ="u"+ helper.IntToStr(uid)
            msg.UserId = uid
            msg.MessageKey = helper.RandString(10)
            msg.CreatedMs = helper.TimeNowMs()
            msg.MessageTypeId = 10
            msg.Text = txt

            if is_image {
                msg.MessageTypeId = 40
                msg.MediaName = helper.RandString(10)+".jpg"
                msg.MediaServerSrc = "http://localhost:5000/public/photo/"+ helper.IntToStr(rand.Intn(21)+1) + "_960.jpg"
                msg.MediaExtension =".jpg"
                msg.MediaSize = 200000
                msg.MediaHeight = 600
                msg.MediaWidth = 960
                msg.MediaThumb64 = `/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEB
AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEB
AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCADVAKADASIA
AhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQA
AAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3
ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWm
p6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEA
AwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSEx
BhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElK
U1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3
uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD+rCug
rn66Cg+XNitisetigDoK2Kx6KANivFPG2qdfwHHt5nvxnp9QvpXqOo6p/wAS0/hn839j0/kR3GT8
h/EjxQTn2I9eoP1/+v26UAcRqOqH+0j14xjP1cDt3znr69xmtvw13/3j/wC1K8T/ALUP9o5564xz
n73Xp0z+nvXqXhjv+H/tSgD6I8Md/wAP/alN8Sf8g1/w/wDQpKNM6N/n+Gt6gD4l8beFj/aJzjtj
GfWTOfw6e/rW18NvC/8AZWefT9SemDz09+h5OQD9Raj4W/tTI9/fsW7Z7g9P043E0zwt/ZeR2455
9X4xn/ZHHPUkjvQBuaZ/yDW/z/FTq2Kx6ACiiigAooooA5+t3Tep+o/9CkrCroKANitiuPrYoA2K
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
M0UV+Tn2PD20fSH/AKVUP//Z`

            }

            cmd:=base.NewCommand(constants.MsgsAddNew)
            cmd.AddSliceData(msg)
            cmd.MakeDataReady()
            //res :=base.WSRes{
            //}
            //res.Commands = []*base.Command{&cmd}
            //sync.AllPipesMap.SendToUser(6,res)
            pipes.AllPipesMap.SendAndStoreCmdToUser(user,cmd)
            time.Sleep(time.Millisecond * time.Duration(dInt))
        }
        //return
    }()
    return nil
}

func SendSampleMesgTable2(a *base.Action) base.AppErr {
    ds:=a.Req.Form.Get("delay")
    dInt := helper.StrToInt(ds,2000)
    go func() {
        for  {
            uid := 6//rand.Intn(80)+1
            msg := models.MessagesTable{}
            msg.RoomKey ="u"+ helper.IntToStr(uid)
            msg.UserId = uid
            msg.MessageKey = helper.RandString(10)
            msg.CreatedTimestampMs = helper.TimeNowMs()
            msg.MessageTypeId = 10
            msg.Text = helper.FactRandStr(15)

            cmd:=base.Command{
                Name: constants.MsgsAddNew,//"addMsg",
            }
            cmd.AddSliceData(msg)
            cmd.MakeDataReady()
            //res :=base.WSRes{
            //}
            //res.Commands = []*base.Command{&cmd}
            //sync.AllPipesMap.SendToUser(6,res)
            sync.AllPipesMap.SendCmdToUser(2,&cmd)
            time.Sleep(time.Millisecond * time.Duration(dInt))
        }
        //return
    }()
    return nil
}

