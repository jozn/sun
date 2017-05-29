package models

import (
	"github.com/golang/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"time"
)

var chanNewMsgPushEventsBuffer = make(chan *x.MsgPushEvent, 20000)

func messageModel_msgsRecicedToUserAddEvents(UserId int, messages []*x.Message) {
	for _, msg := range messages {
		me := &x.MsgPushEvent{
			Id:         0,
			Uid:        helper.RandomUid(),
			ToUserId:   msg.UserId,
			MsgKey:     msg.MessageKey,
			RoomKey:    msg.RoomKey,
			PeerUserId: UserId,
			EventType:  MESSAGE_PUSH_EVENT_RECEIVED_TO_PEER,
			AtTime:     helper.TimeNow(),
		}

		chanNewMsgPushEventsBuffer <- me
	}

}
