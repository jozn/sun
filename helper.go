package main

import (
	"fmt"
	"strconv"
	"time"
)

func devPrintn(p ...interface{}) {
	fmt.Println(p...)
}

func debug(p ...interface{}) {
	if __DEV__ {
		fmt.Println(p...)
	}
}

func seconds(s int) time.Duration {
	return time.Duration(s) * time.Second
}

func minutes(s int) time.Duration {
	return time.Duration(s) * time.Minute
}

func intToStr(i int) string {
	return strconv.Itoa(i)
}
