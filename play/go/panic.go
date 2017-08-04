package main

import (
	"fmt"
	"ms/sun/helper"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i interface{}) {
			/*defer func() {
			    r := recover()
			    if r != nil {
			        fmt.Println(r)
			    }
			}()*/

			/*if r := recover(); r != nil {
			    fmt.Println(r)
			}*/

			defer helper.JustRecover()
			fmt.Println("*")
			panic(i)
			fmt.Println("+")

		}(i)
	}

	time.Sleep(time.Second * 1000)
}
