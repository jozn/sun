package actions

import (
	// "encoding/json"
	"fmt"
	// "strings"
	. "ms/sun/base"
	. "ms/sun/models"
	"strconv"
)

//todo: suport pagination + AmIFollowing
func GetLikesAction(c *Action) {
	devPrintn("\naction likes")
	// c.MustUser()
	// c.Req.ParseForm()
	// uid := c.UserId()
	// print(uid)
	// e(uid)

	pid := (c.Req.Form["post_id"])
	e(pid)
	// print(pid)
	// fmt.Printf("%v %T\n", pid, pid)
	// s := pid[0] + " " + pid[1] + " " + pid[2]
	var likes []Like
	DB.Select(&likes, "select * from likes where PostId = ? limit 50 ", pid[0])
	// DB.Select(&likes, "select * from likes  limit 50 ") //, pid[0])

	// var usersLiked []UserInlineView
	var usersLiked []UserInlineFollowView
	for _, lik := range likes {
		likeView := UserInlineFollowView{}
		likeView.UserInlineView = GetUserView(lik.UserId)
		if c.IsUser() {
			likeView.AmIFollowing = IsUserFollowing(c.UserId(), likeView.UserId)
		}
		// likeView.AmIFollow = true //IsUserFollowing(c.UserId(), likeView.UserId)
		// likeView.IsProfilePrivate = true
		usersLiked = append(usersLiked, likeView)
	}
	c.SendJson(usersLiked)
	// c.SendJson(likes)

}

func PostAddLikeAction(c *Action) {
	devPrintn("\naction likes")
	c.MustUser()
	c.MustPost()
	uid := c.UserId()
	// print(uid)
	e(uid)

	pid := (c.Req.Form["post_id"])
	print(pid)

	lik := Like{}
	lik.UserId = uid
	// lik.PostId = strconv.Itoa(pid[0])
	pidint, err := strconv.Atoi(pid[0])
	noErr(err)
	lik.PostId = pidint
	lik.CreatedTimestamp = now()

	DbInsertStruct(&lik, "likes")
	// c.SendText("ok")
	c.SendJson(lik)
	//	UserMajorInfoMe{}
}

func GetLikesPlayAction(c *Action) {
	devPrintn("\naction likes")
	// c.MustUser()
	c.Req.ParseForm()
	uid := c.UserId()
	// print(uid)
	e(uid)

	pid := (c.Req.Form["post_id"])
	print(pid)
	fmt.Printf("%v %T\n", pid, pid)
	s := pid[0] + " " + pid[1] + " " + pid[2]
	c.SendText(s)

}
