package models

import (
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/helper"
	"ms/sun/models/x"
)

type chatDirect struct {
	MeUserId   int
	PeerUserId int
	MeChat     *x.Chat
	PeerChat   *x.Chat
	Err        error
}

func NewDirectMessagingByUsers(meUserId, peerUserId int) *chatDirect {
	res := &chatDirect{
		MeUserId:   meUserId,
		PeerUserId: peerUserId,
	}
	res.LoadOrCreateRooms()
	return res
}

func NewDirectMessagingByChatId(meUserId, chatId int) (*chatDirect, error) {
	/* if meUserId <=  0 {
	       return  nil,errors.New("not user")
	   }
	   if chatId > 0 {
	       ch, err := x.Store.GetChatByChatId(chatId)
	       if err != nil {

	       }
	   }*/
	ch, err := x.NewChat_Selector().ChatId_Eq(chatId).UserId_Eq(meUserId).GetRow(base.DB)
	if err != nil {
		return nil, err
	}

	s := &chatDirect{
		MeUserId:   meUserId,
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

func (s *chatDirect) LoadOrCreateRooms() error {
	var e1, e2 error
	s.MeChat, e1 = GetOrCreateDirectChatForPeers(s.MeUserId, s.PeerUserId)
	s.PeerChat, e2 = GetOrCreateDirectChatForPeers(s.PeerUserId, s.MeUserId)
	if e1 != nil {
		s.Err = e1
	}

	if e2 != nil {
		s.Err = e2
	}

	if config.IS_DEBUG && s.Err != nil {
		logChat.Printf("%s - Err: %s - %v", "LoadOrCreateRooms() has error: ", s.Err, s)
	}
	return s.Err
}

///////////////// main fuctionalities ///////////////////
func (s *chatDirect) AddMessage(msg *x.DirectMessage) {
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
		Id: helper.NextRowsSeqId(),
		//ChatId:       s.MeChat.ChatId,
		MessageId: msg.MessageId,
		ChatKey:   UsersToChatKey(s.MeUserId, s.PeerUserId),
		//Seq:          s.MeChat.CurrentSeq,
		SourceEnumId: int(x.DirectMessageSourceEnum_COMPOSE_SOURCE),
	}

	d2mPeer := x.DirectToMessage{
		Id: helper.NextRowsSeqId(),
		//ChatId:       s.PeerChat.ChatId,
		ChatKey:   UsersToChatKey(s.PeerUserId, s.MeUserId),
		MessageId: msg.MessageId,
		//Seq:          s.PeerChat.CurrentSeq,
		SourceEnumId: int(x.DirectMessageSourceEnum_COMPOSE_SOURCE),
	}

	tx, err := base.DB.Begin()
	if err != nil {
		if config.IS_DEBUG {
			logChat.Printf(".AddMessage() transextion start has error: %s - %v", err, s)
		}
		return
	}

	s.PeerChat.Save(tx)
	s.MeChat.Save(tx)
	d2mMe.Insert(tx)
	d2mPeer.Insert(tx)
	err = tx.Commit()
	if err != nil {
		if config.IS_DEBUG {
			logChat.Printf(".AddMessage() transextion commit has error: %s - %v", err, s)
		}
		return
	}

	dlNew := x.DirectUpdate{
		Id:            helper.NextRowsSeqId(),
		ToUserId:      s.PeerChat.UserId,
		MessageId:     msg.MessageId,
		MessageFileId: msg.MessageFileId,
		//ChatId:        s.PeerChat.ChatId,
		ChatKey:       UsersToChatKey(s.MeUserId, s.PeerUserId),
		PeerUserId:    s.MeChat.UserId,
		RoomLogTypeId: int(Push_NEW_DIRECT_MESSAGE), //x.RoomLogTypeEnum_NEW_DIRECT_MESSAGE),
		FromSeq:       -1,
		ToSeq:         -1,
		ExtraPB:       []byte{},
		ExtraJson:     "",
		AtTimeMs:      helper.TimeNowMs(),
	}

	dlRec := x.DirectUpdate{
		Id:            helper.NextRowsSeqId(),
		ToUserId:      s.MeChat.UserId,
		MessageId:     msg.MessageId,
		MessageFileId: 0,
		ChatId:        s.MeChat.ChatId,
		PeerUserId:    s.PeerChat.UserId,
		RoomLogTypeId: int(Push_MESSAGE_RECIVED_TO_SERVER), //int(x.RoomLogTypeEnum_MESSAGE_RECIVED_TO_SERVER),
		FromSeq:       -1,
		ToSeq:         -1,
		ExtraPB:       []byte{},
		ExtraJson:     "",
		AtTimeMs:      helper.TimeNowMs(),
	}

	ChatUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: dlNew}
	ChatUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: dlRec}
}

func (s *chatDirect) DeleteMessageFromMe() {

}

func (s *chatDirect) EditMessageFromMe() {

}

//todo: make this more performant with MsgIds passing
func (s *chatDirect) SetMessagesAsSeen(fromSeq, toSeq, time int) {
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
		DeliviryStatusEnumId(int(x.RoomMessageDeliviryStatusEnum_SEEN2)).
		Update(base.DB)

	s.MeChat.LastSeqSeen = toSeq

	/*dlRec := x.DirectUpdate{
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

	  ChatUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: dlNew}*/
}

func (s *chatDirect) SetMessagesStatus() {

}

func (s *chatDirect) DeleteMyHistory() {

}
