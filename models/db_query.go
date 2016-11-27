package models

import (
	"math"
	"ms/sun/base"
	"ms/sun/helper"
)

///ALL Deprecated

///////////////////// Followings /////////////////////////////
/*
func QueryInsertNewFollowing(UserId, FollowedUserId, FollowType int) (*FollowingListMember, error) {
	flm := FollowingListMember{}
	flm.FollowedUserId = FollowedUserId
	flm.UserId = UserId
	flm.FollowType = FollowType //0: not following any more ; 1 following : 2 requested(for private profiles)
	flm.ListId = UserId
	flm.UpdatedTimeMs = helper.TimeNowMs()
	keys, values := helper.StructToFiledsRejectsEscape(&flm, "Id")
	q := "Insert into following_list_member (" + strings.Join(keys, ",") + ") values (" + strings.Join(values, ",") + ")"

	_, err := base.DB.Exec(q)
	if err != nil {
		helper.DebugPrintln(err)
	}
	return &flm, err
}
*/

/*
/////////////////////// Likes ////////////////////////
//dep
func QueryAddPostLike(UserId, PostId int) error {
	l := Like{}
	l.UserId = UserId
	l.PostId = PostId
	l.CreatedTime = helper.TimeNow()

	keys, values := helper.StructToFiledsRejectsEscape(&l, "Id")
	q := "replace into likes (" + strings.Join(keys, ",") + ") values (" + strings.Join(values, ",") + ")"

	_, err := base.DB.Exec(q)
	if err != nil {
		devPrintn(err)
	}
	return err
}

//dep
func QueryReomePostLike(UserId, PostId int) error {
	q := "DELETE FROM likes WHERE UserId = ? AND PostId = ?"
	_, err := base.DB.Exec(q, UserId, PostId)
	if err != nil {
		devPrintn(err)
	}
	return err
}

func QueryIncerPostLikesCount(PostId, CountDiff int) {
	cnt := helper.IntToStr(CountDiff)
	q := "UPDATE post SET LikesCount = LikesCount + " + cnt + " WHERE Id = ?"
	_, err := base.DB.Exec(q, PostId)
	if err != nil {
		devPrintn(err)
	}
}

func QueryDecerPostLikesCount(PostId, CountDiff int) {
	cnt := helper.IntToStr(CountDiff)
	q := "UPDATE post SET LikesCount = LikesCount - " + cnt + " WHERE Id = ?"
	_, err := base.DB.Exec(q, PostId)
	if err != nil {
		devPrintn(err)
	}
}
*/

//////////////////// Comments ///////////////////////
/*

func QueryIncerPostCommentsCount(PostId, CountDiff int) {
	cnt := helper.IntToStr(CountDiff)
	q := "UPDATE post SET CommentsCount = CommentsCount + " + cnt + " WHERE Id = ?"
	_, err := base.DB.Exec(q, PostId)
	if err != nil {
		devPrintn(err)
	}
}

func QueryDecerPostCommentsCount(PostId, CountDiff int) {
	cnt := helper.IntToStr(CountDiff)
	q := "UPDATE post SET CommentsCount = CommentsCount - " + cnt + " WHERE Id = ?"
	_, err := base.DB.Exec(q, PostId)
	if err != nil {
		devPrintn(err)
	}
}
*/

///////////////////// User ///////////////////////////
func QueryUpdateUserActionCounts(UserId, CountDiff int, column string) {
	cnt := helper.IntToStr(int(math.Abs(float64(CountDiff))))
	q := ""
	if CountDiff >= 0 {
		q = "UPDATE user SET " + column + " = " + column + " + " + cnt + " WHERE Id = ?"
	} else {
		q = "UPDATE user SET " + column + " = " + column + " - " + cnt + " WHERE Id = ?"
	}
	_, err := base.DbExecute(q, UserId)
	if err != nil {
		devPrintn(err)
	}
}

func QueryUpdateSessionLastActivitiesOfUsers(UserIds []int) {
	if len(UserIds) == 0 {
		return
	}

	q := "UPDATE user SET LastActivityTime = " + helper.IntToStr(helper.TimeNow()) + " WHERE Id in (" + helper.IntsToSqlIn(UserIds) + ")  "
	_, err := base.DB.Exec(q)
	if err != nil {
		devPrintn(err)
	}
	helper.DebugPrintln(q)
}

/*func QueryRemoveFollowing(UserId, FollowedUserId int)  {
    flm := FollowingListMember{}
    flm.FollowedUserId = FollowedUserId
    flm.UserId = UserId
    flm.FollowType = 1
    flm.ListId = UserId
    flm.UpdatedTimestampMs = helper.TimeNowMs()
    keys,values :=helper.StructToFiledsRejectsEscape(&flm,"Id")
    q := "DELETE FROM following_list_member  "
    _,err:=base.DB.Exec(q)
    if err {
        devPrintn(err)
    }
}*/
