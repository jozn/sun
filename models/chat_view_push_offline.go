package models

import (
	"fmt"
	"github.com/jozn/protobuf/proto"
	"ms/sun/config"
	"ms/sun/models/x"
	"ms/sun/models/x/xconst"
)

func ViewPush_DirectOfflinesList_To_GetDirectUpdatesView(meId int, logs []*x.DirectOffline) *x.PB_Offline_Sync {

	//preload in here
	usersToLoad := make(map[int]bool)
	chatKeysToLoad := make(map[string]bool)
	msgIdsToLoad := []int{}
	msgFileIdsToLoad := []int{}
	for _, log := range logs { //each user
		if log.PBClass == xconst.PB_Offline_NewDirectMessage {
			usersToLoad[log.PeerUserId] = true
			chatKeysToLoad[log.ChatKey] = true
			msgIdsToLoad = append(msgIdsToLoad, log.MessageId)
			if log.MessageFileId > 0 {
				msgFileIdsToLoad = append(msgFileIdsToLoad, log.MessageFileId)
			}
		}
	}

	//
	res := &x.PB_Offline_Sync{}

	//todo add perloadings
	for _, logRow := range logs { //each user
		switch logRow.PBClass {
		case xconst.PB_Offline_NewDirectMessage:
			if v, ok := _pushView_newDirectMessageFromOffline(logRow); ok {
				res.NewMessages = append(res.NewMessages, v)
			}

		case xconst.PB_Offline_ChangeMessageId:
			pb := &x.PB_Offline_ChangeMessageId{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.ChangeMessageIds = append(res.ChangeMessageIds, pb)
			}

		case xconst.PB_Offline_ChangeMessageFileId:
			pb := &x.PB_Offline_ChangeMessageFileId{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.ChangeMessageFileIds = append(res.ChangeMessageFileIds, pb)
			}

		case xconst.PB_Offline_MessageToEdit:
			pb := &x.PB_Offline_MessageToEdit{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.MessagesToEdit = append(res.MessagesToEdit, pb)
			}

		case xconst.PB_Offline_MessageToDelete:
			pb := &x.PB_Offline_MessageToDelete{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.MessagesToDelete = append(res.MessagesToDelete, pb)
			}

			//metas
		case xconst.PB_Offline_MessagesReachedServer:
			pb := &x.PB_Offline_MessagesReachedServer{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.MessagesReachedServer = append(res.MessagesReachedServer, pb)
			}

		case xconst.PB_Offline_MessagesDeliveredToUser:
			pb := &x.PB_Offline_MessagesDeliveredToUser{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.MessagesDeliveredToUser = append(res.MessagesDeliveredToUser, pb)
			}

		case xconst.PB_Offline_MessagesSeenByPeer:
			pb := &x.PB_Offline_MessagesSeenByPeer{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.MessagesSeenByPeer = append(res.MessagesSeenByPeer, pb)
			}

		case xconst.PB_Offline_MessagesDeletedFromServer:
			pb := &x.PB_Offline_MessagesDeletedFromServer{}
			err := proto.UnmarshalMerge(logRow.DataPB, pb)
			if err == nil {
				res.MessagesDeletedFromServer = append(res.MessagesDeletedFromServer, pb)
			}
		}

	}

	if len(usersToLoad) > 0 {
		res.Users = ViewUser_GetUserViewList(meId, usersToLoad)
	}

	if len(chatKeysToLoad) > 0 {
		res.Chats = ViewChat_GetChatViewList_ByChatKeys_map(meId, chatKeysToLoad)
	}

	if len(logs) > 0 {
		res.LastId = int64(logs[len(logs)-1].DirectOfflineId)
	}
	//TODO: add messages files id

	return res
}

func _pushView_newDirectMessageFromOffline(log *x.DirectOffline) (*x.PB_MessageView, bool) {
	if directMsg, ok := x.Store.GetDirectMessageByMessageId(log.MessageId); ok {
		v := PBConv_DirectMessage_to_PB_MessageView(directMsg, log.ChatKey)
		if log.MessageFileId > 0 {
			fmt.Println("log.MessageFileId ", log.MessageFileId)
			msgFile, ok := x.Store.GetMessageFileByMessageFileId(log.MessageFileId)
			if ok {
				v.MessageFileView = PBConvPB_MessageFile_To_MessageFile(msgFile)
				v.MessageFileView.ServerSrc = config.CDN_CHAT_MSG_UPLOAD_URL + v.MessageFileView.Name
				fmt.Println("v.MessageFileView ", v.MessageFileView)
			}
		}
		return v, true
	}
	return nil, false
}
