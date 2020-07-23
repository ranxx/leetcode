package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (37.69%)
 * Likes:    4544
 * Dislikes: 0
 * Total Accepted:    467.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */

/*
 * @解题思路
 *
 * 根据题意可知：链表翻转表示一个数字
 * 实际上我们可以从两个链表的头往后面相加就行了，合大于等于10就向后进1
 * 最后如果哪个链表剩下那么全部赋值过来就行了
 */

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	tmp := root
	flag := false
	for l1 != nil || l2 != nil {
		num := 0
		if flag {
			flag, num = false, num+1
		}
		if l1 != nil {
			num, l1 = num+l1.Val, l1.Next
		}
		if l2 != nil {
			num, l2 = num+l2.Val, l2.Next
		}
		if num >= 10 {
			flag, num = true, num-10
		}
		tmp.Next = &ListNode{Val: num}
		tmp = tmp.Next
	}

	if flag {
		tmp.Next = &ListNode{Val: 1}
	}

	return root.Next
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}

func initListNode(a ...int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, v := range a {
		tmp.Next = &ListNode{
			Val: v,
		}
		tmp = tmp.Next
	}
	return head.Next
}

func (l *ListNode) String() string {
	ret := ""
	for tmp := l; tmp != nil; tmp = tmp.Next {
		ret = fmt.Sprintf("%v", tmp.Val) + ret
	}
	return ret
}

func main() {
	l1 := initListNode(9, 9, 9)
	l2 := initListNode(9, 9, 9)
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println("+")
	fmt.Println(l3)
}
