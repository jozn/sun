package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	opts := opt.Options{
		Strict: opt.DefaultStrict,
	}
	db, err := leveldb.OpenFile("./write21.db", &opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	//	batch := &leveldb.Batch{}
	t2 := time.Now().UnixNano()

	for i := 0; i < 100000; i++ {
		// db.Put([]byte("key-"+strconv.Itoa(rand.Intn(300000))), []byte("val-"+strconv.Itoa(i)), nil)
		db.Put([]byte("key-"+strconv.Itoa(rand.Intn(3))), []byte("val-"+strconv.Itoa(i)), nil)
		//		data, err := db.Get([]byte("bar"), nil)
		//		fmt.Println(data,err)
	}

	//	db.Write(batch,nil)

	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

}
