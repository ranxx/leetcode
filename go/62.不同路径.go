package main

import "fmt"

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (60.99%)
 * Likes:    615
 * Dislikes: 0
 * Total Accepted:    125K
 * Total Submissions: 203.9K
 * Testcase Example:  '3\n2'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 *
 *
 *
 * 示例 1:
 *
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 *
 *
 * 示例 2:
 *
 * 输入: m = 7, n = 3
 * 输出: 28
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 100
 * 题目数据保证答案小于等于 2 * 10 ^ 9
 *
 *
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	visit := make([]int, m*(n+1))
	return uniquePathsByIteration(m, n, visit)
	return uniquePathsByRecursion(m, n, m, visit)
}

func uniquePathsByRecursion(m, n, totalCol int, visit []int) int {
	if n == 0 && m == 0 {
		return 0
	}
	if n-1 == 0 || m-1 == 0 {
		return 1
	}
	v := visit[totalCol*(n-1)+m]
	if v != 0 {
		return v
	}
	visit[totalCol*(n-1)+m] = uniquePathsByRecursion(m-1, n, totalCol, visit) + uniquePathsByRecursion(m, n-1, totalCol, visit)
	return visit[totalCol*(n-1)+m]
}

func uniquePathsByIteration(m, n int, visit []int) int {
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if i == 0 || j == 0 {
				visit[m*j+i] = 1
				continue
			}
			visit[m*j+i] = visit[m*(j-1)+i] + visit[m*j+i-1]
		}
	}
	return visit[m*(n-1)+m-1]
}

// @lc code=end

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(1, 2))
	fmt.Println(uniquePaths(1, 1))
	// fmt.Println(uniquePaths(51, 9))
}
