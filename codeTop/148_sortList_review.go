package codeTop

func sortListReview(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	midNode := middleNode148Review(head)
	left := head
	right := midNode.Next
	midNode.Next = nil

	leftSort := sortList(left)
	rightSort := sortList(right)

	return mergeTwoList148(leftSort, rightSort)
}

func middleNode148Review(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeTwoNodes148Review(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{Next: head1}
	cur := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}

	if head1 != nil {
		cur.Next = head1
	} else if head2 != nil {
		cur.Next = head2
	}
	return dummy.Next
}
