package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	j := 0
	for i := 0; ; i++ {
		n := (rand.Intn(7000000000))
		if n < 7000000 {
			j++
			fmt.Printf("At %d -  %d : %s \n", j, i, humanize.Comma(int64(n)))
		}
		if i%1000 == 1 {
			fmt.Printf("ME: %s \n", humanize.Comma(int64(n)))
		}
		time.Sleep(time.Millisecond * 1)
	}

}
