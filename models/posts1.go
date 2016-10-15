package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

/*
func NewPostAddExtraDetails(post *Post)  {

    AddTagsInPost(post)
    AddUserMentionedInPost(post)
}
*/

func AddNewPostToDbAndItsMeta(post *Post) {
	post.CreatedTime = helper.TimeNow()
	UserMemoryStore.UpdateUserPostsCounts(post.UserId, 1)

	post.Save(base.DB)
	//
	//res, _ := base.DbInsertStruct(post, "post")
	//pid, _ := res.LastInsertId()
	//post.Id = int(pid)

	AddTagsInPost(post)
	AddUserMentionedInPost(post)
}

func GetPosts(PostIds []int) *[]Post {
	sql := "select * from post where Id in(" + helper.IntsToSqlIn(PostIds) + ") order by Id Desc"
	var rs []Post
	base.DB.Select(&rs, sql)

	return &rs
}

func GetPostsAndItsDetails(PostIds []int, UserId int) []*PostAndDetailes {
	rs := GetPosts(PostIds)
	view := PostsToPostsAndDetailesV1(*rs, UserId)
	return view
}
