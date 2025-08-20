package linkedlist

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 109.有序链表转换二叉搜索树（中等）

// 通过转换数组找到中间值
func sortedListToBST(head *ListNode) *TreeNode {
	l := make([]int, 0, 10)
	//先转换数组
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}
	return toBST(l)
}

// 数组转二叉搜索树
func toBST(l []int) *TreeNode {
	if len(l) == 0 {
		return nil
	}
	//找到中间位置
	n := len(l)
	mid := n / 2
	t := &TreeNode{Val: l[mid]}
	t.Left = toBST(l[:mid])
	t.Right = toBST(l[mid+1:])
	return t
}

func TestSortedListToBST(t *testing.T) {
	head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	result := sortedListToBST(head)
	fmt.Println(traversal(result))
}

//--------------------------------------------------------------------------

func sortedListToBST2(head *ListNode) *TreeNode {
	return toBST2(head, nil)
}

// 快慢针找中间值
func getMid(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func toBST2(start, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}
	mid := getMid(start, end)
	t := &TreeNode{Val: mid.Val}
	t.Left = toBST2(start, mid)
	t.Right = toBST2(mid.Next, end)
	return t
}

func TestSortedListToBST2(t *testing.T) {
	head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	result := sortedListToBST2(head)
	fmt.Println(traversal(result))
}

// ----------------------------------------------------------------------------------------------------------

var global *ListNode

// 参考官方解法，用中序遍历的思路
func sortedListToBST3(head *ListNode) *TreeNode {
	l := 0
	p := head
	for p != nil {
		p = p.Next
		l++
	}
	global = head
	return toBST3(0, l-1)
}

func toBST3(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	t := &TreeNode{}
	t.Left = toBST3(left, mid-1)
	t.Val = global.Val
	global = global.Next
	t.Right = toBST3(mid+1, right)
	return t
}

func TestSortedListToBST3(t *testing.T) {
	head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	result := sortedListToBST3(head)
	fmt.Println(traversal(result))
}

//----------------------------------------------------------------------------------------------------------

func traversal(tree *TreeNode) (result []int) {
	if tree == nil {
		return nil
	}
	result = append(result, traversal(tree.Left)...)
	result = append(result, tree.Val)
	result = append(result, traversal(tree.Right)...)
	return
}
