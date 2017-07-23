package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"ms/sun/helper"
	"time"
)

func main() {

	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")

	if err != nil {
		log.Panic(err)
	}

	stm, err := DB.Prepare("INSERT into test2 (UID) VALUES (?) ")

	if err != nil {
		log.Panic(err)
	}

	i := 0
	t := time.Now()
	for j := 0; j < 4; j++ {
		go func() {
			for i < 1000000000 {
				//DB.Exec("INSERT into test2 (UID) VALUES (?) ", helper.RandString(2))
				//DB.Exec("INSERT into test2 (UID) VALUES (?) ", helper.IntToStr(i)+":" + helper.IntToStr(helper.NextRowsSeqId()) )
				//DB.Exec("INSERT into test2 (UID) VALUES (?) ", helper.IntToStr(i%1000)+":"+fmt.Sprintf("%x", helper.NextRowsSeqId()))
				//, helper.IntToStr(i%1000)+":"+fmt.Sprintf("%x", helper.NextRowsSeqId()))

				stm.Exec(helper.IntToStr(i%1000) + ":" + fmt.Sprintf("%x", helper.NextRowsSeqId()))
				if i%10000 == 0 {
					fmt.Println("QPS: ", float64(i)/(time.Now().Sub(t).Seconds()), humanize.FormatInteger("", i))
				}

				//atomic.AddInt(&i,1)
			}
		}()
	}

	time.Sleep(time.Hour * 10)

}
