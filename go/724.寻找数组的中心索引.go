package main

import "fmt"

/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 *
 * https://leetcode-cn.com/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (37.53%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    45.6K
 * Total Submissions: 121.4K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * 给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
 *
 * 我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
 *
 * 如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * nums = [1, 7, 3, 6, 5, 6]
 * 输出：3
 * 解释：
 * 索引 3 (nums[3] = 6) 的左侧数之和 (1 + 7 + 3 = 11)，与右侧数之和 (5 + 6 = 11) 相等。
 * 同时, 3 也是第一个符合要求的中心索引。
 *
 *
 * 示例 2：
 *
 * 输入：
 * nums = [1, 2, 3]
 * 输出：-1
 * 解释：
 * 数组中不存在满足此条件的中心索引。
 *
 *
 *
 * 说明：
 *
 *
 * nums 的长度范围为 [0, 10000]。
 * 任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
 *
 *
 */

// @lc code=start
func pivotIndex(nums []int) int {
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		mid--
	}
	preOp := 0
	fmt.Println("init", mid)
	for mid > 0 && mid < len(nums)-1 {
		leftCount := plus(nums[:mid])
		rightCount := plus(nums[mid+1:])
		fmt.Println(mid, preOp, leftCount, rightCount)
		if leftCount == rightCount {
			return mid
		}
		if preOp != 2 && leftCount < rightCount {
			mid++
			preOp = 1
			continue
		} else if preOp != 1 && leftCount > rightCount {
			mid--
			preOp = 2
			continue
		}
		return -1
	}
	return -1
}
func plus(a []int) (ret int) {
	for _, v := range a {
		ret += v
	}
	return
}

// @lc code=end
func pivotIndexV2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		leftCount := plus(nums[:left+1])
		rightCount := plus(nums[right:])
		fmt.Println(left, right, leftCount, rightCount)
		if leftCount < 0 && rightCount < 0 {

		}
		if leftCount < rightCount {
			left++
			continue
		}
		if leftCount > rightCount {
			right--
			continue
		}
		if right-left == 2 {
			return left + 1
		}
		left++
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, 0, 1}))
}
