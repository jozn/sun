package fact

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	// "strings"
	. "ms/sun/base"
	. "ms/sun/models"
	//	"math/rand"
	//	"strconv"
	//	"time"
)

func FactPhoneContacts(c *Action) {
	// print("factoring pst + media\n")
	var rs []PhoneContactTable
	//	DB.Select(&rs, "select dd* from phone_contacts where user_id = ?",2)
	DB.Select(&rs, "select * from phone_contacts ")
	//	c.SendJson(rs)
	var likes []Like
	DB.Select(&likes, "select * from likes where PostId = 10 limit 50 ")
	c.SendJson(rs)
	//	u := PhoneContactTable{}
	//	u.UserId = rand.Intn(50)
	//	u.Text = factRnddStr(300)
	//	u.CreatedTimestamp = int(time.Now().Unix()) - rand.Intn(50000)
	//	u.TypeId = 10 //text see: gloab_types
	//
	//	DbInsertStruct(&u, "post")

}
