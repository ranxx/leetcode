package main

import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (38.69%)
 * Likes:    460
 * Dislikes: 0
 * Total Accepted:    30.8K
 * Total Submissions: 78.8K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 *
 *
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 *
 *
 *
 *
 *
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 *
 *
 *
 * 示例:
 *
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 *
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	// 栈
	s := newStack(len(heights))
	s.push(-1)
	maxArea := 0
	for i := range heights {
		for s.top() != -1 && heights[s.top()] >= heights[i] {
			maxArea = max(heights[s.pop()]*(i-s.top()-1), maxArea)
		}
		s.push(i)
	}

	for s.top() != -1 {
		maxArea = max(heights[s.pop()]*(len(heights)-s.top()-1), maxArea)
	}
	return maxArea
}

type stack struct {
	data     []int
	len, cap int
}

func (s *stack) String() string {
	return fmt.Sprintf("data:%v, len:%d, cap:%d", s.data, s.len, s.cap)
}

// func (s *stack) len() int {
// 	return s.len
// }
func (s *stack) top() int {
	if s.len > 0 {
		return s.data[s.len-1]
	}
	return -1
}
func (s *stack) pop() int {
	if s.len > 0 {
		s.len--
		return s.data[s.len]
	}
	return 0
}
func (s *stack) push(val int) {
	s.data[s.len] = val
	s.len++
}
func newStack(cap int) *stack {
	return &stack{
		data: make([]int, cap+1),
		cap:  cap,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	// fmt.Println(largestRectangleArea([]int{}))
	// fmt.Println(largestRectangleArea([]int{0}))
	// fmt.Println(largestRectangleArea([]int{2}))
	// fmt.Println(largestRectangleArea([]int{2, 2}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}
