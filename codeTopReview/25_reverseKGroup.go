package codeTopReview

/*
25. k个一组翻转链表
https://leetcode.cn/problems/reverse-nodes-in-k-group/

思路：
备份好前置，后置节点，翻转区间节点。必要时使用哑结点
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}

	pre := dummy

	for head != nil {
		tail := pre

		for n := 0; n < k; n++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		next := tail.Next
		head, tail = my25Reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
	}
	return dummy.Next
}

func my25Reverse(head, tail *ListNode) (*ListNode, *ListNode) {
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
