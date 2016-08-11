package fact

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	// "strings"
	"math/rand"
	. "ms/sun/base"
	. "ms/sun/models"
	"strconv"
	"time"
)

func FactPost1(c *Action) {
	print("factoring pst + media\n")
	u := Post{}
	u.UserId = rand.Intn(50)
	u.Text = factRnddStr(300)
	u.CreatedTime = int(time.Now().Unix()) - rand.Intn(50000)
	u.TypeId = 10 //text see: gloab_types
	// ui := UserInfo{}

	DbInsertStruct(&u, "post")
	// res, err := DbInsertStruct(&u, "post")
	// noErr(err)
	// id, _ := res.LastInsertId()
	// ui.UserId = int(id)
	// DbInsertStruct(&ui, "user_info")
	// e(u)

}

func FactPost2(c *Action) {
	print("factoring pst + media\n")
	u := Post{}
	uid := rand.Intn(50)
	u.UserId = uid
	u.Text = factRnddStr(300)
	u.CreatedTime = int(time.Now().Unix()) - rand.Intn(50000)
	u.TypeId = 10 //text see: gloab_types
	// ui := UserInfo{}
	if rand.Intn(2) == 1 {
		u.TypeId = 11 //photo
		res, err := DbInsertStruct(&u, "post")
		noErr(err)
		id, _ := res.LastInsertId()
		m := Media{}
		m.PostId = int(id)
		m.UserId = uid
		m.Src = "/public/photo/" + strconv.Itoa(rand.Intn(50)) + "_960.jpg"
		_, err = DbInsertStruct(&m, "media")
		noErr(err)
		return
	}

	DbInsertStruct(&u, "post")
	// res, err := DbInsertStruct(&u, "post")
	// noErr(err)
	// id, _ := res.LastInsertId()
	// ui.UserId = int(id)
	// DbInsertStruct(&ui, "user_info")
	// e(u)

}

func FactLike1(c *Action) {
	print("factoring like+comment\n")
	u := Comment{}
	// u.UserId = rand.Intn(50) + 1
	// u.PostId = rand.Intn(450) + 1
	u.UserId = rand.Intn(_factLastUserId()) + 1
	u.PostId = rand.Intn(_factLastPostId()) + 1
	u.Text = factRnddStr(15)
	u.CreatedTime = int(time.Now().Unix()) - rand.Intn(50000)
	DbInsertStruct(&u, "comments")

	l := Like{}
	l.UserId = rand.Intn(_factLastUserId()) + 1
	l.PostId = rand.Intn(_factLastPostId()) + 1
	// l.Text = factRnddStr(15)
	l.CreatedTime = int(time.Now().Unix()) - rand.Intn(50000)
	DbInsertStruct(&l, "likes")

}

func FactComment1(c *Action) {
	print("factoring like+comment\n")
	u := Comment{}
	// u.UserId = rand.Intn(50) + 1
	// u.PostId = rand.Intn(450) + 1
	u.UserId = rand.Intn(_factLastUserId()) + 1
	u.PostId = rand.Intn(_factLastPostId()) + 1
	u.Text = factRnddStr(15)
	u.CreatedTime = int(time.Now().Unix()) - rand.Intn(50000)
	DbInsertStruct(&u, "comments")

}
