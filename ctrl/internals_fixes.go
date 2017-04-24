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
		Cnt    int
	}

	////// Posts
	var res []row

	q := "SELECT UserId, COUNT(*) AS Cnt FROM post GROUP BY UserId"

	err := base.DB.Select(&res, q)
	if err != nil {
		return
	}

	models.NewUser_Updater().PostsCount(0).Update(base.DB)

	for _, row := range res {
		models.NewUser_Updater().PostsCount(row.Cnt).Id_Eq(row.UserId).Update(base.DB)
	}
}

func _fixFolloingCounts() {
	type row struct {
		UserId int
		Cnt    int
	}

	////// Posts
	var res []row

	q := "SELECT UserId, COUNT(*) AS Cnt FROM following_list_member GROUP BY UserId"

	err := base.DB.Select(&res, q)
	if err != nil {
		return
	}

	models.NewUser_Updater().FollowingCount(0).Update(base.DB)

	for _, row := range res {
		models.NewUser_Updater().FollowingCount(row.Cnt).Id_Eq(row.UserId).Update(base.DB)
	}
}

func _fixFollowedCounts() {
	type row struct {
		UserId int
		Cnt    int
	}

	var res []row

	q := "SELECT FollowedUserId As UserId , COUNT(*) AS Cnt FROM following_list_member GROUP BY FollowedUserId"

	err := base.DB.Select(&res, q)
	if err != nil {
		return
	}
	//reset
	models.NewUser_Updater().FollowersCount(0).Update(base.DB)

	for _, row := range res {
		models.NewUser_Updater().FollowersCount(row.Cnt).Id_Eq(row.UserId).Update(base.DB)
	}
}
