package models

////////////////// Posts /////////////////////
func OnPostCreated(post *Post) {

}

func OnPostDeleted(post *Post) {

}

//////////////// Likes - Posts //////////////////////
//dep
func OnPostLiked(like *Like) {
	//Notification_OnPostLiked(like)
}
//dep
func OnPostUnLiked(like *Like) {
	//Notification_OnPostUnLiked(like)
}

/////////////// Comments ////////////////////////
func OnPostCommented(comment *Comment, post *Post) {
	Notification_OnPostCommented(comment, post, true)
}

func OnPostCommentedDelted(comment *Comment, post *Post) {

	Notification_OnPostCommentedDelted(comment, post)

	CacheModels.RemoveCommentById(comment.Id)

}

////////////// Followes /////////////////////////
//dep
func OnFollowed(UserId, FollowedPeerUserId int) {
	//Notification_OnFollowed(UserId, FollowedPeerUserId)
}

//dep
func OnUnFollowed(UserId, FollowedPeerUserId int) {
	Notification_OnUnFollowed(UserId, FollowedPeerUserId)
}
