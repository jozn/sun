package helper

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func ImageResize(imageFilePath, outDirectory, fileNamePrefix string, sizes []int) bool {
	return imageCropThumb(imageFilePath, outDirectory, fileNamePrefix, sizes, false)
}

func ImageCropSquerThumb(imageFilePath, outDirectory, fileNamePrefix string, sizes []int) bool {
	return imageCropThumb(imageFilePath, outDirectory, fileNamePrefix, sizes, true)
}

func imageCropThumb(imageFilePath, outDirectory, fileNamePrefix string, sizes []int, toSquer bool) bool {

	///// File - directory func
	file, err := os.Open(imageFilePath)
	if err != nil {
		log.Println(err)
		return false
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return false
	}
	file.Close()

	os.MkdirAll(outDirectory, 0666)
	if err != nil {
		log.Println(err)
		return false
	}
	///////////////////////
	if toSquer {
		img = makeNewImageToSqurt(img)
	}

	for i := 0; i < len(sizes); i++ {
		sizeX := sizes[i]
		sizeY := 0
		if toSquer {
			sizeY = sizeX
		}
		m := resize.Thumbnail(uint(sizeX), uint(sizeY), img, resize.Lanczos3)
		out, err := os.Create(outDirectory + fileNamePrefix + "" + IntToStr(sizeX) + ".jpg")
		if err != nil {
			log.Println(err)
			continue
		}
		defer out.Close()
		// write new image to file
		jpeg.Encode(out, m, &jpeg.Options{Quality: 90})
	}
	return true
}

func makeNewImageToSqurt(img image.Image) image.Image {
	oldSize := img.Bounds().Max
	minSizeF := math.Min(float64(oldSize.X), float64(oldSize.Y))
	minSize := int(minSizeF)

	newSize := image.Rectangle{}
	newSize.Min = image.Point{0, 0}
	newSize.Max = image.Point{minSize, minSize}

	startPoint := image.Point{(oldSize.X - minSize) / 2, (oldSize.Y - minSize) / 2}

	result := image.NewRGBA(newSize)

	draw.Draw(result, newSize, img, startPoint, draw.Src)
	return result
}
