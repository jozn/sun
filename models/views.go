package models

import "errors"

type _viewImpl int

var Views _viewImpl

func (e _viewImpl) UserViewSync(CurrentUserId, UserId int) *UserSyncAndMeView {
	u, ok := MemoryStore_User.GetUser(UserId)
	if !ok {
		return &UserSyncAndMeView{}
	}

	v := &UserSyncAndMeView{
		UserId:        u.Id,
		FollowingType: MemoryStore.UserFollowingList_GetFollowingTypeForUsers(CurrentUserId, UserId),
		AppVersion:    u.AppVersion,
		Phone:         MemoryStore.GetPhoneForUserIfIsContact(CurrentUserId, UserId),
		UpdatedTime:   u.UpdatedTime,
	}
	v.UserBasic = u.UserBasic
	return v
}

func (e _viewImpl) UserBasicAndMeView(CurrentUserId, UserId int) *UserBasicAndMe {
	u, ok := MemoryStore_User.GetUser(UserId)
	if !ok {
		return &UserBasicAndMe{}
	}

	v := &UserBasicAndMe{
		UserId:        u.Id,
		FollowingType: MemoryStore.UserFollowingList_GetFollowingTypeForUsers(CurrentUserId, UserId),
		UpdatedTime:   u.UpdatedTime,
	}
	v.UserBasic = u.UserBasic
	return v
}
func (e _viewImpl) UserBasicAndMeForUsers(CurrentUserId int, Users []int) []UserBasicAndMe {
	res := make([]UserBasicAndMe, 0, len(Users))

	for _, u := range Users {
		user, ok := MemoryStore_User.GetUser(u)
		if ok {
			v := UserBasicAndMe{
				UserId:        u,
				FollowingType: MemoryStore.UserFollowingList_GetFollowingTypeForUsers(CurrentUserId, u),
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
func (e _viewImpl) GetListOfUserForFollowType(userIds []int, CurrentUserId int) *[]UserBasicAndMe {
	list := make([]UserBasicAndMe, 0, len(userIds))
	for _, uid := range userIds {
		userView := UserBasicAndMe{}
		peerUser, ok := MemoryStore_User.GetUser(uid)
		if ok {
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

func (e _viewImpl) GetUserInlineView(uid int) (UserInlineView, error) {
	u, ok := MemoryStore_User.GetUser(uid)
	if ok {
		//return *u.ToUserInlineView(), nil
		v := UserInlineView{}
		v.FullName = u.GetFullName()
		v.UserId = u.Id
		v.UserName = u.UserName
		v.AvatarUrl = u.AvatarUrl
		return v, nil
	}
	return UserInlineView{}, errors.New("User NOT Fund")
}
