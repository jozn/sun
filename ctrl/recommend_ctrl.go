package ctrl

import (
    "ms/sun/models"
    "ms/sun/base"
    "fmt"
)

func RecommendPostsCtrl(c *base.Action) base.AppErr {
    p := UpdateSessionActivityIfUser(c)

    min := (p.Page -1 ) * p.Limit

    if min < 0 {
        min = 0
    }
    max:= min + p.Limit

    if max > len(models.TopPosts) {
        c.SendJson(models.Recommend_GenTopPosts(200))
        return nil
    }

    fmt.Println(models.TopPosts)

    view:=models.Views.PostsViewsForPostIds(models.TopPosts[min:max],c.UserId())
    //view := models.Views.PostsViews(posts, c.UserId())
    c.SendJson(view)
    return nil
}

