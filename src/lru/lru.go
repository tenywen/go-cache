package lru

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	maxEntries int
	timeout    int64
	mu         sync.RWMutex
	ll         *list.List
	cache      map[interface{}]*list.Element
}

type entry struct {
	timeout    int64
	key, value interface{}
}

func New(maxEntries int, timeout int64) *Cache {
	return &Cache{
		maxEntries: maxEntries,
		timeout:    timeout,
	}
}

func (cache *Cache) Add(key, value interface{}) {
	cache.mu.Lock()
	if cache.cache == nil {
		cache.cache = make(map[interface{}]*list.Element)
		cache.ll = &list.List{}
	}

	now := time.Now().Unix()
	if elem, ok := cache.cache[key]; ok {
		elem.Value.(*entry).value = value
		elem.Value.(*entry).timeout = now + cache.timeout
		cache.ll.MoveToFront(elem)
		cache.mu.Unlock()
		return
	}

	elem := cache.ll.PushFront(&entry{now, key, value})
	cache.cache[key] = elem
	if cache.maxEntries > 0 && cache.ll.Len() > cache.maxEntries {
		cache.removeOldest()
	}
	cache.mu.Unlock()
}

func (cache *Cache) Get(key interface{}) interface{} {
	cache.mu.RLock()
	if cache.cache == nil {
		cache.mu.RUnlock()
		return nil
	}

	now := time.Now().Unix()
	if elem, ok := cache.cache[key]; ok {
		if elem.Value.(*entry).timeout > now {
			cache.ll.MoveToFront(elem)
			cache.mu.RUnlock()
			return elem.Value.(*entry).value
		}
		// del
		delete(cache.cache, key)
		cache.ll.Remove(elem)
	}
	cache.mu.RUnlock()
	return nil
}

func (cache *Cache) Delete(key interface{}) interface{} {
	cache.mu.Lock()
	if cache.cache == nil {
		cache.mu.Unlock()
		return nil
	}
	v := cache.del(key)
	cache.mu.Unlock()
	return v
}

func (cache *Cache) del(key interface{}) interface{} {
	if elem, ok := cache.cache[key]; ok {
		delete(cache.cache, key)
		cache.ll.Remove(elem)
		return elem.Value.(*entry).value
	}
	return nil
}

func (cache *Cache) removeOldest() {
	if cache.cache == nil {
		return
	}
	elem := cache.ll.Remove(cache.ll.Back())
	delete(cache.cache, elem.(*entry).key)
	elem = nil
}

func (cache *Cache) GetKeys() []interface{} {
	cache.mu.RLock()
	keys := make([]interface{}, 0, len(cache.cache))
	for k, _ := range cache.cache {
		keys = append(keys, k)
	}
	cache.mu.RUnlock()
	return keys
}

//----------------------------------------------------- can't delete
func (cache *Cache) Range(isDelete bool, f func(key, value interface{})) {
	cache.mu.Lock()
	for key, value := range cache.cache {
		f(key, value.Value.(*entry).value)
		if isDelete {
			cache.del(key)
		}
	}
	cache.mu.Unlock()
}
