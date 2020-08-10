import "sort"

/*
 * @lc app=leetcode id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 *
 * https://leetcode.com/problems/longest-word-in-dictionary/description/
 *
 * algorithms
 * Easy (47.95%)
 * Likes:    644
 * Dislikes: 747
 * Total Accepted:    63.1K
 * Total Submissions: 130.8K
 * Testcase Example:  '["w","wo","wor","worl","world"]'
 *
 * Given a list of strings words representing an English Dictionary, find the
 * longest word in words that can be built one character at a time by other
 * words in words.  If there is more than one possible answer, return the
 * longest word with the smallest lexicographical order.  If there is no
 * answer, return the empty string.
 *
 * Example 1:
 *
 * Input:
 * words = ["w","wo","wor","worl", "world"]
 * Output: "world"
 * Explanation:
 * The word "world" can be built one character at a time by "w", "wo", "wor",
 * and "worl".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * Output: "apple"
 * Explanation:
 * Both "apply" and "apple" can be built from other words in the dictionary.
 * However, "apple" is lexicographically smaller than "apply".
 *
 *
 *
 * Note:
 * All the strings in the input will only contain lowercase letters.
 * The length of words will be in the range [1, 1000].
 * The length of words[i] will be in the range [1, 30].
 *
 */

// @lc code=start
func longestWord(words []string) string {
	return longestWord1(words)
}

// todo: using trie tree
func longestWord2(words []string) string {
	return ""
}

func longestWord1(words []string) string {
	sort.Strings(words)
	hash, res := map[string]bool{}, ""
	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) == 1 || hash[word[:len(word)-1]] {
			if len(word) > len(res) {
				res = word
			}
			hash[word] = true
		}
	}
	return res
}

// @lc code=end