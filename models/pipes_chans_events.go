package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
	"time"
)

var chanNewMsgPushEventsBuffer = make(chan x.MsgPushEvent, 100000)

func init() {
	go batchechanNewMsgPushEventsBuffer()
}

func batchechanNewMsgPushEventsBuffer() {
	const siz = 50000
	arr := make([]x.MsgPushEvent, 0, siz)
	go func() {
		for m := range chanNewMsgPushEventsBuffer {
			arr = append(arr, m)
		}
	}()

	for {
		time.Sleep(time.Millisecond * 10)
		if len(arr) > 0 {
			//todo this has data racing problem + do it too for msg buffer
			pre := arr
			arr = make([]x.MsgPushEvent, 0, siz)
			processchanNewMsgPushEventsBuffer(pre)
		}
	}
}

func processchanNewMsgPushEventsBuffer(msgEvents []x.MsgPushEvent) {
	mp := make(map[int][]x.MsgPushEvent)
	uids := make([]int, 0, len(msgEvents))

	for _, e := range msgEvents {
		uids = append(uids, e.Uid)
		mp[e.ToUserId] = append(mp[e.ToUserId], e)
	}

	x.MassInsert_MsgPushEvent(msgEvents, base.DB)

	_, err := x.NewMsgPushEvent_Selector().Uid_In(uids).GetRows(base.DB)
	if err != nil {
		return
	}

	for uid, userEvents := range mp { //each user
		MessageModel_PushToPipeMsgEventsToUser(uid, userEvents)
	}

}

func MessageModel_PushToPipeMsgEventsToUser(UserId int, eventsRows []x.MsgPushEvent) {
	if len(eventsRows) == 0 || !AllPipesMap.IsPipeOpen(UserId) {
		return
	}

	pbEvents := []*x.PB_MsgEvent{}

	for _, m := range eventsRows {
		pbEv := PBConv_MsgPushEvent_toNew_PB_MsgEvent(&m)

		pbEvents = append(pbEvents, &pbEv)
	}

	pushReq := &x.PB_PushMsgEvents{
		Push:   nil,
		Events: pbEvents,
	}

	cmd := NewPB_CommandToClient_WithData(PB_PushMsgEvents, pushReq)
	callback := func() {
		uids := make([]int, 0, len(eventsRows))
		for _, m := range eventsRows {
			uids = append(uids, m.Uid)
		}
	}

	AllPipesMap.SendToUserWithCallBack(UserId, cmd, callback)

}
