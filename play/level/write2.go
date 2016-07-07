package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"time"
	"strconv"
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
//	batch := &leveldb.Batch{}
	t2:= time.Now().UnixNano()

	for i := 2000; i < 3000; i++ {
		db.Put([]byte("key-"+strconv.Itoa(i)), []byte("val-"+strconv.Itoa(i) ),nil )
//		data, err := db.Get([]byte("bar"), nil)
//		fmt.Println(data,err)
	}

//	db.Write(batch,nil)

	fmt.Println("time: " ,(time.Now().UnixNano() - t2)/1e6 )

}
