package models

import (
	"github.com/golang/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

var chanNewChatMsgsBuffer = make(chan newChatMsgBuffer, 10000)

//var chanNewMsgPushEventsBuffer = make(chan *x.MsgPushEvent, 20000)

type newChatMsgBuffer struct {
	msgPB      *x.PB_Message
	fromUserId int
	toUserId   int
	roomKey    string
	hashId     int
	uid        int
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
		/*bytes, _ := proto.Marshal(b)
		  json := helper.ToJson(b)
		  mRow := x.Message{
		      Id:            0,
		      Uid:           bm.uid,
		      UserId:        b.RoomKey,
		      MessageKey:    b.MessageKey,
		      RoomKey:       b.RoomKey,
		      MessageType:   b.RoomTypeId,
		      RoomType:      b.RoomTypeId,
		      DataPB:        bytes,
		      DataJson:      json,
		      CreatedTimeMs: b.CreatedMs,
		  }*/

		mRow := PBConv_PB_Message_toNew_Message(b)

		allMsgsRows = append(allMsgsRows, mRow)
		allMsgsKeys = append(allMsgsKeys, b.MessageKey)
	}

	toPushMsgsArr := make([]x.MsgPush, 0, len(msgsBuff))
	for _, bm := range msgsBuff {
		p := x.MsgPush{
			Id:            0,
			Uid:           helper.RandomUid(),
			ToUser:        bm.toUserId,
			MsgUid:        bm.uid,
			CreatedTimeMs: 0,
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
		for _, msgPB := range msgsPB {
			var msgs []*x.Message
			m, ok := x.Store.Message_ByMessageKey(msgPB.msgPB.MessageKey)
			if !ok {
				continue
			}
			msgs = append(msgs, m)

			MessageModel_PushToPipeMsgsToUser(uid, msgs)
		}

	}

}

func MessageModel_PushToPipeMsgsToUser(UserId int, messages []*x.Message) {
	if len(messages) == 0 {
		return
	}
	if AllPipesMap.IsPipeOpen(UserId) {
		//cmd := NewPB_CommandToClient("AddManyMsgs")

		pbMsgs := []*x.PB_Message{}
		userIds := make(map[int]bool)
		pbUsers := []*x.PB_UserWithMe{}

		for _, m := range messages {
			pbMsg := &x.PB_Message{}
			err := proto.Unmarshal(m.DataPB, pbMsg)
			if err == nil {
				pbMsgs = append(pbMsgs, pbMsg)
			}
			userIds[m.UserId] = true
		}

		for uid, _ := range userIds {
			pbUsers = append(pbUsers, (PBNew_PB_UserWithMe(uid, UserId)))
		}

		pushReq := &x.PB_PushMsgAddMany{
			Push:     nil,
			Messages: pbMsgs,
			Users:    pbUsers,
		}

		cmd := NewPB_CommandToClient_WithData("AddManyMsgs", pushReq)
		callback := func() {
			messageModel_onAfterMsgsHasPushedToUser(UserId, messages)
			//messageModel_msgsRecicedToUserAddEvents(UserId, messages)
		}

		AllPipesMap.SendToUserWithCallBack(UserId, cmd, callback)
	}
}

func messageModel_onAfterMsgsHasPushedToUser(UserId int, messages []*x.Message) {
	messageModel_msgsRecicedToPeerAddEvents(UserId, messages)
	messageModel_msgsDeleteFromServer(UserId, messages)
}

func messageModel_msgsDeleteFromServer(i int, messages []*x.Message) {

}

/*//TODO imple this
func messageModel_msgsRecicedToUserAddEvents(UserId int, messages []*x.Message) {

}*/

//////////////////////////////////////////////////////////////
func NewPB_CommandToClient(cmd string) x.PB_CommandToClient {
	p := x.PB_CommandToClient{
		ServerCallId: int64(time.Now().UnixNano()),
		Command:      cmd,
	}
	return p
}

func NewPB_CommandToClient_WithData(cmd string, protoMsg proto.Message) x.PB_CommandToClient {
	m := NewPB_CommandToClient(cmd)
	bytes, err := proto.Marshal(protoMsg)
	if err == nil {
		m.Data = bytes
	} else {
		helper.DebugPrintln("ERROR : proto.Marshal NewPB_CommandToClient_WithData, ", err)
	}
	return m
}
