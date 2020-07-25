/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (61.65%)
 * Likes:    4540
 * Dislikes: 90
 * Total Accepted:    999.9K
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 *
 * Follow up:
 *
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
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
func reverseList(head *ListNode) *ListNode {
	return reverseList1(head)
}

// solution by recursive, Time complexity: O(n), Space complexity: O(n)
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := reverseList2(head.Next)
	// confident
	head.Next.Next = head
	// avoid recycle link
	head.Next = nil
	return tmp
}

// solution by iterative, Time complexity: O(n), Space complexity: O(1)
func reverseList1(head *ListNode) *ListNode {
	var preNode *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode, curNode = curNode, nextNode
	}
	return preNode
}

// @lc code=end