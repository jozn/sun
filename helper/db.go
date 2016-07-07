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
