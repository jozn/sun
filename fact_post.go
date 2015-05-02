package main

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	// "strings"
	"math/rand"
	"strconv"
	"time"
)

func factPost1(c *Action) {
	print("factoring pst + media\n")
	u := Post{}
	u.UserId = rand.Intn(50)
	u.Text = factRnddStr(300)
	u.CreatedTimestamp = int(time.Now().Unix()) - rand.Intn(50000)
	u.TypeId = 10 //text see: gloab_types
	// ui := UserInfo{}

	dbInsertStruct(&u, "post")
	// res, err := dbInsertStruct(&u, "post")
	// noErr(err)
	// id, _ := res.LastInsertId()
	// ui.UserId = int(id)
	// dbInsertStruct(&ui, "user_info")
	// e(u)

}

func factPost2(c *Action) {
	print("factoring pst + media\n")
	u := Post{}
	uid := rand.Intn(50)
	u.UserId = uid
	u.Text = factRnddStr(300)
	u.CreatedTimestamp = int(time.Now().Unix()) - rand.Intn(50000)
	u.TypeId = 10 //text see: gloab_types
	// ui := UserInfo{}
	if rand.Intn(2) == 1 {
		u.TypeId = 11 //photo
		res, err := dbInsertStruct(&u, "post")
		noErr(err)
		id, _ := res.LastInsertId()
		m := Media{}
		m.PostId = int(id)
		m.UserId = uid
		m.Src = "/public/photo/" + strconv.Itoa(rand.Intn(50)) + "_960.jpg"
		_, err = dbInsertStruct(&m, "media")
		noErr(err)
		return
	}

	dbInsertStruct(&u, "post")
	// res, err := dbInsertStruct(&u, "post")
	// noErr(err)
	// id, _ := res.LastInsertId()
	// ui.UserId = int(id)
	// dbInsertStruct(&ui, "user_info")
	// e(u)

}

func factLike1(c *Action) {
	print("factoring like+comment\n")
	u := Comment{}
	u.UserId = rand.Intn(50)
	u.PostId = rand.Intn(100)
	u.Text = factRnddStr(15)
	u.CreatedTimestamp = int(time.Now().Unix()) - rand.Intn(50000)
	dbInsertStruct(&u, "comment")

	l := Like{}
	l.UserId = rand.Intn(50)
	l.PostId = rand.Intn(100)
	// l.Text = factRnddStr(15)
	l.CreatedTimestamp = int(time.Now().Unix()) - rand.Intn(50000)
	dbInsertStruct(&u, "comment")

}
