/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode

// find head and end, and reverse it
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	fast, slow := head, head
	dummy := &ListNode{0, head}
	revereHead := dummy
	for i := 0; i < m-1; i++ {
		revereHead = fast
		fast = fast.Next
	}
	for i := 0; i < n; i++ {
		slow = slow.Next
	}

	revereHead.Next = reverse(fast, slow)
	return dummy.Next
}

func reverse(first, last *ListNode) *ListNode {
	prev := last
	for first != last {
		next := first.Next
		first.Next = prev
		prev = first
		first = next
	}
	return prev
}

// @lc code=end
