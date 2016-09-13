package sync

import (
    "github.com/gorilla/websocket"
    "fmt"
    //"ms/sun/sync"
"ms/sun/base"
)

//todo change UserDevicePipe => *UserDevicePipe
type pipesMap map[int]*UserDevicePipe
var AllPipesMap = make(pipesMap)// make(map[int]UserDevicePipe)
//var AllUserDevicesPipesMap2 = make(pipesMap)// make(map[int]UserDevicePipe)

func (m pipesMap) SendToUser(UserId int, res base.WSRes ) {
    pipe ,ok := m[UserId]
    fmt.Printf("sending to user:%d %v %v ",UserId, ok, res.Commands)
    if ok && pipe.IsOpen{
        defer func() {
            if r := recover(); r != nil {
                pipe.IsOpen = false
                fmt.Println("Recovered in SendToUser: ", r)
            }
        }()
        pipe.SendToUser(res)
        //pipe.ToDeviceChan <- res
    }
}

func (m pipesMap) SendCmdToUser(UserId int, cmd *base.Command ) {
    res := base.WSRes{Status:"OK",ReqKey: ""}
    res.Commands = []*base.Command{cmd}
    m.SendToUser(UserId,res)
}

func (m pipesMap) SendAndStoreCmdToUser(UserId int, cmd *base.Command ) {
    //store
    SaveCmdToRedis(UserId,cmd)

    //send
    res := base.WSRes{Status:"OK",ReqKey: ""}
    res.Commands = []*base.Command{cmd}
    m.SendToUser(UserId,res)
}

//adds a new pip
func (m pipesMap) ServeUserWs(UserId int, ws *websocket.Conn ) {
    pipe := UserDevicePipe{
        UserId: UserId,
        ToDeviceChan: make(chan *base.WSRes,10),
        IsOpen: true,
        Ws: ws,
    }

    fmt.Println("serving user ws for user:",UserId)

    pipe.ServeIncomingReqs()
    pipe.ServeSendToUserDevice()

    m[UserId] = &pipe
}





