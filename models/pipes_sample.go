package models

import (
	"ms/sun/helper"
	"ms/sun/models/x"
)

func SampleSendMessage(toUserId int, msgPb *x.PB_Message) {
	msgBuf := newChatMsgDelayer{
		msgPB:      msgPb,
		fromUserId: int(msgPb.UserId),
		toUserId:   toUserId,
		roomKey:    msgPb.RoomKey,
		hashId:     1,
		uid:        helper.RandomSeqUid(),
	}
	chanNewChatMsgsBuffer <- msgBuf
}
