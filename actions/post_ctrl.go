package actions

import (
	. "ms/sun/base"
	. "ms/sun/models"
)

//todo: implment AmILike
func GetPostsAction(c *Action) {
	devPrintn("action post")
	c.MustUser()
	uid := c.UserId()
	print(uid)
	fids := GetAllPrimiryFollowingIds(uid)
	// dbIns(len(fids))e
	sql := "select * from post where UserId in(" + intsToSqlIn(fids) + ") order by Id Desc limit 100 "

	//debuging
	e(fids)
	// sql := "select * from post  limit 1000 "

	debug(sql)
	var rs []Post
	DB.Select(&rs, sql)
	// noErr(err)
	// print(len(rs))
	// b, _ := json.Marshal(rs)
	// c.SendText(string(b))

	// var viw []PostAndDetailes

	// for _, p := range rs {
	// 	//todo extract this to it's func
	// 	debug(p.UserId)
	// 	v := PostAndDetailes{}
	// 	v.Post = p
	// 	v.TypeName = typeIdToName(p.TypeId)
	// 	v.Sender = GetUserView(p.UserId)
	// 	v.Comments = GetPostLastComments(p.Id)
	// 	v.Likes = GetPostLastLikes(p.Id)
	// 	SetPostImages(&v)
	// 	v.AmIlike = false
	// 	viw = append(viw, v)
	// }
	viw := GetPostAndDetailes(rs)
	c.SendJson(viw)
}

func AddPostAction(c *Action) {
	c.MustUser()
	c.MustPost()

	txts := c.Req.Form["text"]
	text := txts[0]

	post := Post{}
	//TODO security: clean html of text
	post.Text = text
	post.CreatedTimestamp = now()
	post.TypeId = TypeNameToId("text")
	post.UserId = c.UserId()

	res, _ := DbInsertStruct(&post, "post")
	pid, _ := res.LastInsertId()
	post.Id = int(pid)

	parsed := TextParser{}
	parsed.Parse(text)

	//add tags
	AddTagsToPost(post, parsed)
	//# addMentionedPost(post, parsed)
	c.SendJson(parsed)
	devPrintn(parsed)

	// c.SendText("yes")
	e(pid)
}

func AddTagsToPost(post Post, parser TextParser) {
	for _, tag := range parser.Tags {
		var dbTags []Tag
		var dbTag Tag
		DB.Select(&dbTags, "select * from tags where Name = ? ", tag)
		if len(dbTags) == 0 { //not exist ,insert it
			dbTag = Tag{}
			dbTag.Name = tag
			dbTag.CreatedTimestamp = now()
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
		tagPost.CreatedTimestamp = now()

		DbInsertStruct(&tagPost, "tags_posts")
		//TODO increment dbTags.Count
	}
}

func addMentionedPost(post Post, parser TextParser) {

}
