package models

import (
	"errors"
	"ms/sun/config"
	"ms/sun/helper"
	"sync"
	"time"
)

type callRespondCallback struct {
	success      func()
	error        func()
	timeoutAtMs  int   // time second // now + 5 sec
	serverCallId int64 // time nano
}

type _registerMap struct {
	sync.RWMutex
	mp               map[int64]callRespondCallback
	isRunningTimeout bool
}

var callRespondMap _registerMap

func init() {
	callRespondMap = _registerMap{
		mp: make(map[int64]callRespondCallback, 100),
	}
	intervalRunCallsTimeOutChecker()

}

func (m _registerMap) Register(callback callRespondCallback) {
	if callback.serverCallId == 0 {
		helper.DebugErr(errors.New("ERROr: In callRespondCallback, callback.serverCallId must not be 0"))
		callback.serverCallId = int64(helper.RandomUid())
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
	if !ok {
		return nil, errors.New(" serverCallId not found in  map")
	}
	return &callback, nil
}

func (m _registerMap) GetAndRemove(serverCallId int64) (*callRespondCallback, error) {
	if serverCallId == 0 {
		return nil, errors.New(" serverCallId could not be 0")
	}
	m.Lock()
	callback, ok := m.mp[serverCallId]
	delete(m.mp, serverCallId)
	m.Unlock()
	if !ok {
		return nil, errors.New(" serverCallId not found in  map")
	}
	return &callback, nil
}

func (m _registerMap) Remove(serverCallId int64) {
	m.Lock()
	delete(m.mp, serverCallId)
	m.Unlock()
}

func (m _registerMap) runSucceded(serverCallId int64) {
	helper.Debugf("runSucceded() %d", serverCallId)
	callback, err := m.GetAndRemove(serverCallId)
	if err != nil {
		return
	}
	if callback.success != nil {
		callback.success()
	}
}

func (m _registerMap) runError(serverCallId int64) {
	helper.Debugf("runSucceded() %d", serverCallId)
	callback, err := m.GetAndRemove(serverCallId)
	if err != nil {
		return
	}
	if callback.error != nil {
		callback.error()
	}
}

func (m _registerMap) setIsRunningTimeout(is bool) {
	m.Lock()
	m.isRunningTimeout = is
	m.Unlock()
}

func (m _registerMap) runErrorOfTimeouts() {
	if m.isRunningTimeout {
		return
	}
	m.setIsRunningTimeout(true)
	defer m.setIsRunningTimeout(false)

	//helper.DebugPrintln("runErrorOfTimeouts()")
	var arr []callRespondCallback

	m.RLock()
	for _, v := range m.mp {
		if v.timeoutAtMs < helper.TimeNowMs() {
			arr = append(arr, v)
		}
	}
	m.RUnlock()

	m.Lock()
	for _, v := range arr {
		delete(m.mp, v.serverCallId)
	}
	m.Unlock()

	for _, v := range arr {
		if v.error != nil {
			v.error()
		}
	}
}

//TODO: infuter make this run in parralll
func intervalRunCallsTimeOutChecker() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if config.IS_DEBUG {
					helper.DebugPrintln("ERROR PANICED RECOVRED -intervalRunCallsTimeOutChecker - ERR:: ", r)
				}
				intervalRunCallsTimeOutChecker()
			}
		}()

		for {
			time.Sleep(time.Second * 1)
			callRespondMap.runErrorOfTimeouts()
		}
	}()
}

/*
//utils
func getNextCallId() int64 {
	//return time.Now().UnixNano()
	return int64(helper.RandomUid())
}
*/
