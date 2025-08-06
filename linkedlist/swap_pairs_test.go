package linkedlist

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 24. 两两交换链表中的节点(中等)
// 0  1   2   3  4 ->  0  2  1 3  4 ->  0 2  1  4  3
func SwapPairs(head *ListNode) *ListNode {
	ans := &ListNode{Next: head}
	now := ans
	//每次交换两个
	for now.Next != nil && now.Next.Next != nil {
		n1 := now.Next
		n2 := now.Next.Next
		now.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		now = n1
	}
	return ans.Next
}

func TestSwapPairs(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	result := SwapPairs(head)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
