package codeTopReview

/*
23. 合并k个升序链表
https://leetcode.cn/problems/merge-k-sorted-lists/

时间复杂度：故渐进时间复杂度为 O(k^2 n)
空间复杂度：没有用到与 kk 和 nn 规模相关的辅助空间，故渐进空间复杂度为 O(1)O(1)。

*/

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := range lists {
		ans = merge23(ans, lists[i])
	}
	return ans
}

func merge23(node1, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for node1 != nil || node2 != nil {
		if node1 == nil {
			cur.Next = node2
			break
		}
		if node2 == nil {
			cur.Next = node1
			break
		}

		if node1.Val < node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else {
			cur.Next = node2
			node2 = node2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
