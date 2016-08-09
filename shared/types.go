package shared

import (
    "ms/sun/helper"
    "ms/sun/config"
)

func NewAvatarFileName(UserId int) *AvatarFileName{
    a := AvatarFileName{}
    a.Path = "./upload/avatar/"
    a.FileName = helper.IntToStr(UserId) + "_" +helper.RandString(4) + "_"
    a.ServerId = 1
    a.FullUrl = config.CDN_IMAGE_SERVER_FULL_URL + a.Path + a.FileName + "%s.jpg"

    return &a
}

type AvatarFileName struct  {
    FullUrl string
    Path string
    FileName string
    ServerId int
}

