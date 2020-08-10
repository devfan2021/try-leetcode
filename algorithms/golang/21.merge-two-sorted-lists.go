/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (52.86%)
 * Likes:    4313
 * Dislikes: 597
 * Total Accepted:    1M
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new sorted list. The new
 * list should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists2(l1, l2)
}

// 采用递归方式
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}

//新建一个新节点，再分别遍历两个节点
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		} else if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				cur, l1 = cur.Next, l1.Next
			} else {
				cur.Next = l2
				cur, l2 = cur.Next, l2.Next
			}
		}
	}
	return dummy.Next
}

// @lc code=end