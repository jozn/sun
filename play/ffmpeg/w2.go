package main

import (
	"fmt"
	"io/ioutil"
	"ms/sun/helper"
	"os/exec"
	"strings"
    "time"
)

func main() {
    c :=0
	dir, err := ioutil.ReadDir("./")
	ch := make(chan string, 20)
	helper.NoErr(err)

	go func() {
        for _, f := range dir {
            if !f.IsDir() && strings.Contains(f.Name(), ".jpg") {
                ch <- f.Name()
            }
        }
    }()

	for i := 0; i < 20; i++ {
		go func() {
			for fn := range ch {
				cmd := exec.Command("ffmpeg.exe ", "-i", fn, strings.Replace(fn, ".jpg", "+.jpg", -1))
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
				c++
				fmt.Println(c)
			}
		}()
	}

	time.Sleep(time.Second * 10000)
}
