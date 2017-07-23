package models

import (
	"ms/sun/base"
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

	dm := NewDirectMessaging(p.GetUserId(), pid)
	dm.AddMessage(msg)

}

func (rpcMsg) SetRoomActionDoing(i *x.PB_MsgParam_SetRoomActionDoing, p x.RPC_UserParam) (*x.PB_MsgResponse_SetRoomActionDoing, error) {
	panic("implement me")
}

func (rpcMsg) GetMessagesByIds(i *x.PB_MsgParam_GetMessagesByIds, p x.RPC_UserParam) (*x.PB_MsgResponse_GetMessagesByIds, error) {
	panic("implement me")
}

func (rpcMsg) GetMessagesHistory(i *x.PB_MsgParam_GetMessagesHistory, p x.RPC_UserParam) (*x.PB_MsgResponse_GetMessagesHistory, error) {
	panic("implement me")
}

func (rpcMsg) SetMessagesRangeAsSeen(i *x.PB_MsgParam_SetMessagesRangeAsSeen, p x.RPC_UserParam) (*x.PB_MsgResponse_SetMessagesRangeAsSeen, error) {
	panic("implement me")
}

func (rpcMsg) DeleteRoomHistory(i *x.PB_MsgParam_DeleteRoomHistory, p x.RPC_UserParam) (*x.PB_MsgResponse_DeleteRoomHistory, error) {
	panic("implement me")
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
