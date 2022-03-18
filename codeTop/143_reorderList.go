package codeTop

/*
143. 重排链表
https://leetcode-cn.com/problems/reorder-list/

给定一个单链表L的头结点head，单链表L表示为
L0 -> L1 -> ... -> Ln-1 -> Ln
将其重新排列后变为：
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

思路：
1. 寻找链表中点（后面1/2点） + 链表逆序 + 合并链表
2. 线性表，列表寻址
*/

func reorderList2(head *ListNode) {
	var nodeList []*ListNode
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}

	l, r := 0, len(nodeList)-1

	for l < r {
		nodeList[l].Next = nodeList[r]
		l++

		if l == r {
			break
		}

		nodeList[r].Next = nodeList[l]
		r--
	}
	nodeList[l].Next = nil
}

//======================

func reorderList(head *ListNode) {
	l1 := head
	l2 := middleNode(head)

	l2 = reverseList143(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	headSlow := head
	headFast := head

	for headFast != nil {
		headSlow = headSlow.Next
		if headFast.Next != nil {
			headFast = headFast.Next.Next
		} else {
			headFast = nil
		}
	}
	return headSlow
}

func reverseList143(head *ListNode) *ListNode {
	var nowList *ListNode
	for head != nil {
		next := head.Next
		head.Next = nowList
		nowList = head
		head = next
	}
	return nowList
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l2 != nil && l1 != l2 {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp

	}
	// 切除多余链表
	l1.Next = nil
}
