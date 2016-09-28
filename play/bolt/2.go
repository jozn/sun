package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	size := 10000
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("./my13.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin(true)
	if err != nil {
		//		return err
		log.Fatal(err)

	}
	defer tx.Rollback()

	// Use the transaction...
	_, err = tx.CreateBucket([]byte("MyBucket"))
	if err != nil {
		//		return err
		log.Fatal(err)

	}

	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		//		/return err
		log.Fatal(err)

	}
	t2 := time.Now().UnixNano()

	db.Update(func(tx *bolt.Tx) error {
		for i := 1; i < size; i++ {
			b := tx.Bucket([]byte("MyBucket"))
			err = b.Put([]byte("answer-"+strconv.Itoa(i)), []byte(strconv.Itoa(rand.Int())))
		}
		return err
	})

	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

	for i := 1; i < size; i += 1000 {
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			if b == nil {
				log.Fatal("ERRRR", b)
			}
			v := b.Get([]byte("answer-" + strconv.Itoa(i)))
			fmt.Printf("The answer is: %s\n", v)
			return nil
		})
	}
	//...
	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

}
