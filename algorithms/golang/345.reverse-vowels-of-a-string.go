/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 *
 * https://leetcode.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (46.37%)
 * Likes:    1368
 * Dislikes: 1706
 * Total Accepted:    323.9K
 * Total Submissions: 697.3K
 * Testcase Example:  '"hello"'
 *
 * Given a string s, reverse only all the vowels in the string and return it.
 *
 * The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both
 * cases.
 *
 *
 * Example 1:
 * Input: s = "hello"
 * Output: "holle"
 * Example 2:
 * Input: s = "leetcode"
 * Output: "leotcede"
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 3 * 10^5
 * s consist of printable ASCII characters.
 *
 *
 */

// @lc code=start
func reverseVowels(s string) string {
	return solution2(s)
}

func solution1(s string) string {
	if len(s) <= 1 {
		return s
	}

	b := []byte(s)
	begin, end := 0, len(b)-1
	for begin < end {
		if !isVowel(b[begin]) {
			begin++
			continue
		}

		if !isVowel(b[end]) {
			end--
			continue
		}
		b[begin], b[end] = b[end], b[begin]
		begin++
		end--
	}

	return string(b)
}

func isVowel(c byte) bool {
	if c < 'a' {
		c += 32
	}
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func solution2(s string) string {
	if len(s) <= 1 {
		return s
	}

	vowels := "aeiouAEIOU"
	vowelMap := make(map[byte]bool)
	for _, val := range vowels {
		vowelMap[byte(val)] = true
	}

	b := []byte(s)
	begin, end := 0, len(s)-1
	for begin < end {
		if !vowelMap[b[begin]] {
			begin++
			continue
		}

		if !vowelMap[b[end]] {
			end--
			continue
		}

		b[begin], b[end] = b[end], b[begin]
		begin++
		end--
	}
	return string(b)
}

// @lc code=end