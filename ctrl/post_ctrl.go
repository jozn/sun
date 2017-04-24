package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

//psost types: 1:text 2: media/photo

func GetSinglePostAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
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

	view := models.GetPostToPostAndDetailes(&post, c.UserId())
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
	println("LIMIT: ", limit, c.Req.Form.Get("limit"), " +")
	//fids := models.GetAllPrimiryFollowingIds(uid)
	//fids := models.UserMemoryStore.GetAllFollowingsListOfUser(uid).Elements
	fids := models.MemoryStore.UserFollowingList_Get(uid).Elements
	ins := make([]int, 0, len(fids))
	copy(ins, fids)
	ins = append(ins, c.UserId())
	selctor := models.NewPost_Selector().UserId_In(ins).OrderBy_Id_Desc().Limit(limit)

	if last > 0 {
		selctor.Id_LT(last)
	} else if page > 0 {
		selctor.Offset((page - 1) * limit)
	}

	posts, err := selctor.GetRows2(base.DB)
	if err != nil {
		helper.DebugPrintln(err)
		c.SendJson(nil)
		return err
	}

	view := models.PostsToPostsAndDetailesV1(posts, uid)
	c.SendJson(view)
	return nil
}

