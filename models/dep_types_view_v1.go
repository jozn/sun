package models

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
