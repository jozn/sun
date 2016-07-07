package actions

import (
	_"github.com/jmoiron/sqlx"
	"log"
	cacheDrive "github.com/pmylund/go-cache"
	"time"
)

//var DB *sqlx.DB
var __DEV__ bool
var cashe *cacheDrive.Cache

func init()  {
	__DEV__ = true
//	DB, err := sqlx.Connect("mysql", "root:123456@/ms3?charset=utf8mb4")
	// DB, err = sqlx.Connect("mysql", "root:123456@/ms3?charset=ascii")
//	noErr(err)
///	_ = DB
cashe = cacheDrive.New(5*time.Minute, 30*time.Second)

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

