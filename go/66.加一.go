package main

import "fmt"

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (42.98%)
 * Likes:    430
 * Dislikes: 0
 * Total Accepted:    122.5K
 * Total Submissions: 283.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */

// @lc code=start
func plusOne(digits []int) []int {
	plus := false
	digits[len(digits)-1]++
	// ret := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		if plus {
			digits[i]++
		}
		if digits[i] < 10 {
			break
		}
		plus = true
		digits[i] -= 10
	}
	return digits
}

// @lc code=end

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 9}))
	fmt.Println(plusOne([]int{9, 9}))
}
