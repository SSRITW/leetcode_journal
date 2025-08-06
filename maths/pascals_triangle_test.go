package maths

import "testing"

// 118. 杨辉三角（简单）
func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
				continue
			}
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func TestGenerate(t *testing.T) {
	testCases := []struct {
		numRows  int
		expected [][]int
	}{
		{numRows: 5, expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{numRows: 1, expected: [][]int{{1}}},
	}

	t.Run("generate", func(t *testing.T) {
		for _, cases := range testCases {
			result := generate(cases.numRows)
			if len(result) != len(cases.expected) {
				t.Errorf("num:=%v,expected: %v, got: %v", cases.numRows, cases.expected, result)
				continue
			}
			for i := 0; i < len(result); i++ {
				if len(result[i]) != len(cases.expected[i]) {
					t.Errorf("num:=%v,expected: %v, got: %v", cases.numRows, cases.expected, result)
					break
				}
				for j := 0; j < len(result[i]); j++ {
					if result[i][j] != cases.expected[i][j] {
						t.Errorf("num:=%v,expected: %v, got: %v", cases.numRows, cases.expected, result)
						break
					}
				}
			}
		}
	})
}

// 119. 杨辉三角II（简单） 获取指定行
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {

	}
	return ans
}
