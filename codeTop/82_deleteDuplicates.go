package codeTop

/*
82. 删除排序链表中的重复元素II
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/

给定一个已排序的链表头head，删除原始链表中所有重复数字的节点，只留下不同的数字。返回已排序的链表。


！！！ 做过，思路不清
思路：

*/

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil {
		for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		// 如果后面的节点没动，则切换到下一个结点
		if pre.Next == cur {
			pre = pre.Next
		} else {
			// 如果变了，将前节点的下一个节点切换
			pre.Next = cur.Next
		}

		cur = cur.Next
	}
	return dummy.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil {
		// 是否包含重复flag
		flag := false
		for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
			flag = true
			cur = cur.Next
		}
		if flag {
			// 如果包含重复continue
			cur = cur.Next
			continue
		}

		pre.Next = cur
		pre = pre.Next
		cur = cur.Next
	}
	// ! 最后一个一定是nil
	pre.Next = nil
	return dummy.Next
}
