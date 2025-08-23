package linkedlist

import "testing"

//116. 填充每个节点的下一个右侧节点指针（中等）

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 用map记录n节点的第n层的最右端
var rightList map[int]*Node

func connect(root *Node) *Node {
	rightList = make(map[int]*Node)
	connectNode(0, root, nil)
	return root
}

// 用前序遍历
func connectNode(floor int, left, right *Node) {
	if left == nil {
		return
	}

	if _, ok := rightList[floor]; ok {
		rightList[floor].Next = left
	}

	left.Next = right
	connectNode(floor+1, left.Left, left.Right)

	if right == nil {
		return
	}
	rightList[floor] = right
	connectNode(floor+1, right.Left, right.Right)
	if left.Right != nil && right.Left != nil {
		left.Right.Next = right.Left
	}
}

func TestConnect(t *testing.T) {
	tree := &Node{Val: -1,
		Left:  &Node{Val: 0, Left: &Node{Val: 2, Left: &Node{Val: 6}, Right: &Node{Val: 7}}, Right: &Node{Val: 3, Left: &Node{Val: 8}, Right: &Node{Val: 9}}},
		Right: &Node{Val: 1, Left: &Node{Val: 4, Left: &Node{Val: 10}, Right: &Node{Val: 11}}, Right: &Node{Val: 5, Left: &Node{Val: 12}, Right: &Node{Val: 13}}}}
	connect(tree)
}
