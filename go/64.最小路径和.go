package main

import "fmt"

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (65.97%)
 * Likes:    553
 * Dislikes: 0
 * Total Accepted:    112.5K
 * Total Submissions: 168K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 * 示例:
 *
 * 输入:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * 输出: 7
 * 解释: 因为路径 1→3→1→1→1 的总和最小。
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	i, j := 0, 0
	for ; j < len(grid); j++ {
		for i = 0; i < len(grid[j]); i++ {
			if i == j && i == 0 {
				continue
			}
			if i == 0 {
				grid[j][i] += grid[j-1][i]
			} else if j == 0 {
				grid[j][i] += grid[j][i-1]
			} else {
				grid[j][i] += min(grid[j][i-1], grid[j-1][i])
			}
		}
	}
	return grid[j-1][i-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
