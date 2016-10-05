package pipesold

import (
	"github.com/gorilla/websocket"
	"sync"
	//"time"
	"encoding/json"
	"ms/sun/base"
	"ms/sun/helper"
	"os"
    "ms/sun/cmd"
)

type UserDevicePipe struct {
	UserId         int
	IsOpen         bool
	Ws             *websocket.Conn
	ToDeviceChan   chan base.Call
	FromDeviceChan chan base.Call // NOT NEEDED?
	m              sync.RWMutex
}

func (pipe *UserDevicePipe) ServeIncomingCalls() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				helper.Debug("Recovered in ws messaging clinet request", r)
				pipe.ShutDownCompletely()
			}
		}()

		fLog, e := os.OpenFile("./ws_Log_"+helper.IntToStr(helper.TimeNow()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if e != nil {
			panic(e)
		}

		for {
			var call base.Call
			messageType, bytes, err := pipe.Ws.ReadMessage() //blocking

			helper.Debug("messageType: ", " ::", messageType, string(bytes))
			if messageType == websocket.CloseMessage || err != nil {
				pipe.ShutDownCompletely()
				helper.Debugf("closeing pip for userId: %v , messageType:%v , err: %v", pipe.UserId, messageType, err)
				return
			}

			if messageType == websocket.TextMessage {
				err = json.Unmarshal(bytes, &call)
				if err == nil {
					serverWSReqCommands(call, pipe)
				}
				fLog.WriteString(string(bytes) + "\n")
				fLog.Sync()
			}

			if messageType == websocket.BinaryMessage {
				pipe.ShutDownCompletely()
				return
				/*f, _ := os.Create("./upload/ws_" + helper.RandString(8))
				wd, _ := os.Getwd()
				fmt.Println("ws: "+" wd:"+wd+"binary:  ", bytes)
				f.Write(bytes)
				f.Close()*/
			}
		}
	}()
}

func (pipe *UserDevicePipe) ServeSendToUserDevice() {
	go func() {
		//if we panic somehow? in writing json shutdown
		defer func() {
			if r := recover(); r != nil {
				pipe.ShutDownCompletely()
			}
		}()
		for r := range pipe.ToDeviceChan {
			helper.Debug("sending to user fom ToDeviceChan Command size ", r)
			pipe.Ws.WriteJSON(r)
		}
		//after closing chanel
		//fixme :not neccossroy waht about .ShutDownCompletely()?
		err := pipe.Ws.Close()
		pipe.IsOpen = false
		helper.Debug("closed: ", err)
	}()
}

func (pipe *UserDevicePipe) SendToUser(resCall base.Call) {
	if pipe.IsOpen {
		pipe.ToDeviceChan <- resCall
	}
}

func (pipe *UserDevicePipe) ShutDown() {
	pipe.Ws.Close()
	close(pipe.ToDeviceChan)
}

func (pipe *UserDevicePipe) ShutDownCompletely() {
	AllPipesMap.ShutDownUser(pipe.UserId)
}

/////////////// Commands handler //////////////////

func serverWSReqCommands(reqCall base.Call, pipe *UserDevicePipe) {
    //TODO UPGRADE THIS TO NEW Call
	/*arr := make([]int64, 0, len(req.Commands))
	for _, cmd := range req.Commands {
		arr = append(arr, cmd.ClientNanoId)
	}

	if len(arr) > 0 {
		cmdsRecived := base.NewCommand("CommandsReceivedToServer")
		cmdsRecived.ClientNanoId = -1
		cmdsRecived.ServerNanoId = -1
		cmdsRecived.SetData(arr)
		AllPipesMap.SendCmdToUser(pipe.UserId, cmdsRecived)
	}*/

    reqCall.UserId = pipe.UserId
    fnCall := base.CallMapRouter[reqCall.Name]
    helper.Debug("serving Cmd: ", reqCall.Name, " Userid: ", pipe.UserId)
    if fnCall != nil {
        fnCall(pipe)
    } else { //send call func not found -- in debug

    }

    //del
	/*for _, cmd := range req.Commands {
		fnCall := base.CallMapRouter[cmd.Name]
		helper.Debug("serving Cmd: ", cmd.Name, " Userid: ", pipe.UserId)
		action := base.CmdAction{Req: &req, UserId: pipe.UserId, Cmd: &cmd}
		if fnCall != nil {
			fnCall(action)
		} else { //send cmd not found -- in debug
			//wsErrorCommand(&action)
		}
	}*/
}
