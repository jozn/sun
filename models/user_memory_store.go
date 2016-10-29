package models

import (
	. "ms/sun/base"
	"ms/sun/ds"
	"ms/sun/helper"
	"sync"
)

type mapMemoryStoreImpl struct {
	Map map[int]*userMemRow
	sync.RWMutex
}

//type memoryStoreImpl map[int]userMemRow

//+gen slice:"Where,Any"
var UserMemoryStore mapMemoryStoreImpl

func init() {
	UserMemoryStore = mapMemoryStoreImpl{}
	UserMemoryStore.Map = make(map[int]*userMemRow)
	//UserMemoryStore.ReloadAll()
}

//+gen slice:"Where,Count,GroupBy[string]"
type userMemRow struct {
	UserTable
	SessionId           [20]rune
	Followings          *ds.IntList
	FollowingsRequested *ds.IntList

	LikedPost *ds.IntList
	//LikedPost2          []int
	RecentPosts      []int
	RecentActivities []int
}

////////////////////////// Loadings Of values from Dbs ///////////////////////////

func (db *mapMemoryStoreImpl) ReloadAll() {
	db.Map = make(map[int]*userMemRow)
	db.ReloadAllUser() //must called first
	//go
	db.ReloadAllFollowings()
	db.ReloadAllLikes()

}

func (db *mapMemoryStoreImpl) ReloadAllUser() {
	var us []UserTable
	//debug(u)
	err := DB.Select(&us, "select * from user")
	//fmt.Println(us)
	//cashe.Set(key, u[0], 0)
	if err != nil {
		helper.DebugErr(err)
	}
	_ = err
	for _, u := range us {
		row := userMemRow{}
		row.UserTable = u
		db.RWMutex.Lock()
		db.Map[u.Id] = &row
		db.RWMutex.Unlock()
	}
}

func (db *mapMemoryStoreImpl) ReloadAllFollowings() {
	for _, r := range db.Map {
		r.Followings = ds.New()
		r.FollowingsRequested = ds.New()
	}

	var fm []FollowingListMember
	err := DB.Select(&fm, "select * from following_list_member")
	_ = err
	for _, f := range fm {
		_, ok := db.Map[f.UserId]
		if ok {
			if f.FollowType == 1 {
				db.Map[f.UserId].Followings.AddAndSort(f.FollowedUserId)
			} else if f.FollowType == 2 {
				db.Map[f.UserId].FollowingsRequested.AddAndSort(f.FollowedUserId)
			}
			//db.Map[f.UserId].Followings3 = append(db.Map[f.UserId].Followings3 , [2]int{f.FollowedUserId,f.FollowType} )// = 1//f.FollowedUserId
			//db.Map[f.UserId].Followings3[1] = f.FollowType
		}
	}

	for _, r := range db.Map {
		//if r.Followings5 != nil{
		//    r.Followings5.SortDesc()
		_ = r
		//sort.Sort(r.Followings3)
		//}
	}
}

func (db *mapMemoryStoreImpl) ReloadAllLikes() {
	for _, r := range db.Map {
		r.LikedPost = ds.New()
	}

	var ls []Like
	err := DB.Select(&ls, "select * from likes")
	_ = err
	for _, l := range ls {
		_, ok := db.Map[l.UserId]
		if ok {
			db.Map[l.UserId].LikedPost.Add(l.PostId) // = append(db.Map[l.UserId].LikedPost2, l.PostId)
		}
	}

	for _, r := range db.Map {
		if r.LikedPost != nil {
			r.LikedPost.SortDesc()
		}
	}
}

////////////////////////////// ////////////////////////////////////////
func (db *mapMemoryStoreImpl) GetForUser(UserId int) *userMemRow {
	db.RLock()
	u, ok := db.Map[UserId]
	db.RUnlock()
	if ok {
		return u
	}
	return nil
}

func (db *mapMemoryStoreImpl) GetUserTableForUser(UserId int) *UserTable {
	row := db.GetForUser(UserId)
	if row != nil {
		return &row.UserTable
	}
	return nil
}

func (db *mapMemoryStoreImpl) GetAllAsArray() []*userMemRow {
	rows := make([]*userMemRow, 0)

	for _, r := range db.Map {
		rows = append(rows, r)
	}
	return rows
}
func (db *mapMemoryStoreImpl) ReloadUser(UserId int) {
	//_ , ok :=  db.Map[UserId]
	u := db.GetForUser(UserId)
	if u != nil {
		var u2 UserTable
		err := DB.Get(&u2, "select * from user where Id = ? ", UserId)
		if err != nil {
			//db.Lock()
			//db.Map[UserId].UserTable = u
			u.UserTable = u2
			//db.Unlock()
		}
	}
}

//////////////////////// Likes - Post /////////////////////////////////

func (db *mapMemoryStoreImpl) AddPostLike(UserId, PostId int) {
	s, ok := db.Map[UserId]
	if ok {
		if !s.LikedPost.BinaryContains(PostId) { //don't duplicate
			s.LikedPost.AddAndSort(PostId)
		}
		//QueryIncerPostLikesCount(PostId, 1)
		//QueryReomePostLike(UserId,PostId)
		/*err := QueryAddPostLike(UserId, PostId)
		if err == nil {
			QueryIncerPostLikesCount(PostId, 1)
		}*/
	}
}

func (db *mapMemoryStoreImpl) AmILikePost(UserId, PostId int) bool {
	s := db.GetForUser(UserId)
	if s != nil {
		if s.LikedPost.BinaryContains(PostId) {
			return true
		}
	}
	return false
}

func (db *mapMemoryStoreImpl) RemovePostLike(UserId, PostId int) {
	s, ok := db.Map[UserId]
	if ok {
		s.LikedPost.RemoveAndSort(PostId)
		QueryDecerPostLikesCount(PostId, 1)
		/*err := QueryReomePostLike(UserId, PostId)
		if err == nil {
			QueryDecerPostLikesCount(PostId, 1)
		}*/
	}
}

//////////////////////// Followings ///////////////////////////////////
func (db *mapMemoryStoreImpl) GetFollowingTypeForUsers(UserId, ReqFollowedUserId int) int {
	s, ok := db.Map[UserId]
	ftype := 0
	if ok {
		//debug("xx: ",s.Followings)
		if s.Followings.BinaryContains(ReqFollowedUserId) {
			//if s.Followings.Contains(ReqFollowedUserId){
			ftype = 1
		} else if s.FollowingsRequested.BinaryContains(ReqFollowedUserId) {
			ftype = 2
		}
	}
	return ftype
}

func (db *mapMemoryStoreImpl) AddFollow(UserId, FollowedUserId int) {
	s, ok := db.Map[UserId]
	if ok {
		s.Followings.AddAndSort(FollowedUserId)
		//QueryInsertNewFollowing(UserId,FollowedUserId,1)
	}
}

func (db *mapMemoryStoreImpl) RemoveFollow(UserId, FollowedUserId int) {
	s, ok := db.Map[UserId]
	if ok {
		s.Followings.RemoveAndSort(FollowedUserId)
		//QueryInsertNewFollowing(UserId,FollowedUserId,0)
	}
}

func (db *mapMemoryStoreImpl) GetAllFollowingsListOfUser(UserId int) *ds.IntList {
	s, ok := db.Map[UserId]
	if ok {
		return s.Followings
	}
	return ds.New()
}

//////////////// User Actions Counts ////////////////
func (db *mapMemoryStoreImpl) UpdateUserFollowingCounts(UserId int, cnt int) {
	user := db.GetForUser(UserId)
	if user != nil {
		user.UserCounts.FollowingCount += cnt
		QueryUpdateUserActionCounts(UserId, cnt, "FollowingCount")
	}
}

func (db *mapMemoryStoreImpl) UpdateUserFollowersCounts(UserId int, cnt int) {
	user := db.GetForUser(UserId)
	if user != nil {
		user.UserCounts.FollowersCount += cnt
		QueryUpdateUserActionCounts(UserId, cnt, "FollowersCount")
	}
}

func (db *mapMemoryStoreImpl) UpdateUserPostsCounts(UserId int, cnt int) {
	user := db.GetForUser(UserId)
	if user != nil {
		user.UserCounts.PostsCount += cnt
		QueryUpdateUserActionCounts(UserId, cnt, "PostsCount")
	}
}

func (db *mapMemoryStoreImpl) UpdateUserMediasCounts(UserId int, cnt int) {
	user := db.GetForUser(UserId)
	if user != nil {
		user.UserCounts.FollowingCount += cnt
		QueryUpdateUserActionCounts(UserId, cnt, "MediaCount")
	}
}

func (db *mapMemoryStoreImpl) UpdateUserLikesCounts(UserId int, cnt int) {
	user := db.GetForUser(UserId)
	if user != nil {
		user.UserCounts.FollowingCount += cnt
		QueryUpdateUserActionCounts(UserId, cnt, "LikesCount")
	}
}

//////////////////////// Session ///////////////////////////////
func (db *mapMemoryStoreImpl) TryLogin(PhoneEmail, PassWord string) {

}

func (db *mapMemoryStoreImpl) IsUserSession(UserId int, SessionUuid string) bool {
	if UserId < 1 || SessionUuid == "" {
		return false
	}

	user := db.GetForUser(UserId)
	if user != nil {
		if user.SessionUuid == SessionUuid {
			///Temp
			GenrateRecommends(UserId)
			////
			return true
		}
	}
	return false
}

func (db *mapMemoryStoreImpl) IsUserSessionAndUpdateActivity(UserId int, SessionUuid string) bool {
	user := db.GetForUser(UserId)
	is := db.IsUserSession(UserId, SessionUuid)
	if is {
		user.LastActivityTime = helper.TimeNow()
		PeriodaclyUpdateLastActivityOfUser(UserId)
	}
	return is
}

/*
func (db *memoryStoreImpl) UpdateUserCommentsCounts(UserId int, cnt int  ) {
    user  :=  db.GetForUser(UserId)
    if user != nil {
        user.UserCounts.FollowingCount += 1
        QueryUpdateUserActionCounts(UserId,cnt,"LikesCount")
    }
}
*/

//////////////////// End ///////////////////////////////

//all below
//deprecraed
type FollowingType struct {
	FollowedUserId int
	TypeId         int
}

type _followingTypeSorter []FollowingType

func (ft _followingTypeSorter) Len() int          { return len(ft) }
func (p _followingTypeSorter) Less(i, j int) bool { return p[i].FollowedUserId < p[j].FollowedUserId }
func (p _followingTypeSorter) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type TwoArr [][2]int

func (ft TwoArr) Len() int          { return len(ft) }
func (p TwoArr) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p TwoArr) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

/*
func (db *memoryStoreImpl) ReloadAllFollowings(){
    for _ ,r := range db.Map {
        r.Followings = make([]int,0)
    }

    var fm []FollowingListMember
    err := DB.Select(&fm, "select * from following_list_member")
    _ = err
    for _ ,f := range fm {
        _,ok :=db.Map[f.UserId]
        if ok {
            db.Map[f.UserId].Followings = append(db.Map[f.UserId].Followings, f.FollowedUserId)
        }
    }

    for _ ,r := range db.Map {
        if r.Followings != nil{
            sort.Ints(r.Followings)
        }
    }
}

func (db *memoryStoreImpl) ReloadAllFollowings2(){
    for _ ,r := range db.Map {
        r.Followings2 = make(_followingTypeSorter,0)
    }

    var fm []FollowingListMember
    err := DB.Select(&fm, "select * from following_list_member")
    _ = err
    for _ ,f := range fm {
        _,ok :=db.Map[f.UserId]
        if ok {
            db.Map[f.UserId].Followings2 = append(db.Map[f.UserId].Followings2, FollowingType{f.FollowedUserId, f.FollowType})
        }
    }

    for _ ,r := range db.Map {
        if r.Followings2 != nil{
            sort.Sort(r.Followings2)
        }
    }
}

func (db *memoryStoreImpl) ReloadAllFollowings3(){
    for _ ,r := range db.Map {
        r.Followings3 = TwoArr{}
    }

    var fm []FollowingListMember
    err := DB.Select(&fm, "select * from following_list_member")
    _ = err
    for _ ,f := range fm {
        _,ok :=db.Map[f.UserId]
        if ok {
            db.Map[f.UserId].Followings3 = append(db.Map[f.UserId].Followings3 , [2]int{f.FollowedUserId,f.FollowType} )// = 1//f.FollowedUserId
            //db.Map[f.UserId].Followings3[1] = f.FollowType
        }
    }

    for _ ,r := range db.Map {
        if r.Followings3 != nil{
            sort.Sort(r.Followings3)
        }
    }
}
*/
