package pipes

import (
	"ms/sun/base"
	//"time"
	"encoding/json"
	"fmt"
	"ms/sun/helper"
	"ms/sun/models"
	//"ms/sun/docs/del/chat"
	"ms/sun/commands"
	"ms/sun/constants"
	"ms/sun/sync"
)

func MsgAddNew(c *base.CmdAction) {
	fmt.Println("called MsgAddNew :", *c)

	myMsg := models.MessagesTableFromClient{}
	json.Unmarshal([]byte(c.Cmd.Data), &myMsg)

	toUid := helper.StrToInt(myMsg.RoomKey[1:], -1)
	if toUid < 1 {
		return
	}

	//msg: send to peer
	msg := models.CreateNewMsgRecivedForSendingToPeer(&myMsg, toUid, c.UserId) //ref
	meta := models.CreateMsgRecivedToServerMetaResponse(&myMsg)                //ref

	recivedCmd := commands.NewMsgsReceivedToServer(meta)

	sync.AllPipesMap.SendAndStoreCmdToUser(c.UserId, recivedCmd)

	//send msg to peer
	cmd := commands.NewMsgsAddNew(msg)
	sync.AllPipesMap.SendAndStoreCmdToUser(toUid, cmd)
}

func MsgReceivedToPeer(c *base.CmdAction) {
	fmt.Printf("\ncalled MsgReceivedToPeer %v :\n", c)

	metas := []models.MessageSyncMeta{}

	json.Unmarshal([]byte(c.Cmd.Data), &metas)

	for _, met := range metas {
		toUid := helper.StrToInt(met.RoomKey[1:], -1)
		if toUid < 1 {
			continue
		}
		metRes := models.MessageSyncMeta{}
		metRes.RoomKey = "u" + helper.IntToStr(c.UserId)
		metRes.ByUserId = c.UserId
		metRes.MessageKey = met.MessageKey
		metRes.AtTimeMs = met.AtTimeMs

		recivedCmd := base.NewCommand(constants.MsgsReceivedToPeer)
		recivedCmd.AddSliceData(metRes)
		recivedCmd.MakeDataReady()

		sync.AllPipesMap.SendAndStoreCmdToUser(toUid, recivedCmd)

		//send MsgDeletedFromServer to that user
		delCmd := base.NewCommand(constants.MsgsDeletedFromServer)
		delCmd.AddSliceData(metRes)
		delCmd.MakeDataReady()

		sync.AllPipesMap.SendAndStoreCmdToUser(toUid, delCmd)
	}

}

func MsgSeenByPeer(c *base.CmdAction) {
	fmt.Println("called MsgSeenByPeer :", *c)

	metas := []models.MessageSyncMeta{}

	json.Unmarshal([]byte(c.Cmd.Data), &metas)

	for _, met := range metas {
		toUid := helper.StrToInt(met.RoomKey[1:], -1)
		if toUid < 1 {
			continue
		}
		metRes := models.MessageSyncMeta{}
		metRes.RoomKey = "u" + helper.IntToStr(c.UserId)
		metRes.ByUserId = c.UserId
		metRes.MessageKey = met.MessageKey
		metRes.AtTimeMs = met.AtTimeMs
		metRes.ExtraData = met.ExtraData

		recivedCmd := base.NewCommand(constants.MsgsSeenByPeer)
		recivedCmd.AddSliceData(metRes)
		recivedCmd.MakeDataReady()

		sync.AllPipesMap.SendAndStoreCmdToUser(toUid, recivedCmd)
	}

}

///NOT NEEDED IN SERVER
func MsgDeletedFromServer(c *base.CmdAction) {

}

func MsgReceivedToServer(c *base.CmdAction) {

}

//func MsgAddNew(c *base.CmdAction) {
//    fmt.Println("called MsgAddNew :",*c)
//
//    m := models.MessagesTableFromClient{}
//    json.Unmarshal([]byte(c.Cmd.Data),&m)
//
//    toUid := helper.StrToInt(m.RoomKey[1:],-1)
//    if toUid < 1{
//        return
//    }
//
//    //msg: send to peer
//    msg := models.CreateNewMsgRecivedForSendingToPeer(&m,toUid,c.UserId) //copy
//    //msg.UserId = c.UserId
//    //msg.RoomKey = "u"+helper.IntToStr(c.UserId)
//
//    //send MsgReceivedToServer to this user
//    //meta := models.MessageSyncMeta{}
//    //meta.RoomKey = m.RoomKey
//    //meta.ByUserId = -1
//    //meta.MessageKey = m.MessageKey
//    //meta.AtTimeMs = helper.TimeNowMs()
//    meta := models.CreateMsgRecivedToServerMetaResponse(&msg)
//
//    //recivedCmd := base.NewCommand(constants.MsgsReceivedToServer)
//    //recivedCmd.AddSliceData(meta)
//    //recivedCmd.MakeDataReady()
//    recivedCmd := commands.NewMsgsReceivedToServer(&meta)
//
//    sync.AllPipesMap.SendAndStoreCmdToUser(c.UserId,recivedCmd)
//
//    //send msg to peer
//    cmd := commands.NewMsgsAddNew(&msg)
//    //cmd := base.NewCommand(constants.MsgsAddNew)
//    ////cmd.SetData(msg)
//    //cmd.AddSliceData(msg)
//    //cmd.MakeDataReady()
//    sync.AllPipesMap.SendAndStoreCmdToUser(toUid,cmd)
//}
