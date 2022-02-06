/*
 * @lc app=leetcode id=717 lang=golang
 *
 * [717] 1-bit and 2-bit Characters
 *
 * https://leetcode.com/problems/1-bit-and-2-bit-characters/description/
 *
 * algorithms
 * Easy (46.28%)
 * Likes:    648
 * Dislikes: 1663
 * Total Accepted:    100K
 * Total Submissions: 216.3K
 * Testcase Example:  '[1,0,0]'
 *
 * We have two special characters:
 *
 *
 * The first character can be represented by one bit 0.
 * The second character can be represented by two bits (10 or 11).
 *
 *
 * Given a binary array bits that ends with 0, return true if the last
 * character must be a one-bit character.
 *
 *
 * Example 1:
 *
 *
 * Input: bits = [1,0,0]
 * Output: true
 * Explanation: The only way to decode it is two-bit character and one-bit
 * character.
 * So the last character is one-bit character.
 *
 *
 * Example 2:
 *
 *
 * Input: bits = [1,1,1,0]
 * Output: false
 * Explanation: The only way to decode it is two-bit character and two-bit
 * character.
 * So the last character is not one-bit character.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= bits.length <= 1000
 * bits[i] is either 0 or 1.
 *
 *
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	return soluntion2(bits)
}

func soluntion1(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	index := 0
	for index < len(bits) {
		if bits[index] == 1 {
			index += 2
			continue
		}
		if bits[index] == 0 {
			if index == len(bits)-1 {
				return true
			}
			index += 1
			continue
		}
	}
	return false
}

func soluntion2(bits []int) bool {
	l := len(bits)
	for index := 0; index < l; index++ {
		if index == l-1 {
			return true
		}
		if bits[index] == 1 {
			index += 1
		}
	}
	return false
}

// @lc code=end