package main

import (
	"fmt"
	"math/rand"
	"ms/sun/helper"
	"runtime/debug"
	"time"
)

const N = 1000
const N2 = 10000000000

type book struct {
	id   int
	name string
	ref  *book
}

func main() {

	debug.SetGCPercent(200)

	//runtime.GC()

	bk := new(book)
	mp := make(map[int]chan *book, N)

	helper.GcPrintAllPerodicaly(4)
	//helper.GcForceGCPerodicaly(4)

	for i := 0; i < N; i++ {
		ch := make(chan *book, 5)
		mp[i] = ch
		go func() {
			for bks := range ch {
				if bks.id%20000000 == 0 {
					fmt.Println("ch: ", bks.id)
					bks.ref = nil
				}

			}
		}()
	}

	for j := 0; j < N2; j++ {
		b := &book{id: j,
			name: helper.IntToStr(j),
			ref:  bk,
		}
		bk = b

		mp[rand.Intn(N)] <- b
		//mp[i]

		if j%1000 == 0 {
			//fmt.Println(j)
			//time.Sleep(time.Microsecond * 1)
		}
	}

	time.Sleep(time.Second * 60)

}
