package codeTopReview

/*
92. 反转链表II
https://leetcode.cn/problems/reverse-linked-list-ii/

复杂度分析：

时间复杂度：O(N)O(N)，其中 NN 是链表总节点数。最多只遍历了链表一次，就完成了反转。

空间复杂度：O(1)O(1)。只使用到常数个变量。
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	pre := dummy

	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
