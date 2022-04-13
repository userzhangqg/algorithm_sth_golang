package codeTop

/*
83. 删除排序链表中的重复元素
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
给定一个已排序的链表的头head，删除所有重复的元素，使每个元素只出现一次。返回已排序的链表

$ Done
思路：
1. 哑结点
*/
func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	preCur := dummy
	cur := head

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		preCur.Next = cur
		preCur = preCur.Next
		cur = cur.Next
	}
	return dummy.Next
}
