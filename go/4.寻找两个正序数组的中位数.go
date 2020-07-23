package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (38.33%)
 * Likes:    2843
 * Dislikes: 0
 * Total Accepted:    217.6K
 * Total Submissions: 567.8K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
 *
 * 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 *
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

// @lc code=end

// 先排序在取中位数
func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	// 先排序
	tmp := append(nums1, nums2...)
	sort.Ints(tmp)
	if len(tmp)%2 != 0 {
		return float64(tmp[len(tmp)/2])
	}
	mid := len(tmp) / 2
	return float64(tmp[mid]+tmp[mid-1]) / 2.0
}

// 直接循环，len/2 次找到中间的数
func findMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
	_len := len(nums1) + len(nums2)
	start1, start2 := 0, 0
	left, right := 0, 0
	for i := 0; i < _len/2+1; i++ {
		left = right
		// num1 前进
		if len(nums1) == 0 && len(nums2) > start2 {
			right, start2 = nums2[start2], start2+1
			continue
		}
		if len(nums2) == 0 && len(nums1) > start1 {
			right, start1 = nums1[start1], start1+1
			continue
		}
		if len(nums1) > start1 && len(nums2) > start2 && nums1[start1] <= nums2[start2] {
			right, start1 = nums1[start1], start1+1
		} else if len(nums1) <= start1 && len(nums2) > start2 && nums1[start1-1] <= nums2[start2] {
			right, start2 = nums2[start2], start2+1
		} else if len(nums1) > start1 && len(nums2) > start2 && nums1[start1] > nums2[start2] {
			right, start2 = nums2[start2], start2+1
		} else if len(nums2) <= start2 && len(nums1) > start1 && nums1[start1] > nums2[start2-1] {
			right, start1 = nums1[start1], start1+1
		} else {
			fmt.Println("Error")
			break
		}
	}
	if _len%2 != 0 {
		return float64(right)
	}

	return float64(left+right) / 2.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{100001}, []int{100000}))
}
