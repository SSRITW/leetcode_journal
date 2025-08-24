package linkedlist

import "testing"

// 141. 判断是否环形链表(简单)-----------------------------------
// 用hash表去判断
func hasCycle(head *ListNode) bool {
	has := make(map[*ListNode]bool)
	for head != nil {
		if has[head] {
			return true
		}
		has[head] = true
		head = head.Next
	}
	return false
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		head     *ListNode
		expected bool
	}{
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, expected: true},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, expected: true},
		{head: &ListNode{Val: 1}, expected: false},
	}
	testCases[0].head.Next.Next = testCases[0].head.Next
	testCases[1].head.Next = testCases[1].head
	t.Run("hasCycle", func(t *testing.T) {
		for i, cases := range testCases {
			if result := hasCycle(cases.head); result != cases.expected {
				t.Fatalf("%v,expected: %v, got: %v", i, cases.expected, result)
			}
		}
	})
}

// ----------------------------------------------------------
// 使用快慢针判断
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}

	for slow, fast := head, head.Next; slow != fast; slow, fast = slow.Next, fast.Next.Next {
		if fast == nil || fast.Next == nil {
			return false
		}
	}
	return true
}

func TestHasCycle2(t *testing.T) {
	testCases := []struct {
		head     *ListNode
		expected bool
	}{
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, expected: true},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, expected: true},
		{head: &ListNode{Val: 1}, expected: false},
	}
	testCases[0].head.Next.Next = testCases[0].head.Next
	testCases[1].head.Next = testCases[1].head
	t.Run("hasCycle2", func(t *testing.T) {
		for i, cases := range testCases {
			if result := hasCycle2(cases.head); result != cases.expected {
				t.Fatalf("%v,expected: %v, got: %v", i, cases.expected, result)
			}
		}
	})
}
