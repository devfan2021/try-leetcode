import (
	"strings"
)

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (35.79%)
 * Likes:    1749
 * Dislikes: 4733
 * Total Accepted:    472.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */

// @lc code=start
func convert(s string, numRows int) string {
	return convert1(s, numRows)
}

// line 0: indexK = 2 * numRows -2
// numRows -1, indexK = k(2*numRows-2) + numRows -1
// line i: k, 2 * numRows -2 + i  or  (k+1)(2*numRows-2)
/*n=numRows
Δ=2n-2    1                           2n-1                         4n-3
Δ=        2                     2n-2  2n                    4n-4   4n-2
Δ=        3               2n-3        2n+1              4n-5       .
Δ=        .           .               .               .            .
Δ=        .       n+2                 .           3n               .
Δ=        n-1 n+1                     3n-3    3n-1                 5n-5
Δ=2n-2    n                           3n-2                         5n-4
*/
func convert2(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	period := 2*numRows - 2
	res := make([]string, numRows)

	for i, v := range s {
		mod := i % period
		// fmt.Printf("i: %d, mod: %d, period: %d\n",i, mod, period)
		if mod < numRows {
			res[mod] += string(v)
		} else {
			res[period-mod] += string(v)
		}
	}
	return strings.Join(res, "")
}

func convert1(s string, numRows int) string {
	if len(s) == 0 || numRows == 0 || numRows == 1 {
		return s
	}

	rows := min(len(s), numRows)
	retVals := make([]string, rows)

	curRow, goingDown := 0, false
	for i := range s {
		retVals[curRow] += string(s[i : i+1])
		if curRow == 0 || curRow == rows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	vals := ""
	for _, val := range retVals {
		vals += val
	}
	return vals

	// return strings.Join(retVals, "")
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// @lc code=end