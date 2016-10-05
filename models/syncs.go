package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

//TODO: must add update users as well
func SyncGetAllChangedUser(CurrentUserId, LastTime int) []UserViewSync {

	//contacts
	var usersContacts []User
	q := "SELECT u.* FROM `user` AS u JOIN phone_contacts AS pc WHERE pc.UserId=? AND pc.CreatedTime >= ? AND u.Phone = pc.PhoneNormalizedNumber AND u.Phone != '' "
	err1 := base.DB.Select(&usersContacts, q, CurrentUserId, LastTime)

	var usersFollowd []User
	q2 := "SELECT u.* FROM `user` AS u JOIN following_list_member AS fm ON u.Id = fm.FollowedUserId WHERE fm.UserId=? AND fm.UpdatedTimeMs >= ?"
	err2 := base.DB.Select(&usersFollowd, q2, CurrentUserId, LastTime*1000)

	usersRes := make([]UserViewSync, 0, len(usersContacts)+len(usersFollowd))

	//first do for followings - Phone filed maybe overide if a user existi in both

	for _, u := range usersFollowd {
		v := UserViewSync{}
		v.FromUser(CurrentUserId, u, false)
		usersRes = append(usersRes, v)
	}

	for _, u := range usersContacts {
		v := UserViewSync{}
		v.FromUser(CurrentUserId, u, true)
		usersRes = append(usersRes, v)
	}

	helper.Debug("SyncGetAllChangedUser()", len(usersContacts), err1, err2)

	return usersRes
}

func UserTableToUserViewCompation(users []User) []User {

	return nil
}
