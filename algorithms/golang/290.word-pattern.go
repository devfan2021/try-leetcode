import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 *
 * https://leetcode.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (36.75%)
 * Likes:    1106
 * Dislikes: 151
 * Total Accepted:    188.8K
 * Total Submissions: 512.2K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 *
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in str.
 *
 * Example 1:
 *
 *
 * Input: pattern = "abba", str = "dog cat cat dog"
 * Output: true
 *
 * Example 2:
 *
 *
 * Input:pattern = "abba", str = "dog cat cat fish"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: pattern = "aaaa", str = "dog cat cat dog"
 * Output: false
 *
 * Example 4:
 *
 *
 * Input: pattern = "abba", str = "dog dog dog dog"
 * Output: false
 *
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains
 * lowercase letters that may be separated by a single space.
 *
 */

// @lc code=start
// 必须要注意hash的key设计，加上前缀做区分
func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	hash := map[string]int{}
	// "abc"\n"b c a"
	for i := 0; i < len(pattern); i++ {
		letter := "char_" + string(pattern[i])
		// letter := string([]byte{pattern[i]})
		word := "word_" + words[i]
		if _, ok := hash[letter]; !ok {
			hash[letter] = i
		}
		if _, ok := hash[word]; !ok {
			hash[word] = i
		}
		if hash[letter] != hash[word] {
			fmt.Println(hash)
			return false
		}
	}
	return true
}

// @lc code=end