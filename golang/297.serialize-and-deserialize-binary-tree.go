import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 *
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (46.77%)
 * Likes:    3093
 * Dislikes: 150
 * Total Accepted:    333.4K
 * Total Submissions: 704.3K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 * Serialization is the process of converting a data structure or object into a
 * sequence of bits so that it can be stored in a file or memory buffer, or
 * transmitted across a network connection link to be reconstructed later in
 * the same or another computer environment.
 *
 * Design an algorithm to serialize and deserialize a binary tree. There is no
 * restriction on how your serialization/deserialization algorithm should work.
 * You just need to ensure that a binary tree can be serialized to a string and
 * this string can be deserialized to the original tree structure.
 *
 * Example:
 *
 *
 * You may serialize the following tree:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 *
 * as "[1,2,3,null,null,4,5]"
 *
 *
 * Clarification: The above format is the same as how LeetCode serializes a
 * binary tree. You do not necessarily need to follow this format, so please be
 * creative and come up with different approaches yourself.
 *
 * Note: Do not use class member/global/static variables to store states. Your
 * serialize and deserialize algorithms should be stateless.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	separator string
	emptyStr  string
}

func Constructor() Codec {
	return Codec{separator: ",", emptyStr: "null"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.emptyStr + this.separator
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	return strconv.Itoa(root.Val) + this.separator + left + right
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	items := strings.Split(data, this.separator)
	return this.helperDeserialize(&items)
}

func (this *Codec) helperDeserialize(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	nodeVal := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if nodeVal == this.emptyStr {
		return nil
	}

	val, _ := strconv.Atoi(nodeVal)
	curNode := TreeNode{Val: val}
	curNode.Left = this.helperDeserialize(nodes)
	curNode.Right = this.helperDeserialize(nodes)
	return &curNode
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end