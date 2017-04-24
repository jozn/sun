package models

type PostView struct {
	// *Post
	*Post
	TypeName string //for text, photo video
	//Comments []CommentInlineInfo
	//Likes    []Like
	Photo     *Photo
	PhotoView *PhotoView
	//Images   *ImageResHolder
	AmIlike bool //dep
	MyLike  int  //type of like
	Sender  UserInlineView
}

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
			v.MyLike = 0
			v.AmIlike = MemoryStore.UserLikedPostsList_IsLiked(UserId, post.Id) //UserMemoryStore.AmILikePost(UserId, post.Id)
		}
		if post.TypeId == 2 {
			ph, _ := Store.Photo_ByPostId(post.Id)
			v.Photo = ph
			v.PhotoView = Convert_PhotoToNewPhotoView(ph)
		}
		u, err := Views.GetUserInlineView(post.UserId)
		if err == nil {
			v.Sender = u
		}
	}
	return v
}

func (e _viewImpl) PostSingleView_OLD(post *Post, UserId int, mp map[int]*Photo) *PostView {
	v := &PostView{}
	if post != nil {
		v.Post = post
		v.TypeName = PostTypeIdToName(post.TypeId)
		u, err := Views.GetUserInlineView(post.UserId)
		if err == nil {
			v.Sender = u
			//v.Comments = nil //GetPostLastComments(post.Id)
			//v.Likes = nil    //GetPostLastLikes(post.Id)
			//SetPostImages(&v)
			v.AmIlike = MemoryStore.UserLikedPostsList_IsLiked(UserId, post.Id) //UserMemoryStore.AmILikePost(UserId, post.Id)
			if post.TypeId == 2 {
				v.Photo = mp[post.Id]
				v.PhotoView = Convert_PhotoToNewPhotoView(mp[post.Id])
			}
			return v
		}
	}
	return v
}
