import (
	"strings"
)

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 *
 * https://leetcode.com/problems/simplify-path/description/
 *
 * algorithms
 * Medium (32.70%)
 * Likes:    878
 * Dislikes: 1873
 * Total Accepted:    220.1K
 * Total Submissions: 670.8K
 * Testcase Example:  '"/home/"'
 *
 * Given an absolute path for a file (Unix-style), simplify it. Or in other
 * words, convert it to the canonical path.
 *
 * In a UNIX-style file system, a period . refers to the current directory.
 * Furthermore, a double period .. moves the directory up a level.
 *
 * Note that the returned canonical path must always begin with a slash /, and
 * there must be only a single slash / between two directory names. The last
 * directory name (if it exists) must not end with a trailing /. Also, the
 * canonical path must be the shortest string representing the absolute
 * path.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "/home/"
 * Output: "/home"
 * Explanation: Note that there is no trailing slash after the last directory
 * name.
 *
 *
 * Example 2:
 *
 *
 * Input: "/../"
 * Output: "/"
 * Explanation: Going one level up from the root directory is a no-op, as the
 * root level is the highest level you can go.
 *
 *
 * Example 3:
 *
 *
 * Input: "/home//foo/"
 * Output: "/home/foo"
 * Explanation: In the canonical path, multiple consecutive slashes are
 * replaced by a single one.
 *
 *
 * Example 4:
 *
 *
 * Input: "/a/./b/../../c/"
 * Output: "/c"
 *
 *
 * Example 5:
 *
 *
 * Input: "/a/../../b/../c//.//"
 * Output: "/c"
 *
 *
 * Example 6:
 *
 *
 * Input: "/a//b////c/d//././/.."
 * Output: "/a/b/c"
 *
 *
 */

// @lc code=start
func simplifyPath(path string) string {
	return simplifyPath1(path)
}

func simplifyPath1(path string) string {
	if len(path) == 0 {
		return ""
	}

	paths := strings.Split(path, "/")
	stack := []string{}

	for _, v := range paths {
		if v == "" || len(v) == 0 || v == "." {
			continue
		} else if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}

	if len(stack) == 0 {
		return "/"
	} else {
		return "/" + strings.Join(stack, "/")
	}
}

// @lc code=end