package linkedlist

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

//-----------------------------------------------------------------------
