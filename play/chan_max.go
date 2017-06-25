package main

import (
	"fmt"
	"math/rand"
	"ms/sun/helper"
	"strconv"
	"sync/atomic"
	"time"
)

const N = 50000

var Cnt int64 = 0

func main() {
	go func() {
		for {
			time.Sleep(time.Second * 2)
			helper.GcPrintAll()
		}
	}()

	for p := 0; p < N; p++ {

		go func(p int) {
			ch := make(chan string, 2)
			go func() {
				i := 0
				for {
					ch <- strconv.Itoa(i)
					i++
					atomic.AddInt64(&Cnt, 1)
					time.Sleep(time.Millisecond * 100)
				}

			}()

			j := 0
			for _ = range ch {
				j++
				if rand.Intn(N*10) == 2 {
					fmt.Printf("%d -- %d  -- Cnt: %d \n", p, j, Cnt)
				}
			}

		}(p)

	}

	time.Sleep(time.Minute * 10)

}
