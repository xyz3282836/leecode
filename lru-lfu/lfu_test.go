package lrulfu

import (
	"testing"
)

func NewOneLfu() LFUCache {
	lfu := NewLfu(6)
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

	return lfu
}

func TestOneLfu(t *testing.T) {

	lfu := NewOneLfu()
	// list root node is nil 拒绝访问的

	t.Errorf("lfu")
	t.Logf("lfu is %#v", lfu.lists[2])
	t.Logf("lfu is %#v", lfu.lists[2].Front().Prev())
	t.Logf("lfu is %#v", lfu.lists[2].Back().Next())

}
