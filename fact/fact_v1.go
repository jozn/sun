package fact

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"math/rand"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"ms/sun/shared"
	"os"
	"time"
)

const NUM_OF_USERS = 80

func FactPost(c *base.Action) {
	FactPosts()
}
func FactPosts() {
	p := models.Post{}
	p.TypeId = 1
	p.UserId = rand.Intn(80) + 1
	p.Text = helper.FactRandStrEmoji(200, true)

	if rand.Intn(3) == 1 {
		img, W, H := randImage()
		p.MediaServerId = 1
		p.MediaUrl = img
		p.Width = W
		p.Height = H
		p.TypeId = 2
	}

	models.AddNewPostToDbAndItsMeta(&p)
	//base.DbInsertStruct(&p,"post")

}

var imageFiles []os.FileInfo

const dir = "./upload/posts1"

func init() {
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	_ = imageFiles
}

func randImage() (string, int, int) {
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range imageFiles {
		_ = f
		//fmt.Println(f.Name())
	}
	fn := "./upload/posts1" + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds().Size()

	return fn[2:], b.X, b.Y

}

func FactUserAvatars(c *base.Action) {
	dir := "./upload/_avatars"
	imageFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for uid := 0; uid < 100; uid++ {
		u := models.UserMemoryStore.GetForUser(uid)
		if u != nil {
			fn := dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()

			userAvatr := shared.NewAvatarFileName(uid)
			t1 := time.Now().UnixNano()
			helper.ImageCropSquerThumb(fn, userAvatr.Path, userAvatr.FileName, []int{50, 100, 200, 400})
			t := (time.Now().UnixNano() - t1) / 1e6
			u := models.UserMemoryStore.GetForUser(uid)
			u.AvatarUrl = userAvatr.FullUrl
			//u.UserBasic.UpdateToTable()
			base.DbUpdateStruct(&u, "user")
			fmt.Println("time for avatar gen:", t)

		}
	}

}

func FactFollow(c *base.Action) {
	models.Follow(rand.Intn(NUM_OF_USERS)+1, rand.Intn(NUM_OF_USERS)+1)
}

func FactUnFollow(c *base.Action) {
	models.UnFollow(rand.Intn(NUM_OF_USERS)+1, rand.Intn(NUM_OF_USERS)+1)
}

func FactComment(c *base.Action) {
	print("factoring like+comment\n")
	// u.UserId = rand.Intn(50) + 1
	// u.PostId = rand.Intn(450) + 1
	UserId := rand.Intn(_factLastUserId()) + 1
	PostId := rand.Intn(_factLastPostId()) + 1
	n := rand.Intn(50) + 10 //10 at leat
	Text := helper.FactRandStrEmoji(n, true)
	co := models.Comment_Add(UserId, PostId, Text)
	c.SendJson(co)

}

func FactLike(c *base.Action) {
	models.Like_LikePost(rand.Intn(_factLastUserId())+1, rand.Intn(_factLastPostId())+1)
	//models.UserMemoryStore.AddPostLike(rand.Intn(_factLastUserId())+1, rand.Intn(_factLastPostId())+1)
}

func FactUnLike(c *base.Action) {
    models.Like_UnlikePost(rand.Intn(_factLastUserId())+1, rand.Intn(_factLastPostId())+1)
    //models.UserMemoryStore.AddPostLike(rand.Intn(_factLastUserId())+1, rand.Intn(_factLastPostId())+1)
}

var _Cnt_Post int = 0
var _Cnt_User int = 0

func FactLikeFast(c *base.Action) {
	if _Cnt_User == 0 {
		_Cnt_User = _factLastUserId()
	}
	if _Cnt_Post == 0 {
		_Cnt_Post = _factLastPostId()
	}
	models.Like_LikePost(rand.Intn(_Cnt_User)+1, rand.Intn(_Cnt_Post)+1)
	//models.UserMemoryStore.AddPostLike(rand.Intn(_factLastUserId())+1, rand.Intn(_factLastPostId())+1)
}

func FactMix(c *base.Action) {
	i := rand.Intn(4) + 1
	switch i {
	case 1:
		FactLike(c)
	case 2:
		FactComment(c)
	case 3:
		FactFollow(c)
	case 4:
		FactUnFollow(c)
	}
}

func FactDelayMix(c *base.Action) {
    t := c.GetParamInt("delay",10)
    num := c.GetParamInt("limit",10)

    for i:=0;i<num;i++{
        FactMix(c)
        time.Sleep(time.Millisecond*time.Duration(t))
    }
}


func _factLastPostId() int {
	var ps []models.Post
	base.DB.Select(&ps, "select * from post order by Id DESC limit 2 ")
	p := ps[0]

	return p.Id

}

func _factLastUserId() int {
	var ps []models.User
	base.DB.Select(&ps, "select * from user order by Id DESC limit 2 ")
	p := ps[0]

	return p.Id

}

func FactLike2(c *base.Action) {
	print("factoring likes post\n")
	//COUNT = 50
	l := models.Like{}
	l.PostId = rand.Intn(500) + 1
	l.UserId = rand.Intn(80) + 1
	l.CreatedTime = now()
	base.DbInsertStruct(&l, "likes")
}
