package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

//psost types: 1:text 2: media/photo

func AddPostAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	txt := c.Req.Form.Get("text")

	post := models.Post{}
	//TODO security: clean html of text
	post.Text = txt
	post.UserId = c.UserId()
	post.TypeId = models.POST_TYPE_TEXT
	post.CreatedTime = helper.TimeNow()
	post.CommentsCount = 0
	post.LikesCount = 0

	//res, _ := base.DbInsertStruct(&post, "post")
	//pid, _ := res.LastInsertId()
	//post.Id = int(pid)

	//models.AddTagsInPost(post)
	//models.AddUserMentionedInPost(post)
	models.AddNewPostToDbAndItsMeta(&post)
	c.SendJson(post)
	return nil
}

func GetSinglePostAction(c *base.Action) base.AppErr {
	idstr := c.Req.Form.Get("post_id")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}
	var post models.Post
	err := base.DB.Get(&post, "select * from post where Id = ? ", id)
	if err != nil {
		return nil
	}

	view := models.PostToPostAndDetailes(&post)
	c.SendJson(view)
	return nil
}

func PostDeleteAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	idstr := c.Req.Form.Get("post_id")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}

	models.DeletePost(c.UserId(), id)

	return nil
}

func PostUpdateAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	idstr := c.Req.Form.Get("post_id")
	text := c.Req.Form.Get("text")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}

	var post models.Post
	err := base.DB.Get(&post, "select * from post where Id = ? ", id)
	if err != nil {
		return nil
	}

	if post.UserId != c.UserId() {
		return nil
	} else {
		post.Text = text
		base.DbUpdateStruct(&post, "post")
		//todo caches?
	}

	return nil
}

func GetPostsStraemAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
    const LIMIT = 30

	laststr := c.Req.Form.Get("last") //last that have
	pagestr := c.Req.Form.Get("page")
	last := helper.StrToInt(laststr, 0)
	page := helper.StrToInt(pagestr, 0)
    limit := c.GetParamInt("limit", 30)

    _ = last
	_ = page

	uid := c.UserId()
	println("LIMIT: ",limit,c.Req.Form.Get("limit")," +")
	//fids := models.GetAllPrimiryFollowingIds(uid)
	//fids := models.UserMemoryStore.GetAllFollowingsListOfUser(uid).Elements
	fids := models.MemoryStore.UserFollowingList_Get(uid).Elements

	selctor := models.NewPost_Selector().UserId_In(fids).OrderBy_Id_Desc().Limit(limit)

	if last > 0 {
		selctor.Id_LT(last)
	} else if page > 0 {
		selctor.Offset((page - 1) * limit)
	}

	posts, err := selctor.GetRows(base.DB)
	if err != nil {
		helper.DebugPrintln(err)
		c.SendJson(nil)
		return err
	}

	view := models.PostsToPostsAndDetailesV1(posts, uid)
	c.SendJson(view)
	return nil
}

func GetPostsLatestAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	laststr := c.Req.Form.Get("last") //last that have
	pagestr := c.Req.Form.Get("page")
	last := helper.StrToInt(laststr, 0)
	page := helper.StrToInt(pagestr, 0)
    //limit := helper.StrToInt("limit", LIMIT)

    _ = last
	_ = page

	uid := c.UserId()

	// dbIns(len(fids))e
	sql := "select * from post  order by Id Desc limit 100 "

	var rs []models.Post
	base.DB.Select(&rs, sql)

	view := models.PostsToPostsAndDetailesV1(rs, uid)
	c.SendJson(view)
	return nil
}

/*
//////////////// For Profile ///////////////////////
type ProfileRespnse struct {
	User  models.UserTable
	Posts []*models.PostAndDetailes
}


//dep
func GetPostsAndInfoForProfileAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	laststr := c.Req.Form.Get("last") //last that have
	pagestr := c.Req.Form.Get("page")
	last := helper.StrToInt(laststr, 0)
	page := helper.StrToInt(pagestr, 0)
	_ = last
	_ = page

	uid := c.UserId()
	profileId := c.GetParamInt("profile_id", 0)
	mem := models.UserMemoryStore.GetForUser(profileId)

	if mem == nil {
		c.Protocol.Error = "NOT FOUND"
		c.Protocol.Status = "ERROR"
		return nil
	}
	u := mem.UserTable

	sql := "select * from post where UserId =" + helper.IntToStr(profileId) + " order by Id Desc limit 100 "

	var rs []models.Post
	base.DB.Select(&rs, sql)

	view := models.PostsToPostsAndDetailesV1(rs, uid)
	res := ProfileRespnse{}
	res.Posts = view
	res.User = u
	c.SendJson(res)
	return nil
}
*/


/*
func GetPostsStraemAction(c *base.Action) base.AppErr {
    MustBeUserAndUpdate(c)
    laststr := c.Req.Form.Get("last") //last that have
    pagestr := c.Req.Form.Get("page")
    last := helper.StrToInt(laststr, 0)
    page := helper.StrToInt(pagestr, 0)
    _ = last
    _ = page

    const LIMIT  = 30
    uid := c.UserId()
    //print(uid)
    //fids := models.GetAllPrimiryFollowingIds(uid)
    fids := models.UserMemoryStore.GetAllFollowingsListOfUser(uid).Elements
    // dbIns(len(fids))e
    sql := "select * from post where UserId in(" + helper.IntsToSqlIn(fids) + ") order by Id Desc limit 100 "

    selctor := models.NewPost_Selector().UserId_In(fids).OrderBy_Id_Desc().Limit(LIMIT)

    if last > 0 {
        selctor.UserId_LT(last)
    }else if page > 0 {
        selctor.Offset(page*LIMIT)
    }

    posts,err := selctor.GetRows(base.DB)
    if err != nil{
        helper.DebugPrintln(err)
        c.SendJson(nil)
        return err
    }
    var rs []models.Post
    base.DB.Select(&rs, sql)

    view := models.PostsToPostsAndDetailesV1(rs, uid)
    c.SendJson(view)
    return nil
}
*/
