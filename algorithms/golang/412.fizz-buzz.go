/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (65.61%)
 * Likes:    223
 * Dislikes: 39
 * Total Accepted:    582.9K
 * Total Submissions: 881.2K
 * Testcase Example:  '3'
 *
 * Given an integer n, return a string array answer (1-indexed) where:
 *
 *
 * answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
 * answer[i] == "Fizz" if i is divisible by 3.
 * answer[i] == "Buzz" if i is divisible by 5.
 * answer[i] == i (as a string) if none of the above conditions are true.
 *
 *
 *
 * Example 1:
 * Input: n = 3
 * Output: ["1","2","Fizz"]
 * Example 2:
 * Input: n = 5
 * Output: ["1","2","Fizz","4","Buzz"]
 * Example 3:
 * Input: n = 15
 * Output:
 * ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */

// @lc code=start
func fizzBuzz(n int) []string {
	return solution1(n)
}

func solution1(n int) []string {
	retVal := make([]string, n)
	for index := 0; index < n; index++ {
		retVal[index] = fmt.Sprintf("%d", index+1)
	}

	step := 15
	for step <= n {
		retVal[step-1] = "FizzBuzz"
		step += 15
	}

	step = 3
	for step <= n {
		if retVal[step-1] != "FizzBuzz" {
			retVal[step-1] = "Fizz"
		}
		step += 3
	}

	step = 5
	for step <= n {
		if retVal[step-1] != "FizzBuzz" {
			retVal[step-1] = "Buzz"
		}
		step += 5
	}
	return retVal
}

func solution2(n int) []string {
	var retVal []string
	for index := 1; index <= n; index++ {
		if index%3 == 0 && index%5 == 0 {
			retVal = append(retVal, "FizzBuzz")
		} else if index%3 == 0 {
			retVal = append(retVal, "Fizz")
		} else if index%5 == 0 {
			retVal = append(retVal, "Buzz")
		} else {
			retVal = append(retVal, fmt.Sprintf("%d", index))
		}
	}
	return retVal
}

// @lc code=end