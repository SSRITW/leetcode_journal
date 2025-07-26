package hash

import "testing"

// 3487 删除后的最大子数组元素和 （简单）
// 删除数组的元素，使其是一个不重复，非空的子数组，求最大和
// 不能对其重新排序
func MaxSum(nums []int) (sum int) {
	had := map[int]bool{}
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
		if num < 1 {
			continue
		}
		if _, ok := had[num]; ok {
			continue
		}
		had[num] = true
		sum += num
	}
	if len(had) == 0 {
		return maxNum
	}
	return
}

func TestMaxSum(t *testing.T) {
	testCases := []struct {
		num      []int
		expected int
	}{
		//	{num: []int{1, 2, 3, 4, 5}, expected: 15},
		//	{num: []int{1, 1, 0, 1, 1}, expected: 1},
		//	{num: []int{1, 2, -1, -2, 1, 0, -1}, expected: 3},
		{num: []int{-100}, expected: -100},
	}
	t.Run("MaxSum", func(t *testing.T) {
		for _, cases := range testCases {
			if sum := MaxSum(cases.num); sum != cases.expected {
				t.Fatalf("num:=%v,expected: %d, got: %d", cases.num, cases.expected, sum)
			}
		}
	})
}
