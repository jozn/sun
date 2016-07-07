package actions

import (
	"fmt"
	// "reflect"
	"path/filepath"
	"strconv"
	"strings"
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

func int64ToStr(i int64) string {
	return strconv.Itoa(int(i))
}

func now() int {
	return int(time.Now().Unix())
}

func intsToSqlIn(ins []int) string {
	// sql := ""
	sins := make([]string, len(ins))
	for i := 0; i < len(ins); i++ {
		sins[i] = strconv.Itoa(ins[i])
	}
	return strings.Join(sins, ", ")
}

func devNoErr(err error) {
	if __DEV__ {
		noErr(err)
	}

}

func fileExt(path string) string {
	return filepath.Ext(path)
}

//deprecated
