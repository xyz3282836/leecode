package lrulfu

import (
	"testing"
)

func NewOneLfu() LFUCache {
	return NewLfu(6)
}

func TestOneLfu(t *testing.T) {

	lfu := NewOneLfu()
	lfu.Put(1, 101)

	lfu.Put(2, 102)
	lfu.Put(2, 102)

	lfu.Put(22, 202)
	lfu.Put(22, 202)

	lfu.Put(3, 103)
	lfu.Put(3, 103)
	lfu.Put(3, 103)

	lfu.Put(33, 303)
	lfu.Put(33, 303)
	lfu.Put(33, 303)

	lfu.Put(333, 333)
	lfu.Put(333, 333)
	lfu.Put(333, 333)

	lfu.Put(99, 999)
	lfu.Put(99, 999)
	lfu.Put(99, 999)

	t.Errorf("lfu")
	t.Logf("lfu is %#v", lfu.lists[2])

}
