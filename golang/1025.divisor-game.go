/*
 * @lc app=leetcode id=1025 lang=golang
 *
 * [1025] Divisor Game
 *
 * https://leetcode.com/problems/divisor-game/description/
 *
 * algorithms
 * Easy (66.13%)
 * Likes:    444
 * Dislikes: 1217
 * Total Accepted:    61.9K
 * Total Submissions: 93.3K
 * Testcase Example:  '2'
 *
 * Alice and Bob take turns playing a game, with Alice starting first.
 *
 * Initially, there is a number N on the chalkboard.  On each player's turn,
 * that player makes a move consisting of:
 *
 *
 * Choosing any x with 0 < x < N and N % x == 0.
 * Replacing the number N on the chalkboard with N - x.
 *
 *
 * Also, if a player cannot make a move, they lose the game.
 *
 * Return True if and only if Alice wins the game, assuming both players play
 * optimally.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: true
 * Explanation: Alice chooses 1, and Bob has no more moves.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: false
 * Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more
 * moves.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 1000
 *
 *
 *
 */

// @lc code=start
func divisorGame(N int) bool {
	return divisorGame1(N)
}

// time complexity: O(1), space complexity: O(1)
func divisorGame1(N int) bool {
	return N%2 == 0
}

// @lc code=end