/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 *
 * https://leetcode.com/problems/guess-number-higher-or-lower/description/
 *
 * algorithms
 * Easy (42.55%)
 * Likes:    400
 * Dislikes: 1731
 * Total Accepted:    154.5K
 * Total Submissions: 361.2K
 * Testcase Example:  '10\n6'
 *
 * We are playing the Guess Game. The game is as follows:
 *
 * I pick a number from 1 to n. You have to guess which number I picked.
 *
 * Every time you guess wrong, I'll tell you whether the number is higher or
 * lower.
 *
 * You call a pre-defined API guess(int num) which returns 3 possible results
 * (-1, 1, or 0):
 *
 *
 * -1 : My number is lower
 * ⁠1 : My number is higher
 * ⁠0 : Congrats! You got it!
 *
 *
 * Example :
 *
 *
 *
 * Input: n = 10, pick = 6
 * Output: 6
 *
 *
 *
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	return guessNumber2(n)
}

// 类似分三段查找
func guessNumber2(n int) int {
	low, high := 1, n
	for low <= high {
		mid1 := low + (high-low)/2
		val1 := guess(mid1)
		mid2 := high - (high-low)/2
		val2 := guess(mid2)
		if val1 == 0 {
			return mid1
		} else if val2 == 0 {
			return mid2
		} else if val1 < 0 {
			high = mid1 - 1
		} else if val2 > 0 {
			low = mid2 + 1
		} else {
			low = mid1 + 1
			high = mid2 - 1
		}
	}
	return -1
}

// 采用二分查找
func guessNumber1(n int) int {
	low, high := 1, n
	for low <= high {
		mid := (high-low)/2 + low
		val := guess(mid)
		if val == 0 {
			return mid
		} else if val == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// @lc code=end