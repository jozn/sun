package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {

	//log every LOC of apps
	go func() {
		for {
			err := exec.Command(`loc2.exe`, "").Run()
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println("loc2 runned")
            time.Sleep(time.Second * 10 * 60)
		}
	}()

	for {
		err := exec.Command(`C:\Windows\System32\w32tm.exe`, "/resync").Run()
		fmt.Println(err)
		time.Sleep(time.Second * 120)
	}

}
