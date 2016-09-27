package models

import (
    "ms/sun/base"
    "ms/sun/helper"
)

//TODO: must add update users as well
func SyncGetAllChangedUser(UserId , LastTime int ) []User {

    //contacts
    var users1[]User
    q:= "SELECT u.* FROM `user` AS u JOIN phone_contacts AS pc WHERE pc.UserId=? AND pc.CreatedTime >= ? AND u.Phone = pc.PhoneNormalizedNumber"
    err1:=base.DB.Select(&users1,q,UserId,LastTime)

    var users2[]User
    q2:= "SELECT u.* FROM `user` AS u JOIN following_list_member AS fm ON u.Id = fm.FollowedUserId WHERE fm.UserId=? AND fm.UpdatedTimeMs >= ?"
    err2:=base.DB.Select(&users2,q2,UserId,LastTime*1000)

    usersRes := make( []User, 0, len(users1)+len(users2) )

    for _, u := range users1 {
        usersRes = append(usersRes,u)
    }

    for _, u := range users2 {
        usersRes = append(usersRes,u)
    }

    helper.Debug("SyncGetAllChangedUser()",len(users1),err1,err2)

    return usersRes
}

func UserTableToUserViewCompation(users []User) []User {

    return nil
}
