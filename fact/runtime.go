package fact

import (
	"ms/sun/base"
	"ms/sun/helper"
)


func updateUsersAbout() {
	for i := 0; i < _factLastUserId(); i++ {
		s := helper.FactRandStrEmoji(120, true)
		if len(s) > 150 {
			s = s[0:150]
		}
		//s = "SSSSSSS"
		base.DB.Exec("update user set about=? where Id = ?", s, i)

	}
	print("boned updateUsersAbout")
}
