package array

import (
	"testing"
)

//3479. 水果成篮 III（中等）
//未掌握知识点   线段树

// 暴力解法
func numOfUnplacedFruits(fruits []int, baskets []int) (result int) {
	in := make(map[int]bool)
	for i := 0; i < len(fruits); i++ {
		find := false
		for j := 0; j < len(baskets); j++ {
			if in[j] || baskets[j] < fruits[i] {
				continue
			}
			in[j] = true
			find = true
			break
		}
		if !find {
			result++
		}
	}
	return
}

func TestNumOfUnplacedFruits(t *testing.T) {
	testCases := []struct {
		fruits  []int
		baskets []int
		result  int
	}{
		{fruits: []int{4, 2, 5}, baskets: []int{3, 5, 4}, result: 1},
		{fruits: []int{3, 6, 1}, baskets: []int{6, 4, 7}, result: 0},
	}

	t.Run("numOfUnplacedFruits", func(t *testing.T) {
		for _, cases := range testCases {
			if result := numOfUnplacedFruits(cases.fruits, cases.baskets); result != cases.result {
				t.Errorf("fruits:%v,baskets:%v,expected: %v, got: %v", cases.fruits, cases.baskets, cases.result, result)
			}
		}
	})
}
