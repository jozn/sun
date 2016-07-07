package actions

import (
	. "ms/sun/base"
	. "ms/sun/models"
)

func GetProfileAction(c *Action) {

	username := c.Req.Form["username"][0]
	user := GetUserByUsername(username)

	res := ProfileView{}
	res.User = user
	res.UserInfo = GetUserInfo(user.Id)

	var posts []Post
	DB.Select(&posts, " select * from post where UserId = ?", user.Id)

	res.Posts = GetPostAndDetailes(posts)
	c.SendJson(res)

}
