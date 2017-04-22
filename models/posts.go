package models

import (
	"errors"
	"ms/sun/base"
)

const POST_TYPE_TEXT = 1
const POST_TYPE_PHOTO = 2
const POST_TYPE_VIDEO = 3

func Post_AddNewPostToDbAndItsTagsAndCounters(post *Post) {
	err := post.Save(base.DB)
	if err == nil {
		Counter.UpdateUserPostsCounts(post.UserId, 1)
		Tags_AddTagsInPost(post)
		Mentioned_AddUserMentionedInPost(post)
	}
}

func Post_PostsByIds(PostIds []int) ([]*Post, error) {
	if len(PostIds) == 0 {
		return nil, errors.New("PostIds must not be empty")
	}
	return NewPost_Selector().Id_In(PostIds).GetRows(base.DB)
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
