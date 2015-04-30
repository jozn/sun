package main

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	//"strings"
	"time"
)

func factUser1(c *Action) {
	print("factoring user + user_info\n")
	_userFirstNameSamples := []string{"حمید", "سیما", "نیلوفر", "آرش", "Armin", "Leili", "نیاز", "فرخ", "atash", "محمد علی"}
	_userLastNameSamples := []string{"کریمی", "کمانگیر", "بزگ", "فدیایش", "رستگار", "میلانی", "مستانی", "فروهی", "مصداق", "هدایت", "fish", "Nosrat", "Fadaghi"}
	_userUserNameSamples := []string{"fish", "Nosrat", "Fadaghi", "atash", "Jigar", "DooSTi"}

	u := User{}
	// u.Id int
	// u.UserId
	u.UserName = randSilceString(_userUserNameSamples)
	u.FirstName = randSilceString(_userFirstNameSamples)
	u.LastName = randSilceString(_userLastNameSamples)
	u.FullName = u.FirstName + " " + u.LastName
	u.Email = _randomEmail()
	u.CreatedTimestamp = int(time.Now().Unix())

	ui := UserInfo{}

	res, err := dbInsertStruct(&u, "user")
	noErr(err)
	id, _ := res.LastInsertId()
	ui.UserId = int(id)
	dbInsertStruct(&ui, "user_info")
	e(u)

}
