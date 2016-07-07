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
	size := 1000
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
	// _, err = tx.CreateBucket([]byte("MyBucket"))
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

	for j := 5000; j < 6000; j++ {
		t2 := time.Now().UnixNano()
		tx, err := db.Begin(true)
		if err != nil {
			//		return err
			log.Fatal("fatal", err)

		}
		// tx.Rollback()
		buket := "buket-" + strconv.Itoa(j)
		tx.CreateBucketIfNotExists([]byte(buket))
		tx.Commit()
		db.Update(func(tx *bolt.Tx) error {
			// TAG:
			b := tx.Bucket([]byte(buket))
			if b == nil {
				//		return err
				fmt.Println("***fatalLL: ", j, b)
				return nil

			}
			for i := 1; i < 10000; i++ {

				err = b.Put([]byte("answer111-"+strconv.Itoa(i+j*1000)), []byte(strconv.Itoa(rand.Int())))
			}
			return err
		})

		fmt.Println("INSERT TIME( "+strconv.Itoa(j)+" ): ", (time.Now().UnixNano()-t2)/1e6)
	}

	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

	for i := 1; i < size; i += 1 {
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			if b == nil {
				log.Fatal("ERRRR", b)
			}
			v := b.Get([]byte("answer-" + strconv.Itoa(i)))
			_ = v
			//fmt.Printf("The answer is: %s\n", v)
			return nil
		})
	}
	//...
	fmt.Println("time: ", (time.Now().UnixNano()-t2)/1e6)

}
