package main

import (
	"fmt"
	"runtime"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"math/rand"
	"strconv"
	"time"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("sys: %d, alloc: %d, idle: %d\n", m.HeapSys,
		m.HeapAlloc, m.HeapIdle)
}

func main() {
	opts := opt.Options{
		Strict: opt.DefaultStrict,
	}
	db, err := leveldb.OpenFile("./testldbbatch8.db", &opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// ~295MB of data.
	fmt.Println("Creating records...")
	const numRecords = 1
	const recSize = 1550

	fmt.Println("Records created")

	// Print mem stats: ~325M
	printMemStats()

	// Add records to batch.
	fmt.Println("Creating batch...")
	batch := new(leveldb.Batch)
	//	batch.Load(make([]byte, 12, 12+numRecords*recSize+numRecords*11))
	t1 := time.Now().UnixNano()
	for i := 0; i < numRecords; i++ {
		//		batch.Put(records[i], nil)
		batch.Put([]byte("bar-"+strconv.Itoa(rand.Int())), []byte("another%%%%%%%^%&^%&^%&%$^$%^*^&((((((((((( MEEEEEEEEEEEEEE dasda value"))
	}

	fmt.Println("Batch created")

	// Print mem stats: ~680MB
	printMemStats()

	fmt.Println("Writing batch...")
	if err := db.Write(batch, nil); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Batch written")

	t2 := time.Now().UnixNano()
	fmt.Println("TIME", (t2-t1)/1e6)

	// Print mem stats: ~1.8GB!
	printMemStats()
}
