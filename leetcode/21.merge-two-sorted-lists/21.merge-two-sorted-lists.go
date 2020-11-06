/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start

package leetcode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode

// using dummy node to create list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, &ListNode{}}
	for l1.Next != nil && l2.Next != nil {
		
	}

	return dummy.Next
}

// @lc code=end
