package ctrl

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"strings"
)

func SearchTagsCtrl(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	q := c.Req.FormValue("q")

	q = q + "%"

	var tags []models.Tag

	if strings.TrimSpace(q) != "" {
		base.DB.Select(&tags, "select * from tags where `Name` like ? limit 30 ", q)
	}

	c.SendJson(tags)
	return nil
}

func SearchUsersCtrl(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	q := c.Req.FormValue("q")

	//q =  q +"%"
	qs := strings.Split(q, " ")

	if len(qs) > 0 {
		selector := models.NewUser_Selector().Select_Id().Or()

		for _, q0 := range qs {
			if q0 == "" {
				continue
			}
			q0 = q0 + "%"
			selector.
				UserName_Like(q0).
				FirstName_Like(q0).
				LastName_Like(q0)
		}

		UsersIds, err := selector.Limit(30).GetIntSlice(base.DB)

		if err != nil {
			return err
		}

		usersFollow := models.Views.GetListOfUserForFollowType(UsersIds, c.UserId())
		c.SendJson(usersFollow)
		return nil
	}

	c.SendJson(nil)
	return nil
}

func SearchClickedCtrl(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	q := c.Req.FormValue("q")
	target_id := c.GetParamInt("target_id", 0)
	click_type := c.Req.FormValue("click_type")

	if q == "" {
		return nil
	}

	click_type_id := 1 //search
	if click_type == "user" {
		click_type_id = 2
	} else if click_type == "tag" {
		click_type_id = 3
	}

	row := models.SearchClicked{
		Query:     q,
		TargetId:  target_id,
		ClickType: click_type_id,
		UserId:    c.UserId(),
		CreatedAt: helper.TimeNow(),
	}

	row.Save(base.DB)

	return nil
}

func NoftificationsCtrl2(c *base.Action) base.AppErr {
	//MustBeUserAndUpdate(c)

	_ = (models.Notification_GetLastsViews(c.UserId()))
	c.SendJson("asdsa")
	return nil
}

func SearchUsersCtrl_old(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	q := c.Req.FormValue("q")

	q = "%" + q + "%"

	var UsersIds []int

	if strings.TrimSpace(q) != "" {
		base.DB.Select(&UsersIds, "SELECT Id FROM USER WHERE UserName OR FirstName OR LastName LIKE ? LIMIT 20  ", q)
	}

	//usersFollow := models.UsersToInlineFollowView(UsersIds, c.UserId())
	//c.SendJson(usersFollow)
	return nil
}

func SearchTagsCtrl_old(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	q := c.Req.FormValue("q")

	q = "%" + q + "%"

	var tags []models.Tag

	if strings.TrimSpace(q) != "" {
		base.DB.Select(&tags, "select * from tags where `Name` like ? limit 20 ", q)
	}

	c.SendJson(tags)
	return nil
}

type SearchResultView struct {
	Tags  []models.Tag
	Users []models.UserInlineFollowView
}

//dep
func SearchCtrl(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	q := c.Req.FormValue("q")

	q = "%" + q + "%"

	//var users []models.User
	var UsersIds []int
	var tags []models.Tag

	if strings.TrimSpace(q) != "" {
		base.DB.Select(&UsersIds, "SELECT Id FROM USER WHERE UserName OR FirstName OR LastName LIKE ? LIMIT 20  ", q)
		base.DB.Select(&tags, "select * from tags where `Name` like ? limit 20 ", q)
	}

	//usersFollow := models.UsersToInlineFollowView(UsersIds, c.UserId())
	//res := SearchResultView{Tags: tags, Users: usersFollow}
	//_ = res
	c.SendJson("as")
	return nil
}
