package codeTop

/*
234. 回文链表

给你一个单链表的头节点head，请你判断该链表是否是回文链表。如果是，返回true；否则，返回false

进阶：使用O(n)时间复杂度，O(1)空间复杂度

$$$ 部分完成
思路：
1. 列表，双指针判断，O(n)空间复杂度
2. 快慢指针，翻转
*/

func isPalindrome(head *ListNode) bool {
	var values []int

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		if values[i] != values[j] {
			return false
		}
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
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
	// slow节点以前断开，2/2 2/3
	pre.Next = nil

	// 翻转
	var head2 *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}

	// 比对
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}
