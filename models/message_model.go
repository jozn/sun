package models

import (
	"encoding/json"
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
)

const CLIENT_CALL_MsgsReceivedToPeerMany = "MsgsReceivedToPeerMany"
const CLIENT_CALL_MsgsDeletedFromServerMany = "MsgsDeletedFromServerMany"
const CLIENT_CALL_MsgsSeenByPeerMany = "MsgsSeenByPeerMany"
const CLIENT_CALL_MsgAddMany = "MsgAddMany"
const CLIENT_CALL_MsgAddOne = "MsgAddOne"

type messageLoadOne struct {
	Message Message
	User    User
}
type _messageModelImple int

var MessageModel _messageModelImple

func (e _messageModelImple) SendAndStoreMessage(ToUserId int, msg MessagesTableFromClient) {
	helper.DebugPrintln("SendAndStoreMessage()")
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
	call := base.NewCallWithData(CLIENT_CALL_MsgAddOne, data)
	//call.SetData(msg)

	////// perpare for Message table and callbacks ///////

	msgRow := Message{
		FromUserID: msg.UserId,
		ToUserId:   ToUserId,
		MessageKey: msg.MessageKey,
		RoomKey:    msg.RoomKey,
		Data:       helper.ToJson(msg),
		TimeMs:     msg.CreatedMs,
	}

	succ := func() {
		helper.DebugPrintln("SUCESS OF SendAndStoreMessage")
		NewMessage_Deleter().MessageKey_EQ(msg.MessageKey).ToUserId_EQ(ToUserId).Delete(base.DB)

		MessageModel.SendAndStoreMsgsReceivedToPeer([]Message{msgRow})
		MessageModel.SendAndStoreMsgsDeletedFromServer([]Message{msgRow})
	}

	errBack := func() {
		msgRow.Insert(base.DB)
	}

	AllPipesMap.SendToUserWithCallBacks(ToUserId, call, succ, errBack)
}

func (e _messageModelImple) SendAndStoreManyMessages(CuerrntUsrId int, msgs []MessagesTableFromClient) {
	helper.DebugPrintln("SendAndStoreManyMessages()")
	if len(msgs) == 0 {
		return
	}

	/// group for each toSend (Reciver) user
	groupByUser := make(map[int][]MessagesTableFromClient, len(msgs))
	allMsgsRows := make([]Message, 0, len(msgs))

	for _, msg := range msgs {
		toUser, err := RoomKeyToPeerUserId(msg.RoomKey, CuerrntUsrId)
		if err != nil || toUser < 1 {
			continue
		}

		if msg.CreatedMs == 0 {
			t := helper.TimeNowMs()
			msg.CreatedMs = t
		}

		msgRow := Message{
			FromUserID: msg.UserId,
			ToUserId:   toUser,
			MessageKey: msg.MessageKey,
			RoomKey:    msg.RoomKey,
			Data:       helper.ToJson(msg),
			TimeMs:     msg.CreatedMs,
		}

		allMsgsRows = append(allMsgsRows, msgRow)
		groupByUser[toUser] = append(groupByUser[toUser], msg)
	}

	/// Save all msgs
	MassInsert_Message(allMsgsRows, base.DB)

	///send to each online user
	for toUserId, msgsClient := range groupByUser {
		MessageModel.SendManyMessagesClientsToSingleUser(toUserId, msgsClient)
	}

}

func (e _messageModelImple) SendAndStoreMsgsReceivedToPeer(msgs []Message) {
	groupByUsers := make(map[int][]Message)
	for _, msg := range msgs {
		//FromUserID,ok :=groupByUsers[msg.FromUserID]
		groupByUsers[msg.FromUserID] = append(groupByUsers[msg.FromUserID], msg)
	}

	var metaArr []MsgReceivedToPeer
	groupMetaUsers := make(map[int][]MsgReceivedToPeer)
	for toUser, msgs2 := range groupByUsers {
		for _, m := range msgs2 {
			met := MsgReceivedToPeer{
				ToUserId:   toUser,
				PeerUserId: m.ToUserId,
				RoomKey:    m.RoomKey,
				MsgKey:     m.MessageKey,
				AtTime:     helper.TimeNow(),
			}
			metaArr = append(metaArr, met)
			groupMetaUsers[toUser] = append(groupMetaUsers[toUser], met)
			//met.Insert(base.DB)
		}
	}
	err := MassInsert_MsgReceivedToPeer(metaArr, base.DB)
	helper.DebugPrintln(err)

	//// Send to each Message author users : "MsgsRevivedToPeerMany"
	for toUser, metas := range groupMetaUsers {
		helper.DebugPrintln(" calling clinet : MsgsRevivedToPeerMany :", toUser)
		succ := func() {
			var msgKeys []string
			for _, met := range metas {
				msgKeys = append(msgKeys, met.MsgKey)
			}
			NewMsgReceivedToPeer_Deleter().ToUserId_EQ(toUser).MsgKey_In(msgKeys).Delete(base.DB)
		}
		cal := base.NewCallWithData(CLIENT_CALL_MsgsReceivedToPeerMany, metas)
		AllPipesMap.SendToUserWithCallBack(toUser, cal, succ)
	}

}

//copy of SendAndStoreMsgsReceivedToPeer with changes
func (e _messageModelImple) SendAndStoreMsgsDeletedFromServer(msgs []Message) {
	groupByUsers := make(map[int][]Message)
	for _, msg := range msgs {
		//FromUserID,ok :=groupByUsers[msg.FromUserID]
		groupByUsers[msg.FromUserID] = append(groupByUsers[msg.FromUserID], msg)
	}

	var metaArr []MsgDeletedFromServer
	groupMetaUsers := make(map[int][]MsgDeletedFromServer)
	for toUser, msgs2 := range groupByUsers {
		for _, m := range msgs2 {
			met := MsgDeletedFromServer{
				ToUserId:   toUser,
				PeerUserId: m.ToUserId,
				RoomKey:    m.RoomKey,
				MsgKey:     m.MessageKey,
				AtTime:     helper.TimeNow(),
			}
			metaArr = append(metaArr, met)
			groupMetaUsers[toUser] = append(groupMetaUsers[toUser], met)
			//met.Insert(base.DB)
		}
	}
	err := MassInsert_MsgDeletedFromServer(metaArr, base.DB)
	helper.DebugPrintln(err)

	//// Send to each Message author users : "MsgsRevivedToPeerMany"
	for toUser, metas := range groupMetaUsers {
		helper.DebugPrintln(" calling clinet : MsgDeletedFromServer :", toUser)
		succ := func() {
			var msgKeys []string
			for _, met := range metas {
				msgKeys = append(msgKeys, met.MsgKey)
			}
			NewMsgDeletedFromServer_Deleter().ToUserId_EQ(toUser).MsgKey_In(msgKeys).Delete(base.DB)
		}
		cal := base.NewCallWithData(CLIENT_CALL_MsgsDeletedFromServerMany, metas)
		AllPipesMap.SendToUserWithCallBack(toUser, cal, succ)
	}

}

func (e _messageModelImple) SendManyMessagesRowsToSingleUser(ToUserId int, msgRows []Message) {
	helper.DebugPrintln("SendManyMessagesRowsToUser()")

	if len(msgRows) == 0 {
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
		MessageModel.SendAndStoreMsgsDeletedFromServer(msgRows)
	}

	dataSend := struct {
		Messages []*MessagesTableFromClient
		Users    []*UserViewSyncAndMe
	}{
		msgsRes, users,
	}

	call := base.NewCallWithData(CLIENT_CALL_MsgAddMany, dataSend)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
}

func (e _messageModelImple) SendManyMessagesClientsToSingleUser(ToUserId int, msgClients []MessagesTableFromClient) {
	helper.DebugPrintln("SendManyMessagesRowsToUser()")

	if len(msgClients) == 0 {
		return
	}

	mapOfSenders := make(map[int]bool, len(msgClients))
	for _, m := range msgClients {
		mapOfSenders[m.UserId] = true
	}

	users := make([]*UserViewSyncAndMe, 0, len(mapOfSenders))
	for id, _ := range mapOfSenders {
		users = append(users, Views.UserViewSync(ToUserId, id))
	}

	succ := func() {
		helper.DebugPrintln("SUCESS OF FlushAllStoredMessagesToUser()")
		arrMsgs := make([]string, 0, len(mapOfSenders))
		msgsRows := make([]Message, 0, len(msgClients))
		for _, m := range msgClients {
			arrMsgs = append(arrMsgs, m.MessageKey)
			mRow := Message{}
			mRow.FromClientMessageOptimized(ToUserId, m)
			msgsRows = append(msgsRows, mRow)

		}
		NewMessage_Deleter().ToUserId_EQ(ToUserId).MessageKey_In(arrMsgs).Delete(base.DB)
		MessageModel.SendAndStoreMsgsReceivedToPeer(msgsRows)
		MessageModel.SendAndStoreMsgsDeletedFromServer(msgsRows)
	}

	dataSend := struct {
		Messages []MessagesTableFromClient
		Users    []*UserViewSyncAndMe
	}{
		msgClients, users,
	}

	call := base.NewCallWithData(CLIENT_CALL_MsgAddMany, dataSend)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
}

//////////////////////////////////////////////////////////////////////
//////////////////////// Flushes methods /////////////////////////////
//////////////////////////////////////////////////////////////////////
func (e _messageModelImple) FlushAllStoredMessagesToUser(ToUserId int) {
	helper.DebugPrintln("FlushAllStoredMessagesToUser()")

	msgRows, err := NewMessage_Selector().ToUserId_EQ(ToUserId).OrderBy_Id_Asc().GetRows(base.DB) // first msgs rows first in slice
	if err != nil || len(msgRows) == 0 {
		return
	}
	MessageModel.SendManyMessagesRowsToSingleUser(ToUserId, msgRows)

	/*mapOfSenders := make(map[int]bool, len(msgRows))
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
		MessageModel.SendAndStoreMsgsDeletedFromServer(msgRows)
	}

	dataSend := struct {
		Messages []*MessagesTableFromClient
		Users    []*UserViewSyncAndMe
	}{
		msgsRes, users,
	}

	call := base.NewCallWithData(CLIENT_CALL_MsgAddMany, dataSend)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)*/
}

/////////////// Metas flush//////////////
func (e _messageModelImple) FlushAllReceivedMsgsToPeerToUser(ToUserId int) {
	helper.DebugPrintln("FlushAllReceivedMsgsToPeerToUser()")

	metasRows, err := NewMsgReceivedToPeer_Selector().ToUserId_EQ(ToUserId).OrderBy_Id_Asc().GetRows(base.DB) // first msgs rows first in slice
	if err != nil || len(metasRows) == 0 {
		return
	}
	last := metasRows[len(metasRows)-1].Id
	succ := func() {
		helper.DebugPrintln("SUCESS OF FlushAllReceivedMsgsToPeerToUser(): ", ToUserId)

		NewMsgReceivedToPeer_Deleter().ToUserId_EQ(ToUserId).Id_LE(last).Delete(base.DB)
	}

	call := base.NewCallWithData(CLIENT_CALL_MsgsReceivedToPeerMany, metasRows)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
}

func (e _messageModelImple) FlushAllDeletedMsgsToUser(ToUserId int) {
	helper.DebugPrintln("FlushAllDeletedMsgsToUser() ", ToUserId)

	metasRows, err := NewMsgDeletedFromServer_Selector().ToUserId_EQ(ToUserId).OrderBy_Id_Asc().GetRows(base.DB) // first msgs rows first in slice
	if err != nil || len(metasRows) == 0 {
		return
	}
	last := metasRows[len(metasRows)-1].Id
	succ := func() {
		helper.DebugPrintln("SUCESS OF FlushAllDeletedMsgsToUser(): ", ToUserId)

		NewMsgDeletedFromServer_Deleter().ToUserId_EQ(ToUserId).Id_LE(last).Delete(base.DB)
	}

	call := base.NewCallWithData(CLIENT_CALL_MsgsDeletedFromServerMany, metasRows)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
}

func (e _messageModelImple) FlushAllSeenMsgsByPeerToUser(ToUserId int) {
	helper.DebugPrintln("FlushAllSeenMsgsByPeerToUser() ", ToUserId)

	metasRows, err := NewMsgSeenByPeer_Selector().ToUserId_EQ(ToUserId).OrderBy_Id_Asc().GetRows(base.DB) // first msgs rows first in slice
	if err != nil || len(metasRows) == 0 {
		return
	}
	last := metasRows[len(metasRows)-1].Id
	succ := func() {
		helper.DebugPrintln("SUCESS OF FlushAllSeenMsgsByPeerToUser(): ", ToUserId)

		NewMsgSeenByPeer_Deleter().ToUserId_EQ(ToUserId).Id_LE(last).Delete(base.DB)
	}

	call := base.NewCallWithData(CLIENT_CALL_MsgsSeenByPeerMany, metasRows)

	AllPipesMap.SendToUserWithCallBack(ToUserId, call, succ)
}

//deprecated
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

///////////// Utils //////////////////

//format "p142_1569"
func UserIdsToRoomKey(UserId1, UserId2 int) string {
	if UserId2 < UserId1 {
		UserId1, UserId2 = UserId2, UserId1
	}
	return fmt.Sprintf("p%d_%d", UserId1, UserId2)
}
