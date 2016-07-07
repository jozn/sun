package main

import (
	// "math"
	"math/rand"
	"strconv"
	"time"
)

func playRedis() {
	for i := 0; i < 10000; i++ {
		redisServer.Do("GET", "val-"+intToStr(i), i)
	}
}

func playRedis2() {
	for i := 0; i < 10000; i++ {
		redisServer.Do("RPUSH", "user20", i)
	}
}

func redisAction1(c *Action) {
	// playRedis()
	playRedis2()
	c.SendText("redis")

}

func redisAction2(c *Action) {
	// playRedis()
	// playRedis2()
	c.SendText(intToStr(genrateUniqeIdTime()))
	// c.SendText(randIntStrBySize(10))
	// c.SendText(int64ToStr(time.Now().UnixNano()/1e6) + "14526")

}

func genrateUniqeIdTime() int {
	timeMS := time.Now().UnixNano() / 1e6 //millisecons 13digit
	rndStr := randIntStrBySize(3)         // 3digit
	ts := int64ToStr(timeMS)
	strTime := ts + rndStr // 13 + 3 = 16 javascript maximum real integrer value = 2**53 -1 == 16 digit  Number.MAX_SAFE_INTEGER
	debug("strTime: ", strTime)
	return strToInt(strTime)
}

func randIntBySize(n int) int {
	return strToInt(randIntStrBySize(n))
}

func randIntStrBySize(n int) string {
	var str = ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(rand.Intn(9) + 1) //not zero -will bug for first digit
	}
	return str
}

func strToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 62)
	return i
}
func strToInt(str string) int {
	i, _ := strconv.ParseInt(str, 10, 62)
	return int(i)
}
