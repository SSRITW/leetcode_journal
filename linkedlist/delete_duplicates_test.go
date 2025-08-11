package linkedlist

import (
	"fmt"
	"testing"
)

//82. 删除排序链表中的重复元素 II(中等)

func deleteDuplicates(head *ListNode) *ListNode {
	temp := &ListNode{}
	tempP := temp
	dummy := &ListNode{Val: -10000, Next: head}
	l := dummy
	p := dummy.Next
	for p != nil {
		//对比前后，不相同就符合条件
		if p.Val != l.Val && (p.Next == nil || p.Val != p.Next.Val) {
			tempP.Next = p
			tempP = tempP.Next
		} else { //如果这次不相同，要断开链接
			tempP.Next = nil
		}
		l = p
		p = p.Next
	}
	return temp.Next
}

func TestDeleteDuplicates(t *testing.T) {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}
	result := deleteDuplicates(head)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
