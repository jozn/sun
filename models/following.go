package models

////////////////////////////////// new Apis 0.4 ///////////////////////

func UsersToInlineFollowView(userIds []int, cuid int) []UserInlineFollowView {
	var userListView []UserInlineFollowView
	for _, uid := range userIds {
		userView := UserInlineFollowView{}
		userView.UserInlineView = GetUserView(uid)
		if cuid > 0 {
			//userView.AmIFollowing = IsUserFollowing(cuid, userView.UserId)
			userView.IFollowType = MemoryStore.UserFollowingList_GetFollowingTypeForUsers(cuid, uid)
			//debug("GetFollowingTypeForUsers: ", MemoryStore.UserFollowingList_GetFollowingTypeForUsers(cuid, uid))
		}
		userListView = append(userListView, userView)
	}
	return userListView
}

