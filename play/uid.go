package main

import (
	"fmt"
	"ms/sun/helper"
)

func main() {
	mp := make(map[int]bool)

	for i := 0; ; i++ {
		uid := helper.RandomSeqUid()
		if _, ok := mp[uid]; ok {
			fmt.Printf("=========== %2d\n", uid)
		}
		if i%100000 == 0 {
			fmt.Printf("%2d -- %d\n", uid, i)
		}
		mp[uid] = true
	}

}
