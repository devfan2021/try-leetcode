/*
 * @lc app=leetcode id=841 lang=golang
 *
 * [841] Keys and Rooms
 *
 * https://leetcode.com/problems/keys-and-rooms/description/
 *
 * algorithms
 * Medium (63.74%)
 * Likes:    987
 * Dislikes: 80
 * Total Accepted:    76K
 * Total Submissions: 118.4K
 * Testcase Example:  '[[1],[2],[3],[]]'
 *
 * There are N rooms and you start in room 0.  Each room has a distinct number
 * in 0, 1, 2, ..., N-1, and each room may have some keys to access the next
 * room.
 *
 * Formally, each room i has a list of keys rooms[i], and each key rooms[i][j]
 * is an integer in [0, 1, ..., N-1] where N = rooms.length.  A key rooms[i][j]
 * = v opens the room with number v.
 *
 * Initially, all the rooms start locked (except for room 0).
 *
 * You can walk back and forth between rooms freely.
 *
 * Return true if and only if you can enter every room.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1],[2],[3],[]]
 * Output: true
 * Explanation:
 * We start in room 0, and pick up key 1.
 * We then go to room 1, and pick up key 2.
 * We then go to room 2, and pick up key 3.
 * We then go to room 3.  Since we were able to go to every room, we return
 * true.
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,3],[3,0,1],[2],[0]]
 * Output: false
 * Explanation: We can't enter the room with number 2.
 *
 *
 * Note:
 *
 *
 * 1 <= rooms.length <= 1000
 * 0 <= rooms[i].length <= 1000
 * The number of keys in all rooms combined is at most 3000.
 *
 *
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	return canVisitAllRooms2(rooms)
}

// dfs
func canVisitAllRooms2(rooms [][]int) bool {
	visited := make(map[int]bool)
	roomCount := len(rooms)
	stack := []int{0}
	for len(stack) > 0 {
		n := len(stack) - 1
		currRoom := stack[n]
		stack = stack[:n]

		if _, ok := visited[currRoom]; !ok {
			roomCount--
			visited[currRoom] = true
			for _, room := range rooms[currRoom] {
				stack = append(stack, room)
			}
		}
	}
	return roomCount == 0
}

// bfs
func canVisitAllRooms1(rooms [][]int) bool {
	queue := []int{0}
	visited := make(map[int]bool)
	visited[0] = true
	for len(queue) > 0 {
		nextQueue := []int{}
		for _, v := range queue {
			items := rooms[v]
			// add index number in each rooms[v]
			items = append(items, v)
			for _, item := range items {
				if _, ok := visited[item]; !ok {
					visited[item] = true
					nextQueue = append(nextQueue, item)
				}
			}
		}
		if len(visited) == len(rooms) {
			return true
		} else {
			queue = nextQueue
		}
	}
	return false
}

// @lc code=end