/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 *
 * https://leetcode.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (54.94%)
 * Likes:    1434
 * Dislikes: 279
 * Total Accepted:    328.3K
 * Total Submissions: 595K
 * Testcase Example:  '"a"\n"b"'
 *
 * Given two stings ransomNote and magazine, return true if ransomNote can be
 * constructed from magazine and false otherwise.
 *
 * Each letter in magazine can only be used once in ransomNote.
 *
 *
 * Example 1:
 * Input: ransomNote = "a", magazine = "b"
 * Output: false
 * Example 2:
 * Input: ransomNote = "aa", magazine = "ab"
 * Output: false
 * Example 3:
 * Input: ransomNote = "aa", magazine = "aab"
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= ransomNote.length, magazine.length <= 10^5
 * ransomNote and magazine consist of lowercase English letters.
 *
 *
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	return solution2(ransomNote, magazine)
}

func solution1(ransomNote string, magazine string) bool {
	vMap := make(map[rune]int)
	for _, val := range magazine {
		vMap[val] += 1
	}

	for _, val := range ransomNote {
		if vMap[val] > 0 {
			vMap[val] -= 1
		} else {
			return false
		}
	}

	return true
}

func solution2(ransomNote string, magazine string) bool {
	cmpArr1 := make([]int, 26)
	cmpArr2 := make([]int, 26)

	for _, val := range ransomNote {
		cmpArr1[val-'a'] += 1
	}

	for _, val := range magazine {
		cmpArr2[val-'a'] += 1
	}

	for i := 0; i < 26; i++ {
		if cmpArr1[i] > cmpArr2[2] {
			return false
		}
	}
	return true
}

// @lc code=end