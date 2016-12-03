package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strings"
	"time"
	//"ms/sun/sync"
	"strconv"
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

	for j := 0; j < 1; j++ {
		n := j
		go func() {
			ts := time.Now()
			lastI := i
			for i < 10000 {
				addToTable()
				if n == 0 && i%100 == 0 {
					sec := time.Now().Sub(ts).Seconds()
					ts = time.Now()
					k := i
					CNT := k - lastI
					lastI = k
					qps := float64(CNT) / sec
					fmt.Printf("addToTable: qps=%d  - %d \n", int(qps), i)
					//fmt.Println(CNT,sec,i)
				}
				atomic.AddInt64(&i, 1)
			}
		}()
	}

	time.Sleep(time.Second * 100000)

}

func addToTable() {
	N := 1000
	var arr []string
	for i := 0; i < N; i++ {
		arr = append(arr, strconv.Itoa(rand.Intn(50000000)))
	}
	strs := strings.Join(arr, ",")
	DB.MustExec("select * from bench_ints where `Id5` in ( " + strs + ") order by Id5 desc limit 100")
}
