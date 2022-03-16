package codeTop

func detectCycleReview(head *ListNode) *ListNode {
	hashMap := make(map[*ListNode]bool)

	for head != nil {
		if hashMap[head] {
			return head
		}
		hashMap[head] = true
		head = head.Next
	}
	return nil
}
