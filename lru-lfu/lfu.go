package lrulfu

import "container/list"

type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func NewLfu(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	value, ok := lfu.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	lfu.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := lfu.lists[currentNode.frequency]; !ok {
		lfu.lists[currentNode.frequency] = list.New()
	}
	newList := lfu.lists[currentNode.frequency]
	newNode := newList.PushFront(currentNode)
	lfu.nodes[key] = newNode
	if currentNode.frequency-1 == lfu.min && lfu.lists[currentNode.frequency-1].Len() == 0 {
		lfu.min++
	}
	return currentNode.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}
	// 如果存在，更新访问次数
	if currentValue, ok := lfu.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		lfu.Get(key)
		return
	}
	// 如果不存在且缓存满了，需要删除
	if lfu.capacity == len(lfu.nodes) {
		currentList := lfu.lists[lfu.min]
		backNode := currentList.Back()
		delete(lfu.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}
	// 新建结点，插入到 2 个 map 中
	lfu.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := lfu.lists[1]; !ok {
		lfu.lists[1] = list.New()
	}
	newList := lfu.lists[1]
	newNode := newList.PushFront(currentNode)
	lfu.nodes[key] = newNode
}
