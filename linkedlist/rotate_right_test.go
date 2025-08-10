package linkedlist

import (
	"fmt"
	"testing"
)

//61. 旋转链表(中等)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	nLen := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		nLen++
	}
	k %= nLen

	newHeadIdx := nLen - k

	lessNode := &ListNode{}
	headNode := &ListNode{}
	headP := headNode
	lessP := lessNode

	p := head
	for i := 0; i < nLen; i++ {
		next := p.Next
		p.Next = nil
		if newHeadIdx > i {
			lessP.Next = p
			lessP = lessP.Next
		} else {
			headP.Next = p
			headP = headP.Next
		}
		p = next
	}

	headP.Next = lessNode.Next

	return headNode.Next
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
