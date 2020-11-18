/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */

// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcoode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next == nil {
		return slow
	}
	return slow.Next
}

// @lc code=end
