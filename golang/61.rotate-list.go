/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (29.63%)
 * Likes:    1248
 * Dislikes: 1001
 * Total Accepted:    274.1K
 * Total Submissions: 918.7K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, rotate the list to the right by k places, where k is
 * non-negative.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->4->5->NULL, k = 2
 * Output: 4->5->1->2->3->NULL
 * Explanation:
 * rotate 1 steps to the right: 5->1->2->3->4->NULL
 * rotate 2 steps to the right: 4->5->1->2->3->NULL
 *
 *
 * Example 2:
 *
 *
 * Input: 0->1->2->NULL, k = 4
 * Output: 2->0->1->NULL
 * Explanation:
 * rotate 1 steps to the right: 2->0->1->NULL
 * rotate 2 steps to the right: 1->2->0->NULL
 * rotate 3 steps to the right: 0->1->2->NULL
 * rotate 4 steps to the right: 2->0->1->NULL
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	return rotateRight2(head, k)
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	low, fast := head, head
	for i := 0; i < k; i++ {
		if fast.Next == nil { // 巧妙求余，重新循环一次，链表的长度就是i+1 (不用计算链表的长度)
			return rotateRight2(head, k%(i+1))
		}
		fast = fast.Next
	}

	//再次遍历直到链表的末端
	for fast.Next != nil {
		low, fast = low.Next, fast.Next
	}

	// 巧妙的进行了收尾串联
	newHead := low.Next
	low.Next, fast.Next = nil, head
	return newHead
}

// 写的一直有bug, case验证不通过, 节点移位有问题
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	count, cur := 0, head
	for ; cur != nil; cur = cur.Next {
		count++
	}

	count = k % count
	if count == 0 { // 整除，循环，直接返回不用做操作
		return head
	}

	// 先走count步
	low, fast := head, head
	for ; count > 0; count-- {
		fast = fast.Next
	}

	// 这种采用dummy节点的方式，就会出现类似这种case不能通过 [1,2]\n1  ===还是有问题，待进行解决处理
	dummy := &ListNode{}
	node := dummy
	for fast.Next != nil {
		node.Next = fast
		node = node.Next
		low, fast = low.Next, fast.Next
	}

	if fast.Next == nil {
		node.Next = &ListNode{Val: fast.Val}
		node = node.Next
	}

	low.Next = nil
	node.Next = head
	return dummy.Next
}

// @lc code=end