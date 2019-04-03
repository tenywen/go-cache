package cache

import (
	"lru"
)

const (
	Prefix = "CACHE_"
)

type cache struct {
	hotCache  *lru.Cache // in other servers
	mainCache *lru.Cache // local server
}

func newCache(timeout int64) *cache {
	cache := &cache{
		hotCache:  lru.New(64, timeout),
		mainCache: lru.New(1024, timeout),
	}
	return cache
}

func (cache *cache) get(key interface{}) (interface{}, bool) {
	v := cache.mainCache.Get(key)
	if v != nil {
		return v, true
	}

	v = cache.hotCache.Get(key)
	if v != nil {
		return v, true
	}
	return nil, false
}
