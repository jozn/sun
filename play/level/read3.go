package main

import (
	"fmt"
	_ "strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	opts := opt.Options{
		Strict: opt.DefaultStrict,
	}
	db, err := leveldb.OpenFile("./write2.db", &opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	t2 := time.Now().UnixNano()

	for i := 0; i < 1000; i++ {
		//data, err := db.Get([]byte("key-"+strconv.Itoa(i)),nil)
		//fmt.Println(string(data),err)
	}

	//	for i := 2000; i < 3002; i++ {
	//		data, err := db.Get([]byte("key-"+strconv.Itoa(i)),nil)
	//		fmt.Println(string(data),err)
	//	}

	itr := &util.Range{Start: []byte("key-2940")}
	iter := db.NewIterator(itr, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), string(value))

	}
	iter.Release()

	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

}
