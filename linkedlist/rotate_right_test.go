package linkedlist

import (
	"fmt"
	"testing"
)

// 61. 旋转链表(中等)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	nLen := 1
	temp := head
	for temp.Next != nil {
		temp = temp.Next
		nLen++
	}
	k %= nLen
	if k == 0 {
		return head
	}
	//头尾相接
	temp.Next = head
	newTailIdx := nLen - k
	for newTailIdx > 0 {
		temp = temp.Next
		newTailIdx--
	}
	ans := temp.Next
	temp.Next = nil
	return ans
}

func TestRotateRight(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := rotateRight(head, 2)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
