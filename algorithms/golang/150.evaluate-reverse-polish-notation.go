import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (35.80%)
 * Likes:    1031
 * Dislikes: 471
 * Total Accepted:    229.5K
 * Total Submissions: 635.5K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * Evaluate the value of an arithmetic expression in Reverse Polish Notation.
 *
 * Valid operators are +, -, *, /. Each operand may be an integer or another
 * expression.
 *
 * Note:
 *
 *
 * Division between two integers should truncate toward zero.
 * The given RPN expression is always valid. That means the expression would
 * always evaluate to a result and there won't be any divide by zero
 * operation.
 *
 *
 * Example 1:
 *
 *
 * Input: ["2", "1", "+", "3", "*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 *
 *
 * Example 2:
 *
 *
 * Input: ["4", "13", "5", "/", "+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 *
 *
 * Example 3:
 *
 *
 * Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * Output: 22
 * Explanation:
 * ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 *
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	operatorHash := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for i := 0; i < len(tokens); i++ {
		item := tokens[i]
		if _, ok := operatorHash[item]; ok {
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			tmp := 0
			switch item {
			case "+":
				tmp = left + right
			case "-":
				tmp = left - right
			case "*":
				tmp = left * right
			case "/":
				tmp = left / right
			}
			stack[len(stack)-2] = tmp
			stack = stack[:len(stack)-1]
		} else {
			intVal, _ := strconv.Atoi(item)
			stack = append(stack, intVal)
		}
	}
	return stack[0]
}

// @lc code=end