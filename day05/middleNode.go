package day05

/*
876. 链表的中间节点

给定一个头结点为head 的非空单链表， 返回链表的中间结点
如果有两个中间结点，则返回第二个中间结点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var nodeList []*ListNode
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	mid := len(nodeList)/2 + 1
	return nodeList[mid]
}

func middleNodeDemo(head *ListNode) *ListNode {
	/*
		快慢指针实现
	*/
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
