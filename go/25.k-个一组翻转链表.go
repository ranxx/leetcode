package main

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (56.18%)
 * Likes:    379
 * Dislikes: 0
 * Total Accepted:    38.2K
 * Total Submissions: 67.9K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。
 *
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 示例 :
 *
 * 给定这个链表：1->2->3->4->5
 *
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 *
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 * 说明 :
 *
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(begin, end *ListNode) *ListNode {
	pre, cur, next := (*ListNode)(nil), begin, begin
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	midhead := head
	for {
		end, i := midhead.Next, 0
		if midhead == head {
			end = head
		}
		for ; end != nil && i < k; i++ {
			end = end.Next
		}
		if end == nil && i != k {
			return head
		}
		// 这里就是移动到第k的位置上
		pre, cur, next := (*ListNode)(nil), midhead.Next, midhead.Next
		if midhead == head {
			cur, next = midhead, midhead
		}
		tmp := cur
		for cur != end {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		if midhead == head {
			head = pre
			//
		} else {
			midhead.Next = pre
		}
		midhead = tmp
		midhead.Next = end
	}
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	// 用来保存中间的头
	midhead := head
	for {
		tmp := midhead.Next
		if midhead == head {

			tmp = head
		}
		tail, i := tmp, 0
		for ; tail != nil && i < k; i++ {
			tail = tail.Next
		}
		// 如果是i不够，就退出去
		if i != k {
			return head
		}
		khead := reverse(tmp, tail)
		// 如果是第一个k，则更新head,否则更新连接未翻转的
		if midhead == head {
			head = khead
		} else {
			midhead.Next = khead
		}
		// 更新 midhead
		midhead = tmp
		midhead.Next = tail
	}
	return head
}

// @lc code=end
