package lrulfu

import (
	"testing"
)

func NewOneLru() LRUCache {
	return NewLru(4)
}

func TestOne(t *testing.T) {
	t.Run("get test", func(t *testing.T) {
		lru := NewOneLru()
		lru.Put(1, 101)
		lru.Put(2, 102)
		node := lru.Get(2)
		if node == 101 {
			t.Errorf("err")
		}

	})

}
