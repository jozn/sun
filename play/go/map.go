package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strt")
	f11()
	fmt.Println("end")

	time.Sleep(time.Second * 10000)
}

func f10() {
	mp := make(map[int]int, 2000)
	for i := 0; i < 10000000; i++ {
		mp[i] = i
	}
}

func f11() {
	mp := make(map[int]map[int]int, 2000)
	for i := 0; i < 100000; i++ {
		mp[i] = make(map[int]int)
		for j := 0; j < 100; j++ {
			mp[i][j] = j
		}
	}
}
