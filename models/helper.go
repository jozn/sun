package models

import (
	"fmt"
	// "reflect"
	"strconv"
	"time"
)

//all are deprecated
func devPrintn(p ...interface{}) {
	fmt.Println(p...)
}

func debug(p ...interface{}) {
	if __DEV__ {
		fmt.Println(p...)
	}
}

func minutes(s int) time.Duration {
	return time.Duration(s) * time.Minute
}

func intToStr(i int) string {
	return strconv.Itoa(i)
}


func now() int {
	return int(time.Now().Unix())
}
