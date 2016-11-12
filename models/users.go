package models

import (
	"errors"
	. "ms/sun/base"
)

//dep use GetUserById2
func GetUserById(id int) User {
	key := "user_" + intToStr(id)
	if v, ok := cashe.Get(key); ok {
		return v.(User)
	}
	//var u []User
	var u User
	//debug(u)
	err := DB.Get(&u, "select * from user where Id =? limit 1", id)
	//cashe.Set(key, u[0], 0)
	//noErr(err)
	_ = err
	return u
}

//todo to VIews?
func GetUserView(uid int) UserInlineView {
	//debug("GetUserView: ", uid)
	//u := GetUserById(uid)
	u := UserMemoryStore.GetUserTableForUser(uid)
	v := UserInlineView{}
	v.FullName = u.FirstName + " " + u.LastName
	v.UserId = u.Id
	v.UserName = u.UserName
	v.AvatarUrl = u.AvatarUrl
	return v
}
//todo merge
func GetUserViewV2(uid int) (*UserInlineView, error) {
	u := UserMemoryStore.GetUserTableForUser(uid)
	if u != nil {
		return u.ToUserInlineView(), nil
	}
	return &UserInlineView{}, errors.New("User NOT Fund")
}
