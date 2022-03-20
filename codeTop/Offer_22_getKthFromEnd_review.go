package codeTop

func getKthFromEndReview(head *ListNode, k int) *ListNode {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	for n != k {
		n--
		head = head.Next
	}
	return head
}
