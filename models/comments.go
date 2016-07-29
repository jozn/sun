package models

import (
    "ms/sun/helper"
    "ms/sun/base"
)

func AddNewComment(UserId, PostId int, Text string) Comment  {
    cmt := Comment{}
    cmt.UserId = UserId

    cmt.PostId = PostId
    cmt.CreatedTime = helper.TimeNow()
    //todo format text for security and size ,parsing ,tages,...
    cmt.Text = Text

    res,err:= base.DbInsertStruct(&cmt, "comments")
    if err == nil {
        id ,_ := res.LastInsertId()
        cmt.Id = int(id)
        QueryIncerPostCommentsCount(PostId,1)
    }
    return cmt
}

func RemoveComment(UserId, PostId, CommentId int) bool {
    q := "DELETE FROM commonts WHERE UserId = ? AND PostId = ? AND Id = ?"
    res,err:= base.DB.DB.Exec(q, UserId, PostId, CommentId)

    removed:=false
    if err == nil {
        cnt,_ :=res.RowsAffected()
        if cnt == 1 {
            QueryDecerPostCommentsCount(PostId,1)
            removed = true
        }
    }
    return removed
}
