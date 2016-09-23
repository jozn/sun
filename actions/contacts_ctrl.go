package actions

import (
	"encoding/json"
	"ms/sun/base"
	"ms/sun/models"
	//"fmt"
	// "../util"
	"ms/sun/helper"
	//"strings"
	"fmt"
)

func SyncAllContactsAction(a *base.Action) base.AppErr {
	//c.MustPost()
	contactsForm := a.Req.Form.Get("contacts")
	var conts []models.PhoneContactTable
	err:=json.Unmarshal([]byte(contactsForm), &conts)
	if err != nil {
		return nil
	}
	//fmt.Println(conts)
	var _cs []interface{} //just for inserting in mass
	for _, c := range conts {
		c.UserId = a.UserId()
		c.CreatedTime = helper.TimeNow()
		_cs = append(_cs, c)
	}

	//save to DB
	base.DB.Exec("delete from phone_contacts where UserId = ?", a.UserId())
	base.DbMassInsertStruct("phone_contacts", _cs...)

	//get registerd contacts
	var allNumbers []string
	for _, c := range conts {
		if c.PhoneNormalizedNumber != "" {
			allNumbers = append(allNumbers, c.PhoneNormalizedNumber)
		}
	}

	var users []models.UserTable
	//var users []models.UserBasicAndMe
	q := "select * from user where Phone in (" + base.DbSliceStringToSafeIns(allNumbers) + ")"
	helper.DebugPrintln(q)
	_ = base.DB.Select(&users, q)

	users_res := make([]models.UserBasicPhoneAndMe,0,len(users))

	list := models.PreloadFollowingsListTypesForUser(a.UserId())
	for _,u := range users {
		cu := models.UserBasicPhoneAndMe{}
		cu.FromUser(u,list)
		//cu.UserBasic = u.UserBasic
		cu.Phone = u.Phone
		//cu.UpdatedTimestamp = u.UpdatedTimestamp
		//cu.FollowingType = list.FollowingType(u.Id)
		users_res = append(users_res, cu)
	}
	fmt.Println(len(users),len(users_res))
	a.SendJson(users_res)
	return nil
}
