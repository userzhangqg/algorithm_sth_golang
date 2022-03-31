package codeTop

/*
设计一个支持push，pop，top操作，并能在常数时间内检索到最小元素的栈

实现MinStack类：
- MinStack()初始化堆栈对象
- void push(int val) 将元素val推入堆栈
- void pop() 删除堆栈顶部的元素。
- int top() 获取堆栈顶部的元素。
- int getMin() 获取堆栈中的最小元素。

！！！ 未做出
思路：
辅助栈存放最小值实现。
*/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor165() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) > 0 {
		this.minStack = append(this.minStack, min165(this.minStack[len(this.minStack)-1], val))
	} else {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min165(x, y int) int {
	if x < y {
		return x
	}
	return y
}
