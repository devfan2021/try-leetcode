/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 *
 * https://leetcode.com/problems/rotate-string/description/
 *
 * algorithms
 * Easy (50.73%)
 * Likes:    1620
 * Dislikes: 79
 * Total Accepted:    131.9K
 * Total Submissions: 258.1K
 * Testcase Example:  '"abcde"\n"cdeab"'
 *
 * Given two strings s and goal, return true if and only if s can become goal
 * after some number of shifts on s.
 *
 * A shift on s consists of moving the leftmost character of s to the rightmost
 * position.
 *
 *
 * For example, if s = "abcde", then it will be "bcdea" after one shift.
 *
 *
 *
 * Example 1:
 * Input: s = "abcde", goal = "cdeab"
 * Output: true
 * Example 2:
 * Input: s = "abcde", goal = "abced"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, goal.length <= 100
 * s and goal consist of lowercase English letters.
 *
 *
 */
// "awuvupcrsatnwvxsdbonhyszidtchtisxbiqaekyervvjwfrakopukxbsoorjyioppbqhjmmjzvtjijbubhqhsdtsflvopozqufp"
// "qufpawuvupcrsatnwvxsdbonhyszidtchtisxbiqaekyervvjwfrakopukxbsoorjyioppbqhjmmjzvtjijbubhqhsdtsflvopoz"
// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return true
	}

	sIndex, gIndex := 0, 0
	for index1, val1 := range goal {
		for index2 := sIndex; index2 < len(s); index2 += 1 {
			sIndex = index2 + 1
			if (byte)val1 == s[index2] {
				break
			}
		}
		gIndex = index1 + 1
	}

	// fmt.Printf("%d---%d", gIndex, sIndex)
	if gIndex == len(goal) || sIndex != len(s) {
		return false
	}

	for index := 0; index < len(goal)-gIndex; index += 1 {
		if s[index] != goal[gIndex+index] {
			//fmt.Printf("bb")
			return false
		}
	}

	return true
}

// @lc code=end