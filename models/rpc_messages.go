package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

type rpcMsg int

func (rpcMsg) AddNewTextMessage(i *x.PB_MsgParam_AddNewTextMessage, p x.RPC_UserParam) (*x.PB_MsgResponse_AddNewTextMessage, error) {
	pid := int(i.PeerId)
	msg := &x.DirectMessage{
		MessageId:          0,
		RoomKey:            UsersToRoomKey(p.GetUserId(), int(i.GetPeerId())),
		UserId:             p.GetUserId(),
		MessageFileId:      0,
		MessageTypeEnum:    int(x.RoomMessageTypeEnum_TEXT),
		Text:               i.Text,
		Time:               int(i.Time),
		PeerReceivedTime:   0,
		PeerSeenTime:       0,
		DeliviryStatusEnum: int(x.RoomMessageDeliviryStatusEnum_SENT),
	}

	//msg.Insert(base.DB)

	dm := NewDirectMessagingByUsers(p.GetUserId(), pid)
	dm.AddMessage(msg)

}

func (rpcMsg) SetRoomActionDoing(i *x.PB_MsgParam_SetRoomActionDoing, p x.RPC_UserParam) (*x.PB_MsgResponse_SetRoomActionDoing, error) {
	panic("implement me")
}

func (rpcMsg) GetMessagesByIds(i *x.PB_MsgParam_GetMessagesByIds, p x.RPC_UserParam) (*x.PB_MsgResponse_GetMessagesByIds, error) {
	ids := helper.SliceInt64ToInt(i.MessageId)
	x.Store.PreLoadDirectMessageByMessageIds(ids)

	views := make([]*x.PB_MessageView, len(i.MessageId))
	for i, id := range i.MessageId {
		views[i] = ChatViews.MessageIdToMessageView(int(id))
	}

	res := &x.PB_MsgResponse_GetMessagesByIds{
		views,
	}
	return res, nil
}

func (rpcMsg) GetMessagesHistory(i *x.PB_MsgParam_GetMessagesHistory, p x.RPC_UserParam) (*x.PB_MsgResponse_GetMessagesHistory, error) {
	sel := x.NewDirectToMessage_Selector().Select_MessageId().ChatId_Eq(int(i.ChatId))
	if i.FromSeq > 0 {
		sel.Seq_GE(int(i.FromSeq))
	}
	if i.FromSeq > 0 {
		sel.Seq_GE(int(i.FromSeq))
	}

	msgIds, err := sel.OrderBy_Id_Desc().GetIntSlice(base.DB)

	if err != nil {
		return nil, err
	}

	x.Store.PreLoadDirectMessageByMessageIds(msgIds)

	res := &x.PB_MsgResponse_GetMessagesHistory{
		ChatViews.MessageIdsToMessageViews(msgIds),
	}
	return res, nil

}

func (rpcMsg) SetMessagesRangeAsSeen(i *x.PB_MsgParam_SetMessagesRangeAsSeen, p x.RPC_UserParam) (*x.PB_MsgResponse_SetMessagesRangeAsSeen, error) {
	dm, err := NewDirectMessagingByChatId(p.GetUserId(), int(i.ChatId))
	if err != nil {
		return nil, err
	}
	dm.SetMessagesAsSeen(int(i.FromSeq), int(i.EndSeq), int(i.SeenTimeMs))

	return &x.PB_MsgResponse_SetMessagesRangeAsSeen{}, nil
}

func (rpcMsg) DeleteRoomHistory(i *x.PB_MsgParam_DeleteRoomHistory, p x.RPC_UserParam) (*x.PB_MsgResponse_DeleteRoomHistory, error) {
	x.NewDirectToMessage_Deleter().
        ChatId_Eq(int(i.ChatId)).
        Seq_LE(int(i.ToSeq)).
        Delete(base.DB)




}

func (rpcMsg) DeleteMessagesByIds(i *x.PB_MsgParam_DeleteMessagesByIds, p x.RPC_UserParam) (*x.PB_MsgResponse_DeleteMessagesByIds, error) {
	panic("implement me")
}

func (rpcMsg) SetMessagesAsReceived(i *x.PB_MsgParam_SetMessagesAsReceived, p x.RPC_UserParam) (*x.PB_MsgResponse_SetMessagesAsReceived, error) {
	panic("implement me")
}

func (rpcMsg) ForwardMessages(i *x.PB_MsgParam_ForwardMessages, p x.RPC_UserParam) (*x.PB_MsgResponse_ForwardMessages, error) {
	panic("implement me")
}

func (rpcMsg) EditMessage(i *x.PB_MsgParam_EditMessage, p x.RPC_UserParam) (*x.PB_MsgResponse_EditMessage, error) {
	panic("implement me")
}

func (rpcMsg) BroadcastNewMessage(i *x.PB_MsgParam_BroadcastNewMessage, p x.RPC_UserParam) (*x.PB_MsgResponse_BroadcastNewMessage, error) {
	panic("implement me")
}
