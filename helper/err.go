package helper

import "log"

func NoErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
