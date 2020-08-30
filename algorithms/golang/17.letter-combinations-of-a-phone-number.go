/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (46.93%)
 * Likes:    4246
 * Dislikes: 431
 * Total Accepted:    646K
 * Total Submissions: 1.4M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	return letterCombinations1(digits)
}

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	output, index, cur := []string{}, 0, ""
	hash := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	backtrack(hash, &output, digits, cur, index)
	return output
}

func backtrack(hash map[byte]string, output *[]string, digits string, cur string, index int) {
	if index > len(digits) {
		return
	}

	if len(cur) == len(digits) {
		*output = append(*output, cur)
		return
	}

	symbols := hash[digits[index]]
	for _, v := range symbols {
		cur += string(v)
		backtrack(hash, output, digits, cur, index+1)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end