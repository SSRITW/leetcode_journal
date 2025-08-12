package linkedlist

import (
	"fmt"
	"testing"
)

// 86. 分隔链表（中等）
// 比x小的节点放在左边，比x大（或等于）的节点放在右边
func Partition86(head *ListNode, x int) *ListNode {
	l, r := &ListNode{}, &ListNode{}
	lP, rP := l, r
	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val >= x {
			rP.Next = head
			rP = rP.Next
		} else {
			lP.Next = head
			lP = lP.Next
		}
		head = next
	}
	lP.Next = r.Next
	return l.Next
}

func TestPartition86(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	result := Partition86(head, 3)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
