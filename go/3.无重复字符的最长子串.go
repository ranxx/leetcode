package main

import "fmt"

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (33.23%)
 * Likes:    3276
 * Dislikes: 0
 * Total Accepted:    380.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */

/*
 * @解题思路
 * 此题只考虑ASCII码，加上扩展一共为127*2，一般用255长度
 * 数组bits用-1填充，这样可以减少很多判断
 * 用begin记录每次处理字符串的开始下表，初始化为0
 * 循环字符串s，每个字符c的ASCII码值作为数组的下标i
 * 如果bits[i]>=begin，更新begin：begin=bits[i]+1；并计算此时的max= c在s中的下标 - begin
 * 更新bits[i] = c在s中的下标
 * 结束在计算下max
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	bits := make([]int, 255)
	for i := 0; i < 255; i++ {
		bits[i] = -1
	}
	begin, max := 0, 0
	for i := range s {
		index := bits[int(s[i])]
		if index >= begin {
			if max < i-begin {
				max = i - begin
			}
			begin = index + 1
		}
		bits[int(s[i])] = i
	}
	if len(s)-begin > max && len(s) > 0 {
		return len(s) - begin
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("cdd"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
