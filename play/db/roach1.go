package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"ms/sun/helper"
	"time"
)

func main() {

	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")

	if err != nil {
		log.Panic(err)
	}
	sql := "INSERT into test_chat (`Id`, `TimeMs`,`Text` ,`Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) VALUES (?,?,?,?,?,?,?,?,?) "
	//(`Id`, `TimeMs`, `Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) values ('1', '5445', 'sad', '25', '1', '2', '6', '5')
	stm, err := DB.Prepare(sql)

	go func() {
		for {
			stm, err = DB.Prepare("INSERT into test_chat (`Id`, `TimeMs`,`Text` ,`Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) VALUES (?,?,?,?,?,?,?,?,?) ")
			time.Sleep(time.Second * 20)
		}
	}()
	if err != nil {
		log.Panic(err)
	}

	i := 0
	t := time.Now()
	for j := 0; j < 4; j++ {
		go func() {
			for i < 1000000000 {
				r, err := stm.Exec(
					//r, err := DB.Exec(sql,
					helper.NextRowsSeqId(),
					helper.TimeNowMs(),
					helper.FactRandStr(50),
					helper.RandString(40),
					rand.Intn(5000000),

					rand.Intn(50000000),
					rand.Intn(50000000),
					rand.Intn(50000000),
					rand.Intn(50000000),
				)
				if i%10000 == 0 {
					n := int64(0)
					var e2 error
					if err == nil {
						n, e2 = r.LastInsertId()
					}
					fmt.Println("QPS: ", float64(i)/(time.Now().Sub(t).Seconds()), humanize.FormatInteger("", i), err, n, e2)
				}
				i++
			}
		}()
	}

	go func() {
		/*for {
			DB.Exec("UPDATE test_chat set TimeMs = ?  WHERE UserId = ? ", helper.RandomUid64(), rand.Intn(5000000))
		}*/
	}()

	time.Sleep(time.Hour * 10)

}
