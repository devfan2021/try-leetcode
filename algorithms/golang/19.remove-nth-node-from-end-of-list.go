/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (35.06%)
 * Likes:    3335
 * Dislikes: 230
 * Total Accepted:    628.4K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd1(head, n)
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for i := 0; i < n; i++ {
		if fast.Next == nil {
			// assuming all n is valid, so this case
			// the first element to be deleted
			return head.Next
		}
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 为啥要增加一个dummy节点呢？不可以直接操作吗？
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	// 先算总节点数
	length := 0
	first := head
	for first != nil {
		length++
		first = first.Next
	}

	// 再操作需要移动的位置
	first = dummy
	length -= n
	for length > 0 {
		length--
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}

// @lc code=end