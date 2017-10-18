package ctrl

import (
	// "encoding/json"
	// "fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"ms/sun/models/x"
	"strings"
)

const COMMENTS_LIST_PAGE_LIMIT = 30

func GetCommentsAction(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	pid := c.GetParamInt("post_id", 0)
	page := c.GetParamInt("page", 0)
	limit := c.GetParamInt("limit", COMMENTS_LIST_PAGE_LIMIT)
	last := c.GetParamInt("last", 0)

	selector := x.NewComment_Selector().PostId_Eq(pid).OrderBy_Id_Desc().Limit(limit)

	if last > 0 {
		selector.Id_LE(last)
	} else if page > 0 {
		selector.Offset((page - 1) * limit)
	}

	comments, err := selector.GetRows2(base.DB)

	if err != nil {
		helper.DebugErr(err)
		return err
	}

	var commentsInline []models.CommentInlineInfoView
	for _, cmt := range comments {
		cmtView := models.CommentInlineInfoView{}
		cmtView.Comments = cmt
		cmtView.Sender, _ = models.Views.GetUserInlineView(cmt.UserId)
		commentsInline = append(commentsInline, cmtView)
	}
	c.SendJson(commentsInline)
	return nil
}

func PostAddCommentAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	pid := c.GetParamInt("post_id", 0)
	texts := c.Req.Form.Get("text")
	texts = strings.Trim(texts, " ")
	if len(texts) == 0 || pid < 1 {
		return nil
	}

	// c.SendText("ok")
	cmt := models.Comment_Add(c.UserId(), pid, texts)
	c.SendJson(cmt)
	return nil
}

func RemoveCommentAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	pid := c.GetParamInt("post_id", -1)
	comment_id := c.GetParamInt("comment_id", -1)

	boolean := models.Comment_Delete(c.UserId(), pid, comment_id)
	c.SendJson(boolean)
	return nil
}
