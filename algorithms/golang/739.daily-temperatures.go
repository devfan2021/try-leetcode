/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 *
 * https://leetcode.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (62.91%)
 * Likes:    2856
 * Dislikes: 87
 * Total Accepted:    162.6K
 * Total Submissions: 257.3K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 *
 * Given a list of daily temperatures T, return a list such that, for each day
 * in the input, tells you how many days you would have to wait until a warmer
 * temperature.  If there is no future day for which this is possible, put 0
 * instead.
 *
 * For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76,
 * 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
 *
 *
 * Note:
 * The length of temperatures will be in the range [1, 30000].
 * Each temperature will be an integer in the range [30, 100].
 *
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	return dailyTemperatures2(T)
}

// defin Stack type
type Stack struct {
	data []int
	size int
}

func NewStack(capacity int) *Stack {
	return &Stack{data: make([]int, capacity)}
}

func (s *Stack) Push(value int) bool {
	if s.IsFull() {
		return false
	}
	s.data[s.size] = value
	s.size++
	return true
}

func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.size--
	return true
}

func (s *Stack) Front() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[s.size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size > len(s.data)-1
}

// use stack, time complexity:O(n), space complexity:O(n)
func dailyTemperatures2(T []int) []int {
	retVal := make([]int, len(T))
	stack := NewStack(len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && T[i] >= T[stack.Front()] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			retVal[i] = 0
		} else {
			retVal[i] = stack.Front() - i
		}
		stack.Push(i)
	}
	return retVal
}

// two lay iterate, time complexity:O(n*n), space complexity:O(n)
func dailyTemperatures1(T []int) []int {
	retVal := make([]int, len(T))

	for i := 0; i < len(T); i++ {
		flag := false
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				retVal[i] = j - i
				flag = true
				break
			}
		}
		if !flag {
			retVal[i] = 0
		}
	}
	return retVal
}

// @lc code=end