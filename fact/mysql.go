package fact

import (
	//	"encoding/json"
	"fmt"
	. "ms/sun/base"
	"strconv"
	"strings"
	"time"
)

func IsamPlay(c *Action) {
	val := "( 152, 'text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 text1 ', 'chars 1' )"
	var a []string
	for i := 0; i < 1000; i++ {
		a = append(a, val)
	}
	vals := strings.Join(a, ",")
	t2 := time.Now().UnixNano()
	DB.Exec("insert into test2 (Id, Text, Str) values ?", vals)
	r := ("time: " + strconv.Itoa(int((time.Now().UnixNano()-t2)/1e6)))
	fmt.Println("RES:: ", r)
	c.SendJson(r)
}
