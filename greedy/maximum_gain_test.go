package greedy

import (
	"strings"
	"testing"
)

func MaximumGain(s string, x int, y int) (ans int) {
	cutNum := []string{"ba", "ab"}
	p := map[string]int{"ba": y, "ab": x}
	if x > y {
		cutNum = []string{"ab", "ba"}
	}

	for i, reTry := 0, 0; ; i++ {
		thisCutNum := cutNum[reTry]
		lastIdx := strings.LastIndex(s, thisCutNum)
		if lastIdx < 0 {
			if reTry > 0 {
				break
			}
			reTry++
			continue
		}
		ans += p[thisCutNum]
		s = s[:lastIdx] + s[lastIdx+2:]
		reTry = 0
	}
	return
}

func MaximumGain2(s string, x int, y int) (ans int) {
	a, b := 'a', 'b'
	if x < y {
		a, b = b, a
		x, y = y, x
	}
	countA, countB := 0, 0
	for _, c := range s {
		if c == a {
			countA++
		} else if c == b {
			if countA > 0 {
				ans += x
				countA--
			} else {
				countB++
			}
		} else {
			ans += min(countA, countB) * y
			countA, countB = 0, 0
		}
	}
	ans += min(countA, countB) * y
	return
}

func TestMaximumGain(t *testing.T) {
	type TestCase struct {
		s        string
		x        int
		y        int
		expected int
	}

	testCases := []TestCase{
		{s: "cdbcbbaaabab", x: 4, y: 5, expected: 19},
		{s: "aabbaaxybbaabb", x: 5, y: 4, expected: 20},
	}
	t.Run("MaximumGain", func(t *testing.T) {
		for _, testCase := range testCases {
			ans := MaximumGain(testCase.s, testCase.x, testCase.y)
			if ans != testCase.expected {
				t.Errorf("ERROR: For s:%s x:%d,y:%d, expected %d, got %d", testCase.s, testCase.x, testCase.y, testCase.expected, ans)
			}
		}
	})

	t.Run("MaximumGain2", func(t *testing.T) {
		for _, testCase := range testCases {
			ans := MaximumGain2(testCase.s, testCase.x, testCase.y)
			if ans != testCase.expected {
				t.Errorf("ERROR: For s:%s x:%d,y:%d, expected %d, got %d", testCase.s, testCase.x, testCase.y, testCase.expected, ans)
			}
		}
	})
}
