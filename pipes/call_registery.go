package pipes

import (
	"errors"
	"sync"
	"time"
)

type callRespondCallback struct {
	succ         func()
	err          func()
	tiomeoutAt   int64 // time second // now + 10 sec
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
	callback.tiomeoutAt = time.Now().Unix() + 10
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
	if ok {
		return nil, errors.New(" serverCallId not found in  map")
	}
	return &callback, nil
}

//utils
func getNextCallId() int64 {
	return time.Now().UnixNano()
}