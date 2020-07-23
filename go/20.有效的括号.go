package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.85%)
 * Likes:    1413
 * Dislikes: 0
 * Total Accepted:    214.1K
 * Total Submissions: 522.6K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start
func isValid(str string) bool {
	parentheses := map[uint8]uint8{'(': ')', '{': '}', '[': ']'}
	leftIndex := make([]int, 0, len(str))
	for i := 0; i < len(str); i++ {
		if _, ok := parentheses[str[i]]; ok {
			leftIndex = append(leftIndex, i)
			continue
		}
		// 右括号，从 左括号里面拿一个出来比较
		if len(leftIndex) <= 0 {
			return false
		}
		if parentheses[str[leftIndex[len(leftIndex)-1]]] != str[i] {
			return false
		}
		leftIndex = leftIndex[:len(leftIndex)-1]
	}
	if len(leftIndex) != 0 {
		return false
	}
	return true
}

// 实现一个栈
func isValidV2(str string) bool {

	return true
}

// @lc code=end

func main() {
	fmt.Println(isValid("{}{}{}"))
}
