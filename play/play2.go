package main

import (
	// "encoding/json"
	"fmt"
	// "github.com/disintegration/imaging"
	"io"
	// "math/rand"
	"net/url"
	"os"
	// "strconv"
	// "strings"
)

// func uploadTestPlay1(c *Action) {

// 	file, handler, err := c.Req.FormFile("file")
// 	// noErr(err)
// 	if err != nil {
// 		fmt.Println("errr in Req.FormFile", err)
// 		return
// 	}
// 	defer file.Close()

// 	uri, err := url.Parse(handler.Filename)
// 	noErr(err)
// 	name := uri.Path

// 	temFileName := "./public/temp/" + name

// 	f, err := os.OpenFile(temFileName, os.O_WRONLY|os.O_CREATE, 0666)
// 	noErr(err)
// 	defer f.Close()

// 	_, err = io.Copy(f, file)
// 	noErr(err)

// 	//now convert image
// 	img, err := imaging.Open(temFileName)
// 	noErr(err)

// 	ext := fileExt(temFileName)

// 	resizsSizes := []int{50, 100, 200, 400, 600}
// 	userIdS := intToStr(c.UserId())

// 	for _, size := range resizsSizes {
// 		thumb := imaging.Thumbnail(img, size, size, imaging.CatmullRom)
// 		devPrintn(ext)
// 		// avatraName := "./public/avatars/" + intToStr(c.UserId()) + "_" + intToStr(rand.Intn(100000)) + "_" + intToStr(size) + ext
// 		sizeStr := intToStr(size)
// 		avatraName := avatrBuilderPath(userIdS, ext, sizeStr)
// 		err = imaging.Save(thumb, avatraName)
// 		noErr(err)
// 	}

// 	user := GetUserById(c.UserId())
// 	user.AvatarUrl = avatrBuilderPath(userIdS, ext, "%s")
// 	dbUpdateStruct(&user, "user")

// 	c.SendJson(user)
// 	// fmt.Println(c.Req.Form)
// 	// c.SendText(handler.Filename)

// 	// e(file)
// }

func uploadTest2Dep(c *Action) {

	for i := 0; i < 500; i++ {
		// DB.MustExec("INSERT INTO test (text) VALUES (?)", "dssdcsdcs udsbc sducbsdcc scybscysc scscbsdss csdcybsycbsc scsycbsc scsycb cssssssssss")
	}

	file, handler, err := c.Req.FormFile("file")
	// noErr(err)
	if err != nil {
		fmt.Println("errr in Req.FormFile", err)
		return
	}
	defer file.Close()
	uri, err := url.Parse(handler.Filename)
	noErr(err)
	fn := uri.Path
	f, err := os.OpenFile("./public/"+fn, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("errr in io.copy", err)
		return
	}

	fmt.Println(c.Req.Form)
	c.SendText(handler.Filename)

	e(file)
}

type _playJson struct {
	Nmae string
	Nums []int
	Info UserInfo
	User
}

func playJsonAction(c *Action) {
	res := make([]_playJson, 100)
	for i := 0; i < 100; i++ {
		v := _playJson{}
		i := UserInfo{}
		i.ResharedCount = 2986658
		v.Info = i
		v.User = User{}
		v.Nums = []int{1, 25, 6, 2, 5, 36, 75, 2, 5, 6, 5, 77, 2, 2, 514, 55, 9879}
		v.Nmae = "asdjasudhabsdhasbdad asdhbashdgas ,غاشغعیذشسغیشسی شسیعغذشسیذشسی عغإغغ،ر،تیاذشسایذشسیعشسیذشسعغیذل"

		res = append(res, v)
	}

	c.SendJson(res)
}

func playTagsAction(c *Action) {
	var tagsPost []TagPost
	// err = DB.Select(&tagsPost, "select * from tags_posts where TagId = ? limit 50", tag.Id)
	err := DB.Select(&tagsPost, "select * from tags_posts")
	noErr(err)
	c.SendJson(tagsPost)
}
