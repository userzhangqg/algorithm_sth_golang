package codeTop

/*
232. 用栈实现队列

请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作
push、pop、peek、empty

实现MyQueue类：
push x 将元素x推到队列的末尾
pop 从队列的开头移除并返回元素
peek 返回队列开头的元素
empty 如果队列为空，返回true；否则，返回false


!!! 思路错误，不可以直接用简单list，使用双栈先入后出实现
*/

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor232() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		// out空栈才可导入
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	if len(q.inStack) == 0 && len(q.outStack) == 0 {
		return true
	}
	return false
}
