package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
)

func TagsPostsListCtrl(c *base.Action) base.AppErr  {

    UpdateSessionActivityIfUser(c)

    tagName:= c.Req.FormValue("tag")

    tag,ok :=models.TagsMap.GetTag(tagName)

    if !ok {
       return nil
        c.SendJson("")
    }

    var ids []int

    base.DB.Select(&ids, "select PostId from tags_posts where TagId =? order by id desc limit 25", tag.Id)

    res := models.GetPostsAndItsDetails(ids, c.UserId())

    c.SendJson(res)
    return nil

}

