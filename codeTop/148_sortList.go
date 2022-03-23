package codeTop

/*
148. 排序链表
https://leetcode-cn.com/problems/sort-list/
给你链表的头结点head，请将其按升序排列并返回排序后的链表


思路：
归并排序
https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/
*/

func sortList(head *ListNode) *ListNode {
	// 递归结束条件
	if head == nil || head.Next == nil {
		return head
	}

	// 找到链表中间节点，并断开链表。递归下探
	middleNode := middleNode148(head)
	rightNode := middleNode.Next
	middleNode.Next = nil

	left := sortList(head)
	right := sortList(rightNode)

	// 当前层业务操作，合并有序链表
	return mergeTwoList148(left, right)

}

// 找到链表的中间节点（876. 链表的中间节点）
func middleNode148(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 合并两个有序链表（21. 合并两个有序链表）
func mergeTwoList148(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}

	if head1 != nil {
		cur.Next = head1
	} else if head2 != nil {
		cur.Next = head2
	}

	return dummy.Next
}
