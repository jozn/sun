package models

//type UserFollowView struct {
//    UserBasicAndMe
//}

type TopTagsWithPosts struct {
	Tag   Tag
	Posts []*PostAndDetailes
	//Posts []PostAndDetailes
}

type TopTagsWithPostsView struct {
	Tag  Tag
	Post []PostAndDetailes
}

func GetListOfUserForFollowType(userIds []int, CurrentUserId int) *[]UserBasicAndMe {
	list := make([]UserBasicAndMe, 0, len(userIds))
	for _, uid := range userIds {
		userView := UserBasicAndMe{}
		peerUser := UserMemoryStore.GetUserTableForUser(uid)
		if peerUser != nil {
			userView.UserBasic = peerUser.UserBasic
			userView.FullName = peerUser.GetFullName()
			if CurrentUserId > 0 {
				userView.UpdatedTime = peerUser.UpdatedTime
				userView.UserId = peerUser.Id
				userView.FollowingType = MemoryStore.UserFollowingList_GetFollowingTypeForUsers(CurrentUserId, peerUser.Id)
				list = append(list, userView)
			}
		}
	}
	return &list
}

//dep use Views.UserBasicAndMeForUsers
func GetUserBasicAndMe(UserId, CurrentUserId int) *UserBasicAndMe {
	userView := UserBasicAndMe{}
	peerUser := UserMemoryStore.GetUserTableForUser(UserId)
	if peerUser != nil {
		userView.UserBasic = peerUser.UserBasic
		userView.FullName = peerUser.GetFullName()
		if CurrentUserId > 0 {
			userView.UpdatedTime = peerUser.UpdatedTime
			userView.UserId = peerUser.Id
			userView.FollowingType = MemoryStore.UserFollowingList_GetFollowingTypeForUsers(CurrentUserId, peerUser.Id)
		}
	}

	return &userView
}
