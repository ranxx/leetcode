package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(pre, next *ListNode) (*ListNode, *ListNode) {
	if next == nil {
		return pre, nil
	}
	pre.Next = next.Next
	next.Next = pre
	return next, pre
}

func swapPairs(head *ListNode) *ListNode {
	root := &ListNode{Next: head}
	if head == nil || head.Next == nil {
		return root.Next
	}

	head, tail := swap(head, head.Next)
	root.Next = head
	for tail != nil && tail.Next != nil && tail.Next.Next != nil {
		tmp := tail
		head, tail = swap(tail.Next, tail.Next.Next)
		tmp.Next = head
	}

	return root.Next
}
