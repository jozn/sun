package fact

import (
    "ms/sun/models"
    "os"
    "image/jpeg"
    "io/ioutil"
    "math/rand"
    "log"
    "ms/sun/helper"
    "ms/sun/base"
    "time"
    "fmt"
    "ms/sun/shared"
)

const NUM_OF_USERS  = 80

func FactPosts() {
    p := models.Post{}
    p.TypeId = 1
    p.UserId = rand.Intn(80)+1
    p.Text = helper.FactRandStrEmoji(150,true)

    if rand.Intn(3) == 1{
        img , W ,H := randImage()
        p.MediaServerId =1
        p.MediaUrl = img
        p.Width = W
        p.Height = H
        p.TypeId = 2
    }

    models.AddNewPostToDbAndItsMeta(p)
    //base.DbInsertStruct(&p,"post")

}


var imageFiles []os.FileInfo;
const  dir ="./upload/posts1"

func init() {
    imageFiles, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    _ = imageFiles
}

func randImage() (string, int,int) {
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

    return fn[2:], b.X , b.Y

}

func FactUserAvatars(c *base.Action) {
    dir :="./upload/_avatars"
    imageFiles, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for uid :=0 ; uid < 100 ; uid++ {
        u := models.UserMemoryStore.GetForUser(uid)
        if u != nil{
            fn := dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()

            userAvatr := shared.NewAvatarFileName(uid)
            t1:=time.Now().UnixNano()
            helper.ImageCropSquerThumb(fn,userAvatr.Path,userAvatr.FileName, []int{50,100,200,400})
            t := (time.Now().UnixNano() - t1)/ 1e6
            u := models.UserMemoryStore.GetForUser(uid)
            u.AvatarUrl = userAvatr.FullUrl
            u.UserBasic.UpdateToTable()
            fmt.Println("time for avatar gen:",t)

        }
    }

}

func FactFollow(c *base.Action) {
    models.Follow(rand.Intn(NUM_OF_USERS),rand.Intn(NUM_OF_USERS))
}

func FactUnFollow(c *base.Action) {
    models.UnFollow(rand.Intn(NUM_OF_USERS),rand.Intn(NUM_OF_USERS))
}


