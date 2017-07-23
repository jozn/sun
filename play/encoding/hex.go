package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	Tohex()
}

func Tohex() {

	s := "Package hex implements hexadecimal encoding and decoding."
	r := hex.EncodeToString([]byte(s))
	fmt.Println(r)

}
