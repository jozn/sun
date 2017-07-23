package models

import "ms/sun/models/x"

type rpcMsg int

func (s rpcMsg) AddNewTextMessage(i *x.PB_MsgParam_AddNewTextMessage) (*x.PB_MsgResponse_AddNewTextMessage, error) {
    
}

func (s rpcMsg) SetRoomActionDoing(i *x.PB_MsgParam_SetRoomActionDoing) (*x.PB_MsgResponse_SetRoomActionDoing, error) {
    panic("implement me")
}

func (s rpcMsg) GetMessagesByIds(i *x.PB_MsgParam_GetMessagesByIds) (*x.PB_MsgResponse_GetMessagesByIds, error) {
    panic("implement me")
}

func (s rpcMsg) GetMessagesHistory(i *x.PB_MsgParam_GetMessagesHistory) (*x.PB_MsgResponse_GetMessagesHistory, error) {
    panic("implement me")
}

func (s rpcMsg) SetMessagesRangeAsSeen(i *x.PB_MsgParam_SetMessagesRangeAsSeen) (*x.PB_MsgResponse_SetMessagesRangeAsSeen, error) {
    panic("implement me")
}

func (s rpcMsg) DeleteRoomHistory(i *x.PB_MsgParam_DeleteRoomHistory) (*x.PB_MsgResponse_DeleteRoomHistory, error) {
    panic("implement me")
}

func (s rpcMsg) DeleteMessagesByIds(i *x.PB_MsgParam_DeleteMessagesByIds) (*x.PB_MsgResponse_DeleteMessagesByIds, error) {
    panic("implement me")
}

func (s rpcMsg) SetMessagesAsReceived(i *x.PB_MsgParam_SetMessagesAsReceived) (*x.PB_MsgResponse_SetMessagesAsReceived, error) {
    panic("implement me")
}

func (s rpcMsg) ForwardMessages(i *x.PB_MsgParam_ForwardMessages) (*x.PB_MsgResponse_ForwardMessages, error) {
    panic("implement me")
}

func (s rpcMsg) EditMessage(i *x.PB_MsgParam_EditMessage) (*x.PB_MsgResponse_EditMessage, error) {
    panic("implement me")
}

func (s rpcMsg) BroadcastNewMessage(i *x.PB_MsgParam_BroadcastNewMessage) (*x.PB_MsgResponse_BroadcastNewMessage, error) {
    panic("implement me")
}





