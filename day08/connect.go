package day08

/*
116. 填充每个节点的下一个右侧节点指针

给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下
填充它的每个next指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将next指针设置为Null。
初始状态下，所有next指针都被设置为NULL
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect1(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root
	head.Next = nil

	for head != nil {
		cur := head

		for cur != nil {

			if cur.Left != nil {
				cur.Left.Next = cur.Right
			}

			if cur.Next != nil {
				if cur.Right != nil {
					cur.Right.Next = cur.Next.Left
				}
			}

			cur = cur.Next
		}

		head = head.Left
	}

	return root
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root

	for head.Left != nil {

		cur := head

		for cur != nil {
			cur.Left.Next = cur.Right

			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}

			cur = cur.Next
		}

		head = head.Left
	}

	return root
}
