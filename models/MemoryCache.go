package models

import(
    "time"
c "github.com/patrickmn/go-cache"
)

type _StoreImpl int
type _RowCacheImpl int

var Store _StoreImpl
var RowCache *c.Cache

func init() {
	RowCache = c.New(time.Second*6*3600, time.Second*60)
	//Cacher = c.New(time.Second*6*3600, time.Second*60)
}


