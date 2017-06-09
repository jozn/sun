package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

var chanNewChatMsgsBuffer = make(chan newChatMsgDelayer, 10000)

//var chanNewMsgPushEventsBuffer = make(chan *x.MsgPushEvent, 20000)

type newChatMsgDelayer struct {
	msgPB      *x.PB_Message
	fromUserId int
	toUserId   int
	roomKey    string
	hashId     int
	uid        int
}

func init() {
	go batcheNewMsgsBufferProceess()
}

func batcheNewMsgsBufferProceess() {
	const siz = 1000
	arr := make([]newChatMsgDelayer, 0, siz)
	go func() {
		for m := range chanNewChatMsgsBuffer {
			arr = append(arr, m)
		}
	}()

	for {
		time.Sleep(time.Millisecond * 10)
		if len(arr) > 0 {
			pre := arr
			arr = make([]newChatMsgDelayer, 0, siz)
			processNewChatMsgBuffer(pre)
		}
	}
}

func processNewChatMsgBuffer(msgsDelays []newChatMsgDelayer) {
	mp := make(map[int][]newChatMsgDelayer)

	for _, mb := range msgsDelays {
		mp[mb.toUserId] = append(mp[mb.toUserId], mb)
	}

	allMsgsRows := make([]x.Message, 0, len(msgsDelays))
	allMsgsKeys := make([]string, 0, len(msgsDelays))
	for _, md := range msgsDelays {
		b := md.msgPB

		mRow := PBConv_PB_Message_toNew_Message(b)
		mRow.Uid = md.uid

		allMsgsRows = append(allMsgsRows, mRow)
		allMsgsKeys = append(allMsgsKeys, b.MessageKey)
	}

	toPushMsgsArr := make([]x.MsgPush, 0, len(msgsDelays))
	for _, md := range msgsDelays {
		p := x.MsgPush{
			Id:            0,
			Uid:           helper.RandomSeqUid(),
			ToUser:        md.toUserId,
			MsgUid:        md.uid,
			CreatedTimeMs: helper.TimeNowMs(),
		}
		toPushMsgsArr = append(toPushMsgsArr, p)
	}

	x.MassInsert_Message(allMsgsRows, base.DB)
	x.MassInsert_MsgPush(toPushMsgsArr, base.DB)
	//cache it
	_, err := x.NewMessage_Selector().MessageKey_In(allMsgsKeys).GetRows(base.DB)
	if err != nil {
		return
	}

	for uid, msgsPB := range mp { //each user
		var msgs []*x.Message
		for _, msgPB := range msgsPB {
			m, ok := x.Store.Message_ByMessageKey(msgPB.msgPB.MessageKey)
			if !ok {
				continue
			}
			msgs = append(msgs, m)
		}
		MessageModel_PushToPipeMsgsToUser(uid, msgs)

	}

}

func MessageModel_PushToPipeMsgsToUser(UserId int, messages []*x.Message) {
	/*if len(messages) == 0 || !AllPipesMap.IsPipeOpen(UserId) {
		return
	}*/

	msgPusher := sMsgPusher{
		messages: messages,
		toUserId: UserId,
	}
	msgPusher.pushToUser()

}
