package simple

import "testing"

func TestTwoNum(t *testing.T) {
	nums := []int{3, 6, 4, 8}
	target := 9
	ret := TwoNum(nums, target)
	t.Logf("%v", ret)
}
