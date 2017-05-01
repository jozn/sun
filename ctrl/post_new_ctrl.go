package ctrl

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun/models"
	"path/filepath"
	"time"

	"ms/sun/models/x"

	"github.com/nfnt/resize"
)

func AddPostAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	txt := c.Req.Form.Get("text")

	///////////////// start of file functionality //////////////
	upladedFile, fd, err := c.Req.FormFile("file")
	if err == nil { //image
		defer upladedFile.Close()
		buket := models.Buket_GetNextForPostPhoto()
		t := time.Now()
		ext := filepath.Ext(fd.Filename) //.jpg
		photName := fmt.Sprintf("%d_%d_%v_%d_%v_%v_", t.Unix(), c.UserId(), t.Year(), t.Month(), t.Day(), t.Hour())

		photName = photName + "%s" + ext
		PathSrcPath := buket.BucketName + "/" + photName
		imageOrginal, err := jpeg.Decode(upladedFile)
		if err != nil {
			return err
		}

		h := math.Ceil(float64((imageOrginal.Bounds().Dy() * 1080) / imageOrginal.Bounds().Dx()))
		photo := &x.Photo{
			UserId:      c.UserId(),
			Title:       fd.Filename,
			BucketId:    buket.BucketId,
			Width:       1080,   //imageOrginal.Bounds().Dx(),
			Height:      int(h), //imageOrginal.Bounds().Dy(),
			Ratio:       float32(imageOrginal.Bounds().Dx()) / float32(imageOrginal.Bounds().Dy()),
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
			models.Bucket_savePhotoToBucket(photo, buket, img, size)
			lastImage = img
			if size == 1080 {
			}
		}

		//////////////// end of file functionality ////////////////

		post := x.Post{
			Id:           0,
			UserId:       c.UserId(),
			TypeId:       models.POST_TYPE_PHOTO,
			Text:         txt,
			FormatedText: "",
			//MediaUrl: "",
			MediaCount: 0,
			//MediaServerId: 0,
			//Width: 0,
			//Height: 0,
			SharedTo:       0,
			DisableComment: 0,
			HasTag:         0,
			LikesCount:     0,
			CommentsCount:  0,
			EditedTime:     0,
			CreatedTime:    helper.TimeNow(),
		}

        err = photo.Save(base.DB)
        if err != nil {
            return nil
        }

        models.Post_AddNewPostToDbAndItsTagsAndCounters(&post)
		photo.PostId = post.Id
        photo.Save(base.DB)

		c.SendJson(post)
		return nil
	} else { ///just text
		post := x.Post{}
		//TODO security: clean html of text
		post.Text = txt
		post.UserId = c.UserId()
		post.TypeId = models.POST_TYPE_TEXT
		post.CreatedTime = helper.TimeNow()
		post.CommentsCount = 0
		post.LikesCount = 0

		models.Post_AddNewPostToDbAndItsTagsAndCounters(&post)
		c.SendJson(post)
		return nil
	}
}

func GetPostsLatestAction(c *base.Action) base.AppErr {
	param := UpdateSessionActivityIfUser(c)

	selector := x.NewPost_Selector().
		OrderBy_Id_Desc().Limit(param.Limit).
		Offset(param.GetOffset())

	if param.Last > 0 {
		selector.Id_LT(param.Last)
	}

	posts, err := selector.GetRows(base.DB)
	if err != nil {
		return err
	}
	views := models.Views.PostsViews(posts, c.UserId())

	c.SendJson(views)
	return nil
}

func GetPostsStreamAction(c *base.Action) base.AppErr {
	p := MustBeUserAndUpdate(c)

	uid := c.UserId()

	fids := models.MemoryStore.UserFollowingList_Get(uid).Values()
	//var ins = make([]int,0, len(fids)+1)
	ins := append(fids, c.UserId())
	selctor := x.NewPost_Selector().UserId_In(ins).OrderBy_Id_Desc().Limit(p.Limit)

	if p.Last > 0 {
		selctor.Id_LT(p.Last)
	} else if p.Page > 0 {
		selctor.Offset(p.GetOffset())
	}

	posts, err := selctor.GetRows(base.DB)
	if err != nil {
		helper.DebugPrintln(err)
		c.SendJson(nil)
		return err
	}

	view := models.Views.PostsViews(posts, uid)
	c.SendJson(view)
	return nil
}
