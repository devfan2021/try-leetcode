/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 *
 * https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
 *
 * algorithms
 * Easy (45.23%)
 * Likes:    394
 * Dislikes: 507
 * Total Accepted:    74.1K
 * Total Submissions: 163.5K
 * Testcase Example:  '["c","f","j"]\n"a"'
 *
 *
 * Given a list of sorted characters letters containing only lowercase letters,
 * and given a target letter target, find the smallest element in the list that
 * is larger than the given target.
 *
 * Letters also wrap around.  For example, if the target is target = 'z' and
 * letters = ['a', 'b'], the answer is 'a'.
 *
 *
 * Examples:
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "a"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "c"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "d"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "g"
 * Output: "j"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "j"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "k"
 * Output: "c"
 *
 *
 *
 * Note:
 *
 * letters has a length in range [2, 10000].
 * letters consists of lowercase letters, and contains at least 2 unique
 * letters.
 * target is a lowercase letter.
 *
 *
 */

// @lc code=start
// 测试case: ["c","f","j"]\n"j"
func nextGreatestLetter(letters []byte, target byte) byte {
	return nextGreatestLetter2(letters, target)
}

// 采用Hash存储26个字母，再进行遍历比较
func nextGreatestLetter2(letters []byte, target byte) byte {
	hash := make(map[byte]bool, 26)
	for _, v := range letters {
		hash[v] = true
	}

	for {
		target++
		if target > 'z' {
			target = 'a'
		}
		if _, ok := hash[target]; ok {
			return byte(target)
		}
	}
}

// 采用二分查找的方式
func nextGreatestLetter1(letters []byte, target byte) byte {
	low, high := 0, len(letters)
	for low < high {
		mid := low + (high-low)/2
		if (letters[mid] - 'a') <= (target - 'a') {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return letters[low%len(letters)] // 求余，循环回到第一个 ["c", "f", "j"]\n"c"
}

// @lc code=end