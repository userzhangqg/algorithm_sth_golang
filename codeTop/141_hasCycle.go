package codeTop

/*
141. 环形链表
https://leetcode-cn.com/problems/linked-list-cycle/
给你一个链表的头节点head，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。
注意：pos不作为参数进行传递。 仅仅是为了标识链表中的实际情况。
如果链表中存在环，则返回true。否则，返回false

!!! 思路不清
正确思路：
1. 哈希表
2. 快慢指针
一文搞定常见的链表问题
https://leetcode-cn.com/problems/linked-list-cycle/solution/yi-wen-gao-ding-chang-jian-de-lian-biao-wen-ti-h-2/
*/

func hasCycle1(head *ListNode) bool {
	hashMap := make(map[*ListNode]int)

	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		} else {
			hashMap[head] = 0
		}
		head = head.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	// 判断为空情况
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		// 只需判断快指针是否为空即可
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next
	}
	return true
}
