package pipes

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

//type pipesMap map[int]*UserDevicePipe

//var AllPipesMap = make(pipesMap) // make(map[int]UserDevicePipe)
var AllPipesMap = new(pipesMap) // make(map[int]UserDevicePipe)
//var AllUserDevicesPipesMap2 = make(pipesMap)// make(map[int]UserDevicePipe)

func (m pipesMap) SendToUser(UserId int, res base.WSRes) {
	pipe, ok := m.GetUserPipe(UserId)
	helper.Debugf("sending to user:%d %v %v ", UserId, ok, res.Commands)
	if ok && pipe.IsOpen {
		defer func() {
			if r := recover(); r != nil {
				//pipe.IsOpen = false
				pipe.ShutDown()
				helper.Debug("Recovered in SendToUser: ", r)
			}
		}()
		pipe.SendToUser(res)
		//pipe.ToDeviceChan <- res
	}
}

func (m pipesMap) GetUserPipe(UserId int) (*UserDevicePipe, bool) {
	m.m.RLock()
	pipe, ok := m.mp[UserId]
	m.m.RUnlock()
	return pipe, ok
}

func (m pipesMap) SendCmdToUser(UserId int, cmd *base.Command) {
	res := base.WSRes{Status: "OK", ReqKey: ""}
	res.Commands = []*base.Command{cmd}
	m.SendToUser(UserId, res)
}

func (m pipesMap) SendAndStoreCmdToUser(UserId int, cmd *base.Command) {
	//store
	SaveCmdToRedis(UserId, cmd)

	//send
	res := base.WSRes{Status: "OK", ReqKey: ""}
	res.Commands = []*base.Command{cmd}
	m.SendToUser(UserId, res)
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
func (m pipesMap) ServeUserWs(UserId int, ws *websocket.Conn) {
	pipe := UserDevicePipe{
		UserId:       UserId,
		ToDeviceChan: make(chan *base.WSRes, 10),
		IsOpen:       true,
		Ws:           ws,
	}

	helper.Debug("serving user ws for user: ", UserId)

	pipe.ServeIncomingReqs()
	pipe.ServeSendToUserDevice()

	m.AddUserPipe(UserId, &pipe)
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
