package models

import (
	"encoding/json"
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
)

type messageLoadOne struct {
	Message Message
	User    User
}
type _messageModelImple int

var MessageModel _messageModelImple

func (e _messageModelImple) SendAndStoreMessage(ToUserId int, msg MessagesTableFromClient) {
	if msg.CreatedMs == 0 {
		t := helper.TimeNowMs()
		msg.CreatedMs = t
	}
	data := struct {
		Message MessagesTableFromClient
		User    *UserViewSyncAndMe
	}{}
	data.Message = msg
	data.User = Views.UserViewSync(ToUserId, msg.UserId)
	call := base.NewCallWithData("MsgAddOne", data)
	//call.SetData(msg)

	succ := func() {
		helper.DebugPrintln("SUCESS OF SendAndStoreMessage")
		NewMessage_Deleter().MessageKey_EQ(msg.MessageKey).ToUserId_EQ(ToUserId).Delete(base.DB)
		m := Message{
			FromUserID: msg.UserId,
			MessageKey: msg.MessageKey,
		}
		MessageModel.SendAndStoreMsgsReceivedToPeer([]Message{m})
	}

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
	MessageModel.StoreMessage(ToUserId, msg)

}

func (e _messageModelImple) StoreMessage(ToUserId int, msg MessagesTableFromClient) {
	if msg.CreatedMs == 0 {
		t := helper.TimeNowMs()
		msg.CreatedMs = t
	}

	msgT := Message{
		FromUserID: msg.UserId,
		ToUserId:   ToUserId,
		MessageKey: msg.MessageKey,
		Data:       helper.ToJson(msg),
		TimeMs:     msg.CreatedMs,
	}

	err := msgT.Insert(base.DB)
	helper.DebugPrintln(err)
}

func (e _messageModelImple) FlushAllStoredMessagesToUser(ToUserId int) {
	helper.DebugPrintln("FlushAllStoredMessagesToUser()")

	msgRows, err := NewMessage_Selector().ToUserId_EQ(ToUserId).OrderBy_Id_Asc().GetRows(base.DB) // first msgs rows first in slice
	if err != nil || len(msgRows) == 0 {
		return
	}

	mapOfSenders := make(map[int]bool, len(msgRows))
	for _, m := range msgRows {
		mapOfSenders[m.FromUserID] = true
	}

	//arr := make([]int,0,len(mp))
	users := make([]*UserViewSyncAndMe, 0, len(mapOfSenders))
	for id, _ := range mapOfSenders {
		//arr = append(arr,id)
		users = append(users, Views.UserViewSync(ToUserId, id))
	}

	//json unserilaize
	msgsRes := make([]*MessagesTableFromClient, 0, len(mapOfSenders))
	for _, m := range msgRows {
		msgView := &MessagesTableFromClient{}
		json.Unmarshal([]byte(m.Data), msgView)
		msgsRes = append(msgsRes, msgView)
	}

	succ := func() {
		helper.DebugPrintln("SUCESS OF FlushAllStoredMessagesToUser()")
		arrMsgs := make([]string, 0, len(mapOfSenders))
		for _, m := range msgRows {
			arrMsgs = append(arrMsgs, m.MessageKey)
		}
		NewMessage_Deleter().ToUserId_EQ(ToUserId).MessageKey_In(arrMsgs).Delete(base.DB)
		MessageModel.SendAndStoreMsgsReceivedToPeer(msgRows)
	}

	dataSend := struct {
		Messages []*MessagesTableFromClient
		Users    []*UserViewSyncAndMe
	}{
		msgsRes, users,
	}

	call := base.NewCallWithData("MsgAddMany", dataSend)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
}

func (e _messageModelImple) SendAndStoreMsgsReceivedToPeer(msgs []Message) {
	groupByUsers := make(map[int][]Message)
	for _, msg := range msgs {
		//FromUserID,ok :=groupByUsers[msg.FromUserID]
		groupByUsers[msg.FromUserID] = append(groupByUsers[msg.FromUserID], msg)
	}

	var metaArr []MsgReceivedToPeer
	for toUser, msgs2 := range groupByUsers {
		for _, m := range msgs2 {
			met := MsgReceivedToPeer{
				ToUserId:   toUser,
				PeerUserId: m.FromUserID,
				RoomKey:    m.RoomKey,
				MsgKey:     m.MessageKey,
				AtTime:     helper.TimeNow(),
			}
			metaArr = append(metaArr, met)
			//met.Insert(base.DB)
		}
	}
	err := MassInsert_MsgReceivedToPeer(metaArr, base.DB)
	helper.DebugPrintln(err)

}

///////////// Utils //////////////////

//format "p142_1569"
func UserIdsToRoomKey(UserId1, UserId2 int) string {
	if UserId2 < UserId1 {
		UserId1, UserId2 = UserId2, UserId1
	}
	return fmt.Sprintf("p%d_%d", UserId1, UserId2)
}
