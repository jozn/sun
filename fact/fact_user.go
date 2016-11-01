package fact

import (
	// "encoding/json"
	// "fmt"
	// "reflect"
	//"strings"
	"math/rand"
	. "ms/sun/base"
	. "ms/sun/models"
	"time"
    "ms/sun/helper"
)

func FactUser1(c *Action) {
	print("factoring user + user_info\n")
	_userFirstNameSamples := []string{"حمید", "سیما", "نیلوفر", "آرش", "Armin", "Leili", "نیاز", "فرخ", "atash", "محمد علی"}
	_userLastNameSamples := []string{"کریمی", "کمانگیر", "بزگ", "فدیایش", "رستگار", "میلانی", "مستانی", "فروهی", "مصداق", "هدایت", "fish", "Nosrat", "Fadaghi"}
	_userUserNameSamples := []string{"fish", "Nosrat", "Fadaghi", "atash", "Jigar", "DooSTi"}

	u := User{}
	// u.Id int
	// u.UserId
	u.UserName = randSilceString(_userUserNameSamples) + intToStr(rand.Intn(10000))
	u.FirstName = randSilceString(_userFirstNameSamples)
	u.LastName = randSilceString(_userLastNameSamples)
	//u.FullName = u.FirstName + " " + u.LastName
	u.Email = _randomEmail()
	u.AvatarUrl = "public/avatars/" + intToStr(rand.Intn(15)+1) + ".jpg"
	u.CreatedTime = int(time.Now().Unix())

	ui := UserInfo{}

	res, err := DbInsertStruct(&u, "user")
	noErr(err)
	id, _ := res.LastInsertId()
	ui.UserId = int(id)
	DbInsertStruct(&ui, "user_info")

	//folowing list
	fl := FollowingList{}
	fl.UserId = int(id)
	fl.Name = "دنبال میکنم"
	// fl.IsPimiry = true
	//debug(fl)
	r, err := DbInsertStruct(&fl, "following_list")
	// noErr(err)
	//debug(err)
	pid, _ := r.LastInsertId()
	DB.MustExec("update user set PrimaryFollowingList = ? where Id =?", pid, id)
	e(u)

}

func FactUpdateAboutMe(c *Action) {
    user,_ := NewUser_Selector().OrderBy_Id_Asc().GetRows(DB)

    for _ ,u := range user {
        s := helper.FactRandStrEmoji(80, true)
        if len(s) > 150 {
            //s = s[0:150]
        }

        u.About = s
        //u.Id = i+1
        err:=u.Update(DB)
        if err != nil{
            helper.DebugPrintln(err)
        }
    }

}

func FactRealUser(c *Action) {
	//q :=User{}

	us := []User{
	//User{
	//	UserName: "AliAsgh",
	//	FirstName: "علی اصضر",
	//	LastName: "کریمی",
	//	Email:"ali@gmail.com",
	//	Phone: "+989176506200",
	//	AvatarUrl: "public/avatars/100.jpg",
	//},
	//User{
	//	UserName: "abas",
	//	FirstName: "فباس",
	//	LastName: "کریمی",
	//	Email:"abas2@gmail.com",
	//	Phone: "+989178023320",
	//	AvatarUrl: "public/avatars/101.jpg",
	//},
	//User{
	//	UserName: "abas22",
	//	FirstName: "علی الامردشتی",
	//	LastName: "",
	//	Email:"ali22@gmail.com",
	//	Phone: "+989177819068",
	//	AvatarUrl: "public/avatars/102.jpg",
	//},
	//User{
	//	UserName: "ali_bagheri",
	//	FirstName: "ali ",
	//	LastName: "bagrinia",
	//	Email:"ali22@gmail.com",
	//	Phone: "+989176152801",
	//	AvatarUrl: "public/avatars/103.jpg",
	//},
	//User{
	//	UserName: "Alipoor",
	//	FirstName: "خسین",
	//	LastName: "Alipoor",
	//	Email:"hosin@gmail.com",
	//	Phone: "+989178783190",
	//	AvatarUrl: "public/avatars/104.jpg",
	//},
	//User{
	//	UserName: "Edris",
	//	FirstName: "Edris",
	//	LastName: " اعتماد",
	//	Email:"Edris22@gmail.com",
	//	Phone: "+987823623335",
	//	AvatarUrl: "public/avatars/105.jpg",
	//},
	//User{
	//	UserName: "Hosin22",
	//	FirstName: "حسین ",
	//	LastName: "کریمی",
	//	Email:"hossin44@gmail.com",
	//	Phone: "+989125472641",
	//	AvatarUrl: "public/avatars/106.jpg",
	//},
	//User{
	//	UserName: "Motahareh",
	//	FirstName: "Mot",
	//	LastName: "Moti",
	//	Email:"Motahareh@gmail.com",
	//	Phone: "+989361556536",
	//	AvatarUrl: "public/avatars/107.jpg",
	//},
	//User{
	//	UserName: "Aazam",
	//	FirstName: "Azam",
	//	LastName: "کریمی",
	//	Email:"Aazam@gmail.com",
	//	Phone: "+989178541520",
	//	AvatarUrl: "public/avatars/108.jpg",
	//},
	//User{
	//	UserName: "KarimiJAVAd",
	//	FirstName: "جواد",
	//	LastName: "کرمی",
	//	Email:"javad@gmail.com",
	//	Phone: "+989393074664",
	//	AvatarUrl: "public/avatars/109.jpg",
	//},
	}
	for _, user := range us {
		user.CreatedTime = int(time.Now().Unix())
		DbInsertStruct(&user, "user")

	}
}
