package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"

	"github.com/golang/protobuf/proto"
)

type sMsgPusher struct {
	messages []*x.Message
	toUserId int
}

func (p *sMsgPusher) pushToUser() {
	if len(p.messages) == 0 || !AllPipesMap.IsPipeOpen(p.toUserId) {
		return
	}

	logPipes.Printf("sMsgPusher.pushToUser() uid:%d , msgs len:%d \n", p.toUserId, len(p.messages))

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
				logPipes.Println("befor file")
				if m.MsgFileId > 0 {
					fRow, ok := x.Store.GetMsgFileById(m.MsgFileId)
					logPipes.Printf("after file: %b %s", ok, helper.ToJson(fRow))
					if ok {
						filePb := PBConv_MsgFile_toNew_PB_MsgFile(fRow)
						pbMsg.File = &filePb
					}
				}
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
	logPipes.Printf("sMsgPusher.onPushedToUser() uid:%d , msgs len:%d \n", p.toUserId, len(p.messages))

	p.addEventReceivedToPeer()
	p.removeFromMsgPush()
	p.deletedMessageFromServer()

}

func (p *sMsgPusher) removeFromMsgPush() {
	uids := make([]int, 0, len(p.messages))
	for _, msg := range p.messages {
		uids = append(uids, msg.Uid)
	}
	x.NewMsgPush_Deleter().ToUser_Eq(p.toUserId).MsgUid_In(uids).Delete(base.DB)
}

func (p *sMsgPusher) addEventReceivedToPeer() {
	p.addEventToMsgAuther(MESSAGE_PUSH_EVENT_RECEIVED_TO_PEER)
}

func (p *sMsgPusher) deletedMessageFromServer() {
	uids := make([]int, 0, len(p.messages))
	idsFiles := make([]int, 0, len(p.messages))
	for _, msg := range p.messages {
		uids = append(uids, msg.Uid)
		if msg.MsgFileId > 0 {
			idsFiles = append(idsFiles, msg.MsgFileId)
		}
	}
	x.NewMessage_Deleter().Uid_In(uids).Delete(base.DB)
	x.NewMsgFile_Updater().Id_In(idsFiles).CanDel(helper.TimeNow()).Update(base.DB)
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
