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
				userView.FollowingType = UserMemoryStore.GetFollowingTypeForUsers(CurrentUserId, peerUser.Id)
				list = append(list, userView)
			}
		}
	}
	return &list
}
