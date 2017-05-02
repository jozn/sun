package models

import (
	"errors"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
)

const POST_TYPE_TEXT = 1
const POST_TYPE_PHOTO = 2
const POST_TYPE_VIDEO = 3

func Post_AddNewPostToDbAndItsTagsAndCounters(post *x.Post) {
	err := post.Save(base.DB)
	if err == nil {
		Counter.UpdateUserPostsCounts(post.UserId, 1)
		Tags_AddTagsInPost(post)
		Mentioned_AddUserMentionedInPost(post)
	} else {
		helper.DebugErr(err)
	}
}

func Post_PostsByIds(PostIds []int) ([]*x.Post, error) {
	if len(PostIds) == 0 {
		return nil, errors.New("PostIds must not be empty")
	}
	return x.NewPost_Selector().Id_In(PostIds).GetRows(base.DB)
}
func PostTypeIdToName(id int) string {
	switch id {
	case POST_TYPE_TEXT:
		return "text"

	case POST_TYPE_PHOTO:
		return "photo"

	case POST_TYPE_VIDEO:
		return "video"
	}
	return "none"
}
