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

//this is copy of GetListOfUserForFollowType()
//todo merge with UserBasicAndMeForUsers()
//dep use Views.UserBasicAndMeForUsers()
func (e _viewImpl)GetListOfUserForFollowType(userIds []int, CurrentUserId int) *[]UserBasicAndMe {
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
