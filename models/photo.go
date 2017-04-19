package models

import "ms/sun/config"

func Convert_PhotoToNewPhotoView(p *Photo) *PhotoView {
	if p != nil {
		v := &PhotoView{
			PhotoId:     p.PhotoId,
			UserId:      p.UserId,
			PostId:      p.PostId,
			AlbumId:     p.AlbumId,
			ImageTypeId: p.ImageTypeId,
			Width:       p.Width,
			Height:      p.Height,
			Ratio:       p.Ratio,
			UrlFormat:   "",
			//Sizes:[]int
		}
		v.UrlFormat = config.CDN_IMAGE_SERVER_FULL_URL + p.PathSrc
		v.Sizes = make([]int, 0, 6)

		if p.W80 == 1 {
			v.Sizes = append(v.Sizes, 80)
		}
		if p.W160 == 1 {
			v.Sizes = append(v.Sizes, 160)
		}
		if p.W320 == 1 {
			v.Sizes = append(v.Sizes, 320)
		}
		if p.W480 == 1 {
			v.Sizes = append(v.Sizes, 480)
		}
		if p.W720 == 1 {
			v.Sizes = append(v.Sizes, 720)
		}
		if p.W1080 == 1 {
			v.Sizes = append(v.Sizes, 1080)
		}
		return v
	}
	return &PhotoView{}
}
