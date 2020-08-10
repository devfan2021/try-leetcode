import "math"

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (43.90%)
 * Likes:    3460
 * Dislikes: 337
 * Total Accepted:    562K
 * Total Submissions: 1.3M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 *
 * Constraints:
 *
 *
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 *
 *
*/

// @lc code=start
// add help stack, time complexity: O(1), space complexity: O(n)
type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: []int{}, minData: []int{math.MaxInt64}}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	minVal := this.minData[len(this.minData)-1]
	if x < minVal {
		minVal = x
	}
	this.minData = append(this.minData, minVal)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minData = this.minData[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]

}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
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