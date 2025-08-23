package linkedlist

import "testing"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//116. 填充每个节点的下一个右侧节点指针（中等）

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
}

func TestConnect(t *testing.T) {
	tree := &Node{Val: -1,
		Left:  &Node{Val: 0, Left: &Node{Val: 2, Left: &Node{Val: 6}, Right: &Node{Val: 7}}, Right: &Node{Val: 3, Left: &Node{Val: 8}, Right: &Node{Val: 9}}},
		Right: &Node{Val: 1, Left: &Node{Val: 4, Left: &Node{Val: 10}, Right: &Node{Val: 11}}, Right: &Node{Val: 5, Left: &Node{Val: 12}, Right: &Node{Val: 13}}}}
	connect(tree)
}

//-----------------------------------------------------------------------------------------

//117. 填充每个节点的下一个右侧节点指针 II(中等)
//跟116差不多，不一样的是，117不是完美平衡树

// 用map记录n节点的第n层的最右端
var prevMap map[int]*Node

func connect117(root *Node) *Node {
	prevMap = make(map[int]*Node)
	connectNode117(0, root, nil)
	return root
}

// 用前序遍历
func connectNode117(floor int, left, right *Node) {
	prev := prevMap[floor]
	if left != nil {
		if prev != nil {
			prev.Next = left
		}
		connectNode117(floor+1, left.Left, left.Right)
		prev = left
		prevMap[floor] = left
	}

	if right == nil {
		return
	}
	prevMap[floor] = right
	if prev != nil {
		prev.Next = right
	}
	connectNode117(floor+1, right.Left, right.Right)
}

func TestConnect117(t *testing.T) {
	tree := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Right: &Node{Val: 7}}}
	connect117(tree)
}
