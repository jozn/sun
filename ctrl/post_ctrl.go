package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"ms/sun/models/x"
)

//psost types: 1:text 2: media/photo

func GetSinglePostAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	idstr := c.Req.Form.Get("post_id")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}
	var post *x.Post
	err := base.DB.Get(post, "select * from post where Id = ? ", id)
	if err != nil {
		return nil
	}

	view := models.Views.PostSingleView(post, c.UserId())
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

	var post x.Post
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
