package models

import (
	"time"

	c "github.com/patrickmn/go-cache"
	"ms/sun/models/x"
)

type _StoreImpl int
type _RowCacheImpl struct {
	holder *c.Cache
}

var Store _StoreImpl
var Dep_RowCache_DEP *c.Cache
var RowCacheIndex *c.Cache
var RowCache *_RowCacheImpl

func init() {
	Dep_RowCache_DEP = c.New(time.Second*6*3600, time.Second*60)
	RowCacheIndex = c.New(time.Second*6*3600, time.Second*60)
	//Cacher = c.New(time.Second*6*3600, time.Second*60)

	RowCache = &_RowCacheImpl{
		holder: c.New(time.Second*6*3600, time.Second*60),
	}
}

//////////////// Chat ////////////////////
func (c *_RowCacheImpl) SetChatByKey(key string, chat *x.Chat) bool {
	if chat == nil {
		return false
	}
	c.holder.Set(key, chat, time.Hour*0)
	return true
}

func (c *_RowCacheImpl) GetChatByChatId(key string) (*x.Chat, bool) {
	o, ok := c.holder.Get(key)
	if ok {
		if obj, ok := o.(*x.Chat); ok {
			return obj, true
		}
	}
	return nil, false
}

/////////////////////////////////////////
