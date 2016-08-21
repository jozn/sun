package helper

import "ms/sun/config"

func JustRecover() {
    if r := recover(); r != nil {
        if config.IS_DEBUG {
            DebugPrintln("ERROR PANICED RECOVRED - ERR:: ",r)
        }
    }
}
