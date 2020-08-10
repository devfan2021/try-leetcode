import "strings"

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 *
 * https://leetcode.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (20.47%)
 * Likes:    1003
 * Dislikes: 2834
 * Total Accepted:    393.3K
 * Total Submissions: 1.9M
 * Testcase Example:  '"the sky is blue"'
 *
 * Given an input string, reverse the string word by word.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "the sky is blue"
 * Output: "blue is sky the"
 *
 *
 * Example 2:
 *
 *
 * Input: "  hello world!  "
 * Output: "world! hello"
 * Explanation: Your reversed string should not contain leading or trailing
 * spaces.
 *
 *
 * Example 3:
 *
 *
 * Input: "a good   example"
 * Output: "example good a"
 * Explanation: You need to reduce multiple spaces between two words to a
 * single space in the reversed string.
 *
 *
 *
 *
 * Note:
 *
 *
 * A word is defined as a sequence of non-space characters.
 * Input string may contain leading or trailing spaces. However, your reversed
 * string should not contain leading or trailing spaces.
 * You need to reduce multiple spaces between two words to a single space in
 * the reversed string.
 *
 *
 *
 *
 * Follow up:
 *
 * For C programmers, try to solve it in-place in O(1) extra space.
 */

// @lc code=start
func reverseWords(s string) string {
	return reverseWords1(s)
}

func reverseWords2(s string) string {
	sArr := strings.Split(s, " ")
	count := 0
	i := 0
	for {
		if i >= len(sArr) {
			break
		}
		if sArr[i] == "" {
			i++
		} else {
			sArr[count], sArr[i] = sArr[i], sArr[count]
			i++
			count++
		}
	}
	for i := 0; i < count/2; i++ {
		sArr[i], sArr[count-i-1] = sArr[count-i-1], sArr[i]
	}
	return strings.Join(sArr[0:count], " ")
}

func reverseWords1(s string) string {
	sArr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		j := i
		for j = i; j < len(s); j++ {
			if s[j] == ' ' {
				break
			}
		}
		sArr = append(sArr, s[i:j])
		i = j
	}
	ret := ""
	for i := len(sArr) - 1; i > 0; i-- {
		ret += sArr[i] + " "
	}
	if len(sArr) > 0 {
		ret += sArr[0]
	}
	return ret
}

// @lc code=end

