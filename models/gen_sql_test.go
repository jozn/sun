package models

import (
    "testing"
    "ms/sun/base"
    "github.com/jmoiron/sqlx"
)

func Benchmark_Insert(b *testing.B) {
    OnAppStart_Models()
    base.DB, _ = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms?charset=utf8mb4")
    //DB, err = sqlx.Connect("mysql", "root:123456@localhost:3307/ms?charset=utf8mb4")
    //DB.Exec("SET NAMES 'utf8mb4';")
    base.DB.MapperFunc(func(s string) string { return s })

    f:=FollowingListMemberHistory{
    }
    for i := 0; i < b.N; i++ {
        f.Save(base.DB)
    }
}

func Benchmark_UserById(b *testing.B) {

    for i := 0; i < b.N; i++ {
        UserById(base.DB,i)
    }
}

func Benchmark_NewUser_Selector_GetRows(b *testing.B) {

    for i := 0; i < b.N; i++ {
        NewUser_Selector().GetRows(base.DB)
    }
}


