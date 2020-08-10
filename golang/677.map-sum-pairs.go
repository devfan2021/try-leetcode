/*
 * @lc app=leetcode id=677 lang=golang
 *
 * [677] Map Sum Pairs
 *
 * https://leetcode.com/problems/map-sum-pairs/description/
 *
 * algorithms
 * Medium (53.15%)
 * Likes:    506
 * Dislikes: 81
 * Total Accepted:    41.5K
 * Total Submissions: 77.5K
 * Testcase Example:  '["MapSum", "insert", "sum", "insert", "sum"]\n' +
  '[[], ["apple",3], ["ap"], ["app",2], ["ap"]]'
 *
 *
 * Implement a MapSum class with insert, and sum methods.
 *
 *
 *
 * For the method insert, you'll be given a pair of (string, integer). The
 * string represents the key and the integer represents the value. If the key
 * already existed, then the original key-value pair will be overridden to the
 * new one.
 *
 *
 *
 * For the method sum, you'll be given a string representing the prefix, and
 * you need to return the sum of all the pairs' value whose key starts with the
 * prefix.
 *
 *
 * Example 1:
 *
 * Input: insert("apple", 3), Output: Null
 * Input: sum("ap"), Output: 3
 * Input: insert("app", 2), Output: Null
 * Input: sum("ap"), Output: 5
 *
 *
 *
*/

// @lc code=start
// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["a",3], ["ap"], ["b",2], ["a"]]         ===> [null,null,0,null,3]

// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["aa",3], ["a"], ["aa",2], ["a"]]        ===> [null,null,3,null,2]

// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["apple",3], ["ap"], ["app",2], ["ap"]]  ===> [null,null,3,null,5]
type MapSum struct {
	Val      int
	Children map[rune]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{Val: 0, Children: map[rune]*MapSum{}}
}

func (this *MapSum) Insert(key string, val int) {
	parent := this
	for _, c := range key {
		if child, ok := parent.Children[c]; ok {
			parent = child
		} else {
			newChild := &MapSum{Children: map[rune]*MapSum{}}
			parent.Children[c] = newChild
			parent = newChild
		}
	}
	parent.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	parent := this
	for _, c := range prefix {
		if _, ok := parent.Children[c]; !ok {
			return 0
		}
		parent = parent.Children[c]
	}

	return parent.Val + parent.SumOfChildren()
}

func (root *MapSum) SumOfChildren() int {
	sum := 0
	for _, c := range root.Children {
		if c != nil {
			sum += c.Val + c.SumOfChildren()
		}
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end