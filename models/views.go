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

func (e _viewImpl) UserBasicAndMeForUsers(CurrentUserId int, Users []int) []UserBasicAndMe {
	res := make([]UserBasicAndMe, 0, len(Users))

	for _, u := range Users {
		user := UserMemoryStore.GetForUser(u)
		if user != nil {
			v := UserBasicAndMe{
				UserId:        u,
				FollowingType: UserMemoryStore.GetFollowingTypeForUsers(CurrentUserId, u),
			}
			v.UserBasic = user.UserBasic
			v.FullName = v.FirstName + " " + v.LastName
			res = append(res, v)
		}
	}

	return res
}
