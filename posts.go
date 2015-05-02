package main

import (
// "encoding/json"
// "fmt"
)

func getPostsAtction(c *Action) {
	devPrintn("action post")
	var rs []Post
	DB.Select(&rs, "select * from post") //, rand.Intn(1000000))
	// noErr(err)
	// print(len(rs))
	// b, _ := json.Marshal(rs)
	// c.SendText(string(b))

	var viw []PostAndDetailes

	for _, p := range rs {
		v := &PostAndDetailes{}
		v.Post = &p
		v.Comments = getPostLastComments(p.Id)
		v.Likes = getPostLastLikes(p.Id)
		v.AmIlike = false
		viw = append(viw, *v)
	}

	c.SendJson(viw)

	// var rs []test5
	// for i := 0; i < 50; i++ {
	// 	DB.Select(&rs, "select * from test whers id = ?", rand.Intn(1000000))
	// }
	// b, _ := json.Marshal(rs)
	// c.SendText(string(b))
}

// _casheCom :=make(map[int][]Comment,100)
// var _casheCom map[int][]Comment
// var _casheLike map[int][]Like

func getPostLastComments(id int) []Comment {
	debug("getPostLastComments: ", id)
	var rs []Comment
	casheKey := "post-last-comments_" + intToStr(id)
	// v, ok := _casheCom[id]
	v, ok := cashe.Get(casheKey)
	if ok {
		return v.([]Comment)
	}
	err := DB.Select(&rs, "select * from comment where PostId =? order by Id DESC limit 4 ", id)
	debug(err)
	// _casheCom[id] = rs
	cashe.Set(casheKey, rs, minutes(5))
	return rs
}

func getPostLastLikes(id int) []Like {
	var rs []Like
	// _casheLike
	// v, ok := _casheLike[id]
	casheKey := "post-last-likes_" + intToStr(id)
	v, ok := cashe.Get(casheKey)
	if ok {
		return v.([]Like)
	}
	DB.Select(&rs, "select * from like where PostId =? limit 4 order by Id DESC limit 4 ", id)
	// _casheLike[id] = rs
	cashe.Set(casheKey, rs, minutes(5))

	return rs
}
