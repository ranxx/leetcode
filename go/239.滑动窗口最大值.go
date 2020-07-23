package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (43.77%)
 * Likes:    252
 * Dislikes: 0
 * Total Accepted:    32.2K
 * Total Submissions: 72.3K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口中的最大值。
 *
 *
 *
 * 示例:
 *
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7]
 * 解释:
 *
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 *
 * 提示：
 *
 * 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
 *
 *
 *
 * 进阶：
 *
 * 你能在线性时间复杂度内解决此题吗？
 *
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// 双向队列 递减
	de := newDeque(k + 1)
	ret := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for !de.empty() && nums[de.back()] <= nums[i] {
			de.popback()
		}
		de.pushBack(i)

		if de.front() <= i-k {
			de.popfront()
		}
		if i >= k-1 {
			ret = append(ret, nums[de.front()])
		}
	}
	return ret
}

type deque struct {
	data   []int
	l, cap int
}

func (d *deque) String() string {
	return fmt.Sprintf("data:%v, l:%d, cap:%d", d.data, d.l, d.cap)
}

func newDeque(cap int) *deque {
	return &deque{
		data: make([]int, cap),
		cap:  0,
	}
}

func (d *deque) popback() {
	if d.cap > 0 {
		d.cap--
	}
	if d.cap == 0 && d.l > 0 {
		d.data = make([]int, len(d.data))
		d.l = 0
	}
}

func (d *deque) popfront() {
	if d.cap > 0 {
		d.cap--
		d.l++
	}
}

func (d *deque) pushBack(n int) {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("%v, %v", err, d))
		}
	}()
	d.cap++
	d.data[d.l+d.cap-1] = n
}

func (d *deque) front() int {
	if d.cap > 0 {
		return d.data[d.l]
	}
	return 0
}

func (d *deque) back() int {
	if d.cap > 0 {
		return d.data[d.l+d.cap-1]
	}
	return 0
}

func (d *deque) empty() bool {
	if d.cap > 0 {
		return false
	}
	return true
}

func maxSlidingWindowV2(nums []int, k int) []int {
	ret := make([]int, 0, len(nums))
	total := 0
	for i := 0; i < len(nums) && i < k; i++ {
		total += nums[i]
	}
	ret = append(ret, total)
	for i := k; i < len(nums); i++ {
		total = total + nums[i] - nums[i-k]
		ret = append(ret, total)
	}
	return ret[:len(ret)]
}

// @lc code=end

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{}, 3))
	fmt.Println(maxSlidingWindow([]int{2}, 3))
	fmt.Println(maxSlidingWindow([]int{2}, 1))
	fmt.Println(maxSlidingWindow([]int{-95, 92, -85, 59, -59, -14, 88, -39, 2, 92, 94, 79, 78, -58, 37, 48, 63, -91, 91, 74, -28, 39, 90, -9, -72, -88, -72, 93, 38, 14, -83, -2, 21, 4, -75, -65, 3, 63, 100, 59, -48, 43, 35, -49, 48, -36, -64, -13, -7, -29, 87, 34, 56, -39, -5, -27, -28, 10, -57, 100, -43, -98, 19, -59, 78, -28, -91, 67, 41, -64, 76, 5, -58, -89, 83, 26, -7, -82, -32, -76, 86, 52, -6, 84, 20, 51, -86, 26, 46, 35, -23, 30, -51, 54, 19, 30, 27, 80, 45, 22}, 10))
}
