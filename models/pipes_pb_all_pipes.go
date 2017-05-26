package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"sync"

	"github.com/gorilla/websocket"
	"ms/sun/models/x"
)

//todo change UserDevicePipe => *UserDevicePipe
type pipesMap struct {
	mp map[int]*UserDevicePipe
	m  sync.RWMutex
}

func init() {
	AllPipesMap = &pipesMap{
		mp: make(map[int]*UserDevicePipe, 100),
	}
}

var AllPipesMap *pipesMap

func (m pipesMap) SendToUser(UserId int, call base.Call) {
	pipe, ok := m.GetUserPipe(UserId)
	helper.Debugf("sending to user:%d %v %v ", UserId, ok, call.Name)
	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				//pipe.IsOpen = false
				pipe.ShutDownCompletely()
				helper.Debug("Recovered in SendToUser: ", r)
			}
		}()
		pipe.SendToUser(call)
		//pipe.ToDeviceChan <- res
	}
}

func (m pipesMap) SendToUserWithCallBack(UserId int, call base.Call, callback func()) {
	m.SendToUserWithCallBacks(UserId, call, callback, nil)
}

func (m pipesMap) SendToUserWithCallBacks(UserId int, call base.Call, callback func(), errback func()) {
	pipe, ok := m.GetUserPipe(UserId)
	helper.Debugf("sending to user:%d  - pipe is: %v --- %s ", UserId, ok, call.Name)
	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				//pipe.IsOpen = false
				pipe.ShutDownCompletely()
				helper.Debug("Recovered in SendToUser: ", r)
			}
		}()

		resCallback := callRespondCallback{
			serverCallId: call.ServerCallId,
			success:      callback,
			error:        errback,
			timeoutAtMs:  helper.TimeNowMs() + 5000,
		}

		CallRespndMap.Register(resCallback)

		pipe.SendToUser(call)
	} else if errback != nil {
		errback()
	}
}

func (m pipesMap) GetUserPipe(UserId int) (*UserDevicePipe, bool) {
	m.m.RLock()
	pipe, ok := m.mp[UserId]
	m.m.RUnlock()
	return pipe, ok
}

func (m pipesMap) ShutDownUser(UserId int) {
	pipe, ok := m.GetUserPipe(UserId)
	if ok {
		pipe.ShutDown()
		m.m.RLock()
		delete(m.mp, UserId)
		m.m.RUnlock()
	}
}

//adds a new pip
func (m pipesMap) ServeNewHttpWsForUser(UserId int, ws *websocket.Conn) {
	pipe := &UserDevicePipe{
		UserId:       UserId,
		ToDeviceChan: make(chan x.PB_CommandToClient, 10),
		IsOpen:       true,
		Ws:           ws,
	}

	helper.DebugPrintln("serving user ws for user: ", UserId)

	pipe.ServeIncomingCalls()
	pipe.ServeSendToUserDevice()

	m.AddUserPipe(UserId, pipe)

	OnNewUserWsConnected(UserId) //do and send Stored Cmds in here
	MessageModel.FlushAllStoredMessagesToUser(UserId)
	MessageModel.FlushAllReceivedMsgsToPeerToUser(UserId)
	MessageModel.FlushAllDeletedMsgsToUser(UserId)
	MessageModel.FlushAllSeenMsgsByPeerToUser(UserId)

}

func (m pipesMap) AddUserPipe(UserId int, pipe *UserDevicePipe) {
	m.m.Lock()
	m.mp[UserId] = pipe
	m.m.Unlock()
}

func (m pipesMap) DeleteUserPipe(UserId int) {
	m.m.Lock()
	delete(m.mp, UserId)
	m.m.Unlock()
}
