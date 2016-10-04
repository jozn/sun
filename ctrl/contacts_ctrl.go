package ctrl

import (
	"encoding/json"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
)

func GrabAllUserContactsCtrl(a *base.Action) base.AppErr {
	MustBeUserAndUpdate(a)
    a.Req.ParseForm()
	contactsForm := a.Req.Form.Get("contacts")
	//DeviceUuidId := a.Req.Form.Get("uuid")
	var conts []models.PhoneContact
	err := json.Unmarshal([]byte(contactsForm), &conts)
	if err != nil {
		return nil
	}
    helper.Debug("contacts: ",conts)
	//fmt.Println(conts)
	/*var _cs []interface{} //just for inserting in mass
	for _, c := range conts {
		c.Id = 0 //security
		c.UserId = a.UserId()
		c.CreatedTime = helper.TimeNow()
		c.DeviceUuidId = 0 ///DeviceUuidId
		_cs = append(_cs, c)
	}*/

    for _, c := range conts {
        c.Id = 0 //security
        c.UserId = a.UserId()
        c.CreatedTime = helper.TimeNow()
        c.DeviceUuidId = 0 ///DeviceUuidId
        c.Insert(base.DB)
    }

	/*//save to DB
	base.DB.Exec("delete from phone_contacts where UserId = ?", a.UserId())
	base.DbMassInsertStruct("phone_contacts", _cs...)
*/

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

    a.SendJson(allNumbers)

	return nil
}
