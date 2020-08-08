import "fmt"

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 *
 * https://leetcode.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (47.18%)
 * Likes:    582
 * Dislikes: 995
 * Total Accepted:    84.3K
 * Total Submissions: 177.3K
 * Testcase Example:  '0'
 *
 * A binary watch has 4 LEDs on the top which represent the hours (0-11), and
 * the 6 LEDs on the bottom represent the minutes (0-59).
 * Each LED represents a zero or one, with the least significant bit on the
 * right.
 *
 * For example, the above binary watch reads "3:25".
 *
 * Given a non-negative integer n which represents the number of LEDs that are
 * currently on, return all possible times the watch could represent.
 *
 * Example:
 * Input: n = 1Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04",
 * "0:08", "0:16", "0:32"]
 *
 *
 * Note:
 *
 * The order of output does not matter.
 * The hour must not contain a leading zero, for example "01:00" is not valid,
 * it should be "1:00".
 * The minute must be consist of two digits and may contain a leading zero, for
 * example "10:2" is not valid, it should be "10:02".
 *
 *
 */

// @lc code=start
func readBinaryWatch(num int) []string {
	return readBinaryWatch1(num)
}

func readBinaryWatch1(num int) []string {
	retVals := []string{}
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if helper(h)+helper(m) == num {
				retVals = append(retVals, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return retVals
}

func helper(num int) int {
	res := 0
	for num > 0 {
		res += num % 2
		num /= 2
	}
	return res
}

// @lc code=end