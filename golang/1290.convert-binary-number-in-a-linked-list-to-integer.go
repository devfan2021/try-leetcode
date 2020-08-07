/*
 * @lc app=leetcode id=1290 lang=golang
 *
 * [1290] Convert Binary Number in a Linked List to Integer
 *
 * https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/description/
 *
 * algorithms
 * Easy (80.04%)
 * Likes:    471
 * Dislikes: 40
 * Total Accepted:    75.1K
 * Total Submissions: 93.6K
 * Testcase Example:  '[1,0,1]'
 *
 * Given head which is a reference node to a singly-linked list. The value of
 * each node in the linked list is either 0 or 1. The linked list holds the
 * binary representation of a number.
 *
 * Return the decimal value of the number in the linked list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,0,1]
 * Output: 5
 * Explanation: (101) in base 2 = (5) in base 10
 *
 *
 * Example 2:
 *
 *
 * Input: head = [0]
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1]
 * Output: 1
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
 * Output: 18880
 *
 *
 * Example 5:
 *
 *
 * Input: head = [0,0]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * The Linked List is not empty.
 * Number of nodes will not exceed 30.
 * Each node's value is either 0 or 1.
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
func getDecimalValue(head *ListNode) int {
	return getDecimalValue3(head)
}

func getDecimalValue3(head *ListNode) int {
	sum := 0
	for ; head != nil; head = head.Next {
		sum = sum<<1 + head.Val
	}
	return sum
}

func getDecimalValue2(head *ListNode) int {
	if head == nil {
		return 0
	}
	sum := 0
	for head != nil {
		sum = sum<<1 | head.Val
		head = head.Next
	}
	return sum
}

func getDecimalValue1(head *ListNode) int {
	if head == nil {
		return 0
	}
	sum := 0
	for head != nil {
		sum = 2*sum + head.Val
		head = head.Next
	}
	return sum
}

// @lc code=end