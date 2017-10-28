package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

type liveOfflineBuffer struct {
	//HereDirect        chan x.DirectOffline
	HereDirectDelayer_DEP chan OfflineDelayer
	HereDirectDelayer     chan x.DirectOffline
	StoredDirect          chan x.DirectOffline

	HereGroup   chan x.DirectOffline
	StoredGroup chan x.DirectOffline
}

type OfflineDelayer struct {
	directUpdate x.DirectOffline
	/*fromUserId   int
	toUserId     int
	roomKey      string
	hashId       int
	uid          int
	msgFileRowId int*/
}

///////////////////////////////

var LiveOfflineFramer = liveOfflineBuffer{
	//HereDirect:        make(chan x.DirectOffline, 10000),//dep?? use HereDirectDelayer_DEP
	//HereDirectDelayer_DEP: make(chan OfflineDelayer, 10000),
	HereDirectDelayer: make(chan x.DirectOffline, 10000),
	StoredDirect:      make(chan x.DirectOffline, 10000),
	HereGroup:         make(chan x.DirectOffline, 10000),
	StoredGroup:       make(chan x.DirectOffline, 10000),
}

func init() {
	LiveOfflineFramer.StartLoops()
}

func (m *liveOfflineBuffer) StartLoops() {
	go m.loopHereDirect()
	go m.loopDirectToUser()
}

//can causes panic
func (m *liveOfflineBuffer) loopHereDirect() {
	const siz = 50000
	arr := make([]x.DirectOffline, 0, siz)
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
				arr = make([]x.DirectOffline, 0, siz)
				go m.saveNewChatDirectBuffer(pre)
				logChat.Printf(".HereDirectDelayer_DEP is saving for %v", pre)
			}
		}
	}
}

func (m *liveOfflineBuffer) saveNewChatDirectBuffer(msgsDelays []x.DirectOffline) {
	defer helper.JustRecover()

	if len(msgsDelays) == 0 {
		return
	}

	logs := make([]x.DirectOffline, 0, len(msgsDelays))
	for _, md := range msgsDelays {
		logs = append(logs, md)
	}
	err := x.MassInsert_DirectOffline(logs, base.DB)
	helper.DebugErr(err)

	if err != nil {
		logChat.Printf(".saveNewChatDirectBuffer() has error for saving mass OfflineDelayer - Err: %s", err)
	}

	//fixme: is it better to use just use anotehr mechanism - this has data racing??
	for _, md := range msgsDelays {
		m.StoredDirect <- md
	}
}

func (m *liveOfflineBuffer) loopDirectToUser() {
	const siz = 50000
	arr := make([]x.DirectOffline, 0, siz)
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
				fmt.Printf("batch of loopDirectToUser - cnt:%d - len:%d \n", cnt, len(arr))
				pre := arr
				arr = make([]x.DirectOffline, 0, siz)
				go _livePush_sendToUsersOfflineFrame(pre)
			}
		}
	}
}

func _livePush_sendToUsersOfflineFrame(logs []x.DirectOffline) {
	defer helper.JustRecover()

	if len(logs) == 0 {
		return
	}

	mp := make(map[int][]*x.DirectOffline, len(logs))
	msgIds := make([]int, 0, len(logs))
	for _, l := range logs {
		l2 := l
		mp[l.ToUserId] = append(mp[l.ToUserId], &l2)
		if l.MessageId > 0 {
			msgIds = append(msgIds, l.MessageId)
		}
	}

	x.Store.PreLoadDirectMessageByMessageIds(msgIds)

	for UserId, lgs := range mp { //each user
		if AllPipesMap.IsPipeOpen(UserId) && len(lgs) > 0 {
			res := ViewPush_DirectOfflinesList_To_GetDirectUpdatesView(UserId, lgs)

			pb_res := x.PB_AllLivePushes{
				Offline_Messagings: res,
			}
			PushToUserLiveData(UserId, pb_res)
			if config.IS_DEBUG {
				logChat.Printf("_livePush_sendToUsersOfflineFrame() is sending to user: %s", pb_res)
				//logLivePush.Printf("to user: %d - Data: %s\n", UserId, helper.ToJsonPerety(&pb_res))

				fmt.Printf("send to user: %d PushViews : %s", UserId, helper.ToJson(res))
			}

			/*cmd := NewPB_CommandToClient_WithData(PB_PushHolderView, res)
			AllPipesMap.SendToUser(UserId, cmd)
			if config.IS_DEBUG {
				logChat.Printf("_livePush_sendToUsersOfflineFrame() is sending to user: %s", cmd)

				fmt.Printf("send to user: %d PushViews : %s", UserId, helper.ToJson(res))
			}*/
		}
	}
}
