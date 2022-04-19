package codeTop

func swapPairsReview(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	pre := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		next := cur.Next.Next
		pre.Next = cur.Next
		cur.Next = next
		pre.Next.Next = cur

		pre = pre.Next.Next
		cur = next
	}
	return dummy.Next
}
