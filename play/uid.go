package main

import (
	"fmt"
	"github.com/syncthing/syncthing/lib/sync"
	"ms/sun/helper"
	"runtime"
	"time"
)

func main() {
	mp := make(map[int]bool)
	mut := sync.NewRWMutex()
	for n := 0; n < runtime.NumCPU(); n++ {
		go func(n int) {
			for i := 0; ; i++ {
				uid := helper.RandomSeqUid()
				mut.RLock()
				v := mp[uid]
				if v {
					fmt.Printf("=================== %2d\n", uid)
				}
				mut.RUnlock()
				if i%10000 == 0 {
					//fmt.Printf("%2d -- %d -- %d \n", uid, i, n)
					fmt.Printf("len -- %d  -- %d \n", len(mp), n)
				}
				if i%1000 == 0 {
					//runtime.Gosched()
					time.Sleep(time.Microsecond * 1)
				}
				mut.Lock()
				mp[uid] = true
				mut.Unlock()
			}
		}(n)

	}

	time.Sleep(time.Minute * 1)
}
