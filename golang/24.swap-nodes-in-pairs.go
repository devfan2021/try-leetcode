/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (49.75%)
 * Likes:    2327
 * Dislikes: 172
 * Total Accepted:    475.1K
 * Total Submissions: 947K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
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
func swapPairs(head *ListNode) *ListNode {
	return swapPairs1(head)
}

func swapPairs2(head *ListNode) *ListNode {
	ans := &ListNode{0, head}
	dummy := ans
	for dummy.Next != nil && dummy.Next.Next != nil {
		first, second := dummy.Next, dummy.Next.Next
		first.Next = second.Next
		dummy.Next = second
		dummy.Next.Next = first
		dummy = dummy.Next.Next
	}
	return ans.Next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head.Next
	head.Next = swapPairs1(head.Next.Next)
	node.Next = head
	return node
}

// @lc code=end