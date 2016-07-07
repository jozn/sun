package fact

import (
	//	"ms/sun/base"
	"log"
	//	"github.com/jmoiron/sqlx"
)

var __DEV__ bool

//var DB *sqlx.DB

func init() {
	//	DB = base.DB
	__DEV__ = true
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

func e(x interface{}) {

}
