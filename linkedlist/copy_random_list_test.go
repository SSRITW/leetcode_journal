package linkedlist

import "testing"

//138. 随机链表的复制(中等)

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	p := head
	//新的链表
	newHead := &RandomNode{}
	newP := newHead
	//旧链表的节点指向新链表节点
	pMap := map[*RandomNode]*RandomNode{}
	for p != nil {
		newP.Next = &RandomNode{Val: p.Val}
		pMap[p] = newP.Next
		newP = newP.Next
		p = p.Next
	}
	p = head
	newP = newHead.Next
	for p != nil {
		//通过指针得到新节点
		newP.Random = pMap[p.Random]
		newP = newP.Next
		p = p.Next
	}
	return newHead.Next
}

// 使用递归方式实现，不变的是，还是用hash的key记录节点指针。

var nodeMap map[*RandomNode]*RandomNode

func copyRandomList2(head *RandomNode) *RandomNode {
	nodeMap = map[*RandomNode]*RandomNode{}
	return deepCopy(head)
}

func deepCopy(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}
	//已经拷贝过
	if _, ok := nodeMap[head]; ok {
		return nodeMap[head]
	}
	newHead := &RandomNode{Val: head.Val}
	nodeMap[head] = newHead
	newHead.Next = deepCopy(head.Next)
	newHead.Random = deepCopy(head.Random)
	return newHead
}

// -----------------------------------------------------------------------
// 参考官方解题思路
func copyRandomList3(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}
	// 直接在老链表里创建新节点
	for p := head; p != nil; p = p.Next.Next {
		next := p.Next
		p.Next = &RandomNode{Val: p.Val, Next: next}
	}
	//赋值随机节点变量
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	//拼接出新的链表
	newHead := head.Next
	for p := head; p != nil; p = p.Next {
		newP := p.Next
		//将之前拼接的节点移除
		p.Next = p.Next.Next
		if newP.Next != nil {
			newP.Next = newP.Next.Next
		}
	}
	return newHead
}

func TestCopyRandomList3(t *testing.T) {
	tree := &RandomNode{Val: 1, Next: &RandomNode{Val: 2}}
	tree.Random = tree
	tree.Next.Random = tree
	copyRandomList3(tree)
}
