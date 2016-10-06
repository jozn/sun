package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strings"
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
			for i < 100000 {
				bulk_addMassToTable()
				if n == 0 && i%100 == 0 {
					sec := time.Now().Sub(ts).Seconds()
					qps := float64(i) / sec
					fmt.Printf("addToTable: qps=%d  - %d \n", int(qps), i)
				}
				atomic.AddInt64(&i, 1)
			}
		}()
	}

	time.Sleep(time.Second * 100000)
}

func bulk_addMassToTable() {
	N := 1000
	var arr []string
	for i := 0; i < N; i++ {
		s := fmt.Sprintf("('MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS', %d , 'MAS MASS MASS jasndkjas akj', %d )", rand.Intn(50000000), rand.Intn(500000000))
		arr = append(arr, s)
	}
	strs := strings.Join(arr, ",")
	DB.MustExec("insert into bench4 (`Text`,`Time`,`Name`,Indexed) values " + strs)
}
