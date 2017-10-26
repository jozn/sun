package base

import (
	"github.com/jmoiron/sqlx"
	"ms/sun/helper"
)

var DB *sqlx.DB
var __DEV__ bool = true

type AppErr error

func DefultConnectToMysql() {
	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/ms?charset=utf8mb4")
	helper.NoErr(err)
	DB.MapperFunc(func(s string) string { return s })
}
