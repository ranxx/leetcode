package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (36.06%)
 * Likes:    385
 * Dislikes: 0
 * Total Accepted:    89.6K
 * Total Submissions: 246.8K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return uniquePathsWithObstaclesScrollArray(len(obstacleGrid), len(obstacleGrid[0]), obstacleGrid)
	return uniquePathsWithObstaclesIteration(len(obstacleGrid), len(obstacleGrid[0]), obstacleGrid)
}

func uniquePathsWithObstaclesIteration(n, m int, obstacleGrid [][]int) int {
	visit := make([]int, n*m)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if obstacleGrid[j][i] == 1 {
				continue
			}
			if j == 0 && i == 0 {
				visit[m*j+i] = 1
				continue
			}
			if i == 0 {
				visit[m*j+i] = visit[m*(j-1)+i]
				continue
			} else if j == 0 {
				visit[m*j+i] = visit[m*j+i-1]
				continue
			}
			visit[m*j+i] = visit[m*(j-1)+i] + visit[m*j+i-1]
		}
	}
	return visit[m*(n-1)+m-1]
}

func uniquePathsWithObstaclesScrollArray(n, m int, obstacleGrid [][]int) int {
	visit := make([]int, m)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if obstacleGrid[j][i] == 1 {
				visit[i] = 0
				continue
			}
			if j == 0 && i == 0 {
				visit[i] = 1
				continue
			}
			if j == 0 {
				visit[i] = visit[i-1]
				continue
			}
			if i == 0 {
				continue
			}
			visit[i] = visit[i] + visit[i-1]
		}
	}
	return visit[m-1]
}

// @lc code=end

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {0, 0}}))
}
