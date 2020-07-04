import "strings"

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 *
 * https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (69.21%)
 * Likes:    980
 * Dislikes: 87
 * Total Accepted:    199.9K
 * Total Submissions: 288.2K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * Given a string, you need to reverse the order of characters in each word
 * within a sentence while still preserving whitespace and initial word order.
 *
 * Example 1:
 *
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 * Note:
 * In the string, each word is separated by single space and there will not be
 * any extra space in the string.
 *
 */

// @lc code=start
func reverseWords(s string) string {
	return reverseWords3(s)
}

// 采用了系统的strings.Split， append方法
func reverseWords3(s string) string {
	sArr := strings.Split(s, ` `)
	for i := range sArr {
		sArr[i] = reverseWord(sArr[i])
	}
	return strings.Join(sArr, ` `)
}

func reverseWord(s string) string {
	var res []byte
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}

// 思路与第一种方式类似，只是用了strings.Index方法与互相交换逻辑
func reverseWords2(s string) string {
	start := 0
	for {
		if i := strings.Index(s[start:], " "); i == -1 {
			break
		} else {
			s = reverse(s, start, start+i-1)
			start = start + i + 1
		}
	}
	if start < len(s) {
		s = reverse(s, start, len(s)-1)
	}
	return s
}

func reverse(s string, start, end int) string {
	sb := []byte(s)
	for start < end {
		sb[start], sb[end] = sb[end], sb[start]
		start, end = start+1, end-1
	}
	return string(sb)
}

// 自己写的方法，有点啰嗦
func reverseWords1(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	retStr := make([]byte, len(s), len(s))

	for i, j := 0, 0; j < len(s); j = j + 1 {
		if s[j] == 32 {
			for index, k := i, j-1; k >= i; index, k = index+1, k-1 {
				retStr[index] = s[k]
			}
			retStr[j] = ' '
			i = j + 1
		} else if j == len(s)-1 {
			for index, k := i, j; k >= i; index, k = index+1, k-1 {
				retStr[index] = s[k]
			}
			break
		}
	}
	return string(retStr)
}

// @lc code=end

