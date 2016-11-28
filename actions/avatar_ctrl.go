package actions

import (
	// "encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"io"
	"math/rand"
	. "ms/sun/base"
	. "ms/sun/models"
	"net/url"
	"os"
	// "strconv"
	// "strings"
)

func UploadAvatarAction(c *Action) {

	file, handler, err := c.Req.FormFile("file")
	// noErr(err)
	if err != nil {
		fmt.Println("errr in Req.FormFile", err)
		return
	}
	defer file.Close()

	uri, err := url.Parse(handler.Filename)
	noErr(err)
	name := uri.Path

	temFileName := "./public/temp/" + name

	f, err := os.OpenFile(temFileName, os.O_WRONLY|os.O_CREATE, 0666)
	noErr(err)
	defer f.Close()

	_, err = io.Copy(f, file)
	noErr(err)

	//now convert image
	img, err := imaging.Open(temFileName)
	noErr(err)

	ext := fileExt(temFileName)

	rnd := intToStr(rand.Intn(100000))
	//TODO: clean this(remove invoke by params)
	avatrBuilderPath := func(userId, ext, size string) string {
		return "public/avatars/" + userId + "_" + rnd + "_" + size + ext
	}

	resizsSizes := []int{50, 100, 200, 400, 600}
	userIdS := intToStr(c.UserId())

	for _, size := range resizsSizes {
		thumb := imaging.Thumbnail(img, size, size, imaging.CatmullRom)
		devPrintn(ext)
		// avatraName := "./public/avatars/" + intToStr(c.UserId()) + "_" + intToStr(rand.Intn(100000)) + "_" + intToStr(size) + ext
		sizeStr := intToStr(size)
		avatraName := avatrBuilderPath(userIdS, ext, sizeStr)
		err = imaging.Save(thumb, avatraName)
		noErr(err)
	}

	//user := GetUserById(c.UserId())
	user,_ := UserById(DB,c.UserId())
	user.AvatarUrl = avatrBuilderPath(userIdS, ext, "%s")
	DbUpdateStruct(&user, "user")

	c.SendJson(user)

}

func RemoveAvatarAction(c *Action) {
	/*//user := GetUserById(c.UserId())
	user := GetU(c.UserId())
	user.AvatarUrl = ""
	DbUpdateStruct(&user, "user")

	c.SendJson(user)*/
}
