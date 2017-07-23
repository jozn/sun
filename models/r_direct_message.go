package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

type directMessaging struct {
	MeUserId   int
	PeerUserId int
	MeChat     *x.Chat
	PeerChat   *x.Chat
	Err        error
}

func NewDirectMessaging(me, peer int) *directMessaging {
	return &directMessaging{
		MeUserId:   me,
		PeerUserId: peer,
	}
}

func (s *directMessaging) LoadOrCreateRooms() {
	var e1, e2 error
	s.MeChat, e1 = GetOrCreateDirectChatForPeers(s.MeUserId, s.PeerUserId)
	s.PeerChat, e2 = GetOrCreateDirectChatForPeers(s.PeerUserId, s.MeUserId)
	if e1 != nil {
		s.Err = e1
	}

	if e2 != nil {
		s.Err = e2
	}
}

func (s *directMessaging) AddMessage(msg *x.DirectMessage) {
	s.LoadOrCreateRooms()
	if s.Err != nil {
		return
	}

	err := msg.Save(base.DB)
	if err != nil {
		return
	}

	Chat_IncermentForNewMessage(s.MeChat)
	Chat_IncermentForNewMessage(s.PeerChat)

	d2mMe := x.DirectToMessage{
		Id:         helper.NextRowsSeqId(),
		ChatId:     s.MeChat.ChatId,
		MessageId:  msg.MessageId,
		Seq:        s.MeChat.CurrentSeq,
		SourceEnum: 0,
	}

	d2mPeer := x.DirectToMessage{
		Id:         helper.NextRowsSeqId(),
		ChatId:     s.PeerChat.ChatId,
		MessageId:  msg.MessageId,
		Seq:        s.PeerChat.CurrentSeq,
		SourceEnum: 0,
	}

	tx, err := base.DB.Begin()
	if err != nil {
		return
	}

	s.PeerChat.Save(tx)
	s.PeerChat.Save(tx)
	d2mMe.Insert(tx)
	d2mPeer.Insert(tx)
	tx.Commit()

	//todo add to chanels logs
}

func (s *directMessaging) DeleteMessageFromMe() {

}

func (s *directMessaging) EditMessageFromMe() {

}

func (s *directMessaging) SetMessagesAsSeen() {

}

func (s *directMessaging) SetMessagesStatus() {

}

func (s *directMessaging) DeleteMyHistory() {

}
