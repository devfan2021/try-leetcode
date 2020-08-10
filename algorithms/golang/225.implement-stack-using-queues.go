/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (44.36%)
 * Likes:    675
 * Dislikes: 581
 * Total Accepted:    181K
 * Total Submissions: 403.8K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 *
 *
 * Example:
 *
 *
 * MyStack stack = new MyStack();
 *
 * stack.push(1);
 * stack.push(2);
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 *
 *
 */
// @lc code=start
type MyStack struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{data: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.data) > 1 {
		for i := len(this.data) - 1; i > 0; i-- {
			this.data[i], this.data[i-1] = this.data[i-1], this.data[i]
		}
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		panic("MyStack empty")
	}
	retVal := this.data[0]
	this.data = this.data[1:]
	return retVal
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		panic("MyStack empty")
	}
	return this.data[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end