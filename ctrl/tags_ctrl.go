package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
	"strings"
)

func TagsPostsListCtrl(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	limit := c.GetParamInt("limit", 30)
	last := c.GetParamInt("last", 0)
	page := c.GetParamInt("page", 1)

	tagName := c.Req.FormValue("tag")
	tagName = strings.Replace(tagName, "#", "", -1)

	tag, err := models.NewTag_Selector().Name_EQ(tagName).GetRow(base.DB)

	if err != nil {
		c.SendJson(nil)
		return err
	}

	selector := models.NewTagsPost_Selector().Select_PostId().TagId_EQ(tag.Id).OrderBy_Id_Desc().Limit(limit)
	if last > 0 {
		selector.Id_LE(last)
	} else {
		selector.Offset((page - 1) * limit)
	}

	ids, err := selector.GetIntSlice(base.DB)

	if err != nil {
		c.SendJson(nil)
		return err
	}

	res := models.GetPostsAndItsDetails(ids, c.UserId())

	c.SendJson(res)
	return nil
}

func TagsPostsListCtrl_OLD(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	tagName := c.Req.FormValue("tag")
	tagName = strings.Replace(tagName, "#", "", -1)

	tag, ok := models.TagsMap.GetTag(tagName)

	if !ok {
		c.SendJson(nil)
		return nil
	}

	var ids []int

	base.DB.Select(&ids, "select PostId from tags_posts where TagId =? order by id desc limit 25", tag.Id)

	res := models.GetPostsAndItsDetails(ids, c.UserId())

	c.SendJson(res)
	return nil

}
