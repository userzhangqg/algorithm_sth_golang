package day10

/*
21. 合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	newList := ListNode{}
	curList := &newList

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curList.Next = list2
			list2 = list2.Next
		} else {
			curList.Next = list1
			list1 = list1.Next
		}
		curList = curList.Next
	}

	if list1 == nil {
		curList.Next = list2
	} else {
		curList.Next = list1
	}

	return newList.Next
}
