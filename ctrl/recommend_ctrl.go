package ctrl

import (
    "ms/sun/models"
    "ms/sun/base"
)

func RecommendPostsCtrl(c *base.Action) base.AppErr {
    p := UpdateSessionActivityIfUser(c)

    min := (p.Page -1 ) * p.Limit

    if min < 0 {
        min = 0
    }
    max:= min + p.Limit

    //fmt.Println(models.TopPosts)

    if max > len(models.TopPosts) {
        c.SendJson(nil)
        return nil
    }


    view:=models.Views.PostsViewsForPostIds(models.TopPosts[min:max],c.UserId())
    //view := models.Views.PostsViews(posts, c.UserId())
    c.SendJson(view)
    return nil
}

