package models

import (
	"math/rand"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

type rpcMsg int

func (rpcMsg) UploadNewMsg(i *x.PB_Message, p x.RPC_UserParam) (*x.PB_ResRpcAddMsg, error) {
    panic("implement me")
}

func (rpcMsg) GetFreshChatList(i *x.PB_MsgParam_GetFreshChatList, p x.RPC_UserParam) (*x.PB_MsgResponse_GetFreshChatList, error) {
    panic("implement me")
}

func (rpcMsg) GetFreshRoomMessagesList(i *x.PB_MsgParam_GetFreshRoomMessagesList, p x.RPC_UserParam) (*x.PB_MsgResponse_GetFreshRoomMessagesList, error) {
    panic("implement me")
}


func (rpcMsg) Echo(i *x.PB_MsgParam_Echo, p x.RPC_UserParam) (*x.PB_MsgResponse_PB_MsgParam_Echo, error) {
	//fmt.Println("in Echo --> ", i.Text)
	return &x.PB_MsgResponse_PB_MsgParam_Echo{
		Text: i.Text + " ECHO =====" + helper.FactRandStrEmoji(10, true),
	}, nil
}

func (rpcMsg) AddNewTextMessage(i *x.PB_MsgParam_AddNewTextMessage, p x.RPC_UserParam) (*x.PB_MsgResponse_AddNewTextMessage, error) {
	pid := int(i.PeerId)
	msg := &x.DirectMessage{
		MessageId:            helper.NextRowsSeqId(),
		RoomKey:              UsersToRoomKey(p.GetUserId(), int(i.GetPeerId())),
		UserId:               p.GetUserId(),
		MessageFileId:        0,
		MessageTypeEnumId:    int(x.RoomMessageTypeEnum_TEXT),
		Text:                 i.Text,
		Time:                 int(i.Time),
		PeerReceivedTime:     0,
		PeerSeenTime:         0,
		DeliviryStatusEnumId: int(x.RoomMessageDeliviryStatusEnum_SENT),
	}

	dm := NewDirectMessagingByUsers(p.GetUserId(), pid)
	dm.AddMessage(msg)

	res := &x.PB_MsgResponse_AddNewTextMessage{}
	return res, nil
}

func (rpcMsg) AddNewMessage(i *x.PB_MsgParam_AddNewMessage, p x.RPC_UserParam) (*x.PB_MsgResponse_AddNewMessage, error) {

	pid := int(i.PeerId)
	msg := &x.DirectMessage{
		MessageId:            helper.NextRowsSeqId(),
		RoomKey:              UsersToRoomKey(p.GetUserId(), int(i.GetPeerId())),
		UserId:               p.GetUserId(),
		MessageFileId:        0,
		MessageTypeEnumId:    int(x.RoomMessageTypeEnum_TEXT),
		Text:                 i.Text,
		Time:                 int(i.Time),
		PeerReceivedTime:     0,
		PeerSeenTime:         0,
		DeliviryStatusEnumId: int(x.RoomMessageDeliviryStatusEnum_SENT),
	}

	if i.MessageView != nil && i.MessageView.MessageFileView != nil {
		f := i.MessageView.MessageFileView
		igPb := &x.MessageFile{
			MessageFileId:   (helper.NextRowsSeqId()),
			OriginalUserId: p.GetUserId(),
			Name:            f.Name,
			Size:            int(f.Size),
			FileTypeEnumId:  1,
			Width:           int(f.Width),
			Height:          int(f.Height),
			Duration:        0,
			Extension:       ".jpg",
			HashMd5:         helper.MD5BytesToString(i.Blob),
			HashAccess:      int(rand.Int63n(9e12)),
			CreatedSe:       int(helper.TimeNow()),
			ServerSrc:       "",
			ServerPath:      "",
			ServerThumbPath: f.LocalSrc,
			BucketId:        "",
			ServerId:        0,
			CanDel:          0,
		}

		err := igPb.Save(base.DB)
		helper.NoErr(err)
		msg.MessageFileId = igPb.MessageFileId
		msg.MessageTypeEnumId = int(x.RoomMessageTypeEnum_IMAGE_TEXT)
	}

	dm := NewDirectMessagingByUsers(p.GetUserId(), pid)
	dm.AddMessage(msg)

	res := &x.PB_MsgResponse_AddNewMessage{}
	return res, nil
}

func (rpcMsg) SetRoomActionDoing(i *x.PB_MsgParam_SetRoomActionDoing, p x.RPC_UserParam) (*x.PB_MsgResponse_SetRoomActionDoing, error) {
	panic("implement me")
}

func (rpcMsg) GetMessagesByIds(i *x.PB_MsgParam_GetMessagesByIds, p x.RPC_UserParam) (*x.PB_MsgResponse_GetMessagesByIds, error) {
	ids := helper.SliceInt64ToInt(i.MessagesCollections.DirectMessagesIds)
	x.Store.PreLoadDirectMessageByMessageIds(ids)

	views := make([]*x.PB_MessageView, len(ids))
	for i, id := range ids {
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

func (rpcMsg) SetMessagesRangeAsSeen(i *x.PB_MsgParam_SetChatMessagesRangeAsSeen, p x.RPC_UserParam) (*x.PB_MsgResponse_SetChatMessagesRangeAsSeen, error) {
	dm, err := NewDirectMessagingByChatId(p.GetUserId(), int(i.ChatId))
	if err != nil {
		return nil, err
	}
	dm.SetMessagesAsSeen(int(i.FromSeq), int(i.EndSeq), int(i.SeenTimeMs))

	return &x.PB_MsgResponse_SetChatMessagesRangeAsSeen{}, nil
}

func (rpcMsg) DeleteChatHistory(i *x.PB_MsgParam_DeleteChatHistory, p x.RPC_UserParam) (*x.PB_MsgResponse_DeleteChatHistory, error) {
	panic("implement me")
}

func (rpcMsg) DeleteMessagesByIds(i *x.PB_MsgParam_DeleteMessagesByIds, p x.RPC_UserParam) (*x.PB_MsgResponse_DeleteMessagesByIds, error) {
	//dm := NewDirectMessagingByUsers(p.GetUserId(), pid)
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

/////////////////////////////
