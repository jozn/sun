package models

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/models/x"
)

type Push int

const (
	Push_NEW_DIRECT_MESSAGE          Push = 1
	Push_CHANGE_MESSAGE_ID           Push = 2
	Push_MESSAGE_RECIVED_TO_SERVER   Push = 3
	Push_MESSAGE_DELIVIERD_TO_PEER   Push = 4
	Push_MESSAGE_SEEN_BY_PEER        Push = 5
	Push_MESSAGE_DELETED_FROM_SERVER Push = 6
	Push_MESSAGE_UPDATE_BY_USER      Push = 7
	Push_MESSAGE_DELETE_BY_USER      Push = 8
	Push_ROOM_ACTION_DOING           Push = 10
)

func DirectSync_GetSync(me, last int) (*x.PB_SyncResponse_GetDirectUpdates, error) {
	rows, err := x.NewDirectUpdate_Selector().ToUserId_Eq(me).DirectUpdateId_GT(last).OrderBy_DirectUpdateId_Asc().GetRows(base.DB)
	if err != nil {
		return nil, err
	}
	return DirectSync_directUpdatesTo_PB_SyncResponse_GetDirectUpdates(me, rows), nil

}


func DirectSync_directUpdatesTo_PB_SyncResponse_GetDirectUpdates(meId int, logs []*x.DirectUpdate) *x.PB_SyncResponse_GetDirectUpdates {

	//preload in here
	usersToLoad := make(map[int]bool)
	chatKeysToLoad := make(map[string]bool)
	msgIdsToLoad := []int{}
	msgFileIdsToLoad := []int{}
	for _, log := range logs { //each user
		if log.RoomLogTypeId == int(Push_NEW_DIRECT_MESSAGE) {
			usersToLoad[log.PeerUserId] = true
			chatKeysToLoad[log.ChatKey] = true
			msgIdsToLoad = append(msgIdsToLoad, log.MessageId)
			if log.MessageFileId > 0 {
				msgFileIdsToLoad = append(msgFileIdsToLoad, log.MessageFileId)
			}
		}
	}

	//
	res := &x.PB_SyncResponse_GetDirectUpdates{}

	for _, logRow := range logs { //each user
		switch Push(logRow.RoomLogTypeId) {
		case Push_NEW_DIRECT_MESSAGE:
			if v, ok := _pushView_newDirectMessage(logRow); ok {
				res.NewMessages = append(res.NewMessages, v)
			}

			//metas
		case Push_MESSAGE_RECIVED_TO_SERVER:
			if v, ok := _pushView_messageMeta(logRow); ok {
				res.MessagesDelivierdToServer = append(res.MessagesDelivierdToServer, v)
			}

		case Push_MESSAGE_DELIVIERD_TO_PEER:
			if v, ok := _pushView_messageMeta(logRow); ok {
				res.MessagesDelivierdToPeer = append(res.MessagesDelivierdToPeer, v)
			}

		case Push_MESSAGE_SEEN_BY_PEER:
			if v, ok := _pushView_messageMeta(logRow); ok {
				res.MessagesSeenByPeer = append(res.MessagesSeenByPeer, v)
			}

		case Push_MESSAGE_DELETED_FROM_SERVER:
			if v, ok := _pushView_messageMeta(logRow); ok {
				res.MessagesDeletedFromServer = append(res.MessagesDeletedFromServer, v)
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
		res.LastId = int64(logs[len(logs)-1].DirectUpdateId)
	}
	//TODO: add messages files id

	return res
}

func _pushView_newDirectMessage(log *x.DirectUpdate) (*x.PB_MessageView, bool) {
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

func _pushView_messageMeta(log *x.DirectUpdate) (*x.PB_UpdateMessageMeta, bool) {
	v := &x.PB_UpdateMessageMeta{
		MessageId: int64(log.MessageId),
		AtTime:    int64(log.AtTimeMs / 1000),
	}

	return v, true
}
