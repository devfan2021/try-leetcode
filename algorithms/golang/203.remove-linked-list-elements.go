/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (37.59%)
 * Likes:    1615
 * Dislikes: 91
 * Total Accepted:    330.1K
 * Total Submissions: 875.7K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
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
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	preNode, curNode := dummy, head
	for curNode != nil {
		if curNode.Val == val {
			preNode.Next = curNode.Next
		} else {
			preNode = curNode
		}
		curNode = curNode.Next
	}
	return dummy.Next
}

// @lc code=end