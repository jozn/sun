package models

import "ms/sun/base"

type _cntImpl int

var Counter _cntImpl

func (t _cntImpl)UpdateUserFollowingCounts (UserId int, cnt int) {
    NewUser_Updater().Id_EQ(UserId).FollowingCount_Increment(cnt).Update(base.DB)
    row,ok :=MemoryStore_User.GetMemRow(UserId)
    if ok{
        row.Row.FollowingCount += cnt
    }
}

func (t _cntImpl)UpdateUserFollowersCounts (UserId int, cnt int) {
    NewUser_Updater().Id_EQ(UserId).FollowersCount_Increment(cnt).Update(base.DB)
    row,ok :=MemoryStore_User.GetMemRow(UserId)
    if ok{
        row.Row.FollowersCount += cnt
    }
}

func (t _cntImpl)UpdateUserPostsCounts (UserId int, cnt int) {
    NewUser_Updater().Id_EQ(UserId).PostsCount_Increment(cnt).Update(base.DB)
    row,ok :=MemoryStore_User.GetMemRow(UserId)
    if ok{
        row.Row.PostsCount += cnt
    }
}

func (t _cntImpl)IncerPostCommentsCount (PostId, CountDiff int) {
    NewPost_Updater().Id_EQ(PostId).CommentsCount_Increment(CountDiff).Update(base.DB)
}

