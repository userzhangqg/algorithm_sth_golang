package codeTopReview

/*
141. 判断一个链表是否有环
https://leetcode.cn/problems/linked-list-cycle/
*/

func hasCycle(head *ListNode) bool {
	// 哈希表
	/*
		时间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。最坏情况下我们需要遍历每个节点一次。

		空间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。主要为哈希表的开销，最坏情况下我们需要将每个节点插入到哈希表中一次。
	*/
	//hashMap := map[*ListNode]bool{}
	hashMap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := hashMap[head]; !ok {
			hashMap[head] = true
		} else {
			return true
		}
		head = head.Next
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	// 快慢指针法
	/*
		时间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。

		当链表中不存在环时，快指针将先于慢指针到达链表尾部，链表中每个节点至多被访问两次。

		当链表中存在环时，每一轮移动后，快慢指针的距离将减小一。而初始距离为环的长度，因此至多移动 NN 轮。

		空间复杂度：O(1)O(1)。我们只使用了两个指针的额外空间。
	*/
	for head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next
	}
	return true
}
