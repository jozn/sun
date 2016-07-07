package actions

import (
	// "encoding/json"
	// "fmt"
	. "ms/sun/base"
	. "ms/sun/models"
	"strconv"
)

//todo: suport pagination
func GetCommentsAction(c *Action) {
	devPrintn("action comments")
	pid := (c.Req.Form["post_id"])
	e(pid)
	var comments []Comment
	DB.Select(&comments, "select * from comments where PostId = ? limit 50 ", pid[0])
	// DB.Select(&comments, "select * from comments  limit 50 ") //, pid[0])

	var commentsInline []CommentInlineInfo
	for _, cmt := range comments {
		cmtView := CommentInlineInfo{}
		cmtView.Comment = cmt
		cmtView.Sender = GetUserView(cmt.UserId)
		commentsInline = append(commentsInline, cmtView)
	}
	c.SendJson(commentsInline)
	// c.SendJson(comments)

}

func PostAddCommentAction(c *Action) {
	devPrintn("\naction شیی زخئئثدفس")
	c.MustUser()
	c.MustPost()
	uid := c.UserId()
	// print(uid)
	e(uid)

	pid := (c.Req.Form["post_id"])
	texts := (c.Req.Form["text"])
	print(pid)

	cmt := Comment{}
	cmt.UserId = uid
	// lik.PostId = strconv.Itoa(pid[0])
	pidint, err := strconv.Atoi(pid[0])
	noErr(err)
	if err != nil {
		return
	}
	cmt.PostId = pidint
	cmt.CreatedTimestamp = now()
	//todo format text for security and size ,parsing ,tages,...
	cmt.Text = texts[0]
	debug(texts)

	DbInsertStruct(&cmt, "comments")
	// c.SendText("ok")
	c.SendJson(cmt)

}
