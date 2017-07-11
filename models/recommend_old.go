package models

import (
	"ms/sun/base"
	"ms/sun/helper"
	"strings"
)

var TopUserIds []int = []int{}

func ReloadTopUserIds() {
	TopUserIds = []int{}
	base.DB.Select(&TopUserIds, "select * from user order by FollowersCount limit 500")
}

/////////////////////////////////////////////////////////////

/*
type RecommendUser struct {
	Id          int
	UserId      int
	TargetId    int
	CreatedTime int
	Weight      float64
	// xo fields
	_exists, _deleted bool
}
*/

func GenrateRecommends(ForUserId int) {
	//NewRecommendUser_Deleter()
	go func() {
		defer helper.JustRecover()

		helper.DebugPrintln("GenrateRecommends for user: ", ForUserId)
		GenrateRecommendUser(ForUserId)
	}()

}
func GenrateRecommendUser(ForUserId int) {

	////// load user Conatcs Id

	//var contacts []PhoneContactTable
	var PhoneNumbers []string
	base.DB.Select(&PhoneNumbers, "select PhoneNormalizedNumber from phone_contacts where UserId = ? ", ForUserId)

	var UserContactsIds []int
	q := "select Id from user where Phone in ('" + strings.Join(PhoneNumbers, "','") + "' ) "
	base.DB.Select(&UserContactsIds, q)

	/////////////////////////////////////
	var RecomUsers []RecommendUser

	//1:: contacts
	for _, uid := range UserContactsIds {
		if MemoryStore.UserFollowingList_GetFollowingTypeForUsers(ForUserId, uid) == 0 {
			r := RecommendUser{}
			r.UserId = ForUserId
			r.TargetId = uid
			r.Weight = 1
			r.CreatedTime = helper.TimeNow()

			RecomUsers = append(RecomUsers, r)
		}
	}
	/*helper.DebugPrintln("GenrateRecommends data: ",PhoneNumbers)
	  helper.DebugPrintln("GenrateRecommends data: ",UserContactsIds)
	  helper.DebugPrintln("GenrateRecommends data: ",q)
	*/
	// todo: implement this
	//2: follings of my followes

	//3: top user of MS -- for now jus followed count

	l := len(RecomUsers)

	for i := 0; i < l-100; i++ {
		if len(TopUserIds) > i {
			uid := TopUserIds[i]
			if MemoryStore.UserFollowingList_GetFollowingTypeForUsers(ForUserId, uid) == 0 {
				r := RecommendUser{}
				r.UserId = ForUserId
				r.TargetId = uid
				r.Weight = 1
				r.CreatedTime = helper.TimeNow()

				RecomUsers = append(RecomUsers, r)
			}
		}

	}

	//save to db
	//helper.DebugPrintln("GenrateRecommends data: ",RecomUsers)

	var _cs []interface{} //just for inserting in mass
	for _, r := range RecomUsers {
		v := r //it must be a local variable to -- r is just onece intiated
		_cs = append(_cs, &v)
		//base.DbInsertStruct( &r , "recommend_user")
	}

	base.DbExecute("delete from recommend_user where UserId =? ", ForUserId)
	base.DbMassReplacetStructPoninters("recommend_user", _cs...)
}
