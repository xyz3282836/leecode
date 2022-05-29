package main

import "log"

func main() {
	/**
	  x ^ 0 = x
	  x ^ 11111……1111 = ~x
	  x ^ (~x) = 11111……1111
	  x ^ x = 0
	  a ^ b = c  => a ^ c = b  => b ^ c = a (交换律)
	  a ^ b ^ c = a ^ (b ^ c) = (a ^ b）^ c (结合律)

	  将 x 最右边的 n 位清零， x & ( ~0 << n )
	  获取 x 的第 n 位值(0 或者 1)，(x >> n) & 1

	  X & 1 == 1 判断是否是奇数(偶数)
	  X & = (X - 1) 将最低位(LSB)的 1 清零, xxxx...100... & xxxx...011... = xxxx...000... 清空了LSB
	  X & -X 得到最低位(LSB)的 1,就是 xxxx...100... & ((~x)(~x)(~x)(~x)011... +1) eq (~x)(~x)(~x)(~x)100... = 100...
	  X & ~X = 0,也就是 xxx... & (~x)(~x)(~x)... = 0

	  a
	  ~a 就是取反，就是 0xff ^ a
	  -a 就是取反+1，就是 (0xff ^ a) +1
	  a>>1,a是奇数，那么就是 (a-1)/2;a是偶数，就是a/2

	  **/
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
