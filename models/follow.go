package models

import (
    "ms/sun/base"
    "ms/sun/helper"
)

type FollowingList struct {
    Id          int
    UserId      int
    ListType    int
    Name        string
    Count       int
    IsAuto      int
    IsPimiry    int
    CreatedTime int
}

//following_list_member
type FollowingListMember struct {
    Id             int
    ListId         int
    UserId         int
    FollowedUserId int
    FollowType     int
    UpdatedTimeMs  int
}

//table: following_list_member_history
type FollowingListMemberHistory struct {
    FollowingListMember
    FollowId int
}


func Follow(UserId, FollowedPeerUserId int ) {
    if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
        return
    }

    follow,err :=QueryInsertNewFollowing(UserId, FollowedPeerUserId,1)
    if err == nil {
        UserMemoryStore.AddFollow(UserId, FollowedPeerUserId)
        UserMemoryStore.UpdateUserFollowingCounts(UserId, 1)
        UserMemoryStore.UpdateUserFollowersCounts(FollowedPeerUserId,1)

        fh :=FollowingListMemberHistory{}
        fh.FollowingListMember = *follow
        fh.FollowType = 1
        fh.FollowId = 0
        fh.InsertToDb()
    }
}

func UnFollow(UserId, FollowedPeerUserId int )  {
    if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
        return
    }

    res,err := base.DbExecute("DELETE FROM following_list_member WHERE UserId = ? AND FollowedUserId = ?",UserId,FollowedPeerUserId)
    if err == nil {
        n,err := res.RowsAffected()
        if err == nil && n >0 {
            UserMemoryStore.UpdateUserFollowingCounts(UserId,-1)
            UserMemoryStore.UpdateUserFollowersCounts(FollowedPeerUserId , -1)

            fh :=FollowingListMemberHistory{}
            fh.UserId = UserId
            fh.FollowedUserId = FollowedPeerUserId
            fh.UpdatedTimeMs = helper.TimeNowMs()
            fh.FollowType = 2
            fh.InsertToDb()
        }
    }
}

func (fh *FollowingListMemberHistory) InsertToDb ()  {
    base.DbInsertStruct(fh,"following_list_member_history")
}
