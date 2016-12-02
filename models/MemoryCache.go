package models

import (
	"fmt"
	c "github.com/patrickmn/go-cache"
	"ms/sun/base"
	"ms/sun/ds"
	"time"
)

type _StoreImpl int

var Store _StoreImpl

func init() {
	//Cacher = c.New(time.Second*6*3600, time.Second*60)
}


