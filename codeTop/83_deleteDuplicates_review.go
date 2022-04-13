package codeTop

func deleteDuplicates83Review(head *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	cur := head

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		pre.Next = cur
		pre = pre.Next
		cur = cur.Next
	}
	return dummy.Next
}
