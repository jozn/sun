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

func NewDirectMessagingByUsers(me, peer int) *directMessaging {
	res := &directMessaging{
		MeUserId:   me,
		PeerUserId: peer,
	}
	res.LoadOrCreateRooms()
	return res
}

func NewDirectMessagingByChatId(me, chatId int) (*directMessaging, error) {
	/* if me <=  0 {
	       return  nil,errors.New("not user")
	   }
	   if chatId > 0 {
	       ch, err := x.Store.GetChatByChatId(chatId)
	       if err != nil {

	       }
	   }*/
	ch, err := x.NewChat_Selector().ChatId_Eq(chatId).UserId_Eq(me).GetRow(base.DB)
	if err != nil {
		return nil, err
	}

	s := &directMessaging{
		MeUserId:   me,
		PeerUserId: ch.PeerUserId,
	}
	s.MeChat = ch
	var e2 error
	s.PeerChat, e2 = GetOrCreateDirectChatForPeers(s.PeerUserId, s.MeUserId)
	if e2 != nil {
		s.Err = e2
	}

	s.LoadOrCreateRooms()
	return s, nil
}

func (s *directMessaging) LoadOrCreateRooms() error {
	var e1, e2 error
	s.MeChat, e1 = GetOrCreateDirectChatForPeers(s.MeUserId, s.PeerUserId)
	s.PeerChat, e2 = GetOrCreateDirectChatForPeers(s.PeerUserId, s.MeUserId)
	if e1 != nil {
		s.Err = e1
	}

	if e2 != nil {
		s.Err = e2
	}
	return s.Err
}

///////////////// main fuctionalities ///////////////////
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
		SourceEnum: int(x.DirectMessageSourceEnum_COMPOSE_SOURCE),
	}

	d2mPeer := x.DirectToMessage{
		Id:         helper.NextRowsSeqId(),
		ChatId:     s.PeerChat.ChatId,
		MessageId:  msg.MessageId,
		Seq:        s.PeerChat.CurrentSeq,
		SourceEnum: int(x.DirectMessageSourceEnum_COMPOSE_SOURCE),
	}

	tx, err := base.DB.Begin()
	if err != nil {
		return
	}

	s.PeerChat.Save(tx)
	s.MeChat.Save(tx)
	d2mMe.Insert(tx)
	d2mPeer.Insert(tx)
	err = tx.Commit()
	if err != nil {
		return
	}

	dlNew := x.DirectLog{
		Id:            helper.NextRowsSeqId(),
		ToUserId:      s.PeerChat.UserId,
		MessageId:     msg.MessageId,
		ChatId:        s.PeerChat.ChatId,
		PeerUserId:    s.MeChat.UserId,
		RoomLogTypeId: int(x.RoomLogTypeEnum_NEW_DIRECT_MESSAGE),
		FromSeq:       -1,
		ToSeq:         -1,
		ExtraPB:       []byte{},
		ExtraJson:     "",
		AtTimeMs:      helper.TimeNowMs(),
	}

	dlRec := x.DirectLog{
		Id:            helper.NextRowsSeqId(),
		ToUserId:      s.MeChat.UserId,
		MessageId:     msg.MessageId,
		ChatId:        s.MeChat.ChatId,
		PeerUserId:    s.PeerChat.UserId,
		RoomLogTypeId: int(x.RoomLogTypeEnum_MESSAGE_RECIVED_TO_SERVER),
		FromSeq:       -1,
		ToSeq:         -1,
		ExtraPB:       []byte{},
		ExtraJson:     "",
		AtTimeMs:      helper.TimeNowMs(),
	}

	LogUpdater.HereDirectDelayer <- logDelayer{directLog: dlNew}
	LogUpdater.HereDirectDelayer <- logDelayer{directLog: dlRec}
}

func (s *directMessaging) DeleteMessageFromMe() {

}

func (s *directMessaging) EditMessageFromMe() {

}

//todo: make this more performant with MsgIds passing
func (s *directMessaging) SetMessagesAsSeen(fromSeq, toSeq, time int) {
	sel := x.NewDirectToMessage_Selector().Select_MessageId().
		ChatId_Eq(s.MeChat.ChatId)
	if fromSeq > 0 {
		sel.Seq_GE(fromSeq)
	}
	if toSeq > 0 {
		sel.Seq_GE(toSeq)
	}

	msgIds, err := sel.OrderBy_Id_Desc().GetIntSlice(base.DB)
	if err != nil {
		return
	}

	x.NewDirectMessage_Updater().
		MessageId_In(msgIds).
		DeliviryStatusEnum(int(x.RoomMessageDeliviryStatusEnum_SEEN2)).
		Update(base.DB)

	s.MeChat.LastSeqSeen = toSeq

	/*dlRec := x.DirectLog{
	      Id:            helper.NextRowsSeqId(),
	      ToUserId:      s.MeChat.UserId,
	      MessageId:     msg.MessageId,
	      ChatId:        s.MeChat.ChatId,
	      PeerUserId:    s.PeerChat.UserId,
	      RoomLogTypeId: int(x.RoomLogTypeEnum_MESSAGE_RECIVED_TO_SERVER),
	      FromSeq:       -1,
	      ToSeq:         -1,
	      ExtraPB:       []byte{},
	      ExtraJson:     "",
	      AtTimeMs:      helper.TimeNowMs(),
	  }

	  LogUpdater.HereDirectDelayer <- logDelayer{directLog: dlNew}*/
}

func (s *directMessaging) SetMessagesStatus() {

}

func (s *directMessaging) DeleteMyHistory() {

}
