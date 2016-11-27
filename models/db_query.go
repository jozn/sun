package models

import (
	"math"
	"ms/sun/base"
	"ms/sun/helper"
)

///ALL Deprecated

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
