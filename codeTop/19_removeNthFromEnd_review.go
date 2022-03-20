package codeTop

func removeNthFromEndReview(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	k := 0
	for cur != nil {
		k++
		cur = cur.Next
	}

	for k != n {
		k--
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
