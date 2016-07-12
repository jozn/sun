package models

import (
"ms/sun/helper"
"strings"
"ms/sun/base"
)

///////////////////// Followings /////////////////////////////
func QueryInsertNewFollowing(UserId, FollowedUserId, FollowType int )  {
    flm := FollowingListMember{}
    flm.FollowedUserId = FollowedUserId
    flm.UserId = UserId
    flm.FollowType = FollowType //0: not following any more ; 1 following : 2 requested(for private profiles)
    flm.ListId = UserId
    flm.UpdatedTimestampMs = helper.TimeNowMs()
    keys,values :=helper.StructToFiledsRejectsEscape(&flm,"Id")
    q := "replace into following_list_member ("+strings.Join(keys,",") +") values (" +strings.Join(values,",") +")"

    _,err:=base.DB.Exec(q)
    if err !=nil {
        devPrintn(err)
    }
}

/////////////////////// Likes ////////////////////////
func QueryAddPostLike(UserId, PostId int )  {
    l := Like{}
    l.UserId = UserId
    l.PostId = PostId
    l.CreatedTimestamp = helper.TimeNow()

    keys,values :=helper.StructToFiledsRejectsEscape(&l,"Id")
    q := "replace into likes ("+strings.Join(keys,",") +") values (" +strings.Join(values,",") +")"

    _,err:=base.DB.Exec(q)
    if err !=nil {
        devPrintn(err)
    }
}

func QueryReomePostLike(UserId, PostId int )  {
    q := "DELETE FROM likes WHERE UserId = ? AND PostId = ?"
    _,err:=base.DB.Exec(q, UserId ,PostId)
    if err !=nil {
        devPrintn(err)
    }
}

///////////////////// User ///////////////////////////
func QueryUpdateUserActionCounts(UserId,  CountDiff int, column string )  {
    cnt := helper.IntToStr(CountDiff)
    q := "UPDATE user SET "+ column + " = " + column + " + " + cnt  + " WHERE Id = ?"
    _,err:=base.DB.Exec(q, UserId )
    if err !=nil {
        devPrintn(err)
    }
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

