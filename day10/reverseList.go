package day10

/*
206. 反转链表

给你单链表的头节点head，请你反转链表，并返回反转后的链表
*/

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{}
	newHead.Val = head.Val

	for head.Next != nil {
		newHead = &ListNode{head.Next.Val, newHead}
		head = head.Next
	}

	return newHead
}

//双指针
func reverseListDemo1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//递归
func reverseListDemo2(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
