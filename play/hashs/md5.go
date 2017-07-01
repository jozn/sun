package main

import (
	"fmt"
	"ms/sun/helper"
)

func main() {
	s := "sadass"
	fmt.Println(helper.MD5BytesToString([]byte(s)))
	fmt.Println(helper.MD5BytesToString([]byte("a")))
	fmt.Println(helper.MD5BytesToString([]byte("A")))
	fmt.Println(helper.MD5BytesToString([]byte("kdhaksjhdasj")))

}
