package maths

import (
	"testing"
)

// 2683 相邻值的按位异或 （中等）
func doesValidArrayExist(derived []int) bool {
	return valid(derived, 0) || valid(derived, 1)
}

// 用假定的original去推算original列表，然后验证original的下标0^最后一位，是否等于derived的最后一位去判断是否合法
func valid(s []int, start int) bool {
	o := start
	for i := 0; i < len(s)-1; i++ {
		o = o ^ s[i]
	}
	return (o ^ start) == s[len(s)-1]
}

// 根据题目公式：
// derived[i] = original[i] ^ original[i+1]
// 如i==长度-1 derived[i] = original[i] ^ original[0]
// 得出 derived[0] ^ derived[1] ^ ……derived[n-1] = (original[0] ^ original[1]) ^ ( original[1] ^ original[2])^ ……(original[n-1] ^ original[0]) = 0
// 所以 derived 数组的总和 肯定等于 0
func doesValidArrayExist2(derived []int) bool {
	sum := 0
	for i := 0; i < len(derived); i++ {
		sum += derived[i]
	}
	return sum%2 == 0
}

func TestDoesValidArrayExist(t *testing.T) {
	testCases := []struct {
		derived  []int
		expected bool
	}{
		{derived: []int{1, 1, 0}, expected: true},
		{derived: []int{1, 1}, expected: true},
		{derived: []int{1, 0}, expected: false},
		{derived: []int{0, 1, 1}, expected: true},
	}
	t.Run("doesValidArrayExist", func(t *testing.T) {
		for _, cases := range testCases {
			if result := doesValidArrayExist2(cases.derived); result != cases.expected {
				t.Errorf("num:=%v,expected: %v, got: %v", cases.derived, cases.expected, result)
			}
		}
	})
}
