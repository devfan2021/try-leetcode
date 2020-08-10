/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (45.02%)
 * Likes:    1629
 * Dislikes: 115
 * Total Accepted:    474.9K
 * Total Submissions: 1M
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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
// straight-forward approach, time complexity:O(n), space complexity:O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	dummy, cmpVal := &ListNode{}, -1<<31
	curNode := dummy
	for head != nil {
		curVal := head.Val
		if curVal != cmpVal {
			curNode.Next = &ListNode{Val: curVal}
			curNode, cmpVal = curNode.Next, curVal
		}
		head = head.Next
	}
	return dummy.Next
}

// @lc code=end