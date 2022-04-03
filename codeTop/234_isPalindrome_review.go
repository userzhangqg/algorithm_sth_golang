package codeTop

func isPalindromeReview(head *ListNode) bool {
	var nums []int

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}

func isPalindromeReview1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head

	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil

	var head2 *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}

	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}
