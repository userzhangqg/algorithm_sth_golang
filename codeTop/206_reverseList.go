package codeTop

/*
206. 反转链表

给你单链表的头结点head，请你反转链表，并返回反转后的链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var nowList *ListNode
	//cur := head
	for head != nil {
		next := head.Next
		head.Next = nowList
		nowList = head
		head = next
	}
	return nowList
}
