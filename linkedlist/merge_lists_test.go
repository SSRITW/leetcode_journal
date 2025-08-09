package linkedlist

import (
	"fmt"
	"testing"
)

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

// 两两结合
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for n := len(lists); n > 1; {
		dummy := make([]*ListNode, 0, len(lists)/2+1)
		for i := 0; i < len(lists); i += 2 {
			if i+1 == len(lists) {
				dummy = append(dummy, lists[i])
				break
			}
			dummy = append(dummy, mergeTwoLists(lists[i], lists[i+1]))
			n--
		}
		lists = dummy
	}
	return lists[0]
}

// 换成递归实现，思路和2一样，还是两两结合
func mergeKLists3(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if r < l {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func TestMergeKList(t *testing.T) {
	list := make([]*ListNode, 3)
	list[0] = &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list[1] = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list[2] = &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	result := mergeKLists3(list)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
