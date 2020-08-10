/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 *
 * https://leetcode.com/problems/4sum-ii/description/
 *
 * algorithms
 * Medium (52.88%)
 * Likes:    1157
 * Dislikes: 72
 * Total Accepted:    108.5K
 * Total Submissions: 204.6K
 * Testcase Example:  '[1,2]\n[-2,-1]\n[-1,2]\n[0,2]'
 *
 * Given four lists A, B, C, D of integer values, compute how many tuples (i,
 * j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
 *
 * To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤
 * N ≤ 500. All integers are in the range of -2^28 to 2^28 - 1 and the result
 * is guaranteed to be at most 2^31 - 1.
 *
 * Example:
 *
 *
 * Input:
 * A = [ 1, 2]
 * B = [-2,-1]
 * C = [-1, 2]
 * D = [ 0, 2]
 *
 * Output:
 * 2
 *
 * Explanation:
 * The two tuples are:
 * 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 *
 *
 *
 *
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	return fourSumCount2(A, B, C, D)
}

// 每2个数组进行合并累加逻辑进行优化，不需要进行多一次的hash及遍历累加
func fourSumCount2(A []int, B []int, C []int, D []int) int {
	hash := make(map[int]int, len(A)*len(B)) // 尝试开个大空间进行存储
	// hash := make(map[int]int, 0) 类似这种方式，效率会差很多, map会不停的扩展大小
	for _, a := range A {
		for _, b := range B {
			hash[a+b]++
		}
	}

	count := 0
	for _, c := range C {
		for _, d := range D {
			count += hash[0-c-d]
		}
	}
	return count
}

// 直接4次for循环超时, 每2个数组进行合并累加放结果值到hash中，再用两个hash进行比较
func fourSumCount1(A []int, B []int, C []int, D []int) int {
	hash1 := map[int]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			key := A[i] + B[j]
			hash1[key]++
		}
	}

	hash2 := map[int]int{}
	for m := 0; m < len(C); m++ {
		for n := 0; n < len(D); n++ {
			key := C[m] + D[n]
			hash2[key]++
		}
	}

	count := 0
	for k, v := range hash1 {
		if v2, ok := hash2[-k]; ok {
			count += v * v2
		}
	}

	return count
}

// @lc code=end

