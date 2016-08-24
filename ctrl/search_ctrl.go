package ctrl

import (
    "ms/sun/base"
    "strings"
    "ms/sun/models"
)

type SearchResultView struct  {
    Tags []models.Tag
    Users []models.UserInlineFollowView
}

func SearchCtrl(c *base.Action) base.AppErr  {
    UpdateSessionActivityIfUser(c)

    q := c.Req.FormValue("q")

    q = "%"+ q + "%"

    //var users []models.User
    var UsersIds []int
    var tags []models.Tag

    if strings.TrimSpace(q) != "" {
        base.DB.Select(&UsersIds, "SELECT Id FROM USER WHERE UserName OR FirstName OR LastName LIKE ? LIMIT 20  ", q)
        base.DB.Select(&tags, "select * from tags where `Name` like ? limit 20 ", q)
    }

    usersFollow := models.UsersToInlineFollowView(UsersIds, c.UserId())
    res := SearchResultView{Tags: tags, Users: usersFollow}
    c.SendJson(res)
    return nil
}
