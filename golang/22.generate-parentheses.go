/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (61.78%)
 * Likes:    5472
 * Dislikes: 278
 * Total Accepted:    569.4K
 * Total Submissions: 909.7K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	return generateParenthesis1(n)
}

// recursive by character
func generateParenthesis2(n int) []string {
	cache := make(map[int][]string)
	return helper(cache, n)
}

func helper(cache map[int][]string, n int) []string {
	if v, ok := cache[n]; ok {
		return v
	}

	retValus := []string{}
	if n == 0 {
		retValus = append(retValus, "")
	} else {
		for i := 0; i < n; i++ {
			leftResults := helper(cache, i)
			for _, left := range leftResults {
				rightResults := helper(cache, n-i-1)
				for _, right := range rightResults {
					retValus = append(retValus, "("+left+")"+right)
				}
			}
		}
	}
	cache[n] = retValus
	return retValus
}

// backtrack
func generateParenthesis1(n int) []string {
	retVals := []string{}
	backtrack(&retVals, "", 0, 0, n)
	return retVals
}

func backtrack(retVals *[]string, current string, open, close, max int) {
	if len(current) == 2*max {
		*retVals = append(*retVals, current)
		return
	}

	if open < max {
		backtrack(retVals, current+"(", open+1, close, max)
	}
	if open > close {
		backtrack(retVals, current+")", open, close+1, max)
	}
}

// @lc code=end