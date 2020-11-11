/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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

// it is sorted list
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	curr := head
	for curr != nil && curr.Next != nil{
		if curr.Next.Val == curr.Val{
			curr.Next = curr.Next.Next
			continue;
		}
		curr = curr.Next
	}
	return head
}

// @lc code=end
