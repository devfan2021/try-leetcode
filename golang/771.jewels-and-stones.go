import "strings"

/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 *
 * https://leetcode.com/problems/jewels-and-stones/description/
 *
 * algorithms
 * Easy (86.13%)
 * Likes:    2095
 * Dislikes: 366
 * Total Accepted:    512.6K
 * Total Submissions: 594.7K
 * Testcase Example:  '"aA"\n"aAAbbbb"'
 *
 * You're given strings J representing the types of stones that are jewels, and
 * S representing the stones you have.  Each character in S is a type of stone
 * you have.  You want to know how many of the stones you have are also
 * jewels.
 *
 * The letters in J are guaranteed distinct, and all characters in J and S are
 * letters. Letters are case sensitive, so "a" is considered a different type
 * of stone from "A".
 *
 * Example 1:
 *
 *
 * Input: J = "aA", S = "aAAbbbb"
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: J = "z", S = "ZZ"
 * Output: 0
 *
 *
 * Note:
 *
 *
 * S and J will consist of letters and have length at most 50.
 * The characters in J are distinct.
 *
 *
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	return numJewelsInStones2(J, S)
}

func numJewelsInStones2(J string, S string) int {
	count := 0
	for _, v := range J {
		count += strings.Count(S, string(v))
	}
	return count
}

// 利用hashMap
func numJewelsInStones1(J string, S string) int {
	// hash := make(map[rune]bool, 0)
	hash := make(map[string]bool, 0)
	for _, v := range J {
		// hash[v-'A'] = true
		hash[string(v)] = true
	}
	count := 0

	for _, v := range S {
		// if _, ok := hash[v-'A']; ok {
		if _, ok := hash[string(v)]; ok {
			count++
		}
	}
	return count
}

// @lc code=end