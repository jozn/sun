package models

import (
	"github.com/gorilla/websocket"
	"ms/sun/base"
	"ms/sun/helper"
	"sync"
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
    m.SendToUserWithCallBacks(UserId , call , callback,nil )
	/*pipe, ok := m.GetUserPipe(UserId)
	helper.Debugf("sending to user:%d %v %v ", UserId, ok)
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
			timeoutAtMs:  helper.TimeNowMs() + 5000,
		}

		CallRespndMap.Register(resCallback)

		pipe.SendToUser(call)
	}*/
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
		ToDeviceChan: make(chan base.Call, 10),
		IsOpen:       true,
		Ws:           ws,
	}

	helper.Debug("serving user ws for user: ", UserId)

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
	m.m.RLock()
	m.mp[UserId] = pipe
	m.m.RUnlock()
}

func (m pipesMap) DeleteUserPipe(UserId int) {
	m.m.RLock()
	delete(m.mp, UserId)
	m.m.RUnlock()
}

// Deps

/*
func (m pipesMap) SendToUser_DEP(UserId int, call base.Call) {
    pipe, ok := m.GetUserPipe(UserId)
    helper.Debugf("sending to user:%d %v %v ", UserId, ok)
    if ok && pipe.IsOpen {
        defer func() {
            if r := recover(); r != nil {
                //pipe.IsOpen = false
                pipe.ShutDown()
                helper.Debug("Recovered in SendToUser: ", r)
            }
        }()
        pipe.SendToUser(call)
        //pipe.ToDeviceChan <- res
    }
}

func (m pipesMap) SendCmdToUser(UserId int, cmd base.Call) {
    res := base.WSRes{Status: "OK", ReqKey: "", SyncedNanoId: helper.TimeNowNano()}
    res.Commands = []*base.Command{cmd}
    m.SendToUser_DEP(UserId, res)
}

//deperecated //call one by one
func (m pipesMap) SendCmdsToUser(UserId int, cmds []*base.Command) {
    res := base.WSRes{Status: "OK", ReqKey: "", SyncedNanoId: helper.TimeNowNano()}
    res.Commands = cmds
    m.SendToUser_DEP(UserId, res)
}

//deprectaed: use response
func (m pipesMap) SendAndStoreCmdToUser(UserId int, cmd *base.Command) {
    //store
    StoreCommandsToRedis(UserId, cmd)

    //send
    res := base.WSRes{Status: "OK", ReqKey: "", SyncedNanoId: helper.TimeNowNano()}
    res.Commands = []*base.Command{cmd}
    m.SendToUser_DEP(UserId, res)
}

*/
