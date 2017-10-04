package main

import (
	"fmt"
	"io/ioutil"
	"ms/sun/helper"
	"os/exec"
	"strings"
)

func main() {

	dir, err := ioutil.ReadDir("./")

	helper.NoErr(err)

	for _, f := range dir {
		if !f.IsDir() && strings.Contains(f.Name(), ".jpg") {
			cmd := exec.Command("ffmpeg.exe ", "-i", f.Name(), strings.Replace(f.Name(), ".jpg", ".webp", -1))
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
