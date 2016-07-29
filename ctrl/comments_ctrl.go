package ctrl

import (
	// "encoding/json"
	// "fmt"
	"ms/sun/base"
	"ms/sun/models"
	"ms/sun/helper"
	"strings"
)

const COMMENTS_LIST_PAGE_LIMIT = 30

//todo: suport pagination
func GetCommentsAction(c *base.Action) base.AppErr {
    //MustBeUser(c)
	//devPrintn("action comments")
    //time.Sleep(time.Second*1)
	pidStr := c.Req.Form.Get("post_id")
	pid := helper.StrToInt(pidStr,0)

	page:= c.GetPage()
	offset := page * COMMENTS_LIST_PAGE_LIMIT

	var comments []models.Comment
	err:= base.DB.Select(&comments, "select * from comments where PostId = ? order by Id DESC limit ? OFFSET ? ", pid,COMMENTS_LIST_PAGE_LIMIT,offset)
	// DB.Select(&comments, "select * from comments  limit 50 ") //, pid[0])

	if err != nil {
		helper.DebugErr(err)
		return nil
	}

	var commentsInline []models.CommentInlineInfo
	for _, cmt := range comments {
		cmtView := models.CommentInlineInfo{}
		cmtView.Comment = cmt
		cmtView.Sender = models.GetUserView(cmt.UserId)
		commentsInline = append(commentsInline, cmtView)
	}
	c.SendJson(commentsInline)
	// c.SendJson(comments)
	return nil
}


func PostAddCommentAction(c *base.Action) base.AppErr {
    MustBeUserAndUpdate(c)

	pid := c.GetParamInt("post_id",0)
	texts := c.Req.Form.Get("text")
	texts = strings.Trim(texts," ")
	if len(texts) == 0{
		return nil
	}

	/*cmt := models.Comment{}
	cmt.UserId = c.UserId()

	cmt.PostId = pid
	cmt.CreatedTime = helper.TimeNow()
	//todo format text for security and size ,parsing ,tages,...
	cmt.Text = texts

	base.DbInsertStruct(&cmt, "comments")
    models.QueryIncerPostCommentsCount(pid,1)*/
	// c.SendText("ok")
    cmt := models.AddNewComment(c.UserId(),pid,texts)
	c.SendJson(cmt)
	return nil
}

func PostRemoveCommentAction(c *base.Action) base.AppErr {
    MustBeUserAndUpdate(c)

    pid := c.GetParamInt("post_id",0)
    comment_id := c.GetParamInt("comment_id",0)

    boolean := models.RemoveComment(c.UserId(),pid,comment_id)
    c.SendJson(boolean)
    return nil
}

/////////////////////////////////////////////////
