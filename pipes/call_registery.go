package pipes

import (
	"errors"
	"sync"
	"time"
    "ms/sun/helper"
)

type callRespondCallback struct {
    succ         func()
    err          func()
    timeoutAtMs  int64 // time second // now + 10 sec
    serverCallId int64 // time nano
}

var callRespndMap _registerMap

func init() {
	callRespndMap = _registerMap{
		mp: make(map[int64]callRespondCallback, 100),
	}
}

type _registerMap struct {
	*sync.RWMutex
	mp map[int64]callRespondCallback
}

func (m _registerMap) Register(callback callRespondCallback) {
	if callback.serverCallId == 0 {
		callback.serverCallId = getNextCallId()
	}
	callback.timeoutAtMs = helper.TimeNowMs() + 5000
	m.Lock()
	m.mp[callback.serverCallId] = callback
	m.Unlock()
}

func (m _registerMap) Get(serverCallId int64) (*callRespondCallback, error) {
	if serverCallId == 0 {
		return nil, errors.New(" serverCallId could not be 0")
	}
	m.RLock()
	callback, ok := m.mp[serverCallId]
    m.RUnlock()
	if ok {
		return nil, errors.New(" serverCallId not found in  map")
	}
	return &callback, nil
}

func (m _registerMap) Remove(serverCallId int64)  {
    m.Lock()
    delete(m.mp, serverCallId)
    m.Unlock()
}

func (m _registerMap) runSucceded(serverCallId int64)  {
    callback,err := m.Get(serverCallId)
    if err != nil{
        return
    }
    if callback.err != nil{
        callback.err()
    }
}

func (m _registerMap) runErrorOfTimeouts(serverCallId int64){
    var arr []callRespondCallback

    m.RLock()
    for _,v := range m.mp {
        if v.timeoutAtMs < helper.TimeNowMs(){
            arr = append(arr,v)
        }
    }
    m.RUnlock()

    m.Lock()
    for _,v := range arr {
        delete(m.mp, v.serverCallId)
    }
    m.Unlock()

    for _,v := range arr {
        if v.err != nil{
            v.err()
        }
    }
}


func intervalRunCallsTimeOutChecker()  {

}

//utils
func getNextCallId() int64 {
	return time.Now().UnixNano()
}
