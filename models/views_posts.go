package models

func (e _viewImpl) PostsViews(posts []*Post, UserId int) (viw []*PostView) {
	// UserSlice
	var photoIds []int
	for _, p := range posts {
		if p.TypeId == 2 {
			photoIds = append(photoIds, p.Id)
		}
	}
	if len(photoIds) > 0 {
		Store.PreLoadPhoto_ByPostIds(photoIds)
	}

	for _, p := range posts {
		//viw = append(viw, PostToPostAndDetailes(&p))
		viw = append(viw, Views.PostSingleView(p, UserId))

	}
	return viw
}

func (e _viewImpl) PostsViewsForPostIds(postsIds []int, UserId int) (viw []*PostView) {
	if len(postsIds) != 0 {
		posts, _ := Post_PostsByIds(postsIds)
		viw = Views.PostsViews(posts, UserId)
	}
	return
}

func (e _viewImpl) PostSingleView(post *Post, UserId int) *PostView {
	v := &PostView{}
	if post != nil {
		v.Post = post
		v.TypeName = PostTypeIdToName(post.TypeId)
		if UserId > 0 {
			v.MyLike = MemoryStore.UserLikedPostsList_MyLiked(UserId, post.Id)
			//v.AmIlike = MemoryStore.UserLikedPostsList_IsLiked(UserId, post.Id) //UserMemoryStore.AmILikePost(UserId, post.Id)
		}
		if post.TypeId == 2 {
			ph, _ := Store.Photo_ByPostId(post.Id)
			v.PhotoView = Convert_PhotoToNewPhotoView(ph)
		}
		u, err := Views.GetUserInlineView(post.UserId)
		if err == nil {
			v.Sender = u
		}
	}
	return v
}
