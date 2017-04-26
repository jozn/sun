package ctrl

import (
	"ms/sun/base"
	"ms/sun/models"
)

func RecommendPostsCtrl(c *base.Action) base.AppErr {
	p := UpdateSessionActivityIfUser(c)

	min := (p.Page - 1) * p.Limit

	if min < 0 {
		min = 0
	}
	max := min + p.Limit

	if max > len(models.TopPosts) {
		c.SendJson(nil)
		return nil
	}

	view := models.Views.PostsViewsForPostIds(models.TopPosts[min:max], c.UserId())
	c.SendJson(view)
	return nil
}

func RecommendUsersCtrl(c *base.Action) base.AppErr {
	p := MustBeUserAndUpdate(c)

	rowsIds, err := models.NewRecommendUser_Selector().Select_TargetId().UserId_Eq(c.UserId()).OrderBy_Id_Desc().
		Limit(p.Limit).Offset(p.GetOffset()).GetIntSlice(base.DB)
	if err != nil {
		return err
	}

	res := models.Views.GetListOfUserForFollowType(rowsIds, c.UserId())

	c.SendJson(res)

	models.Recommend_Scheduler_GenUserRecomendations(c.UserId())

	return nil
}

func RecommendTagsCtrl(c *base.Action) base.AppErr {
	c.SendJson(models.TopTags)
	return nil
}

func RecommendTagsWithPostsCtrl(c *base.Action) base.AppErr {

	rs := models.TopTagsWithPostsResult

	c.SendJson(rs)
	return nil
}
