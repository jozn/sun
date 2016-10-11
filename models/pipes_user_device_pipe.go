package models

import (
	"github.com/gorilla/websocket"
	"sync"
	//"time"
	"encoding/json"
	"fmt"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/helper"
	"os"
)

type UserDevicePipe struct {
	UserId         int
	IsOpen         bool
	Ws             *websocket.Conn
	ToDeviceChan   chan base.Call
	FromDeviceChan chan base.Call // NOT NEEDED?
	m              sync.RWMutex
	hasShutDown    bool
}

func (pipe *UserDevicePipe) ServeIncomingCalls() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				helper.Debug("Recovered in ws messaging clinet request", r)
				pipe.ShutDownCompletely()
			}
		}()

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
					serverWSReqCalls(call, pipe)
				}
				if config.IS_DEBUG {
					wsDebugLog("-> from device:",pipe.UserId, string(bytes))
				}
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
			//helper.Debug("sending to user fom ToDeviceChan Command size ", r)
			if config.IS_DEBUG {
				wsDebugLog("<- to device:",pipe.UserId, helper.ToJson(r))
			}
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
	if pipe.hasShutDown == true {
		return
	}
	pipe.Ws.Close()
	close(pipe.ToDeviceChan)
	pipe.m.Lock()
	pipe.hasShutDown = true
	pipe.m.Unlock()
}

func (pipe *UserDevicePipe) ShutDownCompletely() {
	defer helper.JustRecover() //maybe double close chanel
	pipe.ShutDown()
	AllPipesMap.ShutDownUser(pipe.UserId)
}

/////////////// Commands handler /////////////////

func serverWSReqCalls(reqCall base.Call, pipe *UserDevicePipe) {

	if reqCall.Name == "CallReceivedToClient" {
		CallRespndMap.runSucceded(reqCall.ServerCallId)
		return
	}

	if reqCall.ClientCallId != 0 {
		callReceived := base.Call{
			Name:         "CallReceivedToServer",
			ClientCallId: reqCall.ClientCallId,
			ServerCallId: 0,
		}

		AllPipesMap.SendToUser(pipe.UserId, callReceived)
	}

	reqCall.UserId = pipe.UserId
	fnCall := base.CallMapRouter[reqCall.Name]
	helper.Debug("serving Cmd: ", reqCall.Name, " Userid: ", pipe.UserId)
	if fnCall != nil {
		fnCall(reqCall)
	} else { //send call func not found -- in debug

	}

}

var _wsLogFile *os.File
var wsDebugLog = func(strings ...interface{}) {
	if config.IS_DEBUG {
		if _wsLogFile == nil {
			_wsLogFile, _ = os.OpenFile("./logs/ws_"+helper.IntToStr(helper.TimeNow())+".json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		}
		_wsLogFile.WriteString(fmt.Sprintln(strings...))
		_wsLogFile.Sync()
	}
}
