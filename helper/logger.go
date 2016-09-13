package helper

import (
    "ms/sun/config"
    "fmt"
)

func Debug(vals ...interface{})  {
    if config.IS_DEBUG {
        fmt.Println(vals...)
    }
}

func Debugf(str string, vals ...interface{}) {
    if config.IS_DEBUG {
        fmt.Printf(str ,vals...)
    }
}

