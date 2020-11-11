/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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

// sorted list, need a dummy node
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			tmp := curr.Next
			for tmp != nil && tmp.Val == curr.Next.Val {
				tmp = tmp.Next
			}
			curr.Next = tmp
			continue
		}
		curr = curr.Next
	}
	return dummy.Next
}

// @lc code=end
