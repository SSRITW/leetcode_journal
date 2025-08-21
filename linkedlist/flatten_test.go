package linkedlist

import (
	"testing"
)

// 114. 二叉树展开为链表(中等)
func flatten(root *TreeNode) {
	flattenToLinkedList(root)
}

func flattenToLinkedList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := root.Right
	root.Right = flattenToLinkedList(root.Left)
	root.Left = nil
	next := root
	for next.Right != nil {
		next = next.Right
	}
	next.Right = flattenToLinkedList(right)
	return root
}

func TestFlatten(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	flatten(tree)
}
