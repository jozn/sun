package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"strconv"
	"time"
)

func printMemStats() {
	//	var m runtime.MemStats
	//	runtime.ReadMemStats(&m)
	//	fmt.Printf("sys: %d, alloc: %d, idle: %d\n", m.HeapSys,
	//		m.HeapAlloc, m.HeapIdle)
}

func main() {
	opts := opt.Options{
		Strict: opt.DefaultStrict,
	}
	t2 := time.Now().UnixNano()

	for i := 0; i < 10000; i++ {
		db, err := leveldb.OpenFile("./mass/"+strconv.Itoa(i)+".db", &opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}

	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)
}
