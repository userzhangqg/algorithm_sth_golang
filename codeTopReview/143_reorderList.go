package codeTopReview

/*
143. 重排链表
https://leetcode.cn/problems/reorder-list/

复杂度分析:

线性表：

时间复杂度：O(N)，其中 NN 是链表中的节点数。
空间复杂度：O(N)，其中 NN 是链表中的节点数。主要为线性表的开销。
*/

func reorderList(head *ListNode) {
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
