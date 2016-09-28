package main

import (
	"fmt"

	"math/rand"
	"time"
)

func shuffle(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		fmt.Println(i, n)
		r := rand.Intn(int(i + 1))
		arr[i], arr[r] = arr[r], arr[i]
	}

	fmt.Println(arr)
}

func shuffleCopy(arr []int) {
	n := len(arr)
	arr2 := make([]int, n)
	for i := n - 1; i > 0; i-- {
		fmt.Println(i, n)
		r := rand.Intn(int(i + 1))
		arr[i], arr[r] = arr[r], arr[i]
	}

	fmt.Println(arr)
}

func main() {
	rand.Seed(time.Now().Unix())
	rand.Seed(55)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(a)
	fmt.Println("Hello, playground")
	fmt.Println(a)
}
