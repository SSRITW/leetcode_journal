package linkedlist

import (
	"fmt"
	"testing"
)

// 143. 重排链表 (中等)

// 用list去记录每个节点的，前后双指针实现重新排序
func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	list := make([]*ListNode, 0, 0)
	for p := head; p != nil; p = p.Next {
		list = append(list, p)
	}

	//至少有三个值才能交互，所以i<j-1
	for i, j := 0, len(list)-1; i < j-1; i, j = i+1, j-1 {
		next := list[i].Next
		//断开连接（新尾节点）
		list[j-1].Next = nil
		//拼接新节点
		list[i].Next, list[j].Next = list[j], next
	}
}

func TestReorderList(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	//list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reorderList(list)
	p := list
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

//--------------------------------------------------------

func reorderList2(head *ListNode) {
	mid := midNode(head)
	right := reverseList(mid.Next)
	mid.Next = nil
	//穿插节点
	for n, m := head, right; n != nil && m != nil; {
		nextN, nextM := n.Next, m.Next
		n.Next = m
		n.Next.Next = nextN
		n, m = nextN, nextM
	}
}

// 快慢指针，获取中间节点
func midNode(head *ListNode) *ListNode {
	n, m := head, head
	for m != nil && m.Next != nil {
		n, m = n.Next, m.Next.Next
	}
	return n
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func TestReorderList2(t *testing.T) {
	//list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reorderList2(list)
	p := list
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
