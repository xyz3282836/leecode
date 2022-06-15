package binaryindextree

import "log"

// BinaryIndexedTree define
type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

// Init define
func (bit *BinaryIndexedTree) Init(nums []int) {
	bit.tree, bit.capacity = make([]int, len(nums)+1), len(nums)+1
	log.Printf("init nums is %v tree is %v cap is %d \r\n", nums, bit.tree, bit.capacity)
	for i := 1; i <= len(nums); i++ {
		bit.tree[i] += nums[i-1]
		log.Printf("i(%d) bit.tree[%d] += nums[%d] %d \r\n", i, i, i-1, bit.tree[i])
		for j := i - 2; j >= i-lowbit(i); j-- {
			bit.tree[i] += nums[j]
			log.Printf("  j=i-2 j=%d bit.tree[%d] + nums[%d](%d) = %d \r\n", j, i, j, nums[j], bit.tree[i])
		}
	}
}

func lowbit(x int) int {
	return x & (-x)
}

// Add define
func (bit *BinaryIndexedTree) Add(index int, val int) {
	for index <= bit.capacity {
		bit.tree[index] += val
		index += lowbit(index)
	}
}

// Query define
func (bit *BinaryIndexedTree) Query(index int) int {
	sum := 0
	for index >= 1 {
		sum += bit.tree[index]
		log.Printf("query index %d is sum = sum + bit.tree[%d] is %v \r\n", index, bit.tree[index], sum)
		index -= lowbit(index)

	}
	return sum
}
