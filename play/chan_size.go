package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 1000000

func main() {
	mp := make(map[int]chan string, 5000)

	for i := 0; i < N; i++ {
		mp[i] = make(chan string, 5)
	}

	for j := 0; j < N; j++ {
		mp[rand.Intn(N)] <- "1"
		//mp[i]
		time.Sleep(time.Microsecond * 1)
		if j%1000 == 0 {
			fmt.Println(j)
		}
	}

	time.Sleep(time.Second * 60)

}
