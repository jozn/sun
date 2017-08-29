package models

import (
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

func pushView_directLogsTo_PB_ChangesHolderView(meId int, logs []x.DirectLog) *x.PB_ChangesHolderView {

	//preload in here
	usersToLoad := make(map[int]bool)
	chatIdsToLoad := make(map[int]bool)
	msgIdsToLoad := []int{}
	for _, log := range logs { //each user
		if log.RoomLogTypeId == int(Push_NEW_DIRECT_MESSAGE) {
			usersToLoad[log.PeerUserId] = true
			chatIdsToLoad[log.ChatId] = true
			msgIdsToLoad = append(msgIdsToLoad, log.MessageId)

		}
	}

	//
	res := &x.PB_ChangesHolderView{}

	for _, logRow := range logs { //each user
		switch x.UpdateLogEnum(logRow.RoomLogTypeId) {
		case Push_NEW_DIRECT_MESSAGE:
			if v, ok := pushView_newDirectMessage(logRow); ok {
				res.NewMessages = append(res.NewMessages, v)
			}

			//metas
		case Push_MESSAGE_RECIVED_TO_SERVER:
			if v, ok := pushView_messageMeta(logRow); ok {
				res.MessagesDelivierdToServer = append(res.MessagesDelivierdToServer, v)
			}

		case Push_MESSAGE_DELIVIERD_TO_PEER:
			if v, ok := pushView_messageMeta(logRow); ok {
				res.MessagesDelivierdToPeer = append(res.MessagesDelivierdToPeer, v)
			}

		case Push_MESSAGE_SEEN_BY_PEER:
			if v, ok := pushView_messageMeta(logRow); ok {
				res.MessagesSeenByPeer = append(res.MessagesSeenByPeer, v)
			}

		case Push_MESSAGE_DELETED_FROM_SERVER:
			if v, ok := pushView_messageMeta(logRow); ok {
				res.MessagesDeletedFromServer = append(res.MessagesDeletedFromServer, v)
			}

		}
	}

	if len(usersToLoad) > 0 {
		res.Users = pushView_userView(meId, usersToLoad)
	}

	if len(chatIdsToLoad) > 0 {
		res.Users = pushView_chatView(meId, chatIdsToLoad)
	}

	//TODO: add messages files id

	return res
}

func pushView_newDirectMessage(log x.DirectLog) (*x.PB_MessageView, bool) {
	v := &x.PB_MessageView{}

	return v, true
}

func pushView_messageMeta(log x.DirectLog) (*x.PB_UpdateMessageMeta, bool) {
	v := &x.PB_UpdateMessageMeta{
		MessageId: int64(log.MessageId),
		AtTime:    int64(log.AtTimeMs / 1000),
	}

	return v, true
}

func pushView_userView(meId int, peerIds map[int]bool) (r []*x.PB_UserView) {
	return
}

func pushView_chatView(meId int, chatIds map[int]bool) (r []*x.PB_UserView) {
	return
}

func pushView_messagesFilesView(meId int, chatIds map[int]bool) (r []*x.PB_UserView) {
	return
}
