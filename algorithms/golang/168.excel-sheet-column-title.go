/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (30.75%)
 * Likes:    1276
 * Dislikes: 255
 * Total Accepted:    223.3K
 * Total Submissions: 718.4K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 *
 * For example:
 *
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "A"
 *
 *
 * Example 2:
 *
 *
 * Input: 28
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: 701
 * Output: "ZY"
 *
 */

// @lc code=start
func convertToTitle(n int) string {
	return convertToTitle1(n)
}

func convertToTitle1(n int) string {
	retVal := ""
	for n > 0 {
		n = n - 1 // key point
		x := 65 + n%26
		retVal = string(x) + retVal
		n = n / 26
	}
	return retVal
}

// @lc code=end