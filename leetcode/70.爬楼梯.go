package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (47.86%)
 * Likes:    840
 * Dislikes: 0
 * Total Accepted:    135.1K
 * Total Submissions: 282.3K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	// 斐波那契(Fibonacci)数列
	if n <= 2 {
		return n
	}
	n1, n2, total := 1, 2, 0
	for i := 3; i <= n; i++ {
		total = n1 + n2
		n1 = n2
		n2 = total
	}
	return total
}

// @lc code=end

func climbStairsV2(n int) int {
	// 递归法求斐波那契数列
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairsV2(n-1) + climbStairsV2(n-2)
}

func main() {
	fmt.Println(climbStairsV2(4))
}
