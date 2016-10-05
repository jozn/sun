package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strings"
	"time"
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

	go func() {
		for i := 0; i < 1000; i++ {
			addMassToTable()
			if i%1000 == 0 {
				fmt.Println(" MASS : ", i)
			}
		}
	}()

	time.Sleep(time.Second * 1)

	go func() {
		for i := 0; i < 1000000; i++ {
			MIX()
			if i%1000 == 0 {
				fmt.Println("MIX: ", i)
			}
		}
	}()

	//DB.Exec("LOCK TABLES a WRITE;")

	go func() {
		for i := 0; i < 1000000; i++ {
			addToTable()
			if i%1000 == 0 {
				fmt.Println("addToTable: ", i)
			}
		}
	}()
	//DB.Exec("UNLOCK TABLES;")

	go func() {
		for i := 0; i < 2000000; i++ {
			update(1000000)
			if i%1000 == 0 {
				fmt.Println("update: ", i)
			}
		}
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			delete()
			if i%1000 == 0 {
				fmt.Println("  DELETE : ", i)
			}
		}
	}()

	for i := 0; i < N; i++ {
		err, cnt := selectoneIndexed(N)
		if i%1000 == 0 {
			fmt.Println(" select index : ", i, " ", cnt, err)
		}
	}

	for i := 0; i < N; i++ {
		err, cnt := selectone(N)
		if i%1000 == 0 {
			fmt.Println(" select on PK : ", i, " ", cnt, err)
		}
	}

	for i := 0; i < 1; i++ {
		err, cnt := selectTable()
		if i%2 == 0 {
			fmt.Println(" select All : ", i, " ", cnt, err)
		}
	}
	time.Sleep(time.Second * 100000)

}

func addToTable() {
	DB.Exec("insert into bench3 (`Text`,`Time`,`Name`,Indexed) values(?,?,?,?)",
		" ahbdas askjn kjdnksd «تسدنتشسی ندنتد jdksdasd asdsajdnas kajdnasd dasdknjasdkad adkdnjadnadad asd diasdasd asda dadadajdbakdjbad m أ یشسینشستی",
		rand.Intn(50000000),
		"jasndkjas akj",
		rand.Intn(500000000)) //500 M)
}

func addMassToTable() {
	return
	N := 1000
	str := strings.Repeat("(?,?,?,?),", N)
	str = str[:len(str)-1]
	var arr []interface{}
	for i := 0; i < N; i++ {
		arr = append(arr, " MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS MASS ")
		arr = append(arr, rand.Intn(50000000))
		arr = append(arr, "MAS MASS MASS jasndkjas akj")
		arr = append(arr, rand.Intn(500000000))
	}

	DB.Exec("replace into bench3 (`Text`,`Time`,`Name`,Indexed) values "+str, arr...)
}

func update(cnt int) {
	DB.Exec("update bench3 set `Text`= ? where id = ?",
		" UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED ",
		rand.Intn(cnt))
}

func selectTable() (error, int) {
	var res []Bench
	err := DB.Unsafe().Select(&res, "select * from bench3 ")
	return err, len(res)
}

func selectone(cnt int) (error, int) {
	var res []Bench
	err := DB.Unsafe().Select(&res, "select * from bench3 where Id = ? ", rand.Intn(cnt))
	return err, len(res)
}

func selectoneIndexed(cnt int) (error, int) {
	var res []Bench
	err := DB.Unsafe().Select(&res, "select * from bench3 where Indexed = ? ", rand.Intn(500000000))
	return err, len(res)
}

func delete() {
	DB.Exec("delete  from bench3 where Indexed = ? ", rand.Intn(500000000)) // 500 Mil
	DB.Exec("delete  from bench3 where Id = ? ", rand.Intn(5000000))        // 5 Mil
}

func MIX() {
	addToTable()
	update(1000000)
	selectone(N)
	delete()
	selectoneIndexed(N)
}
