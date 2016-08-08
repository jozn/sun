package fact

import (
    "ms/sun/models"
    "os"
    "image/jpeg"
    "io/ioutil"
    "math/rand"
    "log"
    "ms/sun/helper"
)

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
