package models

import (
	"sync"

	"github.com/gorilla/websocket"
	//"time"
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/config"
	"ms/sun/helper"
	"ms/sun/models/x"
	"os"
)

type UserDevicePipe struct {
	UserId         int
	IsOpen         bool
	Ws             *websocket.Conn
	ToDeviceChan   chan x.PB_CommandToClient
	FromDeviceChan chan x.PB_CommandToServer // NOT NEEDED?
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
			//var call base.Call
			messageType, bytes, err := pipe.Ws.ReadMessage() //blocking

			helper.Debug("messageType: ", " ::", messageType, string(bytes))
			if messageType == websocket.CloseMessage || err != nil {
				pipe.ShutDownCompletely()
				helper.Debugf("closeing pip for userId: %v , messageType:%v , err: %v", pipe.UserId, messageType, err)
				return
			}

			if messageType == websocket.TextMessage {

			}

			if messageType == websocket.BinaryMessage {
				pb := &x.PB_CommandToServer{}
				err := proto.Unmarshal(bytes, pb)
				if err == nil {
					if config.IS_DEBUG {
						wsDebugLog("-> from device:", pb.Command, helper.ToJson(pb))
					}
					serverWSReqCalls(pb, pipe)
				}
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
				wsDebugLog("<- to device:", pipe.UserId, helper.ToJson(r))
			}
            bts, err := proto.Marshal(r)
            if err == nil{
                pipe.Ws.WriteMessage(websocket.BinaryMessage , bts)
            }
			//pipe.Ws.WriteJSON(r)
		}
		//after closing chanel
		//fixme :not neccossroy waht about .ShutDownCompletely()?
        pipe.ShutDownCompletely()
		//helper.Debug("closed: ", err)
	}()
}

func (pipe *UserDevicePipe) SendToUser(resCall x.PB_CommandToClient) {
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

func serverWSReqCalls(reqCall *x.PB_CommandToServer, pipe *UserDevicePipe) {

	if reqCall.Command == "CallReceivedToClient" {
		CallRespndMap.runSucceded(reqCall.ClientCallId)
		return
	}

	if reqCall.ClientCallId != 0 {
		callReceived := x.PB_CommandToClient{
			Command:      "CallReceivedToServer",
			ServerCallId: reqCall.ClientCallId,
		}

		AllPipesMap.SendToUser(pipe.UserId, callReceived)
	}

	//reqCall.UserId = pipe.UserId
	fnCall := CallMapRouter[reqCall.Command]
	helper.DebugPrintln("serving Cmd: ", reqCall.Command, " Userid: ", pipe.UserId)
	if fnCall != nil {
		fnCall(reqCall, pipe)
	} else { //send call func not found -- in debug
		helper.DebugPrintln("ERROR command nor registerd in CallMapRouter : ", reqCall.Command)
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
