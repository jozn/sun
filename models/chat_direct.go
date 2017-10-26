package models

import (
	"errors"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/helper"
	"ms/sun/models/x"
)

type _chatDirect struct {
	MeUserId    int
	PeerUserId  int
	MeChat      *x.Chat
	PeerChat    *x.Chat
	Err         error
	MeChatKey   string
	PeerChatKey string
}

func NewDirectMessagingByUsers(meUserId, peerUserId int) *_chatDirect {
	res := &_chatDirect{
		MeUserId:    meUserId,
		PeerUserId:  peerUserId,
		MeChatKey:   UsersToChatKey(meUserId, peerUserId),
		PeerChatKey: UsersToChatKey(peerUserId, meUserId),
	}
	res.LoadOrCreateRooms()
	return res
}

func NewDirectMessagingByChatId(meUserId int, chatkEY string) (*_chatDirect, error) {
	if meUserId <= 0 {
		return nil, errors.New("not user")
	}

	ch, err := x.NewChat_Selector().ChatKey_Eq(chatkEY).UserId_Eq(meUserId).GetRow(base.DB)
	if err != nil {
		return nil, err
	}

	s := NewDirectMessagingByUsers(meUserId, ch.PeerUserId)
	//s.MeChat = ch
	/*var e2 error
	s.PeerChat, e2 = GetOrCreateDirectChatForPeers(s.PeerUserId, s.MeUserId)
	if e2 != nil {
		s.Err = e2
	}*/

	//s.LoadOrCreateRooms()
	return s, nil
}

func (s *_chatDirect) LoadOrCreateRooms() error {
	if s.MeChat != nil && s.PeerChat != nil {
		return nil
	}
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
func (s *_chatDirect) AddMessage(msg *x.DirectMessage) {
	s.LoadOrCreateRooms()
	if s.Err != nil {
		return
	}

	err := msg.Save(base.DB)
	if err != nil {
		return
	}

	//Chat_IncermentForNewMessage(s.MeChat)
	//Chat_IncermentForNewMessage(s.PeerChat)

	Chat_setLastMsg(s.MeChat, msg)
	Chat_setLastMsg(s.PeerChat, msg)

	d2mMe := x.DirectToMessage{
		Id:           helper.NextRowsSeqId(),
		MessageId:    msg.MessageId,
		ChatKey:      s.MeChat.ChatKey,
		SourceEnumId: int(x.DirectMessageSourceEnum_COMPOSE_SOURCE),
	}

	d2mPeer := x.DirectToMessage{
		Id:           helper.NextRowsSeqId(),
		ChatKey:      s.PeerChat.ChatKey,
		MessageId:    msg.MessageId,
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
		DirectUpdateId: helper.NextRowsSeqId(),
		ToUserId:       s.PeerChat.UserId,
		MessageId:      msg.MessageId,
		MessageFileId:  msg.MessageFileId,
		//ChatId:        s.PeerChat.ChatId,
		ChatKey:       s.PeerChatKey, //UsersToChatKey(s.MeUserId, s.PeerUserId),
		PeerUserId:    s.MeChat.UserId,
		RoomLogTypeId: int(Push_NEW_DIRECT_MESSAGE), //x.RoomLogTypeEnum_NEW_DIRECT_MESSAGE),
		FromSeq:       -1,
		ToSeq:         -1,
		ExtraPB:       []byte{},
		ExtraJson:     "",
		AtTimeMs:      helper.TimeNowMs(),
	}

	dlRec := x.DirectUpdate{
		DirectUpdateId: helper.NextRowsSeqId(),
		ToUserId:       s.MeChat.UserId,
		MessageId:      msg.MessageId,
		MessageFileId:  0,
		ChatKey:        s.MeChatKey, //UsersToChatKey(s.MeUserId, s.PeerUserId),
		PeerUserId:     s.PeerChat.UserId,
		RoomLogTypeId:  int(Push_MESSAGE_RECIVED_TO_SERVER), //int(x.RoomLogTypeEnum_MESSAGE_RECIVED_TO_SERVER),
		FromSeq:        -1,
		ToSeq:          -1,
		ExtraPB:        []byte{},
		ExtraJson:      "",
		AtTimeMs:       helper.TimeNowMs(),
	}

	LiveUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: dlNew}
	LiveUpdateFramer.HereDirectDelayer <- UpdateDelayer{directUpdate: dlRec}
}

func (s *_chatDirect) DeleteMessageFromMe() {

}

func (s *_chatDirect) EditMessageFromMe() {

}

//todo: make this more performant with MsgIds passing
func (s *_chatDirect) SetMessagesAsSeen(fromSeq, toSeq, time int) {
	sel := x.NewDirectToMessage_Selector().Select_MessageId().
		ChatKey_Eq(s.MeChatKey)
	/*if fromSeq > 0 {
		sel.Seq_GE(fromSeq)
	}
	if toSeq > 0 {
		sel.Seq_GE(toSeq)
	}*/

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

	  LiveUpdateFramer.HereDirectDelayer_DEP <- UpdateDelayer{directUpdate: dlNew}*/
}

func (s *_chatDirect) SetMessagesStatus() {

}

func (s *_chatDirect) DeleteMyHistory() {

}
