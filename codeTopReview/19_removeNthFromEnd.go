package codeTopReview

/*
19. 删除链表的第n个结点
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

复杂度分析

时间复杂度：
O(L)，其中
L 是链表的长度。
空间复杂度：
O(1)。
*/

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nums := 0
	tmp := head

	for tmp != nil {
		nums++
		tmp = tmp.Next
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	k := 0

	for k+n != nums {
		k++
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	pre := dummy
	cur := head

	for i := 0; i < n; i++ {
		cur = cur.Next
	}

	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}

	pre.Next = pre.Next.Next
	return dummy.Next
}
