package fact

import (
	//	"ms/sun/base"
	"log"
	//	"github.com/jmoiron/sqlx"
    "strconv"
    "time"
)

var __DEV__ bool = true

func noErr(err error) {
	if err != nil {
		// log.Fatalln(err) //painc
		log.Panic("***** PANIC ****: ", err)
	}
}

func now() int {
    return int(time.Now().Unix())
}

func e(x interface{}) {

}

func intToStr(i int) string {
    return strconv.Itoa(i)
}
