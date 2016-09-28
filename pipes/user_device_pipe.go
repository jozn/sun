package pipes

import (
	"github.com/gorilla/websocket"
	"sync"
	//"time"
	"encoding/json"
	"ms/sun/base"
	"ms/sun/helper"
	"os"
)

type UserDevicePipe struct {
	UserId         int
	IsOpen         bool
	Ws             *websocket.Conn
	ToDeviceChan   chan *base.WSRes
	FromDeviceChan chan *base.WSRes // NOT NEEDED?
	m              sync.RWMutex
}

func (pipe *UserDevicePipe) ServeIncomingReqs() {
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
			var req base.WSReq
			messageType, bytes, err := pipe.Ws.ReadMessage() //blocking

			helper.Debug("messageType: ", " ::", messageType, string(bytes))
			if messageType == websocket.CloseMessage || err != nil {
				pipe.ShutDownCompletely()
				helper.Debugf("closeing pip for userId: %v , messageType:%v , err: %v", pipe.UserId, messageType, err)
				return
			}

			if messageType == websocket.TextMessage {
				err = json.Unmarshal(bytes, &req)
				if err == nil {
					serverWSReqCommands(req, pipe)
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
			helper.Debug("sending to user fom ToDeviceChan Command size ", len(r.Commands))
			pipe.Ws.WriteJSON(r)
		}
		//after closing chanel
		//fixme :not neccossroy waht about .ShutDownCompletely()?
		err := pipe.Ws.Close()
		pipe.IsOpen = false
		helper.Debug("closed: ", err)
	}()
}

func (pipe *UserDevicePipe) SendToUser(res base.WSRes) {
	if pipe.IsOpen {
		pipe.ToDeviceChan <- &res
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

func serverWSReqCommands(req base.WSReq, pipe *UserDevicePipe) {
	arr := make([]int64, 0, len(req.Commands))
	for _, cmd := range req.Commands {
		arr = append(arr, cmd.ClientNanoId)
	}

	if len(arr) > 0 {
		cmdsRecived := base.NewCommand("CommandsReceivedToServer")
		cmdsRecived.ClientNanoId = -1
		cmdsRecived.ServerNanoId = -1
		cmdsRecived.SetData(arr)
		AllPipesMap.SendCmdToUser(pipe.UserId, cmdsRecived)
	}

	for _, cmd := range req.Commands {
		fncmd := base.CmdMapRouter[cmd.Name]
		helper.Debug("serving Cmd: ", cmd.Name, " Userid: ", pipe.UserId)
		action := base.CmdAction{Req: &req, UserId: pipe.UserId, Cmd: &cmd}
		if fncmd != nil {
			fncmd(&action)
		} else { //send cmd not found -- in debug
			//wsErrorCommand(&action)
		}
	}
}
