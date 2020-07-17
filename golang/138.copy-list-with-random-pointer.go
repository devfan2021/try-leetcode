/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 *
 * https://leetcode.com/problems/copy-list-with-random-pointer/description/
 *
 * algorithms
 * Medium (35.26%)
 * Likes:    3272
 * Dislikes: 665
 * Total Accepted:    413.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[7,null],[13,0],[11,4],[10,2],[1,0]]'
 *
 * A linked list is given such that each node contains an additional random
 * pointer which could point to any node in the list or null.
 *
 * Return a deep copy of the list.
 *
 * The Linked List is represented in the input/output as a list of n nodes.
 * Each node is represented as a pair of [val, random_index] where:
 *
 *
 * val: an integer representing Node.val
 * random_index: the index of the node (range from 0 to n-1) where random
 * pointer points to, or null if it does not point to any node.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
 * Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [[1,1],[2,1]]
 * Output: [[1,1],[2,1]]
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: head = [[3,null],[3,0],[3,null]]
 * Output: [[3,null],[3,0],[3,null]]
 *
 *
 * Example 4:
 *
 *
 * Input: head = []
 * Output: []
 * Explanation: Given linked list is empty (null pointer), so return null.
 *
 *
 *
 * Constraints:
 *
 *
 * -10000 <= Node.val <= 10000
 * Node.random is null or pointing to a node in the linked list.
 * Number of Nodes will not exceed 1000.
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// 题目解法不是太理解
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cloneMap := map[*Node]*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		cloneMap[cur] = &Node{}
	}

	for k, v := range cloneMap {
		v.Val = k.Val
		if k.Next != nil {
			v.Next = cloneMap[k.Next]
		}
		if k.Random != nil {
			v.Random = cloneMap[k.Random]
		}
	}
	return cloneMap[head]
}

// @lc code=end