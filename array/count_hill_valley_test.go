package array

import "testing"

// 2210 找到数组里的峰和谷 (简单)
func countHillValley(nums []int) (count int) {
	numsLen := len(nums)
	var pre, next int
	for i := 1; i < numsLen-1; {
		next = i + 1
		//找到下一个不相等值的下标
		for ; next < numsLen-1 && nums[i] == nums[next]; next++ {

		}
		//判断是否峰或者谷
		if (nums[i] > nums[pre] && nums[i] > nums[next]) || (nums[i] < nums[pre] && nums[i] < nums[next]) {
			count++
		}

		if nums[i] != nums[pre] { //如果与前值不相等，需要移动前值的下标，如果相等则保持原位
			pre = i
			i = next
		} else if nums[i] == nums[next] {
			//i值如果与next的值相等，需要移动到next+1的位置，pre需要移动到next的位置；
			// 反之、 i应该移动到next的位置
			pre = next
			i = next + 1
		} else {
			i = next
		}

	}
	return
}

func TestCountHillValley(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		//{nums: []int{2, 4, 1, 1, 6, 5}, expected: 3},
		//{nums: []int{6, 6, 5, 5, 4, 1}, expected: 0},
		{nums: []int{57, 57, 57, 57, 57, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 85, 85, 85, 86, 86, 86}, expected: 2},
	}

	t.Run("countHillValley", func(t *testing.T) {
		for _, testCase := range testCases {
			if count := countHillValley(testCase.nums); count != testCase.expected {
				t.Errorf("in %v got %v, want %v", testCase.nums, count, testCase.expected)
			}
		}
	})
}
