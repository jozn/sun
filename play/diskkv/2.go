package main

import (
	"fmt"

	"github.com/peterbourgon/diskv"
	"time"
)

func main() {
	d := diskv.New(diskv.Options{
		BasePath:     "./my-diskv-data-directory",
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})

	key := "alpha"
	t2:= time.Now().UnixNano()

	for i := 0; i < 5000; i++ {

		if err := d.Write(key, []byte{'1', '2', '3'}); err != nil {
			panic(err)
		}
	}

	fmt.Println("time: " ,(time.Now().UnixNano() - t2)/1e6 )


	value, err := d.Read(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", value)

	if err := d.Erase(key); err != nil {
		panic(err)
	}
}
