package codeTop

func mergeKListsReview(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := range lists {
		ans = mergeTwoLists23Review(ans, lists[i])
	}
	return ans
}

func mergeTwoLists23Review(node1, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for node1 != nil || node2 != nil {
		if node1 == nil {
			cur.Next = node2
			break
		}
		if node2 == nil {
			cur.Next = node1
			break
		}

		if node1.Val < node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else {
			cur.Next = node2
			node2 = node2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
