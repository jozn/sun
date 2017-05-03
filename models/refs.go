package models

func Ref_Post(PostId int) int {
	return 0
}

func Ref_CommentAdd(CommentId int) int {
	refId := CommentId*1000 + ACTION_TYPE_POST_COMMENTED
	return refId
}

func Ref_LikeAdd(LikeId int) int {
	refId := LikeId*1000 + ACTION_TYPE_POST_LIKED
	return refId
}

func Ref_FollowAdd(FLId int) int {
	refId := FLId*1000 + ACTION_TYPE_FOLLOWED_USER
	return refId
}
