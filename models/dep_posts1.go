package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

func AddNewPostToDbAndItsMeta(post *Post) {
	post.CreatedTime = helper.TimeNow()
	UserMemoryStore.UpdateUserPostsCounts(post.UserId, 1)

	post.Save(base.DB)

	AddTagsInPost(post)
	AddUserMentionedInPost(post)
}

/////////// From version 0.4 /////////////

func DeletePost(UserId, PostId int) bool {
	var post Post
	err := base.DB.Get(&post, "select * from post where Id = ? ", PostId)
	if err != nil {
		return false
	}

	if post.UserId != UserId {
		return false
	}

	base.DbExecute("delete from post where Id = ? limit 1 ", PostId)
	base.DbExecute("delete from likes where PostId = ? And UserId = ? ", PostId, UserId)
	base.DbExecute("delete from comments where PostId = ? And UserId = ? ", PostId, UserId)
	//todo delete from more
	UserMemoryStore.UpdateUserPostsCounts(UserId, -1)

	return true
}

func PostsToPostsAndDetailesV1(posts []Post, UserId int) []*PostAndDetailes {
	var viw []*PostAndDetailes
	// UserSlice
	for _, p := range posts {
		//viw = append(viw, PostToPostAndDetailes(&p))
		viw = append(viw, GetPostToPostAndDetailes(&p, UserId))
	}
	return viw
}

func GetPostToPostAndDetailes(post *Post, UserId int) *PostAndDetailes {
	v := PostAndDetailes{}
	v.Post = *post
	v.TypeName = PostTypeIdToName(post.TypeId)
	u, err := Views.GetUserInlineView(post.UserId)
	if err == nil {
		v.Sender = u
		v.Comments = nil //GetPostLastComments(post.Id)
		v.Likes = nil    //GetPostLastLikes(post.Id)
		//SetPostImages(&v)
		v.AmIlike = MemoryStore.UserLikedPostsList_IsLiked(UserId, post.Id) //UserMemoryStore.AmILikePost(UserId, post.Id)
		return &v
	}
	return nil
}
