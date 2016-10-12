package models

import (
	"encoding/json"
	"fmt"
	"ms/sun/base"
	"ms/sun/constants"
	"ms/sun/helper"
)

func CallReceive_MsgsAddOne(c base.Call) {
	helper.DebugPrintln("called CallReceive_MsgsAddOne() :")


	myMsg := MessagesTableFromClient{}
	json.Unmarshal([]byte(c.Data), &myMsg)

	//toUid := helper.StrToInt(myMsg.RoomKey[1:], -1)
	toUid, err := RoomKeyToPeerUserId(myMsg.RoomKey, c.UserId)
	if err != nil {
        helper.DebugPrintln(err)
		return
	}

    MessageModel.SendAndStoreMessage(toUid,myMsg)

    /*msg := myMsg

    msgTable := Message{
        MessageKey: msg.MessageKey,
        ToUserId:   toUid,
        FromUserID: c.UserId,
        Data:       helper.ToJson(msg),
        TimeMs:     helper.TimeNowMs(),
    }

    call := base.NewCallWithData(constants.MsgAddOne, msg)
    AllPipesMap.SendToUser(toUid, call)

    msgTable.Insert(base.DB)*/

	//meta := CreateMsgRecivedToServerMetaResponse(&myMsg)                //ref

	//recivedCmd := commands.NewMsgsReceivedToServer(meta)

	//AllPipesMap.SendAndStoreCmdToUser(c.UserId, recivedCmd)

	//send msg to peer
	//cmd := commands.NewMsgsAddNew(msg)
	//AllPipesMap.SendAndStoreCmdToUser(toUid, cmd)
}

func CallRecive_MsgReceivedToPeer(c base.Call) {
	fmt.Printf("\ncalled MsgReceivedToPeer %v :\n", c)

	metas := []MessageSyncMeta{}

	json.Unmarshal([]byte(c.Data), &metas)

	for _, met := range metas {
		toUid, err := RoomKeyToPeerUserId(met.RoomKey, c.UserId)
		if err != nil {
			continue
		}
		metRes := MessageSyncMeta{
			MessageKey: met.MessageKey,
			RoomKey:    met.RoomKey,
			ByUserId:   c.UserId,
			AtTimeMs:   met.AtTimeMs, // this is client time
			ExtraData:  met.ExtraData,
		}

		data := []interface{}{metRes}

		//m := MsgReceivedToPeer{}

		recivedCall := base.NewCallWithData(constants.MsgsReceivedToPeer, data)
		AllPipesMap.SendToUser(toUid, recivedCall)

		//send MsgDeletedFromServer to that user
		delCmd := base.NewCallWithData(constants.MsgsDeletedFromServer, data)
		AllPipesMap.SendToUser(toUid, delCmd)
	}

}

func CallRecive_MsgSeenByPeer(c *base.CmdAction) {
	fmt.Println("called MsgSeenByPeer :", *c)

	metas := []MessageSyncMeta{}

	json.Unmarshal([]byte(c.Cmd.Data), &metas)

	for _, met := range metas {
		toUid, err := RoomKeyToPeerUserId(met.RoomKey, c.UserId)
		if err != nil {
			continue
		}
		metRes := MessageSyncMeta{
			MessageKey: met.MessageKey,
			RoomKey:    met.RoomKey,
			ByUserId:   c.UserId,
			AtTimeMs:   met.AtTimeMs, // this is client time
			ExtraData:  met.ExtraData,
		}

		data := []interface{}{metRes}

		recivedCmd := base.NewCallWithData(constants.MsgsSeenByPeer, data)

		AllPipesMap.SendToUser(toUid, recivedCmd)
	}
}

///NOT NEEDED IN SERVER
func CallRecive_MsgDeletedFromServer(c *base.CmdAction) {
	//needs for group
}

func CallRecive_MsgReceivedToServer(c *base.CmdAction) {
	//maybe some futuer work to be more reliable
}
