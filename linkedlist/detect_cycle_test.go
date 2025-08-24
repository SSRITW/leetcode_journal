package linkedlist

// 142. 环形链表 II
//返回进入环的头结点

// 还是先使用hase实现
func detectCycle(head *ListNode) *ListNode {
	has := make(map[*ListNode]bool)
	for head != nil {
		if has[head] {
			return head
		}
		has[head] = true
		head = head.Next
	}
	return nil
}

// 解法2： Floyd 判圈算法
// 头结点到环形入口的节点数为x;入口节点到（fast和slow）相遇节点的节点数为y;相遇节点到入口节点的节点数为z
// 相遇时，slow走过节点 = 2(x+y); fast走过节点 = x+y+n(y+z);从而得出下面的公式
// x+y = n(y+z)
// x = n(y+z)-y
// x = (n-1)(y+z)+z
// 如果n=1的话，x=z;所以得出相遇的时候,一个节点从头节点出发，每次移动一个节点;另一个节点从相遇点出发,每次也移动一个节点,他们第一次相遇的节点就为环入口
func detectCycle2(head *ListNode) *ListNode {
	for fast, slow := head, head; fast != nil && fast.Next != nil; {
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			p := head
			for ; slow != p; p, slow = p.Next, slow.Next {
			}
			return p
		}
	}
	return nil
}
