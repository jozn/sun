package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type AstElem struct {
	num, name, designation, ref string
}

func main() {
	data := getdata()
	fmt.Println("Finished creating data")
	t := time.Now()
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("GC time: ", time.Since(t))
	fmt.Println(len(data))
}

func getdata() []AstElem {
	astelements := make([]AstElem, 2000000)
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Filling the strings with random short strings.
	for i := range astelements {
		astelements[i].num = fmt.Sprint(rd.Float64())
		astelements[i].name = fmt.Sprint(rd.Float64())
		astelements[i].designation = fmt.Sprint(rd.Float64())
		astelements[i].ref = fmt.Sprint(rd.Float64())
	}
	return astelements
}
