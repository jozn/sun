package main

import (
	"fmt"
	"unicode"
)

func main() {

	j := 0
	for i := 0; i < 1000000; i++ {
		if unicode.Is(unicode.Javanese, rune(i)) {
			fmt.Println(i, string(i))
			j++
		}
	}
	fmt.Println(j)

}
