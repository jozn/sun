package ctrl

import (
    "ms/sun/base"
    "ms/sun/models"
)

func FixAllCountsCounts(a *base.Action) base.AppErr {

    _fixPostCounts()
    _fixFolloingCounts()
    _fixFollowedCounts()

    return nil
}

func _fixPostCounts() {
    type row struct {
        UserId int
        Cnt int
    }

    ////// Posts
    var res []row

    q:= "SELECT UserId, COUNT(*) AS Cnt FROM post GROUP BY UserId"

    models.NewUser_Updater().PostsCount(0).Update(base.DB)

    base.DB.Select(&res,q)

    for _, row := range res {
        models.NewUser_Updater().PostsCount(row.Cnt).Id_EQ(row.UserId).Update(base.DB)
    }
}


func _fixFolloingCounts() {
    type row struct {
        UserId int
        Cnt int
    }

    ////// Posts
    var res []row

    q:= "SELECT UserId, COUNT(*) AS Cnt FROM following_list_member GROUP BY UserId"

    base.DB.Select(&res,q)

    models.NewUser_Updater().FollowingCount(0).Update(base.DB)

    for _, row := range res {
        models.NewUser_Updater().FollowingCount(row.Cnt).Id_EQ(row.UserId).Update(base.DB)
    }
}

func _fixFollowedCounts() {
    type row struct {
        UserId int
        Cnt int
    }

    var res []row

    q:= "SELECT FollowedUserId As UserId , COUNT(*) AS Cnt FROM following_list_member GROUP BY FollowedUserId"

    base.DB.Select(&res,q)

    //reset
    models.NewUser_Updater().FollowersCount(0).Update(base.DB)

    for _, row := range res {
        models.NewUser_Updater().FollowersCount(row.Cnt).Id_EQ(row.UserId).Update(base.DB)
    }
}

