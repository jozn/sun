package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

type chatUpdater struct {
	HereDirect        chan x.DirectLog
	HereDirectDelayer chan newChatMsgDelayer
	StoredDirect      chan x.DirectLog

	HereGroup   chan x.DirectLog
	StoredGroup chan x.DirectLog
}

type newChatMsgDelayer struct {
	directLog    x.DirectLog
	fromUserId   int
	toUserId     int
	roomKey      string
	hashId       int
	uid          int
	msgFileRowId int
}

///////////////////////////////

var LogUpdater = chatUpdater{
	HereDirect:        make(chan x.DirectLog, 10000),
	HereDirectDelayer: make(chan newChatMsgDelayer, 10000),
	StoredDirect:      make(chan x.DirectLog, 10000),
	HereGroup:         make(chan x.DirectLog, 10000),
	StoredGroup:       make(chan x.DirectLog, 10000),
}

func (m *chatUpdater) StartLoops() {
	m.loopHereDirect()
}

func (m *chatUpdater) loopHereDirect() {
	const siz = 50000
	arr := make([]newChatMsgDelayer, 0, siz)
	cnt := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case m := <-m.HereDirectDelayer:
			arr = append(arr, m)

		case <-ticker.C:
			if len(arr) > 0 {
				cnt++
				fmt.Printf("batch of chanNewChatMsgsBuffer - cnt:%d - len:%d \n", cnt, len(arr))
				pre := arr
				arr = make([]newChatMsgDelayer, 0, siz)
				go m.saveNewChatDirectBuffer(pre)
			}
		}
	}
}

func (m *chatUpdater) saveNewChatDirectBuffer(msgsDelays []newChatMsgDelayer) {
	defer helper.JustRecover()

	if len(msgsDelays) == 0 {
		return
	}

	logs := make([]x.DirectLog, 0, len(msgsDelays))
	for _, md := range msgsDelays {
		logs = append(logs, md.directLog)
	}
	x.MassInsert_DirectLog(logs, base.DB)

	//fixme: is it better to use just use anotehr mechanism - this has data racing??
	for _, md := range msgsDelays {
		m.StoredDirect <- md.directLog
	}
}

func (m *chatUpdater) loopDirectToUser() {
	const siz = 50000
	arr := make([]x.DirectLog, 0, siz)
	cnt := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case m := <-m.StoredDirect:
			arr = append(arr, m)

		case <-ticker.C:
			if len(arr) > 0 {
				cnt++
				fmt.Printf("batch of chanNewChatMsgsBuffer - cnt:%d - len:%d \n", cnt, len(arr))
				pre := arr
				arr = make([]x.DirectLog, 0, siz)
				go m.sendToUsersUpdates(pre)
			}
		}
	}
}

func (m *chatUpdater) sendToUsersUpdates(logs []x.DirectLog) {
	defer helper.JustRecover()

	if len(logs) == 0 {
		return
	}

	mp := make(map[int][]x.DirectLog)
	for _, l := range logs {
		mp[l.ToUserId] = append(mp[l.ToUserId], l)
	}

	for uid, _ := range mp { //each user
		if AllPipesMap.IsPipeOpen(uid) {

		}
	}
}
