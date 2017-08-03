package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func main() {

	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")

	if err != nil {
		log.Panic(err)
	}
	sql := "INSERT into test_chat (`Id`, `TimeMs`,`Text` ,`Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) VALUES (?,?,?,?,?,?,?,?,?) "
	//(`Id`, `TimeMs`, `Name`, `UserId`, `c2`, `c3`, `c4`, `c5`) values ('1', '5445', 'sad', '25', '1', '2', '6', '5')
	_, err = DB.Prepare(sql)

	go func() {
		for i := 0; i < 1000000; i++ {
			DB.Exec("SELECT * FROM likes AS l JOIN `user` AS u WHERE l.UserId = u.Id AND l.PostId < 100 LIMIT 100 ;")
			//DB.Exec("SELECT * FROM likes LIMIT 100;")
			//time.Sleep(time.Second * 20)
			if i%100 == 0 {
				fmt.Println(i)
			}
		}
	}()
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(time.Hour * 10)

}
