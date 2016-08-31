package main

import (
    "fmt"
    "math/rand"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type Bench struct  {
    Id int
    Text string
    Time int
    Name int
}

func main() {
    DB, _ = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/tops?charset=utf8mb4")

    for i:= 0 ; i < 1000000 ; i++ {
        addToTable()
        if i%1000 ==0 {
            fmt.Println(i)
        }
    }

    for i:= 0 ; i < 0 ; i++ {
        err ,cnt :=selectTable()
        if i%10 ==0 {
            fmt.Println(i," ",cnt ,err)
        }
    }

}

func addToTable() {
    DB.Exec("insert into bench (`Text`,`Time`,`Name`) values(?,?,?)",
     " ahbdas askjn kjdnksd «تسدنتشسی ندنتد jdksdasd asdsajdnas kajdnasd dasdknjasdkad adkdnjadnadad asd diasdasd asda dadadajdbakdjbad m أ یشسینشستی" ,
        rand.Intn(500000),
        "jasndkjas akj")
}

func selectTable() (error, int){
    var res []Bench
    err:= DB.Unsafe().Select(&res,"select * from bench ")
    return  err , len(res)
}
