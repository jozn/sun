package main

import (
    "ms/sun/base"
    "testing"
    . "ms/sun/models"
    "sync"
)

var onc sync.Once
func prepareTest() {
    onc.Do(func() {startApp()})
}

func Benchmark_sql_Insert(b *testing.B) {
    prepareTest()

    for i := 0; i < b.N; i++ {
        f:=FollowingListMemberHistory{}
        f.Save(base.DB)
    }
}

func Benchmark__sql_UserById(b *testing.B) {
    prepareTest()
    for i := 0; i < b.N; i++ {
        UserById(base.DB,i)
    }
}

func Benchmark__sql_NewUser_Selector_GetRows(b *testing.B) {
    prepareTest()
    for i := 0; i < b.N; i++ {
        NewUser_Selector().GetRows2(base.DB)
    }
}

