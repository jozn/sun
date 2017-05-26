package models

import (
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
        return
	}

    for _, msgPb := range req.Messages {
        msgRow := PBConv_PB_Message_toNew_Message(msgPb)
        msgRow.Save(base.DB)
        toUid,err := RoomKeyToPeerUserId(msgPb.RoomKey,pipe.UserId)
        if err == nil{
            PushProxy_PushMsgsReceivedToPeer(toUid,msgRow, msgPb)
        }
    }
}

//Todo : add security layer
func CallRecive_MsgSeenByPeer(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
	fmt.Println("called MsgSeenByPeer :", c)

    req := &x.PB_RequestMsgsSeen{}
    err := proto.Unmarshal(c.GetData(), req)
    if err != nil {
        helper.DebugPrintln("Error :", err)
        return
    }

    var events []x.MsgPushEvent
    for _, seenPb := range req.Seen {
        evetRow := PBConv_PB_MsgSeen_toNew_MsgPushEvent(seenPb)
        evetRow.PeerUserId = pipe.UserId
        evetRow.ToUserId , _ = RoomKeyToPeerUserId(seenPb.RoomKey,pipe.UserId)
        events = append(events, evetRow)
    }

    x.MassInsert_MsgPushEvent(events,base.DB)

    mpGroupByuser := make(map[int][]x.MsgPushEvent)
    for _, seen := range events {
        mpGroupByuser[seen.ToUserId] = append(mpGroupByuser[seen.ToUserId], seen)
    }

    PushProxy_PushMsgsEvents(events)
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
