package ctrl

import (
	"fmt"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"ms/sun/models/x"
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

	rowsIds, err := x.NewRecommendUser_Selector().Select_TargetId().UserId_Eq(c.UserId()).OrderBy_Id_Desc().
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
	p := UpdateSessionActivityIfUser(c)
	/*min := (p.Page - 1 ) * p.Limit
	  if min < 0{
	      min = 0
	  }
	  max := min + p.Limit
	  if len(models.TopTags) < max {
	      if len(models.TopTags) >= min{//err
	          c.SendJson(nil)
	          return nil
	      }
	      max = len(models.TopTags)
	  }*/

	r, ok := helper.MaxPageLimit(len(models.TopPosts), p.Page, p.Limit)

	if !ok {
		c.SendJson(nil)
		return nil
	}

	c.SendJson(models.TopTags[r.Start:r.End])
	return nil
}

func RecommendTagsWithPostsCtrl(c *base.Action) base.AppErr {
	p := UpdateSessionActivityIfUser(c)

	r, ok := helper.MaxPageLimit(len(models.TopTagsWithPostsResult), p.Page, p.Limit)
	fmt.Println(r, ok)

	if !ok {
		c.SendJson(nil)
		return nil
	}

	c.SendJson(models.TopTagsWithPostsResult[r.Start:r.End])
	return nil
}
