package models

import (
	"errors"
	"fmt"
	"github.com/jozn/protobuf/proto"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"ms/sun/models/x/xconst"
	"os"
)

type rpcChat int

func (rpcChat) AddNewMessage(param *x.PB_ChatParam_AddNewMessage, userParam x.RPC_UserParam) (res x.PB_ChatResponse_AddNewMessage, err error) {
	if param.MessageView == nil {
		return
	}

	iMsg := param.MessageView
	timeing := int(iMsg.CreatedSe)
	if timeing < (helper.TimeNow()-24*3600) || timeing > helper.TimeNow() {
		timeing = helper.TimeNow()
	}

	peerId := RoomKeyToOtherUser(iMsg.RoomKey, userParam.GetUserId())
	if peerId < 1 {
		return res, errors.New("peerId is not valid")
	}

	msg := &x.DirectMessage{
		MessageId:            helper.NextRowsSeqId(),
		MessageKey:           iMsg.MessageKey,
		RoomKey:              iMsg.RoomKey,
		UserId:               userParam.GetUserId(),
		MessageFileId:        int(iMsg.MessageFileId),
		MessageTypeEnumId:    int(iMsg.MessageTypeEnumId),
		Text:                 iMsg.Text,
		CreatedSe:            int(timeing),
		PeerReceivedTime:     0,
		PeerSeenTime:         0,
		DeliviryStatusEnumId: int(x.RoomMessageDeliviryStatusEnum_SENT),
	}

	var msgFile *x.MessageFile
	if iMsg.MessageFileView != nil {
		f := iMsg.MessageFileView
		msgFile = &x.MessageFile{
			MessageFileId:   (helper.NextRowsSeqId()),
			MessageFileKey:  f.MessageFileKey, //,KeyNewMessageKey(p.GetUserId()), //todo must set at client
			OriginalUserId:  userParam.GetUserId(),
			Name:            f.Name,
			Size:            int(f.Size),
			FileTypeEnumId:  int(iMsg.MessageTypeEnumId),
			Width:           int(f.Width),
			Height:          int(f.Height),
			Duration:        int(f.Duration),
			Extension:       f.Extension,
			HashMd5:         helper.MD5BytesToString(param.Blob),
			HashAccess:      int(f.HashAccess),
			CreatedSe:       int(helper.TimeNow()),
			ServerSrc:       "",
			ServerPath:      "",
			ServerThumbPath: f.LocalSrc,
			BucketId:        "",
			ServerId:        0,
			CanDel:          0,
		}

		if len(param.Blob) > 0 {
			dirName := fmt.Sprintf("./upload/message_files/")
			os.MkdirAll(dirName, 0666)

			fileName := fmt.Sprintf("%s%d%s", dirName, msgFile.MessageFileId, msgFile.Extension)
			newFile, err1 := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
			if err1 != nil {
				println("add message file : creating file failed")
				return res, err1
			}
			defer newFile.Close()
			newFile.Write(param.Blob)
		}

		err := msgFile.Save(base.DB)
		if err == nil {
			msg.MessageFileId = msgFile.MessageFileId
			msg.MessageTypeEnumId = int(x.RoomMessageTypeEnum_IMAGE_TEXT)
		}
		//helper.NoErr(err)
	}

	dm := NewDirectMessagingByUsers(userParam.GetUserId(), peerId)

	//PUSH NEW MESSAGE
	//sent to both me (for mutliple device and other peer new msg)
	for _, us := range [][]int{{peerId, userParam.GetUserId()}, {userParam.GetUserId(), peerId}} {
		pbNewMsg := &x.PB_Offline_NewDirectMessage{
			ChatKey:       UsersToChatKey(us[0], us[1]),
			FromMessageId: int64(msg.MessageId),
			AtTime:        int64(helper.TimeNowMs()),
		}
		pbByte, _ := proto.Marshal(pbNewMsg)

		dOffNewMsg := x.DirectOffline{
			DirectOfflineId: helper.NextRowsSeqId(),
			ToUserId:        peerId,
			ChatKey:         UsersToChatKey(us[0], us[1]),
			MessageId:       msg.MessageId,
			PBClass:         xconst.PB_Offline_NewDirectMessage,
			DataPB:          pbByte,
			DataJson:        helper.ToJson(pbNewMsg),
			DataTemp:        "",
			AtTimeMs:        helper.TimeNowMs(),
		}
		LiveOfflineFramer.HereDirectDelayer <- dOffNewMsg
	}

	//PUSH CHANGE MESSAGE ID
	pbChangeMsgId := &x.PB_Offline_ChangeMessageId{
		MessageKey:   iMsg.MessageKey,
		NewMessageId: int64(msg.MessageId),
	}
	pbByte, _ := proto.Marshal(pbChangeMsgId)

	dOffNewMsg := x.DirectOffline{
		DirectOfflineId: helper.NextRowsSeqId(),
		ToUserId:        peerId,
		ChatKey:         UsersToChatKey(userParam.GetUserId(), peerId),
		PBClass:         xconst.PB_Offline_ChangeMessageId,
		DataPB:          pbByte,
		DataJson:        helper.ToJson(pbChangeMsgId),
		DataTemp:        "",
		AtTimeMs:        helper.TimeNowMs(),
	}
	LiveOfflineFramer.HereDirectDelayer <- dOffNewMsg

	var pbChangeMsgFileId *x.PB_Offline_ChangeMessageFileId
	//PUSH CHANGE MESSAGE File ID
	if msgFile != nil {
		pbChangeMsgFileId = &x.PB_Offline_ChangeMessageFileId{
			MessageFileKey:   iMsg.MessageFileView.MessageFileKey,
			NewMessageFileId: int64(msgFile.MessageFileId),
		}
		pbByte, _ := proto.Marshal(pbChangeMsgFileId)

		dOff := x.DirectOffline{
			DirectOfflineId: helper.NextRowsSeqId(),
			ToUserId:        peerId,
			ChatKey:         UsersToChatKey(userParam.GetUserId(), peerId),
			PBClass:         xconst.PB_Offline_ChangeMessageFileId,
			DataPB:          pbByte,
			DataJson:        helper.ToJson(pbChangeMsgFileId),
			DataTemp:        "",
			AtTimeMs:        helper.TimeNowMs(),
		}
		LiveOfflineFramer.HereDirectDelayer <- dOff
	}

	//MEASSAGE REACHED SERVER OFFLINE
	pbMsgReachedServer := &x.PB_Offline_MessagesReachedServer{
		MessageKeys: []string{iMsg.MessageKey},
		AtTime:      int64(helper.TimeNow()),
	}
	pbByte, _ = proto.Marshal(pbMsgReachedServer)

	dOffMsgReached := x.DirectOffline{
		DirectOfflineId: helper.NextRowsSeqId(),
		ToUserId:        peerId,
		ChatKey:         UsersToChatKey(userParam.GetUserId(), peerId),
		PBClass:         xconst.PB_Offline_MessagesReachedServer,
		DataPB:          pbByte,
		DataJson:        helper.ToJson(pbMsgReachedServer),
		DataTemp:        "",
		AtTimeMs:        helper.TimeNowMs(),
	}
	LiveOfflineFramer.HereDirectDelayer <- dOffMsgReached

	dm.AddMessage(msg)

	return res, nil
}

func (rpcChat) SetRoomActionDoing(param *x.PB_ChatParam_SetRoomActionDoing, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetRoomActionDoing, err error) {
	panic("implement me")
}

func (rpcChat) SetMessagesRangeAsSeen(param *x.PB_ChatParam_SetChatMessagesRangeAsSeen, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetChatMessagesRangeAsSeen, err error) {
	msgIds, err := x.NewDirectToMessage_Selector().Select_MessageId().ChatKey_Eq(param.ChatKey).
		MessageId_GE(int(param.BottomMessageId)).MessageId_LE(int(param.TopMessageId)).
		GetIntSlice(base.DB)
	if err != nil && len(msgIds) > 0 {
		x.NewDirectMessage_Updater().PeerSeenTime(int(param.SeenTimeMs / 1000)).
			MessageId_In(msgIds).PeerSeenTime_Eq(0).Update(base.DB)

		//MEASSAGE REACHED SERVER OFFLINE
		roomKey, _ := Keys_ChatKeyToRoomKey(param.ChatKey)
		pbMsgSeen := &x.PB_Offline_MessagesSeenByPeer{
			MessageIds: helper.SliceIntToInt64(msgIds),
			RoomKey:    roomKey,
			AtTime:     int64(helper.TimeNow()),
		}
		pbByte, _ := proto.Marshal(pbMsgSeen)
		peerId := RoomKeyToOtherUser(roomKey, userParam.GetUserId())
		if peerId > 0 {
			dOffMsgReached := x.DirectOffline{
				DirectOfflineId: helper.NextRowsSeqId(),
				ToUserId:        peerId,
				ChatKey:         UsersToChatKey(userParam.GetUserId(), peerId),
				PBClass:         xconst.PB_Offline_MessagesSeenByPeer,
				DataPB:          pbByte,
				DataJson:        helper.ToJson(pbMsgSeen),
				DataTemp:        "",
				AtTimeMs:        helper.TimeNowMs(),
			}
			LiveOfflineFramer.HereDirectDelayer <- dOffMsgReached
		}
	}
	return
}

//todo add offlines_for chats
func (rpcChat) DeleteChatHistory(param *x.PB_ChatParam_DeleteChatHistory, userParam x.RPC_UserParam) (res x.PB_ChatResponse_DeleteChatHistory, err error) {
	x.NewDirectToMessage_Deleter().ChatKey_Eq(param.ChatKey).MessageId_LE(int(param.FromMessageId)).Delete(base.DB)
	x.NewChat_Updater().StartMessageIdFrom(int(param.FromMessageId) + 1).
		LastDeletedMessageId(int(param.FromMessageId)).
		ChatKey_Eq(param.ChatKey).Update(base.DB)

	return
}

func (rpcChat) DeleteMessagesByIds(param *x.PB_ChatParam_DeleteMessagesByIds, userParam x.RPC_UserParam) (res x.PB_ChatResponse_DeleteMessagesByIds, err error) {
	panic("implement me")
}

func (rpcChat) SetMessagesAsReceived(param *x.PB_ChatParam_SetMessagesAsReceived, userParam x.RPC_UserParam) (res x.PB_ChatResponse_SetMessagesAsReceived, err error) {
	if len(param.MessageIds) > 0 {
		cnt, _ := x.NewDirectMessage_Updater().PeerReceivedTime(helper.TimeNow()).
			MessageId_In(helper.SliceInt64ToInt(param.MessageIds)).PeerReceivedTime(0).Update(base.DB)

		if cnt > 0 {
			roomKey, _ := Keys_ChatKeyToRoomKey(param.ChatRoom)
			//MEASSAGE REACHED SERVER OFFLINE
			pbMsgRecivied := &x.PB_Offline_MessagesDeliveredToUser{
				MessageIds: param.MessageIds,
				RoomKey:    roomKey,
				AtTime:     int64(helper.TimeNow()),
			}
			pbByte, _ := proto.Marshal(pbMsgRecivied)
			peerId := RoomKeyToOtherUser(roomKey, userParam.GetUserId())
			if peerId > 0 {
				dOffMsgReached := x.DirectOffline{
					DirectOfflineId: helper.NextRowsSeqId(),
					ToUserId:        peerId,
					ChatKey:         UsersToChatKey(userParam.GetUserId(), peerId),
					PBClass:         xconst.PB_Offline_MessagesDeliveredToUser,
					DataPB:          pbByte,
					DataJson:        helper.ToJson(pbMsgRecivied),
					DataTemp:        "",
					AtTimeMs:        helper.TimeNowMs(),
				}
				LiveOfflineFramer.HereDirectDelayer <- dOffMsgReached
			}
		}

	}
	return
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
