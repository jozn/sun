package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
	//"ms/sun/sync"
)

var DB_MY *sqlx.DB
var DB_PG *sqlx.DB

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
	DB_MY, _ = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/tops?charset=utf8mb4")
	DB_PG, _ = sqlx.Connect("postgres", "user=postgres dbname=tag sslmode=disable")

	go func() {
		for i := 1; i < 1000000; i++ {

			DB_PG.Exec(`insert into bench4 ("Text","Time","Name","Indexed") values($1,$2,$3,$4)`,
				" ahbdas askjn kjdnksd «تسدنتشسی ندنتد jdksdasd asdsajdnas kajdnasd dasdknjasdkad adkdnjadnadad asd diasdasd asda dadadajdbakdjbad m أ یشسینشستی",
				rand.Intn(50000000),
				"jasndkjas akj",
				rand.Intn(500000000))
			/*b := Bench4{Text: "asdasdas", Time: 1256}
			  b.Insert(DB_PG)*/

			/*b.Name = "up up"
			  b.Update(DB_PG)*/

			DB_PG.Exec(`update bench4 set "Text"= $1 where "Id" = $2`,
				" UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED ",
				rand.Intn(i))

			NewBench4_Selector().Id_EQ(rand.Intn(int(i))).GetRow(DB_PG)

			DB_PG.Exec(`delete  from bench4 where "Indexed" = $1 `, rand.Intn(500000000)) // 500 Mil
			DB_PG.Exec(`delete  from bench4 where "Id" = $1 `, rand.Intn(5000000))        // 5 Mil

			if i%1000 == 0 {
				fmt.Println("MIX pg: ", i)
			}
		}
	}()

	go func() {
		for i := 1; i < 1000000; i++ {

			DB_MY.Exec("insert into bench5 (`Text`,`Time`,`Name`,Indexed) values(?,?,?,?)",
				" ahbdas askjn kjdnksd «تسدنتشسی ندنتد jdksdasd asdsajdnas kajdnasd dasdknjasdkad adkdnjadnadad asd diasdasd asda dadadajdbakdjbad m أ یشسینشستی",
				rand.Intn(50000000),
				"jasndkjas akj",
				rand.Intn(500000000)) //500 M)

			DB_MY.Exec("update bench5 set `Text`= ? where id = ?",
				" UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED UPDATED ",
				rand.Intn(int(i)))

			var res []Bench
			DB_MY.Unsafe().Select(&res, "select * from bench5 where Id = ? ", rand.Intn(int(i)))

			if i%1000 == 0 {
				fmt.Println("MIX: ", i)
			}
		}
	}()

	time.Sleep(time.Second * 100000)

}

func MIX() {
	/*addToTable()
	update(1000000)
	selectone(N)
	delete()
	selectoneIndexed(N)*/
}

//xo_pg pgsql://postgres:123456@localhost/tag?sslmode=disable --template-path C:\Go\_gopath\src\github.com\jozn\go-orma-pg --single-file -o "model.pg.go" --suffix go --schema "" --package main
//xo mysql://root:123456@localhost:3307/tops --single-file --template-path C:\Go\_gopath\src\github.com\jozn\go-orma --out "model.my.go" --suffix go --schema "" --package main

//xo mysql://root:123456@localhost:3307/tops --single-file --template-path C:\Go\_gopath\src\github.com\jozn\go-orma-events --out "my.events.go" --suffix go --schema "" --package models
//xo mysql://root:123456@localhost:3307/tops --single-file --template-path C:\Go\_gopath\src\github.com\jozn\go-orma-cache --out "my.cache.go" --suffix go --schema "" --package models
