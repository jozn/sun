package store

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
)

const REDIS_SORTED_LIST_ADD = "ZADD"
const REDIS_SORTED_LIST_REMOVE_RANGE_SCORE = "ZREMRANGEBYSCORE"
const REDIS_SORTED_LIST_GET_RANGE_INDEX = "ZRANGE" //by index

func init() {
	//GetRedisPool()
}

var redisPool *pool.Pool
var isRedisCon bool = false

func GetRedisPool() *pool.Pool {
	if isRedisCon {
		return redisPool
	}
	var err error
	redisPool, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		fmt.Println("redis failed")
		return nil
	}
	fmt.Println("redis online")
	isRedisCon = true
	return redisPool
}

func redisInit() {
	var err error
	redisPool, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		fmt.Println("redis failed")
		return
	}
	fmt.Println("redis online")
	isRedisCon = true
	// In another go-routine

	conn, err := redisPool.Get()
	if err != nil {
		// handle error
	}

	if conn.Cmd("SET", "KEYYY", "Val").Err != nil {
		// handle error
	}

	redisPool.Put(conn)
}