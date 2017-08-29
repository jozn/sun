package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

type chatUpdater struct {
	//HereDirect        chan x.DirectLog
	HereDirectDelayer chan logDelayer
	StoredDirect      chan x.DirectLog

	HereGroup   chan x.DirectLog
	StoredGroup chan x.DirectLog
}

type logDelayer struct {
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
	//HereDirect:        make(chan x.DirectLog, 10000),//dep?? use HereDirectDelayer
	HereDirectDelayer: make(chan logDelayer, 10000),
	StoredDirect:      make(chan x.DirectLog, 10000),
	HereGroup:         make(chan x.DirectLog, 10000),
	StoredGroup:       make(chan x.DirectLog, 10000),
}

func init() {
	LogUpdater.StartLoops()
}

func (m *chatUpdater) StartLoops() {
	go m.loopHereDirect()
}

func (m *chatUpdater) loopHereDirect() {
	const siz = 50000
	arr := make([]logDelayer, 0, siz)
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
				arr = make([]logDelayer, 0, siz)
				go m.saveNewChatDirectBuffer(pre)
			}
		}
	}
}

func (m *chatUpdater) saveNewChatDirectBuffer(msgsDelays []logDelayer) {
	defer helper.JustRecover()

	if len(msgsDelays) == 0 {
		return
	}

	logs := make([]x.DirectLog, 0, len(msgsDelays))
	for _, md := range msgsDelays {
		logs = append(logs, md.directLog)
	}
	err := x.MassInsert_DirectLog(logs, base.DB)
	helper.NoErr(err)
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

    //todo add fetching from database every 10ms and push it to x.StoredDirect
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
				go chatLogger_sendToUsersUpdates(pre)
			}
		}
	}
}

func chatLogger_sendToUsersUpdates(logs []x.DirectLog) {
	defer helper.JustRecover()

	if len(logs) == 0 {
		return
	}

	mp := make(map[int][]x.DirectLog, len(logs))
	msgIds := make([]int, 0, len(logs))
	for _, l := range logs {
		mp[l.ToUserId] = append(mp[l.ToUserId], l)
		if l.MessageId > 0 {
			msgIds = append(msgIds, l.MessageId)
		}
	}

	x.Store.PreLoadDirectMessageByMessageIds(msgIds)
	for uid, logs := range mp { //each user
		if AllPipesMap.IsPipeOpen(uid) && len(logs) > 0 {
			rowsView := directLogsToView(logs)
			push := &x.PB_PushDirectLogViewsMany{Rows: rowsView}
			cmd := NewPB_CommandToClient_WithData("PB_PushDirectLogsMany", push)
			AllPipesMap.SendToUser(uid, cmd)
		}
	}
}

func directLogsToView(logs []x.DirectLog) (res []*x.PB_DirectLogView) {
	for _, log := range logs { //each user
		res = append(res, directLogToView(log))
	}
	return
}

func directLogToView(log x.DirectLog) *x.PB_DirectLogView {
	v := &x.PB_DirectLogView{
		Row: PBConv_DirectLog_To_DirectLog(&log),
	}
	switch x.RoomLogTypeEnum(log.RoomLogTypeId) {
	case x.RoomLogTypeEnum_NEW_DIRECT_MESSAGE:
		msg, ok := x.Store.GetDirectMessageByMessageId(log.MessageId)
		if ok {
			_ = msg
			v.NewMessage = &x.PB_MessageView{}
		}
	}
	return v
}
