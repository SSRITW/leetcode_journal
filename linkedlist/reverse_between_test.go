package linkedlist

import (
	"fmt"
	"testing"
)

//92. 反转链表 II(中等)

// 翻转指定的left和right下标间的节点
func ReverseBetween92(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	var leftPre *ListNode
	temp := &ListNode{Next: head}
	p := temp
	for i := 1; p.Next != nil; i++ {
		if i == left {
			leftPre = p
		} else if i == right {
			next := p.Next.Next
			//翻转前先断开链接
			p.Next.Next = nil
			newHead, newTail := reverse(leftPre.Next)
			leftPre.Next = newHead
			newTail.Next = next
			break
		}
		p = p.Next
	}
	return temp.Next
}

// 翻转链表，返回新的头部和尾部
func reverse(head *ListNode) (*ListNode, *ListNode) {
	now := head
	var prev *ListNode
	for now != nil {
		next := now.Next
		now.Next = prev
		now, prev = next, now
	}
	return prev, head
}

func TestReverseBetween92(t *testing.T) {
	//head := &ListNode{Val: 3, Next: &ListNode{Val: 5}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := ReverseBetween92(head, 2, 4)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
