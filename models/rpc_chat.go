package models

import "ms/sun/models/x"

type rpcChat int

func (rpcChat) AddNewMessage(i *x.PB_ChatParam_AddNewMessage, p x.RPC_UserParam) (*x.PB_ChatResponse_AddNewMessage, error) {
    panic("implement me")
}

func (rpcChat) SetRoomActionDoing(i *x.PB_ChatParam_SetRoomActionDoing, p x.RPC_UserParam) (*x.PB_ChatResponse_SetRoomActionDoing, error) {
    panic("implement me")
}

func (rpcChat) SetMessagesRangeAsSeen(i *x.PB_ChatParam_SetChatMessagesRangeAsSeen, p x.RPC_UserParam) (*x.PB_ChatResponse_SetChatMessagesRangeAsSeen, error) {
    panic("implement me")
}

func (rpcChat) DeleteChatHistory(i *x.PB_ChatParam_DeleteChatHistory, p x.RPC_UserParam) (*x.PB_ChatResponse_DeleteChatHistory, error) {
    panic("implement me")
}

func (rpcChat) DeleteMessagesByIds(i *x.PB_ChatParam_DeleteMessagesByIds, p x.RPC_UserParam) (*x.PB_ChatResponse_DeleteMessagesByIds, error) {
    panic("implement me")
}

func (rpcChat) SetMessagesAsReceived(i *x.PB_ChatParam_SetMessagesAsReceived, p x.RPC_UserParam) (*x.PB_ChatResponse_SetMessagesAsReceived, error) {
    panic("implement me")
}

func (rpcChat) EditMessage(i *x.PB_ChatParam_EditMessage, p x.RPC_UserParam) (*x.PB_ChatResponse_EditMessage, error) {
    panic("implement me")
}

func (rpcChat) GetChatList(i *x.PB_ChatParam_GetChatList, p x.RPC_UserParam) (*x.PB_ChatResponse_GetChatList, error) {
    panic("implement me")
}

func (rpcChat) GetChatHistoryToOlder(i *x.PB_ChatParam_GetChatHistoryToOlder, p x.RPC_UserParam) (*x.PB_ChatResponse_GetChatHistoryToOlder, error) {
    panic("implement me")
}

func (rpcChat) GetFreshAllDirectMessagesList(i *x.PB_ChatParam_GetFreshAllDirectMessagesList, p x.RPC_UserParam) (*x.PB_ChatResponse_GetFreshAllDirectMessagesList, error) {
    panic("implement me")
}



