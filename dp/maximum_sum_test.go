package dp

import (
	"testing"
)

// 1186 从数组中选出子数组（不能跳选，需要连续），选出之后可以选择要不要删除其中一个元素，求子数组删除后（如果有删）的最大和 （中等）
// 子数组不能为空
func maximumSum(arr []int) (ans int) {
	arrLen := len(arr)

	//left=以i结尾的最大元素和
	left := make([]int, arrLen)
	//right=以i开头的最大元素和
	right := make([]int, arrLen)
	left[0], right[arrLen-1], ans = arr[0], arr[arrLen-1], arr[0]

	for i := 1; i < arrLen; i++ {
		//如果之前的数已经小于0，重新开始累加
		left[i] = max(left[i-1], 0) + arr[i]
		right[arrLen-i-1] = max(right[arrLen-i], 0) + arr[arrLen-i-1]
		//这里的ans是未删除情况下的最大值
		ans = max(ans, left[i])
	}

	//通过循环left下标i-1和right下标i+1,找出删除1位的最大值
	for i := 1; i < arrLen-1; i++ {
		ans = max(ans, left[i-1]+right[i+1])
	}

	return
}

// 官方思路
func maximumSum2(arr []int) (ans int) {
	arrLen := len(arr)
	//dp1 = 以i为结尾的子数组的总和
	//dp2 = 以i为结尾，会舍弃一个元素，至于舍弃之前的值还是i值，需要对比判断
	dp1, dp2 := arr[0], 0
	ans = arr[0]
	for i := 1; i < arrLen; i++ {
		//如果之前舍弃的dp2加i下标的数，没有不舍数的dp1(还未加上i的值)大，则要舍弃的是下标i的值
		dp2 = max(dp2+arr[i], dp1)
		//如果原值小于0，舍弃之前的累加值。
		dp1 = max(dp1, 0) + arr[i]
		ans = max(ans, dp1, dp2)
	}
	return
}

func TestMaximumSum(t *testing.T) {
	testCases := []struct {
		num      []int
		expected int
	}{
		{num: []int{1, -2, 0, 3}, expected: 4},
		{num: []int{1, -2, -2, 3}, expected: 3},
		{num: []int{-1, -1, -1, -1}, expected: -1},
		{num: []int{1, -4, -5, -2, 5, 0, -1, 2}, expected: 7},
		{num: []int{8, -1, 6, -7, -4, 5, -4, 7, -6}, expected: 17},
	}
	t.Run("maximumSum", func(t *testing.T) {
		for _, cases := range testCases {
			if sum := maximumSum2(cases.num); sum != cases.expected {
				t.Errorf("num:=%v,expected: %d, got: %d", cases.num, cases.expected, sum)
			}
		}
	})
}
