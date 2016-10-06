package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
	//"ms/sun/sync"
	"sync/atomic"
)

var DB *sqlx.DB

type Bench struct {
	Id      int
	Text    string
	Time    int
	Name    int
	Indexed int
}

const NM = 100000 //modify
const N = 100000  //select
func main() {
	DB, _ = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/tops?charset=utf8mb4")

	var i int64 = 0

	for j := 0; j < 4; j++ {
		n := j
		go func() {
			ts := time.Now()
			lastI := i
			for i < 10000 {
				addToTable_tran()
				if n == 0 && i%10 == 0 {
					sec := time.Now().Sub(ts).Seconds()
					ts = time.Now()
					k := i
					CNT := k - lastI
					lastI = k
					qps := float64(CNT) / sec
					fmt.Printf("addToTable Transection: qps=%d  - %d \n", int(qps), i)
					//fmt.Println(CNT,sec,i)
				}
				atomic.AddInt64(&i, 1)
			}
		}()
	}

	time.Sleep(time.Second * 100000)

}

func addToTable_tran() {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("transection err")
	}
	defer tx.Commit()

	for i := 0; i < 1000; i++ {
		tx.Exec("insert into bench4 (`Text`,`Time`,`Name`,Indexed) values(?,?,?,?)",
			" ahbdas askjn kjdnksd «تسدنتشسی ندنتد jdksdasd asdsajdnas kajdnasd dasdknjasdkad adkdnjadnadad asd diasdasd asda dadadajdbakdjbad m أ یشسینشستی",
			rand.Intn(50000000),
			"jasndkjas akj",
			rand.Intn(500000000)) //500 M)
	}

}
