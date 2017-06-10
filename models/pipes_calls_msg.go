package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
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
		toUid, err := RoomKeyToPeerUserId(msgPb.RoomKey, pipe.UserId)
		if err == nil && toUid > 0 {
			msgBuf := newChatMsgDelayer{
				msgPB:      msgPb,
				fromUserId: pipe.UserId,
				toUserId:   toUid,
				roomKey:    msgPb.RoomKey,
				hashId:     1,
				uid:        helper.RandomSeqUid(),
			}
			chanNewChatMsgsBuffer <- msgBuf
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

	//var events []x.MsgPushEvent
	for _, seenPb := range req.Seen {
		eventRow := PBConv_PB_MsgSeen_toNew_MsgPushEvent(seenPb)
		eventRow.PeerUserId = pipe.UserId
		eventRow.ToUserId, _ = RoomKeyToPeerUserId(seenPb.RoomKey, pipe.UserId)
		chanNewMsgPushEventsBuffer <- eventRow
		//events = append(events, eventRow)
	}

	//x.MassInsert_MsgPushEvent(events, base.DB)

	/*mpGroupByuser := make(map[int][]x.MsgPushEvent)
	for _, seen := range events {
		mpGroupByuser[seen.ToUserId] = append(mpGroupByuser[seen.ToUserId], seen)
	}*/

	//PushProxy_PushMsgsEvents(events)
}

func EchoCmd(c *x.PB_CommandToServer, pipe *UserDevicePipe) {
	//b, _ := json.Marshal(c)
	//r := base.WSRes{Status: "BB", ReqKey: string(b)}
	//call := base.NewCallWithData("echo", "sad")
	call := NewPB_CommandToClient("echo")
	AllPipesMap.SendToUser(pipe.UserId, call)
	//AllPipesMap.SendToUser_DEP(c.UserId, r)
}
