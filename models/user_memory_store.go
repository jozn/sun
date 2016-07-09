package models

import (
    . "ms/sun/base"
    "sort"
    "sync"
    "ms/sun/ds"
)

type memoryStoreImpl struct {
    Map map[int]*userMemRow
    sync.RWMutex
}
//type memoryStoreImpl map[int]userMemRow

//+gen slice:"Where,Any"
var UserMemoryStore memoryStoreImpl

func init() {
    UserMemoryStore = memoryStoreImpl{}
    UserMemoryStore.Map = make(map[int]*userMemRow)
}

//+gen slice:"Where,Count,GroupBy[string]"
type userMemRow struct {
    UserTable
    SS                  string
    SessionId           [20]rune
    //Followings []int
    //Followings2 _followingTypeSorter //[]FollowingType
    //Followings3 TwoArr //[]FollowingType
    Followings          *ds.IntList
    FollowingsRequested *ds.IntList

    LikedPost           []int
    RecentPosts         []int
    RecentActivities    []int
}

////////////////////////// Loadings Of values from Dbs ///////////////////////////

func (db *memoryStoreImpl) ReloadAll(){
    db.Map = make(map[int]*userMemRow)
    db.ReloadAllUser()//must called first
    //go
    //db.ReloadAllFollowings()
    //db.ReloadAllFollowings2()
    //db.ReloadAllFollowings3()
    db.ReloadAllFollowings5()
    db.ReloadAllLikes()

}

func (db *memoryStoreImpl) ReloadAllUser(){
    var us []UserTable
    //debug(u)
    err := DB.Select(&us, "select * from user")
    //fmt.Println(us)
    //cashe.Set(key, u[0], 0)
    _ = err
    for _ ,u := range us {
        row := userMemRow{}
        row.UserTable = u
        db.RWMutex.Lock()
        db.Map[u.Id] = &row
        db.RWMutex.Unlock()
        //print(row.SS)
    }
}


func (db *memoryStoreImpl) ReloadAllFollowings5(){
    for _ ,r := range db.Map {
        r.Followings = ds.New()
        r.FollowingsRequested = ds.New()
    }

    var fm []FollowingListMember
    err := DB.Select(&fm, "select * from following_list_member")
    _ = err
    for _ ,f := range fm {
        _,ok :=db.Map[f.UserId]
        if ok {
            if f.FollowType == 1{
                db.Map[f.UserId].Followings.AddAndSort(f.FollowedUserId)
            }else if f.FollowType == 2{
                db.Map[f.UserId].FollowingsRequested.AddAndSort(f.FollowedUserId)
            }
            //db.Map[f.UserId].Followings3 = append(db.Map[f.UserId].Followings3 , [2]int{f.FollowedUserId,f.FollowType} )// = 1//f.FollowedUserId
            //db.Map[f.UserId].Followings3[1] = f.FollowType
        }
    }

    for _ ,r := range db.Map {
        //if r.Followings5 != nil{
        //    r.Followings5.SortDesc()
        _ = r
            //sort.Sort(r.Followings3)
        //}
    }
}

func (db *memoryStoreImpl) ReloadAllLikes(){
    for _ ,r := range db.Map {
        r.LikedPost = make([]int,0)
    }

    var ls []Like
    err := DB.Select(&ls, "select * from likes")
    _ = err
    for _ , l := range ls {
        _,ok :=db.Map[l.UserId]
        if ok {
            db.Map[l.UserId].LikedPost = append(db.Map[l.UserId].LikedPost, l.PostId)
        }
    }

    for _ ,r := range db.Map {
        if r.LikedPost != nil{
            sort.Ints(r.LikedPost)
        }
    }
}

////////////////////////////// ////////////////////////////////////////
func (db *memoryStoreImpl) GetAllAsArray() []*userMemRow {
    rows:= make([]*userMemRow,0)

    for _ ,r := range db.Map {
        rows = append(rows,r)
    }
    return rows
}

//////////////////////// Followings ///////////////////////////////////
func (db *memoryStoreImpl) GetFollowingTypeForUsers(UserId, ReqFollowedUserId int) int {
    s , ok :=  db.Map[UserId]
    ftype:= 0
    if ok {
        if s.Followings.BinaryContains(ReqFollowedUserId){
            ftype = 1
        }else if s.FollowingsRequested.BinaryContains(ReqFollowedUserId) {
            ftype = 2
        }
    }
    return ftype
}

func (db *memoryStoreImpl) AddFollow(UserId, FollowedUserId int) {
    s , ok :=  db.Map[UserId]
    if ok {
        s.Followings.AddAndSort(FollowedUserId)
        QueryInsertNewFollowing(UserId,FollowedUserId,1)
    }
}

func (db *memoryStoreImpl) RemoveFollow(UserId, FollowedUserId int) {
    s , ok :=  db.Map[UserId]
    if ok {
        s.Followings.RemoveAndSort(FollowedUserId)
        QueryInsertNewFollowing(UserId,FollowedUserId,0)
    }
}

////////////////// Posts Functionalities /////////////////////
func (db *memoryStoreImpl) AddPostLike(UserId, PostId int) {
    s , ok :=  db.Map[UserId]
    if ok {
        _ =s
    }
}

func (db *memoryStoreImpl) RemovePostLike(UserId, PostId int) {
    s , ok :=  db.Map[UserId]
    if ok {
        _ =s
    }
}



//////////////////// End ///////////////////////////////

//all below
//deprecraed
type FollowingType struct {
    FollowedUserId int
    TypeId int
}

type _followingTypeSorter []FollowingType

func (ft _followingTypeSorter) Len() int  {return len(ft)}
func (p _followingTypeSorter) Less(i, j int) bool { return p[i].FollowedUserId < p[j].FollowedUserId }
func (p _followingTypeSorter) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type TwoArr [][2]int
func (ft TwoArr) Len() int  {return len(ft)}
func (p TwoArr) Less(i, j int) bool { return p[i][0] < p[j][0]}
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


