package main

import "log"

func main() {
	aa := []int{4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	log.Println(binarySearchMatrix(aa, 6))
	var a = int(28)

	log.Println(a & (-a))
	log.Println(a & ((0xff ^ a) + 1))
	log.Printf("%b", a&(-a))

}

func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
