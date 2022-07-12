package codeTopReview

/*
142. 环形链表II
https://leetcode.cn/problems/linked-list-cycle-ii/

时间复杂度：

1. 哈希表
时间复杂度：O(N)O(N)，其中 NN 为链表中节点的数目。我们恰好需要访问链表中的每一个节点。
空间复杂度：O(N)O(N)，其中 NN 为链表中节点的数目。我们需要将链表中的每个节点都保存在哈希表当中。

2. 快慢指针
https://leetcode.cn/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/
*/

func detectCycle(head *ListNode) *ListNode {
	hashMap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := hashMap[head]; ok {
			return head
		} else {
			hashMap[head] = true
			head = head.Next
		}
	}
	return nil
}
