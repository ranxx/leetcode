package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (40.43%)
 * Likes:    497
 * Dislikes: 0
 * Total Accepted:    100.4K
 * Total Submissions: 247.6K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5,6,7] 和 k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 * 输入: [-1,-100,3,99] 和 k = 2
 * 输出: [3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 * 说明:
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 要求使用空间复杂度为 O(1) 的 原地 算法。
 *
 *
 */

// @lc code=start
func rotate(nums []int, k int) {
	for n := 0; n < k%len(nums); n++ {
		tmp := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp
	}
}

func rotateV3(nums []int, k int) {
	if len(nums) <= 0 || k <= 0 || k%len(nums) == 0 {
		return
	}
	k %= len(nums)
	rotateArray(nums)
	rotateArray(nums[:k])
	rotateArray(nums[k:])
}

func rotateArray(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func rotateV4(nums []int, k int) {
	if len(nums) <= 0 || k <= 0 || k%len(nums) == 0 {
		return
	}
	k %= len(nums)
	for i, step := 0, 0; step < len(nums); i++ {
		start, prev, next := i, nums[i], (i+k)%len(nums)
		for start != next {
			nums[next], prev = prev, nums[next]
			next = (next + k) % len(nums)
			step++
		}
		step++
		nums[start] = prev
	}
}

// @lc code=end

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateV3(nums, 1)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotateV3(nums, 3)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotateV4(nums, 3)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotateV4(nums, 1)
	fmt.Println(nums)
}
