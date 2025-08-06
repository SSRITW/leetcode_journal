package linkedlist

// 23. 合并 K 个升序链表 （困难）

// 暴力两两合并
func mergeKLists1(lists []*ListNode) (result *ListNode) {
	if len(lists) == 0 {
		return
	}
	result = lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeKLists2(lists []*ListNode) (result *ListNode) {
	if len(lists) == 0 {
		return
	}
	result = lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return
}
