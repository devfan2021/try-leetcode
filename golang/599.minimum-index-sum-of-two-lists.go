/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 *
 * https://leetcode.com/problems/minimum-index-sum-of-two-lists/description/
 *
 * algorithms
 * Easy (50.49%)
 * Likes:    580
 * Dislikes: 200
 * Total Accepted:    89.8K
 * Total Submissions: 177.6K
 * Testcase Example:  '["Shogun","Tapioca Express","Burger King","KFC"]\n' +
  '["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]'
 *
 *
 * Suppose Andy and Doris want to choose a restaurant for dinner, and they both
 * have a list of favorite restaurants represented by strings.
 *
 *
 * You need to help them find out their common interest with the least list
 * index sum. If there is a choice tie between answers, output all of them with
 * no order requirement. You could assume there always exists an answer.
 *
 *
 *
 * Example 1:
 *
 * Input:
 * ["Shogun", "Tapioca Express", "Burger King", "KFC"]
 * ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse",
 * "Shogun"]
 * Output: ["Shogun"]
 * Explanation: The only restaurant they both like is "Shogun".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * ["Shogun", "Tapioca Express", "Burger King", "KFC"]
 * ["KFC", "Shogun", "Burger King"]
 * Output: ["Shogun"]
 * Explanation: The restaurant they both like and have the least index sum is
 * "Shogun" with index sum 1 (0+1).
 *
 *
 *
 *
 * Note:
 *
 * The length of both lists will be in the range of [1, 1000].
 * The length of strings in both lists will be in the range of [1, 30].
 * The index is starting from 0 to the list length minus 1.
 * No duplicates in both lists.
 *
 *
*/

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	sum := len(list1) + len(list2) - 2
	lsMap := make(map[string]int)
	for i, v := range list1 {
		lsMap[v] = i
	}
	retArr := make([]string, 0)

	for i, v := range list2 {
		index, ok := lsMap[v]
		if ok {
			if i+index < sum {
				sum = i + index
				retArr = []string{v}
			} else if i+index == sum {
				retArr = append(retArr, v)
			}
		}
	}

	return retArr
}

// @lc code=end