package helper

import (
	"strconv"
	"strings"
)

func IntsToSqlIn(ins []int) string {
	// sql := ""
	sins := make([]string, len(ins))
	for i := 0; i < len(ins); i++ {
		sins[i] = strconv.Itoa(ins[i])
	}
	return strings.Join(sins, ",")
}

func DbQuestionForSqlIn(size int) string {
	s := strings.Repeat("?,", size)
	s = s[0 : len(s)-1] //remove last ','
	return s
}
