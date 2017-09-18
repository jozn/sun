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

func PushView_directLogsTo_PB_ChangesHolderView(meId int, logs []*x.DirectLog) *x.PB_PushHolderView {

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
	res := &x.PB_PushHolderView{}

    for _, logRow := range logs { //each user
        switch Push(logRow.RoomLogTypeId) {
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

	/*for _, logRow := range logs { //each user
		switch Push(logRow.RoomLogTypeId) {
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
	}*/

	if len(usersToLoad) > 0 {
		res.Users = pushView_userView(meId, usersToLoad)
	}

	if len(chatIdsToLoad) > 0 {
		res.Chats = pushView_chatView(meId, chatIdsToLoad)
	}

	//TODO: add messages files id

	return res
}

func pushView_newDirectMessage(log *x.DirectLog) (*x.PB_MessageView, bool) {
	if directMsg, ok := x.Store.GetDirectMessageByMessageId(log.MessageId); ok {
		v := PBConv_DirectMessage_to_PB_MessageView(directMsg, log.ChatId)
		return v, true
	}
	return nil, false
}

func pushView_messageMeta(log *x.DirectLog) (*x.PB_UpdateMessageMeta, bool) {
	v := &x.PB_UpdateMessageMeta{
		MessageId: int64(log.MessageId),
		AtTime:    int64(log.AtTimeMs / 1000),
	}

	return v, true
}

func pushView_userView(meId int, peerIds map[int]bool) (res []*x.PB_UserView) {
	for peerId, _ := range peerIds {
		if user, ok := MemoryStore_User.GetUser(peerId); ok {
			v := &x.PB_UserView{
				UserId:         int32(user.Id),
				UserName:       user.UserName,
				FirstName:      user.FirstName,
				LastName:       user.LastName,
				About:          user.About,
				FullName:       user.FullName,
				AvatarUrl:      user.AvatarUrl,
				PrivacyProfile: int32(user.PrivacyProfile),
				//Phone:            user.Phone,
				//Email:            user.Email,
				IsDeleted:        int32(user.IsDeleted),
				FollowersCount:   int32(user.FollowersCount),
				FollowingCount:   int32(user.FollowingCount),
				PostsCount:       int32(user.PostsCount),
				UpdatedTime:      int32(user.UpdatedTime),
				AppVersion:       int32(user.AppVersion),
				LastActivityTime: int32(user.LastActivityTime),
				FollowingType:    int32(MemoryStore.UserFollowingList_GetFollowingTypeForUsers(meId, peerId)),
			}
			res = append(res, v)
		}
	}
	return
}

func pushView_chatView(meId int, chatIds map[int]bool) (res []*x.PB_ChatView) {
	for chatId, _ := range chatIds {
		if chat, ok := x.Store.GetChatByChatId(chatId); ok {
			chatView := &x.PB_ChatView{
				ChatKey:              chat.ChatKey,
				ChatId:               int64(chat.ChatId),
				RoomTypeEnumId:       int32(chat.RoomTypeEnumId),
				UserId:               int32(chat.UserId),
				PeerUserId:           int32(chat.PeerUserId),
				GroupId:              int64(chat.GroupId),
				CreatedTime:          int32(chat.CreatedTime),
				UpdatedMs:            int64(chat.UpdatedMs),
				LastMessageId:        int64(chat.LastMessageId),
				LastDeletedMessageId: int64(chat.LastDeletedMessageId),
				LastSeenMessageId:    int64(chat.LastSeenMessageId),
				LastSeqSeen:          int32(chat.LastSeqSeen),
				LastSeqDelete:        int32(chat.LastSeqDelete),
				CurrentSeq:           int32(chat.CurrentSeq),
			}

            chatView.User = view_getUserVIew(meId,chat.PeerUserId)

			res = append(res, chatView)
		}
	}

	return
}

func view_getUserVIew(meId, peerId int) (v *x.PB_UserView) {
	if user, ok := MemoryStore_User.GetUser(peerId); ok {
		v = &x.PB_UserView{
			UserId:         int32(user.Id),
			UserName:       user.UserName,
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			About:          user.About,
			FullName:       user.FullName,
			AvatarUrl:      user.AvatarUrl,
			PrivacyProfile: int32(user.PrivacyProfile),
			//Phone:            user.Phone,
			//Email:            user.Email,
			IsDeleted:        int32(user.IsDeleted),
			FollowersCount:   int32(user.FollowersCount),
			FollowingCount:   int32(user.FollowingCount),
			PostsCount:       int32(user.PostsCount),
			UpdatedTime:      int32(user.UpdatedTime),
			AppVersion:       int32(user.AppVersion),
			LastActivityTime: int32(user.LastActivityTime),
			FollowingType:    int32(MemoryStore.UserFollowingList_GetFollowingTypeForUsers(meId, peerId)),
		}
	}
	return
}

func pushView_messagesFilesView(meId int, chatIds map[int]bool) (r []*x.PB_UserView) {
	return
}
