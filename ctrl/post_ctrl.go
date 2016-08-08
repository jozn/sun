package ctrl

import (
"ms/sun/models"
"ms/sun/base"
    "ms/sun/helper"
)

//psost types: 1:text 2: media/photo

func AddPostAction(c *base.Action) base.AppErr {
    c.MustUser()
    //c.MustPost()
    txt := c.Req.Form.Get("text")

    post := models.Post{}
    //TODO security: clean html of text
    post.Text = txt
    post.UserId = c.UserId()
    post.TypeId = models.POST_TYPE_TEXT
    post.CreatedTime = helper.TimeNow()
    post.CommentsCount = 0
    post.LikesCount = 0

    res, _ := base.DbInsertStruct(&post, "post")
    pid, _ := res.LastInsertId()
    post.Id = int(pid)
    models.AddTagsInPost(post)
    models.AddUserMentionedInPost(post)

     c.SendJson(post)
    return nil
}

func GetSinglePostAction(c *base.Action) base.AppErr {
    idstr := c.Req.Form.Get("post_id")
    id:= helper.StrToInt(idstr,0)
    if id < 1{
        return nil
    }
    var post models.Post
    err:= base.DB.Get(&post, "select * from post where Id = ? ", id)
    if err != nil{
        return nil
    }

    view:= models.PostToPostAndDetailes(&post)
    c.SendJson(view)
    return nil
}

func PostDeleteAction(c *base.Action) base.AppErr {
    idstr := c.Req.Form.Get("post_id")
    id:= helper.StrToInt(idstr,0)
    if id < 1{
        return nil
    }

    var post models.Post
    err:= base.DB.Get(&post, "select * from post where Id = ? ", id)
    if err != nil{
        return nil
    }


    if post.UserId != c.UserId() {
        return  nil
    } else {
        base.DB.Exec("delete from post where Id = ? limit 1 ", post.Id)
        base.DB.Exec("delete from likes where PostId = ? And UserId = ? ", post.Id , post.UserId)
        base.DB.Exec("delete from comments where PostId = ? And UserId = ? ", post.Id , post.UserId)
        //todo delete from more
    }

    return nil
}

func PostUpdateAction(c *base.Action) base.AppErr {
    idstr := c.Req.Form.Get("post_id")
    text := c.Req.Form.Get("text")
    id:= helper.StrToInt(idstr,0)
    if id < 1{
        return nil
    }

    var post models.Post
    err:= base.DB.Get(&post, "select * from post where Id = ? ", id)
    if err != nil{
        return nil
    }

    if post.UserId != c.UserId() {
        return  nil
    } else {
       post.Text = text
        base.DbUpdateStruct(&post,"post")
       //todo caches?
    }

    return nil
}

func GetPostsStraemAction(c *base.Action) base.AppErr {
    laststr := c.Req.Form.Get("last")//last that have
    pagestr := c.Req.Form.Get("page")
    last:= helper.StrToInt(laststr,0)
    page:= helper.StrToInt(pagestr,0)
    _ = last; _ = page

    uid := c.UserId()
    print(uid)
    fids := models.GetAllPrimiryFollowingIds(uid)
    // dbIns(len(fids))e
    sql := "select * from post where UserId in(" + helper.IntsToSqlIn(fids) + ") order by Id Desc limit 100 "


    var rs []models.Post
    base.DB.Select(&rs, sql)

    view:= models.PostsToPostsAndDetailes(rs)
    c.SendJson(view)
    return nil
}

func GetPostsLatestAction(c *base.Action) base.AppErr {
    laststr := c.Req.Form.Get("last")//last that have
    pagestr := c.Req.Form.Get("page")
    last:= helper.StrToInt(laststr,0)
    page:= helper.StrToInt(pagestr,0)
    _ = last; _ = page

    // dbIns(len(fids))e
    sql := "select * from post  order by Id Desc limit 100 "


    var rs []models.Post
    base.DB.Select(&rs, sql)

    view:= models.PostsToPostsAndDetailes(rs)
    c.SendJson(view)
    return nil
}

//////////////// For Profile ///////////////////////
type ProfileRespnse struct  {
    User models.UserTable
    Posts []*models.PostAndDetailes

}

func GetPostsForProfileAction(c *base.Action) base.AppErr {
    laststr := c.Req.Form.Get("last")//last that have
    pagestr := c.Req.Form.Get("page")
    last:= helper.StrToInt(laststr,0)
    page:= helper.StrToInt(pagestr,0)
    _ = last; _ = page

    profileId := c.GetParamInt("profile_id",0)
    mem:=models.UserMemoryStore.GetForUser(profileId)

    if mem == nil {
        c.Protocol.Error = "NOT FOUND"
        c.Protocol.Status = "ERROR"
        return nil
    }
    u := mem.UserTable

    sql := "select * from post where UserId =" + helper.IntToStr(profileId) + " order by Id Desc limit 100 "

    var rs []models.Post
    base.DB.Select(&rs, sql)

    view:= models.PostsToPostsAndDetailes(rs)
    res:= ProfileRespnse{}
    res.Posts = view
    res.User = u
    c.SendJson(res)
    return nil
}



