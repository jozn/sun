package models

import (
	"github.com/golang/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

type sMsgPusher struct {
	messages []*x.Message
	toUserId int
}

func (p *sMsgPusher) pushToUser() {
	if len(p.messages) == 0 || !AllPipesMap.IsPipeOpen(p.toUserId) {
		return
	}

	pbMsgs := []*x.PB_Message{}
	userIds := make(map[int]bool)
	pbUsers := []*x.PB_UserWithMe{}

	for _, m := range p.messages {
		pbMsg := &x.PB_Message{}
		//err := proto.Unmarshal(m.DataPB, pbMsg)
		bts, err := helper.FromBase64ToBin(m.Data64)
		if err == nil {
			err := proto.Unmarshal(bts, pbMsg)
			if err == nil {
				pbMsgs = append(pbMsgs, pbMsg)
			}
		}

		userIds[m.UserId] = true
	}

	for uid, _ := range userIds {
		pbUsers = append(pbUsers, PBNew_PB_UserWithMe(uid, p.toUserId))
	}

	pushReq := &x.PB_PushMsgAddMany{
		Push:     nil,
		Messages: pbMsgs,
		Users:    pbUsers,
	}

	cmd := NewPB_CommandToClient_WithData(PB_PushMsgAddMany, pushReq)
	callback := func() {
		p.onPushedToUser()
	}

	AllPipesMap.SendToUserWithCallBack(p.toUserId, cmd, callback)

}

func (p *sMsgPusher) onPushedToUser() {
	p.addEventReceivedToPeer()
	p.removeFromMsgPush()
	p.deletedMessageFromServer()

}

func (p *sMsgPusher) removeFromMsgPush() {
	uids := make([]int, 0, len(p.messages))
	for _, msg := range p.messages {
		uids = append(uids, msg.Uid)
	}
	x.NewMsgPush_Deleter().ToUser_Eq(p.toUserId).Uid_In(uids).Delete(base.DB)
}

func (p *sMsgPusher) addEventReceivedToPeer() {
	p.addEventToMsgAuther(MESSAGE_PUSH_EVENT_RECEIVED_TO_PEER)
}

func (p *sMsgPusher) deletedMessageFromServer() {
	uids := make([]int, 0, len(p.messages))
	for _, msg := range p.messages {
		uids = append(uids, msg.Uid)
	}
	x.NewMessage_Deleter().Uid_In(uids).Delete(base.DB)
	p.addEventDeletedFromServer()
}

func (p *sMsgPusher) addEventDeletedFromServer() {
	p.addEventToMsgAuther(MESSAGE_PUSH_EVENT_DELETED_FROM_SERVER)
}

func (p *sMsgPusher) addEventToMsgAuther(EVENT_TYPE int) {
	for _, msg := range p.messages {
		me := x.MsgPushEvent{
			Id:         0,
			Uid:        helper.RandomSeqUid(),
			ToUserId:   msg.UserId,
			MsgUid:     msg.Uid,
			MsgKey:     msg.MessageKey,
			RoomKey:    msg.RoomKey,
			PeerUserId: p.toUserId,
			EventType:  EVENT_TYPE,
			AtTime:     helper.TimeNow(),
		}
		chanNewMsgPushEventsBuffer <- me
	}
}
