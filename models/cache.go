package models

import (
	c "github.com/patrickmn/go-cache"
	"errors"
	"ms/sun/base"
	"ms/sun/helper"
	"time"
)

type cache struct {
	*c.Cache
}

func NewCache() cache {
	s := c.New(time.Minute*60, time.Minute*1)
	return cache{s}
}

//func (c *cache) ()  {
//    Cache
//}

var Cache cache //has set on_start.go

/////////////////////////////////////////////////////
///////////// CacheModels ///////////////////////////
//Specs: .RemoveSomething() :just remove from cache
type _cacheModels struct {
	cache cache
}

var CacheModels _cacheModels

func (cm *_cacheModels) GetPostById(Id int) (*Post, error) {
	key := "post_" + helper.IntToStr(Id)

	val, ok := cm.cache.Get(key)
	if ok {
		pp, ok := val.(*Post)
		if ok {
			return pp, nil
		}
	}

    p := new(Post)
    err := base.DB.Get(p, " select * from post where Id = ? ", Id)
	if err != nil {
        helper.DebugPrintln("XXX: "+err.Error())
		return nil, errors.New("post not found")
	}

	cm.cache.Set(key, p, 0)

	return p, nil
}

func (cm *_cacheModels) RemovePostById(Id int) {
	key := "post_" + helper.IntToStr(Id)
	cm.cache.Delete(key)
}

func (cm *_cacheModels) GetCommentById(Id int) (*Comment, error) {
	key := "comment_" + helper.IntToStr(Id)

	val, ok := cm.cache.Get(key)
	if ok {
		com2, ok := val.(*Comment)
		if ok {
			return com2, nil
		}
	}

    com:= new(Comment)
    err := base.DB.Get(com, " select * from comments where Id = ? ", Id)
	if err != nil {
		return nil, errors.New("comment not found")
	}

	cm.cache.Set(key, com, 0)

	return com, nil
}

func (cm *_cacheModels) RemoveCommentById(Id int) {
	key := "comment_" + helper.IntToStr(Id)
	cm.cache.Delete(key)
}
