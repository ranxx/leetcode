package main

import "fmt"

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (51.47%)
 * Likes:    394
 * Dislikes: 0
 * Total Accepted:    76.9K
 * Total Submissions: 148.4K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 *
 *
 * 示例:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
*/

// @lc code=start
type MinStack struct {
	stack   []int
	_len    int
	cap     int
	minHeap []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:   make([]int, 512),
		cap:     512,
		minHeap: make([]int, 0, 512),
	}
}

func (this *MinStack) Push(x int) {
	if this._len+1 > this.cap {
		// 重新分配
		tmp := make([]int, this.cap*2)
		copy(tmp, this.stack)
		this.stack = tmp
		this.cap *= 2
	}
	this.stack[this._len] = x
	this._len++
	// 重新堆化
	this.pushHeapify(x)
}

func (this *MinStack) pushHeapify(x int) {
	this.minHeap = append(this.minHeap, x)
	// 最小值就heapify
	if this.minHeap[0] < x {
		return
	}
	this.minHeap = make([]int, this._len)
	copy(this.minHeap, this.stack[:this._len])
	this.heapify()
}

func (this *MinStack) heapify() {
	parent := this._len / 2
	for parent > 0 {
		second, min := parent*2+1, parent*2
		if second <= this._len && this.minHeap[min-1] > this.minHeap[second-1] {
			min = second
		}
		if min <= this._len && this.minHeap[min-1] < this.minHeap[parent-1] {
			this.minHeap[min-1], this.minHeap[parent-1] = this.minHeap[parent-1], this.minHeap[min-1]
		}
		parent--
	}
}

func (this *MinStack) Pop() {
	if this._len > 0 {
		this._len--
	}
	if this._len <= 0 {
		this.minHeap = make([]int, this._len)
		return
	}
	this.popHeapify(this.stack[this._len])
}

func (this *MinStack) popHeapify(x int) {
	if x != this.minHeap[0] {
		return
	}
	this.minHeap = make([]int, this._len)
	copy(this.minHeap, this.stack[:this._len])
	this.heapify()
}

func (this *MinStack) Top() int {
	if this._len > 0 {
		return this.stack[this._len-1]
	}
	return 0
}

func (this *MinStack) len() int {
	return this._len
}

func (this *MinStack) GetMin() int {
	if this._len <= 0 {
		return 0
	}
	return this.minHeap[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

func main() {
	obj := Constructor()
	obj.Push(101)
	fmt.Println(obj.Top())
	obj.Pop()
	param_3 := obj.Top()
	fmt.Println(obj.len())
	fmt.Println(obj.GetMin())
	obj.Push(19999)
	fmt.Println(obj.GetMin())
	obj.Push(1202)
	param_4 := obj.GetMin()
	obj.Pop()
	fmt.Println(param_3, obj.len(), param_4, obj.len())
	fmt.Println(obj.GetMin())
}
