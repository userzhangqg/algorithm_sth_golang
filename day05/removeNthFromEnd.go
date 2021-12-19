package day05

/*
19. 删除链表的倒数第N个结点

给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点
*/
/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	totalNode := 0

	// 计算长度
	cur := head
	for cur != nil {
		totalNode++
		cur = cur.Next
	}

	now := 1
	newCur := head

	// 如果删除开头，直接next
	if totalNode == n {
		return head.Next
	}

	// 如果未到删除结点的前一结点，next
	for totalNode-n > now {
		now++
		newCur = newCur.Next
	}
	newCur.Next = newCur.Next.Next
	return head
}
