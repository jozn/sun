package models

import (
    "ms/sun/models/x"
    "ms/sun/helper"
    "github.com/golang/protobuf/proto"
)

//just return the by value

func PBConv_PB_Message_toNew_Message(pb *x.PB_Message) x.Message {
    msg := &x.Message{
        Id: 0,
        UserId: pb.UserId,
        MessageKey: pb.MessageKey,
        RoomKey: pb.RoomKey,
        MessageType: pb.MessageTypeId,
        RoomType: pb.RoomTypeId,
        DataPB: nil ,
        DataJson: helper.ToJson(pb),
        CreatedTimeMs: helper.TimeNowMs(),
    }
    bs,err :=proto.Marshal(pb)
    if err == nil{
        msg.DataPB =bs
    }

    return msg
}

func PBConv_PB_MsgSeen_toNew_MsgPushEvent(pb *x.PB_MsgSeen) x.MsgPushEvent {
    msg := &x.MsgPushEvent{
        Id: 0,
        ToUserId: -1 , //must set
        MsgKey: pb.MessageKey,
        RoomKey: pb.RoomKey,
        PeerUserId: -1,//must set
        EventType: MESSAGE_PUSH_EVENT_SEEN_BY_PEER,
        AtTime: pb.AtTime,
    }

    return msg
}
