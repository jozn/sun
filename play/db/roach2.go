package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/syncthing/syncthing/lib/sync"
	"ms/sun/helper"
	"strings"
	"time"
)

func main() {
	// Connect to the "bank" database.
	db, err := sqlx.Open("postgres", "postgresql://maxroach@78.46.161.115:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 1000)
			rows, err := db.Query("SELECT count(*) from accounts")
			if err != nil {
				log.Println("cnt2 ERRR: ", err)
				continue
			}
			id := 0
			rows.Next()
			rows.Scan(&id)
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>> count", id)
		}
	}()

	cnt := 0
	mass := func() {
		for j := 0; j < 1000000; j++ {
			//go func(i int) {
			oA := make([]string, 0, 1000)
			for z := 0; z < 100; z++ {
				oA = append(oA, fmt.Sprintf(`(%d,%d,'%s')`, helper.NextRowsSeqId(), helper.NextRowsSeqId(), helper.FactRandStrEmoji(100, true)))
			}

			o := strings.Join(oA, `,`)

			qq := `INSERT INTO accounts (id, balance,txt) VALUES ` + o
			if _, err := db.Exec(qq); err != nil {
				log.Println(err)
			}
			if j%10 == -1 {
				fmt.Println(qq)
			}
			cnt++
			fmt.Println("Mass: ", cnt)
			//g.Done()
			//}(j)
		}
		//g.Wait()

	}

	mass2 := func() {
		db, err := sqlx.Open("postgres", "postgresql://maxroach@78.46.161.115:26257/bank?sslmode=disable")
		if err != nil {
			log.Println("error connecting to the database: ", err)
			return
		}
		for j := 0; j < 1000000; j++ {
			//go func(i int) {
			oA := make([]string, 0, 1000)
			for z := 0; z < 100; z++ {
				oA = append(oA, fmt.Sprintf(`INSERT INTO accounts (id, balance,txt) VALUES (%d,%d,'%s');`, helper.NextRowsSeqId(), helper.NextRowsSeqId(), helper.FactRandStrEmoji(100, true)))
			}

			o := strings.Join(oA, ``)

			qq := o
			if _, err := db.Exec(qq); err != nil {
				log.Println(err)
			}
			if j%10 == -1 {
				fmt.Println(qq)
			}
			cnt++
			fmt.Println("Mass2: ", cnt)
		}
	}

	_ = mass

	for z := 0; z < 100; z++ {
		go mass2()
	}

	// Insert two rows into the "accounts" table.

	cnt2 := 0
	for i := 0; i < 0; i++ {
		g := sync.NewWaitGroup()

		for j := 0; j < 50; j++ {
			g.Add(1)
			go func(i int) {
				if _, err := db.Exec(
					"INSERT INTO accounts (id, balance) VALUES ($1, $2)", helper.NextRowsSeqId(), helper.NextRowsSeqId()); err != nil {
					//log.Println(err)

				}
				cnt2++
				fmt.Println(cnt2)
				//g.Done()
			}(i * j)
		}
		//g.Wait()
	}

	time.Sleep(time.Second * 10000)

	// Print out the balances.
	/*rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %d\n", id, balance)
	}
	*/

}
