package lru

import (
	"fmt"
	"testing"
)

var c *Cache

func init() {
	c = New(10000, 10)
}

func TestAdd(t *testing.T) {
	c.Add("1", "test ==")
}

func TestGet(t *testing.T) {
	c.Add("1", "test ==")
	fmt.Println(c.Get("1").(string))

	c.Add("1", "test")
	fmt.Println(c.Get("1").(string))
}

func TestRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		c.Add(i, i)
	}

	c.Range(false, func(key, value interface{}) {
		fmt.Println("key=", key, "value=", value)
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.Add(i, "testfsf")
	}
}

func BenchmarkGet(b *testing.B) {
	c.Add(1, "testfsf")
	for i := 0; i < b.N; i++ {
		c.Get(2)
	}
}

func BenchmarkDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.Add(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Delete(i)
	}
}

func BenchmarkRange(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		c.Add(i, i)
	}

	b.ResetTimer()
	c.Range(true, func(key, value interface{}) {})
}
