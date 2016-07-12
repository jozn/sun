package models

import (
	. "ms/sun/base"
	"errors"
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

func GetUserById2(id int) UserTable {
	var u UserTable
	DB.Get(&u, "select * from user where Id =? limit 1", id)
	return u
}

//todo cahcse with GetUserById
func GetUserByUsername(username string) User {
	key := "user_" + username
	if v, ok := cashe.Get(key); ok {
		return v.(User)
	}
	var u []User
	debug(u)
	err := DB.Select(&u, "select * from user where UserName =? limit 1", username)
	cashe.Set(key, u[0], 0)
	noErr(err)
	return u[0]
}

func GetUserByUsername2(username string) (User ,error)  {
	key := "user_" + username
	if v, ok := cashe.Get(key); ok {
		return v.(User) ,nil
	}
	var u []User
	debug(u)
	err := DB.Select(&u, "select * from user where UserName =? limit 1", username)
	if err != nil || len(u)==0{
		return User{},errors.New("error babe")
	}
	cashe.Set(key, u[0], 0)
	return u[0],nil
}

func GetUserInfo(id int) UserInfo {
	// key := "user_" + username
	// if v, ok := cashe.Get(key); ok {
	// 	return v.(User)
	// }
	var u []UserInfo
	// debug(u)
	err := DB.Select(&u, "select * from user_info where UserId =? limit 1", id)
	// cashe.Set(key, u[0], 0)
	noErr(err)
	return u[0]
}

func GetUserView(uid int) UserInlineView {
	//debug("GetUserView: ", uid)
	//u := GetUserById(uid)
	u := UserMemoryStore.Map[uid]
	v := UserInlineView{}
	v.FullName = u.FirstName + " " + u.LastName
	v.UserId = u.Id
	v.UserName = u.UserName
	v.AvatarSrc = u.AvatarSrc
	return v
}
