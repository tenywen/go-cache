package cache

import (
	"sync"
)

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

func (group *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	group.mu.Lock()
	if group.calls == nil {
		group.calls = make(map[string]*call)
	}

	if c, ok := group.calls[key]; ok {
		group.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := &call{}
	c.wg.Add(1)
	group.calls[key] = c
	group.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	group.mu.Lock()
	delete(group.calls, key)
	group.mu.Unlock()
	return c.val, c.err
}
