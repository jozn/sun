package models

import (
	"ms/sun/base"
)

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
	//UserMemoryStore.UpdateUserPostsCounts(UserId, -1)
	Counter.UpdateUserPostsCounts(post.UserId, -1)
	return true
}
