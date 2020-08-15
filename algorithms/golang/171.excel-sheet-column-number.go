import "math"

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 *
 * https://leetcode.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (54.11%)
 * Likes:    1207
 * Dislikes: 165
 * Total Accepted:    325.7K
 * Total Submissions: 581.9K
 * Testcase Example:  '"A"'
 *
 * Given a column title as appear in an Excel sheet, return its corresponding
 * column number.
 *
 * For example:
 *
 *
 * ⁠   A -> 1
 * ⁠   B -> 2
 * ⁠   C -> 3
 * ⁠   ...
 * ⁠   Z -> 26
 * ⁠   AA -> 27
 * ⁠   AB -> 28
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: "A"
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: "AB"
 * Output: 28
 *
 *
 * Example 3:
 *
 *
 * Input: "ZY"
 * Output: 701
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 7
 * s consists only of uppercase English letters.
 * s is between "A" and "FXSHRXW".
 *
 *
 */

// @lc code=start
func titleToNumber(s string) int {
	return titleToNumber2(s)
}

func titleToNumber2(s string) int {
	retVal, pos := 0, len(s)-1
	for _, c := range s {
		v := int(c-'A') + 1
		retVal += v * int(math.Pow(26, float64(pos)))
		pos--
	}
	return retVal
}

func titleToNumber1(s string) int {
	retVal := 0
	for i := 0; i < len(s); i++ {
		retVal = retVal*26 + int(s[i]-'A'+1)
	}
	return retVal
}

// @lc code=end