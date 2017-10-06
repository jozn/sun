package helper

import (
	"log"
	"ms/sun/config"
)

func NoErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func NoErrDebug(err error) {
	if config.IS_DEBUG {
		if err != nil {
			log.Panic(err)
		}
	}
}
