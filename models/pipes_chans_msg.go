package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

var chanNewChatMsgsBuffer = make(chan newChatMsgDelayer, 100000)

//var chanNewMsgPushEventsBuffer = make(chan *x.MsgPushEvent, 20000)

type newChatMsgDelayer struct {
	msgPB        *x.PB_Message
	fromUserId   int
	toUserId     int
	roomKey      string
	hashId       int
	uid          int
	msgFileRowId int
}

func init() {
	go batcheNewMsgsBufferProceess()
}

func batcheNewMsgsBufferProceess() {
	const siz = 50000
	arr := make([]newChatMsgDelayer, 0, siz)
	cnt := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {

		case m := <-chanNewChatMsgsBuffer:
			arr = append(arr, m)

		case <-ticker.C:
			if len(arr) > 0 {
				cnt++
				fmt.Printf("batch of chanNewChatMsgsBuffer - cnt:%d - len:%d \n", cnt, len(arr))
				pre := arr
				arr = make([]newChatMsgDelayer, 0, siz)
				go processNewChatMsgBuffer(pre)
			}
		}
	}

	/*go func() {
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
	}*/
}

func processNewChatMsgBuffer(msgsDelays []newChatMsgDelayer) {
	helper.JustRecover()

	mp := make(map[int][]newChatMsgDelayer)

	for _, mb := range msgsDelays {
		mp[mb.toUserId] = append(mp[mb.toUserId], mb)
	}

	allMsgsRows := make([]x.Message, 0, len(msgsDelays))
	allMsgsKeys := make([]string, 0, len(msgsDelays))
	allMsgsFilesIds := make([]int, 0, len(msgsDelays))
	for _, md := range msgsDelays {
		b := md.msgPB

		mRow := PBConv_PB_Message_toNew_Message(b)
		mRow.Uid = md.uid

		if md.msgFileRowId > 0 {
			mRow.MsgFileId = md.msgFileRowId
			allMsgsFilesIds = append(allMsgsFilesIds, md.msgFileRowId)

		}

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

	//cache it (msgs)
	_, err1 := x.NewMessage_Selector().MessageKey_In(allMsgsKeys).GetRows(base.DB)
	_, err2 := x.NewMsgFile_Selector().Id_In(allMsgsFilesIds).GetRows(base.DB)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
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
