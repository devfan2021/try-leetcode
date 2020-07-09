/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (33.61%)
 * Likes:    1314
 * Dislikes: 1892
 * Total Accepted:    552.9K
 * Total Submissions: 1.6M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

// @lc code=start
// 2147395599
func mySqrt(x int) int {
	return mySqrt2(x)
}

func mySqrt2(x int) int {
	ret := x
	for ret*ret > x {
		ret = (ret + x/ret) / 2
	}
	return ret
}

// 常规二分查找，递归
func mySqrt1(x int) int {
	left, right := 0, x
	for {
		mid := (right-left)/2 + left
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			left = mid + 1
		}
	}
}

// @lc code=end