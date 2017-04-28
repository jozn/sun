package models

import (
	"encoding/json"
	"fmt"
	"ms/sun/base"
	"ms/sun/constants"
	"ms/sun/helper"
	"ms/sun/models/x"
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

	MessageModel.SendAndStoreMessage(toUid, myMsg)

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

func CallReceive_MsgsAddMany(c base.Call) {
	helper.DebugPrintln("called CallReceive_MsgsAddMany() :")

	var myMsgs []MessagesTableFromClient
	err := json.Unmarshal([]byte(c.Data), &myMsgs)
	if err != nil {
		return
	}

	MessageModel.SendAndStoreManyMessages(c.UserId, myMsgs)
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

//Todo : add security layer
func CallRecive_MsgSeenByPeer(c base.Call) {
	fmt.Println("called MsgSeenByPeer :", c)

	seensRows := []x.MsgSeenByPeer{}

	json.Unmarshal([]byte(c.Data), &seensRows)

	mpGroupByuser := make(map[int][]x.MsgSeenByPeer)
	for _, seen := range seensRows {
		mpGroupByuser[seen.ToUserId] = append(mpGroupByuser[seen.ToUserId], seen)
	}

	if len(mpGroupByuser) == 0 {
		return
	}

	if len(mpGroupByuser) == 1 {
		var touser int
		var seens []x.MsgSeenByPeer

		for touser, seens = range mpGroupByuser {
		}

		err := func() {
			//fmt.Println("**********************\n*********************\n********************",touser)

			x.MassInsert_MsgSeenByPeer(seens, base.DB)
		}

		call := base.NewCallWithData(CLIENT_CALL_MsgsSeenByPeerMany, seens)

		AllPipesMap.SendToUserWithCallBacks(touser, call, nil, err)
		return
	}

	x.MassInsert_MsgSeenByPeer(seensRows, base.DB)
	for toUserId, seens := range mpGroupByuser {
		MessageModel.SendListOfSeenMsgsByPeerToUser(toUserId, seens)
	}

}

func EchoCmd(c base.Call) {
	//b, _ := json.Marshal(c)
	//r := base.WSRes{Status: "BB", ReqKey: string(b)}
	call := base.NewCallWithData("echo", "sad")
	AllPipesMap.SendToUser(c.UserId, call)
	//AllPipesMap.SendToUser_DEP(c.UserId, r)
}
