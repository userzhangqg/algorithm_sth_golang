package codeTop

/*
92. 反转链表II

给你单链表的头指针head和两个整数left和right，其中left<=right.
请你反转从位置left到位置right的链表节点，返回反转后的链表。

思路：
https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
1. 切开反转，连接
2. 遍历反转（优化）
*/

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}
	var leftHead, rightHead *ListNode
	tmp := head
	hair := &ListNode{Next: head}
	pre := hair

	n := 1
	for tmp != nil {
		if n == left {
			leftHead = tmp
			break
		}
		pre = pre.Next
		tmp = tmp.Next
		n++
	}

	for tmp != nil {
		if n == right {
			rightHead = tmp
			break
		}
		tmp = tmp.Next
		n++
	}

	var tail *ListNode
	if rightHead.Next != nil {
		tail = rightHead.Next
	}
	leftHead, rightHead = reverseListNode(leftHead, rightHead)
	pre.Next = leftHead
	rightHead.Next = tail
	return hair.Next
}

func reverseListNode(leftNode, rightNode *ListNode) (*ListNode, *ListNode) {
	//node := &ListNode{}
	var node *ListNode
	cur := leftNode
	for node != rightNode {
		next := cur.Next
		cur.Next = node
		node = cur
		cur = next
	}
	return rightNode, leftNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	var m, n *ListNode

	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	m = pre.Next
	n = m.Next
	for ; left < right; left++ {
		next := n.Next
		n.Next = m
		m = n
		n = next
	}

	// 连接right到后面节点
	pre.Next.Next = n
	// 切换原node left节点
	pre.Next = m

	return dummy.Next
}
