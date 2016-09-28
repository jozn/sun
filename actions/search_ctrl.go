package actions

import (
	// "strconv"
	// "fmt"
	. "ms/sun/base"
	. "ms/sun/models"
)

type SearchAllView struct {
	Users []User
	Tags  []Tag
}

type TagPostView struct {
	Tag   Tag
	Posts []PostAndDetailes
}

func SearchAllAction(c *Action) {
	devPrintn("Action follow")
	// cuid := c.UserId()

	query := (c.Req.Form["q"])[0]
	devPrintn(query)
	// err := DB.Select(&userRes, "select * from user where UserName like ?% ", query)
	//FIXME: sql-injection
	user := _searchUsers(query)
	tags := _searchTagsCollection(query)
	res := SearchAllView{}
	res.Users = user
	res.Tags = tags
	// fmt.Println(userRes)
	c.SendJson(res)

}

func SearchTagsAction(c *Action) {
	devPrintn("Action searchTagsAction")
	// cuid := c.UserId()

	query := (c.Req.Form["tag"])[0]
	devPrintn(query)
	// err := DB.Select(&userRes, "select * from user where UserName like ?% ", query)
	//FIXME: sql-injection
	var tags []Tag
	sql := "select * from tags where Name = ? limit 1"
	err := DB.Select(&tags, sql, query)
	noErr(err)
	//TODO check for not shuch tag
	tag := tags[0]
	// devPrintn(tag)

	var tagsPost []TagsPost
	err = DB.Select(&tagsPost, "select * from tags_posts where TagId = ? limit 50", tag.Id)
	// err = DB.Select(&tagsPost, "select * from tags_posts") // where TagId = ? limit 50", 1)
	devPrintn(tagsPost)
	noErr(err)
	var pids []int
	for _, tg := range tagsPost {
		pids = append(pids, tg.PostId)
	}
	devPrintn(pids)

	sqlin := intsToSqlIn(pids)
	var posts []Post
	devPrintn(sqlin)
	err = DB.Select(&posts, "select * from post where Id in("+sqlin+") limit 50")
	noErr(err)

	res := TagPostView{}
	res.Posts = GetPostAndDetailes(posts)
	res.Tag = tag

	c.SendJson(res)
	e(tag)
}

func _searchUsers(query string) []User {
	var userRes []User
	sql := "select * from user where UserName like '" + query + "%' or FirstName like '" + query + "%' or LastName like '" + query + "' group by Id limit 10"
	err := DB.Select(&userRes, sql) //, query)
	noErr(err)
	return userRes

}
func _searchTagsCollection(query string) []Tag {
	var tagsRes []Tag
	sql := "select * from tags where Name like '" + query + "%' limit 10"
	err := DB.Select(&tagsRes, sql) //, query)
	noErr(err)
	return tagsRes
}
