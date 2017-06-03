package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
	"time"
)

var chanNewMsgPushEventsBuffer = make(chan x.MsgPushEvent, 20000)

func init() {
	batchechanNewMsgPushEventsBuffer()
}

func batchechanNewMsgPushEventsBuffer() {
	const siz = 10000
	arr := make([]x.MsgPushEvent, 0, siz)
	go func() {
		for m := range chanNewMsgPushEventsBuffer {
			arr = append(arr, m)
		}
	}()

	for {
		time.Sleep(time.Millisecond * 10)
		if len(arr) > 0 {
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
		/*for _, e := range userEvents {
		    var msgs []*x.MsgPushEvent
		    msgs = append(msgs, m)

		    MessageModel_PushToPipeMsgsToUser(uid, msgs)
		}*/

	}

}

func MessageModel_PushToPipeMsgEventsToUser(UserId int, eventsRows []x.MsgPushEvent) {
	if !AllPipesMap.IsPipeOpen(UserId) {
		return
	}
	//cmd := NewPB_CommandToClient("AddManyMsgs")

	pbEvents := []*x.PB_MsgEvent{}

	for _, m := range eventsRows {
		/*pbEv := &x.PB_MsgEvent{
			MessageKey: m.MsgKey,
			RoomKey:    m.RoomKey,
			PeerUserId: m.PeerUserId,
			EventType:  m.EventType,
			AtTime:     m.AtTime,
		}*/
		pbEv := PBConv_MsgPushEvent_toNew_PB_MsgEvent(&m)

		pbEvents = append(pbEvents, &pbEv)
	}

	pushReq := &x.PB_PushMsgEvents{
		Push:   nil,
		Events: pbEvents,
	}

	cmd := NewPB_CommandToClient_WithData("AddManyMsgsEvents", pushReq)
	callback := func() {
		uids := make([]int, 0, len(eventsRows))
		for _, m := range eventsRows {
			uids = append(uids, m.Uid)
		}
	}

	AllPipesMap.SendToUserWithCallBack(UserId, cmd, callback)

}

/*

func messageModel_msgsRecicedToPeerAddEvents(UserId int, messages []*x.Message) {
	for _, msg := range messages {
		me := x.MsgPushEvent{
			Id:         0,
			Uid:        helper.RandomSeqUid(),
			ToUserId:   msg.UserId,
			MsgUid:     msg.Uid,
			MsgKey:     msg.MessageKey,
			RoomKey:    msg.RoomKey,
			PeerUserId: UserId,
			EventType:  MESSAGE_PUSH_EVENT_RECEIVED_TO_PEER,
			AtTime:     helper.TimeNow(),
		}

		chanNewMsgPushEventsBuffer <- me
	}

}
*/
