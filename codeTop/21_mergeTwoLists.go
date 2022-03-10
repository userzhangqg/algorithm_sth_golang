package codeTop

/*
21. 合并两个有序链表
https://leetcode-cn.com/problems/merge-two-sorted-lists/
将两个升序链表合并成一个新的升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

！！！ 未做出
正确思路：递归 or 迭代
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		// 让下一个决出大小告诉我就好
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoListsDD(list1 *ListNode, list2 *ListNode) *ListNode {
	// 迭代解法，使用哑结点连接。return dummy.Next
	dummy := &ListNode{}

	cur := dummy

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next
		} else if list1 == nil {
			cur.Next = list2
			break
		} else if list2 == nil {
			cur.Next = list1
			break
		}
	}
	return dummy.Next
}
