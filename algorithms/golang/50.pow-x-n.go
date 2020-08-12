import "math"

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (29.56%)
 * Likes:    1605
 * Dislikes: 3129
 * Total Accepted:    502.8K
 * Total Submissions: 1.7M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 *
 * Example 1:
 *
 *
 * Input: 2.00000, 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: 2.10000, 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Note:
 *
 *
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 *
 *
 */

// @lc code=start

func myPow(x float64, n int) float64 {
	return myPow3(x, n)
}

// iterative, solution like myPow1
func myPow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	retVal := 1.0
	for n > 0 {
		if n&1 == 1 {
			retVal *= x
		}
		x *= x
		n >>= 1
	}
	return retVal
}

// solution by recursive with memorization
func myPow2(x float64, n int) float64 {
	visited := map[int]float64{}
	N := int(math.Abs(float64(n)))
	val := pow(x, N, visited)
	if n < 0 {
		return 1 / val
	}
	return val
}

func pow(x float64, n int, v map[int]float64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if val, ok := v[n]; ok {
		return val
	}
	// reduce half question
	value := pow(x*x, int(n/2), v) * pow(x, int(n%2), v)
	v[n] = value
	return value
}

// solution by recursive
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	return myPow(x*x, int(n/2)) * myPow(x, int(n%2))
}

// @lc code=end