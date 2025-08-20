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

func traversal(tree *TreeNode) (result []int) {
	if tree == nil {
		return nil
	}
	result = append(result, traversal(tree.Left)...)
	result = append(result, tree.Val)
	result = append(result, traversal(tree.Right)...)
	return
}
