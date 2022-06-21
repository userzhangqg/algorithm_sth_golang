package codeTopReview

/*
206. 反转链表
https://leetcode.cn/problems/reverse-linked-list/
*/

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	//cur := &ListNode{} // 不可以初始化，会产生零值  1，2，3，4，5 翻转后5，4，3，2，1，0

	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}
	return cur
}
