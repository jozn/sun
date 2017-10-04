package models

import "ms/sun/models/x"

func ViewChat_GetChatViewList(meId int, chatIds map[int]bool) (res []*x.PB_ChatView) {
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

			chatView.UserView = UserView_GetUserView(meId, chat.PeerUserId)

			res = append(res, chatView)
		}
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
