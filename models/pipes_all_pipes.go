package models

import (
	"ms/sun/helper"
	"sync"

	"ms/sun/models/x"

	"github.com/gorilla/websocket"
	"ms/sun/base"
	"time"
    "fmt"
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

func (m *pipesMap) IsPipeOpen(UserId int) bool {
	pipe, ok := m.GetUserPipe(UserId)
	if ok && pipe.IsOpen {
		return true
	}
	return false
}

func (m *pipesMap) SendToUser(UserId int, cmd x.PB_CommandToClient) {
	pipe, ok := m.GetUserPipe(UserId)

	//helper.Debugf("pipesMap.SendToUser() to user:%d %v %v ", UserId, ok, cmd.Command)
	logPipes.Infof("pipesMap.SendToUser() no callback to user:%d %v %v ", UserId, ok, cmd.Command)

	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				//pipe.IsOpen = false
				pipe.ShutDownCompletely()
				helper.Debug("Recovered in SendToUser: ", r)
			}
		}()
		pipe.SendToUser(cmd)
		//pipe.ToDeviceChan <- res
	}
}

func (m *pipesMap) SendToUserWithCallBack(UserId int, call x.PB_CommandToClient, callback func()) {
	m.SendToUserWithCallBacks(UserId, call, callback, nil)
}

func (m *pipesMap) SendToUserWithCallBacks(UserId int, call x.PB_CommandToClient, callback func(), errback func()) {
	pipe, ok := m.GetUserPipe(UserId)
	//helper.Debugf("sending to user:%d  - pipe is: %v --- %s ", UserId, ok, call.Command)
	logPipes.Infof("pipesMap sending to user withCallbacks:%d %v %v ", UserId, ok, call.Command)

	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				//pipe.IsOpen = false
				pipe.ShutDownCompletely()
				helper.Debug("Recovered in SendToUser: ", r)
				logPipes.Info("Recovered in SendToUser: ", r)
			}
		}()

		resCallback := callRespondCallback{
			serverCallId: call.ServerCallId,
			success:      callback,
			error:        errback,
			timeoutAtMs:  helper.TimeNowMs() + 5000,
		}

		callRespondMap.register(resCallback)

		pipe.SendToUser(call)
	} else {
		logPipes.Info(" pipesMap SendToUserWithCallBacks: user is not connected (just run errorback if providd) ")
		if errback != nil {
			errback()
		}
	}
}

func (m *pipesMap) GetUserPipe(UserId int) (*UserDevicePipe, bool) {
	m.m.RLock()
	pipe, ok := m.mp[UserId]
	m.m.RUnlock()
	return pipe, ok
}

func (m *pipesMap) ShutDownUser(UserId int) {
	logPipes.Info("pipesMap ShutDownUser ", UserId)

	pipe, ok := m.GetUserPipe(UserId)
	if ok {
		pipe.ShutDown()
		m.DeleteUserPipe(UserId)
	}
}

//adds a new pip
func (m *pipesMap) ServeNewHttpWsForUser(UserId int, ws *websocket.Conn) {
	pipe := &UserDevicePipe{
		UserId:       UserId,
		ToDeviceChan: make(chan x.PB_CommandToClient, 10),
		IsOpen:       true,
		Ws:           ws,
	}

	helper.DebugPrintln("serving user ws for user: ", UserId)
	logPipes.Debug("serving user ws for user: ", UserId)

	pipe.ServeIncomingCalls()
	pipe.ServeSendToUserDevice()

	m.AddUserPipe(UserId, pipe)

    go func() {
        for {
            time.Sleep(time.Second * 5)
            fmt.Println("sendin pushes")
            logs, err := x.NewDirectLog_Selector().ToUserId_Eq(UserId).GetRows(base.DB)
            if err == nil {
                res := PushView_directLogsTo_PB_ChangesHolderView(UserId, logs)
                cmd := NewPB_CommandToClient_WithData("PB_PushHolderView", res)
                AllPipesMap.SendToUser(UserId, cmd)
            }
        }
    }()


	/*flusher_flushPushMsgs(UserId)
	flusher_flushPushEvents(UserId)
	flusher_flushPushNotifications(UserId)*/

	/*OnNewUserWsConnected(UserId) //do and send Stored Cmds in here
	MessageModel.FlushAllStoredMessagesToUser(UserId)
	MessageModel.FlushAllReceivedMsgsToPeerToUser(UserId)
	MessageModel.FlushAllDeletedMsgsToUser(UserId)
	MessageModel.FlushAllSeenMsgsByPeerToUser(UserId)*/

}

func (m *pipesMap) AddUserPipe(UserId int, pipe *UserDevicePipe) {
	logPipes.Info("pipesMap AddUserPipe UserId ", UserId)
	m.m.Lock()
	m.mp[UserId] = pipe
	m.m.Unlock()
}

func (m *pipesMap) DeleteUserPipe(UserId int) {
	logPipes.Info("pipesMap DeleteUserPipe UserId ", UserId)
	m.m.Lock()
	delete(m.mp, UserId)
	m.m.Unlock()
}
