package main

import (
    "image"
    "os"
    "log"
    "image/jpeg"
    "strconv"
    "time"
    "image/color"
)

func main()  {
    srcFileName:="./x.jpg"
    file, err := os.Open(srcFileName)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }

    // decode jpeg into image.Image
    img, _ := jpeg.Decode(file)
    size :=img.Bounds().Size()
    rect :=img.Bounds().Bounds()

    newImg := image.NewRGBA(rect)

    matrix:=[]float32{
        -1, -1, -1,
        -1, 9, -1,
        -1, -1, -1,
    }

    matrix2:=[]float32{
        0, 0, 0,
        0, 0.4, 0,
        0, 0, 0,
    }

    /*-1, -2, -1,
    0, 0, 0,
    1, 2, 1,
    */
    _,_ = matrix2 ,matrix

    for x :=1 ; x < size.X -1 ; x++ {
        for y :=1 ; y < size.Y -1 ; y++ {
            //m[0]*img.At(x-1,y-1) + m[1]*img.At(x,y-1) + m[2]*img.At(x+1,y-1) + //row -1
            //m[0]*img.At(x-1,y) + m[1]*img.At(x-1,y) + m[2]*img.At(x-1,y) + //row 0
            //m[0]*img.At(x-1,y+1) + m[1]*img.At(x-1,y+1) + m[2]*img.At(x-1,y+1) + //row +1
            //
            //
            //newImg.Set(x, y,img.At(x, y).)
            cal(matrix,img ,newImg, x ,y)
        }
    }
     //_ =matrix[100]
    outFileName := srcFileName + strconv.Itoa(int(time.Now().Unix())) + ".jpg"
    toimg, _ := os.Create(outFileName)
    defer toimg.Close()

    jpeg.Encode(toimg, newImg,&jpeg.Options{jpeg.DefaultQuality} )

}

func cal(matrix []float32, srcImg image.Image, newImg *image.RGBA, X,Y int) {
    var rn,gn,bn float32 = 0,0,0// n:new
    var MV float32= 0
    for y:=-1 ; y<= 1; y++{
        for x:=-1 ; x<= 1; x++{
            //if x != 0 || y != 0 {
            //    continue
            //}
            r,g,b,_ :=srcImg.At(X+x,Y+y).RGBA()
            i:= (y+1)*3  + x + 1
            //fmt.Println(i,x,y)
            mv := matrix[i] //matrix val
            //mv = 1.2
            MV += mv
            rn += float32(r/256) * mv
            gn += float32(g/256) * mv
            bn += float32(b/256) * mv
        }
    }

    r,g,b,_ :=srcImg.At(X,Y).RGBA()
    r += r + uint32( MV*(256 - float32(r/256)) )
    g += g + uint32( MV*(256 - float32(g/256)) )
    b += b + uint32( MV*(256 - float32(b/256)) )

    //fmt.Println(r/256)
    //fmt.Println(rn)

    rgb := color.RGBA{ uint8(r), uint8(g), uint8(b),255}
    newImg.Set(X,Y,rgb)
    //newImg.Set(X,Y,srcImg.At(X,Y))
}

