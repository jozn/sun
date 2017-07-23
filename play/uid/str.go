package main

import (
	"fmt"
	"ms/sun/helper"
)

func main() {

	const N = 1000
	c := 1000000000000
	s := "1000000000000:1000000"
	arr := [N]string{}
	for i := 999; i > 1; i-- {
		m := (N - i)
		arr[m] = helper.IntToStr(c + i)
	}

	fmt.Println(arr)
	fmt.Println(len([]byte(fmt.Sprintf("%0b", 1))))
	fmt.Println(len([]byte(s)))
	fmt.Println(fmt.Sprintf("%0320b", 11))

	p := int32(1)

	p = 1 << 30
	fmt.Println(p)

}
