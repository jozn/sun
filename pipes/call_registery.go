package pipes

import (
	"errors"
	"sync"
)

type callRespondCallback struct {
	succ         func()
	err          func()
	tiomeoutAt   uint64 // time nanon
	serverCallId uint64 // time nano
}

var callRespndMap _registerMap

func init() {
	callRespndMap = _registerMap{
		mp: make(map[uint64]callRespondCallback, 100),
	}
}

type _registerMap struct {
	*sync.RWMutex
	mp map[uint64]callRespondCallback
}

func (m _registerMap) Register(callback callRespondCallback) {
	if callback.serverCallId == 0 {
		return
	}
	m.Lock()
	m.mp[callback.serverCallId] = callback
	m.Unlock()
}

func (m _registerMap) Get(serverCallId uint64) (callRespondCallback, error) {
	if serverCallId == 0 {
		return nil, errors.New(" serverCallId could not be 0")
	}
	m.RLock()
	callback, err := m.mp[serverCallId]
	if err != nil {
		return nil, err
	}
	return callback, nil
}
