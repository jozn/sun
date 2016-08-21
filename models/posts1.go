package models

import (
    "ms/sun/helper"
    "ms/sun/base"
)

/*
func NewPostAddExtraDetails(post *Post)  {

    AddTagsInPost(post)
    AddUserMentionedInPost(post)
}
*/

func AddNewPostToDbAndItsMeta(post *Post) {
    post.CreatedTime = helper.TimeNow()
    UserMemoryStore.UpdateUserPostsCounts(post.UserId,1)

    res, _ := base.DbInsertStruct(post, "post")
    pid, _ := res.LastInsertId()
    post.Id = int(pid)

    AddTagsInPost(post)
    AddUserMentionedInPost(post)
}
