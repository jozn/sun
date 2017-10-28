package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

type rpcChat int

func (rpcChat) AddNewMessage(param *x.PB_ChatParam_AddNewMessage, userParam x.RPC_UserParam) (res x.PB_ChatResponse_AddNewMessage, err error) {
	panic("implement me")
}

func (rpcChat) SetRoomActionDoing(param *x.PB_ChatParam_SetRoomActionDoing, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetRoomActionDoing, err error) {
	panic("implement me")
}

func (rpcChat) SetMessagesRangeAsSeen(param *x.PB_ChatParam_SetChatMessagesRangeAsSeen, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetChatMessagesRangeAsSeen, err error) {
	panic("implement me")
}

func (rpcChat) DeleteChatHistory(param *x.PB_ChatParam_DeleteChatHistory, userParam x.RPC_UserParam) (res x.PB_ChatResponse_DeleteChatHistory, err error) {
	panic("implement me")
}

func (rpcChat) DeleteMessagesByIds(param *x.PB_ChatParam_DeleteMessagesByIds, userParam x.RPC_UserParam) (res x.PB_ChatResponse_DeleteMessagesByIds, err error) {
	panic("implement me")
}

func (rpcChat) SetMessagesAsReceived(param *x.PB_ChatParam_SetMessagesAsReceived, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetMessagesAsReceived, err error) {
	panic("implement me")
}

func (rpcChat) EditMessage(param *x.PB_ChatParam_EditMessage, userParam x.RPC_UserParam) (res x.PB_ChatResponse_EditMessage, err error) {
	panic("implement me")
}

func (rpcChat) GetChatList(param *x.PB_ChatParam_GetChatList, userParam x.RPC_UserParam) (res x.PB_ChatResponse_GetChatList, err error) {
	res.Chats, err = Chat_GetChatListForUser(userParam.GetUserId())
	return
}

//fixme :1 this must first load from direct-to_message 2: fix chatId
func (rpcChat) GetChatHistoryToOlder(param *x.PB_ChatParam_GetChatHistoryToOlder, userParam x.RPC_UserParam) (res x.PB_ChatResponse_GetChatHistoryToOlder, err error) {
	roomKey, _ := Keys_ChatKeyToRoomKey(param.ChatKey)
	rows, err := x.NewDirectMessage_Selector().RoomKey_Eq(roomKey).OrderBy_MessageId_Desc().GetRows(base.DB)
	if err != nil {
		return
	}

	for _, r := range rows {
		chat := PBConv_DirectMessage_to_PB_MessageView(r, "")
		res.MessagesViews = append(res.MessagesViews, chat)
	}
	return
}

func (rpcChat) GetFreshAllDirectMessagesList(param *x.PB_ChatParam_GetFreshAllDirectMessagesList, userParam x.RPC_UserParam) (res x.PB_ChatResponse_GetFreshAllDirectMessagesList, err error) {
	chatKeys, err := x.NewChat_Selector().Select_ChatKey().UserId_Eq(userParam.GetUserId()).GetStringSlice(base.DB)
	if err != nil {
		return
	}

	msgIds, err := x.NewDirectToMessage_Selector().Select_MessageId().ChatKey_In(chatKeys).GetIntSlice(base.DB)
	if err != nil {
		return
	}

	res.Messages = ViewChat_GetDirectMessageViewList_ByMsgIds(userParam.GetUserId(), msgIds)

	return
}

/*
func (rpcChat) GetChatList(i *x.PB_ChatParam_GetChatList, p x.RPC_UserParam) (*x.PB_ChatResponse_GetChatList, error) {
    chats, err := Chat_GetChatListForUser(p.GetUserId())
    return &x.PB_ChatResponse_GetChatList{
        Chats: chats,
    }, err
}

func (rpcChat) GetChatHistoryToOlder(i *x.PB_ChatParam_GetChatHistoryToOlder, p x.RPC_UserParam) (*x.PB_ChatResponse_GetChatHistoryToOlder, error) {
    var res []*x.PB_MessageView
    roomKey,_ := Keys_ChatKeyToRoomKey(i.ChatKey)
    rows, err := x.NewDirectMessage_Selector().RoomKey_Eq(roomKey).OrderBy_MessageId_Desc().GetRows(base.DB)
    if err != nil {
        return
    }

    for _, r := range rows {
        chat := PBConv_DirectMessage_to_PB_MessageView(r, "")
        res = append(res, chat)
    }

}

func (rpcChat) GetFreshAllDirectMessagesList(i *x.PB_ChatParam_GetFreshAllDirectMessagesList, p x.RPC_UserParam) (*x.PB_ChatResponse_GetFreshAllDirectMessagesList, error) {
    panic("implement me")
}


*/
