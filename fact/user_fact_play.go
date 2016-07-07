package fact

import (
	"encoding/json"
	"fmt"
	. "ms/sun/base"
	. "ms/sun/models"
	"reflect"
	//"strings"
)

func FactPlay1(c *Action) {
	u := User{}
	// u.Id int
	// u.UserId

	u.UserName = "asd"
	u.FirstName = "string"
	u.LastName = "string"
	u.FullName = "string"
	e(u)
	// dbInsetUpdate(&u)
	// s := reflect.ValueOf(&u).Elem()
	// typeOfT := s.Type()
	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s %s = %v\n", i,
	// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
	// }

	// Email string
	// PasswordHash string
	// PasswordSalt string
	// IsProfilePrivate int
	// CreatedTimestamp int

}

func factPlay2(c *Action) {
	//	u := User{}
	//	u.UserName = "asd"
	//	u.FirstName = "string"
	//	u.LastName = "string"
	//	u.FullName = "string"
	//	DbInsertUpdateStruct(&u, "user", true)
	//
	//	var re []test5
	//	var ins []interface{} //{2, 3, 5}
	//	for i := 1; i < 4; i++ {
	//		ins = append(ins, i)
	//	}
	//	e(ins)
	//	// for i := 0; i < 50; i++ {
	//	DB.Select(&re, "select * from test where id in(?,?,?)", ins...) // 2, 3, 5) //ins...)
	//	// }
	//	b, _ := json.Marshal(re)
	//	c.SendText(string(b))

}

func factPlay3(c *Action) {
	var re []User
	DB.Select(&re, "select * from user")
	for _, u := range re {
		u.Email = "wdji@uiuasd.com"
		DbUpdateStruct(&u, "user")
	}
	b, _ := json.Marshal(re)
	c.SendText(string(b))
}

func factPlay4(c *Action) {
	//	var re []User
	//	var usrInf []ViewUser
	//
	//	if len(users) == 0 {
	//		DB.Select(&re, "select * from user")
	//		users = re
	//	} else {
	//		re = users
	//	}
	//
	//	for _, u := range re {
	//		o := ViewUser{}
	//		o.User = &u
	//		o.PlayPLAY = Play1{}
	//		o.Counts = &UserInfo{}
	//		usrInf = append(usrInf, o)
	//		u.Email = "00000000@uiuasd.com"
	//		u.FullName = "00000000uiuasd.com"
	//		// dbUpdateStruct(&u, "user")
	//	}
	//	b, _ := json.Marshal(usrInf)
	//	c.SendText(string(b))
}

func _printStructFilelds(st interface{}) {
	s := reflect.ValueOf(st)
	fmt.Println(s)
	fmt.Println(s.Kind())
	fmt.Println(s.Elem())
	fmt.Println(s.CanSet())

	s = reflect.ValueOf(st).Elem()

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

}
