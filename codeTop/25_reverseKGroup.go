package codeTop

/*
困难：25. K个一组翻转链表
https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

基础：206.反转链表
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有的顺序。
进阶:
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}

	pre := hair

	for head != nil {
		tail := pre

		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}

		next := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := &ListNode{}
	cur := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
