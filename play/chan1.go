package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 100000)
	go func() {
		i := 0
		for {
			ch <- strconv.Itoa(i)
			i++
			//time.Sleep(time.Millisecond * 10)
		}

	}()

	j := 0
	for _ = range ch {
		j++
		if j%1000000 == 0 {
			fmt.Println(j)
		}
	}

}
