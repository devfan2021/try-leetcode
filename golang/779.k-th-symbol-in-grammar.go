/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 *
 * https://leetcode.com/problems/k-th-symbol-in-grammar/description/
 *
 * algorithms
 * Medium (36.89%)
 * Likes:    482
 * Dislikes: 145
 * Total Accepted:    33.1K
 * Total Submissions: 89.2K
 * Testcase Example:  '1\n1'
 *
 * On the first row, we write a 0. Now in every subsequent row, we look at the
 * previous row and replace each occurrence of 0 with 01, and each occurrence
 * of 1 with 10.
 *
 * Given row N and index K, return the K-th indexed symbol in row N. (The
 * values of K are 1-indexed.) (1 indexed).
 *
 *
 * Examples:
 * Input: N = 1, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 2
 * Output: 1
 *
 * Input: N = 4, K = 5
 * Output: 1
 *
 * Explanation:
 * row 1: 0
 * row 2: 01
 * row 3: 0110
 * row 4: 01101001
 *
 *
 * Note:
 *
 *
 * N will be an integer in the range [1, 30].
 * K will be an integer in the range [1, 2^(N-1)].
 *
 *
 */

// @lc code=start
func kthGrammar(N int, K int) int {
	return kthGrammar2(N, K)
}

// should restudy
// https://leetcode.com/problems/k-th-symbol-in-grammar/discuss/285399/GoLang-Using-recursion
func kthGrammar3(N int, K int) int {
	if K == 1 {
		return 0
	}
	g := kthGrammar3(N-1, (K+1)/2)
	if g == 0 {
		if K%2 != 0 {
			return 0
		}
		return 1
	} else {
		if K%2 != 0 {
			return 1
		}
		return 0
	}
}

// should restudy, runtime error, out memory
func kthGrammar2(N int, K int) int {
	if N == 1 {
		return 0
	}
	return 1 - kthGrammar2(N, (K+1)/2) ^ (K % 2)
}

// solution by recursive, runtime error, out memory
func kthGrammar1(N int, K int) int {
	vals := helper(N)
	if K > len(vals) {
		return -1
	}
	return vals[K-1]
}

func helper(N int) []int {
	if N == 1 {
		return []int{0}
	}
	hash := map[int]int{
		0: 1,
		1: 0,
	}
	vals := helper(N - 1)
	retVals := vals[0:]
	for _, v := range vals {
		retVals = append(retVals, hash[v])
	}
	return retVals
}

// @lc code=end