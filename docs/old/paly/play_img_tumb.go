package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"strconv"
)

func main() {
	for i := 0; i < 50; i++ {
		// open "test.jpg"
		fn := strconv.Itoa(i) + ".jpg"
		file, err := os.Open(fn)
		if err != nil {
			//log.Fatal(err)
			continue
		}

		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			//log.Fatal(err)
			continue
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		// m := resize.Resize(1000, 0, img, resize.NearestNeighbor)
		// m := resize.Resize(1000, 0, img, resize.MitchellNetravali)
		m := resize.Thumbnail(50, 0, img, resize.NearestNeighbor)

		out, err := os.Create(strconv.Itoa(i) + "_Bilinear.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)

	}

}
