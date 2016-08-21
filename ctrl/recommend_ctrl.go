package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
)

func RecommendPostsCtrl(c *base.Action) base.AppErr  {

    return nil
}

func RecommendUsersCtrl(c *base.Action) base.AppErr  {
    MustBeUserAndUpdate(c)

    var ids []int

    base.DB.Select(&ids, "select TargetId from recommend_user where UserId =?", c.UserId())

    res := models.UsersToInlineFollowView(ids, c.UserId())

    c.SendJson(res)
    return nil
}

func RecommendTagsCtrl(c *base.Action) base.AppErr  {
    c.SendJson(models.TopTags)
    return nil
}
