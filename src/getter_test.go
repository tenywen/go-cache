package cache

import (
	"fmt"
	"sync"
	//nats "github.com/nats-io/go-nats"
	"testing"
	//"time"
)

func TestGetter(t *testing.T) {
	getter := New("127.0.0.1:4222", 50, []string{"127.0.0.1:4222"}, func(key string) (interface{}, error) {
		return []byte("ddddddddddd"), nil
	}, 60)

	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(getter.Get("test"))
		}()
	}
	wg.Wait()
}

func BenchmarkGetter(b *testing.B) {
	getter := New("127.0.0.1:4222", 50, []string{"127.0.0.1:4222"}, func(key string) (interface{}, error) {
		return []byte("ddddddddddd"), nil
	}, 60)

	b.N = 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getter.Get("test")
	}
}
