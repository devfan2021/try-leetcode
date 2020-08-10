/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 *
 * https://leetcode.com/problems/design-circular-queue/description/
 *
 * algorithms
 * Medium (43.30%)
 * Likes:    567
 * Dislikes: 81
 * Total Accepted:    64.8K
 * Total Submissions: 148.8K
 * Testcase Example:  '["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * Design your implementation of the circular queue. The circular queue is a
 * linear data structure in which the operations are performed based on FIFO
 * (First In First Out) principle and the last position is connected back to
 * the first position to make a circle. It is also called "Ring Buffer".
 *
 * One of the benefits of the circular queue is that we can make use of the
 * spaces in front of the queue. In a normal queue, once the queue becomes
 * full, we cannot insert the next element even if there is a space in front of
 * the queue. But using the circular queue, we can use the space to store new
 * values.
 *
 * Your implementation should support following operations:
 *
 *
 * MyCircularQueue(k): Constructor, set the size of the queue to be k.
 * Front: Get the front item from the queue. If the queue is empty, return
 * -1.
 * Rear: Get the last item from the queue. If the queue is empty, return
 * -1.
 * enQueue(value): Insert an element into the circular queue. Return true if
 * the operation is successful.
 * deQueue(): Delete an element from the circular queue. Return true if the
 * operation is successful.
 * isEmpty(): Checks whether the circular queue is empty or not.
 * isFull(): Checks whether the circular queue is full or not.
 *
 *
 *
 *
 * Example:
 *
 *
 * MyCircularQueue circularQueue = new MyCircularQueue(3); // set the size to
 * be 3
 * circularQueue.enQueue(1);  // return true
 * circularQueue.enQueue(2);  // return true
 * circularQueue.enQueue(3);  // return true
 * circularQueue.enQueue(4);  // return false, the queue is full
 * circularQueue.Rear();  // return 3
 * circularQueue.isFull();  // return true
 * circularQueue.deQueue();  // return true
 * circularQueue.enQueue(4);  // return true
 * circularQueue.Rear();  // return 4
 *
 *
 *
 * Note:
 *
 *
 * All values will be in the range of [0, 1000].
 * The number of operations will be in the range of [1, 1000].
 * Please do not use the built-in Queue library.
 *
 *
*/

// @lc code=start
// Basic Ideas: Use an Array
// Differentiate two cases: full or empty
// Two pointers: front, rear
//     The cell of rear won't store any items.
//     When rear == front -> empty
//     When (rear+1)%n == front -> full
// Complexity: Time O(1), Space O(n)
// https://leetcode.com/explore/learn/card/queue-stack/228/first-in-first-out-data-structure/1337/discuss/152101/Golang-solution-with-array
type MyCircularQueue struct {
	data              []int
	front, rear, size int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
// 巧妙采用k+1为queue的大小
func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{}
	queue.data = make([]int, k+1)
	queue.front, queue.rear, queue.size = 0, 0, k+1
	return queue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % this.size
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.size
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.rear-1+this.size)%this.size]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.rear == this.front

}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1)%this.size == this.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

