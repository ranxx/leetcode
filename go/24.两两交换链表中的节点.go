package main

import (
	"fmt"
	"time"
)

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (64.28%)
 * Likes:    417
 * Dislikes: 0
 * Total Accepted:    72K
 * Total Submissions: 111.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairsV3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, tmp := (*ListNode)(nil), head
	for tmp != nil && tmp.Next != nil {
		first, second := tmp, tmp.Next
		second.Next, first.Next = first, second.Next
		if pre == nil {
			head = second
		} else {
			pre.Next = second
		}
		pre, tmp = first, first.Next
		print(head)
		time.Sleep(time.Second)
	}
	return head
}

func swapPairsV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head, pre, frist := head.Next, (*ListNode)(nil), head
	for frist != nil && frist.Next != nil {
		tmp := frist.Next.Next
		frist.Next.Next = frist
		if pre != nil {
			pre.Next = frist.Next
		}
		pre = frist
		frist.Next = tmp
		frist = tmp
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	frist, second := head, head.Next
	frist.Next = swapPairs(second.Next)
	second.Next = frist
	return second
}

// @lc code=end
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	print(head)
	print(swapPairsV3(head))
	// print(swapPairsV2(swapPairs(head)))
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
	fmt.Println()
}
