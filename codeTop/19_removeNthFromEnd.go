package codeTop

/*
19. 删除链表的倒数第n个结点

给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	tmp := head
	for tmp != nil {
		size++
		tmp = tmp.Next
	}

	num := 1
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for num+n-1 != size {
		pre = pre.Next
		cur = cur.Next
		num++
	}
	pre.Next = cur.Next
	return dummy.Next
}
