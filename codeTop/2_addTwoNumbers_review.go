package codeTop

func addTwoNumbersReview(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	var tmp int
	for l1 != nil || l2 != nil {
		var sum int
		sum = sum + tmp
		if l1 != nil {
			sum = l1.Val + sum
			l1 = l1.Next
		}
		if l2 != nil {
			sum = l2.Val + sum
			l2 = l2.Next
		}

		tmp = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	if tmp != 0 {
		cur.Next = &ListNode{Val: tmp}
	}
	return dummy.Next
}
