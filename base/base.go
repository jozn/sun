package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"

)
var DB *sqlx.DB
var __DEV__ bool


func init()  {
	__DEV__ = true
//	DB, err := sqlx.Connect("mysql", "root:123456@/ms3?charset=utf8mb4")
//	// DB, err = sqlx.Connect("mysql", "root:123456@/ms3?charset=ascii")
//	noErr(err)
//	DB.MapperFunc(func(s string) string { return s })
	_ = DB
}

func noErr(err error) {
	if err != nil {
		// log.Fatalln(err) //painc
		log.Panic("***** PANIC ****: ", err)
	}
}

func noDevErr(err error) {
	if __DEV__ && err != nil {
		// log.Fatalln(err) //painc
		log.Panic("***** PANIC ****: ", err)
	}
}

//empty for fuell go nasty varible must used
func e(...interface{}) {

}

