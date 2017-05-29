package models

import (
	"github.com/golang/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

var chanNewChatMsgsBuffer = make(chan newChatMsgBuffer, 10000)

type newChatMsgBuffer struct {
	msgPB      *x.PB_Message
	fromUserId int
	toUserId   int
	roomKey    string
	hashId     int
}

func batcheNewMsgsBufferProceess() {
	const siz = 1000
	arr := make([]newChatMsgBuffer, 0, siz)
	go func() {
		for m := range chanNewChatMsgsBuffer {
			arr = append(arr, m)
		}
	}()

	for {
		time.Sleep(time.Millisecond * 5)
		if len(arr) > 0 {
			pre := arr
			arr = make([]newChatMsgBuffer, 0, siz)
			processNewChatMsgBuffer(pre)
		}
	}
}

func processNewChatMsgBuffer(msgsBuff []newChatMsgBuffer) {
	mp := make(map[int][]newChatMsgBuffer)

	for _, mb := range msgsBuff {
		mp[mb.toUserId] = append(mp[mb.toUserId], mb)
	}

	allMsgsRows := make([]x.Message, 0, len(msgsBuff))
	allMsgsKeys := make([]string, 0, len(msgsBuff))
	for _, bm := range msgsBuff {
		b := bm.msgPB
		bytes, _ := proto.Marshal(b)
		json := helper.ToJson(b)
		mRow := x.Message{
			Id:            0,
			UserId:        b.RoomKey,
			MessageKey:    b.MessageKey,
			RoomKey:       b.RoomKey,
			MessageType:   b.RoomTypeId,
			RoomType:      b.RoomTypeId,
			DataPB:        bytes,
			DataJson:      json,
			CreatedTimeMs: b.CreatedMs,
		}

		allMsgsRows = append(allMsgsRows, mRow)
		allMsgsKeys = append(allMsgsKeys, b.MessageKey)
	}

	x.MassInsert_Message(allMsgsRows, base.DB)
	//cache it
	_, err := x.NewMessage_Selector().MessageKey_In(allMsgsKeys).GetRows(base.DB)
	if err != nil {
		return
	}

	for uid, msgsPB := range mp { //each user
		for _, msgPB := range msgsPB {
			var msgs []*x.Message
			m, err := x.Store.Message_ByMessageKey(msgPB.msgPB.MessageKey)
			if err != nil {
				continue
			}
			msgs = append(msgs, m)

			MessageModel_PushToPipeMsgsToUser(uid, msgs)
		}

	}

}

func MessageModel_PushToPipeMsgsToUser(i int, messages []*x.Message) {

}
