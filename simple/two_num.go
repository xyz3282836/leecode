package simple

func TwoNum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if j, ok := numMap[another]; ok {
			return []int{i, j}
		}
		numMap[nums[i]] = i
	}
	return nil
}
