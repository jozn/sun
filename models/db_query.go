package models

import (
"ms/sun/helper"
"strings"
"ms/sun/base"
)

func QueryInsertNewFollowing(UserId, FollowedUserId, FollowType int, )  {
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

