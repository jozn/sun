package main

import (
	"testing"
    "ms/sun/store"
    "strconv"
)

//var onc sync.Once

func prepareTest1() {
	onc.Do(func() { startApp() })
}

func Benchmark_redis_mass_int(b *testing.B) {
    prepareTest()

    p := store.GetRedisPool()
    M:=1000
    arr :=make([]interface{},M)
    for i := 1; i < M; i++{
        arr[i]=i
    }

    arr[0]="list"

	for i := 0; i < b.N; i++ {
        p.Cmd("LPUSH",arr...)
	}
}

func Benchmark_redis_mass_Many_int(b *testing.B) {
    prepareTest()

    p := store.GetRedisPool()
    M:=1000
    arr :=make([]interface{},M)
    for i := 1; i < M; i++{
        arr[i]=i
    }

    for i := 0; i < b.N; i++ {
        arr[0]="l"+strconv.Itoa(i)
        p.Cmd("LPUSH",arr...)
    }
}

