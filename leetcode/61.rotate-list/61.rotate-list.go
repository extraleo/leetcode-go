/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
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


func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := k % getLen(head) // find the node
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	fast.Next = head
	newHead := slow.Next
	slow.Next = nil
	return newHead
}

func getLen(head *ListNode) int {
	len := 0
	for head != nil {
		head = head.Next
		len++
	}
	return len
}

// @lc code=end
