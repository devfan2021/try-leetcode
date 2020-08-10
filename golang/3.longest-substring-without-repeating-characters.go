/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (30.15%)
 * Likes:    9426
 * Dislikes: 567
 * Total Accepted:    1.6M
 * Total Submissions: 5.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 *
 */

// @lc code=start
// 测试例子：" ", "b", "bc", "dvdf", "tmmzuxt", "bbbbb"
// 失败次数比较多，边界点判断有问题, 及清空hash的逻辑
// 待优化改进方法
func lengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring4(s)
}

// brute force
func lengthOfLongestSubstring4(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	begin, end := 0, 0
	for end < len(s) {
		if !isUniqueChars(&s, begin, end) {
			begin++
		}
		end++
	}
	return end - begin
}

func isUniqueChars(s *string, begin, end int) bool {
	for i := begin; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if (*s)[i] == (*s)[j] {
				return false
			}
		}
	}
	return true
}

// The reason is that if s[j] have a duplicate in the range [i, j) with index j' , we don't need to increase i little by little. We can skip all the elements in the range [i, j'] and let i to be j' + 1 directly.
func lengthOfLongestSubstring3(s string) int {
	hash, maxCount := make(map[byte]int), 0
	for begin, end := 0, 0; end < len(s); end++ {
		if index, ok := hash[s[end]]; ok { // check duplicate character, and skip to next index
			if index > begin { // test case: tmmzuxt
				begin = index
			}
		}
		count := end - begin + 1
		if count > maxCount {
			maxCount = count
		}
		hash[s[end]] = end + 1
	}
	return maxCount
}

// sliding window, time complexity:O(n)
func lengthOfLongestSubstring2(s string) int {
	hash, length := make(map[byte]bool), len(s)
	begin, end, maxCount := 0, 0, 0
	for begin < length && end < length {
		// c := int(s[end] - 'a')
		c := s[end]
		if hash[c] {
			hash[s[begin]] = false
			begin++
		} else {
			hash[c] = true
			end++
			count := end - begin
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}

func lengthOfLongestSubstring1(s string) int {
	count, begin := 0, 0
	hash := make(map[string]int)
	for i := 0; i < len(s); {
		val := string(s[i])
		if index, ok := hash[val]; ok {
			// 算起始位置
			if (i - begin) > count {
				count = i - begin
			}
			// 有2种方式处理，清空hash，重新从begin遍历，或者删除begin之前的在map中的值
			hash = make(map[string]int)
			begin, i = index+1, index+1
		} else {
			// 最后一位
			if i == len(s)-1 && (i-begin+1) > count {
				count = i - begin + 1
			}
			hash[val] = i
			i++
		}
	}
	return count
}

// @lc code=end