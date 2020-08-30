import "fmt"

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (55.14%)
 * Likes:    1175
 * Dislikes: 2719
 * Total Accepted:    377.6K
 * Total Submissions: 683.7K
 * Testcase Example:  '3'
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 *
 *
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * For example, two is written as II in Roman numeral, just two one's added
 * together. Twelve is written as, XII, which is simply X + II. The number
 * twenty seven is written as XXVII, which is XX + V + II.
 *
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 *
 *
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 *
 *
 * Given an integer, convert it to a roman numeral. Input is guaranteed to be
 * within the range from 1 to 3999.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: "III"
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "IV"
 *
 * Example 3:
 *
 *
 * Input: 9
 * Output: "IX"
 *
 * Example 4:
 *
 *
 * Input: 58
 * Output: "LVIII"
 * Explanation: L = 50, V = 5, III = 3.
 *
 *
 * Example 5:
 *
 *
 * Input: 1994
 * Output: "MCMXCIV"
 * Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 *
 */

// @lc code=start
func intToRoman(num int) string {
	return intToRoman3(num)
}

func intToRoman3(num int) string {
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	res, index := "", 0
	for num > 0 {
		if num-value[index] >= 0 {
			res += symbol[index]
			num -= value[index]
		} else {
			index++
		}
	}
	return res
}

type roman struct {
	num int
	s   string
}

func intToRoman2(num int) string {
	// should put large in front of array
	romans := []roman{
		roman{num: 1000, s: "M"}, roman{num: 900, s: "CM"}, roman{num: 500, s: "D"}, roman{num: 400, s: "CD"},
		roman{num: 100, s: "C"}, roman{num: 90, s: "XC"}, roman{num: 50, s: "L"}, roman{num: 40, s: "XL"},
		roman{num: 10, s: "X"}, roman{num: 9, s: "IX"}, roman{num: 5, s: "V"}, roman{num: 4, s: "IV"}, roman{num: 1, s: "I"},
	}

	result := ""
	for i := 0; i < len(romans); i++ {
		tmp := num / romans[i].num
		if tmp > 0 {
			for j := 0; j < tmp; j++ {
				result += romans[i].s
			}
		}

		num = num - tmp*romans[i].num
		if num == 0 {
			break
		}
	}
	return result
}

func intToRoman1(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return fmt.Sprintf("%s%s%s%s", M[num/1000], C[(num%1000)/100], X[(num%100)/10], I[num%10])
}

// @lc code=end