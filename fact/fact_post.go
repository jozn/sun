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
    "ms/sun/helper"
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

func FactComment2(c *Action) {
    print("factoring like+comment\n")
    // u.UserId = rand.Intn(50) + 1
    // u.PostId = rand.Intn(450) + 1
    UserId := rand.Intn(_factLastUserId()) + 1
    PostId := rand.Intn(_factLastPostId()) + 1
    Text := helper.FactRandStrEmoji(20,true)
    co :=AddNewComment(UserId,PostId,Text)
    c.SendJson(co)

}

func FactLike2(c *Action) {
	print("factoring likes post\n")
	//COUNT = 50
	l := Like{}
	l.PostId = rand.Intn(500) + 1
	l.UserId = rand.Intn(80) + 1
	l.CreatedTime = now()
	DbInsertStruct(&l, "likes")
}

func FactLike3(c *Action) {
    UserMemoryStore.AddPostLike(rand.Intn(80)+1, rand.Intn(500)+1)
}

func _factLastPostId() int {
	var ps []Post
	DB.Select(&ps, "select * from post order by Id DESC limit 2 ")
	p := ps[0]

	return p.Id

}

func _factLastUserId() int {
	var ps []User
	DB.Select(&ps, "select * from user order by Id DESC limit 2 ")
	p := ps[0]

	return p.Id

}
