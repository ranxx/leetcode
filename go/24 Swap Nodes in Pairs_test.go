package main

import (
	"math/rand"
	"testing"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func generateListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func copySingleListNode(val ...int) []*ListNode {
	lists := make([]*ListNode, len(val))
	for i, v := range val {
		lists[i] = generateListNode(v)
	}
	return lists
}

func generateSwapPairs() []struct{ input, want []*ListNode } {
	rand.Seed(time.Now().Unix())
	rootCount := rand.Int()%50 + 50
	input, output := make([]*ListNode, 0, rootCount), make([]*ListNode, 0, rootCount)
	for i := 0; i < rootCount; i++ {
		pre, next := generateListNode(rand.Int()%100), generateListNode(rand.Int()%100)
		input[i], pre.Next = pre, next
		inputCurHead := next

		lists := copySingleListNode(next.Val, pre.Val)
		output[i], lists[0].Next = lists[0], lists[1]
		outputCurHead = lists[1]

		count := rand.Int() % 18
		for j := 0; j < count; j += 2 {
			pre, next = generateListNode(rand.Int()%100), generateListNode(rand.Int()%100)
			inputCurHead.Next, pre.Next = pre, next
			inputCurHead = next

			lists := copySingleListNode(next.Val, pre.Val)
			outputCurHead.Next, lists[0].Next = lists[0], lists[1]
			outputCurHead = lists[0]
		}
	}
	return []struct{ input, want []*ListNode }{input: input, output: output}
}

func judgeSwapPairs(l1, l2 *ListNode) bool {
	if nil == l1 && nil == l2 {
		return true
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func TestSwapPairs(t *testing.T) {
	tests := generateSwapPairs()
	for _, v := range tests {
		swapPairs(v.input)
	}
}

func BenchmarkSwapPairs(b *testing.B) {
	tests := generateSwapPairs()
	for _, v := range tests {
		swapPairs(v)
	}
}
