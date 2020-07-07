import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (56.04%)
 * Likes:    3512
 * Dislikes: 184
 * Total Accepted:    683.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 *
 * Example:
 *
 *
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * Note:
 *
 *
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	return groupAnagrams4(strs)
}

// 对第二种方法的精简写法
func groupAnagrams4(strs []string) [][]string {
	result := [][]string{}
	hash := map[[26]uint8]int{}

	for i := range strs {
		rCount := [26]uint8{}
		for j := range strs[i] {
			rCount[uint8(strs[i][j]-'a')]++
		}

		if v, ok := hash[rCount]; ok {
			result[v] = append(result[v], strs[i])
		} else {
			result = append(result, []string{strs[i]})
			// 巧妙设置了， 直接用数组为key进行设置，有点想不到，简化流程
			hash[rCount] = len(result) - 1
		}
	}
	return result
}

// 精简写法, 充分利用nil空值特性
func groupAnagrams3(strs []string) [][]string {
	hash := make(map[string][]string, 0)
	result := make([][]string, 0)
	for _, val := range strs {
		tmp := strings.Split(val, "")
		sort.Strings(tmp)
		sortStr := strings.Join(tmp, "")
		hash[sortStr] = append(hash[sortStr], val)
	}

	for _, val := range hash {
		result = append(result, val)
	}
	return result
}

// 先设置26位空间，用对应位字母减去'a'得到对应顺序位
func groupAnagrams2(strs []string) [][]string {
	hashVal := make(map[string][]string, 0)
	for _, v := range strs {
		hashCharcaters := make([]int, 26)
		for _, k := range v {
			hashCharcaters[k-'a']++
		}

		var key string
		for i := 0; i < len(hashCharcaters); i++ {
			key = key + strconv.Itoa(hashCharcaters[i]) + "#"
		}

		if child, ok := hashVal[key]; ok {
			hashVal[key] = append(child, v)
		} else {
			hashVal[key] = []string{v}
		}
	}

	retVal := make([][]string, len(hashVal))

	index := 0
	for _, val := range hashVal {
		retVal[index] = val
		index++
	}
	return retVal
}

// 方法1：对每个字符进行排序做为key进行比较，再放入map中缓存
func groupAnagrams1(strs []string) [][]string {
	hash := make(map[string][]string, 0)
	for _, v := range strs {
		sortV := sortString(v)
		if child, ok := hash[sortV]; ok {
			hash[sortV] = append(child, v)
		} else {
			hash[sortV] = []string{v}
		}
	}

	retVal := make([][]string, len(hash))

	index := 0
	for _, val := range hash {
		retVal[index] = val
		index++
	}
	return retVal
}

// 对string字符进行排序
// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
func sortString(s string) string {
	tmp := strings.Split(s, "")
	sort.Strings(tmp)
	return strings.Join(tmp, "")
}

// @lc code=end

