package codeTop

/*
23. 合并K个升序链表

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表
*/

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = myMerge23(ans, lists[i])
	}
	return ans
}

func myMerge23(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}

		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
