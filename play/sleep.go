package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100000)
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(time.Millisecond * 10)
		}

	}()

	for o := range ch {
		if o%10 != 0 {
			fmt.Println(o)
			//fmt.Println("+")
		} else {
			fmt.Println("-")
		}
	}

}
