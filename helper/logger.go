package helper

import (
	"fmt"
	"ms/sun/config"
)

func Debug(vals ...interface{}) {
	if config.IS_DEBUG {
		fmt.Println(vals...)
	}
}

func Debugf(str string, vals ...interface{}) {
	if config.IS_DEBUG {
		fmt.Printf(str, vals...)
	}
}
