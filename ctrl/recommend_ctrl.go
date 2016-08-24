package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
)

func RecommendPostsCtrl(c *base.Action) base.AppErr  {
    UpdateSessionActivityIfUser(c)
    var posts []models.Post
    err:= base.DB.Select(&posts, "select * from post where TypeId = 2 order by Id Desc limit 50")
    if err != nil{
        return nil
    }

    view:= models.PostsToPostsAndDetailesV1(posts, c.UserId())
    c.SendJson(view)
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


func RecommendTagsWithPostsCtrl(c *base.Action) base.AppErr  {

    rs := models.TopTagsWithPostsResult

    c.SendJson(rs)
    return nil
}
