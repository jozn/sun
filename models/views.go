package models

type _viewImpl int

var Views _viewImpl

func (e _viewImpl) UserViewSync(CurrentUserId, UserId int) *UserViewSyncAndMe {
	u := UserMemoryStore.GetForUser(UserId)
	if u == nil {
		return &UserViewSyncAndMe{}
	}

	v := &UserViewSyncAndMe{
		UserId:        u.Id,
		FollowingType: UserMemoryStore.GetFollowingTypeForUsers(CurrentUserId, UserId),
		AppVersion:    u.AppVersion,
		Phone:         MemoryStore.GetPhoneForUserIfIsContact(CurrentUserId, UserId),
		UpdatedTime:   u.UpdatedTime,
	}
	v.UserBasic = u.UserBasic
	return v
}
