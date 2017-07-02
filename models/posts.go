package models

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"math"
	"math/rand"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models/x"
	"path/filepath"
	"time"
)

const POST_TYPE_TEXT = 1
const POST_TYPE_PHOTO = 2
const POST_TYPE_VIDEO = 3

type NewPostAddParams struct {
	UserId int
	Text   string
	//Image         image.Image
	UploadedFileName string
	UploadedImage    []byte
}

func Post_AddNewPostToDbAndItsTagsAndCounters(post *x.Post) {
    fmt.Println(post)
	err := post.Save(base.DB)
	if err == nil {
		Counter.UpdateUserPostsCounts(post.UserId, 1)
		Tags_AddTagsInPost(post)
		Mentioned_AddUserMentionedInPost(post)
	} else {
		helper.DebugErr(err)
	}
    fmt.Println("54555")
}

func Post_PostsByIds(PostIds []int) ([]*x.Post, error) {
	if len(PostIds) == 0 {
		return nil, errors.New("PostIds must not be empty")
	}
	return x.NewPost_Selector().Id_In(PostIds).GetRows(base.DB)
}
func PostTypeIdToName(id int) string {
	switch id {
	case POST_TYPE_TEXT:
		return "text"

	case POST_TYPE_PHOTO:
		return "photo"

	case POST_TYPE_VIDEO:
		return "video"
	}
	return "none"
}

func PostAddNew(param NewPostAddParams) (*x.Post, error) {
	if param.UserId <= 0 {
		return nil, ErrUserIdMustBePostive
	}

	if len(param.Text) == 0 && len(param.UploadedImage) == 0 {
		return nil, ErrPostIsInvaid
	}

    helper.Debug(param)
    fmt.Println(param)

	var photo *x.Photo
	post := &x.Post{
		Id:             0,
		UserId:         param.UserId,
		TypeId:         POST_TYPE_TEXT,
		Text:           param.Text,
		FormatedText:   "",
		MediaCount:     0,
		SharedTo:       0,
		DisableComment: 0,
		HasTag:         0,
		LikesCount:     0,
		CommentsCount:  0,
		EditedTime:     0,
		CreatedTime:    helper.TimeNow(),
	}

	if len(param.UploadedImage) > 0 {
		buket := Buket_GetNextForPostPhoto()
		t := time.Now()
		ext := filepath.Ext(param.UploadedFileName) //.jpg
		//photName := fmt.Sprintf("%d_%d_%v_%d_%v_%v_", t.Unix(), c.UserId(), t.Year(), t.Month(), t.Day(), t.Hour())
		photName := fmt.Sprintf("%d_%d_%d_", t.Unix(), param.UserId, (rand.Intn(899) + 100))

		photName = photName + "%s" + ext
		PathSrcPath := buket.BucketName + "/" + photName
		imageOrginal, err := jpeg.Decode(bytes.NewBuffer(param.UploadedImage))
		if err == nil {
			helper.DebugErr(err)
			//md5
			md5 := helper.MD5BytesToString(param.UploadedImage)

			h := math.Ceil(float64((imageOrginal.Bounds().Dy() * 1080) / imageOrginal.Bounds().Dx()))
			photo = &x.Photo{
				UserId:      param.UserId,
				Title:       param.UploadedFileName,
				BucketId:    buket.BucketId,
				Width:       1080,   //imageOrginal.Bounds().Dx(),
				Height:      int(h), //imageOrginal.Bounds().Dy(),
				Ratio:       float32(imageOrginal.Bounds().Dx()) / float32(imageOrginal.Bounds().Dy()),
				HashMd5:     md5,
				Color:       helper.ExtractColoer(imageOrginal),
				CreatedTime: helper.TimeNow(),
				PathSrc:     PathSrcPath,
				W1080:       1,
				W720:        1,
				W480:        1,
				W320:        1,
				W160:        1,
				W80:         1,
			}

			var lastImage image.Image = imageOrginal //otimization for faster image
			siezes := []int{1080, 720, 480, 320, 160, 80}
			for _, size := range siezes {
				img := resize.Resize(uint(size), 0, lastImage, resize.Lanczos3)
				Bucket_savePhotoToBucket(photo, buket, img, size)
				lastImage = img
				if size == 1080 {
				}
			}

			err = photo.Save(base.DB)
			if err == nil {
				helper.DebugErr(err)
				post.TypeId = POST_TYPE_PHOTO
			}
		}
	}

	Post_AddNewPostToDbAndItsTagsAndCounters(post)

	if photo != nil {
		photo.PostId = post.Id
		photo.Save(base.DB)
	}

	return post, nil

}
