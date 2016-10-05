package pipes

import (
    "fmt"
    "ms/sun/base"
    "encoding/json"
    "ms/sun/helper"
    "ms/sun/models"
    "ms/sun/constants"
)

func MsgAddOne(c base.Call) {
    fmt.Println("called MsgAddNew :", *c)

    myMsg := models.MessagesTableFromClient{}
    json.Unmarshal([]byte(c.Data), &myMsg)

    //toUid := helper.StrToInt(myMsg.RoomKey[1:], -1)
    toUid,err := RoomKeyToPeerUserId(myMsg.RoomKey, c.UserId)
    if err != nil {
        return
    }

    //msg: send to peer
    //msg := models.CreateNewMsgRecivedForSendingToPeer(&myMsg, toUid, c.UserId) //ref

    msg := myMsg

    msgTable := models.Message{
        MessageKey: msg.MessageKey,
        ToUserId: toUid,
        FromUserID: c.UserId,
        Data: helper.ToJson(msg),
        TimeMs: helper.TimeNowMs(),
    }

    call := base.NewCallWithData(constants.MsgAddOne, msg)
    AllPipesMap.SendToUser(toUid,call)

    msgTable.Insert(base.DB)

    //meta := models.CreateMsgRecivedToServerMetaResponse(&myMsg)                //ref

    //recivedCmd := commands.NewMsgsReceivedToServer(meta)

    //AllPipesMap.SendAndStoreCmdToUser(c.UserId, recivedCmd)

    //send msg to peer
    //cmd := commands.NewMsgsAddNew(msg)
    //AllPipesMap.SendAndStoreCmdToUser(toUid, cmd)
}

func MsgReceivedToPeer(c base.Call) {
    fmt.Printf("\ncalled MsgReceivedToPeer %v :\n", c)

    metas := []models.MessageSyncMeta{}

    json.Unmarshal([]byte(c.Data), &metas)

    for _, met := range metas {
        toUid,err := RoomKeyToPeerUserId(met.RoomKey, c.UserId)
        if err != nil {
            continue
        }
        metRes := models.MessageSyncMeta{
            MessageKey: met.MessageKey,
            RoomKey: met.RoomKey,
            ByUserId: c.UserId,
            AtTimeMs:met.AtTimeMs,// this is client time
        }

        data := []interface{} {metRes}

        recivedCall := base.NewCallWithData(constants.MsgsReceivedToPeer, data)
        AllPipesMap.SendToUser(toUid, recivedCall)

        //send MsgDeletedFromServer to that user
        delCmd := base.NewCallWithData(constants.MsgsDeletedFromServer, data)
        AllPipesMap.SendToUser(toUid, delCmd)
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


