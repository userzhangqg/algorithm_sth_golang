package codeTop

func deleteDuplicatesReview(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil {
		for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
