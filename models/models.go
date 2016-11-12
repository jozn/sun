package models

import (
	//	"github.com/jmoiron/sqlx"
	cacheDrive "github.com/pmylund/go-cache"
	"time"
)

//var DB *sqlx.DB
var __DEV__ bool
var cashe *cacheDrive.Cache

func init() {
	__DEV__ = true
	//	DB, err := sqlx.Connect("mysql", "root:123456@/ms3?charset=utf8mb4")
	//	 DB, err := sqlx.Connect("mysql", "root:123456@/ms3?charset=ascii")
	//	noErr(err)
	//		_ = DB
	cashe = cacheDrive.New(5*time.Minute, 30*time.Second)

}
