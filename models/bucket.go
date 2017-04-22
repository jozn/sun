package models

import (
	"fmt"
	"image"
	"image/jpeg"
	"ms/sun/base"
	"ms/sun/config"
	"ms/sun/helper"
	"os"
	"strconv"
	"strings"
	"time"
)

const BUCKET_TYPE_POST_PHOTO = 1

func Buket_GetNextForPostPhoto() *Bucket {
	buket, err := NewBucket_Selector().
		ContentObjectTypeId_EQ(BUCKET_TYPE_POST_PHOTO).
		OrderBy_BucketId_Desc().
		GetRow(base.DB)

	if err != nil || buket.ContentObjectCount >= config.BUCKET_SIZE {
		buket = bucket_creatNextBucket()
	}

	buket.ContentObjectCount += 1
	buket.Save(base.DB)
	return buket
}

func bucket_creatNextBucket() *Bucket {
	buket := &Bucket{
		BucketName:          helper.RandString(15),
		Server1Id:           1,
		ContentObjectTypeId: BUCKET_TYPE_POST_PHOTO,
		CreatedTime:         helper.TimeNow(),
	}
	buket.Save(base.DB)
	buket.BucketName = bucket_nextName(buket.BucketId)
	buket.Save(base.DB)
	return buket
}

func bucket_nextName(i int) string {
	t := time.Now()
	return fmt.Sprintf("posts/%d_%02d_%02d/b%d", t.Year(), t.Month(), t.Day(), i)
}

func Bucket_savePhotoToBucket(photo *Photo, buket *Bucket, img image.Image, size int) {
	//for now just save to local disk
	dirName := config.UPLOAD_DIR_POSTS + buket.BucketName
	fileName := config.UPLOAD_DIR_POSTS + strings.Replace(photo.PathSrc, "%s", strconv.Itoa(size), -1)
	os.MkdirAll(dirName, 0666)

	newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		print(2, err.Error())
		return //err
	}
	defer newFile.Close()

	opt := &jpeg.Options{Quality: 40}
	jpeg.Encode(newFile, img, opt)

}
