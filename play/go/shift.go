package main

import (
	"fmt"
)

func main() {
	f2()
}

func f2() {
	const (
		PRIVAT  = 1 << 0
		PRIVAT2 = 1 << 1
		PRIVAT3 = 1 << 2
		PRIVAT4 = 1 << 3
	)

	i := PRIVAT | PRIVAT3 | PRIVAT4

	fmt.Printf("%064b \n", i)

	fmt.Printf("is 4 : %v \n", (i&PRIVAT4 > 0))
	fmt.Printf("%64b", (i & PRIVAT4))
}

func f1() {
	for i := uint(0); i < 65; i++ {
		fmt.Println(i+1, 1<<i, fmt.Sprintf("%064b", 1<<i))
	}

	fmt.Println(fmt.Sprintf("%016b", 127))
	fmt.Println(fmt.Sprintf("%b", "س"))
}
