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

	//存储新链表的位置
	nodeList := make([]*RandomNode, 0)
	//旧链表的下标
	pMap := map[*RandomNode]int{}
	for i := 0; p != nil; i++ {
		pMap[p] = i
		newP.Next = &RandomNode{Val: p.Val}
		nodeList = append(nodeList, newP.Next)
		newP = newP.Next
		p = p.Next
	}
	p = head
	newP = newHead.Next
	for p != nil {
		//通过指针获得对应的下标，得到新节点
		if _, ok := pMap[p.Random]; ok {
			newP.Random = nodeList[pMap[p.Random]]
		}
		newP = newP.Next
		p = p.Next
	}
	return newHead.Next
}
