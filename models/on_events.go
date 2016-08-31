package models

////////////////// Posts /////////////////////
func OnPostCreated(post *Post) {

}

func OnPostDeleted(post *Post) {

}

//////////////// Likes - Posts //////////////////////
func OnPostLiked(like *Like) {

}

func OnPostUnLiked(like *Like) {

}

/////////////// Comments ////////////////////////
func OnPostCommented(comment *Comment, post *Post) {
    Notification_OnPostCommented(comment,post,true)
}

func OnPostCommentedDelted(comment *Comment, post *Post) {

    Notification_OnPostCommentedDelted(comment,post)

    CacheModels.RemoveCommentById(comment.Id)

}

////////////// Followes /////////////////////////
func OnFollowed(UserId, FollowedPeerUserId int) {
    Notification_OnFollowed(UserId, FollowedPeerUserId)
}

func OnUnFollowed(UserId, FollowedPeerUserId int) {
    Notification_OnUnFollowed(UserId, FollowedPeerUserId)
}
