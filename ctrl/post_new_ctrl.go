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

	"github.com/nfnt/resize"
)

func AddPostAction(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	txt := c.Req.Form.Get("text")

	///////////////// start of file functionality //////////////
	upladedFile, fd, err := c.Req.FormFile("file")
	if err == nil { //image
		defer upladedFile.Close()
		buket := models.Buket_getNextForPost()
		t := time.Now()
		ext := filepath.Ext(fd.Filename)
		photPath := fmt.Sprintf("%d_%d_%v_%d_%v_%v_", t.Unix(), c.UserId(), t.Year(), t.Month(), t.Day(), t.Hour())

		photPath = photPath + "%s" + ext
		PathSrc := buket.BucketName + "/" + photPath
		imageOrginal, err := jpeg.Decode(upladedFile)
		if err != nil {
			return err
		}

		h := math.Ceil(float64((imageOrginal.Bounds().Dy() * 1080) / imageOrginal.Bounds().Dx()))
		photo := &models.Photo{
			UserId:      c.UserId(),
			Title:       fd.Filename,
			BucketId:    buket.BucketId,
			Width:       1080,   //imageOrginal.Bounds().Dx(),
			Height:      int(h), //imageOrginal.Bounds().Dy(),
			Ratio:       float32(imageOrginal.Bounds().Dx()) / float32(imageOrginal.Bounds().Dy()),
			CreatedTime: helper.TimeNow(),
			PathSrc:     PathSrc,
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
		/*
		   dirName := fmt.Sprintf("./upload/posts/%v/%d/%v/%v/", t.Year(), t.Month(), t.Day(), t.Hour())
		   ext := ".jpg"
		   index := strings.LastIndex(fd.Filename, ".")
		   if index > 0 {
		       ext = fd.Filename[index:]
		   }
		   fileName := fmt.Sprintf("%v%v%v%v%s", dirName, (t.UnixNano() / 1e6), "_", c.UserId(), ext)
		   err = os.MkdirAll(dirName, 0666)
		   newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

		   if err != nil {
		       print(2, err.Error())
		       return err
		   }
		   defer newFile.Close()
		   i, err := io.Copy(newFile, upladedFile)
		   print("Upladed media Size Write: ", i)*/
		//////////////// end of file functionality ////////////////

		post := models.Post{}
		//TODO security: clean html of text
		post.Text = txt
		post.UserId = c.UserId()
		post.TypeId = models.POST_TYPE_PHOTO
		post.CreatedTime = helper.TimeNow()
		post.CommentsCount = 0
		post.LikesCount = 0
		//post.MediaUrl = fileName[2:]
		post.MediaServerId = 1
		post.Height = 640
		post.Width = 640

		models.AddNewPostToDbAndItsMeta(&post)
		photo.PostId = post.Id
		photo.Save(base.DB)

		c.SendJson(post)
		return nil
	} else { ///just text
		post := models.Post{}
		//TODO security: clean html of text
		post.Text = txt
		post.UserId = c.UserId()
		post.TypeId = models.POST_TYPE_TEXT
		post.CreatedTime = helper.TimeNow()
		post.CommentsCount = 0
		post.LikesCount = 0

		models.AddNewPostToDbAndItsMeta(&post)
		c.SendJson(post)
		return nil
	}
}

func GetSinglePostAction_NEW(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)
	idstr := c.Req.Form.Get("post_id")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}
	var post models.Post
	err := base.DB.Get(&post, "select * from post where Id = ? ", id)
	if err != nil {
		return nil
	}

	view := models.GetPostToPostAndDetailes(&post, c.UserId())
	c.SendJson(view)
	return nil
}

func PostDeleteAction_NEW(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)

	idstr := c.Req.Form.Get("post_id")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}

	models.DeletePost(c.UserId(), id)

	return nil
}

func PostUpdateAction_NEW(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	idstr := c.Req.Form.Get("post_id")
	text := c.Req.Form.Get("text")
	id := helper.StrToInt(idstr, 0)
	if id < 1 {
		return nil
	}

	var post models.Post
	err := base.DB.Get(&post, "select * from post where Id = ? ", id)
	if err != nil {
		return nil
	}

	if post.UserId != c.UserId() {
		return nil
	} else {
		post.Text = text
		base.DbUpdateStruct(&post, "post")
		//todo caches?
	}

	return nil
}

func GetPostsStraemAction_NEW(c *base.Action) base.AppErr {
	MustBeUserAndUpdate(c)
	const LIMIT = 30

	laststr := c.Req.Form.Get("last") //last that have
	pagestr := c.Req.Form.Get("page")
	last := helper.StrToInt(laststr, 0)
	page := helper.StrToInt(pagestr, 0)
	limit := c.GetParamInt("limit", 30)

	_ = last
	_ = page

	uid := c.UserId()
	println("LIMIT: ", limit, c.Req.Form.Get("limit"), " +")
	//fids := models.GetAllPrimiryFollowingIds(uid)
	//fids := models.UserMemoryStore.GetAllFollowingsListOfUser(uid).Elements
	fids := models.MemoryStore.UserFollowingList_Get(uid).Elements
	ins := make([]int, 0, len(fids)+1)
	copy(ins, fids)
	ins = append(ins, c.UserId())
	print("ins: ", ins)
	selctor := models.NewPost_Selector().UserId_Ins(fids...).Or().UserId_EQ(c.UserId()).OrderBy_Id_Desc().Limit(limit)

	if last > 0 {
		selctor.Id_LT(last)
	} else if page > 0 {
		selctor.Offset((page - 1) * limit)
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

func GetPostsLatestAction_NEW(c *base.Action) base.AppErr {
	UpdateSessionActivityIfUser(c)

	laststr := c.Req.Form.Get("last") //last that have
	pagestr := c.Req.Form.Get("page")
	last := helper.StrToInt(laststr, 0)
	page := helper.StrToInt(pagestr, 0)
	//limit := helper.StrToInt("limit", LIMIT)

	_ = last
	_ = page

	uid := c.UserId()

	// dbIns(len(fids))e
	sql := "select * from post  order by Id Desc limit 100 "

	var rs []models.Post
	base.DB.Select(&rs, sql)

	view := models.PostsToPostsAndDetailesV1(rs, uid)
	c.SendJson(view)
	return nil
}