package models

import (
	"ms/sun/base"
	"ms/sun/helper"
)

const POST_TYPE_TEXT = 1
const POST_TYPE_PHOTO = 2
const POST_TYPE_VIDEO = 3

func AddNewPostToDbAndItsMeta(post *Post) {
	post.CreatedTime = helper.TimeNow()
	UserMemoryStore.UpdateUserPostsCounts(post.UserId, 1)

	post.Save(base.DB)

	AddTagsInPost(post)
	AddUserMentionedInPost(post)
}

func GetPosts(PostIds []int) *[]Post {
	sql := "select * from post where Id in(" + helper.IntsToSqlIn(PostIds) + ") order by Id Desc"
	var rs []Post
	base.DB.Select(&rs, sql)

	return &rs
}

func GetPostsAndItsDetails(PostIds []int, UserId int) []*PostAndDetailes {
	rs := GetPosts(PostIds)
	view := PostsToPostsAndDetailesV1(*rs, UserId)
	return view
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

func PostsToPostsAndDetailesNew(posts []*Post, UserId int) (viw []*PostAndDetailesNew) {
    // UserSlice
    var photoIds []int
    photoMap:=  make(map[int]*Photo)
    for _, p := range posts {
        if p.TypeId == 2 {
            photoIds = append(photoIds, p.Id)
        }
    }
    if len(photoIds) > 0 {
        phots,err :=NewPhoto_Selector().PostId_In(photoIds).GetRows(base.DB)
        if err != nil{
            return nil
        }
        for _,p:=range phots{
            photoMap[p.PostId] = p
        }
    }

    for _, p := range posts {
        //viw = append(viw, PostToPostAndDetailes(&p))
        viw = append(viw, GetPostToPostAndDetailesNew(p, UserId,photoMap))

    }
    return viw
}

func GetPostToPostAndDetailesNew(post *Post, UserId int,mp map[int]*Photo) *PostAndDetailesNew {
    v := PostAndDetailesNew{}
    v.Post = *post
    v.TypeName = PostTypeIdToName(post.TypeId)
    u, err := Views.GetUserInlineView(post.UserId)
    if err == nil {
        v.Sender = u
        v.Comments = nil //GetPostLastComments(post.Id)
        v.Likes = nil    //GetPostLastLikes(post.Id)
        //SetPostImages(&v)
        v.AmIlike = MemoryStore.UserLikedPostsList_IsLiked(UserId, post.Id) //UserMemoryStore.AmILikePost(UserId, post.Id)
        if post.TypeId == 2{
            v.Photo =mp[post.Id]
        }
        return &v
    }
    return nil
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
