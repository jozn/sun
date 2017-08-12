package main

import (
	"fmt"
	"ms/sun/helper"
	"time"
)

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	//close(queue)

	i := 0
	go func() {
		for {
			i++
			time.Sleep(time.Microsecond * 1)
			queue <- "tmmmwo" + helper.IntToStr(i)
		}
	}()
	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
