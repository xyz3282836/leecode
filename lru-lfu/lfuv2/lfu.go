package lfuv2

import "container/heap"

type LFUCache struct {
	capacity int
	pq       PriorityQueue
	hash     map[int]*Item
	counter  int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value     int // The value of the item; arbitrary.
	key       int
	frequency int // The priority of the item in the queue.
	count     int // use for evicting the oldest element
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].frequency == pq[j].frequency {
		return pq[i].count < pq[j].count
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, frequency int, count int) {
	item.value = value
	item.count = count
	item.frequency = frequency
	heap.Fix(pq, item.index)
}

func NewLfu(capacity int) LFUCache {
	lfu := LFUCache{
		pq:       PriorityQueue{},
		hash:     make(map[int]*Item, capacity),
		capacity: capacity,
	}
	return lfu
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}
	if item, ok := lfu.hash[key]; ok {
		lfu.counter++
		lfu.pq.update(item, item.value, item.frequency+1, lfu.counter)
		return item.value
	}
	return -1
}
func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}
	lfu.counter++
	// 如果存在，增加 frequency，再调整堆
	if item, ok := lfu.hash[key]; ok {
		lfu.pq.update(item, value, item.frequency+1, lfu.counter)
		return
	}
	// 如果不存在且缓存满了，需要删除。在 hashmap 和 pq 中删除。
	if len(lfu.pq) == lfu.capacity {
		item := heap.Pop(&lfu.pq).(*Item)
		delete(lfu.hash, item.key)
	}
	// 新建结点，在 hashmap 和 pq 中添加。
	item := &Item{
		value: value,
		key:   key,
		count: lfu.counter,
	}
	heap.Push(&lfu.pq, item)
	lfu.hash[key] = item
}
