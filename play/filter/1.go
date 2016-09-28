package main

import (
	"github.com/disintegration/gift"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		//fmt.Println("Usage:\tgoimger <file>")
		//os.Exit(1)
	}

	//srcFileName := os.Args[1] || "./x.jpg"
	srcFileName := "./x.jpg"
	srcFile, _ := os.Open(srcFileName)
	src, _, _ := image.Decode(srcFile)

	f := gift.ColorFunc(
		func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
			r = 1 - r0
			g = 1 - g0
			b = 1 - b0
			a = a0

			//r = 1 - r0   // invert the red channel
			//g = g0 + 0.1 // shift the green channel by 0.1
			//b = 0        // set the blue channel to 0
			//a = a0       // preserve the alpha channel
			return
		},
	)
	_ = f

	matrix := gift.Convolution(
		[]float32{
			-1, -2, -1,
			0, 0, 0,
			1, 2, 1,
		},
		false, false, false, 0,
	)

	matrix2 := gift.Convolution(
		[]float32{
			-1, 0, 1,
			-2, 0, 2,
			-1, 0, 1,
		},
		false, false, false, 0,
	)

	matrix3 := gift.Convolution(
		[]float32{
			-1, -1, -1,
			-1, 8, -1,
			-1, -1, -1,
		},
		false, false, false, 0,
	)

	matrix4 := gift.Convolution(
		[]float32{
			3, 10, 3,
			0, 0, 0,
			-3, -10, -3,
		},
		false, false, false, 0,
	)

	matrix5 := gift.Convolution(
		[]float32{
			-1, -1, -1, -1, -1,
			-1, -2, -4, -2, -1,
			-1, -4, 80, -4, -1,
			-1, -2, -4, -2, -1,
			-1, -1, -1, -1, -1,
		},
		true, false, false, 0,
	)

	matrix6 := gift.Convolution(
		[]float32{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
		},
		true, false, false, 0,
	)

	_, _, _, _, _, _ = matrix, matrix2, matrix3, matrix4, matrix5, matrix6

	// let's make a new gift
	g := gift.New(
		//gift.Grayscale(),
		//gift.UnsharpMask(1.0, 0.5, 0.0),
		//gift.Sepia(100),
		//gift.Saturation(500),
		//gift.Sigmoid(0.7, 7.0),
		//gift.Gamma(8),
		//gift.GaussianBlur(5),
		//gift.Hue(-50),
		//gift.Invert(),
		//gift.Maximum(7,false),
		//matrix2 ,
		//matrix,
		matrix5,
		//gift.Invert(),
		//gift.Grayscale(),
		//gift.Median(7,true),

	)

	// dest - output image
	dest := image.NewRGBA(g.Bounds(src.Bounds()))
	// draw result
	g.Draw(dest, src)

	outFileName := srcFileName + strconv.Itoa(int(time.Now().Unix())) + ".jpg"
	toimg, _ := os.Create(outFileName)
	defer toimg.Close()

	jpeg.Encode(toimg, dest, &jpeg.Options{jpeg.DefaultQuality})
}
