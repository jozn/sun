package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	opts := opt.Options{
		Strict: opt.DefaultStrict,
	}
	db, err := leveldb.OpenFile("./write1.db", &opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	t2 := time.Now().UnixNano()

	for i := 0; i < 1000; i++ {
		data, err := db.Get([]byte("key-"+strconv.Itoa(i)), nil)
		fmt.Println(string(data), err)
	}

	for i := 2000; i < 3002; i++ {
		data, err := db.Get([]byte("key-"+strconv.Itoa(i)), nil)
		fmt.Println(string(data), err)
	}

	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

}
