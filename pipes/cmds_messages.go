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

	AllPipesMap.SendAndStoreCmdToUser(c.UserId, recivedCmd)

	//send msg to peer
	cmd := commands.NewMsgsAddNew(msg)
	AllPipesMap.SendAndStoreCmdToUser(toUid, cmd)
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

		AllPipesMap.SendAndStoreCmdToUser(toUid, recivedCmd)

		//send MsgDeletedFromServer to that user
		delCmd := base.NewCommand(constants.MsgsDeletedFromServer)
		delCmd.AddSliceData(metRes)
		delCmd.MakeDataReady()

		AllPipesMap.SendAndStoreCmdToUser(toUid, delCmd)
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

		AllPipesMap.SendAndStoreCmdToUser(toUid, recivedCmd)
	}

}

///NOT NEEDED IN SERVER
func MsgDeletedFromServer(c *base.CmdAction) {

}

func MsgReceivedToServer(c *base.CmdAction) {

}

