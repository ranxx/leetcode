package main

import "fmt"

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (61.28%)
 * Likes:    1125
 * Dislikes: 0
 * Total Accepted:    139.1K
 * Total Submissions: 227K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 */

// @lc code=start
func maxArea(height []int) int {
	// 从I左右两边收敛，哪边低收敛哪边，左右夹逼
	if len(height) <= 0 {
		return 0
	}
	begin, end := 0, len(height)-1
	max := 0
	for begin < end {
		area := (end - begin) * min(height[begin], height[end])
		if area > max {
			max = area
		}
		if height[begin] > height[end] {
			end--
		} else {
			begin++
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

	height = []int{1, 1}
	fmt.Println(maxArea(height))

	height = []int{1, 2}
	fmt.Println(maxArea(height))

	height = []int{1, 2, 1}
	fmt.Println(maxArea(height))

	height = []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(height))

}
