package models

import (
	// "encoding/json"
	// "fmt"
	. "ms/sun/base"
)

var POST_TYPE_TEXT = 1
var POST_TYPE_PHOTO = 2
var POST_TYPE_VIDEO = 3

func GetPostAndDetailes(posts []Post) []PostAndDetailes {
	var viw []PostAndDetailes
	// UserSlice
	for _, p := range posts {
		//todo extract this to it's func
		debug(p.UserId)
		v := PostAndDetailes{}
		v.Post = p
		v.TypeName = TypeIdToName(p.TypeId)
		v.Sender = GetUserView(p.UserId)
		v.Comments = GetPostLastComments(p.Id)
		v.Likes = GetPostLastLikes(p.Id)
		SetPostImages(&v)
		v.AmIlike = false
		viw = append(viw, v)
	}
	return viw
}

// _casheCom :=make(map[int][]Comment,100)
// var _casheCom map[int][]Comment
// var _casheLike map[int][]Like

func GetPostLastComments(id int) []CommentInlineInfo {
	debug("GetPostLastComments: ", id)
	var rs []CommentInlineInfo
	casheKey := "post-last-comments-view_" + intToStr(id)
	// v, ok := _casheCom[id]
	v, ok := cashe.Get(casheKey)
	if ok {
		return v.([]CommentInlineInfo)
	}
	var rsCom []Comment
	err := DB.Select(&rsCom, "select * from comments where PostId =? order by Id DESC limit 4 ", id)
	noErr(err)
	// debug(...)
	for i := 0; i < len(rsCom); i++ {
		r := CommentInlineInfo{}
		r.Comment = rsCom[i]
		r.Sender = GetUserView(r.UserId)
		rs = append(rs, r)
	}
	// _casheCom[id] = rs
	cashe.Set(casheKey, rs, minutes(5))
	return rs
}

func GetPostLastLikes(id int) []Like {
	var rs []Like
	// _casheLike
	// v, ok := _casheLike[id]
	casheKey := "post-last-likes_" + intToStr(id)
	v, ok := cashe.Get(casheKey)
	if ok {
		return v.([]Like)
	}
	DB.Select(&rs, "select * from like where PostId =? limit 4 order by Id DESC limit 4 ", id)
	// _casheLike[id] = rs
	cashe.Set(casheKey, rs, minutes(5))

	return rs
}

func SetPostImages(pv *PostAndDetailes) {
	debug("pv.TypeId: ", pv.TypeId)
	if pv.TypeId == 11 { //photo
		var ms []Media //ImageResHolder //for now only one image per post
		var img ImageResHolder
		casheKey := "post-images_" + intToStr(pv.Id)
		_, ok := cashe.Get(casheKey)
		if ok {
			// return v.([]Like)
		}
		DB.Select(&ms, "select * from media where PostId = ? limit 1 ", pv.Id)
		// DB.Select(&ms, "select * from media ") //where PostId = ?  ", pv.Id)
		debug("*********************8media resuults: ", ms)
		if len(ms) == 0 {
			return
		}
		m := ms[0]
		if m.PostId != 0 {
			img = ImageResHolder{}
			//now bulid high low thumb ... resolutions
			hr := ImageResRow{}
			hr.Url = m.Src

			st := ImageResRow{}
			st.Url = m.Src

			lr := ImageResRow{}
			lr.Url = m.Src

			tr := ImageResRow{}
			tr.Url = m.Src

			img.HighRes = &hr
			img.StandardRes = &st
			img.LowRes = &lr
			img.ThumbRes = &tr

			//now set it to orginal post
			pv.Images = &img
		} else {
			// pv.Images = nil
		}
		// _casheLike[id] = rs
		// cashe.Set(casheKey, rs, minutes(5))

	}

}

func GetMedias() {
	var ms []Media
	var us []User
	err := DB.Select(&ms, "select * from media ") //where PostId = ?  ", pv.Id)
	debug(err)
	DB.Select(&us, "select * from User ") //where PostId = ?  ", pv.Id)
	debug("media: ", ms)
	debug("media-user: ", len(us))
}

/////////////

//////////////////  New Apis   ///////////////////////////
func AddTagsInPost(post Post) {
	parser := TextParser{}
	parser.Parse(post.Text)
	for _, tag := range parser.Tags {
		var dbTags []Tag
		var dbTag Tag
		DB.Select(&dbTags, "select * from tags where Name = ? ", tag)
		if len(dbTags) == 0 { //not exist ,insert it
			dbTag = Tag{}
			dbTag.Name = tag
			dbTag.CreatedTime = now()
			res, _ := DbInsertStruct(&dbTag, "tags")
			tid, _ := res.LastInsertId()
			dbTag.Id = int(tid)
		} else {
			dbTag = dbTags[0]
		}

		tagPost := TagPost{}
		tagPost.TagId = dbTag.Id
		tagPost.PostId = post.Id
		tagPost.TypeId = post.TypeId
		tagPost.CreatedTime = now()

		DbInsertStruct(&tagPost, "tags_posts")
		//TODO increment dbTags.Count
	}
}

func AddUserMentionedInPost(post Post) {

}

func PostToPostAndDetailes(post *Post) *PostAndDetailes {
	//todo extract this to it's func
	//debug(p.UserId)
	v := PostAndDetailes{}
	v.Post = *post
	v.TypeName = TypeIdToName(post.TypeId)
	v.Sender = GetUserView(post.UserId)
	v.Comments = GetPostLastComments(post.Id)
	v.Likes = GetPostLastLikes(post.Id)
	SetPostImages(&v)
	v.AmIlike = false

	return &v
}

func PostsToPostsAndDetailes(posts []Post) []*PostAndDetailes {
	var viw []*PostAndDetailes
	// UserSlice
	for _, p := range posts {
		viw = append(viw, PostToPostAndDetailes(&p))
	}
	return viw
}


