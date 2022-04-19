package codeTop

/*
24. 两两交换链表中的节点

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头结点。
你必须在不修改节点内部的值的情况下完成本题

$ Done

*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	pre := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		next := cur.Next.Next
		pre.Next = cur.Next
		pre.Next.Next = cur
		cur.Next = next

		pre = pre.Next.Next
		cur = cur.Next
	}
	return dummy.Next
}
