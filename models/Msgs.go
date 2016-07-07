package models

import "ms/sun/helper"

func CreateMsgRecivedToServerMetaResponse(msg *MessagesTableFromClient) *MessageSyncMeta {
    meta := MessageSyncMeta{}
    meta.RoomKey = msg.RoomKey
    meta.ByUserId = -1
    meta.MessageKey = msg.MessageKey
    meta.AtTimeMs = helper.TimeNowMs()
    return &meta
}

func CreateNewMsgRecivedForSendingToPeer(msg *MessagesTableFromClient, toPeerUserId, fromUserId int) *MessagesTableFromClient {
    m := *msg //copy
    m.UserId = fromUserId
    m.RoomKey = "u"+helper.IntToStr(fromUserId)
    return &m
}

func RoomKeyToUserId(roomKey string ) int  {
   return helper.StrToInt(roomKey[1:],-1)
}
