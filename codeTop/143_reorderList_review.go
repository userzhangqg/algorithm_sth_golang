package codeTop

func reorderListReview(head *ListNode) {
	var dp []*ListNode

	tmp := head

	for tmp != nil {
		dp = append(dp, tmp)
		tmp = tmp.Next
	}

	l, r := 0, len(dp)-1

	for l < r {
		dp[l].Next = dp[r]
		l++

		if l < r {
			dp[r].Next = dp[l]
		}
		r--
	}
	dp[l].Next = nil
}

func reorderListReview2(head *ListNode) {
	cur := head
	right := getMidRightNode(cur)
	right = reverse143(right)

	merge143(cur, right)
}

func getMidRightNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}
	return slow
}

func reverse143(head *ListNode) *ListNode {
	var cur *ListNode
	tmp := head
	for tmp != nil {
		next := tmp.Next
		tmp.Next = cur
		cur = tmp
		tmp = next
	}
	return cur
}

func merge143(head1, head2 *ListNode) {
	var next1, next2 *ListNode
	for head2 != nil && head1 != head2 {
		next1 = head1.Next
		next2 = head2.Next

		head1.Next = head2
		head1 = next1

		head2.Next = head1
		head2 = next2
	}
	head1.Next = nil
}
