package models

import (
	"log"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/models/x"
)

func ViewChat_GetChatViewList_ByChatKeys_map(meId int, chatKeysMap map[string]bool) (res []*x.PB_ChatView) {
	chatKeys := make([]string, 0, len(chatKeysMap))
	for chatId, _ := range chatKeysMap {
		chatKeys = append(chatKeys, chatId)
	}

	return ViewChat_GetChatViewList_ByChatKeys(meId, chatKeys)
}

func ViewChat_GetChatViewList_ByChatKeys(meId int, chatKeys []string) (res []*x.PB_ChatView) {
	for _, chatId := range chatKeys {
		if chat, ok := x.Store.GetChatByChatKey(chatId); ok {
			if config.IS_DEBUG && !Keys_IsMyChatKey(chatId, meId) {
				log.Panic("try to make chat view for a chat id not belong to current user: in ViewChat_GetChatViewList_ByChatKeys ", chatId, meId)
			}
			chatView := &x.PB_ChatView{
				ChatKey:              chat.ChatKey,
				RoomKey:              chat.RoomKey,
				RoomTypeEnumId:       int32(chat.RoomTypeEnumId),
				UserId:               int32(chat.UserId),
				PeerUserId:           int32(chat.PeerUserId),
				GroupId:              int64(chat.GroupId),
				CreatedSe:            int32(chat.CreatedSe),
				UpdatedMs:            int64(chat.UpdatedMs),
				LastMessageId:        int64(chat.LastMessageId),
				LastDeletedMessageId: int64(chat.LastDeletedMessageId),
				LastSeenMessageId:    int64(chat.LastSeenMessageId),
			}

			chatView.UserView = UserView_GetUserView(meId, chat.PeerUserId)

			res = append(res, chatView)
		}
	}

	return
}

func ViewChat_GetDirectMessageViewList_ByMsgIds(meId int, msgIds []int) (res []*x.PB_MessageView) {

	msgs, err := x.NewDirectMessage_Selector().MessageId_In(msgIds).GetRows(base.DB)
	if err != nil {
		return
	}
	msgFileIdsToLoad := []int{}

	for _, msg := range msgs {
		if msg.MessageFileId > 0 {
			msgFileIdsToLoad = append(msgFileIdsToLoad, msg.MessageFileId)
		}
	}

	x.Store.PreLoadMessageFileByMessageFileIds(msgFileIdsToLoad)

	for _, msg := range msgs {
		v := PBConv_DirectMessage_to_PB_MessageView(msg, "not_implemented")
		if msg.MessageFileId > 0 {
			//fmt.Println("log.MessageFileId ", log.MessageFileId)
			msgFile, ok := x.Store.GetMessageFileByMessageFileId(msg.MessageFileId)
			if ok {
				v.MessageFileView = PBConvPB_MessageFile_To_MessageFile(msgFile)
				v.MessageFileView.ServerSrc = config.CDN_CHAT_MSG_UPLOAD_URL + v.MessageFileView.Name
				//fmt.Println("v.MessageFileView ", v.MessageFileView)
			}
		}
		res = append(res, v)
	}

	return
}

//migrate to useview.go
func ViewUser_GetUserViewList(meId int, peerIds map[int]bool) (res []*x.PB_UserView) {
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

func UserView_GetUserView(meId, peerId int) (v *x.PB_UserView) {
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
