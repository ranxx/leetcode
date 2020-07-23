package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes2(nums []int) {
	// 两个指针，一个指向0，另一个指向非零
	var zero, nonZero int
	numsLen := len(nums)
	for {
		for zero < numsLen && (nums[zero] != 0) {
			zero++
		}
		for nonZero < numsLen && (nums[nonZero] == 0) {
			nonZero++
		}
		if !(zero < numsLen && nonZero < numsLen) || zero == numsLen-1 {
			break
		}

		// 如果0在后面，意味需要找0后面的数字
		if zero > nonZero {
			nonZero = zero
			continue
		}

		nums[zero], nums[nonZero] = nums[nonZero], nums[zero]
	}
}

func moveZeroes(nums []int) {
	// 把所有非零的放前面，后面的补上零
	var nonZero int
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		nums[nonZero] = nums[i]
		nonZero++
	}
	for ; nonZero < len(nums); nonZero++ {
		nums[nonZero] = 0
	}
}

// @lc code=end

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0, 1, 0, 3, 12, 0, 111}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1, 0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0, 1}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1, 0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}
