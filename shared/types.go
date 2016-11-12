//todo move to baase? and change base to shared?
package shared

import (
	"ms/sun/config"
	"ms/sun/helper"
)

func NewAvatarFileName(UserId int) *AvatarFileName {
	a := AvatarFileName{}
	a.Path = "./upload/avatar/" + helper.IntToStr(UserId%config.NUMBER_OF_USER_AVATAR_DIRS) + "/"
	a.FileName = helper.IntToStr(UserId) + "_" + helper.RandString(4) + "_"
	a.ServerId = 1
	a.FullUrl = config.CDN_IMAGE_SERVER_FULL_URL + a.Path[2:] + a.FileName + "%s.jpg"

	return &a
}

type AvatarFileName struct {
	FullUrl  string
	Path     string
	FileName string
	ServerId int
}
