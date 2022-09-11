import java.util.List;

/*
 * @lc app=leetcode id=57 lang=java
 *
 * [57] Insert Interval
 *
 * https://leetcode.com/problems/insert-interval/description/
 *
 * algorithms
 * Medium (36.84%)
 * Likes:    5851
 * Dislikes: 400
 * Total Accepted:    591.6K
 * Total Submissions: 1.6M
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * You are given an array of non-overlapping intervals intervals where
 * intervals[i] = [starti, endi] represent the start and the end of the i^th
 * interval and intervals is sorted in ascending order by starti. You are also
 * given an interval newInterval = [start, end] that represents the start and
 * end of another interval.
 * 
 * Insert newInterval into intervals such that intervals is still sorted in
 * ascending order by starti and intervals still does not have any overlapping
 * intervals (merge overlapping intervals if necessary).
 * 
 * Return intervals after the insertion.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with
 * [3,5],[6,7],[8,10].
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^5
 * intervals is sorted by starti in ascending order.
 * newInterval.length == 2
 * 0 <= start <= end <= 10^5
 * 
 * 
 */

// @lc code=start
class Solution {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        return solution2(intervals, newInterval);
    }

    public int[][] solution1(int[][] intervals, int[] newInterval) {
        List<int[]> retVals = new ArrayList();
        int i = 0;
        while (i < intervals.length && intervals[i][1] < newInterval[0]) {
            retVals.add(intervals[i]);
            i++;
        }
        while (i < intervals.length && intervals[i][0] <= newInterval[1]) {
            newInterval[0] = Math.min(newInterval[0], intervals[i][0]);
            newInterval[1] = Math.max(newInterval[1], intervals[i][1]);
            i++;
        }
        retVals.add(newInterval);
        while (i < intervals.length) {
            retVals.add(intervals[i]);
            i++;
        }
        return retVals.toArray(new int[0][]);
    }

    public int[][] solution2(int[][] intervals, int[] newInterval) {
        List<int[]> retVals = new ArrayList();
        int i = 0;
        int start = newInterval[0];
        int end = newInterval[1];
        while (i < intervals.length && intervals[i][1] < start) {
            retVals.add(intervals[i]);
            i++;
        }
        while (i < intervals.length && intervals[i][0] <= end) {
            start = Math.min(intervals[i][0], start);
            end = Math.max(intervals[i][1], end);
            i++;
        }
        retVals.add(new int[] { start, end });
        while (i < intervals.length) {
            retVals.add(intervals[i]);
            i++;
        }

        int rets[][] = new int[retVals.size()][];
        for (int j = 0; j < retVals.size(); j++) {
            rets[j] = retVals.get(j);
        }
        return rets;
    }
}
// @lc code=end