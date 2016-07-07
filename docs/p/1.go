package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	fmt.Printf("%x", time.Now().UnixNano())
}
