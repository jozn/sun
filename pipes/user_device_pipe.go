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
			}
		}()

		fLog, e := os.OpenFile("./ws_Log_"+helper.IntToStr(helper.TimeNow()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if e != nil {
			panic(e)
		}
		//FIXME:this is sequence message proccessing do we need multi process
		for {

			var req base.WSReq
			messageType, bytes, err := pipe.Ws.ReadMessage() //blocking
			//cnt++
			helper.Debug("messageType: ", " ::", messageType, string(bytes))
			if messageType == websocket.CloseMessage || err != nil {
				pipe.ShutDownCompletely()
				helper.Debug("closeding ", messageType, " error: ", err)
				return
			}
			//TODO: another go fn here for many command proccessing
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
		//if we panic somehow? in writing json to writer, just close ws
		defer func() {
			if r := recover(); r != nil {
				pipe.Ws.Close()
			}
		}()
		for r := range pipe.ToDeviceChan {
			pipe.Ws.WriteJSON(r)
		}
		//after closing chanel
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

/////////////////////////////////

func serverWSReqCommands(req base.WSReq, pipe *UserDevicePipe) {
	arr := make([]int64, 0, len(req.Commands))
	serveCmdsRec := true
	for _, cmd := range req.Commands {
		arr = append(arr, cmd.CmdId)
		if cmd.CmdId < 0 {
			serveCmdsRec = false
		}
	}

	if serveCmdsRec {
		cmdsRecived := base.NewCommand("CommandsReceivedToServer")
		cmdsRecived.CmdId = -1
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

func nn() {

}
