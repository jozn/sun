package models

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

func CallReceive_MsgsAddMany(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
	helper.DebugPrintln("called CallReceive_MsgsAddMany() :")

	req := &x.PB_RequestMsgAddMany{}
	err := proto.Unmarshal(c.GetData(), req)
	if err != nil {
		helper.DebugPrintln("Error :", err)
	}

	MessageModel.SendAndStoreManyMessages(c.UserId, myMsgs)
}

//Todo : add security layer
func CallRecive_MsgSeenByPeer(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
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

func EchoCmd(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
	//b, _ := json.Marshal(c)
	//r := base.WSRes{Status: "BB", ReqKey: string(b)}
	call := base.NewCallWithData("echo", "sad")
	AllPipesMap.SendToUser(c.UserId, call)
	//AllPipesMap.SendToUser_DEP(c.UserId, r)
}

/*
func CallRecive_MsgReceivedToPeer(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
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


*/

/*
//dep
func CallReceive_MsgsAddOne(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
	helper.DebugPrintln("called CallReceive_MsgsAddOne() :")

	msg := &x.PB_Message{}
	err := proto.Unmarshal(c.GetData(), msg)
	if err != nil {
		helper.DebugPrintln("Error :", err)
	}

	myMsg := MessagesTableFromClient{}
	json.Unmarshal([]byte(c.Data), &myMsg)

	//toUid := helper.StrToInt(myMsg.RoomKey[1:], -1)
	toUid, err := RoomKeyToPeerUserId(myMsg.RoomKey, c.UserId)
	if err != nil {
		helper.DebugPrintln(err)
		return
	}

	MessageModel.SendAndStoreMessage(toUid, myMsg)
}
*/
