package codeTopReview

/*
232. 用栈实现队列
https://leetcode.cn/problems/implement-queue-using-stacks/
*/

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor232() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		for i := len(this.inStack) - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.inStack = []int{}
	}
	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		for i := len(this.inStack) - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.inStack = []int{}
	}
	x := this.outStack[len(this.outStack)-1]
	return x
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
